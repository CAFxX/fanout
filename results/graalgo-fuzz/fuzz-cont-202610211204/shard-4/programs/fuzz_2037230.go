package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64; F1 int64; F2 int }
type GenS3 struct { F0 bool; F1 uint64; F2 int32 }
type GenS4 struct { F0 bool; F1 uint16; F2 int64 }
type GenS5 struct { F0 string; F1 uint16; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 34
    var v3 float64 = -2.25
    v2 = v2
    var v4 uint16 = 42
    v3 -= v3
    if (-685 < 9223372036854775807) {
        v2 = -27
    } else {
        delete(v1, 569)
    }
    var v5 uint64 = 2
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            v4 += (v4 * 1)
            switch "" {
            case "foo bar":
                v8 := GenS1{F0: 1, F1: -32768, F2: -1}
                _ = v8
            default:
                v2 -= 127
            }
            _ = fi7
        }
        v2 = (v2 & -1)
        v4 = (v4 & (v4 / (((v4 & v4)) | 1)))
        _ = fi6
    }
    var v9 uint = 65535
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 float32 = ((1e10 * 1e10) * 1e10)
        if true {
            v2 = v2
        } else {
            v2 = (-(v2))
            if true {
                var v12 uint = 747
                switch "foo bar" {
                case "\n\t":
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v14 := GenS1{F0: 128, F1: 1, F2: 902}
                        var v15 uint = v12
                        var v16 string = "a"
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                case "hello":
                    v17 := GenS1{F0: 38, F1: -32768, F2: 305}
                    _ = v17
                }
                switch (("foo bar" + "\n\t") + "\n\t") {
                case "":
                    v5 += v5
                case "\n\t":
                    delete(v1, 65535)
                case "hello":
                    fmt.Println(v9, v3, v3)
                default:
                    delete(v1, 0)
                }
                _ = v12
            }
            var v18 uint32 = 2
            _ = v18
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 int = fi10
            v5 += (v5 % ((0) | 1))
            _ = fi19
            _ = v20
        }
        _ = fi10
        _ = v11
    }
    v21 := GenS2{F0: 42, F1: -1, F2: 32768}
    if (0 <= 1) {
        v2 = (v2 << ((18446744073709551615) & 63))
        var v22 int64 = -1
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := GenS3{F0: true, F1: v5, F2: 42}
            v25 := GenS4{F0: false, F1: v4, F2: -1}
            if (false || (!(2147483647 <= 2147483647))) {
                switch (fi23 / ((2147483647) | 1)) {
                case -1:
                    v4 -= v4
                case 128:
                    var v26 uint32 = 1
                    _ = v26
                case 4294967295:
                    v22 = 9223372036854775807
                }
                v27 := GenS1{F0: 154, F1: 949, F2: (v22 / ((42) | 1))}
                _ = v27
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v22
    } else {
        if true {
            _ = len(v1)
        }
        v3 = (v3 - float64(((0.1 + (0.0 * (1e10 / 1e10))) * -2.25)))
        var v28 uint16 = (v4 - (v4 + (v4 << (((^(v9))) & 63))))
        _ = v28
    }
    v29 := GenS2{F0: (^(v5)), F1: 1, F2: 0}
    if (((("hello" + ("" + ("" + ("" + "foo bar")))) + (("\n\t" + "foo bar") + "")) + ("a" + ("hello" + "hello"))) == "") {
        switch -57 {
        case -1:
            for fi30 := 0; fi30 < 9; fi30++ {
                v31 := GenS5{F0: "a", F1: (v4 & 0), F2: (v9 * (^(v9)))}
                var v32 uint64 = v5
                v33 := GenS5{F0: ("\n\t" + ("foo bar" + "\n\t")), F1: (^(v4)), F2: v9}
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
        case 32768:
            v34 := []int{201, 32767, -564, -1}
            _ = v34
        case -9223372036854775808:
            var v35 int = 486
            _ = v35
        }
        var v36 int64 = -484
        var v37 float64 = (v3 / (v3 - 0.0))
        _ = v36
        _ = v37
    } else {
        v38 := v1[513]
        _ = v38
    }
    v5 = v5
    v3 = 3.14159
    v39 := GenS5{F0: "a", F1: (v4 % (((^(v4))) | 1)), F2: v9}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v9
}
