package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int64; F2 int16 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int32; F2 int8 }

func main() {
    var v1 int64 = -1
    v2 := []string{("hello" + "")}
    var v3 uint = 0
    v4 := make(map[int]uint32)
    v5 := []uint32{2, uint32(v3), 2147483647, 0}
    var v6 int = 42
    var v7 string = "hello"
    var v8 int16 = -927
    switch (v6 & (^(v6))) {
    case 1:
        var v9 int32 = 42
        _ = v9
    case -1:
        if false {
            if (18446744073709551615 != 4294967295) {
                switch (("a" + v7) + ((v7 + v7) + v7)) {
                case "\n\t":
                    if (!false) {
                        var v10 int32 = 1
                        var v11 int = 0
                        _ = v10
                        _ = v11
                    }
                case "a":
                    var v12 float64 = ((float64(1e10) - 1e10) / (0.1 * 1e10))
                    _ = v12
                }
                var v13 string = (v7 + v7)
                for fi14 := 0; fi14 < 4; fi14++ {
                    v15 := GenS1{F0: v3}
                    var v16 uint16 = 307
                    if len(v5) > 0 {
                        v5[0] = 345
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v13
            }
            fmt.Println(v6, v7)
            if false {
                if ((^(v3)) != (^(v3))) {
                    _ = len(v5)
                    var v17 uint = (v3 / ((v3) | 1))
                    _ = v17
                } else {
                    var v18 float32 = 1.5
                    v19 := GenS1{F0: 0}
                    fmt.Println(v3, v18, v8)
                    _ = v18
                    _ = v19
                }
                if (!false) {
                    switch v6 {
                    case 9223372036854775807:
                        for fi20 := 0; fi20 < 14; fi20++ {
                            for fi21 := 0; fi21 < 7; fi21++ {
                                v6 = (v6 & v6)
                                v22 := GenS2{F0: (691 > 32767), F1: v1, F2: -223}
                                v6 = (v6 | (fi21 & fi20))
                                var v23 int64 = v1
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            var v24 int16 = (v8 / ((v8) | 1))
                            if ((275 == 32767) || ((^(v3)) < (v3 >> (((v3 % ((65535) | 1))) & 63)))) {
                                for fi25 := 0; fi25 < 7; fi25++ {
                                    _ = len(v2)
                                    if (127 >= 42) {
                                        v8 = v24
                                    } else {
                                        _ = len(v5)
                                        if ((!(2 != 42)) && ((!true) && (!(!(18446744073709551615 != 17))))) {
                                            var v26 int64 = 42
                                            if (v26 >= (v26 * (v26 & (v26 + -1)))) {
                                                var v27 int8 = 16
                                                fmt.Println(v27, v1)
                                                _ = v27
                                            }
                                            _ = v26
                                        } else {
                                            v28 := GenS1{F0: v3}
                                            _ = v28
                                        }
                                        v29 := GenS3{F0: v3, F1: 42}
                                        _ = v29
                                    }
                                    v8 -= v8
                                    _ = fi25
                                }
                                for fi30 := 0; fi30 < 20; fi30++ {
                                    switch fi30 {
                                    case 127:
                                        _ = len(v2)
                                    case -2147483648:
                                        v31 := GenS4{F0: 100}
                                        _ = v31
                                    default:
                                        var v32 int32 = 0
                                        _ = v32
                                    }
                                    var v33 float32 = 1e10
                                    var v34 uint8 = 157
                                    _ = fi30
                                    _ = v33
                                    _ = v34
                                }
                            } else {
                                v1 -= 773
                                v3 += v3
                                fmt.Println(v7, v1)
                            }
                            _ = fi20
                            _ = v24
                        }
                    case 1:
                        v6 -= 4294967295
                    case 0:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
            } else {
                var v35 uint8 = 255
                fmt.Println(v7, v3, v7)
                v3 = (v3 << ((v3) & 63))
                _ = v35
            }
        }
    default:
        v8 = (v8 - -32768)
    }
    v4[(-(v6))] = 2147483648
    v36 := GenS5{F0: false}
    v37 := GenS5{F0: false}
    v38 := GenS1{F0: 0}
    v39 := GenS1{F0: v3}
    delete(v4, (v6 >> (((v3 - v3)) & 63)))
    if (!false) {
        v40 := v4[v6]
        _ = v40
    }
    var v41 int64 = -724
    var v42 int32 = 0
    switch ("hello" + (v7 + v7)) {
    case "hello":
        if true {
            v3 = v3
            v42 -= v42
        }
    }
    var v43 int64 = v1
    for fi44 := 0; fi44 < 17; fi44++ {
        var v45 string = "foo bar"
        for fi46 := 0; fi46 < 6; fi46++ {
            var v47 int8 = 1
            fmt.Println(v41, v43, v43)
            if (v7 == "a") {
                fmt.Println(v7)
                switch fi44 {
                case -278:
                    v42 += v42
                case 421:
                    fmt.Println(v8, v45, v43)
                    var v48 float64 = -2.25
                    fmt.Println(v8, v3)
                    _ = v48
                case 2147483647:
                    v49 := GenS6{F0: v8, F1: v42, F2: v47}
                    _ = v49
                default:
                    if ((v41 + 9223372036854775807) <= v41) {
                        _ = len(v4)
                    }
                }
                fmt.Println(v6, v41, v8)
            }
            _ = fi46
            _ = v47
        }
        _ = fi44
        _ = v45
    }
    for fi50 := 0; fi50 < 10; fi50++ {
        var v51 uint64 = uint64(v3)
        v52 := GenS6{F0: (^(v8)), F1: (v42 + (v42 << ((18446744073709551615) & 63))), F2: 1}
        fmt.Println(v1, v41, v42)
        _ = fi50
        _ = v51
        _ = v52
    }
    if true {
        v3 = (^(v3))
        if true {
            var v53 uint32 = 2147483648
            _ = v53
        }
        var v54 uint16 = 247
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
