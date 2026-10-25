package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int32; F1 int32 }
func genFn0(p0 int64, p1 uint32, p2 bool) uint8 {
    p0 += 230
    var v6 string = "foo bar"
    var v7 string = v6
    fmt.Println(p0, p1, p1)
    v6 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = 1e10
    v3 := make(map[uint32]bool)
    v4 := []string{"\n\t"}
    var v5 float32 = -2.25
    _ = len(v3)
    v2 = v2
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(v5)
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: -2147483648}
            fmt.Println(v2, fi8)
            if false {
                switch "a" {
                case "a":
                    v1 = 983
                case "foo bar":
                    if ("" != "") {
                        if false {
                            if false {
                                var v11 float64 = 3.14159
                                _ = v11
                            }
                        } else {
                            var v12 uint64 = v1
                            _ = len(v3)
                            _ = v12
                        }
                        v5 -= v2
                    } else {
                        if false {
                            if false {
                                v2 = (((v2 + (v5 + float32(((0.1 / (0.0 / ((((0.1 * (0.0 * 1e10)) + -2.25) / (-2.25 + ((float64(v5) - -2.25) - (((3.14159 - 0.1) / 1.0) * 0.1)))) + 0.1))) + (3.14159 - ((1.5 * -2.25) + 3.14159)))))) + v2) / v5)
                                var v13 float64 = ((((1e10 - ((((float64(float32(0.1)) / 1.5) + -2.25) * 1.5) + 3.14159)) / 3.14159) + 1.5) + ((1.5 + -2.25) * ((3.14159 / 1e10) + (1e10 + 0.1))))
                                _ = v13
                            }
                            v2 = v2
                        }
                    }
                case "\n\t":
                    fmt.Println(fi9, v2, v2)
                default:
                    v14 := GenS1{F0: 42}
                    _ = v14
                }
                v5 = v5
                switch ("a" + "\n\t") {
                case "\n\t":
                    v15 := GenS1{F0: 1}
                    _ = v15
                case "hello":
                    if true {
                        var v16 bool = true
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v17 string = ("" + "hello")
                        _ = v16
                        _ = v17
                    } else {
                        v4 = append(v4, ("hello" + ("hello" + (("hello" + "hello") + (("\n\t" + "a") + "a")))))
                        v18 := GenS2{F0: -745, F1: (v1 << ((4294967295) & 63)), F2: (-1 < -693)}
                        _ = v18
                    }
                }
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v19 uint64 = (v1 * (v1 & (v1 % (((v1 - (v1 / ((v1) | 1)))) | 1))))
    var v20 bool = (false && false)
    var v21 uint8 = 255
    fmt.Println(v20)
    if v20 {
        if (1e10 > 0.0) {
            var v22 bool = (1 <= -1)
            _ = v22
        } else {
            v5 -= 1.5
            fmt.Println(v19)
            v23 := GenS1{F0: 1}
            _ = v23
        }
        v19 = v19
    } else {
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := GenS3{F0: 440, F1: -578}
            var v26 int64 = 730
            v27 := GenS1{F0: 42}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        if ((65535 == 64) && v20) {
            v28 := v3[113]
            v29 := GenS1{F0: 2147483647}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v28
            _ = v29
        } else {
            for fi30 := 0; fi30 < 14; fi30++ {
                if (65535 == 255) {
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v32 := GenS1{F0: 42}
                        delete(v3, 65535)
                        _ = fi31
                        _ = v32
                    }
                    v19 = (v1 % (((v19 ^ v1)) | 1))
                }
                if v20 {
                    switch (fi30 | (fi30 ^ 1)) {
                    case 127:
                        delete(v3, 2147483647)
                    }
                    v33 := GenS1{F0: 1}
                    _ = v33
                }
                _ = fi30
            }
        }
        var v34 uint64 = (v19 / (((v1 >> ((18446744073709551615) & 63))) | 1))
        _ = v34
    }
    var v35 uint = 255
    v1 = (v1 / ((v19) | 1))
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        fmt.Println(v1, v2, v1)
    default:
        var v36 int8 = 0
        _ = v36
    }
    if v20 {
        if (v21 >= genFn0(0, 186, v20)) {
            v37 := GenS4{F0: -2147483648, F1: 1}
            if v20 {
                var v38 string = ("a" + "")
                fmt.Println(v21, v35, v5)
                _ = v38
            }
            _ = v37
        }
    } else {
        var v39 int = 985
        var v40 float32 = (v5 - (((v5 + -2.25) * 1e10) + v2))
        var v41 uint32 = 65535
        _ = v39
        _ = v40
        _ = v41
    }
    switch ("\n\t" + ("hello" + ("hello" + ("foo bar" + "a")))) {
    case "\n\t":
        var v42 int32 = 823
        _ = v42
    default:
        if (128 <= -2147483648) {
            fmt.Println(v19, v19)
        }
    }
    v20 = ((!v20) && v20)
    v2 += v5
    var v43 uint16 = 255
    var v44 string = ("" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
