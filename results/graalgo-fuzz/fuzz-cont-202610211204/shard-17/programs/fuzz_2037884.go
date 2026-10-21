package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint; F1 int16; F2 uint16 }
type GenS3 struct { F0 int; F1 string; F2 int64 }
func genFn0(p0 int8) uint16 {
    v7 := GenS1{F0: 18446744073709551615, F1: 128}
    var v8 int64 = -1
    if false {
        for fi9 := 0; fi9 < 6; fi9++ {
            if (true && true) {
                p0 = (p0 << ((18446744073709551615) & 63))
                v10 := []int{(fi9 * (fi9 & fi9))}
                p0 -= 91
                _ = v10
            } else {
                for fi11 := 0; fi11 < 3; fi11++ {
                    v8 = (v8 % ((v8) | 1))
                    var v12 int = fi11
                    var v13 uint64 = 0
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v14 := []int{fi9, (-(fi9)), 2147483647, (fi9 >> ((379) & 63))}
                if (1 != 528) {
                    var v15 int16 = -631
                    v8 -= (v8 ^ (v8 << ((1) & 63)))
                    _ = v15
                } else {
                    if len(v14) > 0 {
                        v14[0] = fi9
                    }
                    if true {
                        var v16 int32 = 0
                        _ = v16
                    } else {
                        var v17 int = fi9
                        var v18 string = ("foo bar" + "foo bar")
                        if (0 > 624) {
                            var v19 float32 = ((1.5 / ((((3.14159 * 3.14159) / (((((1e10 / ((-2.25 - ((-2.25 / ((((float32(1e10) / 0.1) * -2.25) - 1.5) + 0.0)) * (3.14159 / 1.5))) - 1e10)) + 0.0) - 1e10) - (1e10 * 1e10)) * 1e10)) / 1.0) * 1e10)) / 1.0)
                            v8 = v8
                            v17 += (v17 / ((v17) | 1))
                            _ = v19
                        } else {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        }
                        _ = v17
                        _ = v18
                    }
                }
                _ = v14
            }
            v20 := []int{(fi9 * (fi9 + (fi9 | fi9))), fi9}
            _ = fi9
            _ = v20
        }
        fmt.Println(v8, p0)
    }
    p0 = p0
    switch ("\n\t" + "hello") {
    case "foo bar":
        var v21 uint = 325
        _ = v21
    }
    _ = p0
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 2
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    v5 := []int{0, 963, 387}
    var v6 int16 = int16(v2)
    var v22 int64 = 42
    var v23 int8 = -76
    var v24 bool = true
    v25 := GenS2{F0: 1, F1: 32767, F2: 42}
    switch -447 {
    case 2147483647:
        fmt.Println(v23)
    case -9223372036854775808:
        v24 = (100 <= v23)
    default:
        v6 += v6
    }
    v24 = ((4294967295 > 1) && (genFn0(v23) >= genFn0(v23)))
    var v26 uint16 = 2
    v27 := GenS3{F0: 4294967295, F1: "hello", F2: v22}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
