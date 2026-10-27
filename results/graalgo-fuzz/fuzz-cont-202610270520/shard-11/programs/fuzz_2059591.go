package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64; F1 int; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 int16 }
type GenS4 struct { F0 string; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int32; F1 int16; F2 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 394
    var v3 float32 = ((1.5 - float32((0.1 / 1.5))) / -2.25)
    switch 32768 {
    case 127:
        v4 := GenS1{F0: 2, F1: 42}
        _ = v4
    case 2147483647:
        if (!true) {
            v2 = v2
        }
    }
    var v5 float64 = (float64(0.0) / ((0.1 - 1.5) / 3.14159))
    switch "\n\t" {
    case "":
        for fi6 := 0; fi6 < 2; fi6++ {
            if (false && true) {
                v3 = v3
            }
            v7 := GenS1{F0: 1, F1: 1}
            _ = fi6
            _ = v7
        }
    case "\n\t":
        var v8 string = ("" + (("a" + "\n\t") + ""))
        _ = v8
    default:
        if true {
            var v9 int16 = 660
            fmt.Println(v9)
            _ = v9
        }
    }
    v5 = 1e10
    v10 := GenS1{F0: 0, F1: 723}
    v11 := []int{-9223372036854775808, 599}
    var v12 int16 = 46
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int32 = 0
        v1[218] = ""
        _ = fi13
        _ = v14
    }
    var v15 int64 = -977
    switch 32768 {
    case 0:
        fmt.Println(v2, v2)
    case 1:
        for fi16 := 0; fi16 < 14; fi16++ {
            switch (fi16 / ((-1) | 1)) {
            case -1:
                v17 := GenS1{F0: 2, F1: (^(v12))}
                _ = v17
            case -186:
                if false {
                    switch ("a" + "") {
                    case "foo bar":
                        var v18 uint = 4294967295
                        _ = v18
                    case "\n\t":
                        for fi19 := 0; fi19 < 17; fi19++ {
                            v15 -= -223
                            _ = fi19
                        }
                    case "hello":
                        _ = len(v1)
                    default:
                        switch (^(fi16)) {
                        case 32768:
                            fmt.Println(v3, v2)
                        case -2147483648:
                            v2 = v2
                        case 0:
                            switch (fi16 % ((fi16) | 1)) {
                            case -605:
                                v20 := GenS1{F0: 32767, F1: 1000}
                                _ = v20
                            case 32767:
                                for fi21 := 0; fi21 < 10; fi21++ {
                                    var v22 float64 = v5
                                    v23 := GenS2{F0: v15, F1: -304, F2: 42}
                                    var v24 uint16 = 32767
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                            }
                        default:
                            v12 += (^(v12))
                        }
                    }
                    var v25 int = (^(fi16))
                    fmt.Println(v2, v15)
                    _ = v25
                }
            default:
                if false {
                    fmt.Println(v5, v15, v2)
                    if (true || false) {
                        if (15 != 1) {
                            var v26 string = ""
                            fmt.Println(v2, v26, v3)
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 12; fi27++ {
                                var v28 int8 = 1
                                v28 = (v28 - (v28 * (v28 * v28)))
                                v29 := GenS1{F0: 706, F1: (v12 & v12)}
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            var v30 int8 = int8(fi16)
                            var v31 uint16 = 1
                            _ = v30
                            _ = v31
                        }
                        var v32 float64 = v5
                        _ = v32
                    }
                    v5 -= (v5 / 1e10)
                }
            }
            _ = fi16
        }
    case -9223372036854775808:
        switch "a" {
        case "hello":
            v3 -= 1e10
        case "foo bar":
            var v33 bool = (!(131 == 0))
            _ = v33
        }
    default:
        var v34 uint64 = v2
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        v36 := GenS1{F0: 505, F1: (^(v12))}
        v1[2147483647] = ((((("foo bar" + "") + "\n\t") + "a") + "hello") + ((("hello" + "foo bar") + "hello") + ("hello" + "foo bar")))
        for fi37 := 0; fi37 < 16; fi37++ {
            if (false && false) {
                if (!false) {
                    v2 -= (v2 % ((v2) | 1))
                }
                v38 := GenS3{F0: (v12 / ((v12) | 1)), F1: 27, F2: int16(fi35)}
                v2 = (^(v2))
                _ = v38
            } else {
                v5 = v5
                v39 := GenS4{F0: "", F1: 189, F2: 2147483648}
                v40 := GenS5{F0: fi35, F1: 0, F2: 124}
                _ = v39
                _ = v40
            }
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    var v41 float32 = 0.0
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if (v12 < (^(v12))) {
        v42 := GenS6{F0: -1, F1: v12, F2: ("hello" + "foo bar")}
        v1[2] = ("foo bar" + "foo bar")
        _ = v42
    } else {
        var v43 uint = 803
        if len(v11) > 0 {
            v11[0] = -252
        }
        for fi44 := 0; fi44 < 9; fi44++ {
            if true {
                v45 := v1[65535]
                var v46 int16 = (v12 << (((v43 + (v43 / (((^(v43))) | 1)))) & 63))
                v41 += v41
                _ = v45
                _ = v46
            }
            var v47 int = fi44
            v2 -= v2
            _ = fi44
            _ = v47
        }
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v41
    _ = v5
}
