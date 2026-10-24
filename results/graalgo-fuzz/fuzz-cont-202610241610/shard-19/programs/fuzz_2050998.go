package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int8, p1 uint64) int64 {
    switch ("hello" + ("a" + "hello")) {
    case "\n\t":
        var v4 uint32 = 753
        _ = v4
    case "hello":
        v5 := []int{-2147483648}
        _ = v5
    }
    switch -9223372036854775808 {
    case 42:
        var v6 uint32 = 896
        _ = v6
    case -1:
        var v7 uint16 = 42
        _ = v7
    case 127:
        var v8 float32 = 0.1
        _ = v8
    default:
        var v9 float64 = 1.5
        _ = v9
    }
    fmt.Println(p1, p0)
    p0 = (p0 * p0)
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    var v1 float32 = (1.5 * -2.25)
    v2 := make(map[string]int)
    var v3 int32 = 0
    v1 = (1e10 / -2.25)
    v10 := GenS1{F0: 1, F1: 65535, F2: 32767}
    fmt.Println(v3, v1)
    v11 := []int{-1}
    for fi12 := 0; fi12 < 20; fi12++ {
        if true {
            v13 := GenS2{F0: 1000, F1: 9223372036854775807}
            _ = v13
        } else {
            v1 += ((v1 - (((((v1 / v1) / -2.25) / v1) - v1) - 1e10)) * float32(0.0))
            switch fi12 {
            case 996:
                v14 := v2["foo bar"]
                _ = v14
            case 127:
                switch fi12 {
                case 13:
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                case -750:
                    v1 = (1e10 * (((float32(((1.5 + ((-2.25 / 1e10) / (-2.25 - 1e10))) + 3.14159)) + 1e10) * v1) / (v1 / -2.25)))
                case 182:
                    v3 = 42
                default:
                    var v15 uint = 18446744073709551615
                    _ = v15
                }
            case 166:
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint32 = 4294967295
                    switch int(v17) {
                    case -926:
                        var v18 int64 = -488
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
            default:
                v19 := v2[(("foo bar" + "\n\t") + "hello")]
                _ = v19
            }
        }
        v1 = v1
        var v20 int64 = -1
        _ = fi12
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        switch 127 {
        case 1:
            v22 := GenS3{F0: (v3 % (((^(v3))) | 1))}
            _ = v22
        default:
            if false {
                if true {
                    fmt.Println(v1)
                    v3 += 0
                }
                v23 := v2["\n\t"]
                v24 := GenS2{F0: -958, F1: genFn0(0, 744)}
                _ = v23
                _ = v24
            }
        }
        _ = len(v11)
        _ = fi21
    }
    v3 += (v3 % ((v3) | 1))
    delete(v2, "")
    var v25 float64 = (float64(v1) + (((0.1 + 0.0) + 0.1) * 0.1))
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 int32 = (v3 / ((v3) | 1))
        _ = fi26
        _ = v27
    }
    v25 -= 1e10
    var v28 uint8 = 2
    var v29 uint8 = v28
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
}
