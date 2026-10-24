package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }

func main() {
    var v1 int64 = -1
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    v4 := []int{128}
    v5 := []int{-1, 0}
    var v6 int16 = 32767
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(fi7, fi7, v1)
        for fi8 := 0; fi8 < 18; fi8++ {
            v6 = (v6 * v6)
            switch "hello" {
            case "hello":
                v9 := GenS1{F0: -1, F1: 128, F2: fi7}
                _ = v9
            case "":
                var v10 float32 = -2.25
                _ = v10
            }
            if len(v4) > 0 {
                v4[0] = 42
            }
            _ = fi8
        }
        _ = fi7
    }
    if len(v5) > 0 {
        v5[0] = -141
    }
    var v11 uint = 65535
    var v12 float32 = 1e10
    if true {
        for fi13 := 0; fi13 < 8; fi13++ {
            if false {
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(v12, v12, v1)
                    _ = fi14
                }
                var v15 uint8 = 42
                _ = v15
            }
            if len(v4) > 0 {
                v4[0] = (^(fi13))
            }
            _ = fi13
        }
    }
    switch 127 {
    case 0:
        var v16 uint32 = 2147483647
        _ = v16
    case -1:
        for fi17 := 0; fi17 < 8; fi17++ {
            v3[1] = true
            if (!((fi17 >> ((v11) & 63)) > (fi17 / ((fi17) | 1)))) {
                for fi18 := 0; fi18 < 1; fi18++ {
                    v11 = v11
                    _ = fi18
                }
                if false {
                    var v19 uint8 = 128
                    _ = v19
                }
                var v20 uint32 = uint32(v11)
                _ = v20
            }
            v21 := GenS1{F0: -20, F1: 42, F2: fi17}
            _ = fi17
            _ = v21
        }
    case 128:
        v22 := GenS1{F0: -128, F1: 42, F2: 128}
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := GenS1{F0: 1, F1: 255, F2: 0}
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 float64 = 3.14159
            if true {
                v12 += v12
            } else {
                _ = len(v5)
                var v27 int64 = (v1 + v1)
                var v28 int8 = 127
                fmt.Println(v11, fi25)
                switch ("\n\t" + (("a" + "foo bar") + "a")) {
                case "a":
                    v29 := GenS1{F0: v28, F1: 128, F2: -428}
                    _ = v29
                case "":
                    v30 := GenS1{F0: v28, F1: 220, F2: (fi23 - 42)}
                    _ = v30
                case "hello":
                    fmt.Println(v26)
                }
                _ = v27
                _ = v28
            }
            if false {
                if true {
                    if len(v5) > 0 {
                        v5[0] = (fi25 ^ (fi23 % ((int(v6)) | 1)))
                    }
                } else {
                    if (255 < 65535) {
                        if false {
                            fmt.Println(fi25, fi25, v6)
                            fmt.Println(v1, v1, fi25)
                        } else {
                            var v31 int16 = v6
                            _ = v31
                        }
                        v12 += float32(v26)
                        switch (fi23 % ((-9223372036854775808) | 1)) {
                        case 2147483647:
                            if (-1 == 54) {
                                if true {
                                    v3[2147483648] = true
                                    v6 = v6
                                    v32 := GenS2{F0: 32767, F1: 0, F2: 2147483647}
                                    switch int(v6) {
                                    case 2147483647:
                                        delete(v3, 65535)
                                    }
                                    if false {
                                        if (v12 < (float32(v26) * v12)) {
                                            var v33 uint8 = 1
                                            v34 := GenS1{F0: int8(v6), F1: (v33 | (v33 & (v33 >> ((v11) & 63)))), F2: fi25}
                                            _ = v33
                                            _ = v34
                                        }
                                        v12 -= (v12 + v12)
                                    } else {
                                        if false {
                                            v6 = v6
                                        } else {
                                            var v35 float64 = 0.1
                                            _ = v35
                                        }
                                        var v36 uint = 65535
                                        if true {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        } else {
                                            if true {
                                                switch "foo bar" {
                                                case "":
                                                    v37 := GenS1{F0: 1, F1: 128, F2: 93}
                                                    _ = v37
                                                case "foo bar":
                                                    var v38 int16 = 42
                                                    _ = v38
                                                default:
                                                    if false {
                                                        v39 := GenS1{F0: 100, F1: 0, F2: fi23}
                                                        var v40 int64 = 1
                                                        v36 = v11
                                                        v26 = 1e10
                                                        _ = v39
                                                        _ = v40
                                                    } else {
                                                        if true {
                                                            fmt.Println(fi23)
                                                            v3[744] = ((!false) && (v1 <= 0))
                                                        }
                                                        var v41 float64 = ((((v26 / (v26 + (float64(0.0) + v26))) - float64(((v12 - (v12 + (((v12 * (v12 - ((((v12 * v12) / ((0.0 - v12) / float32(0.1))) / (((3.14159 * v12) / (v12 * 1.5)) - v12)) - v12))) - (((1.5 - v12) * v12) + -2.25)) * (1e10 * v12)))) / 3.14159))) / v26) + -2.25)
                                                        var v42 uint16 = 166
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                }
                                                if (!false) {
                                                    v43 := GenS2{F0: 447, F1: 2, F2: -732}
                                                    _ = v43
                                                } else {
                                                    var v44 uint32 = 65535
                                                    _ = v44
                                                }
                                            }
                                        }
                                        v26 = (-2.25 + v26)
                                        switch fi25 {
                                        case 0:
                                            fmt.Println(fi25)
                                        case 111:
                                            v11 += (v11 + v36)
                                        }
                                        v26 = ((v26 + v26) * (v26 / ((v26 - float64(v12)) * (3.14159 * v26))))
                                        v12 = (float32(-2.25) / v12)
                                        _ = v36
                                    }
                                    _ = v32
                                } else {
                                    var v45 int = fi23
                                    _ = len(v2)
                                    var v46 uint16 = 42
                                    _ = v45
                                    _ = v46
                                }
                            }
                        case 127:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            var v47 float64 = (v26 * float64(1.5))
                            _ = v47
                        }
                    }
                    delete(v3, 230)
                    v1 = v1
                    v12 += (((3.14159 * v12) - v12) - v12)
                    var v48 uint16 = 65535
                    var v49 int16 = v6
                    _ = v48
                    _ = v49
                }
            }
            _ = fi25
            _ = v26
        }
        for fi50 := 0; fi50 < 13; fi50++ {
            var v51 uint16 = 32767
            _ = fi50
            _ = v51
        }
        _ = fi23
        _ = v24
    }
    var v52 uint64 = 1
    for fi53 := 0; fi53 < 14; fi53++ {
        for fi54 := 0; fi54 < 8; fi54++ {
            v1 = v1
            var v55 float32 = -2.25
            _ = fi54
            _ = v55
        }
        v56 := GenS2{F0: 255, F1: 0, F2: 2147483647}
        v1 = v1
        _ = fi53
        _ = v56
    }
    var v57 uint64 = 4294967295
    for fi58 := 0; fi58 < 20; fi58++ {
        fmt.Println(v12)
        _ = fi58
    }
    _ = len(v2)
    v11 = 859
    v11 += (v11 ^ v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v57
    _ = v6
}
