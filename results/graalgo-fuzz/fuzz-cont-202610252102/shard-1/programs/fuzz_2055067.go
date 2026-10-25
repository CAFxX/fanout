package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int32; F1 uint8; F2 int16 }
type GenS3 struct { F0 bool; F1 string; F2 uint64 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]string)
    var v3 int32 = -814
    var v4 int64 = 1
    if true {
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            v1 += 956
        case "\n\t":
            v3 = (v3 >> ((534) & 63))
        case "a":
            _ = len(v2)
        }
        switch ("\n\t" + "a") {
        case "":
            var v5 int64 = (v4 + (^(v4)))
            _ = v5
        case "\n\t":
            v1 -= v1
        case "foo bar":
            v1 = (v1 - v1)
        default:
            v4 = v4
        }
        v6 := GenS1{F0: (42 < -107), F1: 765}
        _ = v6
    } else {
        v3 -= (v3 / (((-(v3))) | 1))
    }
    fmt.Println(v3, v3, v3)
    var v7 int8 = -19
    for fi8 := 0; fi8 < 8; fi8++ {
        switch "hello" {
        case "\n\t":
            for fi9 := 0; fi9 < 2; fi9++ {
                v1 -= v1
                delete(v2, "foo bar")
                if (true && true) {
                    v10 := []int{fi9, fi8}
                    var v11 float64 = (1.5 / 1e10)
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint = 1
                    var v13 int32 = v3
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
        default:
            switch (fi8 - (fi8 | fi8)) {
            case 625:
                fmt.Println(v7)
            }
        }
        _ = fi8
    }
    if true {
        v1 = 153
    }
    if (1 < 65535) {
        v4 -= (-(v4))
        for fi14 := 0; fi14 < 10; fi14++ {
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 float32 = 1.5
                var v17 float64 = float64(1e10)
                v16 = ((v16 + v16) - (((float32(((v17 * 1.5) + v17)) * v16) - 1.5) * (v16 + (v16 / (float32(1.5) * 1.5)))))
                if false {
                    var v18 uint = 400
                    v1 += v1
                    v19 := []int{fi14, -833, fi15}
                    _ = v18
                    _ = v19
                } else {
                    if (!(0 > 133)) {
                        v4 = v4
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            _ = len(v2)
                        case "a":
                            v20 := GenS2{F0: (^(v3)), F1: 2, F2: 649}
                            _ = v20
                        }
                        if true {
                            if false {
                                switch "\n\t" {
                                case "foo bar":
                                    var v21 float64 = v17
                                    v22 := GenS3{F0: (((2 != 2) && ((fi14 - 127) != -9223372036854775808)) || ((v4 & (v4 + v4)) != v4)), F1: ("\n\t" + ""), F2: 0}
                                    v7 = -128
                                    _ = v21
                                    _ = v22
                                }
                                v7 = (v7 ^ -128)
                            } else {
                                var v23 uint8 = 128
                                fmt.Println(fi14, v4)
                                v2[""] = ("hello" + ("foo bar" + "foo bar"))
                                _ = v23
                            }
                            var v24 string = ""
                            _ = v24
                        }
                        v25 := []int{2147483647}
                        _ = v25
                    }
                }
                if (false && (112 <= 64)) {
                    var v26 string = ("\n\t" + (("a" + ("foo bar" + ("" + "\n\t"))) + ("" + "a")))
                    var v27 int16 = 0
                    _ = v26
                    _ = v27
                } else {
                    v7 += v7
                    v28 := GenS2{F0: (v3 & (v3 ^ (v3 + v3))), F1: 128, F2: 42}
                    switch fi14 {
                    case 250:
                        fmt.Println(v3, fi15)
                    case 1:
                        v29 := []int{(fi14 * int(v1)), (fi14 * (fi14 - -1))}
                        _ = v29
                    default:
                        fmt.Println(v4, v16, fi14)
                    }
                    _ = v28
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            v30 := GenS4{F0: (^(v4))}
            var v31 float64 = 1e10
            _ = fi14
            _ = v30
            _ = v31
        }
    } else {
        v7 = (-(v7))
    }
    var v32 int8 = v7
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 int = fi33
        _ = fi33
        _ = v34
    }
    var v35 float32 = (-2.25 * 1.5)
    for fi36 := 0; fi36 < 15; fi36++ {
        fmt.Println(v4, v3)
        _ = fi36
    }
    v1 += v1
    fmt.Println(v4, v32)
    v35 = 1e10
    var v37 uint8 = 42
    var v38 int8 = -75
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v7
}
