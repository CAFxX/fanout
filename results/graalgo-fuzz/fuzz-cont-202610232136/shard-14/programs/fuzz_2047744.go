package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint16, p1 int, p2 string) int16 {
    p2 = (p2 + "a")
    if (false && false) {
        p0 -= p0
    } else {
        v4 := GenS1{F0: 4294967295}
        v5 := GenS1{F0: 1}
        _ = v4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint64 = 1
    v2 := []string{"", "a", ("" + (("hello" + "\n\t") + "hello"))}
    v3 := []int{}
    if len(v3) > 0 {
        v3[0] = 127
    }
    fmt.Println(v1)
    if true {
        switch -2147483648 {
        case 42:
            var v6 uint64 = v1
            _ = v6
        default:
            _ = len(v3)
        }
    }
    if (-128 != -128) {
        for fi7 := 0; fi7 < 6; fi7++ {
            switch ("" + ("" + "hello")) {
            case "\n\t":
                fmt.Println(v1, v1)
            case "a":
                v1 = v1
            case "":
                v1 -= v1
            default:
                switch 551 {
                case 32767:
                    if false {
                        for fi8 := 0; fi8 < 3; fi8++ {
                            v1 = 0
                            _ = fi8
                        }
                    } else {
                        for fi9 := 0; fi9 < 16; fi9++ {
                            var v10 string = "a"
                            _ = fi9
                            _ = v10
                        }
                        for fi11 := 0; fi11 < 18; fi11++ {
                            v1 = 649
                            v1 = v1
                            _ = fi11
                        }
                        var v12 bool = false
                        _ = v12
                    }
                case 9223372036854775807:
                    for fi13 := 0; fi13 < 17; fi13++ {
                        fmt.Println(fi7, v1, fi7)
                        fmt.Println(fi13)
                        _ = fi13
                    }
                case 2147483647:
                    v1 -= (v1 & v1)
                default:
                    var v14 string = "foo bar"
                    _ = v14
                }
            }
            v1 -= (v1 << ((uint(v1)) & 63))
            _ = fi7
        }
        fmt.Println(v1)
    }
    v1 += v1
    v15 := GenS1{F0: 42}
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 int64 = -54
        v1 = (v1 & (v1 ^ 4294967295))
        v1 = (v1 - v1)
        _ = fi16
        _ = v17
    }
    fmt.Println(v1)
    v1 += (^(v1))
    if len(v3) > 0 {
        v3[0] = 856
    }
    switch 9223372036854775807 {
    case -2147483648:
        v3 = append(v3, -1)
    case 32768:
        var v18 int = -665
        _ = v18
    case 127:
        var v19 int8 = 5
        _ = v19
    }
    if (false || false) {
        var v20 uint64 = v1
        v2 = append(v2, "")
        _ = v20
    } else {
        v1 = v1
        var v21 bool = (0.1 >= -2.25)
        v21 = (!((-111 > -75) || (true || v21)))
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
}
