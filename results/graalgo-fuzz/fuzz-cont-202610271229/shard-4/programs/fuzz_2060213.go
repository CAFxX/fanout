package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 uint8; F2 int64 }
type GenS9 struct { F0 int64 }
func genFn0() uint64 {
    v4 := GenS1{F0: 2, F1: -128, F2: -1}
    var v5 int8 = 1
    var v6 int = -40
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -600
    var v3 int64 = (-(v2))
    if false {
        var v7 int32 = -2147483648
        v8 := GenS1{F0: 2, F1: 1, F2: 127}
        var v9 float32 = 1.5
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        if (465 >= 515) {
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := GenS2{F0: genFn0()}
                fmt.Println(v2)
                var v13 bool = ((3.14159 + 1.5) <= 0.0)
                _ = fi11
                _ = v12
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := []int{}
                if true {
                    fmt.Println(fi14, fi10, v3)
                }
                delete(v1, 596)
                switch (fi10 / ((9223372036854775807) | 1)) {
                case -1:
                    switch ("\n\t" + "hello") {
                    case "hello":
                        v16 := GenS3{F0: ("a" + ("a" + ("\n\t" + (("a" + ((("a" + ("hello" + "foo bar")) + "foo bar") + ("foo bar" + "\n\t"))) + ("hello" + (("hello" + ("hello" + "hello")) + ("hello" + ("" + "hello")))))))), F1: 127}
                        _ = v16
                    case "":
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    default:
                        if true {
                            switch (fi14 - (fi10 >> ((4294967295) & 63))) {
                            case -899:
                                if false {
                                    if true {
                                        var v17 uint8 = 11
                                        _ = v17
                                    } else {
                                        v3 = (v3 / ((v2) | 1))
                                        v18 := GenS4{F0: genFn0(), F1: 255}
                                        if (0 < v2) {
                                            switch fi10 {
                                            case 127:
                                                fmt.Println(v3, fi14)
                                            case 32768:
                                                v3 = (^(v3))
                                            default:
                                                if len(v15) > 0 {
                                                    v15[0] = fi10
                                                }
                                            }
                                            var v19 int32 = 981
                                            v20 := GenS2{F0: 392}
                                            _ = v19
                                            _ = v20
                                        } else {
                                            v1[483] = "hello"
                                            if (127 != 0) {
                                                if len(v15) >= 2 {
                                                    v15 = v15[1:len(v15)]
                                                }
                                                var v21 uint8 = 127
                                                if (!true) {
                                                    fmt.Println(fi10, v21)
                                                    v21 += 127
                                                    _ = len(v1)
                                                }
                                                _ = v21
                                            }
                                        }
                                        _ = v18
                                    }
                                    fmt.Println(v2, fi10)
                                    switch "foo bar" {
                                    case "\n\t":
                                        if len(v15) > 0 {
                                            v15[0] = fi10
                                        }
                                    case "foo bar":
                                        switch "hello" {
                                        case "":
                                            if (true || (-164 != 1)) {
                                                v22 := GenS5{F0: 314, F1: 1}
                                                var v23 int32 = -764
                                                var v24 uint16 = 65535
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            }
                                        case "foo bar":
                                            v3 -= v3
                                        case "a":
                                            var v25 int32 = 1
                                            _ = v25
                                        }
                                    }
                                } else {
                                    v2 += v3
                                }
                            case -793:
                                v26 := GenS4{F0: 4294967295, F1: 32767}
                                var v27 float32 = (float32((0.0 - float64((1.5 * (1e10 + 0.0))))) / 3.14159)
                                var v28 float32 = v27
                                if ((fi14 - fi14) > fi14) {
                                    v29 := GenS6{F0: false}
                                    v30 := GenS7{F0: false}
                                    v2 += (^(v3))
                                    var v31 float32 = 1.5
                                    if (!(!false)) {
                                        v2 += -1
                                        v2 -= v2
                                        if false {
                                            v32 := GenS5{F0: (fi14 + fi14), F1: 42}
                                            v3 = v2
                                            _ = v32
                                        } else {
                                            if (v2 <= int64(fi10)) {
                                                fmt.Println(fi10, v3, v3)
                                            }
                                            var v33 uint = 42
                                            _ = v33
                                        }
                                        fmt.Println(fi10, v2, v3)
                                        v31 -= (float32(3.14159) + (3.14159 + ((1.5 + (1e10 * v28)) + v27)))
                                    } else {
                                        var v34 bool = ((-(fi14)) > (fi14 << ((1) & 63)))
                                        v2 += (v3 - -9223372036854775808)
                                        v35 := GenS2{F0: 2}
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                                if false {
                                    v3 -= -1
                                    var v36 int = fi14
                                    _ = v36
                                }
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                            v37 := GenS8{F0: (fi14 & (fi10 / ((182) | 1))), F1: 2, F2: -9223372036854775808}
                            if (true && (fi14 < (fi10 | (fi14 % ((-9223372036854775808) | 1))))) {
                                var v38 string = "a"
                                var v39 uint = 987
                                var v40 uint = (v39 - 4294967295)
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            _ = v37
                        } else {
                            v2 = (v3 << ((480) & 63))
                        }
                    }
                }
                var v41 string = ""
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                if (!((v2 > v2) && true)) {
                    var v42 uint16 = 1
                    _ = v42
                }
                _ = fi14
                _ = v15
                _ = v41
            }
            v3 = v2
        }
        v3 += v2
        _ = fi10
    }
    var v43 int32 = -1
    if (42 <= 127) {
        for fi44 := 0; fi44 < 8; fi44++ {
            var v45 bool = false
            if (-2.25 > 1e10) {
                v46 := GenS5{F0: (-(fi44)), F1: -169}
                v47 := GenS1{F0: 0, F1: 13, F2: 730}
                _ = v46
                _ = v47
            }
            _ = fi44
            _ = v45
        }
    }
    v1[4294967295] = (("" + ("\n\t" + "a")) + ("hello" + "hello"))
    var v48 float32 = 1.5
    var v49 uint8 = 42
    fmt.Println(v43, v2, v49)
    v50 := v1[2147483647]
    for fi51 := 0; fi51 < 12; fi51++ {
        for fi52 := 0; fi52 < 10; fi52++ {
            v53 := v1[965]
            var v54 uint8 = (v49 - (v49 ^ (v49 >> ((65535) & 63))))
            if (false || false) {
                var v55 uint8 = v49
                v2 -= (v3 & (v3 >> ((2) & 63)))
                var v56 bool = ((!true) || ((2 < 678) || true))
                _ = v55
                _ = v56
            }
            var v57 uint = 745
            v58 := GenS1{F0: 0, F1: -128, F2: (^(fi51))}
            var v59 uint16 = 326
            v60 := []int{32768, (fi52 - fi51), fi51}
            _ = fi52
            _ = v53
            _ = v54
            _ = v57
            _ = v58
            _ = v59
            _ = v60
        }
        _ = fi51
    }
    var v61 float32 = float32(float64(-2.25))
    v50 = "\n\t"
    v62 := []int{128, 42, 128, 2147483647}
    var v63 string = (v50 + v50)
    v64 := GenS5{F0: 32768, F1: 0}
    v65 := GenS2{F0: 1}
    var v66 uint32 = 1
    switch ((v63 + v63) + "hello") {
    case "hello":
        v67 := GenS9{F0: (-(v3))}
        _ = v67
    case "foo bar":
        v49 = (v49 & (v49 | (v49 << ((18446744073709551615) & 63))))
    default:
        if true {
            v66 -= (v66 << ((987) & 63))
        } else {
            v43 -= (v43 % ((-1) | 1))
            for fi68 := 0; fi68 < 13; fi68++ {
                switch (-(fi68)) {
                case 709:
                    if true {
                        var v69 uint32 = v66
                        _ = v69
                    }
                case 134:
                    if (!(18446744073709551615 >= 301)) {
                        v63 = "hello"
                        v70 := GenS9{F0: (v3 % ((9223372036854775807) | 1))}
                        _ = v70
                    }
                }
                _ = fi68
            }
            var v71 uint16 = 255
            _ = v71
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v63:", v63)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v48
    _ = v49
    _ = v50
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v66
}
