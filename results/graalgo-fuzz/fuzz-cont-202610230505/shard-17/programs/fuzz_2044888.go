package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint8; F1 int; F2 int16 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint8) int8 {
    p0 = p0
    v6 := []int{128, 0, 153, -2147483648}
    var v7 float32 = (1e10 - 3.14159)
    var v8 int16 = 32767
    var v9 uint = 1
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int) bool {
    var v10 uint64 = 2
    p0 += (p0 & -319)
    if ((1.5 == float64(3.14159)) || (true && (true || false))) {
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 int16 = -893
                switch fi11 {
                case 127:
                    var v14 int32 = -189
                    _ = v14
                case 898:
                    v10 -= 18446744073709551615
                case 9223372036854775807:
                    var v15 string = (("hello" + "") + "foo bar")
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            p0 -= 128
            if true {
                var v16 uint64 = v10
                _ = v16
            } else {
                var v17 uint8 = 79
                var v18 int16 = 42
                _ = v17
                _ = v18
            }
            _ = fi11
        }
    } else {
        fmt.Println(v10, p0)
    }
    _ = p0
    _ = v10
    return (!false)
}

func main() {
    var v1 int = 779
    v2 := make(map[string]uint32)
    var v3 int = (-(v1))
    var v4 int8 = -14
    var v5 uint8 = 2
    v1 += v1
    var v19 string = "foo bar"
    var v20 int = v1
    switch v19 {
    case "hello":
        switch v19 {
        case "foo bar":
            var v21 int16 = int16(v4)
            _ = v21
        case "hello":
            if true {
                fmt.Println(v5, v1, v20)
                var v22 float32 = 1e10
                for fi23 := 0; fi23 < 18; fi23++ {
                    if false {
                        v3 = 0
                    } else {
                        for fi24 := 0; fi24 < 9; fi24++ {
                            if false {
                                v22 += ((v22 * v22) * 0.1)
                            }
                            v25 := []int{(fi24 ^ (-(fi23))), v1, 1, 128}
                            if true {
                                fmt.Println(fi23)
                                v26 := GenS1{F0: 9223372036854775807, F1: 864}
                                v27 := GenS1{F0: 42, F1: 128}
                                _ = v26
                                _ = v27
                            }
                            _ = fi24
                            _ = v25
                        }
                    }
                    v1 += -9223372036854775808
                    v28 := GenS2{F0: -716, F1: "hello"}
                    _ = fi23
                    _ = v28
                }
                _ = v22
            } else {
                if true {
                    switch ("" + v19) {
                    case "a":
                        fmt.Println(v1)
                    default:
                        delete(v2, "hello")
                    }
                    v3 = v1
                } else {
                    v29 := GenS3{F0: 208, F1: -1}
                    _ = v29
                }
            }
        case "":
            switch v19 {
            case "hello":
                _ = len(v2)
            case "":
                fmt.Println(v19, v3)
            case "\n\t":
                var v30 int = v20
                _ = v30
            }
        default:
            var v31 float32 = 0.1
            _ = v31
        }
    case "foo bar":
        v3 -= (v3 * v3)
    case "":
        for fi32 := 0; fi32 < 15; fi32++ {
            switch (^(v20)) {
            case 4294967295:
                if ((v19 + v19) == "foo bar") {
                    if ((678 > 0) && true) {
                        v33 := v2[v19]
                        if false {
                            var v34 int8 = 102
                            _ = v34
                        } else {
                            if true {
                                v20 = (fi32 >> ((65535) & 63))
                                var v35 uint32 = (v33 / ((v33) | 1))
                                for fi36 := 0; fi36 < 3; fi36++ {
                                    v37 := GenS4{F0: 960, F1: -446}
                                    switch ("" + "") {
                                    case "foo bar":
                                        var v38 float32 = float32(3.14159)
                                        _ = v38
                                    case "hello":
                                        var v39 uint16 = 443
                                        _ = v39
                                    default:
                                        var v40 float64 = 0.1
                                        _ = v40
                                    }
                                    _ = fi36
                                    _ = v37
                                }
                                _ = v35
                            }
                        }
                        fmt.Println(v4, v3, v1)
                        _ = v33
                    }
                    v3 -= (v20 & (-(v3)))
                }
            case 128:
                for fi41 := 0; fi41 < 11; fi41++ {
                    var v42 int64 = 1
                    var v43 uint16 = 496
                    if (1 < 269) {
                        if true {
                            v4 = (v4 >> ((960) & 63))
                            var v44 uint16 = v43
                            _ = v44
                        }
                        if true {
                            v45 := GenS4{F0: 810, F1: int32(v43)}
                            v2[""] = 2147483648
                            var v46 uint = 14
                            _ = v45
                            _ = v46
                        }
                    }
                    v3 += fi41
                    var v47 float64 = ((((1.5 / 0.1) - 0.0) * float64(1e10)) - 1.5)
                    var v48 int16 = 1
                    v20 = (v3 - (fi41 ^ (^(fi41))))
                    _ = fi41
                    _ = v42
                    _ = v43
                    _ = v47
                    _ = v48
                }
            }
            v49 := []int{-907, fi32, v1, (v20 * (^(v20)))}
            v50 := GenS5{F0: 1, F1: v1, F2: -1}
            _ = fi32
            _ = v49
            _ = v50
        }
    default:
        if (-1 > 81) {
            v51 := GenS6{F0: v4}
            _ = v51
        } else {
            for fi52 := 0; fi52 < 3; fi52++ {
                v5 = v5
                if false {
                    var v53 int16 = -852
                    var v54 string = v19
                    for fi55 := 0; fi55 < 12; fi55++ {
                        var v56 float32 = float32((-2.25 / float64((float32(0.1) + (1e10 / ((-2.25 / 1e10) / 1.5))))))
                        _ = fi55
                        _ = v56
                    }
                    _ = v53
                    _ = v54
                }
                for fi57 := 0; fi57 < 9; fi57++ {
                    v1 -= (v1 * (^(fi57)))
                    _ = fi57
                }
                _ = fi52
            }
            if false {
                if true {
                    v19 = ((v19 + v19) + (v19 + v19))
                    v58 := GenS6{F0: (v4 * v4)}
                    _ = v58
                }
            } else {
                v1 += v3
                fmt.Println(v20)
            }
        }
    }
    v59 := GenS4{F0: 833, F1: -1}
    fmt.Println(v19, v1, v19)
    switch (v20 << ((32) & 63)) {
    case 694:
        v60 := GenS2{F0: 2147483647, F1: v19}
        _ = v60
    case 246:
        for fi61 := 0; fi61 < 1; fi61++ {
            if true {
                for fi62 := 0; fi62 < 10; fi62++ {
                    var v63 uint32 = 1
                    v64 := GenS2{F0: 42, F1: "\n\t"}
                    _ = fi62
                    _ = v63
                    _ = v64
                }
            } else {
                var v65 string = v19
                if (9223372036854775807 > 167) {
                    for fi66 := 0; fi66 < 5; fi66++ {
                        v19 = v19
                        _ = fi66
                    }
                }
                _ = v65
            }
            v67 := GenS1{F0: int64(v20), F1: 65535}
            _ = fi61
            _ = v67
        }
    }
    var v68 int16 = -32768
    switch 9223372036854775807 {
    case 2147483647:
        v20 = (^(v20))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v68
}
