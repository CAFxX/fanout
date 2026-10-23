package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 uint8 }

func main() {
    var v1 int8 = 41
    v2 := make(map[string]int)
    var v3 uint8 = 203
    var v4 float32 = 1.5
    var v5 float64 = 0.0
    v6 := make(map[string]int)
    v3 -= v3
    if false {
        v7 := GenS1{F0: ("\n\t" + ("a" + "\n\t"))}
        var v8 uint = 255
        _ = v7
        _ = v8
    }
    v1 = 1
    var v9 uint64 = 0
    v10 := GenS1{F0: "hello"}
    v5 = (float64(v4) * ((float64((((((v4 * 0.1) - (1e10 - 0.0)) + v4) + (v4 + float32(float64(v4)))) + float32(v5))) * v5) - (((-2.25 + v5) * v5) + (v5 * float64(v4)))))
    switch -2147483648 {
    case 843:
        if false {
            v11 := []int{550}
            _ = v11
        } else {
            v3 = uint8(v9)
        }
    case 9223372036854775807:
        switch 1 {
        case 148:
            var v12 int32 = -1
            _ = v12
        case 566:
            v4 = 1.5
        default:
            for fi13 := 0; fi13 < 6; fi13++ {
                if false {
                    v14 := GenS1{F0: (("\n\t" + "\n\t") + ("hello" + "\n\t"))}
                    if true {
                        v15 := []int{}
                        _ = v15
                    }
                    _ = v14
                } else {
                    v9 = v9
                    v3 += v3
                    switch (fi13 << ((2) & 63)) {
                    case 264:
                        v16 := GenS2{F0: 1}
                        _ = v16
                    case -1:
                        for fi17 := 0; fi17 < 13; fi17++ {
                            v5 += ((v5 - (v5 + float64(0.0))) * 0.1)
                            _ = fi17
                        }
                    case 479:
                        v9 = (v9 >> ((2) & 63))
                    default:
                        switch "\n\t" {
                        case "a":
                            if false {
                                var v18 uint8 = (v3 % (((v3 | (v3 - 127))) | 1))
                                v1 -= 100
                                _ = v18
                            }
                        case "\n\t":
                            v1 = (v1 % ((-1) | 1))
                        }
                    }
                }
                if (false || false) {
                    var v19 uint32 = 65535
                    if (!(18446744073709551615 == 4294967295)) {
                        var v20 string = "foo bar"
                        v9 -= (v9 << ((65535) & 63))
                        _ = v20
                    } else {
                        fmt.Println(v5, v3)
                        v21 := GenS2{F0: 32767}
                        if (1 <= 4294967295) {
                            v22 := GenS3{F0: 0}
                            v3 = uint8(v9)
                            v23 := v2[("a" + ("" + (("hello" + ("foo bar" + ("" + "a"))) + "a")))]
                            _ = v22
                            _ = v23
                        } else {
                            var v24 float32 = float32((1e10 - ((v5 - ((0.0 * -2.25) - v5)) + (((v5 / 1.5) / (v5 * v5)) + (v5 / v5)))))
                            v5 -= float64(((1.5 * 0.0) - v4))
                            _ = v24
                        }
                        _ = v21
                    }
                    if true {
                        delete(v6, "foo bar")
                        switch (("a" + "hello") + ("\n\t" + "\n\t")) {
                        case "foo bar":
                            v25 := GenS3{F0: 0}
                            _ = v25
                        case "hello":
                            v6["a"] = 2147483647
                        }
                    } else {
                        for fi26 := 0; fi26 < 14; fi26++ {
                            if false {
                                switch (-(fi26)) {
                                case -9223372036854775808:
                                    v27 := GenS1{F0: "a"}
                                    _ = v27
                                case 218:
                                    v28 := GenS4{F0: 720}
                                    _ = v28
                                case 9223372036854775807:
                                    fmt.Println(v1, v19, v5)
                                }
                                v29 := GenS5{F0: 42, F1: uint8(v19)}
                                _ = v29
                            }
                            _ = fi26
                        }
                        for fi30 := 0; fi30 < 16; fi30++ {
                            v6["hello"] = 0
                            _ = fi30
                        }
                        v3 = (v3 / ((v3) | 1))
                    }
                    _ = v19
                }
                _ = fi13
            }
        }
    case 295:
        for fi31 := 0; fi31 < 7; fi31++ {
            v1 -= v1
            v9 -= (v9 << ((361) & 63))
            v32 := GenS5{F0: 1, F1: 140}
            _ = fi31
            _ = v32
        }
    default:
        var v33 float64 = v5
        _ = v33
    }
    for fi34 := 0; fi34 < 18; fi34++ {
        for fi35 := 0; fi35 < 9; fi35++ {
            fmt.Println(fi34, v1, v9)
            if (!true) {
                fmt.Println(v4, v3, v5)
                var v36 int8 = -128
                _ = v36
            }
            v37 := GenS2{F0: 0}
            _ = fi35
            _ = v37
        }
        var v38 bool = false
        v39 := GenS4{F0: 65535}
        _ = fi34
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
