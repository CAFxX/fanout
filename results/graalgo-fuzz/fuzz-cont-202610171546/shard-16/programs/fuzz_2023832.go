package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int16; F2 int16 }

func main() {
    v1 := []string{("\n\t" + ("\n\t" + "foo bar")), ""}
    var v2 int16 = 1000
    v3 := make(map[int]bool)
    v4 := []string{"\n\t"}
    for fi5 := 0; fi5 < 18; fi5++ {
        v2 -= -436
        _ = fi5
    }
    v2 = (v2 ^ -1)
    v6 := v3[-2147483648]
    switch "hello" {
    case "":
        var v7 uint8 = 73
        _ = v7
    default:
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := v3[32767]
            if len(v1) > 0 {
                v1[0] = ""
            }
            _ = fi8
            _ = v9
        }
    }
    switch 571 {
    case 32768:
        for fi10 := 0; fi10 < 3; fi10++ {
            if v6 {
                if len(v4) > 0 {
                    v4[0] = ""
                }
            }
            var v11 uint = 42
            if (!(1 <= 9223372036854775807)) {
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 string = "hello"
                    _ = fi12
                    _ = v13
                }
            }
            _ = fi10
            _ = v11
        }
    case -68:
        v14 := GenS1{F0: 529}
        _ = v14
    }
    v6 = v6
    v15 := GenS1{F0: 751}
    v16 := GenS2{F0: 18446744073709551615, F1: 1000, F2: v2}
    fmt.Println(v6)
    v2 = 1
    var v17 uint8 = 165
    v6 = (1e10 <= 1e10)
    _ = len(v3)
    switch -135 {
    case -2147483648:
        v18 := GenS1{F0: 523}
        _ = v18
    default:
        v6 = v6
    }
    var v19 uint64 = 130
    var v20 int64 = -9223372036854775808
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
}
