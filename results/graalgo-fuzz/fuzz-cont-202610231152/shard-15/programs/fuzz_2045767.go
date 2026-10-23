package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint8 }
func genFn0() int {
    switch ("foo bar" + ("hello" + "")) {
    case "\n\t":
        switch 9223372036854775807 {
        case 127:
            v4 := []int{32768, 32767, -182, -1}
            _ = v4
        default:
            v5 := []int{}
            _ = v5
        }
    case "hello":
        v6 := GenS1{F0: 0, F1: 197}
        _ = v6
    }
    var v7 uint16 = 16
    v8 := []int{127, 145, 32767, 161}
    var v9 int32 = -288
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = 0.1
    var v3 uint32 = 42
    v10 := GenS1{F0: 776, F1: 2}
    v11 := []int{127, 646, 9223372036854775807}
    for fi12 := 0; fi12 < 1; fi12++ {
        if (127 <= 42) {
            switch fi12 {
            case 63:
                if len(v11) > 0 {
                    v11[0] = 589
                }
            case 440:
                if (4294967295 < 804) {
                    fmt.Println(v1)
                    fmt.Println(v3, v2)
                    v1 += v1
                }
            }
            fmt.Println(fi12)
            if true {
                v13 := GenS2{F0: 2147483647, F1: (-112 > 42), F2: 255}
                v3 = 394
                switch (fi12 ^ (fi12 ^ fi12)) {
                case 42:
                    var v14 bool = true
                    _ = v14
                }
                _ = v13
            }
        } else {
            if len(v11) > 0 {
                v11[0] = (fi12 % (((fi12 | fi12)) | 1))
            }
        }
        v15 := GenS1{F0: 18446744073709551615, F1: 176}
        _ = fi12
        _ = v15
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v11 = append(v11, fi16)
        _ = len(v11)
        _ = fi16
    }
    if len(v11) > 0 {
        v11[0] = genFn0()
    }
    v17 := GenS1{F0: 693, F1: 240}
    v2 += v2
    if ((225 != 255) || false) {
        v11 = append(v11, genFn0())
        var v18 uint64 = v1
        _ = v18
    } else {
        if len(v11) > 0 {
            v11[0] = 9223372036854775807
        }
        for fi19 := 0; fi19 < 17; fi19++ {
            v11 = append(v11, (fi19 - fi19))
            _ = fi19
        }
    }
    v1 += 0
    var v20 int64 = -1
    fmt.Println(v1)
    v20 -= (-(v20))
    v1 = (v1 ^ uint64(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v20
    _ = v3
}
