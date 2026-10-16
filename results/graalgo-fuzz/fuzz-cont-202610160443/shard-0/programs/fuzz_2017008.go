package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 int32; F1 int32; F2 int32 }
type GenS3 struct { F0 string; F1 uint; F2 string }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int32; F1 int32; F2 int8 }
type GenS6 struct { F0 int; F1 int16 }
func genFn0(p0 uint8) int64 {
    v7 := GenS1{F0: -1, F1: ("\n\t" + ("hello" + ("\n\t" + ((("foo bar" + "hello") + "foo bar") + "")))), F2: 32767}
    var v8 uint8 = (p0 | p0)
    if (!true) {
        if (0 >= -97) {
            v8 = (^(p0))
            v9 := GenS1{F0: 32767, F1: "foo bar", F2: 42}
            _ = v9
        }
        v10 := GenS2{F0: 1, F1: -1, F2: 2147483647}
        fmt.Println(v8)
        _ = v10
    } else {
        p0 = 128
        if (255 != 1) {
            p0 -= 118
            var v11 int64 = 0
            for fi12 := 0; fi12 < 15; fi12++ {
                v13 := []int{(fi12 ^ 682), fi12, 42}
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                var v14 bool = (0 < 200)
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
    }
    var v15 uint = 652
    fmt.Println(v15, v15)
    _ = p0
    _ = v15
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 string = "hello"
    var v3 uint16 = 255
    var v4 uint32 = 65535
    var v5 uint = 841
    var v6 int64 = (v1 % ((-775) | 1))
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS1{F0: 1000, F1: v2, F2: 827}
        var v18 int32 = -2147483648
        var v19 string = v2
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 int8 = 0
    v21 := []int{-1, 757, 4294967295, -801}
    v2 = (v2 + "a")
    v22 := GenS1{F0: 0, F1: "hello", F2: -32768}
    for fi23 := 0; fi23 < 2; fi23++ {
        v5 -= v5
        v20 += int8(v4)
        _ = fi23
    }
    if (v5 != v5) {
        switch int(v1) {
        case -575:
            for fi24 := 0; fi24 < 16; fi24++ {
                v3 = (v3 ^ v3)
                if true {
                    var v25 uint32 = (^(v4))
                    _ = v25
                }
                _ = fi24
            }
        case 42:
            for fi26 := 0; fi26 < 4; fi26++ {
                for fi27 := 0; fi27 < 5; fi27++ {
                    switch -525 {
                    case 127:
                        v3 = v3
                        switch fi27 {
                        case 0:
                            if true {
                                var v28 string = "hello"
                                _ = v28
                            } else {
                                var v29 bool = true
                                if false {
                                    var v30 int = (fi27 | (fi26 + (-(fi26))))
                                    _ = v30
                                } else {
                                    var v31 int64 = (v1 << ((v5) & 63))
                                    _ = v31
                                }
                                _ = v29
                            }
                        case 9223372036854775807:
                            v3 = v3
                        case 232:
                            if false {
                                fmt.Println(v2, v4)
                                var v32 string = v2
                                fmt.Println(v3, v5, v2)
                                _ = v32
                            } else {
                                switch v2 {
                                case "foo bar":
                                    var v33 uint64 = 18446744073709551615
                                    switch v2 {
                                    case "\n\t":
                                        if len(v21) > 0 {
                                            v21[0] = (fi26 >> ((2) & 63))
                                        }
                                    }
                                    v34 := GenS2{F0: 0, F1: -303, F2: int32(v3)}
                                    _ = v33
                                    _ = v34
                                case "a":
                                    v5 = v5
                                case "\n\t":
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                default:
                                    var v35 float64 = (float64(float32(float64((((1e10 + 0.0) / 0.1) - 1.5)))) + (-2.25 + 1e10))
                                    _ = v35
                                }
                                switch (fi26 ^ 32767) {
                                case 42:
                                    v36 := GenS2{F0: 2147483647, F1: 42, F2: -2147483648}
                                    _ = v36
                                default:
                                    v37 := GenS3{F0: "foo bar", F1: (v5 ^ (v5 / ((v5) | 1))), F2: (v2 + "\n\t")}
                                    _ = v37
                                }
                                v4 = (^(v4))
                            }
                            v6 -= v1
                            if (!((^(v3)) == 0)) {
                                if ((float32(((((0.1 * 1.5) - ((0.1 + (0.1 + 1e10)) - ((1.5 - 0.0) * float64((3.14159 - (1e10 / (0.0 + ((3.14159 * ((0.1 - 1e10) + 1e10)) - 1e10)))))))) * 1.5) + float64(1e10))) >= 1.5) || false) {
                                    v38 := GenS3{F0: ("foo bar" + ""), F1: (v5 % (((v5 + v5)) | 1)), F2: v2}
                                    var v39 bool = (false && false)
                                    v40 := GenS4{F0: (^(v5)), F1: int64(v20)}
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                                v41 := GenS1{F0: -1, F1: ((("foo bar" + v2) + (v2 + v2)) + "foo bar"), F2: -1}
                                _ = v41
                            }
                            _ = len(v21)
                            v5 = (v5 / (((v5 * 65535)) | 1))
                        }
                        if len(v21) > 0 {
                            v21[0] = -9223372036854775808
                        }
                        if true {
                            var v42 uint = 18446744073709551615
                            _ = v42
                        }
                    case -1:
                        if true {
                            if false {
                                switch (fi26 ^ (fi27 % ((fi27) | 1))) {
                                case -2147483648:
                                    fmt.Println(v5)
                                case 128:
                                    var v43 string = "\n\t"
                                    _ = v43
                                default:
                                    v44 := GenS3{F0: "a", F1: v5, F2: v2}
                                    _ = v44
                                }
                                v45 := GenS1{F0: 42, F1: ("foo bar" + v2), F2: 0}
                                switch fi26 {
                                case 32768:
                                    v46 := GenS2{F0: -1, F1: 42, F2: 195}
                                    _ = v46
                                case 642:
                                    v47 := GenS4{F0: uint(v4), F1: (-(v1))}
                                    _ = v47
                                case -9223372036854775808:
                                    switch fi27 {
                                    case 32768:
                                        v6 += genFn0(42)
                                    case -1:
                                        if (((3.14159 - 1e10) - float32(1e10)) == (0.1 * 1e10)) {
                                            _ = len(v21)
                                            v48 := GenS5{F0: 2147483647, F1: -1, F2: v20}
                                            v3 += (^(v3))
                                            _ = v48
                                        }
                                        var v49 uint64 = 835
                                        _ = v49
                                    }
                                default:
                                    switch "" {
                                    case "a":
                                        _ = len(v21)
                                    case "foo bar":
                                        var v50 string = "a"
                                        _ = v50
                                    }
                                }
                                _ = v45
                            } else {
                                var v51 uint32 = (^(v4))
                                v52 := GenS2{F0: 1, F1: 42, F2: -2147483648}
                                v51 += (v4 - v51)
                                if (!(false || false)) {
                                    if (v3 == (v3 / ((v3) | 1))) {
                                        switch "foo bar" {
                                        case "hello":
                                            var v53 uint16 = v3
                                            _ = v53
                                        case "foo bar":
                                            var v54 int32 = 607
                                            _ = v54
                                        case "":
                                            if (75 == 255) {
                                                fmt.Println(v51, v20, fi26)
                                                v1 -= -724
                                                if (1.5 != (0.1 * 0.0)) {
                                                    var v55 uint32 = (v51 - v51)
                                                    _ = v55
                                                } else {
                                                    if (true && true) {
                                                        var v56 int8 = (v20 ^ v20)
                                                        _ = v56
                                                    }
                                                    if (!(int16(v20) >= 898)) {
                                                        v6 += genFn0(182)
                                                        v57 := GenS2{F0: 868, F1: 1, F2: 1}
                                                        _ = v57
                                                    } else {
                                                        v21 = append(v21, (-(fi26)))
                                                        v21 = append(v21, fi26)
                                                        if (-85 <= 655) {
                                                            v2 = v2
                                                            v1 -= v1
                                                            fmt.Println(v5, v4, v5)
                                                            var v58 int16 = -1
                                                            var v59 uint = v5
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                    }
                                                    if len(v21) > 0 {
                                                        v21[0] = (fi27 / ((fi26) | 1))
                                                    }
                                                }
                                                v21 = append(v21, (fi27 * (fi26 >> ((v5) & 63))))
                                                fmt.Println(v20)
                                            }
                                        default:
                                            var v60 uint = v5
                                            _ = v60
                                        }
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                        switch fi27 {
                                        case 128:
                                            v61 := GenS3{F0: v2, F1: 4294967295, F2: v2}
                                            _ = v61
                                        case -9223372036854775808:
                                            var v62 uint16 = v3
                                            _ = v62
                                        default:
                                            var v63 int32 = 2147483647
                                            _ = v63
                                        }
                                    }
                                    var v64 int32 = int32(v6)
                                    _ = v64
                                } else {
                                    var v65 uint8 = uint8(v51)
                                    v66 := GenS6{F0: (^(fi27)), F1: 0}
                                    _ = len(v21)
                                    _ = v65
                                    _ = v66
                                }
                                if len(v21) > 0 {
                                    v21[0] = 0
                                }
                                switch v2 {
                                case "hello":
                                    fmt.Println(v51)
                                default:
                                    v51 = v51
                                }
                                var v67 int = fi27
                                v51 = v4
                                _ = v51
                                _ = v52
                                _ = v67
                            }
                            var v68 string = ("a" + (((v2 + v2) + v2) + "a"))
                            v4 = v4
                            _ = v68
                        }
                    }
                    _ = fi27
                }
                _ = fi26
            }
        }
        var v69 int32 = 701
        _ = v69
    }
    switch -175 {
    case 1:
        var v70 bool = true
        _ = v70
    }
    var v71 float32 = (1.5 + (0.1 - (-2.25 - (-2.25 * (1.5 + 1e10)))))
    fmt.Println(v4)
    var v72 string = v2
    switch "\n\t" {
    case "\n\t":
        switch ("\n\t" + "a") {
        case "hello":
            v71 += v71
        case "\n\t":
            var v73 uint8 = 255
            _ = v73
        case "a":
            var v74 int16 = 1
            _ = v74
        default:
            v72 = v2
        }
    default:
        var v75 string = v72
        _ = v75
    }
    switch 633 {
    case 0:
        v2 = "hello"
    default:
        fmt.Println(v71, v3, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
    _ = v72
}
