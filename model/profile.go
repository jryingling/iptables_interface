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
			return Ruleset{}
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
