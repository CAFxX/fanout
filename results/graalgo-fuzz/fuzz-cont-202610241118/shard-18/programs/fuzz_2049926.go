package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0() uint32 {
    switch "\n\t" {
    case "a":
        var v5 uint32 = 1
        _ = v5
    case "":
        switch 214 {
        case -9223372036854775808:
            var v6 string = ""
            _ = v6
        default:
            if (((!(!true)) || true) || (!(979 == 255))) {
                var v7 bool = (true && false)
                v8 := GenS1{F0: 42}
                _ = v7
                _ = v8
            } else {
                v9 := GenS2{F0: 32767}
                _ = v9
            }
        }
    }
    v10 := GenS1{F0: 1000}
    var v11 uint16 = 0
    _ = v10
    _ = v11
    return 42
}
func genFn1(p0 uint16) uint8 {
    for fi12 := 0; fi12 < 12; fi12++ {
        p0 -= (p0 * (p0 / ((p0) | 1)))
        _ = fi12
    }
    p0 = 864
    var v13 uint16 = (p0 + 331)
    _ = p0
    _ = v13
    return 42
}

func main() {
    var v1 uint = 563
    v2 := []int{4294967295, -2147483648}
    var v3 float64 = 3.14159
    v4 := make(map[uint32]uint32)
    var v14 uint = (v1 / ((v1) | 1))
    if ((v3 == 1.5) || ((32767 < -310) && (false && (1 == 401)))) {
        fmt.Println(v1)
    } else {
        v14 = (v14 & (v1 | v1))
        v15 := GenS1{F0: 32767}
        _ = v15
    }
    fmt.Println(v14, v1, v3)
    switch 423 {
    case 9223372036854775807:
        for fi16 := 0; fi16 < 19; fi16++ {
            for fi17 := 0; fi17 < 9; fi17++ {
                v4[65535] = 2147483647
                var v18 uint16 = 0
                v19 := GenS3{F0: 211}
                _ = len(v4)
                v1 += 65535
                switch fi17 {
                case -923:
                    v20 := GenS3{F0: 1}
                    v14 -= (v1 * (v14 & (v14 ^ (v1 | v1))))
                    _ = v20
                case 32767:
                    _ = len(v4)
                case 128:
                    _ = len(v2)
                default:
                    var v21 float32 = 1.5
                    _ = v21
                }
                v3 = v3
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
    }
    var v22 int32 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, -9223372036854775808)
    _ = len(v2)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
