package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint16 }

func main() {
    v1 := []string{"foo bar", "a", "foo bar"}
    var v2 uint64 = 2
    var v3 uint = 1
    var v4 float64 = (1.5 / 1.0)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v5 int8 = 1
    var v6 float64 = v4
    switch (("\n\t" + ("" + "hello")) + ("foo bar" + "hello")) {
    case "\n\t":
        v2 = v2
    default:
        v3 = v3
    }
    fmt.Println(v5)
    fmt.Println(v5, v5, v2)
    var v7 int8 = 0
    var v8 float32 = ((-2.25 + 0.0) / (1e10 - ((3.14159 * 3.14159) + float32(v4))))
    v1 = append(v1, ("\n\t" + "a"))
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(v6, v4, v3)
        var v10 int32 = -2147483648
        v11 := GenS1{F0: uint(v2), F1: 0}
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v2, v4, v4)
    v3 += 255
    switch ("hello" + "a") {
    case "\n\t":
        if (4294967295 <= 65535) {
            var v12 bool = false
            _ = v12
        } else {
            var v13 uint32 = 536
            _ = v13
        }
    case "foo bar":
        if (-814 >= -783) {
            v8 -= v8
        }
    case "a":
        if (1 < 37) {
            v14 := GenS2{F0: 0}
            v6 = (v4 * (v4 / (v4 - v4)))
            v3 = v3
            _ = v14
        }
    default:
        v6 = ((v6 + -2.25) / ((float64((-2.25 - (v8 * v8))) * v4) / (v4 - 1e10)))
    }
    v7 += v7
    var v15 int16 = 0
    for fi16 := 0; fi16 < 19; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            switch "\n\t" {
            case "a":
                v3 = 1
                var v18 float64 = float64((v8 - 1.5))
                fmt.Println(fi16)
                _ = v18
            case "hello":
                if (("" + "a") <= "\n\t") {
                    if (!(false || (!false))) {
                        v3 = v3
                        var v19 string = ("foo bar" + "\n\t")
                        v3 += 87
                        if true {
                            v20 := GenS1{F0: (^(v3)), F1: 0}
                            var v21 bool = (0 > v3)
                            _ = v20
                            _ = v21
                        }
                        _ = v19
                    } else {
                        v5 = v7
                        if (false && true) {
                            if false {
                                var v22 uint16 = 65535
                                v2 -= (v2 >> ((635) & 63))
                                _ = len(v1)
                                fmt.Println(fi16, v3)
                                if false {
                                    var v23 uint16 = v22
                                    v24 := GenS3{F0: 2, F1: 2147483647, F2: (v2 * v2)}
                                    v25 := GenS4{F0: 255, F1: -1, F2: v23}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                _ = v22
                            }
                            if (29 >= 190) {
                                fmt.Println(v15, v2, v6)
                                if true {
                                    var v26 int64 = 0
                                    _ = v26
                                }
                                v6 = v4
                                var v27 uint16 = 0
                                _ = v27
                            }
                            if len(v1) > 0 {
                                v1[0] = ("" + "")
                            }
                        }
                        v15 -= v15
                    }
                } else {
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        var v28 uint8 = 0
                        _ = v28
                    default:
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                    }
                }
            }
            _ = fi17
        }
        fmt.Println(v2)
        _ = fi16
    }
    if true {
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 int = fi29
            if (250 <= uint8(v2)) {
                fmt.Println(fi29, v7)
                if (true && (3.14159 > (0.1 * (float64(v8) * v6)))) {
                    switch fi29 {
                    case 32767:
                        var v31 int = fi29
                        _ = v31
                    case -2147483648:
                        for fi32 := 0; fi32 < 11; fi32++ {
                            v6 -= (3.14159 - 0.1)
                            var v33 float64 = (float64((v8 + float32(v4))) * (v4 + v4))
                            if false {
                                var v34 uint8 = 0
                                _ = v34
                            } else {
                                fmt.Println(v33, v8, v30)
                            }
                            v2 = 42
                            fmt.Println(v5)
                            _ = fi32
                            _ = v33
                        }
                    default:
                        v1 = append(v1, "hello")
                    }
                }
                for fi35 := 0; fi35 < 13; fi35++ {
                    v6 = v6
                    var v36 int8 = (v7 / ((95) | 1))
                    _ = fi35
                    _ = v36
                }
            }
            _ = fi29
            _ = v30
        }
    } else {
        v2 = (v2 << (((v3 * 1)) & 63))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
