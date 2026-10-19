package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 string; F2 uint64 }
type GenS6 struct { F0 string; F1 int8; F2 uint64 }
func genFn0(p0 int16) int64 {
    p0 += p0
    var v5 uint8 = 228
    var v6 uint = 1
    p0 = (p0 >> ((2) & 63))
    v7 := []int{}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int8, p1 bool, p2 int) string {
    v8 := GenS1{F0: 2147483647, F1: 212, F2: 18446744073709551615}
    fmt.Println(p2, p2)
    v9 := []int{}
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := GenS2{F0: 42, F1: 42}
        fmt.Println(p2, p2, p0)
        fmt.Println(p0, fi10, p2)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return ""
}

func main() {
    v1 := []uint32{18, 240, 958, 32}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    v4 := []int{32767, -1, 42, -528}
    if (1000 > 42) {
        v12 := GenS3{F0: -34, F1: (-675 != 1)}
        _ = v12
    } else {
        delete(v2, 2147483647)
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    var v14 int8 = -30
    var v15 string = ("\n\t" + "foo bar")
    for fi16 := 0; fi16 < 1; fi16++ {
        if ((4294967295 > 2147483647) || (true && (!true))) {
            fmt.Println(v15, v14, v15)
            v15 = ((v15 + v15) + v15)
            switch "" {
            case "foo bar":
                if false {
                    var v17 bool = ("a" == "a")
                    _ = v17
                }
            case "":
                if (!true) {
                    var v18 bool = (-9223372036854775808 < genFn0(42))
                    _ = len(v2)
                    _ = v18
                } else {
                    switch (^(fi16)) {
                    case 128:
                        v19 := GenS2{F0: 2147483647, F1: 127}
                        _ = v19
                    case 127:
                        for fi20 := 0; fi20 < 6; fi20++ {
                            switch (fi16 % ((fi20) | 1)) {
                            case -2147483648:
                                v3[450] = 32767
                            case -1:
                                v21 := GenS3{F0: 0, F1: ((!(!((fi20 * fi20) <= fi16))) && true)}
                                _ = v21
                            case 32768:
                                v15 = ""
                            }
                            _ = fi20
                        }
                    default:
                        var v22 uint8 = 213
                        _ = v22
                    }
                    var v23 float64 = 0.0
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v25 := GenS4{F0: -747}
                        v15 = v15
                        v15 = v15
                        _ = fi24
                        _ = v25
                    }
                    _ = v23
                }
            default:
                switch (((((v15 + v15) + v15) + (v15 + v15)) + (((v15 + ((((v15 + v15) + (v15 + (v15 + ((v15 + ((v15 + (("" + v15) + ("foo bar" + ("\n\t" + v15)))) + v15)) + "foo bar")))) + v15) + v15)) + v15) + v15)) + v15) {
                case "foo bar":
                    var v26 int32 = -1
                    _ = v26
                case "\n\t":
                    var v27 int32 = 42
                    _ = v27
                case "a":
                    for fi28 := 0; fi28 < 7; fi28++ {
                        v29 := GenS5{F0: 874, F1: v15, F2: 42}
                        switch v15 {
                        case "":
                            if (v15 >= "foo bar") {
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch fi16 {
                                    case 9223372036854775807:
                                        v3[2147483648] = fi16
                                    case 127:
                                        _ = len(v4)
                                    case 562:
                                        v15 = ("\n\t" + "a")
                                    default:
                                        var v30 int = fi16
                                        _ = v30
                                    }
                                } else {
                                    var v31 int32 = -2147483648
                                    _ = v31
                                }
                                var v32 bool = false
                                v14 -= 0
                                _ = v32
                            } else {
                                v14 -= 45
                                v14 -= v14
                                if (18446744073709551615 != 155) {
                                    _ = len(v3)
                                }
                                v15 = ""
                            }
                        case "hello":
                            if false {
                                switch ((v15 + "foo bar") + v15) {
                                case "":
                                    if false {
                                        if ((194 > 131) && ((float32((0.1 - (1.5 * 1e10))) * (-2.25 / 3.14159)) > -2.25)) {
                                            v15 = "hello"
                                        }
                                        _ = len(v2)
                                    } else {
                                        var v33 int64 = genFn0(-1)
                                        v15 = v15
                                        _ = v33
                                    }
                                case "a":
                                    if true {
                                        v15 = "foo bar"
                                        var v34 uint64 = 2
                                        _ = v34
                                    }
                                default:
                                    var v35 int32 = 42
                                    _ = v35
                                }
                            } else {
                                if false {
                                    var v36 string = (v15 + ((v15 + ("hello" + (v15 + (v15 + v15)))) + "a"))
                                    if ((!(4294967295 < 4294967295)) && false) {
                                        v36 = "\n\t"
                                        var v37 int = 4294967295
                                        _ = v37
                                    } else {
                                        switch v36 {
                                        case "foo bar":
                                            fmt.Println(v36)
                                        case "hello":
                                            var v38 float64 = float64(1e10)
                                            switch fi16 {
                                            case 4294967295:
                                                var v39 int64 = -9223372036854775808
                                                _ = v39
                                            case -1:
                                                if (!(1 >= (fi16 - -66))) {
                                                    v14 -= v14
                                                    var v40 uint8 = 1
                                                    fmt.Println(v40)
                                                    _ = v40
                                                }
                                            case -9223372036854775808:
                                                v15 = ""
                                            }
                                            _ = v38
                                        default:
                                            v41 := GenS3{F0: 42, F1: false}
                                            _ = v41
                                        }
                                    }
                                    _ = v36
                                } else {
                                    if (!(false && (4294967295 > 2))) {
                                        if false {
                                            v15 = "a"
                                            v14 += v14
                                        }
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = 690
                                        }
                                        var v42 int16 = 0
                                        var v43 uint64 = 1
                                        _ = v42
                                        _ = v43
                                    }
                                    v15 = v15
                                    v44 := GenS3{F0: 42, F1: true}
                                    _ = v44
                                }
                                var v45 int16 = -930
                                _ = v45
                            }
                        case "foo bar":
                            v15 = (v15 + (("\n\t" + v15) + (v15 + "hello")))
                        default:
                            v14 = -1
                        }
                        _ = fi28
                        _ = v29
                    }
                default:
                    v46 := v3[2147483647]
                    _ = v46
                }
            }
        } else {
            switch "" {
            case "hello":
                switch "" {
                case "foo bar":
                    v47 := GenS1{F0: 42, F1: 42, F2: 65535}
                    _ = v47
                case "\n\t":
                    v14 = -128
                }
            default:
                switch (v15 + "foo bar") {
                case "a":
                    v14 -= v14
                case "":
                    if true {
                        var v48 uint16 = 42
                        var v49 bool = (1 != -676)
                        _ = v48
                        _ = v49
                    }
                case "\n\t":
                    v14 += v14
                default:
                    v15 = (v15 + v15)
                }
            }
            for fi50 := 0; fi50 < 1; fi50++ {
                v14 -= (v14 * 0)
                _ = fi50
            }
            fmt.Println(v14)
        }
        v51 := GenS5{F0: genFn0(1), F1: ((v15 + v15) + ("hello" + (("hello" + (v15 + "")) + v15))), F2: 332}
        var v52 float32 = 1e10
        _ = fi16
        _ = v51
        _ = v52
    }
    var v53 int = 44
    if (313 <= 335) {
        var v54 uint = 844
        _ = v54
    } else {
        for fi55 := 0; fi55 < 18; fi55++ {
            v56 := GenS6{F0: v15, F1: (v14 >> ((1) & 63)), F2: 2}
            _ = fi55
            _ = v56
        }
        v14 += v14
    }
    v14 += (v14 + 100)
    fmt.Println(v14)
    var v57 uint32 = 2147483647
    var v58 int64 = -1
    _ = len(v4)
    var v59 uint16 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v53:", v53)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v57
    _ = v58
    _ = v59
}
