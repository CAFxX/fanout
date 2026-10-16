package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 bool }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint32) uint16 {
    if (-9223372036854775808 != 9223372036854775807) {
        if ((2147483647 > 2147483647) && (127 >= 114)) {
            v7 := []int{533, 127, -1}
            _ = v7
        }
        p0 -= 891
    } else {
        p0 -= (p0 >> ((42) & 63))
        if (0 >= 266) {
            p0 = 2
        }
        if false {
            var v8 int64 = 0
            if false {
                if (0 == 127) {
                    v9 := GenS1{F0: 18446744073709551615}
                    fmt.Println(v8)
                    _ = v9
                } else {
                    v10 := []int{-9223372036854775808, 127, 42, 1}
                    v8 += -712
                    var v11 uint64 = 42
                    _ = v10
                    _ = v11
                }
                v12 := []int{127, 9223372036854775807, -2147483648}
                _ = v12
            }
            if (42 >= 130) {
                v13 := GenS1{F0: 2}
                v14 := []int{2147483647, 128, -9223372036854775808, 32767}
                _ = v13
                _ = v14
            }
            _ = v8
        }
    }
    p0 += p0
    _ = p0
    return 375
}
func genFn1(p0 int64, p1 int64) int16 {
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := []int{32767, 32767, fi15}
        if len(v16) > 0 {
            v16[0] = -9223372036854775808
        }
        _ = fi15
        _ = v16
    }
    var v17 uint = 1
    _ = p0
    _ = p1
    _ = v17
    return -32768
}

func main() {
    var v1 int32 = 309
    var v2 int64 = 9223372036854775807
    var v3 uint = 18446744073709551615
    var v4 int = -410
    var v5 int32 = v1
    var v6 int32 = (v5 / (((v5 >> (((^(v3))) & 63))) | 1))
    var v18 string = ((("\n\t" + (("" + "a") + ("hello" + ("hello" + "hello")))) + "hello") + (((("foo bar" + ("\n\t" + ("\n\t" + (((("foo bar" + "hello") + "hello") + "a") + (("foo bar" + "a") + "a"))))) + "") + ("hello" + "a")) + ("hello" + "a")))
    fmt.Println(v4, v1, v1)
    v5 = (v6 * v5)
    if ((!false) && (true || (!true))) {
        v4 = v4
        for fi19 := 0; fi19 < 18; fi19++ {
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 float32 = (0.0 / 3.14159)
                fmt.Println(v1)
                var v22 float32 = 1.5
                _ = fi20
                _ = v21
                _ = v22
            }
            if (2 > 914) {
                v23 := GenS1{F0: 933}
                _ = v23
            }
            _ = fi19
        }
        var v24 uint16 = 255
        _ = v24
    } else {
        for fi25 := 0; fi25 < 1; fi25++ {
            if false {
                v6 = (v1 - (v5 * v6))
                for fi26 := 0; fi26 < 18; fi26++ {
                    switch v18 {
                    case "":
                        if true {
                            v1 += (v5 + v5)
                            var v27 bool = ((((v4 ^ fi26) >= -1) && false) && (!false))
                            var v28 int32 = (v5 % (((v6 * -575)) | 1))
                            _ = v27
                            _ = v28
                        } else {
                            if false {
                                v4 = -521
                                v29 := GenS2{F0: 0, F1: 2, F2: (genFn1(v2, v2) > 1000)}
                                var v30 int8 = 42
                                _ = v29
                                _ = v30
                            }
                            var v31 int32 = 2147483647
                            v32 := []int{fi25, fi26, 2147483647}
                            _ = v31
                            _ = v32
                        }
                    case "hello":
                        v4 += fi25
                    case "a":
                        v33 := GenS2{F0: 2, F1: 2147483647, F2: false}
                        _ = v33
                    default:
                        v34 := GenS3{F0: 128, F1: 1, F2: 1}
                        _ = v34
                    }
                    v5 = 206
                    _ = fi26
                }
                v35 := []int{}
                _ = v35
            }
            for fi36 := 0; fi36 < 1; fi36++ {
                if false {
                    if true {
                        var v37 int32 = (v6 << ((v3) & 63))
                        v38 := []int{fi25}
                        _ = v37
                        _ = v38
                    }
                }
                _ = fi36
            }
            _ = fi25
        }
        v4 = v4
        v39 := GenS4{F0: (^(v3)), F1: 426}
        _ = v39
    }
    fmt.Println(v1)
    var v40 uint16 = 0
    var v41 uint8 = 2
    v42 := []int{2147483647, 4294967295}
    v43 := GenS5{F0: 100}
    v4 -= (v4 >> (((v3 & 1)) & 63))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
