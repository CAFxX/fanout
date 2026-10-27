package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 bool; F2 string }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint; F1 int }
func genFn0(p0 string, p1 string, p2 uint64) uint {
    var v7 float32 = 1e10
    p0 = p0
    if true {
        v8 := []int{486}
        for fi9 := 0; fi9 < 19; fi9++ {
            _ = len(v8)
            if true {
                var v10 int64 = 1
                var v11 int64 = 106
                _ = v10
                _ = v11
            } else {
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 bool = (true || false)
                    var v14 string = p0
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                fmt.Println(v7, p0, p0)
                if len(v8) > 0 {
                    v8[0] = fi9
                }
            }
            p1 = (p0 + p1)
            _ = fi9
        }
        _ = v8
    } else {
        if (true || true) {
            p1 = p1
            fmt.Println(v7, p1, v7)
        } else {
            p0 = p0
            v15 := GenS1{F0: 18446744073709551615, F1: 2147483647}
            _ = v15
        }
        if (9223372036854775807 >= 42) {
            var v16 string = p1
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 uint16, p1 string, p2 int32) int64 {
    if (0.0 < 1e10) {
        fmt.Println(p2)
    } else {
        fmt.Println(p0)
        v17 := []int{964, -228}
        p1 = p1
        _ = v17
    }
    switch "foo bar" {
    case "hello":
        var v18 int32 = int32(p0)
        _ = v18
    case "foo bar":
        v19 := []int{9223372036854775807, 127, -743, 9223372036854775807}
        _ = v19
    }
    v20 := []int{413}
    var v21 uint16 = (p0 ^ 42)
    fmt.Println(p2, v21, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 9223372036854775807
}

func main() {
    v1 := []string{"hello"}
    v2 := []int{-842}
    var v3 uint = 0
    var v4 float32 = (0.1 + 1.5)
    var v5 uint64 = 4294967295
    var v6 float32 = v4
    var v22 int16 = 1
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 int16 = (v22 % ((v22) | 1))
        v6 -= ((0.0 * v4) + float32((float64(v4) - 3.14159)))
        var v25 uint32 = 2
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 float64 = 1e10
    v4 += ((v6 - 0.1) * (1.5 * float32((v26 - ((float64(float32(v26)) * (3.14159 + (v26 + v26))) / v26)))))
    v22 += (v22 - (v22 - v22))
    v27 := GenS2{F0: 2}
    v28 := GenS2{F0: 42}
    for fi29 := 0; fi29 < 3; fi29++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            var v30 bool = true
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            v31 := GenS3{F0: ""}
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 string = ("" + "")
            v34 := GenS4{F0: fi32, F1: 255}
            v4 += 3.14159
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi29
    }
    if true {
        for fi35 := 0; fi35 < 7; fi35++ {
            if false {
                fmt.Println(v6)
            } else {
                if false {
                    if (!false) {
                        var v36 uint64 = v5
                        var v37 int64 = -9223372036854775808
                        _ = len(v2)
                        _ = v36
                        _ = v37
                    } else {
                        var v38 int32 = 1
                        _ = v38
                    }
                    var v39 int64 = 91
                    _ = v39
                }
                if false {
                    var v40 int32 = -314
                    fmt.Println(v6, v6)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v40
                }
            }
            v41 := GenS3{F0: ""}
            if (!(v4 > 0.0)) {
                v26 = float64(v4)
                if false {
                    fmt.Println(v26)
                } else {
                    var v42 int64 = -654
                    var v43 uint64 = v5
                    _ = v42
                    _ = v43
                }
                if true {
                    v4 -= v4
                    v26 += (v26 + float64(v4))
                } else {
                    for fi44 := 0; fi44 < 4; fi44++ {
                        v45 := GenS5{F0: ("foo bar" + ("a" + "foo bar"))}
                        _ = fi44
                        _ = v45
                    }
                    v46 := GenS6{F0: 185}
                    _ = v46
                }
            }
            _ = fi35
            _ = v41
        }
        var v47 int32 = 615
        v1 = append(v1, "a")
        _ = v47
    } else {
        for fi48 := 0; fi48 < 6; fi48++ {
            switch ("hello" + ("hello" + "foo bar")) {
            case "foo bar":
                fmt.Println(v6, v26)
            }
            _ = fi48
        }
        switch ("a" + "a") {
        case "foo bar":
            switch "\n\t" {
            case "hello":
                if true {
                    v3 = v3
                    v49 := GenS3{F0: ("a" + ((("foo bar" + ("\n\t" + ("foo bar" + "a"))) + "hello") + "\n\t"))}
                    var v50 bool = (true || true)
                    _ = v49
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 6; fi51++ {
                        v52 := GenS5{F0: "hello"}
                        _ = fi51
                        _ = v52
                    }
                    if false {
                        v26 += (float64(v4) + v26)
                    } else {
                        if false {
                            var v53 float64 = v26
                            v54 := GenS7{F0: v22, F1: (145 >= 255), F2: "hello"}
                            v55 := GenS3{F0: "\n\t"}
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                        if false {
                            var v56 string = "hello"
                            for fi57 := 0; fi57 < 10; fi57++ {
                                var v58 uint16 = 1
                                v59 := GenS2{F0: (v58 | 42)}
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                            _ = v56
                        } else {
                            fmt.Println(v26, v6)
                        }
                        v60 := GenS1{F0: 2, F1: 684}
                        _ = v60
                    }
                }
            case "":
                v5 -= (^(v5))
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            }
        case "":
            v61 := GenS8{F0: (-351 < 194), F1: 42}
            _ = v61
        default:
            var v62 int16 = v22
            _ = v62
        }
        v63 := GenS8{F0: true, F1: 4294967295}
        _ = v63
    }
    for fi64 := 0; fi64 < 17; fi64++ {
        switch (fi64 - 4294967295) {
        case -2147483648:
            if false {
                switch ("a" + "foo bar") {
                case "foo bar":
                    if true {
                        var v65 int8 = 0
                        _ = v65
                    }
                case "":
                    var v66 int8 = -1
                    _ = v66
                case "a":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    fmt.Println(v26, v4, v22)
                }
            } else {
                for fi67 := 0; fi67 < 3; fi67++ {
                    if true {
                        if len(v1) > 0 {
                            v1[0] = (("foo bar" + "") + "a")
                        }
                        var v68 uint8 = 31
                        v69 := GenS9{F0: v3, F1: int(v22)}
                        _ = v68
                        _ = v69
                    }
                    _ = fi67
                }
                var v70 uint = (v3 | 0)
                _ = v70
            }
        case -690:
            var v71 int8 = -128
            _ = v71
        case 9223372036854775807:
            for fi72 := 0; fi72 < 6; fi72++ {
                v5 -= v5
                var v73 uint32 = 0
                _ = fi72
                _ = v73
            }
        default:
            v5 = (v5 + 42)
        }
        fmt.Println(v22, v6)
        _ = fi64
    }
    if false {
        for fi74 := 0; fi74 < 14; fi74++ {
            switch ("hello" + (("hello" + ("a" + "foo bar")) + ("" + "a"))) {
            case "\n\t":
                _ = len(v1)
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = fi74
                }
            default:
                if (!false) {
                    v75 := GenS7{F0: v22, F1: false, F2: "hello"}
                    _ = v75
                } else {
                    v2 = append(v2, fi74)
                }
            }
            fmt.Println(v4, v26, v22)
            _ = fi74
        }
        v22 = (v22 << ((28) & 63))
    }
    v26 = v26
    v76 := GenS2{F0: 65535}
    switch "\n\t" {
    case "":
        var v77 uint32 = 706
        _ = v77
    case "foo bar":
        var v78 uint64 = v5
        _ = v78
    default:
        fmt.Println(v22)
    }
    var v79 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v79
}
