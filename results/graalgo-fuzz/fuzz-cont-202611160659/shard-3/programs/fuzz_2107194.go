package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 int8 }
func genFn0(p0 int) int32 {
    v4 := GenS1{F0: 74, F1: -2147483648, F2: -89}
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS2{F0: 0, F1: 6, F2: 18446744073709551615}
        if (!true) {
            if (true || (!(!false))) {
                if false {
                    p0 -= (fi5 % ((-795) | 1))
                }
                for fi7 := 0; fi7 < 5; fi7++ {
                    var v8 int16 = -1
                    _ = fi7
                    _ = v8
                }
            }
            v9 := []int{p0, (fi5 ^ -198), p0}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 4; fi10++ {
                p0 = p0
                p0 = p0
                v11 := GenS2{F0: -662, F1: 1, F2: 61}
                _ = fi10
                _ = v11
            }
        }
        _ = fi5
        _ = v6
    }
    var v12 float64 = 0.1
    _ = p0
    _ = v12
    _ = v4
    return 369
}
func genFn1(p0 int16, p1 uint, p2 uint) uint64 {
    fmt.Println(p1, p0)
    var v13 int8 = -128
    var v14 float64 = 3.14159
    p0 = p0
    switch "foo bar" {
    case "\n\t":
        v15 := GenS3{F0: 2147483648, F1: 1, F2: (v13 * (v13 / ((v13) | 1)))}
        _ = v15
    case "hello":
        fmt.Println(v13, p1)
    default:
        var v16 int8 = (v13 << ((p1) & 63))
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 0
}

func main() {
    v1 := []uint32{70, 65535}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v17 bool = (2 >= 42)
    switch -845 {
    case -2147483648:
        if true {
            var v18 float32 = -2.25
            _ = v18
        } else {
            v17 = false
        }
    case 1:
        var v19 int16 = 1
        _ = v19
    case 2147483647:
        v17 = v17
    }
    var v20 float64 = ((1e10 + 0.1) * float64(float32(-2.25)))
    if len(v1) > 0 {
        v1[0] = 0
    }
    fmt.Println(v17)
    switch -1 {
    case 1:
        if v17 {
            if v17 {
                v20 = (0.0 - (-2.25 - 3.14159))
                fmt.Println(v20, v17)
            } else {
                v21 := GenS3{F0: 2147483648, F1: 0, F2: 94}
                _ = v21
            }
        } else {
            switch -393 {
            case 42:
                v17 = v17
            case -2147483648:
                v20 += 0.1
            case 127:
                v22 := GenS1{F0: 127, F1: 758, F2: 100}
                _ = v22
            default:
                switch 32768 {
                case 4294967295:
                    for fi23 := 0; fi23 < 5; fi23++ {
                        v20 = (v20 * (v20 + 1e10))
                        if false {
                            if false {
                                v24 := GenS2{F0: -1, F1: 195, F2: 0}
                                _ = v24
                            }
                            fmt.Println(v17, v20)
                        }
                        _ = fi23
                    }
                case 128:
                    var v25 int = 127
                    _ = v25
                case 2147483647:
                    if (!v17) {
                        fmt.Println(v20, v20)
                    } else {
                        v26 := GenS1{F0: 255, F1: 1, F2: 42}
                        if (5 > -9223372036854775808) {
                            if (float64((1e10 / (3.14159 * 3.14159))) <= v20) {
                                switch "" {
                                case "hello":
                                    var v27 int32 = 189
                                    _ = v27
                                case "foo bar":
                                    v17 = (!(!false))
                                case "":
                                    fmt.Println(v20, v17)
                                }
                                v28 := GenS4{F0: (!v17), F1: -947}
                                v29 := GenS5{F0: 2147483647}
                                _ = v28
                                _ = v29
                            }
                            v17 = v17
                        } else {
                            var v30 bool = (true && (255 <= 0))
                            _ = v30
                        }
                        v31 := GenS2{F0: -32768, F1: 665, F2: 914}
                        _ = v26
                        _ = v31
                    }
                }
            }
            switch "a" {
            case "\n\t":
                switch ("a" + "") {
                case "hello":
                    switch 1 {
                    case 128:
                        switch 32768 {
                        case -1:
                            v32 := GenS6{F0: -2147483648, F1: -6}
                            _ = v32
                        default:
                            var v33 float32 = float32(v20)
                            _ = v33
                        }
                    }
                }
            default:
                for fi34 := 0; fi34 < 5; fi34++ {
                    for fi35 := 0; fi35 < 19; fi35++ {
                        if (!(127 > -1)) {
                            var v36 int64 = 133
                            switch (fi35 | 128) {
                            case 9223372036854775807:
                                var v37 int64 = -9223372036854775808
                                _ = v37
                            }
                            _ = len(v2)
                            _ = v36
                        } else {
                            fmt.Println(fi34, fi35)
                            switch (("hello" + "a") + "hello") {
                            case "hello":
                                var v38 float32 = (float32((1e10 - (v20 * 0.0))) * (-2.25 / 1e10))
                                _ = v38
                            case "foo bar":
                                if v17 {
                                    if (!(0 <= 42)) {
                                        var v39 int32 = genFn0(fi34)
                                        switch fi35 {
                                        case 1:
                                            var v40 int16 = -1
                                            _ = v40
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v39
                                    } else {
                                        v17 = (-271 >= -156)
                                    }
                                    v1 = append(v1, 4294967295)
                                    var v41 bool = (!v17)
                                    var v42 uint32 = 65535
                                    v43 := GenS5{F0: -2147483648}
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                            case "":
                                v17 = false
                            default:
                                v44 := GenS6{F0: (fi35 ^ (fi34 << ((660) & 63))), F1: 1}
                                _ = v44
                            }
                        }
                        switch 32768 {
                        case -2147483648:
                            v17 = (-32768 == -223)
                        default:
                            fmt.Println(fi35, fi35)
                            v45 := GenS2{F0: -537, F1: 4294967295, F2: 4294967295}
                            v46 := GenS2{F0: 0, F1: 2, F2: 998}
                            v20 += float64(1.5)
                            var v47 int16 = 979
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        if true {
                            if (128 != 255) {
                                v48 := GenS2{F0: 42, F1: 123, F2: 42}
                                if (-1 != int16(fi35)) {
                                    var v49 float64 = float64(-2.25)
                                    fmt.Println(v20, v17)
                                    fmt.Println(fi34, fi35)
                                    _ = v49
                                }
                                _ = v48
                            } else {
                                var v50 int16 = -558
                                fmt.Println(fi35, v50, v50)
                                v17 = (2147483647 >= 1)
                                _ = v50
                            }
                        }
                        var v51 uint8 = 127
                        _ = fi35
                        _ = v51
                    }
                    _ = fi34
                }
            }
        }
    case 428:
        v20 = (float64(-2.25) * float64(0.1))
    }
    fmt.Println(v20, v17)
    v17 = false
    var v52 float64 = v20
    if v17 {
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
        _ = len(v3)
        if v17 {
            v17 = true
            if ((!v17) && v17) {
                v52 = v52
            } else {
                var v53 uint = 2
                _ = v53
            }
            if v17 {
                v52 -= (v52 + v52)
            }
        }
    } else {
        for fi54 := 0; fi54 < 15; fi54++ {
            for fi55 := 0; fi55 < 8; fi55++ {
                v56 := v2[0]
                _ = fi55
                _ = v56
            }
            _ = fi54
        }
        if ((32767 != -996) && false) {
            var v57 bool = v17
            var v58 float32 = 3.14159
            _ = v57
            _ = v58
        }
        v52 = v52
    }
    if len(v1) > 0 {
        v1[0] = 2
    }
    v59 := v3[886]
    v60 := GenS5{F0: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v52
    _ = v59
    _ = v60
}
