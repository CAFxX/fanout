package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 uint; F1 uint64 }

func main() {
    v1 := []uint32{2147483648}
    v2 := []uint32{2147483648, 1}
    var v3 int = 9223372036854775807
    var v4 int16 = 42
    fmt.Println(v3)
    fmt.Println(v4, v3)
    switch (v3 + (v3 >> ((65535) & 63))) {
    case -9223372036854775808:
        v3 -= v3
    }
    var v5 int64 = 0
    v5 -= v5
    v6 := GenS1{F0: 42, F1: 128}
    var v7 int64 = (v5 << ((1) & 63))
    _ = len(v2)
    if (0 <= 286) {
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 uint16 = 413
                var v11 int64 = (v7 / ((0) | 1))
                _ = fi9
                _ = v10
                _ = v11
            }
            v4 = -1
            _ = fi8
        }
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        if true {
            if false {
                var v13 uint64 = 698
                _ = v13
            } else {
                fmt.Println(v4, v5, v7)
                var v14 float32 = (float32(float64(1e10)) * 0.1)
                if (140 <= 2147483648) {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        _ = len(v2)
                        switch ("hello" + (("" + "hello") + ("" + "hello"))) {
                        case "hello":
                            switch ("" + "\n\t") {
                            case "foo bar":
                                switch "hello" {
                                case "foo bar":
                                    v14 += v14
                                }
                            case "a":
                                fmt.Println(fi15, v14, v14)
                            case "\n\t":
                                v16 := GenS1{F0: 1, F1: 128}
                                _ = v16
                            default:
                                switch ("hello" + "") {
                                case "foo bar":
                                    v4 = -1
                                }
                            }
                            var v17 uint16 = 2
                            _ = v17
                        case "foo bar":
                            v4 += v4
                        case "\n\t":
                            v4 += v4
                        }
                        var v18 uint8 = 42
                        _ = fi15
                        _ = v18
                    }
                }
                _ = v14
            }
        }
        if false {
            if false {
                v1 = append(v1, 2147483647)
                if false {
                    if (-4 != 0) {
                        v19 := GenS1{F0: 237, F1: 128}
                        v20 := GenS1{F0: 2, F1: 0}
                        for fi21 := 0; fi21 < 4; fi21++ {
                            v22 := GenS1{F0: 127, F1: 100}
                            _ = fi21
                            _ = v22
                        }
                        _ = v19
                        _ = v20
                    } else {
                        if false {
                            v4 = v4
                            for fi23 := 0; fi23 < 11; fi23++ {
                                fmt.Println(fi23, fi23, v4)
                                _ = fi23
                            }
                        } else {
                            var v24 uint32 = 65535
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 1; fi25++ {
                            switch fi25 {
                            case 128:
                                v5 = 9223372036854775807
                            case 42:
                                v26 := GenS2{F0: -101, F1: "a"}
                                var v27 int16 = (v4 + v4)
                                _ = v26
                                _ = v27
                            case -89:
                                var v28 uint16 = 32767
                                _ = v28
                            }
                            switch -568 {
                            case -138:
                                if (true || (2 < 168)) {
                                    if ((v5 % (((v5 ^ (v5 / (((v7 ^ 1)) | 1)))) | 1)) > v5) {
                                        v7 = 9223372036854775807
                                        v3 = v3
                                        v4 -= (v4 ^ (-(v4)))
                                    }
                                } else {
                                    fmt.Println(fi12, v4)
                                    switch ("\n\t" + "hello") {
                                    case "hello":
                                        _ = len(v1)
                                    }
                                    v3 -= (fi12 << ((255) & 63))
                                }
                            case -1:
                                var v29 uint16 = 1
                                _ = v29
                            default:
                                switch 32768 {
                                case 32768:
                                    v30 := GenS3{F0: int8(fi25), F1: 765}
                                    _ = v30
                                case -552:
                                    switch (fi25 % ((-1) | 1)) {
                                    case 0:
                                        var v31 int16 = v4
                                        _ = v31
                                    }
                                case 9223372036854775807:
                                    var v32 uint8 = 0
                                    _ = v32
                                default:
                                    var v33 int64 = 1
                                    v34 := GenS4{F0: 0, F1: 127}
                                    switch "\n\t" {
                                    case "a":
                                        v4 = 32767
                                    case "hello":
                                        switch ("hello" + (("" + "") + "")) {
                                        case "\n\t":
                                            fmt.Println(fi12, v5, fi25)
                                        case "":
                                            v35 := GenS5{F0: -65, F1: false}
                                            fmt.Println(v3, v3, v4)
                                            fmt.Println(fi25)
                                            _ = v35
                                        default:
                                            if ((955 == 959) || ((412 < 32) && false)) {
                                                if true {
                                                    if true {
                                                        v36 := GenS2{F0: 80, F1: "hello"}
                                                        v2 = append(v2, 42)
                                                        var v37 int = fi25
                                                        switch ("a" + (("" + "\n\t") + "")) {
                                                        case "":
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v38 int64 = (-(v5))
                                                                _ = v38
                                                            }
                                                        case "a":
                                                            if (false || (!false)) {
                                                                var v39 int = 128
                                                                v40 := GenS1{F0: 127, F1: 255}
                                                                _ = v39
                                                                _ = v40
                                                            }
                                                        }
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                }
                                                v5 += (v33 / ((v5) | 1))
                                            }
                                        }
                                    default:
                                        if (!true) {
                                            var v41 uint64 = 0
                                            if (42 != v4) {
                                                fmt.Println(v7)
                                            }
                                            _ = v41
                                        } else {
                                            fmt.Println(v4, fi12, v5)
                                        }
                                    }
                                    _ = v33
                                    _ = v34
                                }
                            }
                            _ = fi25
                        }
                    }
                } else {
                    var v42 uint8 = 0
                    if false {
                        v3 += fi12
                    }
                    _ = v42
                }
                v43 := GenS6{F0: 294, F1: 775}
                _ = v43
            }
            _ = len(v1)
        }
        if len(v2) > 0 {
            v2[0] = 127
        }
        _ = fi12
    }
    var v44 uint8 = 42
    var v45 float32 = -2.25
    var v46 float64 = (3.14159 * -2.25)
    for fi47 := 0; fi47 < 19; fi47++ {
        var v48 uint = 979
        _ = fi47
        _ = v48
    }
    var v49 float32 = v45
    if ("hello" != "\n\t") {
        switch (v3 - (v3 & v3)) {
        case -37:
            v50 := GenS3{F0: 100, F1: 127}
            _ = v50
        case -809:
            for fi51 := 0; fi51 < 11; fi51++ {
                for fi52 := 0; fi52 < 17; fi52++ {
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v49 -= v49
                        v53 := GenS2{F0: -128, F1: (("" + "foo bar") + ("\n\t" + ("hello" + ("\n\t" + ("hello" + (("\n\t" + "") + "foo bar"))))))}
                        switch fi52 {
                        case 42:
                            v54 := GenS1{F0: (v44 - (^(v44))), F1: 127}
                            _ = v54
                        case -484:
                            v3 += v3
                        default:
                            if true {
                                v3 += 2147483647
                                v5 = -9223372036854775808
                                fmt.Println(v5)
                            }
                        }
                        _ = v53
                    }
                    _ = fi52
                }
                _ = fi51
            }
        default:
            _ = len(v2)
        }
        v3 = 2147483647
        if len(v1) > 0 {
            v1[0] = 0
        }
    }
    fmt.Println(v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
