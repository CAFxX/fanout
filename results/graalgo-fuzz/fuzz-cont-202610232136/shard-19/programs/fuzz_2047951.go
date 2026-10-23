package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 int16 }
type GenS3 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -1
    var v3 float32 = -2.25
    fmt.Println(v2, v3)
    v2 += v2
    v2 = 419
    for fi4 := 0; fi4 < 17; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 float32 = 0.1
            _ = fi5
            _ = v6
        }
        switch "" {
        case "hello":
            v7 := GenS1{F0: (fi4 >> ((42) & 63)), F1: 701}
            _ = v7
        case "":
            if false {
                for fi8 := 0; fi8 < 13; fi8++ {
                    v9 := GenS1{F0: 883, F1: 2147483647}
                    var v10 uint32 = 65535
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
        case "foo bar":
            if (!(!false)) {
                switch "foo bar" {
                case "":
                    v2 = 42
                case "hello":
                    v11 := GenS1{F0: -9223372036854775808, F1: 2}
                    _ = v11
                default:
                    var v12 uint16 = 42
                    _ = v12
                }
                v13 := GenS1{F0: (v2 | (v2 * -2147483648)), F1: 2}
                _ = v13
            }
        }
        _ = fi4
    }
    if (105 <= -128) {
        for fi14 := 0; fi14 < 4; fi14++ {
            v2 = v2
            _ = fi14
        }
        fmt.Println(v3, v3)
        if (false && true) {
            v2 += (v2 | -2147483648)
            fmt.Println(v2, v2)
            if false {
                var v15 uint32 = 4294967295
                v16 := v1[(("foo bar" + "hello") + (("hello" + ("foo bar" + (("hello" + "") + ("hello" + "\n\t")))) + ("" + "foo bar")))]
                v3 -= (((v3 * v3) - 3.14159) - float32(-2.25))
                _ = v15
                _ = v16
            } else {
                _ = len(v1)
                switch ("a" + "foo bar") {
                case "\n\t":
                    if (!true) {
                        for fi17 := 0; fi17 < 18; fi17++ {
                            v2 += (-(v2))
                            var v18 int = v2
                            v19 := GenS2{F0: -451, F1: false, F2: -476}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        var v20 int8 = 100
                        _ = v20
                    }
                case "":
                    v3 += (v3 + 0.1)
                default:
                    var v21 int = v2
                    _ = v21
                }
                var v22 bool = true
                _ = v22
            }
        }
    } else {
        if true {
            var v23 int32 = 42
            _ = v23
        } else {
            fmt.Println(v3, v3)
        }
    }
    fmt.Println(v2, v3)
    if (128 >= 255) {
        v3 += v3
    }
    var v24 float32 = (float32(1e10) + v3)
    var v25 float32 = (((v3 / v24) - (v24 - (v24 / 3.14159))) - 3.14159)
    for fi26 := 0; fi26 < 10; fi26++ {
        if (42 > 100) {
            v27 := []int{}
            var v28 uint64 = 0
            v29 := v1[("a" + "")]
            _ = v27
            _ = v28
            _ = v29
        }
        if true {
            v25 -= float32(1.5)
            v3 = 3.14159
        }
        _ = fi26
    }
    if false {
        var v30 float32 = v24
        if true {
            for fi31 := 0; fi31 < 11; fi31++ {
                for fi32 := 0; fi32 < 9; fi32++ {
                    var v33 string = ("foo bar" + ("\n\t" + "foo bar"))
                    switch (v33 + (v33 + "hello")) {
                    case "a":
                        if false {
                            switch (fi32 << ((255) & 63)) {
                            case -875:
                                var v34 int32 = -965
                                _ = v34
                            }
                        } else {
                            if (true && false) {
                                var v35 int = 32767
                                fmt.Println(fi31, fi31)
                                var v36 uint8 = 127
                                _ = v35
                                _ = v36
                            }
                            var v37 int = (v2 - (v2 / (((fi32 >> ((0) & 63))) | 1)))
                            var v38 int16 = 924
                            _ = v37
                            _ = v38
                        }
                    case "":
                        v2 -= 128
                        var v39 float64 = 1.5
                        _ = v39
                    case "foo bar":
                        if (127 == 9) {
                            fmt.Println(fi31, v30)
                        }
                    default:
                        v2 = (v2 | fi32)
                    }
                    if (0.0 == 1.5) {
                        var v40 float64 = 1.5
                        _ = v40
                    } else {
                        var v41 int32 = 0
                        v42 := GenS2{F0: 1000, F1: false, F2: 1000}
                        _ = v41
                        _ = v42
                    }
                    _ = fi32
                    _ = v33
                }
                _ = fi31
            }
        }
        _ = v30
    }
    delete(v1, "\n\t")
    v2 -= (v2 >> ((613) & 63))
    v43 := GenS3{F0: -9223372036854775808, F1: -1}
    var v44 uint16 = 2
    v45 := GenS2{F0: -1, F1: false, F2: -619}
    var v46 uint8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
