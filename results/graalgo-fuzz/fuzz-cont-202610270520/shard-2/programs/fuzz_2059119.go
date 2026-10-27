package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint; F2 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 int16; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int16; F2 int32 }
type GenS8 struct { F0 uint; F1 string; F2 int8 }
func genFn0(p0 uint32) string {
    v5 := []int{-468, -498, 201}
    v6 := GenS1{F0: 42}
    _ = p0
    _ = v5
    _ = v6
    return ("a" + (("hello" + ("" + "\n\t")) + ((("" + ("" + ("hello" + "hello"))) + "foo bar") + ("foo bar" + ""))))
}
func genFn1() uint32 {
    v7 := []int{-455}
    var v8 int8 = 100
    fmt.Println(v8)
    var v9 int64 = 42
    if (!false) {
        v10 := GenS1{F0: 255}
        var v11 float64 = 3.14159
        _ = v10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 294
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []int{216, 2147483647, 4294967295, 83}
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    if false {
        var v12 uint32 = (v1 ^ (v1 - 1))
        _ = v12
    }
    fmt.Println(v1)
    v13 := GenS1{F0: 0}
    var v14 uint8 = 0
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            if true {
                if ("foo bar" < "hello") {
                    if (32767 >= 2) {
                        v14 += v14
                    }
                    var v17 float64 = (-2.25 - 1e10)
                    v14 = v14
                    v14 += (v14 & v14)
                    if true {
                        if false {
                            if true {
                                v14 = v14
                            }
                        } else {
                            switch "\n\t" {
                            case "":
                                _ = len(v4)
                            default:
                                v18 := GenS1{F0: 255}
                                _ = v18
                            }
                            if (((fi15 % (((fi15 + int(v1))) | 1)) < -227) && false) {
                                v14 += (v14 | v14)
                                var v19 float32 = 1.5
                                var v20 uint64 = 286
                                _ = v19
                                _ = v20
                            } else {
                                v17 += ((0.0 / v17) * 1e10)
                                v21 := GenS1{F0: 1}
                                v22 := GenS2{F0: -30, F1: 2}
                                _ = v21
                                _ = v22
                            }
                            v17 += (v17 * v17)
                        }
                    } else {
                        var v23 uint64 = 416
                        _ = v23
                    }
                    _ = v17
                } else {
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                    fmt.Println(v1, fi15, fi15)
                    v1 = (v1 >> ((42) & 63))
                    if (("\n\t" + ("\n\t" + "foo bar")) >= "hello") {
                        switch ("a" + (("" + ((("hello" + "\n\t") + "foo bar") + "foo bar")) + "foo bar")) {
                        case "":
                            fmt.Println(fi15, v14)
                        case "a":
                            switch (fi15 & fi16) {
                            case 32768:
                                v24 := GenS2{F0: -261, F1: 72}
                                _ = v24
                            case -803:
                                if ((!((1e10 + 0.0) >= 0.0)) || ((3.14159 < 1e10) && true)) {
                                    if (-32768 > 1000) {
                                        v3[(^(fi15))] = v1
                                        var v25 int64 = -1
                                        v26 := GenS3{F0: 0, F1: 4294967295, F2: 489}
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v27 := GenS4{F0: 515}
                                        _ = v27
                                    }
                                } else {
                                    var v28 uint16 = 42
                                    if (((false && (968 < 0)) || false) && true) {
                                        var v29 uint8 = (v14 | v14)
                                        var v30 uint16 = uint16(v1)
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v28 -= (^(v28))
                                        v31 := GenS5{F0: (fi15 / (((fi16 + fi16)) | 1)), F1: 0, F2: v1}
                                        fmt.Println(fi15, v14)
                                        _ = v31
                                    }
                                    var v32 int64 = -9223372036854775808
                                    _ = v28
                                    _ = v32
                                }
                            case 1:
                                var v33 uint32 = v1
                                _ = v33
                            default:
                                v14 += v14
                            }
                            var v34 uint16 = 1
                            _ = v34
                        default:
                            if (true && (!false)) {
                                if false {
                                    v35 := GenS4{F0: 179}
                                    var v36 int8 = -117
                                    var v37 float32 = (((1e10 * 1.5) * 3.14159) / (float32(1.5) / (1.5 + -2.25)))
                                    if (v1 > (v1 << ((1) & 63))) {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        fmt.Println(fi15, v37)
                                        switch "a" {
                                        case "foo bar":
                                            v38 := GenS6{F0: (true && false), F1: 2}
                                            var v39 string = "foo bar"
                                            _ = v38
                                            _ = v39
                                        default:
                                            var v40 int16 = 679
                                            _ = v40
                                        }
                                        if false {
                                            if false {
                                                var v41 uint = 437
                                                delete(v4, fi15)
                                                v37 += v37
                                                _ = v41
                                            }
                                        } else {
                                            v42 := GenS2{F0: -2147483648, F1: 4294967295}
                                            _ = v42
                                        }
                                    }
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                } else {
                                    v43 := GenS7{F0: genFn1(), F1: 1, F2: 278}
                                    v14 += (v14 * v14)
                                    _ = v43
                                }
                                fmt.Println(fi16, v14)
                                v14 += v14
                            } else {
                                v1 -= genFn1()
                            }
                        }
                    } else {
                        v14 += v14
                        v1 += 65535
                    }
                    v14 = (v14 << ((312) & 63))
                }
                var v44 uint16 = uint16(v14)
                fmt.Println(v44)
                _ = v44
            } else {
                var v45 int8 = 1
                var v46 bool = (!(!false))
                _ = v45
                _ = v46
            }
            _ = fi16
        }
        _ = len(v4)
        v47 := GenS8{F0: 255, F1: "", F2: 76}
        _ = fi15
        _ = v47
    }
    var v48 uint8 = (v14 * 28)
    for fi49 := 0; fi49 < 15; fi49++ {
        var v50 uint64 = 2
        _ = fi49
        _ = v50
    }
    v2 = append(v2, -9223372036854775808)
    var v51 int16 = 275
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v51
}
