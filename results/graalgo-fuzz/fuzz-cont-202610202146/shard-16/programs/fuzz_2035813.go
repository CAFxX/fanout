package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint8; F1 int8; F2 string }
func genFn0(p0 int64, p1 bool, p2 string) int64 {
    p2 = p2
    var v5 uint32 = 2147483647
    v6 := []int{-101}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 string = "hello"
    v2 := []string{("hello" + (v1 + v1))}
    var v3 int32 = 2147483647
    var v4 float32 = 1.5
    v7 := GenS1{F0: 2, F1: 1}
    var v8 bool = false
    v1 = (v1 + "")
    fmt.Println(v4, v3, v1)
    var v9 int32 = v3
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int = (fi10 % ((fi10) | 1))
        var v12 int8 = -128
        v13 := GenS1{F0: 937, F1: 1}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    switch 2147483647 {
    case 0:
        switch -276 {
        case 0:
            if len(v2) > 0 {
                v2[0] = (v1 + ((v1 + "\n\t") + "hello"))
            }
        case -9223372036854775808:
            var v14 uint8 = 42
            _ = v14
        default:
            for fi15 := 0; fi15 < 17; fi15++ {
                fmt.Println(fi15, v8, v4)
                if len(v2) > 0 {
                    v2[0] = v1
                }
                _ = fi15
            }
        }
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        switch fi16 {
        case 1:
            switch (-(fi16)) {
            case 9223372036854775807:
                var v17 uint8 = 20
                _ = v17
            case 66:
                if (v1 > v1) {
                    v9 += v9
                    v3 = -1
                    fmt.Println(v8, v3)
                }
            case -734:
                v8 = (65 <= 42)
            }
        case -9223372036854775808:
            switch 32767 {
            case 163:
                v2 = append(v2, "")
            case -1:
                var v18 uint16 = 0
                _ = v18
            default:
                for fi19 := 0; fi19 < 10; fi19++ {
                    if (128 < 127) {
                        v8 = v8
                        var v20 string = (v1 + v1)
                        v8 = (false && v8)
                        _ = v20
                    }
                    v4 += (v4 + ((v4 * v4) - v4))
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    _ = fi19
                }
            }
        case -135:
            v21 := GenS1{F0: 4294967295, F1: 9223372036854775807}
            _ = v21
        }
        switch (fi16 << ((42) & 63)) {
        case -9223372036854775808:
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 float64 = -2.25
                v24 := GenS2{F0: 0, F1: 100, F2: ("a" + v1)}
                var v25 bool = false
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
        default:
            v26 := GenS2{F0: 57, F1: -1, F2: v1}
            _ = v26
        }
        _ = fi16
    }
    v3 += -2147483648
    v4 -= 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
