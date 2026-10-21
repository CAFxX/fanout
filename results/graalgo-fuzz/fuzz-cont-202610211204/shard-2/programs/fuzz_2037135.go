package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 int8; F1 uint32; F2 string }
type GenS7 struct { F0 uint; F1 uint32; F2 bool }

func main() {
    var v1 uint32 = 426
    var v2 int32 = 2147483647
    v3 := []uint32{}
    var v4 float32 = 0.0
    v5 := make(map[int]int)
    v2 = (-(v2))
    delete(v5, 138)
    fmt.Println(v2, v1, v2)
    if true {
        for fi6 := 0; fi6 < 11; fi6++ {
            fmt.Println(v1, v4)
            for fi7 := 0; fi7 < 4; fi7++ {
                if (0.1 != 1.5) {
                    v8 := GenS1{F0: true, F1: int64(fi6)}
                    _ = v8
                }
                fmt.Println(fi6, v1)
                if (((!true) && false) || true) {
                    delete(v5, fi6)
                    if true {
                        switch "" {
                        case "a":
                            var v9 int = (fi7 * fi7)
                            _ = v9
                        default:
                            v1 = v1
                        }
                    }
                    var v10 int64 = 0
                    if len(v3) > 0 {
                        v3[0] = (v1 - (v1 ^ v1))
                    }
                    fmt.Println(v10, fi6, v2)
                    switch (("" + "") + ("\n\t" + ("\n\t" + "foo bar"))) {
                    case "hello":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v10
                } else {
                    v11 := GenS1{F0: (false && (!(!true))), F1: 813}
                    v5[(fi6 * fi6)] = 501
                    delete(v5, (fi7 - (fi6 + (-(fi7)))))
                    _ = len(v5)
                    if false {
                        v2 += (v2 | v2)
                        var v12 uint64 = 4294967295
                        switch (fi7 >> ((uint(v1)) & 63)) {
                        case 83:
                            fmt.Println(fi6, v4, v4)
                        case 32768:
                            if (42 <= 0) {
                                v4 = (v4 / float32((float64(v4) * -2.25)))
                                var v13 int16 = 0
                                _ = v13
                            } else {
                                fmt.Println(fi6, v12, v2)
                            }
                        case 20:
                            var v14 uint8 = 127
                            _ = v14
                        default:
                            var v15 int16 = 1000
                            _ = v15
                        }
                        v12 = (v12 / ((v12) | 1))
                        switch (("" + "hello") + ("a" + (("hello" + "a") + "\n\t"))) {
                        case "\n\t":
                            var v16 int16 = 32767
                            _ = v16
                        case "":
                            fmt.Println(v1)
                        default:
                            v4 = v4
                        }
                        _ = v12
                    }
                    v17 := GenS1{F0: false, F1: -9223372036854775808}
                    _ = v11
                    _ = v17
                }
                v2 -= -2147483648
                v1 = 2147483647
                v2 -= v2
                v2 = -2147483648
                v1 += (^(v1))
                _ = fi7
            }
            var v18 int8 = 1
            _ = fi6
            _ = v18
        }
        switch ("" + "") {
        case "":
            var v19 uint = 2
            _ = v19
        case "a":
            var v20 uint64 = 1
            _ = v20
        case "hello":
            if true {
                switch 910 {
                case 391:
                    v21 := GenS1{F0: (42 == 572), F1: 42}
                    _ = v21
                case -837:
                    if (0.0 <= 0.0) {
                        var v22 float64 = 1e10
                        for fi23 := 0; fi23 < 5; fi23++ {
                            for fi24 := 0; fi24 < 3; fi24++ {
                                v25 := GenS2{F0: 133}
                                v26 := GenS3{F0: 255, F1: 0, F2: 979}
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            var v27 uint16 = 2
                            for fi28 := 0; fi28 < 2; fi28++ {
                                var v29 int8 = int8(fi28)
                                switch "" {
                                case "a":
                                    v27 += (v27 << ((1) & 63))
                                }
                                switch "hello" {
                                case "a":
                                    v1 -= v1
                                case "":
                                    if false {
                                        var v30 bool = (uint64(v27) >= 18446744073709551615)
                                        if v30 {
                                            if (0 > (v1 >> ((626) & 63))) {
                                                v29 -= (v29 << ((18446744073709551615) & 63))
                                                v22 -= (-2.25 / v22)
                                                v31 := GenS4{F0: 293}
                                                _ = v31
                                            }
                                            if v30 {
                                                var v32 uint = 2
                                                _ = v32
                                            } else {
                                                var v33 float32 = v4
                                                var v34 int64 = 1
                                                _ = v33
                                                _ = v34
                                            }
                                        } else {
                                            var v35 string = "foo bar"
                                            _ = v35
                                        }
                                        _ = v30
                                    } else {
                                        var v36 uint8 = 42
                                        _ = v36
                                    }
                                case "\n\t":
                                    if false {
                                        var v37 uint16 = (v27 / ((v27) | 1))
                                        _ = len(v3)
                                        _ = v37
                                    } else {
                                        if true {
                                            switch "" {
                                            case "hello":
                                                var v38 int8 = v29
                                                _ = v38
                                            case "a":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v39 uint16 = (v27 << ((0) & 63))
                                                    _ = v39
                                                }
                                            case "foo bar":
                                                _ = len(v5)
                                            default:
                                                var v40 float64 = v22
                                                _ = v40
                                            }
                                            if ((2 >= 127) && ((v29 - -95) != v29)) {
                                                v41 := GenS1{F0: false, F1: 9223372036854775807}
                                                var v42 uint64 = 374
                                                delete(v5, (fi28 >> ((295) & 63)))
                                                fmt.Println(v4)
                                                _ = v41
                                                _ = v42
                                            }
                                            switch (fi23 >> ((18446744073709551615) & 63)) {
                                            case -429:
                                                v43 := GenS5{F0: -75, F1: (fi23 ^ fi28)}
                                                _ = v43
                                            default:
                                                var v44 uint16 = v27
                                                _ = v44
                                            }
                                        }
                                    }
                                }
                                _ = fi28
                                _ = v29
                            }
                            _ = fi23
                            _ = v27
                        }
                        var v45 bool = true
                        _ = v22
                        _ = v45
                    } else {
                        v46 := GenS3{F0: 255, F1: -1, F2: 42}
                        v47 := GenS1{F0: false, F1: 0}
                        v1 -= v1
                        _ = v46
                        _ = v47
                    }
                }
                var v48 bool = (((-2.25 + float32(((0.1 - (1.5 + 0.1)) / 0.1))) * v4) != (v4 * v4))
                var v49 int32 = (v2 | v2)
                _ = v48
                _ = v49
            } else {
                v1 = (v1 + (v1 & v1))
                v50 := GenS6{F0: 127, F1: v1, F2: ((("" + "a") + ("a" + ("" + "a"))) + "")}
                _ = v50
            }
        }
        v51 := GenS1{F0: (1 == 1), F1: 42}
        _ = v51
    }
    delete(v5, 128)
    if false {
        for fi52 := 0; fi52 < 13; fi52++ {
            var v53 bool = ("foo bar" == "")
            var v54 int8 = 1
            _ = fi52
            _ = v53
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 18; fi55++ {
        var v56 int = fi55
        _ = fi55
        _ = v56
    }
    var v57 int = 563
    fmt.Println(v57, v57)
    v1 = (v1 / ((v1) | 1))
    v58 := GenS7{F0: 18446744073709551615, F1: (v1 >> ((722) & 63)), F2: (!true)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
}
