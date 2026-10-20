package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint16; F2 int32 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint8, p1 int8) uint8 {
    v5 := GenS1{F0: 127}
    p1 = p1
    v6 := GenS2{F0: 915}
    p0 -= (p0 % (((p0 - 128)) | 1))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int16) int {
    p0 = (-(p0))
    var v7 float64 = float64(3.14159)
    p0 -= (p0 * p0)
    p0 -= (p0 * -882)
    _ = p0
    _ = v7
    return 636
}

func main() {
    var v1 int = 42
    v2 := []int{(-(v1)), (-(v1)), v1}
    var v3 uint = 18446744073709551615
    var v4 uint64 = 4294967295
    if (false || (!(true && (0 >= 4294967295)))) {
        for fi8 := 0; fi8 < 8; fi8++ {
            if len(v2) > 0 {
                v2[0] = (-(fi8))
            }
            v9 := GenS2{F0: (v3 / (((v3 >> (((v3 - v3)) & 63))) | 1))}
            if (640 <= 1) {
                v1 = 32767
                v1 += (v1 % ((fi8) | 1))
            } else {
                fmt.Println(v4)
                var v10 float64 = (((float64(((0.0 * (-2.25 + 3.14159)) - (1e10 * 0.0))) - 0.0) / 1e10) + 1.5)
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        v11 := GenS3{F0: -32768, F1: 255, F2: 2147483647}
        _ = v11
    } else {
        var v12 int32 = 1
        if ((!true) && true) {
            var v13 uint = (^(v3))
            _ = v13
        }
        fmt.Println(v12, v12)
        _ = v12
    }
    if true {
        switch v1 {
        case 138:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 918:
            var v14 uint32 = 2147483647
            _ = v14
        }
        if false {
            var v15 float32 = 1.5
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v15
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v1, v1)
        }
        fmt.Println(v1)
    }
    var v16 int16 = 1000
    fmt.Println(v3, v16)
    v4 = (v4 | v4)
    if false {
        switch (("" + "hello") + ("foo bar" + (("" + "\n\t") + "\n\t"))) {
        case "":
            var v17 uint = v3
            _ = v17
        case "foo bar":
            switch (v1 - (v1 << ((v3) & 63))) {
            case 4294967295:
                for fi18 := 0; fi18 < 20; fi18++ {
                    fmt.Println(fi18, v3)
                    _ = fi18
                }
            case 884:
                var v19 float32 = -2.25
                _ = v19
            case -359:
                v4 = (v4 / ((v4) | 1))
            }
        case "hello":
            var v20 uint64 = v4
            _ = v20
        }
    } else {
        var v21 uint8 = 37
        for fi22 := 0; fi22 < 4; fi22++ {
            if true {
                v23 := GenS4{F0: (v21 & (^(v21))), F1: (v16 | (v16 | 0)), F2: 65535}
                if true {
                    fmt.Println(v3)
                    switch (("a" + "\n\t") + (("foo bar" + "\n\t") + "hello")) {
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v24 int = -806
                    _ = v24
                } else {
                    switch fi22 {
                    case 1:
                        v25 := GenS2{F0: (v3 % ((v3) | 1))}
                        _ = v25
                    case 32768:
                        var v26 uint8 = (^(v21))
                        _ = v26
                    default:
                        var v27 uint8 = (v21 << ((42) & 63))
                        _ = v27
                    }
                    var v28 uint32 = 1
                    var v29 uint64 = (^(v4))
                    _ = v28
                    _ = v29
                }
                if (((fi22 << ((1) & 63)) < fi22) && (-2.25 > (1e10 * (1.5 / 1e10)))) {
                    fmt.Println(v3, v21)
                    v16 = -32768
                    var v30 int = v1
                    _ = v30
                } else {
                    var v31 float64 = 0.0
                    _ = v31
                }
                _ = v23
            } else {
                v2 = append(v2, genFn1((v16 | 820)))
                for fi32 := 0; fi32 < 13; fi32++ {
                    if (548 >= 42) {
                        switch "hello" {
                        case "hello":
                            var v33 uint64 = v4
                            _ = v33
                        case "":
                            var v34 uint16 = 2
                            _ = v34
                        default:
                            v35 := GenS5{F0: 838, F1: -1}
                            _ = v35
                        }
                        var v36 string = "a"
                        fmt.Println(v21, v16)
                        _ = v36
                    }
                    switch (v1 | (v1 ^ fi22)) {
                    case 32768:
                        var v37 int32 = -1
                        _ = v37
                    default:
                        v2 = append(v2, genFn1(v16))
                    }
                    _ = fi32
                }
            }
            _ = fi22
        }
        switch "\n\t" {
        case "hello":
            v3 -= 18446744073709551615
        case "a":
            if (9223372036854775807 != 377) {
                fmt.Println(v1, v3)
                if (0.1 <= 3.14159) {
                    v16 -= v16
                }
            }
        case "":
            if true {
                v21 += v21
                var v38 string = ("a" + ("hello" + ((("foo bar" + ("foo bar" + ("" + ""))) + (("hello" + ("\n\t" + "hello")) + "")) + ("foo bar" + ""))))
                _ = v38
            } else {
                switch "\n\t" {
                case "\n\t":
                    v39 := GenS4{F0: v21, F1: (v16 / (((^(v16))) | 1)), F2: 65535}
                    _ = v39
                }
                v1 -= (v1 | 32767)
            }
        }
        _ = v21
    }
    v1 += 128
    var v40 bool = (v16 >= 32767)
    switch 127 {
    case -157:
        var v41 int8 = 38
        _ = v41
    case -1:
        for fi42 := 0; fi42 < 10; fi42++ {
            v3 += (^(v3))
            var v43 uint32 = 4294967295
            v43 = (v43 + 65535)
            _ = fi42
            _ = v43
        }
    case 127:
        var v44 int32 = -415
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
