package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS5 struct { F0 int32; F1 int16; F2 uint8 }
type GenS6 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint8, p1 int64) uint {
    v5 := []int{-2147483648}
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 2147483647)
    var v6 uint64 = 42
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    var v3 int32 = 42
    v4 := []string{("\n\t" + "a"), ""}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint64 = 105
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: -1, F1: 9223372036854775807}
            v11 := GenS2{F0: -874, F1: 232}
            _ = fi9
            _ = v10
            _ = v11
        }
        if (-1 != 9223372036854775807) {
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS2{F0: -1, F1: 10}
                v8 += v8
                if (false || (176 >= 42)) {
                    if false {
                        switch (fi7 * int(v8)) {
                        case -2147483648:
                            var v14 float64 = 1.5
                            _ = v14
                        case -940:
                            v15 := GenS2{F0: 32767, F1: 128}
                            _ = v15
                        }
                        v8 += (v8 << ((1) & 63))
                        switch ("hello" + ("foo bar" + "\n\t")) {
                        case "hello":
                            if true {
                                fmt.Println(v3)
                                delete(v2, 127)
                            } else {
                                v16 := GenS3{F0: 128, F1: 4294967295}
                                var v17 uint = genFn0(127, -525)
                                _ = v16
                                _ = v17
                            }
                        case "a":
                            var v18 uint16 = 0
                            _ = v18
                        default:
                            v19 := GenS3{F0: fi12, F1: fi12}
                            _ = v19
                        }
                        v8 -= (v8 | v8)
                    } else {
                        var v20 float32 = 0.1
                        _ = v20
                    }
                    var v21 float64 = 1.5
                    switch "a" {
                    case "hello":
                        var v22 int8 = 127
                        _ = v22
                    case "a":
                        if ((571 != 1) || false) {
                            if (!true) {
                                var v23 int = (fi7 / ((2147483647) | 1))
                                _ = v23
                            }
                            var v24 uint32 = 4294967295
                            v21 -= v21
                            _ = v24
                        } else {
                            var v25 float32 = 1e10
                            v25 = float32((float64(v25) / ((v21 - v21) + (v21 - (v21 - v21)))))
                            _ = v25
                        }
                    default:
                        v21 -= (v21 + 1e10)
                    }
                    _ = v21
                }
                _ = fi12
                _ = v13
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi7
        _ = v8
    }
    v3 = -2147483648
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    var v26 float32 = 1e10
    var v27 int32 = v3
    delete(v1, 2147483648)
    v28 := GenS4{F0: 2, F1: v3, F2: 65535}
    v29 := GenS1{F0: 42, F1: 9223372036854775807}
    var v30 int32 = v3
    v26 -= float32(0.0)
    switch "a" {
    case "foo bar":
        switch "foo bar" {
        case "hello":
            if (!(!(!(605 >= 2147483648)))) {
                _ = len(v4)
                _ = len(v2)
                _ = len(v4)
            } else {
                var v31 int8 = 100
                v26 = 3.14159
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v31
            }
        case "foo bar":
            var v32 int64 = -9223372036854775808
            _ = v32
        default:
            if false {
                if true {
                    var v33 uint = 1
                    if ((false || (-1 <= 1000)) && ((false || false) && false)) {
                        var v34 uint64 = 609
                        _ = v34
                    } else {
                        fmt.Println(v30, v26, v30)
                        var v35 uint = (v33 & (^(v33)))
                        if (true || false) {
                            v27 += (^(v30))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v36 int8 = -1
                            _ = v36
                        }
                        _ = v35
                    }
                    if (!false) {
                        v30 -= (^(v30))
                        var v37 float32 = v26
                        v38 := GenS4{F0: 292, F1: v30, F2: 734}
                        _ = v37
                        _ = v38
                    }
                    _ = v33
                }
                v39 := GenS5{F0: (^(v27)), F1: 1, F2: 42}
                _ = v39
            } else {
                v40 := GenS2{F0: -512, F1: 255}
                _ = v40
            }
        }
    case "a":
        var v41 uint32 = 65535
        _ = v41
    default:
        if true {
            if (true && true) {
                var v42 uint64 = 2
                var v43 uint64 = 0
                _ = v42
                _ = v43
            } else {
                v44 := GenS2{F0: -32768, F1: 2}
                _ = v44
            }
            fmt.Println(v30, v26, v3)
            fmt.Println(v30, v3, v3)
        }
    }
    if true {
        v3 += 1
        var v45 float32 = (float32(0.1) / 1e10)
        for fi46 := 0; fi46 < 15; fi46++ {
            if ((true && (fi46 < 127)) || (v26 > (v26 * 0.1))) {
                v27 += -834
                for fi47 := 0; fi47 < 13; fi47++ {
                    switch (("\n\t" + ("\n\t" + "\n\t")) + "\n\t") {
                    case "\n\t":
                        var v48 int32 = (v3 + (v30 << ((0) & 63)))
                        _ = v48
                    default:
                        v49 := GenS4{F0: 42, F1: -1, F2: 1}
                        _ = v49
                    }
                    v27 = v27
                    v50 := GenS6{F0: 0, F1: 0}
                    _ = fi47
                    _ = v50
                }
            } else {
                var v51 string = "foo bar"
                _ = v51
            }
            _ = fi46
        }
        _ = v45
    }
    v27 -= v27
    switch ((("" + "\n\t") + "\n\t") + ("foo bar" + "")) {
    case "hello":
        v52 := v2[-1]
        _ = v52
    default:
        switch 128 {
        case -857:
            var v53 int16 = -32768
            _ = v53
        case 570:
            for fi54 := 0; fi54 < 16; fi54++ {
                switch ("" + (("\n\t" + (("\n\t" + "foo bar") + "foo bar")) + "")) {
                case "a":
                    if true {
                        for fi55 := 0; fi55 < 17; fi55++ {
                            if (!true) {
                                v56 := GenS1{F0: -128, F1: 1}
                                var v57 int64 = -9223372036854775808
                                _ = v56
                                _ = v57
                            }
                            if (18446744073709551615 > 18446744073709551615) {
                                switch ("hello" + "\n\t") {
                                case "":
                                    v2[(fi54 | (fi55 & fi55))] = ("foo bar" + "foo bar")
                                default:
                                    if true {
                                        var v58 bool = false
                                        switch (fi55 - 897) {
                                        case -54:
                                            v58 = (("hello" + ("\n\t" + ("a" + "foo bar"))) > "hello")
                                        case 2147483647:
                                            v58 = true
                                        case 4294967295:
                                            var v59 int8 = 1
                                            _ = v59
                                        }
                                        _ = v58
                                    }
                                }
                            } else {
                                var v60 float32 = v26
                                v27 += (v3 + v3)
                                v30 -= (-(v30))
                                _ = v60
                            }
                            _ = fi55
                        }
                    }
                case "\n\t":
                    if true {
                        var v61 string = ""
                        fmt.Println(v3, v27)
                        var v62 bool = true
                        _ = v61
                        _ = v62
                    } else {
                        for fi63 := 0; fi63 < 7; fi63++ {
                            var v64 string = "hello"
                            v27 = (v27 | 0)
                            v65 := GenS3{F0: 128, F1: 4294967295}
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                    }
                case "":
                    var v66 uint32 = 2
                    _ = v66
                default:
                    var v67 uint16 = 1
                    _ = v67
                }
                switch fi54 {
                case 128:
                    v3 += v27
                case -383:
                    delete(v2, (fi54 & (fi54 | fi54)))
                case -445:
                    var v68 int = (^(fi54))
                    _ = v68
                default:
                    v1[65535] = 128
                }
                switch fi54 {
                case 348:
                    v69 := v1[0]
                    _ = v69
                case 32768:
                    v70 := GenS2{F0: 0, F1: 2}
                    _ = v70
                default:
                    _ = len(v2)
                }
                _ = fi54
            }
        case -1:
            v4 = append(v4, "a")
        }
    }
    v30 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
