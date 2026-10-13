package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 float32 = 3.14159
    var v4 uint = 1
    fmt.Println(v3)
    var v5 uint32 = 926
    v6 := []int{2147483647, -2147483648, 884, 42}
    switch ("" + "foo bar") {
    case "hello":
        v5 = v5
    default:
        if (false || true) {
            switch 666 {
            case -728:
                for fi7 := 0; fi7 < 17; fi7++ {
                    switch "\n\t" {
                    case "":
                        var v8 uint32 = 484
                        _ = v8
                    case "a":
                        v9 := GenS1{F0: 9223372036854775807, F1: (fi7 - (fi7 << ((0) & 63))), F2: 42}
                        _ = v9
                    default:
                        if false {
                            if true {
                                v10 := GenS1{F0: 42, F1: fi7, F2: -9223372036854775808}
                                _ = v10
                            }
                            var v11 int16 = 32767
                            _ = v11
                        } else {
                            fmt.Println(v4, v3)
                            switch "a" {
                            case "a":
                                var v12 uint = v4
                                _ = v12
                            default:
                                switch (fi7 / (((fi7 * 127)) | 1)) {
                                case 1:
                                    for fi13 := 0; fi13 < 2; fi13++ {
                                        var v14 int = fi7
                                        v5 += (^(v5))
                                        if false {
                                            var v15 uint32 = v5
                                            if (true || (!(false || (v14 == 128)))) {
                                                fmt.Println(v3, v15, fi13)
                                                _ = len(v2)
                                            } else {
                                                var v16 uint32 = (v5 / (((v5 << ((v4) & 63))) | 1))
                                                fmt.Println(fi13, v4)
                                                _ = v16
                                            }
                                            _ = v15
                                        } else {
                                            switch ("\n\t" + "foo bar") {
                                            case "a":
                                                v14 = -1
                                            }
                                            if (true && false) {
                                                if true {
                                                    v14 -= 1
                                                }
                                                if (true && true) {
                                                    if (false || false) {
                                                        switch ("a" + "a") {
                                                        case "a":
                                                            fmt.Println(v5, fi7, fi13)
                                                        default:
                                                            v1[v5] = fi7
                                                        }
                                                    } else {
                                                        var v17 uint32 = v5
                                                        v18 := GenS2{F0: v5}
                                                        var v19 int16 = -299
                                                        _ = v17
                                                        _ = v18
                                                        _ = v19
                                                    }
                                                    var v20 uint = (v4 & (v4 & v4))
                                                    switch fi7 {
                                                    case -1:
                                                        var v21 bool = false
                                                        v22 := GenS2{F0: (v5 + (v5 / ((517) | 1)))}
                                                        _ = v21
                                                        _ = v22
                                                    case 1:
                                                        switch fi13 {
                                                        case 105:
                                                            v14 = 168
                                                            fmt.Println(v4, v20, v5)
                                                        case 32768:
                                                            v23 := GenS1{F0: -801, F1: (fi13 | -449), F2: 0}
                                                            _ = v23
                                                        }
                                                    case 42:
                                                        switch fi13 {
                                                        case 128:
                                                            v24 := GenS3{F0: 0}
                                                            _ = v24
                                                        default:
                                                            var v25 uint8 = 253
                                                            _ = v25
                                                        }
                                                    default:
                                                        var v26 uint8 = 127
                                                        _ = v26
                                                    }
                                                    switch "foo bar" {
                                                    case "a":
                                                        if false {
                                                            v3 = float32(1.5)
                                                            if ((((("\n\t" + ("a" + "foo bar")) + ("foo bar" + ((("" + ("\n\t" + (("hello" + "foo bar") + "\n\t"))) + "a") + (("\n\t" + ("a" + ("a" + ("foo bar" + "foo bar")))) + "")))) + "") + "a") == "hello") {
                                                                var v27 uint8 = 224
                                                                if false {
                                                                    switch (^(fi13)) {
                                                                    case 128:
                                                                        var v28 string = "foo bar"
                                                                        _ = v28
                                                                    case -953:
                                                                        if ((167 > 0) && ("foo bar" > "hello")) {
                                                                            v27 = (^(v27))
                                                                            fmt.Println(v5, v20, v5)
                                                                        }
                                                                    case 42:
                                                                        v20 = 2
                                                                    }
                                                                    var v29 int = v14
                                                                    if true {
                                                                        v20 = v20
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            fmt.Println(v14)
                                                                        case "":
                                                                            v30 := GenS4{F0: 4294967295, F1: 0}
                                                                            _ = v30
                                                                        default:
                                                                            var v31 uint = v4
                                                                            _ = v31
                                                                        }
                                                                        v32 := GenS3{F0: 42}
                                                                        _ = v32
                                                                    } else {
                                                                        _ = len(v6)
                                                                    }
                                                                    _ = v29
                                                                }
                                                                fmt.Println(fi13)
                                                                _ = v27
                                                            }
                                                        }
                                                    case "":
                                                        v33 := GenS2{F0: (v5 & 697)}
                                                        _ = v33
                                                    default:
                                                        fmt.Println(v14, v3)
                                                    }
                                                    _ = v20
                                                } else {
                                                    if true {
                                                        if len(v6) > 0 {
                                                            v6[0] = 394
                                                        }
                                                        fmt.Println(v4, v4, v3)
                                                        var v34 uint16 = 835
                                                        v4 += 1
                                                        _ = v34
                                                    }
                                                    switch ("\n\t" + ("" + "")) {
                                                    case "foo bar":
                                                        fmt.Println(v3)
                                                    case "a":
                                                        v2[fi13] = false
                                                    case "\n\t":
                                                        v3 += v3
                                                    }
                                                    var v35 int16 = 32767
                                                    _ = v35
                                                }
                                                v36 := v2[fi7]
                                                _ = v36
                                            }
                                        }
                                        _ = fi13
                                        _ = v14
                                    }
                                case -425:
                                    var v37 int16 = -32768
                                    _ = v37
                                case -2147483648:
                                    for fi38 := 0; fi38 < 15; fi38++ {
                                        var v39 uint16 = 1
                                        v40 := GenS5{F0: 1, F1: v4}
                                        var v41 string = (("\n\t" + "\n\t") + (("\n\t" + ((("a" + ("\n\t" + "hello")) + "\n\t") + ("hello" + (("a" + "\n\t") + "")))) + "foo bar"))
                                        _ = fi38
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                default:
                                    var v42 float32 = v3
                                    _ = v42
                                }
                            }
                        }
                    }
                    delete(v1, 73)
                    v4 = (v4 - v4)
                    _ = fi7
                }
            case 84:
                switch 128 {
                case -690:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case -2147483648:
                    if len(v6) > 0 {
                        v6[0] = 32768
                    }
                case 653:
                    for fi43 := 0; fi43 < 8; fi43++ {
                        v44 := GenS3{F0: 201}
                        v45 := GenS5{F0: 255, F1: (v4 >> ((373) & 63))}
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                default:
                    var v46 bool = true
                    _ = v46
                }
            default:
                var v47 uint32 = (v5 + (v5 ^ v5))
                _ = v47
            }
            switch ("" + ("hello" + "foo bar")) {
            case "foo bar":
                v48 := GenS6{F0: 4294967295}
                _ = v48
            }
        }
    }
    if ("" > "\n\t") {
        for fi49 := 0; fi49 < 4; fi49++ {
            if len(v6) > 0 {
                v6[0] = fi49
            }
            v3 = 1e10
            var v50 int64 = -1
            _ = fi49
            _ = v50
        }
    }
    if len(v6) > 0 {
        v6[0] = -511
    }
    v1[(v5 << ((v4) & 63))] = 127
    for fi51 := 0; fi51 < 15; fi51++ {
        v52 := GenS7{F0: "a"}
        fmt.Println(v4, v3)
        _ = fi51
        _ = v52
    }
    if true {
        v3 -= (v3 / 0.1)
        for fi53 := 0; fi53 < 6; fi53++ {
            var v54 uint = 18446744073709551615
            _ = fi53
            _ = v54
        }
        for fi55 := 0; fi55 < 14; fi55++ {
            v5 += v5
            var v56 int64 = 42
            _ = fi55
            _ = v56
        }
    } else {
        v5 += v5
        if true {
            v6 = append(v6, -9223372036854775808)
        }
    }
    if (238 == 42) {
        var v57 float32 = (v3 + (v3 - v3))
        _ = v57
    } else {
        v4 = 255
        v58 := GenS1{F0: 9223372036854775807, F1: 32768, F2: -9223372036854775808}
        var v59 float32 = v3
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 16; fi60++ {
        v61 := GenS2{F0: (v5 * v5)}
        _ = fi60
        _ = v61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
