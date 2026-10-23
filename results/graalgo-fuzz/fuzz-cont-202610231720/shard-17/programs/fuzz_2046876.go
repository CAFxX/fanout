package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int8; F2 int }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 int16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint; F1 int16 }
type GenS11 struct { F0 uint8; F1 int32; F2 int8 }
func genFn0() int64 {
    var v5 uint = 0
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: true}
        v5 = v5
        if true {
            var v8 int64 = -476
            if ((false || ((-117 == -2147483648) || false)) || (!false)) {
                if (float64(1.5) != 3.14159) {
                    switch "a" {
                    case "hello":
                        v8 += v8
                    }
                    if (763 < 2) {
                        v9 := GenS1{F0: false}
                        for fi10 := 0; fi10 < 14; fi10++ {
                            var v11 uint8 = 207
                            _ = fi10
                            _ = v11
                        }
                        var v12 uint64 = 42
                        _ = v12
                        _ = v9
                    } else {
                        v13 := []int{}
                        switch 32767 {
                        case 0:
                            switch ("hello" + (("\n\t" + "hello") + "\n\t")) {
                            case "a":
                                var v14 uint16 = 42
                                _ = v14
                            case "":
                                if len(v13) > 0 {
                                    v13[0] = (fi6 * (fi6 / (((^(fi6))) | 1)))
                                }
                            case "\n\t":
                                var v15 uint = (v5 ^ 787)
                                _ = v15
                            }
                        case -1:
                            var v16 uint = (^(v5))
                            _ = v16
                        default:
                            var v17 uint16 = 1
                            _ = v17
                        }
                        _ = v13
                    }
                    if true {
                        if (65535 != (^(v5))) {
                            for fi18 := 0; fi18 < 2; fi18++ {
                                var v19 int8 = 1
                                _ = fi18
                                _ = v19
                            }
                            var v20 string = ""
                            v8 -= -1
                            _ = v20
                        } else {
                            if (0 > 128) {
                                v21 := []int{fi6, fi6}
                                var v22 float32 = 0.0
                                _ = v21
                                _ = v22
                            } else {
                                v23 := []int{fi6, fi6}
                                v24 := GenS2{F0: -77}
                                _ = v23
                                _ = v24
                            }
                        }
                    } else {
                        if false {
                            if true {
                                switch "" {
                                case "":
                                    fmt.Println(v8, v5, v5)
                                }
                            } else {
                                if false {
                                    var v25 uint64 = 304
                                    var v26 uint16 = uint16(v25)
                                    fmt.Println(v25)
                                    _ = v25
                                    _ = v26
                                } else {
                                    if (-628 == -1) {
                                        for fi27 := 0; fi27 < 16; fi27++ {
                                            v8 = (v8 % ((v8) | 1))
                                            v28 := GenS3{F0: 42, F1: 417, F2: 336}
                                            fmt.Println(fi27, v5)
                                            _ = fi27
                                            _ = v28
                                        }
                                        if false {
                                            var v29 int8 = -81
                                            _ = v29
                                        } else {
                                            v8 -= (v8 + -358)
                                            var v30 float64 = float64(((0.1 + (0.1 + ((float32(float64(1.5)) - (((1.5 / 1.0) + (1e10 - 0.1)) - 1.5)) / 3.14159))) - 1e10))
                                            var v31 uint32 = 0
                                            _ = v30
                                            _ = v31
                                        }
                                    }
                                    v32 := []int{}
                                    _ = v32
                                }
                                for fi33 := 0; fi33 < 18; fi33++ {
                                    fmt.Println(v8, v8, fi6)
                                    v34 := []int{0, fi33}
                                    v5 = v5
                                    _ = fi33
                                    _ = v34
                                }
                                v8 = (v8 << (((v5 * (v5 - (v5 | 255)))) & 63))
                            }
                            for fi35 := 0; fi35 < 13; fi35++ {
                                var v36 uint64 = 584
                                _ = fi35
                                _ = v36
                            }
                            v8 -= v8
                        } else {
                            if false {
                                v37 := GenS4{F0: 32767}
                                if false {
                                    if (true && false) {
                                        v8 = (-(v8))
                                        v8 -= v8
                                        var v38 uint32 = 2
                                        _ = v38
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            var v39 uint32 = 0
                                            _ = v39
                                        default:
                                            var v40 int64 = (v8 << ((v5) & 63))
                                            _ = v40
                                        }
                                    }
                                    v8 += (v8 | -1)
                                    var v41 int8 = 100
                                    _ = v41
                                }
                                v5 -= v5
                                _ = v37
                            }
                            v42 := []int{fi6}
                            _ = v42
                        }
                        v43 := GenS5{F0: 32768, F1: -1, F2: fi6}
                        _ = v43
                    }
                }
            }
            var v44 int16 = 32767
            _ = v44
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    for fi45 := 0; fi45 < 7; fi45++ {
        v5 += (^(v5))
        _ = fi45
    }
    v46 := GenS6{F0: -73, F1: ("foo bar" + ("foo bar" + ("a" + ("\n\t" + "a"))))}
    var v47 uint64 = 958
    _ = v46
    _ = v47
    _ = v5
    return 1
}

func main() {
    v1 := []string{"hello", (("hello" + ("hello" + "")) + "foo bar")}
    var v2 uint8 = 58
    var v3 uint8 = 128
    v4 := make(map[string]string)
    var v48 int = -989
    switch v48 {
    case 169:
        v49 := GenS7{F0: 42}
        _ = v49
    case 2147483647:
        v50 := GenS8{F0: 1, F1: -32768}
        _ = v50
    case 32767:
        var v51 bool = false
        _ = v51
    default:
        switch ("foo bar" + ("foo bar" + "\n\t")) {
        case "hello":
            v52 := GenS2{F0: -111}
            _ = v52
        default:
            for fi53 := 0; fi53 < 13; fi53++ {
                switch v48 {
                case 209:
                    var v54 bool = (!(127 == 100))
                    _ = v54
                case -901:
                    v48 -= (-(v48))
                }
                for fi55 := 0; fi55 < 12; fi55++ {
                    v56 := GenS9{F0: 1}
                    v57 := GenS10{F0: 0, F1: -32768}
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                _ = fi53
            }
        }
    }
    v48 = v48
    v48 = (v48 + v48)
    if false {
        if true {
            var v58 uint = 65535
            _ = v58
        }
        fmt.Println(v3, v3)
    } else {
        v48 = v48
        fmt.Println(v2, v48, v48)
        switch (v48 | v48) {
        case 448:
            var v59 string = ("" + "foo bar")
            _ = v59
        default:
            v1 = append(v1, "\n\t")
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi60 := 0; fi60 < 1; fi60++ {
        for fi61 := 0; fi61 < 19; fi61++ {
            fmt.Println(v48)
            _ = fi61
        }
        _ = fi60
    }
    var v62 int = v48
    var v63 uint16 = 32767
    v64 := GenS11{F0: v3, F1: 1, F2: 100}
    fmt.Println(v48, v2, v62)
    v3 -= (^(v3))
    var v65 float32 = (1.5 * 1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
