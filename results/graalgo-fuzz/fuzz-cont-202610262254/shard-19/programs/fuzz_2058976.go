package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint32; F1 int8; F2 int32 }
type GenS3 struct { F0 string }
func genFn0(p0 int32) string {
    p0 = (p0 & p0)
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 bool = (!(42 > 2))
            v9 := GenS1{F0: -1, F1: v8}
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        if (1 > 1) {
            v10 := GenS2{F0: 65535, F1: 0, F2: (p0 % ((0) | 1))}
            _ = v10
        } else {
            v11 := []int{1, -1}
            for fi12 := 0; fi12 < 16; fi12++ {
                p0 = p0
                p0 = p0
                v13 := GenS3{F0: ("foo bar" + ("hello" + "foo bar"))}
                _ = fi12
                _ = v13
            }
            if false {
                p0 += 240
            }
            _ = v11
        }
        if true {
            v14 := []int{-167, -355, 42, 128}
            var v15 float32 = 0.0
            var v16 int32 = p0
            _ = v14
            _ = v15
            _ = v16
        }
    }
    _ = p0
    return "hello"
}
func genFn1(p0 uint, p1 uint64, p2 uint32) int64 {
    for fi17 := 0; fi17 < 14; fi17++ {
        if true {
            switch "a" {
            case "foo bar":
                if (-206 >= -964) {
                    var v18 uint8 = 128
                    var v19 uint16 = 385
                    v20 := []int{(fi17 / ((76) | 1))}
                    _ = v18
                    _ = v19
                    _ = v20
                }
            }
        } else {
            if (!true) {
                v21 := GenS2{F0: p2, F1: -128, F2: -952}
                v22 := GenS2{F0: 4294967295, F1: -77, F2: -437}
                _ = v21
                _ = v22
            } else {
                if (false || true) {
                    var v23 uint16 = 65535
                    _ = v23
                } else {
                    var v24 bool = true
                    _ = v24
                }
            }
            var v25 uint32 = 42
            var v26 float32 = ((1e10 * 1e10) - (((1e10 + float32(((float64(1.5) / (((((0.1 + 1e10) - -2.25) + float64((1.5 / (float32(3.14159) * 3.14159)))) / -2.25) * ((1e10 / -2.25) + (0.0 + float64(0.1))))) * 3.14159))) * 1.5) + (float32((1.5 + 1e10)) / 0.1)))
            _ = v25
            _ = v26
        }
        _ = fi17
    }
    var v27 float64 = (1e10 / (float64(1e10) * -2.25))
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 9223372036854775807
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]uint32)
    var v3 string = ("" + "foo bar")
    var v4 uint32 = 2147483647
    var v5 int64 = 0
    v6 := []string{v3, (v3 + ""), v3, "foo bar"}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v4)
    v28 := GenS2{F0: (^(v4)), F1: 0, F2: 2147483647}
    v4 = (v4 & (v4 ^ 2))
    var v29 uint8 = 2
    for fi30 := 0; fi30 < 13; fi30++ {
        if false {
            if (18446744073709551615 != 4294967295) {
                var v31 uint = v1
                _ = v31
            }
            var v32 uint = v1
            _ = v32
        }
        var v33 int16 = 579
        _ = fi30
        _ = v33
    }
    v34 := GenS2{F0: 4294967295, F1: -30, F2: -1}
    var v35 float32 = 3.14159
    v36 := GenS1{F0: 1, F1: (v4 > (v4 + v4))}
    v1 -= (v1 << ((v1) & 63))
    v37 := GenS3{F0: v3}
    var v38 int32 = -1
    if (!true) {
        var v39 string = ""
        _ = v39
    } else {
        for fi40 := 0; fi40 < 4; fi40++ {
            var v41 int32 = v38
            for fi42 := 0; fi42 < 1; fi42++ {
                var v43 uint16 = 65535
                _ = len(v2)
                v6 = append(v6, v3)
                v44 := GenS1{F0: int16(v43), F1: (42 < -32768)}
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = fi40
            _ = v41
        }
        var v45 int16 = 32767
        _ = v45
    }
    fmt.Println(v5, v3, v38)
    v3 = (v3 + v3)
    var v46 float32 = 1.5
    if (!(0.1 >= (1e10 * float64(0.0)))) {
        var v47 uint64 = 18446744073709551615
        v48 := GenS3{F0: ("a" + v3)}
        _ = v47
        _ = v48
    } else {
        var v49 int32 = -881
        v50 := GenS3{F0: "a"}
        _ = v49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
