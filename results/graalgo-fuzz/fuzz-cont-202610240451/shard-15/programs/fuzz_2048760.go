package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
type GenS3 struct { F0 uint; F1 int16; F2 string }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint16; F1 int32; F2 uint32 }

func main() {
    var v1 uint = 12
    var v2 bool = (false && (!(178 >= 65535)))
    v3 := []string{"hello"}
    v4 := []uint32{}
    v5 := make(map[int]string)
    v6 := GenS1{F0: 2147483648, F1: 2, F2: 658}
    var v7 bool = true
    v8 := GenS2{F0: -854, F1: -2147483648, F2: 42}
    v9 := GenS3{F0: 615, F1: 767, F2: ("hello" + "")}
    switch 0 {
    case -971:
        var v10 uint64 = 0
        _ = v10
    case -2147483648:
        v1 -= v1
    case 151:
        if (!v2) {
            v11 := GenS2{F0: -1, F1: -1, F2: 0}
            delete(v5, 42)
            _ = v11
        }
    }
    var v12 float32 = (((3.14159 * 0.0) + (((float32(1e10) / 3.14159) - (1.5 * ((0.0 * (0.1 - 3.14159)) * (0.1 * 0.0)))) * (0.1 * (((float32(1.5) / 3.14159) / 1.0) / -2.25)))) / -2.25)
    if (1 <= -1) {
        _ = len(v5)
        var v13 uint16 = 2
        for fi14 := 0; fi14 < 6; fi14++ {
            fmt.Println(v2, v7, v12)
            var v15 uint16 = 932
            for fi16 := 0; fi16 < 8; fi16++ {
                v2 = (v1 >= (v1 | (^(v1))))
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    var v17 int64 = 0
    if ((!v7) || (42 < v17)) {
        v4 = append(v4, 65535)
        for fi18 := 0; fi18 < 9; fi18++ {
            _ = len(v3)
            _ = fi18
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            if (0 != 32767) {
                v7 = (v7 || false)
                for fi20 := 0; fi20 < 12; fi20++ {
                    v21 := v5[(fi20 - 42)]
                    if (-759 > (v17 / ((v17) | 1))) {
                        v12 -= ((3.14159 * v12) + v12)
                        var v22 uint8 = 255
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v22
                    } else {
                        var v23 string = (("" + "\n\t") + v21)
                        _ = v23
                    }
                    _ = fi20
                    _ = v21
                }
            } else {
                var v24 int8 = 127
                if v7 {
                    fmt.Println(fi19)
                    fmt.Println(v2)
                    fmt.Println(fi19, v24)
                } else {
                    v1 += 413
                    if (!(!((!(!true)) && v2))) {
                        switch "hello" {
                        case "hello":
                            switch fi19 {
                            case 4294967295:
                                v17 -= v17
                            case 32767:
                                var v25 float32 = 1e10
                                _ = v25
                            case -9223372036854775808:
                                v7 = v7
                            }
                        case "\n\t":
                            switch fi19 {
                            case 32767:
                                v17 -= -9223372036854775808
                            }
                        case "foo bar":
                            switch (fi19 / ((fi19) | 1)) {
                            case 338:
                                fmt.Println(v7)
                            }
                        default:
                            v26 := GenS4{F0: -401, F1: false}
                            _ = v26
                        }
                        if v2 {
                            v1 -= v1
                            for fi27 := 0; fi27 < 18; fi27++ {
                                var v28 float32 = 1e10
                                v1 = (^(v1))
                                _ = fi27
                                _ = v28
                            }
                        }
                        fmt.Println(v12, v17, v12)
                    }
                    var v29 uint16 = 32767
                    _ = v29
                }
                _ = v24
            }
            for fi30 := 0; fi30 < 11; fi30++ {
                v1 = v1
                v1 += (^(v1))
                _ = fi30
            }
            for fi31 := 0; fi31 < 18; fi31++ {
                _ = len(v5)
                switch ("" + "a") {
                case "foo bar":
                    v1 -= v1
                default:
                    switch 32767 {
                    case 166:
                        if v7 {
                            v17 = (v17 * (v17 | (v17 + (^(v17)))))
                            var v32 uint64 = 1
                            var v33 uint16 = 42
                            _ = v32
                            _ = v33
                        }
                    case 42:
                        v34 := GenS1{F0: 0, F1: 116, F2: 1}
                        _ = v34
                    case 1:
                        if (!v2) {
                            v35 := GenS5{F0: 176, F1: -2147483648, F2: 2}
                            var v36 uint8 = uint8(v1)
                            var v37 int16 = 981
                            if (v36 < v36) {
                                var v38 float32 = 1.5
                                var v39 float64 = (float64(1e10) * -2.25)
                                var v40 float32 = float32(0.0)
                                _ = len(v5)
                                v2 = v7
                                _ = v38
                                _ = v39
                                _ = v40
                            } else {
                                v2 = ((0.0 + (v12 - float32(-2.25))) < (3.14159 * v12))
                                v41 := GenS2{F0: 42, F1: -1, F2: (v37 >> ((523) & 63))}
                                var v42 uint = v1
                                _ = v41
                                _ = v42
                            }
                            _ = v35
                            _ = v36
                            _ = v37
                        } else {
                            v2 = (!false)
                        }
                    }
                    var v43 int32 = -1
                    fmt.Println(v1)
                    _ = v43
                }
                if ("a" < "foo bar") {
                    var v44 float64 = (-2.25 + 3.14159)
                    var v45 uint = 65535
                    v46 := GenS4{F0: 42, F1: v7}
                    _ = v44
                    _ = v45
                    _ = v46
                }
                v5[fi19] = "hello"
                v1 -= 0
                v17 = (^(v17))
                _ = len(v5)
                _ = fi31
            }
            _ = fi19
        }
    }
    v47 := GenS4{F0: -557, F1: ((false || (!(-2147483648 == 40))) || false)}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
