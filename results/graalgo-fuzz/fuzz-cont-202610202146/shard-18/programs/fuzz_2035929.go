package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint8, p1 string) int16 {
    switch (("a" + p1) + (p1 + ("hello" + ("hello" + p1)))) {
    case "a":
        var v6 uint = 724
        _ = v6
    default:
        v7 := []int{-2147483648, -72}
        _ = v7
    }
    fmt.Println(p0, p0)
    if (42 < 42) {
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                fmt.Println(p0)
                switch fi8 {
                case 32768:
                    var v10 int = (fi8 << ((196) & 63))
                    _ = v10
                case 127:
                    var v11 int = fi9
                    _ = v11
                case 9223372036854775807:
                    var v12 int32 = 42
                    _ = v12
                }
                v13 := GenS1{F0: -32768, F1: 2147483647}
                v14 := GenS2{F0: 42, F1: (p0 / (((p0 >> ((33) & 63))) | 1))}
                _ = fi9
                _ = v13
                _ = v14
            }
            p1 = ((("a" + p1) + (p1 + "\n\t")) + "foo bar")
            _ = fi8
        }
    } else {
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(fi15, p1, fi15)
            _ = fi15
        }
        v16 := GenS1{F0: 1000, F1: 42}
        v17 := GenS1{F0: 813, F1: 42}
        _ = v16
        _ = v17
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 876
}
func genFn1(p0 int8, p1 uint16, p2 int8) int {
    switch 450 {
    case 32768:
        var v18 int16 = -988
        _ = v18
    case -9223372036854775808:
        var v19 int16 = -779
        _ = v19
    default:
        if (757 == 42) {
            if ("" == "\n\t") {
                var v20 string = "hello"
                switch (((("foo bar" + v20) + (v20 + (v20 + v20))) + v20) + (((("hello" + v20) + ("hello" + v20)) + ("a" + "a")) + v20)) {
                case "foo bar":
                    switch v20 {
                    case "":
                        v21 := GenS2{F0: p2, F1: 255}
                        _ = v21
                    case "foo bar":
                        if (!false) {
                            fmt.Println(v20, p2, p2)
                            var v22 int16 = 32767
                            _ = v22
                        }
                    default:
                        v23 := []int{-9223372036854775808, 179, 32767}
                        _ = v23
                    }
                case "hello":
                    v24 := []int{}
                    _ = v24
                }
                _ = v20
            }
            p1 = (p1 & 0)
        } else {
            fmt.Println(p0)
            v25 := GenS3{F0: uint64(p1), F1: 65535, F2: 4294967295}
            var v26 uint = 0
            _ = v25
            _ = v26
        }
    }
    switch "hello" {
    case "a":
        var v27 int8 = p0
        _ = v27
    case "hello":
        var v28 uint64 = 817
        _ = v28
    }
    v29 := GenS1{F0: 0, F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return 127
}

func main() {
    var v1 float32 = (3.14159 * 1e10)
    var v2 uint8 = 240
    var v3 uint = 255
    var v4 float32 = (v1 / v1)
    v5 := make(map[uint32]string)
    var v30 int16 = 1
    switch 32768 {
    case 0:
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 int8 = -128
            _ = fi31
            _ = v32
        }
    case 32768:
        var v33 int64 = -1
        _ = v33
    case 1:
        for fi34 := 0; fi34 < 8; fi34++ {
            v3 = v3
            var v35 uint32 = 65535
            var v36 float64 = (float64(((v1 + v1) + (v4 + v1))) - ((1.5 * 0.0) / (3.14159 / (float64(float32((1e10 - ((0.0 * 3.14159) - -2.25)))) / ((0.1 * (((-2.25 * 3.14159) + -2.25) - ((1.5 * 3.14159) * (-2.25 - (0.1 / 1e10))))) - ((3.14159 / (float64(3.14159) + float64(v1))) * 1.5))))))
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    v37 := []int{0, 128, -9223372036854775808}
    for fi38 := 0; fi38 < 19; fi38++ {
        if (18446744073709551615 <= 42) {
            var v39 int8 = -1
            _ = v39
        } else {
            v40 := GenS3{F0: 4294967295, F1: (^(v3)), F2: 42}
            _ = v40
        }
        var v41 int = (fi38 % ((fi38) | 1))
        var v42 uint64 = 488
        _ = fi38
        _ = v41
        _ = v42
    }
    v43 := GenS4{F0: (3.14159 < (float64((float32(1e10) - v4)) / 3.14159))}
    for fi44 := 0; fi44 < 19; fi44++ {
        v1 = v4
        v4 += v1
        _ = fi44
    }
    if len(v37) >= 2 {
        v37 = v37[1:len(v37)]
    }
    var v45 int32 = -550
    _ = len(v37)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v43
    _ = v45
    _ = v5
}
