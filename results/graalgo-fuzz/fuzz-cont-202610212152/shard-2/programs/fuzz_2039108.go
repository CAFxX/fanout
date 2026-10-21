package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 bool }
type GenS3 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint16 = 0
    v2 := []string{"hello"}
    var v3 uint = 4294967295
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    var v6 float32 = float32(-2.25)
    var v7 bool = true
    v3 = (v3 << (((^(v3))) & 63))
    switch 9223372036854775807 {
    case 2:
        v8 := GenS1{F0: 920}
        _ = v8
    default:
        var v9 int = 32767
        _ = v9
    }
    var v10 uint = (v3 << ((v3) & 63))
    v10 += 255
    if (1.5 >= 1e10) {
        v11 := GenS2{F0: uint64(v3), F1: 0, F2: v7}
        var v12 uint32 = 299
        fmt.Println(v10, v1, v12)
        _ = v11
        _ = v12
    }
    switch 9223372036854775807 {
    case 32768:
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 uint = (v3 + (v3 & (v10 | v3)))
            for fi15 := 0; fi15 < 6; fi15++ {
                if (!v7) {
                    var v16 int8 = -128
                    switch 2147483647 {
                    case 42:
                        v7 = false
                    case -325:
                        if v7 {
                            v7 = v7
                        } else {
                            v6 = ((v6 / (v6 / v6)) + v6)
                        }
                    case -1:
                        var v17 int16 = -476
                        _ = v17
                    default:
                        fmt.Println(v1, v3)
                    }
                    _ = v16
                }
                _ = fi15
            }
            var v18 uint32 = 149
            _ = fi13
            _ = v14
            _ = v18
        }
    }
    v6 = (1.5 - v6)
    if (-9223372036854775808 < 100) {
        var v19 uint16 = v1
        var v20 float64 = 1.5
        switch "foo bar" {
        case "":
            var v21 float32 = float32(-2.25)
            _ = v21
        }
        _ = v19
        _ = v20
    }
    if (((127 != -120) && v7) || ((((0.1 - (float32(1e10) * v6)) / float32(3.14159)) * (v6 / 0.1)) < v6)) {
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 uint16 = 65535
            if (-235 == -2147483648) {
                switch (-(fi22)) {
                case 9223372036854775807:
                    fmt.Println(v7, v6, fi22)
                case -556:
                    var v24 float64 = 1.5
                    _ = v24
                }
            } else {
                v25 := GenS3{F0: 128, F1: -32768, F2: 1}
                var v26 int = 2147483647
                for fi27 := 0; fi27 < 3; fi27++ {
                    v28 := GenS4{F0: 4294967295}
                    _ = fi27
                    _ = v28
                }
                _ = v25
                _ = v26
            }
            _ = fi22
            _ = v23
        }
        for fi29 := 0; fi29 < 19; fi29++ {
            if true {
                var v30 uint16 = (v1 & (v1 & v1))
                _ = v30
            } else {
                var v31 uint16 = (v1 & (v1 << ((v3) & 63)))
                if ((float32(3.14159) != (3.14159 + 0.0)) && true) {
                    var v32 string = (("a" + (("foo bar" + (((((((((("" + "hello") + "") + ("\n\t" + ("a" + "hello"))) + "hello") + "\n\t") + "hello") + "foo bar") + ("a" + ("a" + "\n\t"))) + (("a" + "a") + ("\n\t" + "\n\t"))) + "")) + ((("" + ("foo bar" + ((("hello" + "hello") + "") + "foo bar"))) + ("" + "a")) + "hello"))) + "hello")
                    if (!v7) {
                        switch (v32 + "") {
                        case "foo bar":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "\n\t":
                            var v33 int32 = -1
                            _ = v33
                        }
                        v6 += (-2.25 / ((1e10 / v6) - v6))
                        var v34 uint32 = 2
                        _ = v34
                    } else {
                        if (float32(((0.1 - (float64((v6 * (float32(0.1) + v6))) + -2.25)) * (1e10 + 0.1))) == v6) {
                            fmt.Println(v7, v1)
                        }
                    }
                    _ = v32
                }
                if v7 {
                    delete(v5, "")
                    for fi35 := 0; fi35 < 14; fi35++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v36 bool = (!v7)
                        v10 = 427
                        _ = fi35
                        _ = v36
                    }
                }
                _ = v31
            }
            for fi37 := 0; fi37 < 20; fi37++ {
                fmt.Println(fi37, v3, fi37)
                v4[247] = "hello"
                v38 := GenS2{F0: 2, F1: 42, F2: false}
                _ = fi37
                _ = v38
            }
            _ = fi29
        }
    } else {
        _ = len(v2)
    }
    v10 += (v10 * uint(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
