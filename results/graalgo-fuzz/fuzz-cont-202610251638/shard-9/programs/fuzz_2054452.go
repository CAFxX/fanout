package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 uint64 }

func main() {
    var v1 uint8 = 131
    v2 := make(map[string]string)
    var v3 uint = 0
    v4 := []string{"hello", "foo bar", ((("" + ("\n\t" + ("a" + "hello"))) + "foo bar") + "")}
    var v5 int32 = 2147483647
    var v6 uint8 = v1
    v7 := GenS1{F0: 0}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 += 128
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint64 = 18446744073709551615
        if (-9223372036854775808 != 0) {
            for fi10 := 0; fi10 < 16; fi10++ {
                v6 = v1
                _ = fi10
            }
        }
        _ = fi8
        _ = v9
    }
    v1 -= uint8(v3)
    for fi11 := 0; fi11 < 9; fi11++ {
        if (((((1.5 * 0.0) * (0.1 - (0.0 - (0.0 * 0.0)))) * float32(0.1)) < 0.0) && (!true)) {
            fmt.Println(fi11)
            if (2 <= 379) {
                if ((1.5 - (1.5 - ((((0.0 - -2.25) * 1e10) * 3.14159) / (-2.25 * float32(1e10))))) <= float32(float64(((1.5 / 3.14159) + 3.14159)))) {
                    if true {
                        if (255 <= 42) {
                            v1 = (v1 >> ((v3) & 63))
                            v6 = v1
                            var v12 float64 = float64((3.14159 - (0.1 + 1e10)))
                            _ = v12
                        }
                    }
                    v3 += v3
                    if true {
                        switch (fi11 ^ fi11) {
                        case 0:
                            for fi13 := 0; fi13 < 4; fi13++ {
                                v6 -= v1
                                if true {
                                    fmt.Println(fi13)
                                    fmt.Println(fi11, fi13)
                                } else {
                                    fmt.Println(v3, fi11, v3)
                                    if true {
                                        v1 -= (v1 + v6)
                                        v6 -= v1
                                        var v14 uint32 = 1
                                        _ = v14
                                    } else {
                                        if true {
                                            fmt.Println(v3, v1, fi11)
                                            fmt.Println(v1)
                                            v3 -= (^(v3))
                                            fmt.Println(v3, fi11, v6)
                                        }
                                    }
                                    if true {
                                        _ = len(v2)
                                        v3 = v3
                                        var v15 uint16 = 1
                                        _ = v15
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = (("foo bar" + "a") + "a")
                                        }
                                        switch 128 {
                                        case 127:
                                            if false {
                                                if len(v4) > 0 {
                                                    v4[0] = "hello"
                                                }
                                                _ = len(v2)
                                                fmt.Println(fi13, v1, v3)
                                            } else {
                                                _ = len(v4)
                                                v3 += v3
                                            }
                                        case 128:
                                            switch "foo bar" {
                                            case "foo bar":
                                                v5 -= v5
                                            }
                                        case 42:
                                            if true {
                                                var v16 int16 = -1
                                                fmt.Println(v3, v16)
                                                if true {
                                                    fmt.Println(v6, v6, v6)
                                                    var v17 int16 = (v16 - (v16 & (v16 + (v16 & v16))))
                                                    v17 = v16
                                                    fmt.Println(fi11)
                                                    _ = v17
                                                }
                                                var v18 float64 = 1e10
                                                v16 -= (v16 + int16(v5))
                                                _ = v16
                                                _ = v18
                                            }
                                        }
                                        if (false && (!true)) {
                                            var v19 float32 = 0.1
                                            v19 = (0.0 / 1.0)
                                            v4 = append(v4, ("hello" + "foo bar"))
                                            var v20 float64 = 0.0
                                            _ = v19
                                            _ = v20
                                        } else {
                                            var v21 int32 = (-(v5))
                                            _ = v21
                                        }
                                        v1 = (^(v6))
                                    }
                                }
                                v22 := GenS1{F0: (fi13 | -2147483648)}
                                var v23 bool = ((!(4294967295 < 32768)) && true)
                                var v24 uint16 = uint16(v3)
                                _ = fi13
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        case 42:
                            v25 := GenS1{F0: 4294967295}
                            _ = v25
                        default:
                            _ = len(v4)
                        }
                    } else {
                        var v26 uint64 = 538
                        _ = v26
                    }
                } else {
                    v27 := GenS1{F0: fi11}
                    _ = len(v4)
                    _ = v27
                }
                switch ("foo bar" + "hello") {
                case "":
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 int8 = 1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi28
                        _ = v29
                    }
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                case "foo bar":
                    v3 += (^(v3))
                default:
                    v30 := GenS1{F0: int(v5)}
                    _ = v30
                }
                v5 -= (v5 & (v5 % ((v5) | 1)))
            } else {
                v3 = v3
            }
            var v31 uint16 = 0
            _ = v31
        }
        if (32767 != 2) {
            switch (fi11 / (((^(fi11))) | 1)) {
            case 1:
                v32 := GenS1{F0: fi11}
                _ = v32
            case 32768:
                var v33 uint = (v3 % ((42) | 1))
                _ = v33
            case 127:
                switch (fi11 + fi11) {
                case 128:
                    for fi34 := 0; fi34 < 17; fi34++ {
                        if true {
                            switch ("a" + ("a" + "")) {
                            case "":
                                v35 := GenS2{F0: v5, F1: 320, F2: 974}
                                _ = v35
                            default:
                                fmt.Println(fi34)
                            }
                            v36 := GenS3{F0: 418, F1: "foo bar"}
                            if true {
                                var v37 float64 = 0.0
                                _ = v37
                            }
                            _ = v36
                        }
                        if ("a" < ("hello" + "a")) {
                            v38 := GenS4{F0: (745 > 2)}
                            v39 := GenS4{F0: false}
                            _ = v38
                            _ = v39
                        } else {
                            var v40 float32 = float32(0.0)
                            v3 = 65535
                            switch 42 {
                            case 42:
                                if ((1 < 518) && (!true)) {
                                    if false {
                                        if true {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            if false {
                                                var v41 string = "\n\t"
                                                switch fi11 {
                                                case -609:
                                                    v3 += (^(v3))
                                                case 811:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                case -873:
                                                    switch v41 {
                                                    case "a":
                                                        v5 = (v5 * v5)
                                                    }
                                                default:
                                                    var v42 uint32 = 605
                                                    _ = v42
                                                }
                                                var v43 uint = v3
                                                _ = v41
                                                _ = v43
                                            }
                                        } else {
                                            switch (fi11 + -9223372036854775808) {
                                            case 32768:
                                                v44 := GenS1{F0: fi34}
                                                _ = v44
                                            case 711:
                                                if false {
                                                    v5 -= v5
                                                    fmt.Println(v3, v1, v5)
                                                    if (!true) {
                                                        v45 := v2["a"]
                                                        var v46 uint32 = 2147483648
                                                        v45 = v45
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        var v47 int64 = -9223372036854775808
                                                        var v48 float64 = -2.25
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    fmt.Println(v3, v40, v6)
                                                    v1 = (v6 >> ((v3) & 63))
                                                }
                                            case 4294967295:
                                                var v49 float32 = (v40 * v40)
                                                _ = v49
                                            }
                                            v5 = v5
                                            var v50 uint16 = 325
                                            switch "foo bar" {
                                            case "a":
                                                if false {
                                                    if ("" <= "hello") {
                                                        v3 -= 98
                                                    }
                                                    var v51 string = ((("hello" + ("" + "\n\t")) + "a") + ("a" + "a"))
                                                    _ = v51
                                                }
                                            default:
                                                v52 := GenS3{F0: 0, F1: "foo bar"}
                                                _ = v52
                                            }
                                            if (!(127 == 42)) {
                                                if len(v4) > 0 {
                                                    v4[0] = "a"
                                                }
                                                var v53 float64 = 3.14159
                                                v40 += v40
                                                fmt.Println(v6, v6)
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                if (v1 == (v1 ^ (v6 % ((v6) | 1)))) {
                                                    if (0.1 >= (v40 / v40)) {
                                                        delete(v2, "\n\t")
                                                        fmt.Println(fi11, v53)
                                                        if (!false) {
                                                            v3 -= v3
                                                            v54 := GenS4{F0: ((v40 + (v40 - float32(((v53 - 1.5) - (v53 / v53))))) > ((v40 + v40) / float32((v53 - v53))))}
                                                            _ = len(v2)
                                                            _ = v54
                                                        } else {
                                                            var v55 uint32 = 413
                                                            _ = v55
                                                        }
                                                    } else {
                                                        v3 += v3
                                                    }
                                                    v56 := GenS1{F0: (fi34 - fi34)}
                                                    _ = v56
                                                }
                                                v57 := v2[(("" + "") + ("a" + "a"))]
                                                if false {
                                                    v58 := GenS4{F0: (v57 != v57)}
                                                    _ = v58
                                                }
                                                v50 += (v50 << (((v3 * (^(v3)))) & 63))
                                                v59 := GenS5{F0: (v6 % ((v6) | 1)), F1: 719}
                                                _ = v53
                                                _ = v57
                                                _ = v59
                                            } else {
                                                var v60 int32 = -2147483648
                                                _ = v60
                                            }
                                            _ = v50
                                        }
                                        delete(v2, "foo bar")
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    } else {
                                        v61 := GenS1{F0: fi11}
                                        fmt.Println(v1, v5)
                                        _ = v61
                                    }
                                } else {
                                    _ = len(v2)
                                }
                            case 0:
                                if true {
                                    v62 := GenS2{F0: (v5 % (((v5 & v5)) | 1)), F1: v3, F2: 42}
                                    _ = v62
                                }
                            case 289:
                                var v63 int = fi11
                                _ = v63
                            }
                            _ = v40
                        }
                        _ = fi34
                    }
                case 1:
                    v64 := v2[""]
                    _ = v64
                }
            }
        }
        switch -9223372036854775808 {
        case 970:
            for fi65 := 0; fi65 < 16; fi65++ {
                var v66 float32 = 1.5
                var v67 bool = false
                fmt.Println(fi65, v1)
                _ = fi65
                _ = v66
                _ = v67
            }
        case -731:
            if (!true) {
                v1 -= (v1 | 0)
                _ = len(v4)
                v6 = 0
            } else {
                v1 = (^(v1))
                for fi68 := 0; fi68 < 9; fi68++ {
                    if ((v3 >= 1) && (!true)) {
                        var v69 float32 = -2.25
                        v70 := GenS6{F0: 9223372036854775807}
                        var v71 uint8 = (v1 * v1)
                        _ = v69
                        _ = v70
                        _ = v71
                    }
                    v72 := GenS7{F0: ("foo bar" + "\n\t"), F1: 18446744073709551615}
                    _ = fi68
                    _ = v72
                }
            }
        }
        _ = fi11
    }
    v73 := GenS6{F0: -764}
    var v74 int16 = -938
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v73
    _ = v74
}
