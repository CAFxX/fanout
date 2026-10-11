package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 bool; F2 int32 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint16, p2 uint16) int32 {
    p1 -= 586
    fmt.Println(p1, p2)
    var v4 int16 = 1000
    v5 := []int{696, 2147483647, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -519
}
func genFn1(p0 int16) int {
    fmt.Println(p0)
    if false {
        v6 := GenS1{F0: -38}
        _ = v6
    } else {
        var v7 uint8 = 127
        var v8 float64 = ((3.14159 / ((-2.25 / float64((1.5 / 1.0))) + 3.14159)) / 3.14159)
        v7 = v7
        _ = v7
        _ = v8
    }
    _ = p0
    return 2147483647
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []string{"foo bar"}
    v3 := []uint32{0, 1, 0}
    var v9 uint8 = 0
    var v10 uint64 = 42
    switch ("hello" + (("" + ("hello" + (("a" + "a") + ("" + "")))) + "\n\t")) {
    case "hello":
        if ((true && false) && false) {
            var v11 bool = true
            v12 := GenS2{F0: 1}
            v10 = v10
            _ = v11
            _ = v12
        }
    }
    var v13 string = ("hello" + "")
    fmt.Println(v13)
    v1 -= (^(v1))
    var v14 int64 = 59
    v9 = (^(v9))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 5; fi16++ {
            if (-1 > 0) {
                if len(v2) > 0 {
                    v2[0] = v13
                }
            } else {
                var v17 float32 = (1.5 + 3.14159)
                _ = v17
            }
            v9 = 255
            if false {
                fmt.Println(fi15, v1, v10)
                v10 += v10
            } else {
                v18 := GenS1{F0: -1}
                var v19 uint16 = 0
                _ = v18
                _ = v19
            }
            _ = len(v2)
            _ = fi16
        }
        var v20 uint8 = (v9 * (^(v9)))
        var v21 string = v13
        _ = fi15
        _ = v20
        _ = v21
    }
    if (804 <= 100) {
        fmt.Println(v9, v14)
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 uint64 = v10
            _ = fi22
            _ = v23
        }
        v24 := GenS2{F0: genFn0(9223372036854775807, 32767, uint16(v9))}
        _ = v24
    } else {
        v13 = v13
        v25 := GenS1{F0: 127}
        var v26 uint16 = 1
        _ = v25
        _ = v26
    }
    if (-55 != 42) {
        if (!(-1 < 571)) {
            fmt.Println(v9)
            v9 = v9
        } else {
            switch 0 {
            case 4294967295:
                v3 = append(v3, 4294967295)
            default:
                v13 = v13
            }
            if true {
                if true {
                    for fi27 := 0; fi27 < 18; fi27++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi27
                    }
                    for fi28 := 0; fi28 < 1; fi28++ {
                        if (v1 <= v1) {
                            if false {
                                if false {
                                    var v29 int = (fi28 ^ 522)
                                    v9 = (v9 | (v9 - (v9 & (v9 / (((v9 & 127)) | 1)))))
                                    var v30 float32 = 0.0
                                    _ = v29
                                    _ = v30
                                }
                                v9 += v9
                                var v31 float64 = (1.5 - 3.14159)
                                _ = v31
                            } else {
                                if true {
                                    v32 := GenS2{F0: 42}
                                    var v33 int = (fi28 * (-(fi28)))
                                    _ = v32
                                    _ = v33
                                }
                                for fi34 := 0; fi34 < 7; fi34++ {
                                    if (false || true) {
                                        if false {
                                            if (!(58 != -12)) {
                                                v14 = (v14 - (v14 << (((^(v1))) & 63)))
                                                fmt.Println(v13, fi34, v13)
                                                var v35 uint8 = (v9 % (((v9 | v9)) | 1))
                                                _ = v35
                                            } else {
                                                var v36 int64 = (v14 - v14)
                                                _ = v36
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            switch fi34 {
                                            case -733:
                                                v10 += (v10 * v10)
                                            case 780:
                                                var v37 uint16 = 42
                                                _ = v37
                                            default:
                                                if false {
                                                    switch 127 {
                                                    case 4294967295:
                                                        if true {
                                                            v2 = append(v2, v13)
                                                        }
                                                    }
                                                    var v38 int8 = 42
                                                    switch fi34 {
                                                    case 4294967295:
                                                        v9 -= 255
                                                    case 322:
                                                        var v39 int = fi28
                                                        _ = v39
                                                    default:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    _ = v38
                                                } else {
                                                    switch fi34 {
                                                    case 42:
                                                        v3 = append(v3, 2147483647)
                                                    }
                                                    var v40 bool = false
                                                    var v41 uint = (v1 ^ v1)
                                                    _ = v40
                                                    _ = v41
                                                }
                                            }
                                            v42 := GenS3{F0: 42, F1: (v10 / (((v10 / (((^(v10))) | 1))) | 1)), F2: 2}
                                            if false {
                                                v9 -= (v9 % ((v9) | 1))
                                            } else {
                                                fmt.Println(v10, v10)
                                                fmt.Println(v1, fi28)
                                            }
                                            _ = v42
                                        }
                                        v43 := GenS4{F0: -1, F1: 4294967295}
                                        _ = v43
                                    } else {
                                        if (3.14159 == (3.14159 / ((0.0 / 1.0) - (0.1 / 0.1)))) {
                                            var v44 float32 = 3.14159
                                            var v45 int = genFn1(-404)
                                            _ = v44
                                            _ = v45
                                        } else {
                                            switch (v13 + (v13 + v13)) {
                                            case "hello":
                                                var v46 uint64 = (v10 - 345)
                                                _ = v46
                                            default:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            switch v13 {
                                            case "a":
                                                var v47 uint16 = 102
                                                _ = v47
                                            case "hello":
                                                var v48 int32 = -2147483648
                                                _ = v48
                                            }
                                            if false {
                                                switch ((("hello" + (v13 + v13)) + "\n\t") + (((v13 + v13) + ("hello" + v13)) + v13)) {
                                                case "\n\t":
                                                    if len(v3) > 0 {
                                                        v3[0] = 65535
                                                    }
                                                case "a":
                                                    var v49 int64 = v14
                                                    _ = v49
                                                }
                                                switch ((v13 + "") + v13) {
                                                case "hello":
                                                    if true {
                                                        fmt.Println(v13, fi34)
                                                        switch v13 {
                                                        case "hello":
                                                            v13 = "hello"
                                                        }
                                                    } else {
                                                        _ = len(v3)
                                                        if false {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v1 += (v1 - (v1 << (((v1 >> ((v1) & 63))) & 63)))
                                                        } else {
                                                            var v50 int8 = 1
                                                            if ((v50 >> (((v1 >> ((v1) & 63))) & 63)) == (-(v50))) {
                                                                fmt.Println(v13)
                                                            } else {
                                                                _ = len(v3)
                                                            }
                                                            _ = v50
                                                        }
                                                    }
                                                case "a":
                                                    fmt.Println(v9, fi34, v9)
                                                case "foo bar":
                                                    _ = len(v2)
                                                }
                                                v51 := GenS5{F0: 42, F1: 216}
                                                _ = v51
                                            } else {
                                                var v52 int16 = 389
                                                fmt.Println(fi28, v52, v52)
                                                v13 = (v13 + "a")
                                                _ = v52
                                            }
                                        }
                                        if (v14 == v14) {
                                            v9 -= (v9 >> ((v1) & 63))
                                        }
                                    }
                                    _ = fi34
                                }
                            }
                            var v53 int32 = 42
                            _ = len(v2)
                            _ = v53
                        } else {
                            v10 += (^(v10))
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (v13 >= v13) {
                            v54 := GenS6{F0: 32767}
                            for fi55 := 0; fi55 < 1; fi55++ {
                                switch (v13 + v13) {
                                case "foo bar":
                                    fmt.Println(v13, v1)
                                case "hello":
                                    switch "a" {
                                    case "foo bar":
                                        var v56 int = (^(fi55))
                                        _ = v56
                                    case "":
                                        var v57 int32 = 0
                                        _ = v57
                                    default:
                                        v10 = v10
                                    }
                                }
                                if len(v3) > 0 {
                                    v3[0] = 65535
                                }
                                v58 := GenS6{F0: -648}
                                _ = fi55
                                _ = v58
                            }
                            _ = v54
                        } else {
                            v2 = append(v2, v13)
                            if (genFn0(0, 2, 2) <= 2147483647) {
                                v59 := GenS7{F0: true, F1: (true || (true && ((fi28 | (fi28 / ((fi28) | 1))) > fi28))), F2: -882}
                                for fi60 := 0; fi60 < 20; fi60++ {
                                    v61 := GenS8{F0: 65535}
                                    _ = fi60
                                    _ = v61
                                }
                                _ = v59
                            } else {
                                fmt.Println(v13, v10)
                                v9 -= (v9 & 0)
                                v13 = v13
                            }
                        }
                        _ = fi28
                    }
                    var v62 uint16 = 42
                    _ = v62
                }
                v63 := GenS2{F0: 2147483647}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v63
            }
            var v64 uint32 = 905
            _ = v64
        }
        for fi65 := 0; fi65 < 12; fi65++ {
            if (65535 != 4294967295) {
                var v66 uint64 = v10
                var v67 int = (fi65 ^ fi65)
                v67 += genFn1(42)
                _ = v66
                _ = v67
            }
            v3 = append(v3, 165)
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            _ = fi65
        }
    } else {
        switch ((v13 + ("hello" + (v13 + ""))) + v13) {
        case "a":
            if false {
                var v68 int8 = 0
                _ = v68
            }
        case "":
            for fi69 := 0; fi69 < 6; fi69++ {
                fmt.Println(v9)
                _ = fi69
            }
        case "hello":
            v70 := GenS1{F0: 1}
            _ = v70
        }
    }
    var v71 bool = (-9223372036854775808 == int64(v10))
    fmt.Println(v10, v71)
    var v72 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v71
    _ = v72
    _ = v9
}
