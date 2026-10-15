package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int32 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint16; F1 int; F2 uint32 }
type GenS5 struct { F0 uint; F1 int }
func genFn0() uint {
    var v7 float32 = (0.1 / (1e10 - 3.14159))
    for fi8 := 0; fi8 < 13; fi8++ {
        switch (^(fi8)) {
        case 401:
            var v9 int32 = 1
            _ = v9
        case 127:
            v10 := GenS1{F0: 449, F1: 172, F2: 42}
            _ = v10
        case 32767:
            v11 := GenS2{F0: 115, F1: 1}
            _ = v11
        }
        v7 = v7
        _ = fi8
    }
    var v12 uint32 = 555
    var v13 int8 = -1
    v14 := []int{-9223372036854775808, 128, 9223372036854775807, -1}
    _ = v12
    _ = v13
    _ = v14
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := []uint32{858, 2}
    var v2 uint32 = 624
    var v3 uint8 = 2
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    var v6 float32 = 3.14159
    if ((42 < 0) && true) {
        v6 = (v6 + ((1.5 + 1e10) * 0.0))
        if (true || (!true)) {
            if true {
                if (((v3 & 127) <= v3) || true) {
                    if true {
                        v15 := GenS3{F0: 255, F1: 48}
                        _ = v15
                    }
                }
                var v16 int8 = -1
                _ = v16
            } else {
                for fi17 := 0; fi17 < 3; fi17++ {
                    v6 = ((v6 / v6) / (v6 / (v6 / 1.5)))
                    _ = len(v5)
                    _ = fi17
                }
                if false {
                    var v18 float32 = (1e10 / v6)
                    _ = len(v4)
                    _ = v18
                } else {
                    var v19 uint = 42
                    _ = v19
                }
                if (3.14159 <= 0.1) {
                    if false {
                        var v20 string = "a"
                        var v21 int32 = -1
                        v22 := GenS3{F0: 42, F1: 100}
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        v6 = v6
                        v1 = append(v1, (v2 * (v2 | v2)))
                        fmt.Println(v2, v2, v3)
                    }
                    var v23 int16 = -32768
                    var v24 uint32 = (v2 * (^(v2)))
                    _ = v23
                    _ = v24
                }
            }
            for fi25 := 0; fi25 < 9; fi25++ {
                v2 = v2
                _ = fi25
            }
            var v26 uint = 4294967295
            _ = v26
        } else {
            _ = len(v1)
            v1 = append(v1, v2)
        }
    }
    v5[718] = (v2 << ((935) & 63))
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 uint = 65535
        var v29 bool = true
        for fi30 := 0; fi30 < 7; fi30++ {
            fmt.Println(v28, fi27)
            _ = len(v5)
            if len(v1) > 0 {
                v1[0] = v2
            }
            _ = fi30
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    v2 += 953
    v2 = (v2 * (v2 ^ uint32(v3)))
    for fi31 := 0; fi31 < 2; fi31++ {
        v4[2147483647] = 726
        switch "hello" {
        case "a":
            if (false || (("hello" + "foo bar") == ("" + ((("foo bar" + (((("foo bar" + ("" + "foo bar")) + "hello") + ("a" + "hello")) + "a")) + "a") + "a")))) {
                for fi32 := 0; fi32 < 13; fi32++ {
                    var v33 uint16 = 76
                    _ = fi32
                    _ = v33
                }
            } else {
                for fi34 := 0; fi34 < 13; fi34++ {
                    v2 = (v2 & v2)
                    _ = fi34
                }
            }
        default:
            v6 += v6
        }
        v35 := GenS4{F0: 255, F1: fi31, F2: (^(v2))}
        _ = fi31
        _ = v35
    }
    v36 := GenS5{F0: 42, F1: -2147483648}
    _ = len(v1)
    var v37 int32 = 0
    v2 = v2
    fmt.Println(v6)
    var v38 int = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
