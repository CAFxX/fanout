package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int32; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 string; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS6 struct { F0 uint8 }
func genFn0() uint {
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 int32 = 2147483647
        _ = fi4
        _ = v5
    }
    switch (("foo bar" + (("hello" + "\n\t") + ("\n\t" + (("" + ("\n\t" + "\n\t")) + ("a" + ("" + ("\n\t" + (("foo bar" + "") + ("foo bar" + ((((("\n\t" + (("" + "hello") + "foo bar")) + "hello") + "foo bar") + "a") + "\n\t")))))))))) + "foo bar") {
    case "":
        if false {
            for fi6 := 0; fi6 < 2; fi6++ {
                v7 := []int{fi6}
                _ = fi6
                _ = v7
            }
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := GenS1{F0: -536}
                var v10 int8 = 100
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 bool = true
            _ = v11
        } else {
            if true {
                if true {
                    switch 32767 {
                    case -347:
                        switch 623 {
                        case 1:
                            var v12 bool = (128 < 32768)
                            _ = v12
                        case 9223372036854775807:
                            var v13 uint16 = 42
                            _ = v13
                        default:
                            if (662 > 2147483648) {
                                var v14 uint32 = 353
                                _ = v14
                            }
                        }
                    case 696:
                        var v15 string = "a"
                        _ = v15
                    }
                    v16 := GenS1{F0: 32768}
                    var v17 int32 = 509
                    _ = v16
                    _ = v17
                } else {
                    var v18 uint = 326
                    _ = v18
                }
                var v19 int16 = 1
                _ = v19
            } else {
                var v20 bool = true
                v21 := []int{9223372036854775807, 2147483647}
                _ = v20
                _ = v21
            }
        }
    case "\n\t":
        v22 := []int{}
        _ = v22
    case "a":
        if (true && (true && false)) {
            var v23 uint32 = 983
            var v24 int16 = -1
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 uint64 = 0
                _ = fi25
                _ = v26
            }
            var v27 string = "hello"
            var v28 uint64 = 1
            _ = v27
            _ = v28
        }
    default:
        v29 := []int{}
        _ = v29
    }
    if true {
        var v30 string = (((("a" + "") + ("\n\t" + "hello")) + "foo bar") + "a")
        v31 := GenS2{F0: 652, F1: 274, F2: (!(true && (1000 < -289)))}
        _ = v30
        _ = v31
    }
    v32 := GenS2{F0: 18446744073709551615, F1: 0, F2: false}
    switch -541 {
    case -138:
        v33 := GenS2{F0: 466, F1: -2147483648, F2: false}
        _ = v33
    }
    _ = v32
    return 448
}
func genFn1(p0 int32) int16 {
    if ((("" + (("" + "\n\t") + "foo bar")) + "foo bar") <= "\n\t") {
        var v34 int = 1
        v35 := []int{(v34 & v34), v34, (v34 % ((v34) | 1)), 114}
        _ = v34
        _ = v35
    }
    v36 := []int{}
    if len(v36) > 0 {
        v36[0] = 718
    }
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    _ = p0
    _ = v36
    return -419
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint16 = 373
    v3 := make(map[uint32]uint32)
    if false {
        v37 := GenS3{F0: 377}
        v38 := []int{32767, 2147483647, 32768, 2147483647}
        for fi39 := 0; fi39 < 12; fi39++ {
            v1 = v1
            _ = fi39
        }
        _ = v37
        _ = v38
    } else {
        var v40 float32 = 1.5
        _ = v40
    }
    v41 := GenS4{F0: 1, F1: ("" + ((("hello" + ("a" + "hello")) + ("\n\t" + "foo bar")) + "foo bar")), F2: 1}
    if true {
        var v42 float64 = v1
        _ = v42
    } else {
        v3[0] = 2
        for fi43 := 0; fi43 < 1; fi43++ {
            var v44 bool = ((!((fi43 ^ fi43) == fi43)) || true)
            _ = fi43
            _ = v44
        }
    }
    v45 := []int{127, 32767, -646, 984}
    for fi46 := 0; fi46 < 12; fi46++ {
        v47 := v3[4294967295]
        v48 := GenS1{F0: fi46}
        var v49 uint = 255
        _ = fi46
        _ = v47
        _ = v48
        _ = v49
    }
    fmt.Println(v1, v2)
    if false {
        if (true || true) {
            if false {
                var v50 bool = true
                if v50 {
                    for fi51 := 0; fi51 < 8; fi51++ {
                        var v52 float32 = (-2.25 + 1e10)
                        _ = len(v45)
                        _ = fi51
                        _ = v52
                    }
                    if v50 {
                        fmt.Println(v1, v50)
                        if v50 {
                            v53 := GenS5{F0: 2, F1: 2147483647, F2: v2}
                            fmt.Println(v1, v50, v2)
                            switch 2147483647 {
                            case 32768:
                                v50 = (genFn1(0) < 42)
                            case 9223372036854775807:
                                v3[2147483647] = 2147483647
                            }
                            _ = v53
                        } else {
                            v2 = v2
                            switch "foo bar" {
                            case "hello":
                                v50 = (1 != 0)
                            case "foo bar":
                                v50 = v50
                            case "a":
                                _ = len(v45)
                            }
                        }
                    } else {
                        var v54 int64 = -9223372036854775808
                        fmt.Println(v1)
                        _ = v54
                    }
                }
                _ = v50
            }
            v1 += v1
        } else {
            for fi55 := 0; fi55 < 19; fi55++ {
                var v56 uint16 = (v2 + (v2 + 2))
                _ = fi55
                _ = v56
            }
        }
    } else {
        var v57 int16 = 32767
        switch "a" {
        case "foo bar":
            var v58 uint32 = 924
            _ = v58
        case "hello":
            var v59 int = -928
            _ = v59
        }
        v60 := v3[2147483647]
        _ = v57
        _ = v60
    }
    for fi61 := 0; fi61 < 18; fi61++ {
        if false {
            var v62 float64 = (0.0 - -2.25)
            switch "\n\t" {
            case "":
                var v63 int8 = 127
                _ = v63
            default:
                var v64 string = ("a" + "foo bar")
                _ = v64
            }
            var v65 bool = (!true)
            _ = v62
            _ = v65
        } else {
            if false {
                fmt.Println(fi61, v1)
                for fi66 := 0; fi66 < 16; fi66++ {
                    delete(v3, 0)
                    _ = len(v3)
                    _ = fi66
                }
            } else {
                fmt.Println(fi61, v1)
            }
        }
        v1 = v1
        v2 = (v2 - (v2 - v2))
        _ = fi61
    }
    if ((18446744073709551615 == genFn0()) && ((1.5 < ((v1 * v1) * (float64(1.5) + v1))) && (true && ((((float64(3.14159) * float64(-2.25)) - v1) / v1) == (((float64(((float32(-2.25) + (3.14159 / float32(float64((1e10 * ((0.1 + 1e10) * (0.1 - ((1.5 / 0.1) - 0.1)))))))) - -2.25)) - float64(0.1)) / 1.0) * 1e10))))) {
        v45 = append(v45, -9223372036854775808)
    } else {
        v1 += v1
    }
    fmt.Println(v1)
    if (float32(v1) == (3.14159 + 1e10)) {
        var v67 string = ""
        v68 := GenS6{F0: 197}
        _ = v67
        _ = v68
    } else {
        if len(v45) >= 2 {
            v45 = v45[1:len(v45)]
        }
        var v69 int8 = 0
        _ = v69
    }
    var v70 uint32 = 1
    _ = len(v45)
    var v71 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v45
    _ = v70
    _ = v71
}
