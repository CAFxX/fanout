package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 int64; F1 int32; F2 int32 }
type GenS6 struct { F0 int64; F1 uint32; F2 int32 }

func main() {
    var v1 int8 = 1
    v2 := []string{}
    var v3 int8 = v1
    var v4 uint16 = 0
    var v5 uint64 = 0
    var v6 uint8 = 128
    if (3.14159 <= (0.0 - 1e10)) {
        if true {
            var v7 uint32 = 455
            if ("hello" >= "foo bar") {
                var v8 int16 = -1
                fmt.Println(v3, v4)
                _ = v8
            }
            for fi9 := 0; fi9 < 1; fi9++ {
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 uint32 = (v7 >> ((449) & 63))
                    var v12 uint16 = (v4 % ((280) | 1))
                    v13 := GenS1{F0: 32767}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                var v14 uint64 = (v5 / ((v5) | 1))
                _ = fi9
                _ = v14
            }
            _ = v7
        }
        var v15 float32 = ((1e10 + ((0.1 + 1.5) / 0.1)) / 1.5)
        _ = v15
    }
    v16 := GenS2{F0: 128}
    if false {
        switch "foo bar" {
        case "hello":
            v1 = (^(v3))
        default:
            v1 = v1
        }
        v1 += v3
    } else {
        fmt.Println(v3)
    }
    v17 := GenS1{F0: -321}
    if false {
        fmt.Println(v5, v3)
        for fi18 := 0; fi18 < 16; fi18++ {
            if false {
                var v19 int = 4294967295
                _ = v19
            } else {
                if (v4 < v4) {
                    switch (fi18 ^ fi18) {
                    case 934:
                        var v20 string = ""
                        _ = v20
                    default:
                        var v21 float32 = 1.5
                        _ = v21
                    }
                    v3 -= (^(v1))
                }
            }
            var v22 float32 = 3.14159
            _ = fi18
            _ = v22
        }
    } else {
        v4 -= 748
        for fi23 := 0; fi23 < 17; fi23++ {
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 int16 = 32767
                _ = fi24
                _ = v25
            }
            if false {
                v6 += v6
                fmt.Println(v5, v4, v4)
            }
            fmt.Println(fi23, v4, fi23)
            _ = fi23
        }
        fmt.Println(v5)
    }
    v1 += v3
    if (2147483647 >= 2) {
        fmt.Println(v1, v5)
        var v26 uint64 = v5
        _ = v26
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 int16 = int16(v6)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi27
            _ = v28
        }
    }
    switch ((((("hello" + "") + "a") + "a") + ((("\n\t" + ("a" + "foo bar")) + ("" + "\n\t")) + "hello")) + "hello") {
    case "\n\t":
        v6 += v6
    case "a":
        v29 := GenS3{F0: "foo bar", F1: (v1 ^ (-(v1)))}
        _ = v29
    case "hello":
        for fi30 := 0; fi30 < 1; fi30++ {
            v31 := GenS4{F0: 42, F1: -554}
            _ = fi30
            _ = v31
        }
    }
    var v32 float64 = (float64(1.5) - 3.14159)
    v32 -= v32
    if false {
        v1 = -128
        if (-1 == -1) {
            v33 := GenS5{F0: -9223372036854775808, F1: -2147483648, F2: -1}
            for fi34 := 0; fi34 < 6; fi34++ {
                v32 = v32
                switch (fi34 * (fi34 * fi34)) {
                case 4294967295:
                    if len(v2) > 0 {
                        v2[0] = (("hello" + (("foo bar" + ("foo bar" + "")) + (("foo bar" + "foo bar") + (("foo bar" + "hello") + ((((("" + "a") + "hello") + "hello") + "hello") + "\n\t"))))) + "\n\t")
                    }
                case 32767:
                    v2 = append(v2, ("\n\t" + ("a" + "foo bar")))
                default:
                    v6 = (v6 ^ v6)
                }
                _ = len(v2)
                _ = fi34
            }
            for fi35 := 0; fi35 < 12; fi35++ {
                var v36 string = "\n\t"
                v5 -= 2
                v37 := GenS6{F0: 1, F1: 2, F2: 1}
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = v33
        }
    }
    var v38 int8 = (v1 + 127)
    var v39 uint64 = v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
