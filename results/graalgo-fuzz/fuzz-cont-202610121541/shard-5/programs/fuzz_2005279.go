package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint; F1 uint16 }
type GenS10 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS11 struct { F0 int32; F1 bool; F2 uint32 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int; F1 int16; F2 int8 }

func main() {
    v1 := []int{}
    v2 := make(map[int]uint32)
    var v3 int = -1
    v4 := make(map[uint32]uint32)
    v5 := GenS1{F0: 42}
    v6 := GenS2{F0: (^(v3)), F1: -98, F2: 696}
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(v3)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3 = fi8
            _ = fi8
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS2{F0: (v3 * 127), F1: -669, F2: 587}
            v11 := GenS3{F0: 64, F1: -836, F2: 0}
            var v12 uint16 = 255
            v13 := GenS4{F0: -128, F1: (v12 / ((v12) | 1)), F2: -313}
            _ = len(v1)
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
            _ = v13
        }
        switch ((("" + ("foo bar" + ("" + "foo bar"))) + "\n\t") + (("\n\t" + "hello") + "")) {
        case "hello":
            v14 := GenS5{F0: 641}
            _ = v14
        case "foo bar":
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 float32 = 0.0
                switch "a" {
                case "":
                    v17 := GenS3{F0: 127, F1: 326, F2: 0}
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
        case "\n\t":
            v18 := GenS6{F0: 414}
            _ = v18
        }
        _ = fi7
    }
    v19 := GenS7{F0: -128}
    var v20 int32 = -840
    if false {
        var v21 uint16 = 732
        _ = v21
    } else {
        if (true && (!true)) {
            var v22 string = "\n\t"
            fmt.Println(v3, v22)
            _ = v22
        } else {
            if (32767 < 857) {
                v23 := GenS5{F0: 65535}
                v1 = append(v1, (-(v3)))
                _ = v23
            }
            var v24 int32 = 1
            var v25 uint16 = 32767
            _ = v24
            _ = v25
        }
        v20 -= (^(v20))
        switch "hello" {
        case "a":
            switch (v3 & (v3 << ((2) & 63))) {
            case 218:
                var v26 string = "foo bar"
                _ = v26
            case 2147483647:
                if len(v1) > 0 {
                    v1[0] = v3
                }
            case 9223372036854775807:
                if true {
                    if (((554 != 714) && false) && false) {
                        v3 = (v3 + 1)
                        v20 = v20
                    } else {
                        var v27 int64 = 0
                        v28 := GenS8{F0: 65535}
                        _ = v27
                        _ = v28
                    }
                    switch v3 {
                    case -2147483648:
                        var v29 float32 = (float32(3.14159) - (float32(float64(1e10)) - float32(1e10)))
                        _ = v29
                    case 0:
                        v3 += (^(v3))
                    case -1:
                        var v30 uint64 = 4294967295
                        _ = v30
                    default:
                        var v31 int32 = (v20 / (((v20 * (v20 ^ (v20 & 0)))) | 1))
                        _ = v31
                    }
                    v32 := v4[65535]
                    _ = v32
                }
            default:
                v20 += v20
            }
        }
    }
    var v33 int16 = 42
    v34 := GenS1{F0: 301}
    fmt.Println(v20)
    for fi35 := 0; fi35 < 16; fi35++ {
        var v36 int32 = 0
        v37 := GenS9{F0: 65535, F1: 32767}
        if (!false) {
            var v38 uint32 = 775
            switch fi35 {
            case -467:
                var v39 int32 = v20
                _ = v39
            case 9223372036854775807:
                v40 := GenS4{F0: 127, F1: 487, F2: v36}
                _ = v40
            case -2147483648:
                v20 += (v36 ^ (v20 / (((v20 + (v36 >> ((358) & 63)))) | 1)))
            }
            _ = v38
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    var v41 int32 = -658
    v42 := GenS10{F0: 42, F1: 42, F2: 38}
    if true {
        fmt.Println(v20, v3, v33)
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        switch (fi43 | fi43) {
        case 128:
            v20 = v20
        case 2147483647:
            if false {
                v44 := GenS11{F0: (v41 & v20), F1: false, F2: 967}
                v3 -= (fi43 ^ -628)
                var v45 int8 = 1
                _ = v44
                _ = v45
            } else {
                if true {
                    fmt.Println(v3, v33)
                    v20 = 665
                    switch (("foo bar" + ("hello" + "")) + "hello") {
                    case "hello":
                        var v46 uint8 = 255
                        _ = v46
                    case "\n\t":
                        v3 += fi43
                    case "":
                        fmt.Println(fi43)
                    }
                } else {
                    fmt.Println(v41, v3, fi43)
                    switch v3 {
                    case 9223372036854775807:
                        v2[-2147483648] = 909
                    case 32768:
                        v3 -= (^(v3))
                    case 978:
                        _ = len(v1)
                    }
                    if (-2.25 == 3.14159) {
                        var v47 int16 = (v33 & v33)
                        var v48 uint = 18446744073709551615
                        v49 := GenS12{F0: false}
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                }
            }
        }
        _ = fi43
    }
    v50 := GenS13{F0: (v3 % ((v3) | 1)), F1: int16(v3), F2: 125}
    switch "" {
    case "":
        var v51 float32 = (0.1 - 3.14159)
        _ = v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v50
    _ = v6
}
