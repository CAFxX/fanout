package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint; F1 int32; F2 bool }
type GenS4 struct { F0 int16 }
func genFn0(p0 bool, p1 int64) string {
    fmt.Println(p1)
    switch ("foo bar" + "\n\t") {
    case "hello":
        for fi7 := 0; fi7 < 20; fi7++ {
            p0 = false
            _ = fi7
        }
    case "foo bar":
        fmt.Println(p1, p1)
    }
    var v8 int16 = 32767
    _ = p0
    _ = p1
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    var v3 int32 = -1
    v4 := make(map[uint32]bool)
    var v5 uint32 = 42
    var v6 int64 = -489
    v9 := GenS1{F0: 128, F1: 32767, F2: 42}
    v10 := []int{32768, -2147483648, 2147483647, 305}
    var v11 int16 = 550
    if (!(724 > 65535)) {
        var v12 int8 = -119
        v13 := GenS1{F0: 127, F1: -9223372036854775808, F2: 837}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 string = ("a" + (("a" + (("\n\t" + "") + (("" + (("a" + "hello") + (((("\n\t" + ("" + "")) + "foo bar") + "\n\t") + ("hello" + "")))) + ("hello" + ("hello" + "hello"))))) + ("hello" + "\n\t")))
        v3 = -2147483648
        v3 -= -560
        _ = fi14
        _ = v15
    }
    var v16 uint = 4294967295
    var v17 int32 = 1
    var v18 float64 = float64(3.14159)
    v18 = v18
    v19 := GenS2{F0: v5, F1: v11, F2: 42}
    if (107 <= 255) {
        v3 += (v17 / (((v17 ^ 986)) | 1))
        var v20 float32 = (3.14159 - ((-2.25 / -2.25) + ((1.5 / (0.1 + 3.14159)) / ((-2.25 + ((-2.25 / (float32((-2.25 + v18)) * 1.5)) * (float32(float64(1.5)) - (-2.25 - 1e10)))) * (((((((3.14159 + float32(v18)) + 0.1) - 3.14159) / 0.1) / ((3.14159 / (float32((((float64(3.14159) * 0.1) / float64(1e10)) * v18)) - ((-2.25 / float32(v18)) - (((3.14159 + 0.1) / 3.14159) + 3.14159)))) - 1.5)) - ((1.5 - 0.0) - ((3.14159 * -2.25) + 3.14159))) / 1.0)))))
        if false {
            fmt.Println(v20)
            var v21 uint = (v16 << (((v16 >> ((2) & 63))) & 63))
            _ = v21
        } else {
            v5 = v5
            fmt.Println(v17)
        }
        _ = v20
    } else {
        var v22 uint64 = 18446744073709551615
        _ = v22
    }
    var v23 uint64 = 18446744073709551615
    v24 := GenS1{F0: 128, F1: -1, F2: 42}
    switch 32768 {
    case 2147483647:
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(v5)
            v5 = 4294967295
            _ = fi25
        }
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        if false {
            fmt.Println(v6, v18, v6)
            if false {
                fmt.Println(v16, v16)
                v10 = append(v10, fi26)
            }
        } else {
            fmt.Println(v5)
            if true {
                fmt.Println(fi26, v6, v23)
            }
            for fi27 := 0; fi27 < 19; fi27++ {
                if false {
                    switch ("foo bar" + ("" + "a")) {
                    case "hello":
                        var v28 uint = 18446744073709551615
                        v6 = -390
                        v16 -= v16
                        _ = v28
                    default:
                        if true {
                            v29 := GenS3{F0: (v16 << (((v16 % ((665) | 1))) & 63)), F1: v3, F2: false}
                            if false {
                                if true {
                                    v18 = 1.5
                                } else {
                                    v6 += v6
                                    var v30 uint16 = 84
                                    _ = v30
                                }
                                if true {
                                    var v31 float64 = v18
                                    v6 -= v6
                                    v3 -= v3
                                    if true {
                                        var v32 bool = false
                                        if (!(!true)) {
                                            v31 = v31
                                        } else {
                                            var v33 bool = true
                                            switch (-(fi26)) {
                                            case 127:
                                                v4[v5] = (!(549 != 706))
                                            }
                                            v32 = v33
                                            _ = v33
                                        }
                                        _ = v32
                                    }
                                    v3 = (v3 / ((v3) | 1))
                                    _ = v31
                                }
                            } else {
                                if true {
                                    var v34 float32 = -2.25
                                    _ = v34
                                } else {
                                    fmt.Println(v18, v5, v3)
                                }
                                fmt.Println(v3)
                                switch "\n\t" {
                                case "hello":
                                    switch (("\n\t" + ("" + "\n\t")) + "") {
                                    case "":
                                        var v35 uint = 2
                                        _ = v35
                                    case "foo bar":
                                        var v36 int = int(v23)
                                        _ = v36
                                    case "a":
                                        v18 += v18
                                    default:
                                        switch fi27 {
                                        case 4294967295:
                                            var v37 float32 = (0.0 - (((1.5 + 3.14159) - 1.5) / float32((((1e10 + (v18 / ((((0.0 / v18) / (((1.5 + -2.25) - (1e10 + (float64(-2.25) + (v18 - (v18 * v18))))) / float64(3.14159))) / 0.1) * (v18 - float64(3.14159))))) / v18) - v18))))
                                            _ = v37
                                        default:
                                            fmt.Println(v11)
                                        }
                                    }
                                    fmt.Println(fi26, v11, v23)
                                case "":
                                    v18 = (v18 * (v18 - float64((-2.25 + 0.1))))
                                case "a":
                                    v38 := GenS4{F0: (v11 % ((1000) | 1))}
                                    _ = v38
                                }
                            }
                            if true {
                                v18 += (v18 * -2.25)
                                switch ("\n\t" + "hello") {
                                case "\n\t":
                                    v39 := GenS4{F0: (v11 % ((v11) | 1))}
                                    _ = v39
                                default:
                                    v17 += -2147483648
                                }
                            } else {
                                _ = len(v1)
                                var v40 uint64 = (^(v23))
                                var v41 uint8 = 255
                                _ = v40
                                _ = v41
                            }
                            _ = v29
                        }
                    }
                    var v42 uint32 = 1
                    _ = v42
                } else {
                    v3 -= (v17 % ((v17) | 1))
                    var v43 int32 = 463
                    _ = v43
                }
                v44 := GenS2{F0: v5, F1: 0, F2: 32767}
                switch 32768 {
                case 452:
                    v16 -= v16
                case 4294967295:
                    v11 += v11
                case -67:
                    fmt.Println(v3, v17)
                default:
                    if (256 > 591) {
                        switch fi27 {
                        case 162:
                            var v45 int32 = 2147483647
                            _ = v45
                        case 479:
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        default:
                            v46 := GenS4{F0: v11}
                            _ = v46
                        }
                    }
                }
                _ = fi27
                _ = v44
            }
        }
        v47 := GenS3{F0: v16, F1: v17, F2: ((828 >= 65535) && true)}
        _ = fi26
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
