package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint16; F1 bool; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int; F2 int8 }
func genFn0(p0 uint32, p1 uint8) int64 {
    v4 := GenS1{F0: 2147483647, F1: -1}
    if (!(0 <= -455)) {
        fmt.Println(p1)
        v5 := []int{}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1() uint32 {
    v6 := []int{42, 32768}
    if (((0 != -2147483648) || ((759 != 4294967295) || (-76 < -1))) || false) {
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 string = "a"
                var v10 int16 = 557
                var v11 uint32 = 42
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 uint16 = 65535
            if len(v6) > 0 {
                v6[0] = fi7
            }
            _ = fi7
            _ = v12
        }
        var v13 int32 = 42
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v13
    } else {
        var v14 int = 128
        if len(v6) > 0 {
            v6[0] = v14
        }
        fmt.Println(v14)
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int64 = -709
        _ = fi15
        _ = v16
    }
    switch 127 {
    case 9223372036854775807:
        if true {
            var v17 uint8 = 1
            _ = v17
        }
    default:
        var v18 int16 = -700
        _ = v18
    }
    _ = v6
    return 439
}

func main() {
    var v1 int16 = -1
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    fmt.Println(v1)
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint32 = 1
        _ = fi19
        _ = v20
    }
    v1 += v1
    v1 -= (v1 - v1)
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 uint32 = 565
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS2{F0: 32767, F1: false, F2: 45}
            var v26 int8 = -1
            v27 := v3[fi23]
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        v28 := []int{-1, (fi23 | (fi23 * fi23)), fi23}
        v1 = v1
        _ = fi23
        _ = v28
    }
    v1 = (v1 ^ (v1 ^ v1))
    if (255 > 2) {
        if true {
            if (42 >= 207) {
                var v29 int16 = (v1 | (-(v1)))
                _ = v29
            }
        } else {
            v1 -= -782
        }
        var v30 int = 959
        fmt.Println(v30, v1)
        _ = v30
    }
    if (42 <= 18446744073709551615) {
        fmt.Println(v1)
    }
    switch 569 {
    case 128:
        var v31 int32 = 179
        _ = v31
    case -9223372036854775808:
        fmt.Println(v1)
    case -104:
        var v32 uint64 = 486
        _ = v32
    }
    v33 := []int{127, 128, -485, 2147483647}
    for fi34 := 0; fi34 < 8; fi34++ {
        switch ("\n\t" + "a") {
        case "hello":
            if (true && false) {
                v35 := GenS3{F0: 255}
                v36 := GenS4{F0: 406, F1: fi34, F2: int8(fi34)}
                _ = v35
                _ = v36
            }
        case "a":
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 uint32 = 65535
                var v39 uint16 = 640
                switch ("a" + "") {
                case "a":
                    v1 = (v1 * (^(v1)))
                case "\n\t":
                    fmt.Println(fi37, v38)
                default:
                    var v40 string = ("hello" + ("a" + ""))
                    var v41 float32 = 0.1
                    _ = v40
                    _ = v41
                }
                _ = fi37
                _ = v38
                _ = v39
            }
        default:
            if (124 != 42) {
                switch "hello" {
                case "foo bar":
                    v1 += (v1 * int16(fi34))
                case "a":
                    var v42 int16 = v1
                    _ = v42
                case "hello":
                    v1 -= v1
                }
                v1 = (v1 << ((18446744073709551615) & 63))
                var v43 int16 = v1
                _ = v43
            }
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
}
