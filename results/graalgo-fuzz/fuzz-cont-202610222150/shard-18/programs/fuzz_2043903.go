package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 int64; F1 int32; F2 int64 }
type GenS6 struct { F0 uint16; F1 int; F2 int16 }
type GenS7 struct { F0 int16 }
func genFn0(p0 int64) uint8 {
    p0 -= p0
    var v7 int16 = -412
    v8 := []int{}
    if (!(!(0 >= -580))) {
        p0 = p0
    } else {
        v9 := GenS1{F0: 368, F1: (v7 >> ((131) & 63)), F2: 65535}
        var v10 bool = ((-2.25 - ((float64((1.5 + 1e10)) / 1.0) * (0.1 - 1.5))) <= 0.1)
        var v11 float64 = 3.14159
        _ = v10
        _ = v11
        _ = v9
    }
    if true {
        fmt.Println(p0)
        if true {
            for fi12 := 0; fi12 < 6; fi12++ {
                p0 = 1
                p0 = -9223372036854775808
                p0 += (p0 ^ (-(p0)))
                _ = fi12
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                if len(v8) > 0 {
                    v8[0] = fi13
                }
                v14 := GenS1{F0: 1, F1: -32768, F2: 569}
                _ = fi13
                _ = v14
            }
        } else {
            v7 -= (v7 + (v7 - v7))
            for fi15 := 0; fi15 < 7; fi15++ {
                for fi16 := 0; fi16 < 4; fi16++ {
                    v17 := GenS1{F0: (fi15 - 128), F1: (v7 ^ v7), F2: 65535}
                    var v18 uint64 = 0
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                p0 -= p0
                _ = fi15
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                v7 = v7
                v20 := GenS2{F0: 32767, F1: -9223372036854775808, F2: true}
                fmt.Println(p0, v7)
                _ = fi19
                _ = v20
            }
        }
        var v21 int16 = v7
        _ = v21
    }
    _ = p0
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int, p1 int64, p2 uint8) int {
    p1 -= (^(p1))
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 int32 = 1
    var v2 string = ("\n\t" + "foo bar")
    v3 := make(map[string]int)
    var v4 int16 = int16(v1)
    var v5 uint8 = 2
    var v6 uint32 = 913
    if false {
        v4 -= v4
    } else {
        var v22 int32 = (^(v1))
        if false {
            var v23 uint32 = 819
            var v24 int64 = -239
            if (v4 != (v4 & v4)) {
                v24 = v24
            } else {
                var v25 uint16 = 65535
                _ = v25
            }
            _ = v23
            _ = v24
        } else {
            v1 += (^(v1))
            for fi26 := 0; fi26 < 16; fi26++ {
                v27 := GenS3{F0: (v4 + v4), F1: -632}
                _ = fi26
                _ = v27
            }
            v6 = (v6 << ((291) & 63))
        }
        switch v2 {
        case "foo bar":
            var v28 uint = 0
            _ = v28
        case "\n\t":
            var v29 int16 = v4
            _ = v29
        }
        _ = v22
    }
    if ((true || false) || true) {
        for fi30 := 0; fi30 < 9; fi30++ {
            for fi31 := 0; fi31 < 4; fi31++ {
                v32 := []int{}
                v1 += (v1 + v1)
                v6 = (v6 % (((v6 >> ((1) & 63))) | 1))
                switch "\n\t" {
                case "hello":
                    v33 := GenS4{F0: (v4 / (((v4 >> ((uint(v5)) & 63))) | 1)), F1: v5}
                    _ = v33
                case "foo bar":
                    if len(v32) > 0 {
                        v32[0] = (^(fi30))
                    }
                case "\n\t":
                    var v34 uint = 4294967295
                    _ = v34
                default:
                    v35 := GenS1{F0: (fi31 * fi31), F1: -32768, F2: 268}
                    _ = v35
                }
                if (!false) {
                    v36 := GenS5{F0: 42, F1: (v1 >> ((18446744073709551615) & 63)), F2: 686}
                    _ = v36
                }
                switch v2 {
                case "a":
                    var v37 string = v2
                    _ = v37
                case "foo bar":
                    if (false && false) {
                        v38 := GenS2{F0: 1, F1: fi31, F2: true}
                        var v39 int32 = int32(v5)
                        if ((true && true) || false) {
                            v5 = v5
                            delete(v3, (("" + ("" + v2)) + (((("\n\t" + v2) + "") + (v2 + (v2 + v2))) + (v2 + v2))))
                            var v40 uint64 = 1
                            if (v2 != v2) {
                                v6 += (v6 - 65535)
                                fmt.Println(v2, fi30)
                            }
                            _ = v40
                        } else {
                            v41 := GenS4{F0: v4, F1: 19}
                            fmt.Println(v6)
                            _ = v41
                        }
                        v4 = v4
                        var v42 string = "\n\t"
                        _ = v38
                        _ = v39
                        _ = v42
                    } else {
                        var v43 int64 = 42
                        _ = v43
                    }
                case "\n\t":
                    _ = len(v3)
                }
                v32 = append(v32, (fi31 >> ((54) & 63)))
                v44 := GenS1{F0: fi30, F1: (-(v4)), F2: 4}
                v5 = genFn0(-9223372036854775808)
                _ = fi31
                _ = v32
                _ = v44
            }
            v45 := []int{fi30, (^(fi30)), 1}
            _ = fi30
            _ = v45
        }
        for fi46 := 0; fi46 < 12; fi46++ {
            fmt.Println(fi46, v6, v2)
            _ = fi46
        }
    } else {
        var v47 uint = 919
        delete(v3, (v2 + v2))
        _ = v47
    }
    v48 := GenS4{F0: v4, F1: v5}
    fmt.Println(v5, v6, v1)
    var v49 int64 = 1
    for fi50 := 0; fi50 < 5; fi50++ {
        v2 = v2
        _ = fi50
    }
    v2 = v2
    if true {
        v2 = v2
        if ((("hello" + "hello") + "\n\t") >= ((v2 + v2) + v2)) {
            fmt.Println(v2, v49)
            for fi51 := 0; fi51 < 20; fi51++ {
                v52 := GenS6{F0: 32767, F1: fi51, F2: v4}
                _ = fi51
                _ = v52
            }
        }
    } else {
        for fi53 := 0; fi53 < 2; fi53++ {
            fmt.Println(v6, v6)
            _ = fi53
        }
        v54 := GenS7{F0: (v4 | int16(v49))}
        v55 := GenS4{F0: v4, F1: 172}
        _ = v54
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
