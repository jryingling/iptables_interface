package model

type ProfileType int

const (
	Desktop ProfileType = iota // start at 0, inc 1 for each constant
	Server
	Router
	Custom
)

/*
preset choice gives us ruleset for that preset
*/
func DefaultRuleset(preset ProfileType) Ruleset {
	switch preset {
		case Desktop:
			return Ruleset{
				RulesetName: "Desktop",
				Input: Chain{
					ChainName: INPUT,
					Policy: DROP,
					// conntrack, interfaces, protocols (order matters)
					Rules: []Rule{
						{
							Comment: "Allow established/related connections",
							ConntrackState: "ESTABLISHED,RELATED",
							Target: ACCEPT,
							
						},
						{
							Comment: "Allow loopback traffic on Host",
							InIface: "lo",
							Target: ACCEPT,
						},
						{
							Comment: "Allow ICMP",
							Protocol: ICMP,
							Target: ACCEPT,
						},
					},
				},
				Output: Chain {
					ChainName: OUTPUT,
					Policy: ACCEPT,
					Rules: []Rule{},
				},
				Forward: Chain {
					ChainName: FORWARD,
					Policy: DROP,
					Rules: []Rule{},
				},
			}
		case Server:
			return Ruleset{}
		case Router:
			return Ruleset{}
		case Custom:
			/*
				empty ruleset with 
					INPUT and FORWARD policy set to accept
					Output accept too, by default for all
			*/
			return Ruleset{
				RulesetName: "Custom",
				Input: Chain{
					ChainName: INPUT,
					Policy: ACCEPT,
					Rules: []Rule{},
				},
				Output: Chain {
					ChainName: OUTPUT,
					Policy: ACCEPT,
					Rules: []Rule{},
				},
				Forward: Chain {
					ChainName: FORWARD,
					Policy: ACCEPT,
					Rules: []Rule{},
				},
			}
		default:
			return Ruleset{} // return empty ruleset struct/type
	}
}
