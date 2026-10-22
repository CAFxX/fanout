package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 uint) bool {
    var v4 int32 = 42
    v5 := GenS1{F0: (p0 / ((249) | 1))}
    _ = p0
    _ = v4
    _ = v5
    return (644 <= 137)
}
func genFn1(p0 uint8) uint {
    v6 := GenS1{F0: 65535}
    v7 := []int{433}
    _ = p0
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint32 = 974
    var v2 int = 9223372036854775807
    v3 := []int{v2, (^(v2)), v2}
    v1 = 0
    if ((v1 ^ v1) >= 2147483647) {
        for fi8 := 0; fi8 < 18; fi8++ {
            v1 -= (v1 * (v1 % (((v1 + v1)) | 1)))
            v9 := GenS2{F0: v1}
            _ = fi8
            _ = v9
        }
    }
    var v10 string = ("\n\t" + "hello")
    if false {
        v1 -= (v1 * v1)
        switch v10 {
        case "hello":
            if len(v3) > 0 {
                v3[0] = (v2 / ((v2) | 1))
            }
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "a":
            var v11 float64 = (float64(1e10) + 1.5)
            _ = v11
        default:
            for fi12 := 0; fi12 < 10; fi12++ {
                v2 += v2
                switch "hello" {
                case "a":
                    fmt.Println(fi12, v10)
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "foo bar":
                    v13 := GenS2{F0: (v1 >> ((993) & 63))}
                    _ = v13
                default:
                    v14 := GenS1{F0: genFn1(128)}
                    _ = v14
                }
                switch 128 {
                case 9223372036854775807:
                    v1 = v1
                default:
                    if len(v3) > 0 {
                        v3[0] = fi12
                    }
                }
                _ = fi12
            }
        }
        fmt.Println(v2, v1, v10)
    } else {
        if len(v3) > 0 {
            v3[0] = -173
        }
    }
    v15 := GenS3{F0: "hello", F1: 0}
    v16 := GenS3{F0: v10, F1: 42}
    v17 := GenS2{F0: (^(v1))}
    v1 += 4294967295
    if (128 <= 179) {
        for fi18 := 0; fi18 < 12; fi18++ {
            if len(v3) > 0 {
                v3[0] = v2
            }
            v19 := GenS3{F0: "a", F1: -1}
            if true {
                var v20 uint64 = 42
                _ = v20
            } else {
                v3 = append(v3, v2)
            }
            _ = fi18
            _ = v19
        }
    } else {
        var v21 uint16 = 656
        var v22 float32 = -2.25
        _ = len(v3)
        _ = v21
        _ = v22
    }
    var v23 uint32 = (v1 << ((65535) & 63))
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
}
