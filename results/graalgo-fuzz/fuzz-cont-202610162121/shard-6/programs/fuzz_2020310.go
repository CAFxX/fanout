package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint; F1 uint; F2 uint64 }
func genFn0(p0 string) int {
    switch 32768 {
    case 9223372036854775807:
        fmt.Println(p0)
    case -256:
        v4 := GenS1{F0: 127, F1: 2, F2: -598}
        _ = v4
    default:
        var v5 string = "hello"
        _ = v5
    }
    v6 := GenS1{F0: 0, F1: 18446744073709551615, F2: -536}
    _ = p0
    _ = v6
    return 32767
}
func genFn1(p0 string) string {
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p0, fi7)
        _ = fi7
    }
    var v8 int32 = -2147483648
    for fi9 := 0; fi9 < 19; fi9++ {
        v8 += (v8 & (v8 | v8))
        var v10 uint16 = 65535
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 255
    _ = p0
    _ = v11
    _ = v8
    return p0
}

func main() {
    var v1 int32 = 1
    var v2 int8 = 0
    var v3 uint64 = 18446744073709551615
    var v12 bool = true
    var v13 string = "hello"
    fmt.Println(v13, v12, v2)
    v14 := []int{}
    var v15 uint32 = 996
    v15 = 161
    v16 := GenS2{F0: 90, F1: 0, F2: (v3 >> ((65535) & 63))}
    v15 -= 2147483648
    _ = len(v14)
    v12 = (2 > 235)
    var v17 float64 = -2.25
    var v18 int64 = -9223372036854775808
    for fi19 := 0; fi19 < 12; fi19++ {
        v3 = v3
        _ = fi19
    }
    if v12 {
        if (v2 <= 42) {
            for fi20 := 0; fi20 < 2; fi20++ {
                for fi21 := 0; fi21 < 2; fi21++ {
                    switch v13 {
                    case "hello":
                        fmt.Println(v12, fi21)
                    case "foo bar":
                        fmt.Println(v1, v13, v15)
                    default:
                        fmt.Println(v13, v15)
                    }
                    _ = fi21
                }
                var v22 uint32 = v15
                if (!(76 > uint8(v3))) {
                    v3 = (^(v3))
                    if (255 != 2) {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            v24 := GenS1{F0: 255, F1: (v3 << ((563) & 63)), F2: 0}
                            _ = fi23
                            _ = v24
                        }
                    } else {
                        for fi25 := 0; fi25 < 17; fi25++ {
                            v15 -= v15
                            _ = fi25
                        }
                        v26 := GenS2{F0: 1, F1: 252, F2: v3}
                        _ = v26
                    }
                    v1 += int32(v15)
                } else {
                    v27 := GenS2{F0: 65535, F1: 0, F2: v3}
                    v15 = (v22 >> ((309) & 63))
                    _ = v27
                }
                _ = fi20
                _ = v22
            }
        }
        var v28 uint64 = 4294967295
        _ = v28
    }
    v14 = append(v14, -509)
    var v29 int32 = v1
    var v30 uint32 = (v15 + v15)
    if len(v14) > 0 {
        v14[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
