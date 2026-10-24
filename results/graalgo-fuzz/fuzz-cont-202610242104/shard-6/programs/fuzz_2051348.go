package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int16; F1 int32; F2 bool }
type GenS4 struct { F0 uint8; F1 int64; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8; F1 string; F2 uint }
type GenS10 struct { F0 string; F1 int8; F2 bool }
func genFn0(p0 uint16, p1 int8) uint8 {
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            if (165 <= 127) {
                if false {
                    var v8 int64 = -1
                    _ = v8
                }
            }
            _ = fi7
        }
        var v9 int = fi6
        _ = fi6
        _ = v9
    }
    p1 += -1
    switch -827 {
    case 76:
        var v10 uint8 = 177
        _ = v10
    case 2147483647:
        v11 := GenS1{F0: 1, F1: 1}
        _ = v11
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := []string{}
    var v2 uint8 = 255
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    var v5 string = ("\n\t" + ("\n\t" + ("hello" + ("\n\t" + ("foo bar" + ("a" + "foo bar"))))))
    switch ("a" + (("" + v5) + "foo bar")) {
    case "foo bar":
        v12 := GenS2{F0: 0, F1: false}
        _ = v12
    }
    v13 := GenS3{F0: 336, F1: -850, F2: true}
    v2 = 249
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 bool = (!true)
        var v16 uint = 0
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 int16 = 42
            delete(v4, fi17)
            var v19 bool = false
            _ = fi17
            _ = v18
            _ = v19
        }
        if true {
            v20 := GenS2{F0: -516, F1: false}
            var v21 int = -434
            switch (v21 ^ -139) {
            case -76:
                v2 += v2
            case 4294967295:
                v22 := GenS1{F0: 4294967295, F1: 0}
                _ = v22
            }
            _ = v20
            _ = v21
        } else {
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint = 357
                if (0 != 42) {
                    v25 := v4[(fi23 << ((v24) & 63))]
                    v26 := GenS4{F0: v2, F1: 840, F2: 42}
                    var v27 int16 = 0
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    if (!(false && true)) {
                        v2 = v2
                    } else {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            switch (v5 + (v5 + v5)) {
                            case "foo bar":
                                v5 = "foo bar"
                            case "":
                                v24 += (^(v24))
                            case "a":
                                v29 := GenS5{F0: 2147483647}
                                _ = v29
                            default:
                                delete(v4, (fi28 & (fi23 + (fi23 | (fi28 | (fi23 | (^(fi23))))))))
                            }
                            if true {
                                if (32767 != 65535) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v24 += v24
                                } else {
                                    if true {
                                        v30 := GenS1{F0: 42, F1: -802}
                                        var v31 float32 = 1.5
                                        _ = v30
                                        _ = v31
                                    } else {
                                        if false {
                                            v2 = (^(v2))
                                            if true {
                                                v5 = "\n\t"
                                                v24 -= (v24 | v24)
                                                var v32 int8 = 42
                                                _ = v32
                                            }
                                            var v33 int8 = 0
                                            _ = v33
                                        } else {
                                            _ = len(v1)
                                            v34 := GenS6{F0: (640 <= 32767)}
                                            _ = v34
                                        }
                                    }
                                }
                                delete(v4, fi28)
                                if ((!(false && ((v2 | genFn0(0, 0)) < (v2 ^ (v2 | (v2 >> ((v24) & 63))))))) || false) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v35 := GenS7{F0: -1, F1: 549, F2: 1}
                                    _ = v35
                                } else {
                                    switch v5 {
                                    case "":
                                        var v36 uint16 = 32767
                                        _ = v36
                                    default:
                                        if (true || (false || (v24 >= (v24 | (v24 / (((v24 / ((v24) | 1))) | 1)))))) {
                                            fmt.Println(fi23, fi23, v24)
                                        }
                                    }
                                }
                            } else {
                                var v37 uint32 = 0
                                if (65535 < 42) {
                                    v38 := GenS8{F0: 42}
                                    var v39 int16 = -1
                                    var v40 uint = (v24 / ((v24) | 1))
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                } else {
                                    v24 = v24
                                }
                                _ = v37
                            }
                            var v41 int64 = 0
                            _ = fi28
                            _ = v41
                        }
                    }
                    for fi42 := 0; fi42 < 7; fi42++ {
                        if false {
                            if false {
                                fmt.Println(fi23)
                                var v43 uint8 = genFn0(42, 65)
                                v5 = "a"
                                _ = v43
                            }
                        } else {
                            if ((true || ((3.14159 - 0.1) == 1.5)) && (56 > -1)) {
                                v5 = (("" + (v5 + "\n\t")) + "a")
                            }
                            var v44 int16 = -32768
                            _ = v44
                        }
                        var v45 uint = v24
                        _ = fi42
                        _ = v45
                    }
                    switch v5 {
                    case "\n\t":
                        switch fi23 {
                        case 127:
                            for fi46 := 0; fi46 < 17; fi46++ {
                                fmt.Println(fi46, v5)
                                if (924 == 65535) {
                                    if true {
                                        v24 = (v24 | v24)
                                        var v47 float64 = (3.14159 + ((3.14159 * 0.1) - float64(0.0)))
                                        if (v24 != 1) {
                                            var v48 int16 = 903
                                            if len(v1) > 0 {
                                                v1[0] = "hello"
                                            }
                                            _ = v48
                                        } else {
                                            v49 := GenS9{F0: -128, F1: v5, F2: v24}
                                            v1 = append(v1, (v5 + v5))
                                            _ = v49
                                        }
                                        _ = v47
                                    }
                                } else {
                                    fmt.Println(v5, v5)
                                }
                                _ = fi46
                            }
                        case -10:
                            var v50 uint16 = uint16(v24)
                            _ = v50
                        }
                    }
                }
                _ = fi23
                _ = v24
            }
            v4[128] = 771
        }
    }
    if true {
        for fi51 := 0; fi51 < 6; fi51++ {
            v2 = 2
            _ = fi51
        }
    }
    for fi52 := 0; fi52 < 18; fi52++ {
        v53 := GenS10{F0: v5, F1: 42, F2: (-64 > 42)}
        var v54 float32 = float32(1.5)
        if true {
            for fi55 := 0; fi55 < 6; fi55++ {
                var v56 bool = true
                v5 = ("" + v5)
                _ = fi55
                _ = v56
            }
        }
        _ = fi52
        _ = v53
        _ = v54
    }
    v5 = ("hello" + v5)
    if true {
        v2 += (^(v2))
    } else {
        v57 := GenS10{F0: "\n\t", F1: 0, F2: (false && (-9223372036854775808 <= 128))}
        _ = v57
    }
    delete(v3, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
