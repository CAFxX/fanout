package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint64 }
type GenS2 struct { F0 string; F1 int8; F2 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 uint64 = 1
    var v2 int32 = -1
    v3 := []int{128, -9223372036854775808}
    var v4 float32 = -2.25
    var v5 int8 = 127
    if (2 < 128) {
        for fi6 := 0; fi6 < 2; fi6++ {
            for fi7 := 0; fi7 < 20; fi7++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if ((v1 & v1) != v1) {
                    fmt.Println(v5, v5)
                    v2 = (v2 >> ((4294967295) & 63))
                    switch fi7 {
                    case 689:
                        _ = len(v3)
                    case 899:
                        var v8 uint = 42
                        _ = v8
                    default:
                        if false {
                            switch "\n\t" {
                            case "a":
                                var v9 int64 = 1
                                _ = v9
                            case "\n\t":
                                if ((-735 > 9223372036854775807) && ((^(v2)) < -1)) {
                                    if len(v3) > 0 {
                                        v3[0] = (fi6 | (fi6 | 1))
                                    }
                                    if ((^(v1)) == (v1 / (((v1 & 4294967295)) | 1))) {
                                        var v10 int16 = 1000
                                        v4 = v4
                                        _ = v10
                                    }
                                }
                            }
                            fmt.Println(v2, v5)
                        } else {
                            var v11 int32 = v2
                            fmt.Println(fi7, fi7)
                            v1 += v1
                            _ = v11
                        }
                    }
                } else {
                    v4 = v4
                }
                v1 = (v1 + (v1 + v1))
                _ = fi7
            }
            _ = fi6
        }
        if false {
            switch "\n\t" {
            case "a":
                v12 := GenS1{F0: 65535, F1: 813, F2: v1}
                _ = v12
            }
            if len(v3) > 0 {
                v3[0] = -891
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                if false {
                    var v14 uint16 = 0
                    switch fi13 {
                    case 1:
                        v2 -= (v2 % ((v2) | 1))
                    }
                    for fi15 := 0; fi15 < 13; fi15++ {
                        v16 := GenS1{F0: (v14 << ((1) & 63)), F1: fi13, F2: (v1 << ((0) & 63))}
                        var v17 int32 = (v2 & (v2 - v2))
                        var v18 uint16 = (^(v14))
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    _ = v14
                }
                if true {
                    var v19 int16 = -778
                    var v20 int16 = (v19 % (((v19 << ((960) & 63))) | 1))
                    _ = v19
                    _ = v20
                } else {
                    switch "" {
                    case "hello":
                        var v21 int8 = (v5 & (v5 * 0))
                        _ = v21
                    }
                    fmt.Println(v4, v2, v1)
                    var v22 uint16 = 554
                    _ = v22
                }
                for fi23 := 0; fi23 < 20; fi23++ {
                    var v24 uint8 = uint8(v1)
                    fmt.Println(v2, v4)
                    if true {
                        var v25 int32 = 1
                        v26 := GenS2{F0: (((("a" + (("\n\t" + ("foo bar" + ("hello" + "\n\t"))) + "foo bar")) + "") + "foo bar") + ""), F1: (v5 & v5), F2: (v25 * v25)}
                        _ = v25
                        _ = v26
                    } else {
                        var v27 uint32 = 967
                        switch fi13 {
                        case 137:
                            _ = len(v3)
                        case 932:
                            v5 = (v5 >> ((592) & 63))
                        }
                        if false {
                            switch 127 {
                            case 0:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                var v28 int8 = (-(v5))
                                _ = v28
                            }
                        } else {
                            v29 := GenS3{F0: -917}
                            v30 := GenS4{F0: 42}
                            var v31 string = "foo bar"
                            v5 -= v5
                            var v32 int16 = 1
                            _ = v29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        _ = v27
                    }
                    var v33 uint8 = v24
                    var v34 bool = (!false)
                    _ = fi23
                    _ = v24
                    _ = v33
                    _ = v34
                }
                _ = fi13
            }
        } else {
            var v35 int64 = 660
            var v36 int16 = 1
            v36 += (^(v36))
            _ = v35
            _ = v36
        }
        v4 = v4
    } else {
        switch 127 {
        case 134:
            var v37 int8 = (v5 << ((112) & 63))
            _ = v37
        case 42:
            var v38 uint16 = 42
            _ = v38
        }
    }
    var v39 float64 = (0.0 / -2.25)
    var v40 uint64 = v1
    v40 = v1
    var v41 float32 = v4
    switch 32767 {
    case 718:
        if (!((!false) && false)) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    fmt.Println(v4, v41, v2)
    var v42 uint8 = 255
    v39 = float64((float32(-2.25) - (0.0 + float32(v39))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
