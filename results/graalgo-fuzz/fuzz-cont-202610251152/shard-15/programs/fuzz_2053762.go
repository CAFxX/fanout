package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int32; F2 int8 }
func genFn0(p0 int16, p1 bool) int {
    var v5 int8 = 42
    v5 -= v5
    v6 := GenS1{F0: 42, F1: -1, F2: true}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int8, p1 bool, p2 int) uint {
    if p1 {
        fmt.Println(p0, p2)
        var v7 uint = 587
        var v8 int = p2
        _ = v7
        _ = v8
    }
    v9 := []int{}
    v10 := GenS2{F0: 42, F1: (224 == 194)}
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 uint32 = 1
        if p1 {
            switch 32767 {
            case 127:
                var v13 int64 = -9223372036854775808
                _ = v13
            case -2147483648:
                if p1 {
                    switch ((("foo bar" + ("\n\t" + "\n\t")) + "hello") + "a") {
                    case "a":
                        var v14 string = "foo bar"
                        _ = v14
                    case "hello":
                        var v15 uint8 = 2
                        _ = v15
                    }
                    var v16 int16 = 0
                    v17 := GenS1{F0: 1, F1: v16, F2: p1}
                    _ = v16
                    _ = v17
                } else {
                    v12 = v12
                    if (!(p1 && (!(p1 && p1)))) {
                        if false {
                            switch "\n\t" {
                            case "hello":
                                switch -2147483648 {
                                case -622:
                                    var v18 int32 = 0
                                    _ = v18
                                case 32768:
                                    p0 += (p0 - 1)
                                default:
                                    fmt.Println(p1, p2)
                                }
                            case "\n\t":
                                switch int(p0) {
                                case 1:
                                    switch "" {
                                    case "hello":
                                        v19 := GenS3{F0: 182}
                                        _ = v19
                                    default:
                                        var v20 uint = 255
                                        _ = v20
                                    }
                                case 9223372036854775807:
                                    v12 += v12
                                case 698:
                                    v21 := GenS1{F0: 226, F1: 1, F2: true}
                                    _ = v21
                                default:
                                    if (127 > 164) {
                                        if p1 {
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            p0 = p0
                                        }
                                    } else {
                                        if p1 {
                                            var v22 float64 = (0.1 / (0.1 / float64(-2.25)))
                                            switch "hello" {
                                            case "a":
                                                p2 += -1
                                            }
                                            _ = v22
                                        } else {
                                            var v23 uint32 = v12
                                            fmt.Println(v12, p0, fi11)
                                            _ = v23
                                        }
                                    }
                                }
                            }
                            v24 := GenS2{F0: 0, F1: (p1 || ((65535 < 1) && ((false || (p1 || p1)) && p1)))}
                            _ = v24
                        } else {
                            p0 = 61
                            p2 = (p2 & genFn0(-32768, (458 > 745)))
                        }
                        for fi25 := 0; fi25 < 18; fi25++ {
                            fmt.Println(p1, v12)
                            _ = fi25
                        }
                    }
                    if p1 {
                        var v26 uint64 = 1
                        _ = v26
                    } else {
                        if true {
                            var v27 int32 = -120
                            v28 := GenS2{F0: 896, F1: false}
                            _ = v27
                            _ = v28
                        } else {
                            var v29 int16 = 32767
                            if (-776 < p2) {
                                v30 := GenS3{F0: uint8(v12)}
                                _ = v30
                            } else {
                                if len(v9) > 0 {
                                    v9[0] = (p2 + (p2 << ((18446744073709551615) & 63)))
                                }
                                _ = len(v9)
                            }
                            for fi31 := 0; fi31 < 20; fi31++ {
                                fmt.Println(v12, p2)
                                var v32 float32 = 0.1
                                if (((1e10 - -2.25) * 0.0) >= 1.5) {
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                }
                                v29 += (-(v29))
                                if true {
                                    var v33 uint64 = 42
                                    p1 = (167 >= (^(v12)))
                                    v34 := GenS2{F0: 255, F1: p1}
                                    _ = v33
                                    _ = v34
                                } else {
                                    v32 += v32
                                    fmt.Println(fi11, v32)
                                }
                                _ = fi31
                                _ = v32
                            }
                            _ = v29
                        }
                        var v35 float64 = 3.14159
                        _ = v35
                    }
                }
            case 4294967295:
                p0 -= 127
            }
        } else {
            switch -9223372036854775808 {
            case -9223372036854775808:
                if len(v9) > 0 {
                    v9[0] = p2
                }
            default:
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            }
            if (850 == 1) {
                switch (-(fi11)) {
                case 42:
                    v12 -= v12
                }
            }
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(p0, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 4294967295
}

func main() {
    var v1 int = 555
    v2 := make(map[string]int)
    var v3 bool = (("foo bar" + ("foo bar" + "a")) < "\n\t")
    v4 := make(map[string]uint32)
    fmt.Println(v1)
    if v3 {
        v36 := GenS3{F0: 128}
        if true {
            var v37 uint = 18446744073709551615
            v38 := []int{(v1 ^ genFn0(42, v3)), v1, (v1 % ((v1) | 1))}
            for fi39 := 0; fi39 < 1; fi39++ {
                if false {
                    v40 := GenS2{F0: 120, F1: (true || (genFn1(-1, v3, fi39) > genFn1(-128, (1.5 < -2.25), v1)))}
                    for fi41 := 0; fi41 < 17; fi41++ {
                        var v42 uint8 = 0
                        _ = fi41
                        _ = v42
                    }
                    _ = v40
                } else {
                    v43 := GenS4{F0: (v37 >> ((v37) & 63)), F1: -647, F2: 100}
                    _ = v43
                }
                switch (fi39 / ((genFn0(42, v3)) | 1)) {
                case 0:
                    v3 = v3
                case 127:
                    var v44 int16 = 777
                    _ = v44
                case 4294967295:
                    switch (fi39 ^ fi39) {
                    case -164:
                        _ = len(v38)
                    default:
                        var v45 uint16 = 1
                        _ = v45
                    }
                }
                v46 := GenS1{F0: 441, F1: -32768, F2: v3}
                _ = fi39
                _ = v46
            }
            _ = v37
            _ = v38
        }
        v47 := v2[("a" + ("" + "\n\t"))]
        _ = v36
        _ = v47
    }
    v4["foo bar"] = 681
    v1 -= v1
    v3 = false
    var v48 uint64 = 4294967295
    if ((v1 * v1) < v1) {
        if v3 {
            delete(v2, "a")
            v48 = v48
        } else {
            v49 := []int{v1, v1}
            _ = v49
        }
    }
    v1 += (v1 / (((v1 - v1)) | 1))
    var v50 uint8 = 2
    var v51 float32 = ((0.1 - 0.1) / -2.25)
    if (4294967295 > 877) {
        v3 = v3
    }
    v52 := GenS2{F0: 0, F1: false}
    fmt.Println(v50, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v50
    _ = v51
    _ = v52
}
