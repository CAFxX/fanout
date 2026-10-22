package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint }

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]bool)
    var v3 uint16 = 65535
    if (!(0 < -2147483648)) {
        v4 := GenS1{F0: 518}
        _ = v4
    }
    var v5 float32 = 1.5
    var v6 int = -498
    for fi7 := 0; fi7 < 12; fi7++ {
        if true {
            var v8 uint = 18446744073709551615
            var v9 string = ("" + "foo bar")
            var v10 uint16 = 32767
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    v1 += v1
    v1 = (^(v1))
    var v11 string = ("\n\t" + "\n\t")
    var v12 uint32 = 0
    for fi13 := 0; fi13 < 20; fi13++ {
        switch (("" + ("a" + "\n\t")) + "a") {
        case "\n\t":
            var v14 uint32 = v12
            _ = v14
        case "foo bar":
            var v15 int64 = 9223372036854775807
            _ = v15
        case "a":
            switch (v11 + "\n\t") {
            case "hello":
                v11 = "\n\t"
            case "a":
                for fi16 := 0; fi16 < 19; fi16++ {
                    fmt.Println(v5, v11)
                    v6 = (^(fi13))
                    _ = fi16
                }
            }
        default:
            v12 -= 0
        }
        v6 += fi13
        if (255 != uint8(v3)) {
            v3 -= v3
            for fi17 := 0; fi17 < 20; fi17++ {
                if true {
                    fmt.Println(v3)
                    var v18 int64 = 432
                    _ = v18
                }
                _ = fi17
            }
            delete(v2, v11)
        } else {
            for fi19 := 0; fi19 < 19; fi19++ {
                if (159 == 244) {
                    v6 = 2147483647
                    var v20 float64 = (float64(1.5) / -2.25)
                    _ = v20
                }
                var v21 uint8 = 128
                v1 -= -535
                _ = fi19
                _ = v21
            }
        }
        _ = fi13
    }
    var v22 int8 = -1
    if true {
        for fi23 := 0; fi23 < 11; fi23++ {
            if true {
                var v24 float64 = -2.25
                var v25 int8 = int8(v1)
                for fi26 := 0; fi26 < 8; fi26++ {
                    fmt.Println(v22, v11)
                    v1 -= 1000
                    var v27 string = "hello"
                    _ = fi26
                    _ = v27
                }
                _ = v24
                _ = v25
            }
            for fi28 := 0; fi28 < 9; fi28++ {
                _ = len(v2)
                if (1000 <= v1) {
                    v22 -= (^(v22))
                    _ = len(v2)
                    v29 := GenS1{F0: 4294967295}
                    switch (-(fi23)) {
                    case 47:
                        if ((v22 << ((65535) & 63)) != (v22 - v22)) {
                            v30 := GenS2{F0: fi28, F1: ("" + (("" + v11) + "hello"))}
                            fmt.Println(v6, v1, v5)
                            v31 := GenS3{F0: (v11 + v11), F1: 65}
                            _ = v30
                            _ = v31
                        } else {
                            v32 := []int{(^(fi28))}
                            v22 += (v22 - (-(v22)))
                            _ = v32
                        }
                    case -182:
                        v33 := GenS1{F0: 2}
                        _ = v33
                    case 0:
                        v6 += fi28
                    }
                    v34 := GenS4{F0: (v22 - 1), F1: (("hello" + "foo bar") >= v11)}
                    _ = v29
                    _ = v34
                } else {
                    if ((v11 + "foo bar") != v11) {
                        v6 -= (v6 & v6)
                    }
                    var v35 uint = 0
                    _ = v35
                }
                _ = fi28
            }
            _ = fi23
        }
        if false {
            v12 = 2147483648
        } else {
            var v36 float64 = float64(0.0)
            v2[v11] = false
            _ = v36
        }
        v37 := GenS5{F0: 2147483647}
        _ = v37
    }
    v38 := GenS2{F0: v6, F1: "\n\t"}
    for fi39 := 0; fi39 < 19; fi39++ {
        v40 := GenS2{F0: (fi39 * -851), F1: v11}
        v12 = v12
        for fi41 := 0; fi41 < 6; fi41++ {
            switch v6 {
            case 131:
                v12 += v12
            case -667:
                v42 := []int{}
                _ = v42
            }
            fmt.Println(v5, v5)
            _ = fi41
        }
        _ = fi39
        _ = v40
    }
    if false {
        switch v11 {
        case "a":
            switch ("" + v11) {
            case "":
                var v43 uint64 = 325
                _ = v43
            case "hello":
                v44 := []int{v6, v6, (-(v6))}
                _ = v44
            }
        case "":
            delete(v2, ((((v11 + (v11 + ((v11 + ((((("foo bar" + v11) + v11) + (v11 + "hello")) + v11) + v11)) + (v11 + "")))) + v11) + "\n\t") + v11))
        }
        switch 424 {
        case 1:
            if true {
                var v45 string = (v11 + v11)
                _ = v45
            } else {
                for fi46 := 0; fi46 < 4; fi46++ {
                    v47 := v2[v11]
                    var v48 int16 = v1
                    fmt.Println(v47)
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                switch (-(v6)) {
                case -9223372036854775808:
                    if false {
                        var v49 string = v11
                        _ = v49
                    } else {
                        for fi50 := 0; fi50 < 9; fi50++ {
                            if false {
                                v51 := GenS1{F0: 42}
                                for fi52 := 0; fi52 < 15; fi52++ {
                                    switch v11 {
                                    case "foo bar":
                                        var v53 uint64 = 1
                                        _ = v53
                                    }
                                    _ = fi52
                                }
                                _ = v51
                            } else {
                                if (v1 != 202) {
                                    v54 := []int{32767}
                                    switch (v11 + v11) {
                                    case "":
                                        var v55 int8 = v22
                                        _ = v55
                                    case "hello":
                                        var v56 bool = false
                                        _ = v56
                                    }
                                    _ = v54
                                } else {
                                    delete(v2, ("foo bar" + v11))
                                }
                                switch v6 {
                                case 32767:
                                    if (float64(v5) == -2.25) {
                                        v57 := GenS6{F0: v6, F1: -9223372036854775808}
                                        _ = v57
                                    } else {
                                        switch "foo bar" {
                                        case "\n\t":
                                            v58 := []int{v6, v6}
                                            _ = v58
                                        case "a":
                                            v59 := GenS1{F0: 26}
                                            _ = v59
                                        case "":
                                            v60 := []int{v6, (fi50 >> ((18446744073709551615) & 63)), (v6 / ((v6) | 1))}
                                            _ = v60
                                        }
                                        var v61 int8 = (v22 % ((0) | 1))
                                        _ = len(v2)
                                        _ = v61
                                    }
                                default:
                                    v1 += v1
                                }
                                v62 := GenS1{F0: 18446744073709551615}
                                _ = v62
                            }
                            _ = fi50
                        }
                    }
                case -601:
                    v3 = 65535
                case 22:
                    fmt.Println(v1, v12)
                default:
                    switch "" {
                    case "a":
                        _ = len(v2)
                    }
                }
                fmt.Println(v1)
            }
        case -627:
            var v63 int8 = -92
            _ = v63
        }
    } else {
        v64 := v2[""]
        _ = v64
    }
    for fi65 := 0; fi65 < 12; fi65++ {
        for fi66 := 0; fi66 < 20; fi66++ {
            if true {
                switch -701 {
                case 0:
                    switch fi66 {
                    case 42:
                        var v67 uint32 = v12
                        _ = v67
                    case 127:
                        v68 := []int{}
                        _ = v68
                    case -9223372036854775808:
                        switch (fi66 - fi65) {
                        case 32768:
                            var v69 int16 = (-(v1))
                            _ = v69
                        case 694:
                            var v70 uint16 = (v3 ^ v3)
                            _ = v70
                        default:
                            fmt.Println(v11, v3)
                        }
                    }
                case -609:
                    v71 := GenS2{F0: (fi65 | (-(fi65))), F1: v11}
                    _ = v71
                case 451:
                    v72 := []int{(fi66 - (-(fi66)))}
                    _ = v72
                }
                v6 += (fi66 % ((785) | 1))
                v3 -= 42
            }
            if ((v1 >> ((473) & 63)) != -32768) {
                switch v11 {
                case "hello":
                    v73 := []int{v6, fi65}
                    _ = v73
                case "\n\t":
                    switch v11 {
                    case "foo bar":
                        fmt.Println(v3, v11, v12)
                    }
                case "foo bar":
                    v22 -= 93
                    if false {
                        v74 := GenS7{F0: uint(v12)}
                        _ = v74
                    } else {
                        v3 += v3
                    }
                }
                v3 -= (v3 ^ uint16(v12))
            }
            _ = fi66
        }
        _ = fi65
    }
    var v75 int64 = 1
    v76 := []int{v6, v6}
    _ = len(v76)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v75:", v75)
    fmt.Println("slicelen v76:", len(v76), v76)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v38
    _ = v5
    _ = v6
    _ = v75
    _ = v76
}
