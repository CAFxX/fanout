package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint; F1 int16 }
func genFn0(p0 uint32, p1 uint8) int64 {
    p0 = 2147483648
    switch 127 {
    case -464:
        var v5 int16 = 1
        _ = v5
    case 127:
        v6 := []int{32767, 4294967295, -232, -760}
        _ = v6
    case -746:
        var v7 string = "a"
        _ = v7
    }
    v8 := []int{-1, -2147483648, 0}
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: p0, F1: 65535}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return 0
}
func genFn1(p0 int, p1 int32, p2 int64) string {
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := []int{p0, 128}
            _ = fi12
            _ = v13
        }
        var v14 int32 = (p1 ^ p1)
        var v15 uint = 255
        _ = fi11
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 590, F1: 139}
    var v17 float32 = 0.0
    v18 := GenS1{F0: 0, F1: 30}
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return ("\n\t" + (("a" + (("foo bar" + "hello") + "a")) + ("hello" + ("hello" + ""))))
}

func main() {
    var v1 uint16 = 32767
    var v2 int64 = 1
    v3 := []uint32{920, 525, 2147483647, 2147483648}
    var v4 int = 2147483647
    v19 := GenS1{F0: 4294967295, F1: 372}
    if len(v3) > 0 {
        v3[0] = 0
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 uint = 4294967295
        v1 += 255
        _ = fi20
        _ = v21
    }
    v3 = append(v3, 4294967295)
    v2 -= 587
    fmt.Println(v1, v1, v1)
    if (false || true) {
        if (4294967295 <= 1) {
            for fi22 := 0; fi22 < 13; fi22++ {
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 int = 42
                    var v25 float32 = 0.1
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                if true {
                    switch (v4 + (v4 % ((v4) | 1))) {
                    case 128:
                        v26 := GenS2{F0: 333, F1: 2, F2: fi22}
                        _ = v26
                    case 32767:
                        for fi27 := 0; fi27 < 5; fi27++ {
                            var v28 string = ""
                            v1 = (v1 % (((v1 ^ 0)) | 1))
                            if false {
                                _ = len(v3)
                            } else {
                                var v29 uint64 = 204
                                _ = v29
                            }
                            _ = fi27
                            _ = v28
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                    v1 += v1
                }
                _ = fi22
            }
            v3 = append(v3, 4294967295)
        }
    } else {
        if len(v3) > 0 {
            v3[0] = 1
        }
        var v30 uint64 = 4294967295
        if (!(-2147483648 <= 1)) {
            var v31 uint8 = 141
            fmt.Println(v31, v2, v2)
            if (true && (0 > 127)) {
                v31 = (^(v31))
                var v32 float32 = 0.0
                _ = v32
            }
            _ = v31
        }
        _ = v30
    }
    fmt.Println(v1)
    v4 += v4
    _ = len(v3)
    switch (v4 >> ((65535) & 63)) {
    case 702:
        if true {
            if (!((!false) && (((1e10 + float64((((float32(((0.0 - 1.5) + (3.14159 / 1.0))) * (0.1 / float32((0.1 * 1.5)))) * 0.0) / (1e10 * (((0.0 + 3.14159) - 0.1) + -2.25))))) - 1e10) > 1e10))) {
                if len(v3) > 0 {
                    v3[0] = 92
                }
                var v33 uint8 = 248
                v34 := GenS2{F0: 450, F1: 2, F2: v4}
                _ = v33
                _ = v34
            }
        }
    case 42:
        if (v1 > (v1 ^ (^(v1)))) {
            if len(v3) > 0 {
                v3[0] = 1
            }
            v1 = v1
        }
    default:
        for fi35 := 0; fi35 < 3; fi35++ {
            v2 = genFn0(4294967295, 128)
            switch fi35 {
            case 2147483647:
                if (((1.5 - 0.1) + 3.14159) < 0.1) {
                    v36 := GenS3{F0: -610, F1: -236}
                    _ = v36
                }
            case 388:
                v37 := GenS4{F0: 42, F1: 42}
                _ = v37
            default:
                v38 := GenS5{F0: 0, F1: 42}
                _ = v38
            }
            _ = fi35
        }
    }
    v3 = append(v3, 2147483648)
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
