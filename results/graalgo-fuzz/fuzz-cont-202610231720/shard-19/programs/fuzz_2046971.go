package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int64 = 847
    v2 := []string{"foo bar"}
    var v3 uint16 = 580
    v4 := make(map[string]int)
    var v5 int16 = 1000
    switch 9223372036854775807 {
    case 812:
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 bool = false
            var v8 string = "hello"
            v2 = append(v2, v8)
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 4; fi9++ {
            if (((((!(!false)) || (false || (!(42 >= (v1 & (^(v1))))))) && ((true || true) && (v3 != (v3 >> ((65535) & 63))))) || false) || ((fi9 * 9223372036854775807) != fi9)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 = (v1 - int64(fi9))
                var v10 uint = 847
                _ = v10
            }
            if false {
                for fi11 := 0; fi11 < 9; fi11++ {
                    v12 := GenS1{F0: (949 == 0)}
                    _ = len(v4)
                    v5 = (v5 - (v5 % (((v5 / ((v5) | 1))) | 1)))
                    _ = fi11
                    _ = v12
                }
            }
            _ = fi9
        }
    }
    delete(v4, "hello")
    var v13 uint32 = 2
    switch "foo bar" {
    case "foo bar":
        if (0.1 == (1e10 - -2.25)) {
            var v14 uint32 = v13
            v3 = v3
            switch "\n\t" {
            case "\n\t":
                delete(v4, (("foo bar" + "foo bar") + ("hello" + "hello")))
            }
            _ = v14
        }
    case "\n\t":
        var v15 int64 = (v1 << ((347) & 63))
        _ = v15
    }
    fmt.Println(v1, v5)
    v16 := GenS1{F0: ((526 > -477) && true)}
    var v17 float64 = (3.14159 + 0.1)
    fmt.Println(v17)
    var v18 int16 = (v5 >> ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
