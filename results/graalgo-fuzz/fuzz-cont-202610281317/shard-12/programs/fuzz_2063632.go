package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint8) uint16 {
    for fi7 := 0; fi7 < 14; fi7++ {
        p0 -= p0
        p0 = 9223372036854775807
        _ = fi7
    }
    if true {
        for fi8 := 0; fi8 < 2; fi8++ {
            if (!false) {
                v9 := GenS1{F0: (fi8 << ((255) & 63)), F1: 42, F2: 376}
                for fi10 := 0; fi10 < 7; fi10++ {
                    v11 := []int{(fi8 * fi8)}
                    _ = fi10
                    _ = v11
                }
                _ = v9
            }
            p1 = p1
            _ = fi8
        }
        var v12 int64 = (p0 + p0)
        for fi13 := 0; fi13 < 2; fi13++ {
            switch "" {
            case "\n\t":
                v14 := GenS1{F0: 4294967295, F1: 0, F2: -907}
                _ = v14
            default:
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 uint = 65535
                    var v17 int64 = (p0 % ((-9223372036854775808) | 1))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = fi13
        }
        _ = v12
    } else {
        var v18 string = "foo bar"
        _ = v18
    }
    v19 := []int{}
    _ = p0
    _ = p1
    _ = v19
    return 786
}

func main() {
    v1 := []string{}
    var v2 int = 1
    v3 := make(map[string]int)
    var v4 bool = false
    v5 := make(map[string]bool)
    var v6 int64 = -1
    for fi20 := 0; fi20 < 16; fi20++ {
        v6 += v6
        var v21 int16 = -1
        for fi22 := 0; fi22 < 3; fi22++ {
            v23 := v5["foo bar"]
            v24 := GenS2{F0: 42}
            switch "\n\t" {
            case "a":
                var v25 uint64 = 0
                _ = v25
            default:
                if v23 {
                    v26 := v3["\n\t"]
                    _ = v26
                }
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi20
        _ = v21
    }
    v4 = (v4 && (true || (1 < 0)))
    for fi27 := 0; fi27 < 10; fi27++ {
        v5["hello"] = v4
        _ = fi27
    }
    var v28 int8 = 42
    var v29 int32 = 0
    v28 += -128
    v30 := GenS1{F0: (v2 + (v2 % ((-2147483648) | 1))), F1: -249, F2: v29}
    for fi31 := 0; fi31 < 15; fi31++ {
        v32 := GenS1{F0: fi31, F1: 69, F2: (v29 + (v29 ^ v29))}
        _ = fi31
        _ = v32
    }
    fmt.Println(v29, v4, v4)
    var v33 int = (v2 * (v2 >> ((65535) & 63)))
    var v34 float32 = 0.1
    fmt.Println(v2)
    _ = len(v3)
    var v35 int32 = (v29 - 0)
    var v36 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
