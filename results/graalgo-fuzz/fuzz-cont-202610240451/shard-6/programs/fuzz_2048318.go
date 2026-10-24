package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint8; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int64; F1 int16 }
func genFn0(p0 bool) int16 {
    v5 := GenS1{F0: -1}
    p0 = p0
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := []int{(fi6 & fi6)}
        v8 := GenS2{F0: 825, F1: 80, F2: ((p0 || (81 > 1)) || (-697 > -32768))}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint8 = 14
        v11 := []int{fi9, 32767, (fi9 / ((fi9) | 1)), fi9}
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{}
    _ = p0
    _ = v12
    _ = v5
    return -1
}
func genFn1(p0 bool, p1 uint64, p2 uint8) uint32 {
    if (p0 || (32767 >= 255)) {
        fmt.Println(p2, p2, p1)
        var v13 uint32 = 65535
        p0 = (genFn0((535 <= 42)) >= -750)
        _ = v13
    }
    p0 = false
    if p0 {
        var v14 string = "a"
        v15 := []int{-9223372036854775808, 128, 9223372036854775807}
        _ = v14
        _ = v15
    }
    var v16 float64 = -2.25
    v17 := []int{2147483647, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 0
}

func main() {
    v1 := []uint32{65535, 2, 2147483647}
    v2 := make(map[uint32]bool)
    var v3 float64 = 3.14159
    var v4 uint8 = 42
    v18 := GenS3{F0: 0}
    if (!(255 > 420)) {
        var v19 int16 = 823
        switch "" {
        case "":
            v20 := v2[2147483647]
            _ = v20
        case "hello":
            var v21 float64 = float64((float32((v3 + 1.5)) + 0.0))
            _ = v21
        }
        switch (("" + "") + "") {
        case "hello":
            var v22 uint8 = (v4 ^ v4)
            _ = v22
        default:
            v23 := GenS2{F0: (v19 >> ((18446744073709551615) & 63)), F1: 127, F2: (((true && ("" < "foo bar")) || true) || true)}
            _ = v23
        }
        _ = v19
    } else {
        switch -2147483648 {
        case -9223372036854775808:
            v4 -= (^(v4))
        default:
            v4 = 42
        }
    }
    fmt.Println(v4, v3)
    v24 := GenS4{F0: -477, F1: genFn0((("hello" + "hello") != "a"))}
    var v25 float64 = 3.14159
    if true {
        var v26 bool = ((42 > 2147483647) || true)
        _ = v26
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if ((0 < 2) || (9223372036854775807 == -9223372036854775808)) {
            v27 := GenS3{F0: 865}
            v25 = float64((3.14159 + (1e10 - ((0.0 * 1.5) + float32(float64((3.14159 + (((-2.25 - 3.14159) + 3.14159) - (1e10 * (3.14159 / ((float32(v25) - 3.14159) + (0.0 - -2.25))))))))))))
            _ = v27
        } else {
            v28 := GenS4{F0: 32767, F1: 641}
            v25 += ((v25 * v3) * v25)
            _ = v28
        }
    }
    v29 := GenS3{F0: 0}
    if (true || (!true)) {
        fmt.Println(v4, v4)
        var v30 int32 = 0
        _ = v30
    }
    v31 := GenS5{F0: -9223372036854775808, F1: int16(v4)}
    if (42 > 0) {
        var v32 int = 9223372036854775807
        for fi33 := 0; fi33 < 2; fi33++ {
            v3 = 1e10
            for fi34 := 0; fi34 < 18; fi34++ {
                v25 += (1e10 / v25)
                var v35 int = fi34
                if false {
                    switch (^(v35)) {
                    case -377:
                        switch fi34 {
                        case -644:
                            v36 := GenS2{F0: -32768, F1: 128, F2: true}
                            _ = v36
                        case -836:
                            var v37 uint16 = 480
                            _ = v37
                        default:
                            v38 := GenS3{F0: 620}
                            _ = v38
                        }
                    case 32767:
                        _ = len(v1)
                    case 4294967295:
                        fmt.Println(v3)
                    }
                }
                _ = fi34
                _ = v35
            }
            _ = fi33
        }
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v31
    _ = v4
}
