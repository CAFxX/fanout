package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
func genFn0(p0 uint64) int64 {
    var v5 uint64 = 1
    v6 := []int{}
    if (918 == 521) {
        if len(v6) > 0 {
            v6[0] = int(p0)
        }
    } else {
        switch -1 {
        case 9223372036854775807:
            fmt.Println(v5, v5)
        default:
            switch 32767 {
            case -169:
                var v7 uint16 = 255
                _ = v7
            case -412:
                if (!((945 <= 42) && false)) {
                    _ = len(v6)
                    fmt.Println(p0, v5)
                } else {
                    fmt.Println(v5, p0)
                    var v8 int16 = 32767
                    _ = v8
                }
            case 32767:
                if (68 == 1) {
                    p0 = (v5 ^ (v5 + (^(p0))))
                }
            default:
                v9 := GenS1{F0: 248, F1: 2147483647, F2: ("hello" + "a")}
                _ = v9
            }
        }
        p0 += (p0 | 1)
        var v10 bool = (!true)
        _ = v10
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 42
    var v3 float64 = ((float64((float32((((1.5 - 0.1) * 0.0) + (((-2.25 - (((((0.1 / float64(3.14159)) - ((3.14159 - 1e10) + float64((1.5 / ((1.5 * ((1e10 - (0.0 / 1.5)) - 0.0)) / ((-2.25 - (3.14159 - (3.14159 + (0.1 - (float32(float64(((0.1 * 1.5) - 1e10))) * 0.0))))) - 0.1)))))) - 3.14159) * 0.0) - 3.14159)) - -2.25) - 1.5))) / 0.1)) + (1e10 * 0.1)) * 0.1)
    var v4 uint16 = 943
    v2 = v2
    if (true && (uint32(v4) > v2)) {
        if true {
            switch "hello" {
            case "\n\t":
                fmt.Println(v2)
            case "a":
                var v11 int32 = 742
                _ = v11
            }
        } else {
            var v12 bool = false
            _ = v12
        }
    } else {
        var v13 uint16 = 1
        fmt.Println(v2, v2, v4)
        _ = v13
    }
    var v14 uint8 = 2
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := v1[(fi15 % ((fi15) | 1))]
        _ = fi15
        _ = v16
    }
    v4 = (v4 | (v4 + (^(v4))))
    var v17 float32 = float32(((1.5 - (0.0 / (3.14159 + v3))) / ((v3 + v3) * v3)))
    v18 := GenS1{F0: v4, F1: 0, F2: (("foo bar" + "\n\t") + ("foo bar" + ""))}
    if true {
        if (!((true || true) && (0 != -187))) {
            switch "" {
            case "hello":
                _ = len(v1)
            }
        } else {
            if false {
                v4 -= (v4 ^ (v4 & (v4 % (((^(v4))) | 1))))
                if true {
                    var v19 bool = false
                    _ = v19
                } else {
                    var v20 uint64 = 4294967295
                    fmt.Println(v2)
                    _ = v20
                }
            }
            v21 := []int{-2147483648}
            switch 32768 {
            case -643:
                v14 = v14
            }
            _ = v21
        }
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 uint32 = v2
            _ = fi22
            _ = v23
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
