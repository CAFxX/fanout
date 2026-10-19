package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 int8; F2 int }
type GenS9 struct { F0 bool; F1 int32; F2 int32 }
type GenS10 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint16, p1 uint16) int8 {
    var v5 int8 = -69
    var v6 int32 = 42
    for fi7 := 0; fi7 < 5; fi7++ {
        p1 = (p1 | (^(p0)))
        var v8 int32 = (v6 / ((-1) | 1))
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (v5 / ((v5) | 1))
}
func genFn1(p0 int16, p1 uint32) uint16 {
    if (1 > 128) {
        v9 := []int{994, 42}
        p1 += 2
        _ = v9
    } else {
        p0 = p0
    }
    p0 = 643
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"foo bar"}
    var v3 uint16 = 0
    v4 := make(map[int]uint32)
    var v10 uint32 = 2147483647
    _ = len(v2)
    v11 := GenS1{F0: true, F1: -9223372036854775808}
    var v12 uint = 0
    v13 := GenS1{F0: ((0 == 9223372036854775807) && (1 < uint16(v10))), F1: 42}
    v2 = append(v2, (("hello" + "foo bar") + "hello"))
    v14 := v4[-961]
    v12 -= (v12 ^ v12)
    if (((3.14159 - 1.5) * 1e10) >= (1.5 * 0.0)) {
        var v15 uint32 = v10
        var v16 float32 = 0.1
        _ = v15
        _ = v16
    } else {
        v17 := GenS2{F0: ("foo bar" + "foo bar"), F1: v12}
        _ = v17
    }
    if false {
        if true {
            v2 = append(v2, "")
        }
        switch 9223372036854775807 {
        case -831:
            for fi18 := 0; fi18 < 8; fi18++ {
                v3 -= genFn1(-143, v10)
                var v19 float32 = ((((1e10 * (1.5 + 0.1)) * (0.1 - 3.14159)) - float32(float64(3.14159))) - (((1e10 * -2.25) - 0.1) - 1e10))
                v14 -= (v10 | v10)
                _ = fi18
                _ = v19
            }
        case 4294967295:
            var v20 float64 = 1.5
            _ = v20
        }
    } else {
        switch "" {
        case "foo bar":
            for fi21 := 0; fi21 < 3; fi21++ {
                v12 = (^(v12))
                v22 := GenS3{F0: 0, F1: v12, F2: (^(v10))}
                _ = fi21
                _ = v22
            }
        case "\n\t":
            var v23 int64 = 9223372036854775807
            _ = v23
        case "":
            var v24 int64 = -1
            _ = v24
        default:
            if true {
                for fi25 := 0; fi25 < 8; fi25++ {
                    if (42 != 587) {
                        var v26 uint = uint(v10)
                        switch fi25 {
                        case 361:
                            if (1 != 42) {
                                switch 9223372036854775807 {
                                case -2147483648:
                                    var v27 int8 = 63
                                    _ = v27
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                if ((-2.25 <= -2.25) && (0.0 < 1.5)) {
                                    var v28 int8 = -1
                                    fmt.Println(fi25)
                                    fmt.Println(v26)
                                    _ = v28
                                } else {
                                    for fi29 := 0; fi29 < 15; fi29++ {
                                        v30 := GenS4{F0: (("foo bar" + (("\n\t" + (("foo bar" + "\n\t") + "a")) + "\n\t")) + ("a" + ""))}
                                        switch (("\n\t" + "\n\t") + (("\n\t" + ((("\n\t" + "\n\t") + "hello") + "foo bar")) + "foo bar")) {
                                        case "hello":
                                            if (!false) {
                                                var v31 uint = (v12 + uint(v14))
                                                _ = v31
                                            } else {
                                                v32 := GenS5{F0: 2, F1: 99}
                                                _ = v32
                                            }
                                            if false {
                                                v26 -= 65535
                                            } else {
                                                v10 = v10
                                                var v33 float64 = (1e10 - -2.25)
                                                var v34 int64 = -218
                                                fmt.Println(v26, v34)
                                                switch "hello" {
                                                case "foo bar":
                                                    var v35 int8 = 42
                                                    _ = v35
                                                case "":
                                                    fmt.Println(v12)
                                                case "a":
                                                    fmt.Println(v33, v34)
                                                default:
                                                    v36 := GenS6{F0: 1}
                                                    _ = v36
                                                }
                                                _ = v33
                                                _ = v34
                                            }
                                        case "foo bar":
                                            v14 += (v14 ^ v10)
                                        }
                                        v37 := GenS7{F0: fi25}
                                        _ = fi29
                                        _ = v30
                                        _ = v37
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v38 := GenS8{F0: ((v3 ^ (v3 | (v3 / (((v3 * v3)) | 1)))) < genFn1(32767, v14)), F1: 100, F2: (fi25 | fi25)}
                                    _ = v38
                                }
                            }
                        case 128:
                            fmt.Println(v10, v14)
                        case 4294967295:
                            v2 = append(v2, "hello")
                        }
                        _ = v26
                    } else {
                        if (!true) {
                            v39 := GenS2{F0: "", F1: 4294967295}
                            if len(v2) > 0 {
                                v2[0] = (("\n\t" + "") + ("foo bar" + "foo bar"))
                            }
                            _ = v39
                        } else {
                            var v40 bool = (true || true)
                            v41 := GenS6{F0: -128}
                            v42 := GenS9{F0: (2147483647 > (fi25 & (^(fi25)))), F1: int32(fi25), F2: -313}
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        v43 := GenS10{F0: 2, F1: -591}
                        _ = v43
                    }
                    var v44 uint = (^(v12))
                    _ = fi25
                    _ = v44
                }
            }
        }
        var v45 uint16 = 65535
        _ = v45
    }
    var v46 string = ("hello" + "\n\t")
    fmt.Println(v46)
    v10 -= v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v46
}
