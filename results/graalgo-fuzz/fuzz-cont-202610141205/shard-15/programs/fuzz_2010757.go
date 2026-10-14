package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }
type GenS3 struct { F0 bool; F1 uint64; F2 uint }
func genFn0(p0 int16) uint32 {
    p0 -= (p0 + p0)
    switch -742 {
    case -310:
        fmt.Println(p0)
    case 635:
        if false {
            var v6 bool = true
            if v6 {
                fmt.Println(v6)
                v6 = (!v6)
                for fi7 := 0; fi7 < 20; fi7++ {
                    p0 -= p0
                    v8 := GenS1{F0: 1, F1: 18446744073709551615}
                    _ = fi7
                    _ = v8
                }
            }
            v9 := GenS2{F0: 65535, F1: 42, F2: 104}
            _ = v6
            _ = v9
        }
    }
    _ = p0
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 uint = 2
    var v3 uint32 = 344
    var v4 int32 = -1
    var v5 bool = (((v3 - (v3 | v3)) != 42) || true)
    v10 := GenS1{F0: v3, F1: 671}
    v11 := GenS1{F0: 2, F1: 18446744073709551615}
    v12 := []int{-226, int(v2), 9223372036854775807}
    fmt.Println(v2, v5, v5)
    if ((3.14159 - 1e10) <= 1e10) {
        v4 = 42
    }
    if (100 < 127) {
        var v13 string = ""
        if true {
            v13 = v13
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS2{F0: (^(v1)), F1: -1, F2: (v2 - v2)}
                _ = fi14
                _ = v15
            }
        }
        _ = v13
    }
    v2 -= 0
    var v16 uint16 = v1
    if (0.1 < 3.14159) {
        if false {
            v1 += 2
            v3 = 332
            if true {
                switch 2147483647 {
                case 32768:
                    if len(v12) > 0 {
                        v12[0] = 924
                    }
                case -9223372036854775808:
                    var v17 int32 = 42
                    _ = v17
                }
                fmt.Println(v16, v2, v5)
            } else {
                for fi18 := 0; fi18 < 1; fi18++ {
                    fmt.Println(v2)
                    v4 -= (v4 / (((^(v4))) | 1))
                    for fi19 := 0; fi19 < 16; fi19++ {
                        v1 = (v1 * 32767)
                        _ = fi19
                    }
                    _ = fi18
                }
            }
        } else {
            fmt.Println(v3, v2)
            fmt.Println(v3)
            fmt.Println(v1, v16, v1)
        }
        v20 := GenS1{F0: v3, F1: 0}
        if false {
            switch "\n\t" {
            case "foo bar":
                v21 := GenS3{F0: v5, F1: 42, F2: (v2 | v2)}
                _ = v21
            case "hello":
                fmt.Println(v4, v2, v1)
            default:
                v3 = (v3 % ((v3) | 1))
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(v1, v2)
                v4 = v4
                var v23 uint32 = genFn0(-32768)
                _ = fi22
                _ = v23
            }
            v2 += (v2 << ((v2) & 63))
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
