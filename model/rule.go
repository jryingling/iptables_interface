package model

type Rule struct {
    Comment string
    Protocol Protocol
    SrcIP, DstIP string
    SrcPort, DstPort string
    InIface, OutIface string
    ConntrackState string
    Target Target
}

type Chain struct {
    ChainName ChainName
    Policy Target // with packets that dont match any rule
    Rules []Rule
}

// rules live in chains
// chains live in ruleset

type Ruleset struct {
    RulesetName string
    // dont store as slice, easier to access this way
    // chains are fixed
    Input Chain
    Output Chain
    Forward Chain
}

/* 
custom string type for:
- type safety
    - always passing a Protocol, not just any string
- constants
- custom methods later on
*/
type Protocol string 

const (
    TCP Protocol = "tcp"
    UDP Protocol = "udp"
    ICMP Protocol = "icmp"
    ALL Protocol = "all"
)
    
type Target string

const (
    ACCEPT Target = "ACCEPT"
    DROP Target = "DROP"
    REJECT Target = "REJECT"
    LOG Target = "LOG"
    RETURN Target = "RETURN"
    QUEUE Target = "QUEUE"
)

type ChainName string

const (
    INPUT ChainName = "INPUT"
    OUTPUT ChainName = "OUTPUT"
    FORWARD ChainName = "FORWARD"
)
