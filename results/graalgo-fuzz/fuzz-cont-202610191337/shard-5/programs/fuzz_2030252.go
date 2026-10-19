package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int8; F2 uint32 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32; F1 int32; F2 uint8 }

func main() {
    var v1 bool = true
    v2 := []uint32{2147483648}
    var v3 int32 = -1
    v4 := []string{}
    v5 := GenS1{F0: 65535}
    v6 := GenS2{F0: 42, F1: 916}
    var v7 float64 = (0.1 / -2.25)
    _ = len(v2)
    switch -279 {
    case 9223372036854775807:
        v8 := GenS3{F0: 2147483648}
        _ = v8
    case 32768:
        for fi9 := 0; fi9 < 16; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                if (true && true) {
                    v3 += v3
                    _ = len(v4)
                    var v11 string = "foo bar"
                    _ = v11
                } else {
                    switch ("foo bar" + "a") {
                    case "":
                        v4 = append(v4, "a")
                    case "\n\t":
                        v1 = v1
                        if false {
                            v2 = append(v2, 740)
                        } else {
                            v12 := GenS2{F0: 42, F1: (v3 << ((42) & 63))}
                            var v13 int = (fi9 - -1)
                            v14 := GenS4{F0: -1, F1: -128, F2: 1}
                            fmt.Println(fi9, v7, fi9)
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        var v15 string = "a"
                        _ = v15
                    case "hello":
                        switch (fi9 - fi9) {
                        case 0:
                            v1 = (2147483647 >= 392)
                        case 4294967295:
                            switch ("a" + "") {
                            case "foo bar":
                                switch ("hello" + ((("a" + "") + (("" + "a") + ("hello" + "\n\t"))) + "\n\t")) {
                                case "":
                                    v1 = (!(!v1))
                                }
                            }
                        case -9223372036854775808:
                            v16 := GenS5{F0: 32767, F1: 9}
                            _ = v16
                        }
                    }
                    if v1 {
                        if (0 > 881) {
                            switch "\n\t" {
                            case "foo bar":
                                switch "" {
                                case "hello":
                                    v3 = -668
                                case "\n\t":
                                    v7 = 3.14159
                                case "":
                                    var v17 uint32 = 0
                                    _ = v17
                                }
                            case "\n\t":
                                if true {
                                    var v18 int16 = -1
                                    _ = len(v2)
                                    v19 := GenS6{F0: "hello"}
                                    _ = v18
                                    _ = v19
                                }
                            case "hello":
                                if (238 <= 128) {
                                    v4 = append(v4, "hello")
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = 391
                                    }
                                    fmt.Println(v1, v1, fi9)
                                }
                            default:
                                v3 += (v3 - (v3 & (v3 % ((v3) | 1))))
                            }
                            if (42 < 128) {
                                if (true || (!(42 > 0))) {
                                    var v20 int32 = 42
                                    var v21 int32 = v20
                                    v1 = (130 < 0)
                                    if true {
                                        switch (fi10 ^ (fi10 + (fi10 % (((fi9 ^ (fi10 | fi9))) | 1)))) {
                                        case 8:
                                            var v22 int16 = -85
                                            _ = v22
                                        case 4294967295:
                                            if (v1 || v1) {
                                                var v23 float64 = 0.0
                                                switch (fi10 * (fi9 - fi9)) {
                                                case 127:
                                                    var v24 uint16 = 42
                                                    _ = v24
                                                default:
                                                    var v25 int = fi10
                                                    _ = v25
                                                }
                                                var v26 uint = 65535
                                                _ = v23
                                                _ = v26
                                            } else {
                                                switch (fi9 >> ((4294967295) & 63)) {
                                                case 42:
                                                    var v27 uint64 = 42
                                                    _ = v27
                                                }
                                                v28 := GenS2{F0: 127, F1: (v3 / ((42) | 1))}
                                                var v29 uint = 65535
                                                _ = v28
                                                _ = v29
                                            }
                                        }
                                        v1 = (18446744073709551615 == 2)
                                        var v30 uint64 = 2
                                        _ = v30
                                    }
                                    _ = v20
                                    _ = v21
                                }
                                v31 := GenS7{F0: (v3 * v3), F1: -2147483648, F2: 87}
                                _ = v31
                            } else {
                                v32 := GenS3{F0: 65535}
                                v2 = append(v2, 884)
                                _ = v32
                            }
                            v7 += 0.0
                        }
                        if v1 {
                            if v1 {
                                var v33 bool = v1
                                var v34 uint8 = 127
                                switch (("" + "a") + "foo bar") {
                                case "hello":
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                default:
                                    v33 = v1
                                }
                                _ = v33
                                _ = v34
                            } else {
                                if (0 > -32768) {
                                    fmt.Println(fi9, v7)
                                } else {
                                    v3 -= -1
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if v1 {
                                        v1 = (1 != 0)
                                    } else {
                                        v3 = -643
                                    }
                                }
                                fmt.Println(fi9, v7, fi10)
                            }
                            var v35 string = ("foo bar" + "foo bar")
                            _ = v35
                        } else {
                            var v36 uint32 = 72
                            _ = v36
                        }
                        switch fi10 {
                        case 0:
                            if (2 < 2) {
                                var v37 int = (fi9 * (fi9 % ((fi9) | 1)))
                                var v38 uint8 = 127
                                _ = v37
                                _ = v38
                            } else {
                                var v39 float64 = v7
                                _ = v39
                            }
                        }
                        fmt.Println(fi10)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch 1 {
                        case 42:
                            var v40 int64 = 9223372036854775807
                            _ = v40
                        case 4294967295:
                            var v41 float64 = float64(3.14159)
                            _ = v41
                        }
                    }
                    var v42 uint16 = 65535
                    var v43 float64 = -2.25
                    v3 += (v3 % ((v3) | 1))
                    _ = v42
                    _ = v43
                }
                v44 := GenS6{F0: "\n\t"}
                fmt.Println(v3, v3, v3)
                _ = fi10
                _ = v44
            }
            var v45 uint32 = 1
            _ = fi9
            _ = v45
        }
    case -709:
        for fi46 := 0; fi46 < 1; fi46++ {
            for fi47 := 0; fi47 < 9; fi47++ {
                fmt.Println(v3, v3, fi47)
                switch "foo bar" {
                case "foo bar":
                    v7 += (v7 - 0.0)
                case "\n\t":
                    v48 := GenS6{F0: ("\n\t" + "\n\t")}
                    _ = v48
                }
                _ = fi47
            }
            if false {
                v4 = append(v4, "\n\t")
            } else {
                for fi49 := 0; fi49 < 16; fi49++ {
                    if (!true) {
                        switch 247 {
                        case -2147483648:
                            v1 = v1
                        case 9223372036854775807:
                            fmt.Println(fi46, v3)
                        case -9223372036854775808:
                            var v50 int8 = 0
                            _ = v50
                        }
                        var v51 uint = 255
                        _ = v51
                    }
                    _ = fi49
                }
                v1 = ((-1 >= -9223372036854775808) || (v1 || false))
                v1 = false
            }
            var v52 int8 = 98
            _ = fi46
            _ = v52
        }
    default:
        var v53 int8 = -27
        _ = v53
    }
    var v54 uint16 = 2
    v3 -= v3
    if len(v2) > 0 {
        v2[0] = 0
    }
    for fi55 := 0; fi55 < 14; fi55++ {
        var v56 string = "a"
        if ((fi55 * (fi55 | (fi55 - fi55))) < fi55) {
            fmt.Println(fi55, v54, v1)
        }
        var v57 int16 = -1
        _ = fi55
        _ = v56
        _ = v57
    }
    fmt.Println(v54, v7)
    for fi58 := 0; fi58 < 8; fi58++ {
        v7 -= 0.0
        var v59 uint16 = 2
        var v60 float32 = 0.1
        _ = fi58
        _ = v59
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v54:", v54)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v7
}
