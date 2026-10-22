package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 string; F1 uint32; F2 int16 }
type GenS4 struct { F0 string; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 int64 }
func genFn0(p0 int32) uint8 {
    v7 := []int{}
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            p0 = 249
            var v10 uint32 = 132
            switch fi8 {
            case 128:
                if (("\n\t" + ("foo bar" + "")) != (("\n\t" + "\n\t") + "\n\t")) {
                    v11 := GenS1{F0: "foo bar", F1: 408, F2: 32767}
                    _ = v11
                } else {
                    switch fi8 {
                    case 9223372036854775807:
                        if (false || true) {
                            var v12 uint16 = 42
                            var v13 int16 = -32768
                            _ = v12
                            _ = v13
                        } else {
                            if false {
                                switch 2147483647 {
                                case -9223372036854775808:
                                    var v14 float64 = float64(1e10)
                                    _ = v14
                                case 42:
                                    fmt.Println(p0)
                                default:
                                    fmt.Println(p0, p0, fi9)
                                }
                                v10 = 2147483648
                                v10 -= (^(v10))
                            }
                            v15 := GenS2{F0: 32767, F1: -1, F2: 42}
                            if ((3.14159 / (((1.5 - float64(0.0)) / float64(-2.25)) - (-2.25 * (-2.25 / (float64(0.1) * float64(((float32(3.14159) + 0.0) - -2.25))))))) <= 1.5) {
                                switch (("hello" + ("foo bar" + "hello")) + ("foo bar" + "a")) {
                                case "":
                                    v7 = append(v7, (fi9 / ((660) | 1)))
                                default:
                                    var v16 uint64 = 4294967295
                                    fmt.Println(v10)
                                    switch 32767 {
                                    case -135:
                                        var v17 uint8 = 0
                                        _ = v17
                                    case 772:
                                        fmt.Println(p0, v10)
                                    default:
                                        v18 := GenS3{F0: "foo bar", F1: (v10 & v10), F2: -32768}
                                        _ = v18
                                    }
                                    _ = v16
                                }
                                switch "foo bar" {
                                case "a":
                                    if len(v7) > 0 {
                                        v7[0] = -161
                                    }
                                default:
                                    if len(v7) > 0 {
                                        v7[0] = 128
                                    }
                                }
                            }
                            v19 := GenS2{F0: 255, F1: -1, F2: 0}
                            fmt.Println(v10)
                            if ((128 >= 41) && (2 < 877)) {
                                var v20 uint8 = 2
                                fmt.Println(v10, v10)
                                if len(v7) > 0 {
                                    v7[0] = fi9
                                }
                                fmt.Println(p0, v20, fi9)
                                p0 = (^(p0))
                                var v21 uint = 1
                                _ = v20
                                _ = v21
                            } else {
                                if true {
                                    v22 := GenS4{F0: "foo bar", F1: 565, F2: 2}
                                    _ = v22
                                } else {
                                    v10 -= 1
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    fmt.Println(fi8, fi8, v10)
                                }
                                fmt.Println(p0)
                                if (!false) {
                                    switch "hello" {
                                    case "\n\t":
                                        p0 += -407
                                    case "":
                                        var v23 uint64 = 1
                                        _ = v23
                                    case "hello":
                                        var v24 bool = (!(((1.5 - (0.1 + 3.14159)) / 1.0) == -2.25))
                                        _ = v24
                                    default:
                                        if true {
                                            p0 -= -2147483648
                                            p0 += p0
                                            var v25 uint16 = 2
                                            _ = v25
                                        }
                                    }
                                    if (!(0 < 729)) {
                                        v10 = (v10 + (^(v10)))
                                    }
                                    switch "foo bar" {
                                    case "":
                                        var v26 int64 = -9223372036854775808
                                        _ = v26
                                    default:
                                        fmt.Println(fi8, fi8)
                                        if (!true) {
                                            var v27 uint32 = 2147483648
                                            var v28 int16 = 652
                                            _ = v27
                                            _ = v28
                                        }
                                        var v29 uint16 = 32767
                                        p0 = 857
                                        fmt.Println(p0, v29)
                                        _ = v29
                                    }
                                }
                            }
                            _ = v15
                            _ = v19
                        }
                    case -1:
                        if true {
                            v10 = v10
                        } else {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            var v30 uint64 = 206
                            _ = v30
                        }
                    case 127:
                        var v31 uint16 = 2
                        _ = v31
                    }
                    if false {
                        v32 := GenS4{F0: ("a" + "\n\t"), F1: 32767, F2: 326}
                        v10 -= v10
                        if len(v7) > 0 {
                            v7[0] = (fi9 * 336)
                        }
                        _ = v32
                    }
                }
            }
            fmt.Println(fi9, v10)
            v33 := GenS5{F0: 660}
            p0 = 1
            if len(v7) > 0 {
                v7[0] = 32767
            }
            _ = fi9
            _ = v10
            _ = v33
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    p0 = 354
    _ = p0
    _ = v7
    return 127
}

func main() {
    var v1 uint16 = 335
    var v2 bool = (false || false)
    var v3 float64 = -2.25
    v4 := []string{"a", ("\n\t" + "a")}
    v5 := make(map[int]int)
    v6 := make(map[string]int)
    v2 = v2
    v34 := v5[128]
    var v35 uint8 = 127
    v34 = (-(v34))
    v35 = v35
    v34 += v34
    v36 := GenS1{F0: "foo bar", F1: 1, F2: -1}
    v1 = (v1 * v1)
    if (1000 > -527) {
        var v37 uint8 = 0
        _ = v37
    } else {
        v38 := GenS5{F0: (^(v1))}
        _ = v38
    }
    delete(v5, 42)
    var v39 int64 = 9223372036854775807
    switch (v34 % ((v34) | 1)) {
    case 127:
        if ((float32(v3) + 1.5) == 1e10) {
            delete(v6, "a")
            switch (("\n\t" + "\n\t") + (((("foo bar" + "a") + (("" + "") + "foo bar")) + "hello") + "")) {
            case "":
                for fi40 := 0; fi40 < 3; fi40++ {
                    if v2 {
                        var v41 int8 = 127
                        switch ("hello" + "") {
                        case "":
                            v42 := GenS1{F0: "hello", F1: 767, F2: 32767}
                            _ = v42
                        case "foo bar":
                            for fi43 := 0; fi43 < 1; fi43++ {
                                v44 := GenS5{F0: (v1 % ((v1) | 1))}
                                v45 := GenS6{F0: v35, F1: (v39 + (v39 >> ((65535) & 63)))}
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                        default:
                            var v46 string = ("a" + ("hello" + ""))
                            _ = v46
                        }
                        _ = v41
                    } else {
                        v47 := GenS3{F0: "foo bar", F1: 2147483647, F2: -32768}
                        _ = v47
                    }
                    v34 = fi40
                    var v48 bool = (v2 || true)
                    _ = fi40
                    _ = v48
                }
            case "foo bar":
                delete(v5, (-(v34)))
            default:
                var v49 uint32 = 228
                _ = v49
            }
        }
    case 73:
        switch ("a" + "a") {
        case "hello":
            v3 += v3
        case "":
            var v50 float64 = v3
            _ = v50
        case "a":
            v35 = (v35 | (v35 & (v35 << ((42) & 63))))
        }
    case -146:
        var v51 int32 = -1
        _ = v51
    default:
        var v52 uint = 503
        _ = v52
    }
    switch 127 {
    case -605:
        v4 = append(v4, "foo bar")
    case 146:
        for fi53 := 0; fi53 < 6; fi53++ {
            for fi54 := 0; fi54 < 5; fi54++ {
                v3 = -2.25
                v1 += v1
                _ = fi54
            }
            for fi55 := 0; fi55 < 12; fi55++ {
                v34 -= fi55
                var v56 float32 = 0.1
                _ = fi55
                _ = v56
            }
            v2 = (819 >= 1)
            _ = fi53
        }
    }
    if false {
        fmt.Println(v2, v35, v35)
    }
    v57 := GenS4{F0: (("foo bar" + ("hello" + "hello")) + "\n\t"), F1: (v1 >> ((42) & 63)), F2: 42}
    var v58 int = v34
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
