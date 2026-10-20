package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint16; F2 int16 }
type GenS5 struct { F0 bool; F1 uint8; F2 int16 }

func main() {
    var v1 int32 = 516
    var v2 uint = 18446744073709551615
    v3 := make(map[string]uint32)
    if true {
        if ((100 != 1) || false) {
            var v4 int8 = 100
            _ = v4
        }
        v1 -= 0
        var v5 float64 = 1.5
        _ = v5
    }
    v1 -= v1
    var v6 uint8 = 128
    var v7 uint32 = 65535
    v8 := []int{}
    fmt.Println(v2, v7, v2)
    switch "foo bar" {
    case "\n\t":
        v2 -= (v2 & 743)
    case "hello":
        switch ((((("a" + "hello") + ((("a" + "") + "a") + "a")) + "hello") + ("\n\t" + ("\n\t" + "\n\t"))) + "a") {
        case "":
            v9 := GenS1{F0: 1}
            _ = v9
        case "a":
            var v10 int32 = 135
            _ = v10
        case "hello":
            v3[(("foo bar" + "a") + "foo bar")] = (^(v7))
        default:
            for fi11 := 0; fi11 < 12; fi11++ {
                v1 -= (v1 / ((v1) | 1))
                _ = fi11
            }
        }
    case "":
        var v12 float64 = 0.1
        _ = v12
    }
    var v13 int8 = 42
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint = 1
        var v16 uint8 = (v6 & v6)
        if false {
            for fi17 := 0; fi17 < 14; fi17++ {
                v18 := GenS1{F0: fi17}
                var v19 int8 = (v13 + 46)
                if len(v8) > 0 {
                    v8[0] = 9223372036854775807
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS2{F0: (v1 | (v1 & -1)), F1: int16(v16)}
            if (!(0 <= 255)) {
                var v21 uint32 = v7
                var v22 int64 = -701
                _ = v21
                _ = v22
            }
            _ = v20
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    if false {
        if true {
            v23 := v3["\n\t"]
            var v24 string = ("foo bar" + "foo bar")
            v1 = v1
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 17; fi25++ {
                for fi26 := 0; fi26 < 10; fi26++ {
                    switch 32768 {
                    case -994:
                        _ = len(v8)
                    case -879:
                        fmt.Println(v1)
                    case -729:
                        v27 := GenS2{F0: 0, F1: 32767}
                        _ = v27
                    default:
                        fmt.Println(fi26, v13)
                    }
                    _ = fi26
                }
                v13 -= v13
                v1 -= v1
                _ = fi25
            }
            var v28 uint64 = 42
            var v29 uint64 = v28
            _ = v28
            _ = v29
        }
    }
    v1 = v1
    var v30 int16 = 0
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 uint = 4294967295
        switch fi31 {
        case -759:
            if len(v8) > 0 {
                v8[0] = fi31
            }
        case -2147483648:
            v33 := GenS3{F0: "\n\t"}
            _ = v33
        case 389:
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 uint16 = 1
                _ = fi34
                _ = v35
            }
        default:
            v8 = append(v8, (fi31 << ((v32) & 63)))
        }
        v36 := GenS4{F0: v30, F1: 562, F2: v30}
        _ = fi31
        _ = v32
        _ = v36
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        for fi38 := 0; fi38 < 5; fi38++ {
            v13 = (v13 % (((v13 - v13)) | 1))
            switch (fi38 % ((127) | 1)) {
            case 1:
                v39 := GenS5{F0: (881 == uint64(v7)), F1: (v6 / (((v6 & 255)) | 1)), F2: (v30 * (v30 % ((v30) | 1)))}
                _ = v39
            case -598:
                var v40 uint16 = 204
                _ = v40
            default:
                var v41 uint64 = 42
                _ = v41
            }
            _ = len(v8)
            _ = fi38
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v6
    _ = v7
    _ = v8
}
