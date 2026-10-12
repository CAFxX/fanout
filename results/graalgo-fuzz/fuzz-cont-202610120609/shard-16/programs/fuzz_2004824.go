package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 string; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS6 struct { F0 uint; F1 int64; F2 uint64 }
type GenS7 struct { F0 int8 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    v3 := []uint32{410, 2147483647, 1}
    var v4 int32 = 0
    v5 := make(map[int]bool)
    var v6 string = "\n\t"
    fmt.Println(v4, v4)
    fmt.Println(v6, v4)
    fmt.Println(v6, v4)
    for fi7 := 0; fi7 < 8; fi7++ {
        if (true || (188 > 866)) {
            v6 = v6
        }
        v8 := GenS1{F0: v6}
        for fi9 := 0; fi9 < 16; fi9++ {
            if ((false || ((1e10 / 1e10) <= 3.14159)) || true) {
                v1 = append(v1, v6)
                v3 = append(v3, 625)
                var v10 int32 = 0
                _ = v10
            } else {
                v4 = (-(v4))
                v4 = 2147483647
                v3 = append(v3, 42)
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    v11 := GenS2{F0: -194, F1: (v4 >> ((4294967295) & 63))}
    _ = len(v2)
    for fi12 := 0; fi12 < 6; fi12++ {
        if false {
            if (-870 >= 9223372036854775807) {
                delete(v2, 65535)
            }
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := GenS2{F0: 1000, F1: 2147483647}
                if len(v3) > 0 {
                    v3[0] = 597
                }
                _ = fi13
                _ = v14
            }
        }
        _ = fi12
    }
    v4 += (v4 / (((v4 / ((v4) | 1))) | 1))
    fmt.Println(v4, v6)
    if false {
        switch (v6 + ((v6 + v6) + (v6 + v6))) {
        case "":
            switch v6 {
            case "\n\t":
                if (false && false) {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 float32 = float32(0.0)
                        v4 += (v4 - (^(v4)))
                        _ = fi15
                        _ = v16
                    }
                    var v17 int32 = (v4 + (-(v4)))
                    _ = v17
                } else {
                    v4 = v4
                    v18 := GenS3{F0: v6, F1: 127, F2: 394}
                    v6 = ((v6 + v6) + ("hello" + v6))
                    _ = v18
                }
            }
        case "hello":
            fmt.Println(v4, v6)
        }
    }
    var v19 float64 = ((0.1 / ((3.14159 + float64(3.14159)) - 0.1)) + (((0.0 + float64(-2.25)) + -2.25) - ((-2.25 - 0.1) + ((3.14159 + 0.1) * (-2.25 + 0.0)))))
    if (0 == 127) {
        v19 = 1e10
    } else {
        for fi20 := 0; fi20 < 18; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                switch v6 {
                case "a":
                    var v22 uint16 = 1
                    _ = v22
                case "\n\t":
                    v23 := GenS4{F0: 19, F1: 178}
                    _ = v23
                case "":
                    var v24 uint32 = 2147483648
                    _ = v24
                default:
                    v2[4294967295] = fi21
                    v25 := GenS5{F0: 4294967295, F1: 74, F2: 285}
                    _ = v25
                }
                _ = fi21
            }
            fmt.Println(v19, v4)
            if (42 < 255) {
                for fi26 := 0; fi26 < 10; fi26++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v19 = v19
                    var v27 int32 = v4
                    _ = fi26
                    _ = v27
                }
                v28 := GenS3{F0: "", F1: 66, F2: 2}
                v29 := GenS5{F0: 1, F1: -1, F2: 32767}
                _ = v28
                _ = v29
            }
            _ = fi20
        }
        var v30 float32 = 1.5
        v31 := GenS4{F0: 42, F1: 42}
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        for fi33 := 0; fi33 < 6; fi33++ {
            if true {
                var v34 uint = 95
                v35 := GenS1{F0: "foo bar"}
                _ = v34
                _ = v35
            }
            if (0 != 2) {
                v6 = v6
                var v36 uint16 = 65535
                _ = v36
            }
            if (127 >= -1) {
                fmt.Println(v4)
                v6 = "\n\t"
                if (!(v4 >= v4)) {
                    var v37 int64 = 993
                    var v38 uint = 1
                    _ = v37
                    _ = v38
                }
            } else {
                var v39 int16 = 42
                if (42 != 962) {
                    delete(v2, 42)
                    v39 = v39
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v4)
                    var v40 uint16 = 32767
                    if ((!(242 < 42)) && false) {
                        v41 := GenS6{F0: 18446744073709551615, F1: 184, F2: 18446744073709551615}
                        var v42 bool = ("foo bar" != "")
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v43 int8 = -94
                        fmt.Println(v42, v39)
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    _ = v40
                }
                fmt.Println(v19)
                if (false && (3.14159 < 1.5)) {
                    var v44 float64 = 0.1
                    _ = v44
                } else {
                    var v45 uint8 = 127
                    if true {
                        fmt.Println(fi32)
                        var v46 float64 = (float64(0.0) + 0.1)
                        _ = v46
                    } else {
                        var v47 uint8 = 195
                        if (!true) {
                            var v48 uint16 = 255
                            _ = v48
                        } else {
                            delete(v2, 42)
                            var v49 int64 = 9223372036854775807
                            _ = len(v3)
                            _ = v49
                        }
                        var v50 uint8 = 127
                        _ = v47
                        _ = v50
                    }
                    _ = v45
                }
                var v51 int16 = (v39 - -1)
                _ = v39
                _ = v51
            }
            _ = fi33
        }
        _ = fi32
    }
    for fi52 := 0; fi52 < 12; fi52++ {
        v53 := GenS7{F0: -128}
        _ = fi52
        _ = v53
    }
    fmt.Println(v19, v4, v4)
    v4 -= v4
    var v54 int32 = 42
    fmt.Println(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
