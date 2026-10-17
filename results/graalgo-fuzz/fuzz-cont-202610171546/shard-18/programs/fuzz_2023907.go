package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int64; F1 string; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint; F2 uint32 }
func genFn0(p0 uint16) int16 {
    if (65535 < 4294967295) {
        if true {
            v6 := GenS1{F0: 128, F1: p0}
            v7 := GenS2{F0: 9223372036854775807, F1: ("hello" + ("foo bar" + "\n\t")), F2: -789}
            _ = v6
            _ = v7
        }
    } else {
        fmt.Println(p0)
    }
    v8 := []int{32768}
    p0 += (p0 & p0)
    _ = p0
    _ = v8
    return -733
}
func genFn1(p0 string, p1 uint64) string {
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 float32 = 3.14159
        _ = fi9
        _ = v10
    }
    p0 = "foo bar"
    fmt.Println(p0, p0)
    p0 = (("hello" + p0) + "")
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 uint8 = 1
    var v2 int8 = 54
    var v3 int8 = (^(v2))
    var v4 uint64 = 18446744073709551615
    v5 := []uint32{708, 42, 65535}
    switch 9223372036854775807 {
    case 1:
        v5 = append(v5, 2147483647)
    case 976:
        fmt.Println(v2)
    case 4294967295:
        if (false && (-9223372036854775808 == -1)) {
            var v11 float64 = 3.14159
            v12 := GenS3{F0: -705}
            _ = v11
            _ = v12
        }
    }
    v13 := GenS2{F0: 9223372036854775807, F1: (("foo bar" + "") + (("hello" + ((("hello" + "hello") + "a") + ((("hello" + "foo bar") + "foo bar") + (("foo bar" + (("\n\t" + "a") + "")) + "foo bar")))) + "\n\t")), F2: 932}
    fmt.Println(v4)
    if (2147483647 >= int(v3)) {
        for fi14 := 0; fi14 < 11; fi14++ {
            v1 += 127
            var v15 int32 = 68
            _ = fi14
            _ = v15
        }
    }
    if ((true && false) && ((!false) || (0.0 != (3.14159 * 1e10)))) {
        fmt.Println(v1)
    }
    v4 += 0
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            if (v4 >= (v4 % ((v4) | 1))) {
                var v18 uint = 1
                fmt.Println(v2)
                _ = v18
            }
            _ = fi17
        }
        v4 = v4
        _ = fi16
    }
    if ((v3 % ((-128) | 1)) >= v2) {
        if false {
            if (127 < v1) {
                var v19 int32 = -84
                var v20 int8 = (v2 | v3)
                _ = v19
                _ = v20
            }
        } else {
            var v21 int64 = 601
            for fi22 := 0; fi22 < 2; fi22++ {
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 uint16 = 65535
                    switch fi23 {
                    case -483:
                        fmt.Println(v4)
                    case -1:
                        switch (fi23 | 0) {
                        case 4294967295:
                            if false {
                                _ = len(v5)
                                fmt.Println(v4, v2)
                            }
                        case 127:
                            fmt.Println(v21, v1)
                        case 526:
                            v25 := GenS4{F0: 42, F1: 255, F2: 1}
                            _ = v25
                        default:
                            var v26 uint64 = 0
                            _ = v26
                        }
                    case -637:
                        v24 = v24
                        v5 = append(v5, 406)
                        var v27 int64 = 0
                        _ = v27
                    }
                    _ = fi23
                    _ = v24
                }
                var v28 int64 = v21
                if (781 == 584) {
                    switch (fi22 >> ((18446744073709551615) & 63)) {
                    case 9223372036854775807:
                        v3 = v3
                    default:
                        fmt.Println(v2, v4, v2)
                    }
                } else {
                    var v29 int16 = -204
                    _ = v29
                }
                _ = fi22
                _ = v28
            }
            var v30 string = ("hello" + "\n\t")
            _ = v21
            _ = v30
        }
        if len(v5) > 0 {
            v5[0] = 2
        }
    } else {
        v4 -= v4
        var v31 uint8 = v1
        _ = v31
    }
    var v32 int64 = 0
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 uint16 = 168
        _ = fi33
        _ = v34
    }
    if (2147483647 > 830) {
        v5 = append(v5, 4294967295)
    } else {
        v35 := GenS3{F0: (v32 / ((v32) | 1))}
        v3 = (^(v2))
        for fi36 := 0; fi36 < 11; fi36++ {
            v1 = (v1 / (((^(v1))) | 1))
            _ = fi36
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
