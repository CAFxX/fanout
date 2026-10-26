package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 int64 = -128
    var v4 uint64 = 4294967295
    var v5 bool = false
    v5 = (!true)
    v6 := GenS1{F0: 9223372036854775807}
    v7 := GenS1{F0: -9223372036854775808}
    switch ("a" + "a") {
    case "foo bar":
        v8 := GenS1{F0: 9223372036854775807}
        _ = v8
    case "a":
        v9 := []int{-460}
        _ = v9
    default:
        for fi10 := 0; fi10 < 12; fi10++ {
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 string = "\n\t"
                switch -2147483648 {
                case -2147483648:
                    if (3.14159 < 0.0) {
                        var v13 uint64 = (v4 - v4)
                        v4 += 2
                        v5 = (!(fi10 > fi11))
                        v14 := GenS1{F0: (fi10 & fi10)}
                        _ = v13
                        _ = v14
                    }
                case 964:
                    v5 = (1e10 < float64(1e10))
                case 9223372036854775807:
                    v12 = (("" + ((v12 + v12) + (v12 + v12))) + v12)
                }
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    v15 := GenS1{F0: -9223372036854775808}
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 float32 = 1e10
        v4 = (v4 << ((2) & 63))
        v18 := GenS1{F0: (fi16 * (fi16 ^ 4294967295))}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: -1}
    v1["foo bar"] = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
