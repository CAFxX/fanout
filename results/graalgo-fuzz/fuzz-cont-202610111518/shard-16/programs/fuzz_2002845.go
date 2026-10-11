package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int64; F1 int; F2 int32 }
func genFn0() string {
    switch "a" {
    case "a":
        var v7 int8 = 0
        _ = v7
    default:
        var v8 uint64 = 604
        _ = v8
    }
    if (3.14159 > (float64(float32(0.0)) + 0.0)) {
        var v9 float64 = 0.0
        _ = v9
    }
    var v10 uint32 = 962
    v10 += 716
    _ = v10
    return ("a" + "hello")
}
func genFn1(p0 int64, p1 int) uint {
    v11 := GenS1{F0: (1 == 4294967295), F1: 860, F2: 1}
    p0 += (p0 | (p0 % (((p0 + (-(p0)))) | 1)))
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    var v1 float64 = (0.1 - 0.0)
    var v2 int16 = -642
    v3 := make(map[string]uint32)
    var v4 uint64 = 42
    var v5 uint16 = 836
    v6 := []uint32{343, 140, 1, 49}
    var v12 int = -2147483648
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v4, v2)
    default:
        v4 -= (^(v4))
    }
    var v13 string = ("foo bar" + "")
    var v14 string = (v13 + "\n\t")
    switch (v12 / (((v12 - -381)) | 1)) {
    case -1:
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 float64 = v1
            for fi17 := 0; fi17 < 5; fi17++ {
                fmt.Println(v16, v2)
                v4 -= (v4 & v4)
                var v18 int16 = (v2 & v2)
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
    case 0:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v1 = v1
    var v19 uint16 = (v5 ^ (v5 + 255))
    v2 = (v2 << ((196) & 63))
    v20 := GenS2{F0: -9223372036854775808, F1: v2}
    if (42 < 705) {
        v19 -= v5
        switch v12 {
        case 32768:
            if (float32(((float64(1e10) + (v1 + v1)) - v1)) > -2.25) {
                var v21 int16 = (v2 / ((v2) | 1))
                _ = v21
            } else {
                var v22 uint8 = 0
                for fi23 := 0; fi23 < 11; fi23++ {
                    v14 = (("a" + "") + (v13 + "hello"))
                    _ = fi23
                }
                var v24 int64 = 1
                _ = v22
                _ = v24
            }
        default:
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 uint = 0
                v26 -= 602
                _ = fi25
                _ = v26
            }
        }
        var v27 float32 = 3.14159
        _ = v27
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        for fi29 := 0; fi29 < 18; fi29++ {
            if (0 >= int64(v19)) {
                var v30 string = (v14 + "foo bar")
                switch ((v30 + (v30 + v30)) + "foo bar") {
                case "":
                    _ = len(v6)
                    v1 = v1
                case "foo bar":
                    var v31 uint8 = 0
                    _ = v31
                default:
                    switch (v14 + v30) {
                    case "":
                        v14 = "foo bar"
                    case "foo bar":
                        v1 -= v1
                    default:
                        var v32 uint16 = (v5 % ((v19) | 1))
                        _ = v32
                    }
                }
                _ = v30
            } else {
                _ = len(v3)
                fmt.Println(v14)
                var v33 int32 = -2147483648
                _ = v33
            }
            if true {
                v13 = v13
                v34 := GenS3{F0: 9223372036854775807, F1: (v12 << ((0) & 63)), F2: 42}
                _ = v34
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v35 int32 = 42
                switch "" {
                case "\n\t":
                    v12 = 32767
                default:
                    if (true && false) {
                        v4 = (v4 + v4)
                    }
                }
                _ = v35
            }
            _ = fi29
        }
        _ = fi28
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        v37 := GenS3{F0: -831, F1: v12, F2: 257}
        var v38 int = (fi36 % ((v12) | 1))
        fmt.Println(v14)
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
