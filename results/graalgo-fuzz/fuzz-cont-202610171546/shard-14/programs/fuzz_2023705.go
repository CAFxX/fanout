package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }

func main() {
    v1 := []uint32{}
    var v2 uint64 = 1
    var v3 string = "foo bar"
    v4 := make(map[string]uint32)
    var v5 int32 = -1
    var v6 string = v3
    v2 += 2
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int64 = 9223372036854775807
        switch v6 {
        case "foo bar":
            if (-507 == 42) {
                switch -781 {
                case 691:
                    var v9 int64 = (-(v8))
                    _ = v9
                }
                v2 = v2
                for fi10 := 0; fi10 < 5; fi10++ {
                    v11 := GenS1{F0: (fi7 - fi7), F1: 46}
                    if (!true) {
                        v5 = (v5 ^ (v5 >> ((4294967295) & 63)))
                    } else {
                        v5 += (v5 >> ((18446744073709551615) & 63))
                        fmt.Println(v6, v3)
                    }
                    if true {
                        var v12 uint64 = v2
                        _ = v12
                    }
                    _ = fi10
                    _ = v11
                }
            }
        }
        _ = fi7
        _ = v8
    }
    var v13 string = ((v3 + v3) + v6)
    fmt.Println(v13, v5)
    v2 = (v2 * (v2 & 0))
    v13 = v6
    v14 := GenS1{F0: 910, F1: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
