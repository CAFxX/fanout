package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS4 struct { F0 int8; F1 uint8; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint8 = 127
    var v2 uint32 = 2
    var v3 uint16 = 718
    v4 := make(map[string]int)
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(v2)
            _ = fi6
        }
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int8 = -1
            _ = fi7
            _ = v8
        }
        _ = fi5
    }
    fmt.Println(v3)
    v2 = 2147483648
    if ("\n\t" != "a") {
        v9 := GenS1{F0: -1, F1: 1, F2: 1}
        if (-9223372036854775808 >= -842) {
            var v10 float64 = 3.14159
            v11 := GenS1{F0: -9223372036854775808, F1: 1, F2: -1}
            if ((v10 + float64((float32((float64(0.1) - v10)) + float32(float64((1e10 * 1.5)))))) != float64(float32(v10))) {
                for fi12 := 0; fi12 < 10; fi12++ {
                    if true {
                        var v13 int64 = 0
                        _ = v13
                    }
                    var v14 uint16 = 42
                    _ = fi12
                    _ = v14
                }
            } else {
                v15 := []int{2147483647, -2147483648, 0, 42}
                _ = v15
            }
            _ = v10
            _ = v11
        } else {
            var v16 bool = (false || (!(-1 == 0)))
            _ = v16
        }
        for fi17 := 0; fi17 < 8; fi17++ {
            if (128 < 128) {
                for fi18 := 0; fi18 < 7; fi18++ {
                    v19 := GenS1{F0: 1, F1: 1, F2: 768}
                    var v20 uint64 = uint64(v3)
                    var v21 int32 = 1
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                var v22 string = "foo bar"
                _ = v22
            } else {
                for fi23 := 0; fi23 < 4; fi23++ {
                    switch "" {
                    case "":
                        var v24 float32 = 1e10
                        _ = v24
                    }
                    _ = fi23
                }
            }
            var v25 uint32 = 65535
            v26 := []int{fi17, (fi17 - -211)}
            _ = fi17
            _ = v25
            _ = v26
        }
        _ = v9
    } else {
        for fi27 := 0; fi27 < 8; fi27++ {
            fmt.Println(v3, v2, fi27)
            v3 = uint16(v1)
            v2 -= v2
            _ = fi27
        }
        fmt.Println(v1)
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        v29 := GenS1{F0: 42, F1: 1, F2: 0}
        if (true || true) {
            fmt.Println(v2, v3, fi28)
        }
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := GenS1{F0: 9223372036854775807, F1: 0, F2: 42}
        v1 += (^(v1))
        _ = fi30
        _ = v31
    }
    v3 = v3
    for fi32 := 0; fi32 < 8; fi32++ {
        for fi33 := 0; fi33 < 18; fi33++ {
            fmt.Println(v3, v2)
            if false {
                if false {
                    var v34 uint32 = (v2 + 0)
                    delete(v4, "hello")
                    if (false && ((("hello" + "foo bar") + "") <= "foo bar")) {
                        v35 := []int{int(v1), fi32}
                        var v36 float32 = float32(0.1)
                        v37 := GenS2{F0: v1}
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    _ = v34
                }
                fmt.Println(fi32)
                if false {
                    switch fi32 {
                    case -570:
                        delete(v4, "\n\t")
                    case -694:
                        var v38 int = fi32
                        _ = v38
                    case 9223372036854775807:
                        v3 = v3
                    }
                    fmt.Println(v2)
                }
            } else {
                v39 := []int{}
                switch (("" + "\n\t") + "") {
                case "hello":
                    if true {
                        v40 := GenS2{F0: v1}
                        v1 = (^(v1))
                        _ = v40
                    } else {
                        var v41 uint64 = 18446744073709551615
                        if (!true) {
                            v1 = 39
                        }
                        _ = v41
                    }
                default:
                    v42 := GenS3{F0: 2147483648, F1: v2, F2: -48}
                    _ = v42
                }
                _ = v39
            }
            var v43 int16 = 246
            _ = fi33
            _ = v43
        }
        switch fi32 {
        case 32768:
            v44 := []int{(fi32 << ((4294967295) & 63)), fi32}
            _ = v44
        }
        if false {
            if true {
                switch (fi32 >> ((265) & 63)) {
                case -9223372036854775808:
                    switch fi32 {
                    case 32767:
                        var v45 bool = (47 <= -128)
                        _ = v45
                    }
                case -2147483648:
                    if false {
                        if (!false) {
                            var v46 uint8 = (v1 % ((v1) | 1))
                            _ = v46
                        }
                        var v47 int8 = 63
                        _ = v47
                    }
                case 611:
                    v3 = (v3 * (v3 % (((v3 >> ((52) & 63))) | 1)))
                default:
                    v48 := GenS3{F0: v2, F1: uint32(v3), F2: 0}
                    _ = v48
                }
                if (194 == v1) {
                    var v49 uint64 = 0
                    var v50 uint = 2
                    fmt.Println(fi32)
                    _ = v49
                    _ = v50
                }
            } else {
                _ = len(v4)
            }
            switch "" {
            case "foo bar":
                var v51 float64 = 1e10
                _ = v51
            default:
                v52 := []int{}
                _ = v52
            }
        } else {
            for fi53 := 0; fi53 < 9; fi53++ {
                switch ("foo bar" + ("a" + "\n\t")) {
                case "\n\t":
                    fmt.Println(fi32, v2)
                }
                var v54 uint32 = 698
                _ = fi53
                _ = v54
            }
            fmt.Println(v3)
            v2 -= uint32(v3)
        }
        _ = fi32
    }
    var v55 bool = (-982 > 127)
    v56 := GenS4{F0: -68, F1: (v1 - 255), F2: 0}
    for fi57 := 0; fi57 < 5; fi57++ {
        for fi58 := 0; fi58 < 2; fi58++ {
            v59 := v4["a"]
            v60 := GenS5{F0: (fi57 | fi58)}
            switch fi57 {
            case 422:
                v61 := GenS3{F0: (v2 & uint32(v1)), F1: 468, F2: -26}
                _ = v61
            }
            _ = fi58
            _ = v59
            _ = v60
        }
        _ = fi57
    }
    var v62 int64 = 935
    for fi63 := 0; fi63 < 14; fi63++ {
        if v55 {
            fmt.Println(fi63, v55)
            fmt.Println(v2)
        }
        var v64 uint64 = 794
        v62 -= 9223372036854775807
        _ = fi63
        _ = v64
    }
    switch (("foo bar" + "hello") + (("hello" + ("foo bar" + "a")) + ("a" + "foo bar"))) {
    case "":
        var v65 uint = 101
        _ = v65
    case "a":
        _ = len(v4)
    case "\n\t":
        for fi66 := 0; fi66 < 9; fi66++ {
            for fi67 := 0; fi67 < 6; fi67++ {
                if v55 {
                    var v68 uint8 = (v1 << ((916) & 63))
                    _ = v68
                } else {
                    v1 = (v1 << ((717) & 63))
                    switch "hello" {
                    case "":
                        var v69 int = (fi67 / ((fi66) | 1))
                        _ = v69
                    case "hello":
                        v55 = false
                    case "foo bar":
                        v3 += v3
                    }
                    v3 = v3
                }
                var v70 uint64 = 0
                v4[((("" + "a") + ("foo bar" + ("foo bar" + "\n\t"))) + ("" + ""))] = (fi67 * fi66)
                _ = fi67
                _ = v70
            }
            v71 := GenS3{F0: v2, F1: (^(v2)), F2: 1}
            for fi72 := 0; fi72 < 13; fi72++ {
                var v73 float64 = 0.0
                var v74 int32 = 344
                if ("a" > ("hello" + "hello")) {
                    v75 := []int{fi66, (fi66 ^ (-(fi72)))}
                    v76 := GenS3{F0: v2, F1: 1, F2: 127}
                    v77 := GenS6{F0: 65535}
                    _ = v75
                    _ = v76
                    _ = v77
                } else {
                    v78 := GenS5{F0: (-(fi72))}
                    switch -156 {
                    case 1:
                        v1 = uint8(v3)
                    case 321:
                        v2 = (v2 ^ v2)
                        fmt.Println(v55)
                    case 128:
                        switch "hello" {
                        case "\n\t":
                            var v79 int32 = v74
                            _ = v79
                        default:
                            var v80 int16 = -1
                            _ = v80
                        }
                    default:
                        fmt.Println(v3, v2)
                    }
                    v81 := []int{fi72, (fi66 - 2147483647), (fi66 >> ((237) & 63)), (fi66 * -340)}
                    v1 = 0
                    _ = v78
                    _ = v81
                }
                _ = fi72
                _ = v73
                _ = v74
            }
            _ = fi66
            _ = v71
        }
    }
    v2 = (v2 - (v2 & v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v62
}
