package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS3 struct { F0 int32; F1 string; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int16 }
func genFn0() int64 {
    if false {
        var v4 uint = 42
        v5 := GenS1{F0: 255, F1: 1}
        _ = v4
        _ = v5
    }
    switch -289 {
    case -651:
        v6 := GenS1{F0: 255, F1: 2}
        _ = v6
    }
    var v7 string = ("\n\t" + "\n\t")
    if true {
        var v8 int8 = 42
        if true {
            v8 = v8
            v7 = ""
            for fi9 := 0; fi9 < 20; fi9++ {
                if true {
                    var v10 uint32 = 4294967295
                    _ = v10
                } else {
                    fmt.Println(v7, v7)
                    fmt.Println(fi9, v7, fi9)
                }
                _ = fi9
            }
        } else {
            if (!(true && false)) {
                var v11 uint8 = 42
                _ = v11
            }
            v12 := []int{}
            _ = v12
        }
        _ = v8
    } else {
        fmt.Println(v7)
        fmt.Println(v7)
        fmt.Println(v7)
    }
    var v13 int64 = 9223372036854775807
    _ = v13
    _ = v7
    return -1
}
func genFn1() string {
    v14 := []int{9223372036854775807, 4294967295, 2147483647, 708}
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 12; fi16++ {
            fmt.Println(fi16, fi16)
            fmt.Println(fi16, fi15)
            _ = fi16
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            switch "" {
            case "\n\t":
                if false {
                    if (2 < 717) {
                        var v18 bool = false
                        _ = v18
                    }
                    if ((!true) || true) {
                        v19 := GenS2{F0: 2147483648, F1: 0, F2: -1}
                        var v20 uint32 = 42
                        v21 := GenS3{F0: 1, F1: ("" + ("" + "foo bar")), F2: -32768}
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        var v22 bool = (fi17 < fi15)
                        var v23 uint16 = 255
                        switch ("a" + (("\n\t" + ("hello" + ("\n\t" + (("hello" + (("\n\t" + ("foo bar" + "hello")) + ("foo bar" + (("\n\t" + "a") + (((("a" + "") + "foo bar") + "\n\t") + ("a" + "")))))) + "")))) + "")) {
                        case "foo bar":
                            fmt.Println(v22, v22, fi17)
                        }
                        _ = v22
                        _ = v23
                    }
                }
            }
            var v24 float64 = ((((1.5 * 1.5) * -2.25) / 0.1) - -2.25)
            v24 = v24
            _ = fi17
            _ = v24
        }
        _ = fi15
    }
    var v25 int64 = 0
    _ = v14
    _ = v25
    return ("foo bar" + "hello")
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{696, 42, 394, 65535}
    v3 := make(map[int]int)
    delete(v1, 32768)
    var v26 bool = ((true && (1 >= 37)) || false)
    v26 = false
    if (472 >= 0) {
        v26 = (-815 >= 1)
        v27 := GenS3{F0: -2147483648, F1: "", F2: -1}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 1; fi28++ {
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 int64 = -194
                _ = fi29
                _ = v30
            }
            if (2 <= 15) {
                v3[fi28] = fi28
                var v31 uint8 = 1
                if true {
                    for fi32 := 0; fi32 < 11; fi32++ {
                        v31 -= (v31 / (((v31 * v31)) | 1))
                        _ = fi32
                    }
                }
                _ = v31
            } else {
                var v33 int64 = -379
                if (!v26) {
                    v26 = true
                }
                _ = v33
            }
            _ = fi28
        }
    }
    var v34 float64 = -2.25
    if v26 {
        v35 := GenS4{F0: genFn0()}
        _ = v35
    }
    delete(v1, -403)
    if (255 == 42) {
        if false {
            fmt.Println(v26, v26)
            if v26 {
                v36 := GenS3{F0: 2147483647, F1: "foo bar", F2: -32768}
                v37 := GenS4{F0: 675}
                for fi38 := 0; fi38 < 14; fi38++ {
                    if v26 {
                        var v39 float32 = 0.1
                        _ = v39
                    } else {
                        v40 := GenS3{F0: -41, F1: ("hello" + "hello"), F2: -1}
                        _ = v40
                    }
                    fmt.Println(v34, v26, fi38)
                    for fi41 := 0; fi41 < 6; fi41++ {
                        switch (-(fi41)) {
                        case -407:
                            v42 := GenS3{F0: -1, F1: ("a" + "foo bar"), F2: 1000}
                            _ = v42
                        case 72:
                            var v43 float32 = -2.25
                            _ = v43
                        default:
                            _ = len(v3)
                        }
                        switch (fi41 >> ((281) & 63)) {
                        case 127:
                            if v26 {
                                if v26 {
                                    if v26 {
                                        v34 -= (3.14159 / v34)
                                    } else {
                                        fmt.Println(v26, fi41)
                                        var v44 bool = (113 < 2)
                                        if v26 {
                                            v34 += (v34 / (v34 / (v34 + 0.0)))
                                        } else {
                                            switch 32767 {
                                            case -2147483648:
                                                _ = len(v2)
                                            case 1:
                                                v45 := GenS2{F0: 322, F1: 128, F2: -1}
                                                _ = v45
                                            default:
                                                v26 = v26
                                            }
                                            fmt.Println(fi38)
                                            v44 = (v44 && (156 == 9))
                                        }
                                        _ = v44
                                    }
                                } else {
                                    v26 = (1000 == -322)
                                }
                                v34 -= v34
                            } else {
                                var v46 int32 = 0
                                v47 := GenS2{F0: 580, F1: 127, F2: v46}
                                var v48 uint8 = 96
                                if len(v2) > 0 {
                                    v2[0] = 2147483648
                                }
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                        default:
                            v26 = (65535 != 65535)
                        }
                        _ = fi41
                    }
                    _ = fi38
                }
                _ = v36
                _ = v37
            }
        }
        if v26 {
            for fi49 := 0; fi49 < 18; fi49++ {
                v50 := GenS5{F0: 926, F1: 0}
                _ = fi49
                _ = v50
            }
            if ((v26 && ((v34 / v34) > float64(3.14159))) && true) {
                v34 = (v34 * 3.14159)
                for fi51 := 0; fi51 < 20; fi51++ {
                    v34 = (0.0 - float64(float32(v34)))
                    v26 = (864 < 951)
                    _ = fi51
                }
            } else {
                switch (((("a" + ("foo bar" + "")) + (("hello" + ("foo bar" + ("foo bar" + ("" + "foo bar")))) + "a")) + ("" + ("foo bar" + "hello"))) + "a") {
                case "a":
                    var v52 int32 = 42
                    _ = v52
                case "foo bar":
                    v26 = (3.14159 != float32((1e10 * ((float64((0.0 * 0.1)) * v34) / 1.5))))
                default:
                    var v53 uint8 = 156
                    _ = v53
                }
                var v54 uint64 = 18446744073709551615
                var v55 int = 2147483647
                _ = v54
                _ = v55
            }
            _ = len(v3)
        } else {
            var v56 uint = 255
            v57 := GenS1{F0: 126, F1: 29}
            _ = v56
            _ = v57
        }
        fmt.Println(v34)
    } else {
        _ = len(v1)
        fmt.Println(v34)
        var v58 uint8 = 1
        _ = v58
    }
    v59 := GenS2{F0: 0, F1: 128, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v59
}
