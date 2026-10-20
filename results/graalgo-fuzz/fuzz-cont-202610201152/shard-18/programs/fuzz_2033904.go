package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint; F1 uint16; F2 int64 }
func genFn0(p0 int) uint32 {
    p0 += p0
    var v7 int64 = 1
    var v8 uint64 = 870
    for fi9 := 0; fi9 < 16; fi9++ {
        if false {
            var v10 float32 = (3.14159 - -2.25)
            if (674 < 0) {
                var v11 int16 = 827
                _ = v11
            }
            _ = v10
        }
        var v12 uint = 695
        _ = fi9
        _ = v12
    }
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int = 127
    var v2 uint16 = 238
    v3 := []int{-73, (v1 ^ v1), (-(v1)), (v1 / ((v1) | 1))}
    var v4 int16 = 803
    var v5 uint64 = 4294967295
    var v6 string = ("" + "")
    if len(v3) > 0 {
        v3[0] = int(v2)
    }
    v2 = 126
    _ = len(v3)
    if false {
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 int32 = 2147483647
    if true {
        if false {
            v2 = v2
        } else {
            fmt.Println(v4, v6, v6)
            switch v6 {
            case "":
                fmt.Println(v4, v6)
            case "foo bar":
                var v14 int64 = -9223372036854775808
                _ = v14
            case "a":
                v15 := GenS1{F0: 2147483648, F1: (v5 ^ 952), F2: v4}
                _ = v15
            }
        }
        if false {
            switch (v1 * v1) {
            case 756:
                if len(v3) > 0 {
                    v3[0] = (v1 * (v1 ^ v1))
                }
            case -2147483648:
                if true {
                    var v16 float32 = ((((3.14159 * 3.14159) * -2.25) - -2.25) * 0.0)
                    _ = v16
                }
            }
            fmt.Println(v2)
            v13 -= -2147483648
        } else {
            if false {
                v17 := GenS1{F0: 65535, F1: (v5 * 496), F2: v4}
                _ = v17
            } else {
                v18 := GenS2{F0: 375, F1: (v2 | v2), F2: -816}
                _ = v18
            }
        }
    }
    v5 -= (v5 & v5)
    v6 = ("\n\t" + ((v6 + v6) + v6))
    if len(v3) > 0 {
        v3[0] = 1
    }
    v3 = append(v3, 2147483647)
    if true {
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 int32 = 0
            v3 = append(v3, (fi19 / ((fi19) | 1)))
            _ = fi19
            _ = v20
        }
    }
    v1 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
