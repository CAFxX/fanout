package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 bool; F1 uint64 }

func main() {
    var v1 uint8 = 145
    var v2 uint64 = 2
    var v3 int8 = -128
    v4 := make(map[int]bool)
    var v5 uint8 = uint8(v2)
    fmt.Println(v2)
    if (true && true) {
        delete(v4, 9223372036854775807)
        v4[42] = (!true)
        for fi6 := 0; fi6 < 1; fi6++ {
            if (!true) {
                for fi7 := 0; fi7 < 8; fi7++ {
                    v8 := GenS1{F0: 439}
                    if (!(335 != 255)) {
                        v9 := v4[-2147483648]
                        var v10 int16 = 1
                        _ = v10
                        _ = v9
                    } else {
                        if (!false) {
                            if (3.14159 > (1.5 / (float32(-2.25) / 0.1))) {
                                if false {
                                    v11 := GenS2{F0: (v3 | (-(v3))), F1: 42}
                                    var v12 uint = 621
                                    v4[fi7] = (v12 >= 42)
                                    v1 = (v1 * (^(v1)))
                                    _ = v11
                                    _ = v12
                                } else {
                                    switch ((("foo bar" + (("hello" + ("" + (("a" + ("" + ("foo bar" + "a"))) + ("hello" + ("" + "hello"))))) + "\n\t")) + "\n\t") + "foo bar") {
                                    case "a":
                                        v13 := []int{fi6, fi6}
                                        _ = v13
                                    case "\n\t":
                                        var v14 string = ("hello" + "")
                                        _ = v14
                                    default:
                                        _ = len(v4)
                                    }
                                    switch (fi6 / (((-(fi7))) | 1)) {
                                    case 1:
                                        switch 42 {
                                        case 4294967295:
                                            v15 := []int{-379}
                                            _ = v15
                                        case 1:
                                            var v16 uint = 65535
                                            _ = v16
                                        case -1:
                                            v1 = (v1 >> ((978) & 63))
                                        }
                                        v3 -= (v3 | 113)
                                        v17 := []int{fi7, fi7}
                                        _ = v17
                                    case -1:
                                        v2 = (v2 * v2)
                                    default:
                                        var v18 int32 = 454
                                        _ = v18
                                    }
                                    var v19 int8 = v3
                                    v3 += 26
                                    switch (fi7 << ((515) & 63)) {
                                    case -237:
                                        var v20 uint8 = v1
                                        _ = v20
                                    case -351:
                                        fmt.Println(fi7, v5, fi6)
                                    case -2147483648:
                                        var v21 float64 = (-2.25 / (-2.25 * (1.5 / (1e10 * (-2.25 - 1.5)))))
                                        _ = v21
                                    default:
                                        fmt.Println(v5)
                                        var v22 int64 = 9223372036854775807
                                        if (fi7 != fi7) {
                                            if true {
                                                switch "a" {
                                                case "hello":
                                                    v23 := []int{(fi6 / (((fi7 | fi7)) | 1))}
                                                    _ = v23
                                                }
                                                var v24 uint64 = v2
                                                var v25 int8 = v3
                                                _ = v24
                                                _ = v25
                                            }
                                        }
                                        _ = v22
                                    }
                                    _ = v19
                                }
                                var v26 string = "\n\t"
                                if false {
                                    if (958 == (fi6 + (fi7 >> ((4294967295) & 63)))) {
                                        v27 := GenS3{F0: false, F1: (v2 << ((18446744073709551615) & 63))}
                                        delete(v4, int(v3))
                                        _ = v27
                                    }
                                } else {
                                    if true {
                                        fmt.Println(v2, fi7, v3)
                                    }
                                }
                                _ = v26
                            } else {
                                var v28 int64 = 0
                                if true {
                                    v2 = (v2 - v2)
                                    v1 += (v5 << ((576) & 63))
                                    var v29 float32 = float32(float64(((1.5 * (-2.25 / (float32(((0.1 + float64(1.5)) * (float64(-2.25) * 0.1))) - 3.14159))) / 1.0)))
                                    _ = v29
                                }
                                _ = v28
                            }
                            fmt.Println(v2, fi6, fi6)
                            if (!true) {
                                var v30 uint = 1
                                _ = v30
                            } else {
                                if true {
                                    var v31 uint8 = v1
                                    _ = v31
                                }
                                var v32 uint8 = v1
                                var v33 float32 = 1e10
                                _ = v32
                                _ = v33
                            }
                        } else {
                            var v34 string = "hello"
                            var v35 string = (v34 + (v34 + v34))
                            _ = v34
                            _ = v35
                        }
                        var v36 int32 = -600
                        _ = v36
                    }
                    _ = fi7
                    _ = v8
                }
                for fi37 := 0; fi37 < 19; fi37++ {
                    fmt.Println(v3, v3, v1)
                    v1 = 2
                    _ = fi37
                }
                fmt.Println(v1, v3)
            }
            v5 = 1
            v2 += 4294967295
            _ = fi6
        }
    } else {
        v38 := []int{32768, -2147483648, 0, 128}
        if len(v38) >= 2 {
            v38 = v38[1:len(v38)]
        }
        switch ("a" + "") {
        case "a":
            v3 = (v3 | (v3 * (^(v3))))
        case "\n\t":
            if len(v38) >= 2 {
                v38 = v38[1:len(v38)]
            }
        case "hello":
            v5 = v1
        }
        _ = v38
    }
    fmt.Println(v3)
    fmt.Println(v2)
    if true {
        v39 := v4[235]
        v3 -= (v3 - (v3 << ((42) & 63)))
        _ = v39
    } else {
        var v40 string = "foo bar"
        _ = v40
    }
    v5 = v5
    var v41 float64 = 3.14159
    _ = len(v4)
    var v42 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
