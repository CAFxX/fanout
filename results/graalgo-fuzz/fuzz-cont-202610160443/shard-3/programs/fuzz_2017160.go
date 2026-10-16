package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int8; F1 int; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 string }
type GenS5 struct { F0 uint64; F1 string; F2 uint8 }
func genFn0() bool {
    var v7 int = 406
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := []int{fi8}
        switch "a" {
        case "foo bar":
            v10 := GenS1{F0: 62, F1: 844}
            _ = v10
        case "hello":
            v11 := GenS1{F0: 387, F1: 4294967295}
            _ = v11
        case "a":
            v9 = append(v9, (v7 * (fi8 - (fi8 + (fi8 % (((v7 ^ (^(v7)))) | 1))))))
        default:
            v9 = append(v9, fi8)
        }
        var v12 int16 = -1
        _ = fi8
        _ = v12
        _ = v9
    }
    switch v7 {
    case 2147483647:
        switch (v7 % ((v7) | 1)) {
        case -439:
            var v13 int32 = 1
            _ = v13
        case 32768:
            fmt.Println(v7)
        }
    case -510:
        var v14 uint32 = 0
        _ = v14
    }
    v7 = v7
    var v15 int32 = 2147483647
    _ = v15
    _ = v7
    return (false || false)
}
func genFn1() uint16 {
    var v16 int16 = 1000
    v17 := GenS1{F0: 2147483647, F1: 433}
    if false {
        if false {
            var v18 int64 = 870
            fmt.Println(v16, v18)
            v19 := GenS1{F0: -1, F1: 65535}
            _ = v18
            _ = v19
        }
        v16 -= v16
        v16 -= v16
    }
    _ = v16
    _ = v17
    return 32767
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{2147483647, 42, 0, -158}
    var v3 float32 = -2.25
    v4 := make(map[int]uint32)
    v5 := []uint32{4294967295}
    var v6 int16 = 1000
    fmt.Println(v3, v3)
    if true {
        v6 = 1000
        if false {
            fmt.Println(v3, v3)
        }
    } else {
        v3 += 1e10
        v1[-9223372036854775808] = ("hello" + ("a" + "hello"))
        fmt.Println(v6, v3)
    }
    var v20 int64 = -15
    v21 := GenS1{F0: 1, F1: 65535}
    var v22 uint16 = genFn1()
    v6 = 1
    v20 -= v20
    v6 = (v6 * (v6 ^ int16(v22)))
    switch 612 {
    case 0:
        v20 -= v20
    case 42:
        v23 := GenS2{F0: v22, F1: true}
        _ = v23
    default:
        switch -9223372036854775808 {
        case -946:
            v22 = genFn1()
        case 779:
            for fi24 := 0; fi24 < 13; fi24++ {
                v20 = (-(v20))
                if (!(true && (!(!false)))) {
                    v25 := GenS2{F0: genFn1(), F1: true}
                    for fi26 := 0; fi26 < 4; fi26++ {
                        switch ("a" + "a") {
                        case "hello":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "":
                            v27 := GenS1{F0: -777, F1: 625}
                            _ = v27
                        case "foo bar":
                            v28 := GenS3{F0: -116, F1: (fi26 - (-(fi24))), F2: 128}
                            _ = v28
                        }
                        v29 := GenS4{F0: 0, F1: 0, F2: "foo bar"}
                        _ = fi26
                        _ = v29
                    }
                    _ = v25
                } else {
                    delete(v1, fi24)
                }
                _ = fi24
            }
        default:
            v6 += (^(v6))
        }
    }
    v30 := GenS1{F0: 42, F1: 42}
    v22 -= v22
    var v31 uint32 = 681
    fmt.Println(v22, v31, v20)
    if true {
        var v32 float32 = float32(((0.1 + (((0.1 / (float64(v3) / 0.1)) * (1e10 + (3.14159 * 3.14159))) * 1.5)) / 1.5))
        v33 := GenS5{F0: 18446744073709551615, F1: "", F2: 255}
        _ = v32
        _ = v33
    }
    v6 += v6
    switch 32767 {
    case -9223372036854775808:
        var v34 uint64 = 0
        _ = v34
    case 2147483647:
        if true {
            v6 = (^(v6))
        } else {
            fmt.Println(v6, v31, v31)
            for fi35 := 0; fi35 < 3; fi35++ {
                v36 := GenS1{F0: 2147483647, F1: (^(v31))}
                _ = fi35
                _ = v36
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
