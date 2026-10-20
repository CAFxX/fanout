package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 int64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 uint8; F1 int; F2 int }
func genFn0() uint16 {
    var v5 uint32 = 982
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 bool = (-32768 <= 253)
        v5 = v5
        if v7 {
            fmt.Println(fi6, v5)
        }
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 42}
    v5 += (v5 & 1)
    _ = v5
    _ = v8
    return 42
}
func genFn1() int {
    var v9 int8 = -128
    if ((float32(1e10) + 3.14159) != float32((0.1 * ((-2.25 - 0.0) + 0.1)))) {
        v9 -= v9
    }
    _ = v9
    return 4294967295
}

func main() {
    v1 := []string{}
    var v2 uint = 718
    var v3 uint32 = 2
    v4 := make(map[uint32]bool)
    if (42 >= -661) {
        if ((v3 >= (^(v3))) || (384 >= 32767)) {
            var v10 int = genFn1()
            _ = v10
        }
        if false {
            fmt.Println(v3, v3)
            var v11 uint64 = 1
            switch "hello" {
            case "a":
                if (!(genFn1() > 246)) {
                    if false {
                        v2 += v2
                        v11 = 370
                        v12 := v4[(v3 / ((652) | 1))]
                        _ = v12
                    } else {
                        var v13 bool = false
                        _ = v13
                    }
                    var v14 float64 = -2.25
                    _ = v14
                }
            default:
                v4[283] = (int32(v11) != -2147483648)
            }
            _ = v11
        } else {
            var v15 bool = true
            for fi16 := 0; fi16 < 3; fi16++ {
                v2 = v2
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v17 := GenS2{F0: 2, F1: 42, F2: 1}
                _ = fi16
                _ = v17
            }
            v2 += 255
            _ = v15
        }
    } else {
        v2 = v2
        var v18 uint8 = 255
        v19 := GenS3{F0: ("foo bar" + "\n\t"), F1: (^(v3))}
        _ = v18
        _ = v19
    }
    var v20 float64 = (3.14159 * 0.0)
    v21 := GenS2{F0: 2, F1: 9223372036854775807, F2: -9223372036854775808}
    switch 128 {
    case 626:
        if ((32768 == genFn1()) && (true && (113 != 42))) {
            for fi22 := 0; fi22 < 18; fi22++ {
                v3 = uint32(v2)
                v2 += v2
                if (!("hello" < "foo bar")) {
                    var v23 uint8 = uint8(v2)
                    _ = v23
                } else {
                    v2 -= 92
                    var v24 uint16 = 1
                    _ = v24
                }
                _ = fi22
            }
        }
    default:
        for fi25 := 0; fi25 < 12; fi25++ {
            if false {
                if (genFn0() > 1) {
                    var v26 float32 = float32(((v20 / v20) * v20))
                    v2 -= v2
                    v27 := GenS1{F0: (v2 ^ (v2 >> ((v2) & 63)))}
                    _ = v26
                    _ = v27
                }
            }
            var v28 uint16 = 255
            _ = fi25
            _ = v28
        }
    }
    var v29 int = 2147483647
    for fi30 := 0; fi30 < 4; fi30++ {
        if true {
            for fi31 := 0; fi31 < 6; fi31++ {
                v2 = (v2 - v2)
                var v32 uint64 = 1
                v29 = (v29 << ((v2) & 63))
                _ = fi31
                _ = v32
            }
        }
        var v33 float32 = 1e10
        if (-809 == -292) {
            v34 := GenS4{F0: (^(v2)), F1: (v3 * 0)}
            if (-2.25 != 0.1) {
                delete(v4, v3)
                for fi35 := 0; fi35 < 15; fi35++ {
                    if ((true && false) && true) {
                        var v36 uint16 = 629
                        if (-1 != 42) {
                            v37 := GenS5{F0: -9223372036854775808}
                            var v38 uint16 = 255
                            _ = v37
                            _ = v38
                        } else {
                            if true {
                                fmt.Println(v29)
                                v39 := GenS6{F0: genFn1()}
                                switch (fi35 >> ((42) & 63)) {
                                case 0:
                                    v36 += v36
                                case 1:
                                    if true {
                                        var v40 uint32 = 4294967295
                                        if ((0.1 > float64(v33)) || true) {
                                            v2 += v2
                                        }
                                        v33 = v33
                                        _ = v40
                                    } else {
                                        v3 += v3
                                        v41 := GenS2{F0: (^(v36)), F1: -117, F2: 0}
                                        v4[v3] = false
                                        switch "\n\t" {
                                        case "a":
                                            var v42 float32 = v33
                                            _ = v42
                                        }
                                        v33 = v33
                                        v43 := GenS7{F0: 1}
                                        if (((("a" + "a") > "hello") || false) || false) {
                                            if (!false) {
                                                fmt.Println(v33)
                                                var v44 float64 = 0.0
                                                _ = len(v1)
                                                v45 := GenS8{F0: -342, F1: v36}
                                                _ = v44
                                                _ = v45
                                            }
                                        } else {
                                            switch 820 {
                                            case -167:
                                                v33 = v33
                                            case 416:
                                                v2 = (v2 + v2)
                                            case 127:
                                                var v46 int16 = -32768
                                                _ = v46
                                            default:
                                                if true {
                                                    var v47 float64 = float64(v33)
                                                    _ = v47
                                                } else {
                                                    var v48 float64 = (v20 + v20)
                                                    _ = v48
                                                }
                                            }
                                        }
                                        _ = v41
                                        _ = v43
                                    }
                                }
                                _ = v39
                            } else {
                                v49 := GenS3{F0: ("hello" + ("\n\t" + "a")), F1: 234}
                                v50 := GenS9{F0: 128, F1: 520, F2: 32768}
                                _ = v49
                                _ = v50
                            }
                            v4[65535] = true
                        }
                        _ = v36
                    }
                    switch "a" {
                    case "\n\t":
                        if false {
                            var v51 int8 = -81
                            _ = v51
                        }
                        v29 = (^(v29))
                        var v52 uint16 = 42
                        _ = v52
                    default:
                        var v53 int32 = 42
                        _ = v53
                    }
                    v54 := GenS6{F0: -2147483648}
                    _ = fi35
                    _ = v54
                }
            } else {
                v33 = v33
                v55 := v4[456]
                if (267 != 1) {
                    for fi56 := 0; fi56 < 5; fi56++ {
                        _ = len(v4)
                        _ = fi56
                    }
                }
                _ = v55
            }
            _ = v34
        }
        _ = fi30
        _ = v33
    }
    if (v3 != 42) {
        v57 := v4[v3]
        var v58 uint16 = uint16(v2)
        _ = v57
        _ = v58
    } else {
        var v59 uint64 = 4294967295
        _ = v59
    }
    v29 -= (v29 * (v29 - v29))
    delete(v4, 65)
    for fi60 := 0; fi60 < 2; fi60++ {
        v20 = v20
        var v61 int8 = 1
        _ = fi60
        _ = v61
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
}
