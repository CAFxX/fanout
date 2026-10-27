package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint32; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 int16; F1 uint32; F2 uint64 }
func genFn0() string {
    var v5 float32 = float32((1e10 + 0.0))
    fmt.Println(v5)
    switch -1 {
    case -203:
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 float64 = (0.0 * 3.14159)
            _ = fi6
            _ = v7
        }
    case -1:
        var v8 float64 = 3.14159
        _ = v8
    default:
        v5 = float32(1.5)
    }
    v5 += float32(0.0)
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS1{F0: -1, F1: -493}
        fmt.Println(v5, v5)
        _ = fi9
        _ = v10
    }
    _ = v5
    return ""
}
func genFn1(p0 int, p1 bool, p2 uint) uint8 {
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint = p2
        var v13 uint32 = 2147483647
        if p1 {
            v14 := []int{-2147483648}
            _ = v14
        } else {
            var v15 uint = (v12 | 2)
            _ = v15
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    switch ("\n\t" + ("foo bar" + "\n\t")) {
    case "a":
        var v16 bool = (!(-1 != 9223372036854775807))
        _ = v16
    case "hello":
        if (true && ((!p1) && (-208 <= 8))) {
            switch (^(p0)) {
            case 42:
                v17 := GenS2{F0: 2}
                _ = v17
            case 128:
                p1 = p1
            }
            if (!(!p1)) {
                v18 := []int{-1}
                switch (("a" + "\n\t") + "") {
                case "\n\t":
                    for fi19 := 0; fi19 < 2; fi19++ {
                        var v20 float64 = (0.0 - 0.0)
                        switch "a" {
                        case "foo bar":
                            switch "a" {
                            case "":
                                for fi21 := 0; fi21 < 13; fi21++ {
                                    p0 = fi21
                                    switch p0 {
                                    case 128:
                                        var v22 int = fi21
                                        _ = v22
                                    }
                                    _ = fi21
                                }
                            case "\n\t":
                                switch "hello" {
                                case "":
                                    v23 := GenS2{F0: 65535}
                                    _ = v23
                                default:
                                    p0 = fi19
                                }
                            case "a":
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            default:
                                var v24 uint = (p2 | 42)
                                _ = v24
                            }
                        }
                        for fi25 := 0; fi25 < 19; fi25++ {
                            p1 = p1
                            var v26 uint32 = 42
                            switch fi25 {
                            case -2147483648:
                                fmt.Println(p0)
                                if (("" + "\n\t") != ("\n\t" + "foo bar")) {
                                    if len(v18) > 0 {
                                        v18[0] = (fi25 * p0)
                                    }
                                    var v27 int32 = -1
                                    switch (fi19 | fi25) {
                                    case 1:
                                        if p1 {
                                            v28 := GenS1{F0: 1000, F1: v27}
                                            _ = v28
                                        } else {
                                            p2 -= p2
                                        }
                                    case 127:
                                        v29 := GenS3{F0: 2147483647, F1: 216, F2: (v27 & (^(v27)))}
                                        _ = v29
                                    default:
                                        fmt.Println(fi25)
                                    }
                                    _ = v27
                                } else {
                                    if (int32(v26) < 42) {
                                        fmt.Println(v20, v26, v20)
                                        v30 := GenS3{F0: int32(fi25), F1: (^(v26)), F2: 0}
                                        _ = v30
                                    }
                                }
                            default:
                                p0 += fi25
                            }
                            _ = fi25
                            _ = v26
                        }
                        _ = fi19
                        _ = v20
                    }
                }
                _ = v18
            }
        }
    case "foo bar":
        fmt.Println(p0)
    }
    if (!(p1 && p1)) {
        fmt.Println(p0)
    }
    p2 += (^(p2))
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 75
    v3 := []uint32{656, 77, 4294967295}
    var v4 bool = true
    if (4294967295 >= 4294967295) {
        if (!v4) {
            switch "a" {
            case "\n\t":
                if false {
                    var v31 int32 = -1
                    if (42 != 4) {
                        v3 = append(v3, 420)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    fmt.Println(v2)
                    _ = v31
                }
            default:
                switch "" {
                case "foo bar":
                    var v32 int8 = v2
                    _ = v32
                case "hello":
                    var v33 int8 = -128
                    _ = v33
                default:
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v4, v4)
                    case "hello":
                        if false {
                            fmt.Println(v4)
                            var v34 int32 = -2147483648
                            v2 -= (v2 << ((42) & 63))
                            _ = v34
                        } else {
                            v2 += (v2 >> ((255) & 63))
                            for fi35 := 0; fi35 < 19; fi35++ {
                                v4 = (127 < (v2 - (v2 ^ (v2 / (((v2 - v2)) | 1)))))
                                v2 += (v2 | 100)
                                v1[""] = -873
                                _ = fi35
                            }
                            var v36 uint32 = 0
                            _ = v36
                        }
                    default:
                        v2 -= v2
                    }
                }
            }
            if v4 {
                for fi37 := 0; fi37 < 17; fi37++ {
                    var v38 int = 0
                    _ = fi37
                    _ = v38
                }
            }
        }
    } else {
        if v4 {
            if true {
                v4 = v4
                v2 -= v2
                fmt.Println(v2, v4)
            } else {
                switch (("foo bar" + "foo bar") + "") {
                case "hello":
                    switch 2147483647 {
                    case 233:
                        var v39 uint = 4294967295
                        _ = v39
                    case 468:
                        switch "a" {
                        case "":
                            var v40 uint64 = 42
                            _ = v40
                        case "\n\t":
                            v3 = append(v3, 0)
                        case "a":
                            switch 960 {
                            case -2147483648:
                                var v41 int32 = 1
                                _ = v41
                            case 668:
                                if false {
                                    v2 -= -57
                                    var v42 uint = 65535
                                    _ = v42
                                }
                            default:
                                var v43 int8 = (v2 ^ (v2 * (v2 ^ (v2 ^ (v2 + v2)))))
                                _ = v43
                            }
                        default:
                            v2 -= (v2 / ((v2) | 1))
                        }
                    default:
                        switch (("" + "\n\t") + "") {
                        case "foo bar":
                            for fi44 := 0; fi44 < 18; fi44++ {
                                v2 = v2
                                var v45 uint16 = 98
                                _ = fi44
                                _ = v45
                            }
                        case "":
                            for fi46 := 0; fi46 < 18; fi46++ {
                                var v47 int16 = -685
                                var v48 int8 = 34
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                        case "hello":
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v4)
                            case "":
                                if true {
                                    if true {
                                        if v4 {
                                            v49 := GenS3{F0: int32(v2), F1: 345, F2: 0}
                                            _ = v49
                                        } else {
                                            var v50 uint8 = 128
                                            _ = v50
                                        }
                                        v4 = v4
                                    } else {
                                        fmt.Println(v4)
                                        if v4 {
                                            v4 = false
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            switch ("foo bar" + "\n\t") {
                                            case "foo bar":
                                                var v51 uint16 = 65535
                                                _ = v51
                                            case "\n\t":
                                                var v52 uint = 2
                                                _ = v52
                                            default:
                                                if v4 {
                                                    var v53 uint8 = 0
                                                    _ = v53
                                                } else {
                                                    for fi54 := 0; fi54 < 15; fi54++ {
                                                        v4 = v4
                                                        _ = fi54
                                                    }
                                                    v55 := GenS2{F0: 945}
                                                    v56 := GenS4{F0: 0}
                                                    _ = v55
                                                    _ = v56
                                                }
                                            }
                                        }
                                    }
                                    var v57 int32 = 0
                                    var v58 int64 = -1
                                    _ = v57
                                    _ = v58
                                }
                            default:
                                _ = len(v1)
                            }
                        }
                    }
                }
                v4 = v4
                if v4 {
                    v2 += v2
                    var v59 int16 = 1
                    _ = v59
                }
            }
        } else {
            var v60 float64 = (0.0 * (0.1 / -2.25))
            var v61 int = 544
            _ = v60
            _ = v61
        }
    }
    v3 = append(v3, 42)
    for fi62 := 0; fi62 < 8; fi62++ {
        var v63 float64 = (0.1 / 1e10)
        _ = fi62
        _ = v63
    }
    v4 = (0 == 182)
    v64 := GenS2{F0: 42}
    v65 := GenS3{F0: -383, F1: 2147483647, F2: -160}
    v3 = append(v3, 2)
    fmt.Println(v2, v4)
    var v66 uint64 = 1
    v67 := GenS2{F0: 159}
    fmt.Println(v4)
    for fi68 := 0; fi68 < 19; fi68++ {
        v4 = ((!v4) || (!("foo bar" > ("a" + ""))))
        _ = fi68
    }
    for fi69 := 0; fi69 < 20; fi69++ {
        if v4 {
            for fi70 := 0; fi70 < 8; fi70++ {
                v71 := GenS1{F0: 0, F1: 0}
                _ = fi70
                _ = v71
            }
        } else {
            fmt.Println(v4, v4, v4)
            v3 = append(v3, 2147483648)
        }
        switch fi69 {
        case 312:
            v2 = v2
        case 688:
            v72 := GenS5{F0: 255, F1: fi69}
            _ = v72
        case 128:
            switch "hello" {
            case "":
                switch ("hello" + ("a" + "hello")) {
                case "foo bar":
                    v2 = v2
                case "hello":
                    _ = len(v3)
                }
            case "hello":
                for fi73 := 0; fi73 < 4; fi73++ {
                    var v74 int32 = 0
                    v2 = v2
                    v4 = (v66 >= (v66 ^ v66))
                    _ = fi73
                    _ = v74
                }
            case "foo bar":
                var v75 uint64 = (v66 - v66)
                _ = v75
            default:
                v4 = (370 == 1)
            }
        default:
            if (!((("a" + "foo bar") + "a") != ("a" + ("" + "hello")))) {
                v76 := GenS2{F0: 1}
                var v77 uint32 = 0
                delete(v1, (("foo bar" + "hello") + ""))
                _ = v76
                _ = v77
            } else {
                var v78 float64 = (1.5 * 0.0)
                v79 := GenS6{F0: 1000, F1: 2147483648, F2: (^(v66))}
                fmt.Println(v78, v4, v78)
                _ = v78
                _ = v79
            }
        }
        _ = fi69
    }
    var v80 int8 = (v2 * v2)
    v81 := GenS4{F0: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v80
    _ = v81
}
