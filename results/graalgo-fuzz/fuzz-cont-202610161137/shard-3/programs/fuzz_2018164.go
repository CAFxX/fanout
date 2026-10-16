package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 uint8; F1 uint; F2 int32 }
type GenS9 struct { F0 int64; F1 uint64; F2 int }
type GenS10 struct { F0 int64; F1 uint8 }
type GenS11 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -765
    v3 := []int{}
    delete(v1, 461)
    var v4 uint16 = 32767
    v5 := GenS1{F0: 675, F1: 47}
    v2 = (v2 & (v2 - 0))
    var v6 int32 = 1
    fmt.Println(v2, v4, v6)
    if true {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint32 = 2
            _ = fi7
            _ = v8
        }
    }
    if false {
        fmt.Println(v6, v6, v2)
    } else {
        _ = len(v3)
    }
    if false {
        if ((!((0.0 / 1.0) == ((1e10 * 0.1) - 3.14159))) && (!(-128 == 126))) {
            var v9 uint64 = 42
            for fi10 := 0; fi10 < 9; fi10++ {
                fmt.Println(v9, v6, v9)
                _ = fi10
            }
            _ = v9
        }
        v11 := GenS1{F0: uint32(v4), F1: v6}
        fmt.Println(v2, v6)
        _ = v11
    } else {
        v12 := GenS2{F0: -108, F1: 561}
        _ = v12
    }
    v3 = append(v3, 553)
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int = 1
            v6 += (v6 | (v6 | (v6 ^ -566)))
            v16 := GenS3{F0: 2, F1: (v4 % (((v4 * 1)) | 1)), F2: 890}
            v15 = (-(v15))
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 float32 = 1e10
        _ = fi13
        _ = v17
    }
    if false {
        var v18 int16 = 32767
        v19 := GenS4{F0: 144}
        _ = v18
        _ = v19
    }
    if false {
        v6 -= 1
        if (true && false) {
            if false {
                v20 := v1[980]
                v21 := v1[v20]
                v22 := GenS5{F0: uint16(v20), F1: (v4 | 2), F2: v20}
                _ = v20
                _ = v21
                _ = v22
            }
            fmt.Println(v2, v6)
        }
        v23 := GenS6{F0: (v4 & (v4 & 65535)), F1: ""}
        _ = v23
    } else {
        if (((((0.0 * ((0.1 / 1.0) / (1.5 + (3.14159 - 0.0)))) / 0.1) + 0.1) / 1.5) <= 1e10) {
            fmt.Println(v2, v6)
        } else {
            v24 := GenS2{F0: -1, F1: 1000}
            if (true && false) {
                _ = len(v3)
            } else {
                v25 := GenS7{F0: false, F1: 0}
                v2 -= -1
                v3 = append(v3, -9223372036854775808)
                _ = v25
            }
            _ = v24
        }
        _ = len(v3)
        for fi26 := 0; fi26 < 8; fi26++ {
            fmt.Println(v6, fi26)
            switch (fi26 >> ((65535) & 63)) {
            case 4294967295:
                fmt.Println(v4)
            }
            for fi27 := 0; fi27 < 1; fi27++ {
                switch "a" {
                case "foo bar":
                    var v28 uint32 = 0
                    _ = v28
                case "hello":
                    delete(v1, 65535)
                }
                if true {
                    v2 -= 0
                    fmt.Println(v6, fi26)
                } else {
                    if (((1e10 != (1e10 * ((1.5 / -2.25) / -2.25))) && (!true)) || (0 >= uint8(v4))) {
                        switch ("a" + ((("foo bar" + ("" + "")) + "a") + "\n\t")) {
                        case "a":
                            if true {
                                fmt.Println(v4, v2, v6)
                                v29 := GenS8{F0: 24, F1: 255, F2: (v6 >> ((409) & 63))}
                                fmt.Println(fi27, fi26)
                                _ = v29
                            }
                        default:
                            var v30 uint64 = 824
                            v4 += 32767
                            var v31 uint32 = 2
                            _ = v30
                            _ = v31
                        }
                        v32 := GenS9{F0: 989, F1: 2, F2: fi27}
                        _ = v32
                    }
                    v33 := GenS2{F0: 127, F1: 32767}
                    var v34 float32 = (1e10 - float32(1e10))
                    fmt.Println(fi27, fi27, v6)
                    if true {
                        if ((42 >= -67) || (!(-711 > fi26))) {
                            if (v2 != (^(v2))) {
                                v35 := GenS7{F0: (((1.5 * 3.14159) + 1.5) >= (((3.14159 + 1e10) * 1e10) * 0.1)), F1: 0}
                                _ = v35
                            } else {
                                var v36 uint16 = v4
                                delete(v1, 2)
                                _ = v36
                            }
                            v4 = v4
                            v37 := GenS10{F0: v2, F1: 127}
                            _ = v37
                        } else {
                            v38 := GenS9{F0: v2, F1: 42, F2: (fi27 % (((fi27 % (((fi27 / ((fi26) | 1))) | 1))) | 1))}
                            v2 += 0
                            var v39 int = (^(fi26))
                            switch (fi26 | fi26) {
                            case 107:
                                if false {
                                    if true {
                                        var v40 uint64 = 1
                                        _ = v40
                                    }
                                    v41 := GenS4{F0: 42}
                                    _ = v41
                                }
                            case -1:
                                v42 := GenS10{F0: (v2 / ((v2) | 1)), F1: 255}
                                _ = v42
                            case 32768:
                                v43 := GenS11{F0: 834, F1: v6}
                                _ = v43
                            default:
                                var v44 int32 = v6
                                _ = v44
                            }
                            var v45 string = "\n\t"
                            _ = v38
                            _ = v39
                            _ = v45
                        }
                        var v46 uint16 = 42
                        if (true && (1 == 18446744073709551615)) {
                            if (372 == 2147483647) {
                                var v47 float64 = (1.5 / -2.25)
                                v48 := GenS2{F0: 1, F1: 38}
                                _ = v47
                                _ = v48
                            } else {
                                var v49 uint = 47
                                delete(v1, 0)
                                _ = v49
                            }
                            var v50 int = fi27
                            fmt.Println(v6)
                            _ = v50
                        } else {
                            v4 -= 2
                        }
                        _ = v46
                    } else {
                        var v51 string = (("\n\t" + "foo bar") + "")
                        var v52 float32 = (((float32(float64(1e10)) - (1.5 + (v34 - 1.5))) * (v34 - (v34 - 0.1))) / v34)
                        var v53 int32 = (v6 - v6)
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    _ = v33
                    _ = v34
                }
                _ = fi27
            }
            _ = fi26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
