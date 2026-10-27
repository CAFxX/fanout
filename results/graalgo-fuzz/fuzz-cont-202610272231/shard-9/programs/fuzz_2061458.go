package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 int64; F2 int16 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint32, p1 int16, p2 uint8) int8 {
    var v4 int64 = 785
    var v5 float32 = ((float32(0.0) / 3.14159) * ((((0.0 - 0.0) + (((0.1 / 1.0) + (float32(-2.25) / 1.0)) - 1e10)) - 1e10) - 0.0))
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := GenS1{F0: 65535}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS2{F0: 0}
        _ = fi8
        _ = v9
    }
    var v10 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v4
    _ = v5
    return 100
}
func genFn1(p0 int, p1 int32, p2 string) uint32 {
    v11 := GenS2{F0: 4294967295}
    v12 := []int{(p0 + p0)}
    if (0 != 2) {
        if (true || ((1 <= 128) && true)) {
            if false {
                var v13 string = (p2 + (p2 + p2))
                var v14 int = p0
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v13
                _ = v14
            } else {
                var v15 int8 = 100
                v15 -= (v15 % ((v15) | 1))
                _ = v15
            }
            switch "hello" {
            case "\n\t":
                if (!(9223372036854775807 == 1)) {
                    var v16 uint32 = 65535
                    _ = v16
                }
            }
            switch p2 {
            case "hello":
                v17 := GenS3{F0: (!(!true))}
                _ = v17
            case "a":
                p2 = p2
            case "":
                p1 += 0
            default:
                switch (p0 + (p0 >> ((65535) & 63))) {
                case 32767:
                    var v18 uint8 = 26
                    _ = v18
                case -9223372036854775808:
                    var v19 int16 = -192
                    _ = v19
                case 1:
                    for fi20 := 0; fi20 < 13; fi20++ {
                        fmt.Println(p0, p2)
                        _ = fi20
                    }
                }
            }
        }
        v12 = append(v12, int(p1))
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    } else {
        v21 := GenS2{F0: 0}
        v22 := GenS1{F0: 1}
        switch "hello" {
        case "foo bar":
            if false {
                p1 = (p1 ^ (p1 + p1))
            }
        case "a":
            p0 += (p0 / (((p0 >> ((1) & 63))) | 1))
        case "hello":
            p0 = (p0 >> ((18446744073709551615) & 63))
        }
        _ = v21
        _ = v22
    }
    if (-1 == -9223372036854775808) {
        fmt.Println(p1, p0, p1)
        for fi23 := 0; fi23 < 11; fi23++ {
            fmt.Println(fi23, p1, fi23)
            var v24 uint8 = 128
            if true {
                var v25 uint = 4294967295
                _ = v25
            }
            _ = fi23
            _ = v24
        }
    } else {
        if true {
            for fi26 := 0; fi26 < 8; fi26++ {
                if (33 > genFn0(0, 1, 0)) {
                    v27 := GenS4{F0: 927, F1: 42, F2: -837}
                    _ = v27
                } else {
                    v28 := GenS5{F0: (-2.25 == (((0.1 / float64(1e10)) + float64(float32(((0.0 / (1.5 * 1e10)) - (1.5 * 1.5))))) / float64(1e10)))}
                    var v29 uint16 = 932
                    if true {
                        p0 -= int(v29)
                    }
                    _ = v28
                    _ = v29
                }
                _ = fi26
            }
        } else {
            switch p2 {
            case "\n\t":
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            case "hello":
                var v30 int = (^(p0))
                _ = v30
            case "":
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            default:
                switch (p2 + (p2 + "hello")) {
                case "a":
                    p0 += p0
                }
            }
            p0 = (^(p0))
            var v31 int64 = -1
            _ = v31
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 2
}

func main() {
    v1 := []uint32{511, 849, 2147483647, 42}
    v2 := make(map[string]uint32)
    var v3 int8 = 100
    var v32 uint32 = 301
    v3 = v3
    v32 = (v32 - (v32 ^ (v32 * (^(v32)))))
    v33 := GenS3{F0: false}
    delete(v2, "hello")
    fmt.Println(v3, v32)
    for fi34 := 0; fi34 < 2; fi34++ {
        if true {
            var v35 int16 = 42
            v32 = v32
            _ = v35
        } else {
            for fi36 := 0; fi36 < 3; fi36++ {
                v37 := GenS3{F0: (42 < 18446744073709551615)}
                if ((1.5 < (-2.25 + float32(0.1))) && (!(255 != 127))) {
                    _ = len(v1)
                }
                _ = fi36
                _ = v37
            }
            fmt.Println(v32, fi34)
            v32 -= 302
        }
        var v38 uint = 1
        _ = fi34
        _ = v38
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        _ = len(v2)
        var v40 float64 = (3.14159 * 0.0)
        _ = fi39
        _ = v40
    }
    var v41 bool = (false || (4294967295 <= 32768))
    if (0 == 1000) {
        v3 = (v3 * (v3 + (^(v3))))
        v41 = v41
        v42 := GenS4{F0: (^(v32)), F1: 42, F2: 32767}
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v41
}
