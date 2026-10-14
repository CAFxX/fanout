package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int16; F1 int16; F2 uint }
type GenS4 struct { F0 int64; F1 int32; F2 uint64 }
type GenS5 struct { F0 int16; F1 int16; F2 int64 }
func genFn0(p0 int32, p1 uint16) int32 {
    v7 := GenS1{F0: (true || true), F1: "hello", F2: 32768}
    fmt.Println(p1, p0)
    v8 := GenS2{F0: (p1 > p1), F1: true}
    switch "a" {
    case "hello":
        var v9 float32 = (-2.25 - (1.5 * 1.5))
        _ = v9
    case "foo bar":
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 27
}
func genFn1(p0 uint32) uint16 {
    v10 := GenS2{F0: true, F1: true}
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := []int{(fi11 - 829)}
        p0 += p0
        p0 += (p0 % ((4294967295) | 1))
        _ = fi11
        _ = v12
    }
    v13 := GenS3{F0: 0, F1: -1, F2: 183}
    _ = p0
    _ = v10
    _ = v13
    return 355
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 100
    v3 := []int{128, -793}
    v4 := make(map[string]bool)
    v5 := []string{(("\n\t" + ("" + "")) + ("\n\t" + ("\n\t" + ""))), (("hello" + "hello") + ("foo bar" + ("" + ""))), ("\n\t" + ("hello" + (("a" + "foo bar") + ""))), "hello"}
    v6 := make(map[string]string)
    v14 := GenS4{F0: -1, F1: 1, F2: 18446744073709551615}
    if true {
        var v15 uint = 0
        var v16 uint = v15
        var v17 float32 = 1e10
        _ = v15
        _ = v16
        _ = v17
    } else {
        v2 = (-(v2))
        v18 := GenS4{F0: 880, F1: -1, F2: 4294967295}
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 int16 = -57
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    var v21 int8 = (v2 & v2)
    var v22 int32 = -2147483648
    v23 := GenS5{F0: 32767, F1: 32767, F2: 9223372036854775807}
    v22 += v22
    switch 4294967295 {
    case 32768:
        var v24 int8 = (v21 + v21)
        _ = v24
    case 4294967295:
        switch 128 {
        case 1:
            var v25 uint16 = 255
            _ = v25
        case -587:
            var v26 uint32 = 0
            _ = v26
        case 2147483647:
            var v27 int = 860
            _ = v27
        }
    }
    v22 -= genFn0(1, 378)
    if false {
        v21 = (v21 - -50)
        if (v2 <= (^(v21))) {
            v2 += v2
        }
        var v28 uint32 = 385
        _ = v28
    } else {
        switch 1 {
        case 127:
            if (994 < 65535) {
                if true {
                    if false {
                        if (((4294967295 < 715) && (!true)) || true) {
                            fmt.Println(v21, v2)
                            var v29 uint8 = 133
                            var v30 int8 = (v21 << ((438) & 63))
                            _ = v29
                            _ = v30
                        } else {
                            fmt.Println(v2, v22)
                        }
                        if true {
                            v22 = (v22 / (((v22 << ((2) & 63))) | 1))
                            fmt.Println(v2, v21)
                        }
                    }
                    v31 := GenS4{F0: -966, F1: (^(v22)), F2: 4294967295}
                    v22 = v22
                    _ = v31
                }
                v32 := GenS2{F0: true, F1: true}
                _ = v32
            }
        case 2147483647:
            v33 := GenS3{F0: 32767, F1: -32768, F2: 221}
            _ = v33
        default:
            delete(v4, "")
        }
    }
    v2 = (v21 * 61)
    var v34 bool = (!(128 > 198))
    v22 -= v22
    for fi35 := 0; fi35 < 6; fi35++ {
        var v36 uint16 = 339
        _ = fi35
        _ = v36
    }
    switch ("hello" + "\n\t") {
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v37 uint64 = 4294967295
    v22 = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
