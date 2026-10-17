package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }
type GenS3 struct { F0 int16; F1 int16; F2 uint8 }
func genFn0() uint16 {
    var v7 int = 4294967295
    var v8 uint8 = 42
    for fi9 := 0; fi9 < 19; fi9++ {
        v8 = (v8 >> ((42) & 63))
        _ = fi9
    }
    var v10 int8 = 115
    _ = v10
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int, p1 int64, p2 uint64) string {
    v11 := GenS1{F0: 430, F1: 483}
    switch "hello" {
    case "foo bar":
        v12 := []int{128}
        _ = v12
    default:
        var v13 int64 = (p1 % ((p1) | 1))
        _ = v13
    }
    if (0 != 765) {
        switch (p0 * p0) {
        case 4294967295:
            v14 := GenS1{F0: 32767, F1: genFn0()}
            _ = v14
        default:
            switch "a" {
            case "a":
                v15 := []int{p0, (^(p0)), (p0 ^ p0)}
                _ = v15
            case "\n\t":
                v16 := GenS2{F0: uint16(p2), F1: "foo bar", F2: 0}
                _ = v16
            }
        }
    } else {
        fmt.Println(p2, p0, p1)
        var v17 uint = 0
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return "a"
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 float32 = (-2.25 / float32((1.5 - -2.25)))
    v5 := make(map[string]int)
    var v6 uint32 = 1
    v18 := GenS2{F0: 65535, F1: "\n\t", F2: -622}
    var v19 uint8 = v1
    v20 := GenS1{F0: 65535, F1: genFn0()}
    switch "" {
    case "":
        if ((^(v6)) >= 2147483648) {
            if (-1 == -1) {
                switch "foo bar" {
                case "a":
                    v1 = (v19 + 42)
                case "hello":
                    v21 := GenS3{F0: 885, F1: -1, F2: v1}
                    _ = v21
                }
                var v22 uint64 = 343
                _ = v22
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                switch (fi23 << ((4294967295) & 63)) {
                case 4294967295:
                    if (false && (!(v4 == ((1e10 / v4) - 3.14159)))) {
                        switch (("hello" + "\n\t") + "a") {
                        case "":
                            var v24 int64 = 387
                            _ = v24
                        case "hello":
                            switch "hello" {
                            case "":
                                delete(v3, (v6 | (v6 << ((0) & 63))))
                            default:
                                v25 := GenS1{F0: 2, F1: 0}
                                _ = v25
                            }
                        default:
                            v19 += 1
                        }
                        var v26 uint16 = 891
                        fmt.Println(v1, v1, v19)
                        _ = v26
                    } else {
                        fmt.Println(v4, fi23)
                    }
                case 569:
                    var v27 uint64 = uint64(v6)
                    _ = v27
                }
                v28 := GenS2{F0: 255, F1: "hello", F2: 15}
                _ = fi23
                _ = v28
            }
            var v29 uint32 = (v6 ^ v6)
            _ = v29
        } else {
            var v30 uint32 = 827
            _ = v30
        }
    case "foo bar":
        v31 := GenS1{F0: 65535, F1: 32767}
        _ = v31
    case "hello":
        fmt.Println(v6, v19, v1)
    }
    switch 42 {
    case 128:
        var v32 float32 = 1.5
        _ = v32
    case -1:
        var v33 bool = true
        _ = v33
    case 164:
        for fi34 := 0; fi34 < 13; fi34++ {
            v1 = (v1 << ((517) & 63))
            v35 := []int{(^(fi34)), 2147483647}
            v2[(fi34 | fi34)] = fi34
            _ = fi34
            _ = v35
        }
    default:
        var v36 uint = 18446744073709551615
        _ = v36
    }
    v37 := []int{1, 9223372036854775807}
    v6 = 711
    var v38 int16 = 1000
    if len(v37) >= 2 {
        v37 = v37[1:len(v37)]
    }
    if true {
        v39 := GenS2{F0: 1, F1: ("a" + "foo bar"), F2: 248}
        _ = v39
    } else {
        v4 -= (float32(-2.25) + float32(1e10))
        switch -9223372036854775808 {
        case -31:
            v4 = v4
        }
    }
    fmt.Println(v6, v6, v6)
    var v40 int64 = 9223372036854775807
    var v41 int16 = v38
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
