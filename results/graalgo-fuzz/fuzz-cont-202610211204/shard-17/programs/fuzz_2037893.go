package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint64 }
func genFn0(p0 int) uint16 {
    p0 += p0
    var v4 uint8 = 255
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{fi5}
        for fi7 := 0; fi7 < 19; fi7++ {
            v4 = 127
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return 0
}

func main() {
    var v1 uint16 = 877
    v2 := []uint32{965, 2147483648}
    v3 := make(map[uint32]bool)
    _ = len(v2)
    var v8 uint16 = v1
    var v9 uint64 = 4294967295
    var v10 int32 = -2147483648
    var v11 int64 = -9223372036854775808
    fmt.Println(v10, v1)
    var v12 int64 = (v11 + (v11 | v11))
    fmt.Println(v11)
    if false {
        if true {
            v8 = v8
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 uint = 2
                switch ("a" + ("foo bar" + "hello")) {
                case "a":
                    v14 += (v14 ^ v14)
                case "":
                    v11 -= v12
                default:
                    v1 = (v8 << ((v14) & 63))
                }
                switch fi13 {
                case -9223372036854775808:
                    v1 = 0
                case -161:
                    var v15 float64 = 1.5
                    _ = v15
                case 773:
                    for fi16 := 0; fi16 < 14; fi16++ {
                        if (-128 != -53) {
                            _ = len(v3)
                            var v17 uint32 = 0
                            _ = v17
                        } else {
                            v18 := v3[0]
                            fmt.Println(v14, v10)
                            _ = v18
                        }
                        _ = fi16
                    }
                }
                _ = fi13
                _ = v14
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                v10 -= v10
                var v20 string = ("foo bar" + "\n\t")
                v1 -= v8
                _ = fi19
                _ = v20
            }
        }
        var v21 float64 = -2.25
        var v22 float64 = ((1e10 + -2.25) + (v21 / ((v21 * ((v21 - v21) * 0.1)) - float64(0.0))))
        _ = v21
        _ = v22
    } else {
        v23 := GenS1{F0: (v11 ^ (v12 + (v11 * (^(v11))))), F1: (v9 + (v9 >> ((42) & 63))), F2: 1}
        _ = v23
    }
    if true {
        switch 972 {
        case 381:
            v11 = v11
        default:
            v11 += v12
        }
        var v24 bool = (32768 != 9223372036854775807)
        _ = v24
    }
    var v25 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v8
    _ = v9
}
