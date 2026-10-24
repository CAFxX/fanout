package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint16 = 1
    var v2 int64 = 1
    var v3 int16 = 1000
    v4 := []int{0, 466, 2147483647}
    fmt.Println(v1, v1)
    switch "a" {
    case "a":
        if (128 <= 127) {
            if (false || true) {
                v1 = (^(v1))
                var v5 bool = true
                var v6 bool = v5
                _ = v5
                _ = v6
            } else {
                fmt.Println(v1, v1, v2)
            }
        }
    case "":
        for fi7 := 0; fi7 < 6; fi7++ {
            if false {
                if (!((("foo bar" + "") + "hello") >= "foo bar")) {
                    if (!false) {
                        v8 := GenS1{F0: 72, F1: v2, F2: false}
                        for fi9 := 0; fi9 < 13; fi9++ {
                            v3 += (v3 << ((2) & 63))
                            if true {
                                var v10 uint16 = 32767
                                if len(v4) > 0 {
                                    v4[0] = int(v2)
                                }
                                if (201 == 255) {
                                    var v11 int16 = v3
                                    var v12 int32 = 0
                                    v3 = 0
                                    _ = v11
                                    _ = v12
                                }
                                v13 := GenS1{F0: 42, F1: 848, F2: true}
                                _ = v10
                                _ = v13
                            }
                            _ = fi9
                        }
                        v1 = (v1 % ((0) | 1))
                        _ = v8
                    }
                    var v14 float64 = 0.0
                    var v15 uint = 65535
                    _ = v14
                    _ = v15
                }
                var v16 float64 = -2.25
                if (538 == -1) {
                    v16 = float64(((0.1 * -2.25) * float32(1.5)))
                } else {
                    var v17 int16 = (v3 - (v3 + v3))
                    switch fi7 {
                    case 9223372036854775807:
                        var v18 int = (fi7 * -788)
                        _ = v18
                    }
                    if true {
                        var v19 string = "a"
                        fmt.Println(fi7, v17, v19)
                        v3 -= -238
                        _ = v19
                    }
                    _ = v17
                }
                _ = v16
            }
            v20 := GenS1{F0: -51, F1: v2, F2: (2147483648 < 4294967295)}
            _ = fi7
            _ = v20
        }
    }
    v3 += 1000
    v21 := GenS2{F0: 65535, F1: 0, F2: 4294967295}
    var v22 uint8 = 255
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(v22, v2)
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 int = fi24
            switch "\n\t" {
            case "foo bar":
                v26 := GenS3{F0: 803, F1: 65535}
                _ = v26
            case "":
                v27 := GenS4{F0: -1, F1: v22, F2: 0}
                _ = v27
            case "\n\t":
                fmt.Println(v25, v22, v25)
            default:
                v1 += 0
            }
            v2 += v2
            var v28 int32 = 42
            _ = fi24
            _ = v25
            _ = v28
        }
        v1 += (^(v1))
        _ = fi23
    }
    switch "hello" {
    case "foo bar":
        if (!(v1 <= (^(v1)))) {
            if true {
                v29 := GenS1{F0: 36, F1: 1, F2: true}
                _ = v29
            } else {
                v2 = v2
                for fi30 := 0; fi30 < 8; fi30++ {
                    var v31 string = ("\n\t" + "hello")
                    var v32 float32 = 0.1
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        } else {
            for fi33 := 0; fi33 < 16; fi33++ {
                var v34 int = 2147483647
                var v35 bool = false
                _ = fi33
                _ = v34
                _ = v35
            }
            var v36 string = "\n\t"
            v37 := GenS2{F0: 4294967295, F1: 120, F2: 255}
            _ = v36
            _ = v37
        }
    case "\n\t":
        fmt.Println(v2, v3, v1)
    default:
        var v38 int64 = v2
        _ = v38
    }
    switch ("" + (((("foo bar" + ("hello" + ("hello" + ("hello" + ("hello" + ("\n\t" + "a")))))) + ("foo bar" + "hello")) + ("\n\t" + "\n\t")) + "hello")) {
    case "":
        fmt.Println(v2, v2, v22)
    }
    switch 899 {
    case 1:
        for fi39 := 0; fi39 < 5; fi39++ {
            var v40 bool = (!((((v1 + v1) <= v1) && (4294967295 != 565)) || (v22 >= 66)))
            v3 = (v3 << ((494) & 63))
            for fi41 := 0; fi41 < 16; fi41++ {
                v42 := GenS1{F0: 42, F1: (v2 * -432), F2: v40}
                _ = fi41
                _ = v42
            }
            _ = fi39
            _ = v40
        }
    case 32768:
        v43 := GenS2{F0: 821, F1: uint8(v1), F2: 414}
        _ = v43
    case 128:
        var v44 int = 1
        _ = v44
    default:
        for fi45 := 0; fi45 < 14; fi45++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v1 += (v1 / ((2) | 1))
            switch "" {
            case "hello":
                v46 := GenS5{F0: uint16(v22), F1: -1}
                _ = v46
            default:
                v47 := GenS6{F0: 42}
                _ = v47
            }
            _ = fi45
        }
    }
    for fi48 := 0; fi48 < 20; fi48++ {
        fmt.Println(fi48)
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
