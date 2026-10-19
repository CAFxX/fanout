package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int16; F1 string; F2 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 bool, p1 uint32) int64 {
    v7 := []int{-2147483648}
    fmt.Println(p1)
    var v8 uint = 1
    p0 = p0
    if p0 {
        var v9 uint32 = (p1 >> (((v8 ^ v8)) & 63))
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 15; fi11++ {
                v9 -= (v9 | (p1 * (v9 ^ v9)))
                var v12 float32 = -2.25
                var v13 int = fi10
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
        _ = v9
    } else {
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS1{F0: 0, F1: -727, F2: 9223372036854775807}
            _ = fi14
            _ = v15
        }
        switch "a" {
        case "foo bar":
            var v16 int = 621
            _ = v16
        default:
            if ((0 <= 65535) && ((-111 != 47) && p0)) {
                v8 = (v8 / ((v8) | 1))
                if ((0.0 - 3.14159) > float64(3.14159)) {
                    p1 = (p1 >> (((v8 >> ((v8) & 63))) & 63))
                    switch ("" + ("" + "a")) {
                    case "\n\t":
                        if len(v7) > 0 {
                            v7[0] = -1
                        }
                    default:
                        fmt.Println(p0, v8, p0)
                    }
                } else {
                    p0 = (true && (-2147483648 <= -884))
                    fmt.Println(p1)
                }
                if true {
                    fmt.Println(p0)
                }
            } else {
                for fi17 := 0; fi17 < 15; fi17++ {
                    switch "a" {
                    case "a":
                        var v18 int = int(v8)
                        _ = v18
                    case "hello":
                        for fi19 := 0; fi19 < 10; fi19++ {
                            var v20 string = ("foo bar" + "a")
                            var v21 string = ("" + "foo bar")
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    default:
                        if (false && p0) {
                            v22 := GenS2{F0: 32767, F1: 2147483647, F2: 910}
                            if ((18446744073709551615 >= 18446744073709551615) || ((^(fi17)) < 1)) {
                                v8 = (v8 | (^(v8)))
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            } else {
                                fmt.Println(p1)
                            }
                            p1 = p1
                            _ = v22
                        } else {
                            var v23 int32 = 2147483647
                            _ = v23
                        }
                    }
                    _ = fi17
                }
                v8 += (v8 << (((v8 * 16)) & 63))
            }
        }
        v24 := GenS2{F0: 255, F1: 174, F2: 255}
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := []uint32{449}
    var v2 float32 = (1e10 + (float32(1.5) + 0.0))
    var v3 uint32 = 2147483648
    var v4 int = 1
    var v5 uint64 = 991
    v6 := []uint32{v3, 2, 2147483647, 2}
    v25 := GenS3{F0: -880, F1: "hello", F2: 0}
    v26 := GenS1{F0: 289, F1: 0, F2: 42}
    if (0 == 32767) {
        v3 = v3
    } else {
        fmt.Println(v3)
    }
    var v27 float64 = (3.14159 / 1e10)
    var v28 float64 = 1e10
    v29 := GenS1{F0: 1, F1: 1, F2: 1}
    var v30 int = 2147483647
    var v31 uint16 = 1
    for fi32 := 0; fi32 < 3; fi32++ {
        fmt.Println(v5, v27)
        _ = fi32
    }
    v33 := GenS4{F0: -1}
    v34 := GenS3{F0: 348, F1: "a", F2: 42}
    var v35 uint16 = (v31 ^ v31)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
