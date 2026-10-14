package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int = 127
    v2 := []int{v1, -1}
    v3 := []string{"", "\n\t", ("hello" + ""), (("a" + ("hello" + "\n\t")) + ("foo bar" + (("a" + "") + (("hello" + "") + (((("" + ("foo bar" + "\n\t")) + "foo bar") + ("a" + ("hello" + ("\n\t" + "hello")))) + "")))))}
    v1 = (v1 * -901)
    if true {
        for fi4 := 0; fi4 < 20; fi4++ {
            if len(v2) > 0 {
                v2[0] = v1
            }
            var v5 uint8 = 42
            _ = fi4
            _ = v5
        }
        if (false && (511 <= 2)) {
            v6 := GenS1{F0: v1, F1: "a"}
            switch (v1 * -2147483648) {
            case 311:
                for fi7 := 0; fi7 < 4; fi7++ {
                    if true {
                        _ = len(v3)
                    }
                    for fi8 := 0; fi8 < 3; fi8++ {
                        fmt.Println(v1)
                        v1 = (^(v1))
                        var v9 int = -1
                        _ = fi8
                        _ = v9
                    }
                    var v10 float32 = 0.1
                    _ = fi7
                    _ = v10
                }
            case 128:
                var v11 uint32 = 2147483648
                _ = v11
            }
            _ = v6
        } else {
            v1 += 32768
            v1 -= -197
        }
    } else {
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 11; fi13++ {
                v1 -= 4294967295
                switch fi12 {
                case -1:
                    fmt.Println(fi12, v1)
                case 9223372036854775807:
                    switch (-(fi12)) {
                    case 2147483647:
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    }
                default:
                    v2 = append(v2, 842)
                }
                var v14 uint64 = 18446744073709551615
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        if true {
            v1 += (v1 & v1)
            var v15 int32 = -2147483648
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 uint32 = 2147483648
                var v18 bool = false
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        } else {
            v1 = v1
        }
        fmt.Println(v1)
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        v1 = fi19
        v1 += v1
        fmt.Println(v1)
        _ = fi19
    }
    switch ("a" + "") {
    case "foo bar":
        v1 += v1
    }
    fmt.Println(v1)
    var v20 float32 = 3.14159
    v21 := GenS1{F0: v1, F1: ("" + ("foo bar" + (("\n\t" + "foo bar") + "a")))}
    var v22 uint = 2
    var v23 uint16 = 32767
    switch "hello" {
    case "":
        for fi24 := 0; fi24 < 19; fi24++ {
            var v25 uint64 = 42
            switch "foo bar" {
            case "a":
                v1 += v1
            case "":
                if false {
                    if true {
                        switch (^(fi24)) {
                        case -2147483648:
                            _ = len(v2)
                        case 127:
                            v26 := GenS2{F0: (v22 << ((v22) & 63)), F1: 1}
                            _ = v26
                        case 9223372036854775807:
                            v1 -= v1
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v1 = 9223372036854775807
                        fmt.Println(fi27)
                        v1 = 9223372036854775807
                        v20 = -2.25
                        _ = fi27
                    }
                }
            default:
                for fi28 := 0; fi28 < 4; fi28++ {
                    fmt.Println(v25, fi24, v23)
                    _ = fi28
                }
            }
            var v29 uint64 = 1
            _ = fi24
            _ = v25
            _ = v29
        }
    case "\n\t":
        if (-9223372036854775808 >= 892) {
            if false {
                v3 = append(v3, "foo bar")
                var v30 float32 = (v20 + float32(1e10))
                _ = v30
            } else {
                var v31 uint8 = 255
                switch "\n\t" {
                case "foo bar":
                    var v32 uint64 = 932
                    _ = v32
                case "a":
                    _ = len(v2)
                case "":
                    v31 = (v31 * (v31 % ((82) | 1)))
                }
                v33 := GenS2{F0: (v22 | (^(v22))), F1: 127}
                _ = v31
                _ = v33
            }
            if false {
                if (18446744073709551615 <= 669) {
                    if true {
                        v23 += v23
                    }
                } else {
                    var v34 uint8 = 142
                    v35 := GenS2{F0: v22, F1: 2}
                    v22 = v22
                    _ = v34
                    _ = v35
                }
                var v36 int16 = 42
                if (false && (123 >= 125)) {
                    var v37 int16 = v36
                    switch "" {
                    case "a":
                        fmt.Println(v1)
                    }
                    var v38 float64 = 3.14159
                    _ = v37
                    _ = v38
                } else {
                    var v39 uint16 = 255
                    for fi40 := 0; fi40 < 20; fi40++ {
                        var v41 string = "\n\t"
                        v42 := GenS3{F0: fi40}
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    _ = v39
                }
                _ = v36
            } else {
                for fi43 := 0; fi43 < 4; fi43++ {
                    _ = len(v2)
                    v44 := GenS1{F0: fi43, F1: ("a" + ("" + "hello"))}
                    switch (-(fi43)) {
                    case 127:
                        v45 := GenS3{F0: (fi43 | 204)}
                        _ = v45
                    case -2147483648:
                        fmt.Println(v23)
                    default:
                        v46 := GenS1{F0: 32768, F1: ""}
                        _ = v46
                    }
                    _ = fi43
                    _ = v44
                }
                v1 = v1
                var v47 int32 = 2147483647
                _ = v47
            }
            var v48 int64 = 0
            _ = v48
        }
    case "foo bar":
        v20 = (0.0 * v20)
    }
    var v49 int8 = 78
    if (-1 <= int64(v23)) {
        v50 := GenS1{F0: (v1 | v1), F1: "foo bar"}
        for fi51 := 0; fi51 < 1; fi51++ {
            for fi52 := 0; fi52 < 17; fi52++ {
                if (!true) {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = ("a" + "foo bar")
                        }
                        switch v1 {
                        case 42:
                            var v53 int32 = -2147483648
                            _ = v53
                        }
                        fmt.Println(v22, v49, v1)
                    }
                    var v54 uint8 = 127
                    _ = v54
                }
                v49 -= v49
                var v55 float64 = (3.14159 - 1.5)
                _ = fi52
                _ = v55
            }
            v20 = (v20 / 3.14159)
            _ = fi51
        }
        fmt.Println(v22)
        _ = v50
    } else {
        for fi56 := 0; fi56 < 12; fi56++ {
            v49 += v49
            _ = fi56
        }
        switch v1 {
        case 42:
            var v57 uint32 = 237
            _ = v57
        }
        var v58 uint32 = uint32(v22)
        _ = v58
    }
    v59 := GenS1{F0: 893, F1: "hello"}
    if true {
        v20 = v20
        var v60 float32 = 1e10
        switch -2147483648 {
        case 2147483647:
            fmt.Println(v23, v60)
        case 56:
            v49 = (v49 >> (((^(v22))) & 63))
        case 510:
            switch v1 {
            case 2147483647:
                v61 := GenS4{F0: uint32(v22)}
                _ = v61
            case 32767:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 32768:
                var v62 int = 4294967295
                _ = v62
            default:
                var v63 int64 = 1
                _ = v63
            }
        }
        _ = v60
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v64 uint16 = v23
    switch v1 {
    case 32767:
        if (!(!(((!(!false)) && false) || true))) {
            if (42 < 42) {
                for fi65 := 0; fi65 < 4; fi65++ {
                    v66 := GenS3{F0: -9223372036854775808}
                    if len(v2) > 0 {
                        v2[0] = (fi65 * v1)
                    }
                    _ = fi65
                    _ = v66
                }
                v64 += v23
            } else {
                var v67 uint16 = v23
                _ = v67
            }
            var v68 uint64 = 925
            _ = v68
        }
    case -100:
        if false {
            var v69 float32 = -2.25
            if (!(((0.1 / 1e10) * 0.1) != 0.0)) {
                if (0.1 == (-2.25 - v20)) {
                    fmt.Println(v1, v49)
                    var v70 uint16 = v23
                    if true {
                        v71 := GenS5{F0: v64}
                        var v72 bool = ((v70 % ((v64) | 1)) <= (v23 + (v70 / ((v70) | 1))))
                        _ = v71
                        _ = v72
                    }
                    _ = v70
                } else {
                    var v73 uint64 = 4294967295
                    _ = v73
                }
                if (v64 <= (v23 ^ v64)) {
                    fmt.Println(v22, v1, v64)
                    if true {
                        v2 = append(v2, v1)
                    } else {
                        var v74 int8 = 100
                        fmt.Println(v64, v22)
                        _ = v74
                    }
                } else {
                    if false {
                        v75 := GenS3{F0: v1}
                        _ = v75
                    } else {
                        v76 := GenS1{F0: v1, F1: "hello"}
                        switch "hello" {
                        case "a":
                            if true {
                                _ = len(v3)
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v76
                    }
                    v20 = float32(float64((float32(-2.25) + float32(0.1))))
                }
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v69
        } else {
            v49 = 1
        }
    case 42:
        v23 = v64
    default:
        if true {
            fmt.Println(v1)
            switch "" {
            case "\n\t":
                switch v1 {
                case 816:
                    fmt.Println(v20)
                default:
                    v3 = append(v3, "hello")
                }
            case "foo bar":
                for fi77 := 0; fi77 < 14; fi77++ {
                    v22 = (v22 & v22)
                    var v78 int64 = 0
                    _ = fi77
                    _ = v78
                }
            case "":
                v3 = append(v3, (("hello" + "a") + ("" + (("" + ("" + "\n\t")) + "a"))))
            }
        } else {
            var v79 int32 = 2147483647
            _ = v79
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v49:", v49)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v49
    _ = v59
    _ = v64
}
