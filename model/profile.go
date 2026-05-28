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
router pointer to routerconfig object, for required router wan, lan, subnet info
optional by passing nil for other presets
*/
func DefaultRuleset(preset ProfileType, router *RouterConfig) Ruleset {
	switch preset {
		case Desktop:
			return Ruleset{
				RulesetName: "Desktop",
				Input: Chain{
					ChainName: INPUT,
					Policy: DROP,
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
			return Ruleset{
				RulesetName: "Server",
				Input: Chain{
					ChainName: INPUT,
					Policy: DROP,
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
						{
							Comment: "Allow SSH",
							Protocol: TCP,
							DstPort: "22",
							Target: ACCEPT,
						},
						{
							Comment: "Allow HTTP",
							Protocol: TCP,
							DstPort: "80",
							Target: ACCEPT,
						},
						{
							Comment: "Allow HTTPS",
							Protocol: TCP,
							DstPort: "443",
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
		case Router:
			var wanIf, lanIf string
			if router == nil {
				wanIf = "eth0"
				lanIf = "eth1"
			} else {
				wanIf = router.WanIF
				lanIf = router.LanIF
			}
			return Ruleset{
				RulesetName: "Router",
				Input: Chain{
					ChainName: INPUT,
					Policy: DROP,
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
						{
							Comment: "Allow SSH",
							Protocol: TCP,
							DstPort: "22",
							InIface: wanIf,
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
					Rules: []Rule{
						{
							Comment: "Allow established/related connections",
							ConntrackState: "ESTABLISHED,RELATED",
							Target: ACCEPT,
							
						},
						{
							Comment: "LAN > WAN",
							InIface: lanIf,
							OutIface: wanIf,
							Target: ACCEPT,
						},
					},
				},
			}
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
