package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 uint64 }

func main() {
    var v1 float32 = 1e10
    var v2 float32 = 0.1
    var v3 bool = (0 < 1)
    if v3 {
        if false {
            v4 := GenS1{F0: -124}
            v1 = ((1.5 - (v1 / 3.14159)) * (v2 + (v1 / -2.25)))
            _ = v4
        }
        var v5 uint = 220
        v6 := GenS1{F0: 127}
        _ = v5
        _ = v6
    } else {
        var v7 int32 = -88
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 float64 = -2.25
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 80}
    var v12 int32 = 42
    v13 := GenS2{F0: 305, F1: 200}
    v14 := GenS1{F0: 124}
    switch ("hello" + ((("foo bar" + "a") + "") + "\n\t")) {
    case "":
        if (0.0 != -2.25) {
            v15 := GenS3{F0: v3, F1: 32767}
            v16 := []int{128, -2147483648}
            var v17 uint8 = 127
            _ = v15
            _ = v16
            _ = v17
        }
    case "hello":
        switch 42 {
        case -921:
            switch "" {
            case "foo bar":
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 uint16 = 2
                    _ = fi18
                    _ = v19
                }
            default:
                for fi20 := 0; fi20 < 19; fi20++ {
                    switch ("" + "\n\t") {
                    case "hello":
                        var v21 string = "hello"
                        _ = v21
                    case "\n\t":
                        v12 += v12
                    case "":
                        v2 = v1
                    default:
                        var v22 uint16 = 1
                        _ = v22
                    }
                    if v3 {
                        v12 -= 1
                        if ("hello" >= "hello") {
                            if v3 {
                                v23 := GenS3{F0: true, F1: 1000}
                                if (!(!(0 <= 42))) {
                                    var v24 int32 = v12
                                    _ = v24
                                }
                                fmt.Println(v12)
                                _ = v23
                            }
                            for fi25 := 0; fi25 < 18; fi25++ {
                                if (!(!(!v3))) {
                                    switch ("" + ("foo bar" + "hello")) {
                                    case "":
                                        switch "a" {
                                        case "foo bar":
                                            v26 := []int{fi20}
                                            _ = v26
                                        case "hello":
                                            v27 := GenS2{F0: 155, F1: 7}
                                            _ = v27
                                        }
                                    case "a":
                                        if (42 == 128) {
                                            v28 := []int{fi20}
                                            v29 := GenS3{F0: (0 >= 42), F1: 1000}
                                            if v3 {
                                                v30 := GenS3{F0: v3, F1: -1}
                                                _ = v30
                                            }
                                            _ = v28
                                            _ = v29
                                        }
                                    default:
                                        var v31 uint64 = 1
                                        _ = v31
                                    }
                                    var v32 uint64 = 760
                                    _ = v32
                                }
                                switch (fi20 << ((677) & 63)) {
                                case 9223372036854775807:
                                    var v33 int16 = 1000
                                    _ = v33
                                case 127:
                                    var v34 string = (("" + "hello") + "foo bar")
                                    _ = v34
                                case -946:
                                    var v35 float64 = (1e10 - 3.14159)
                                    _ = v35
                                default:
                                    v2 = v1
                                }
                                _ = fi25
                            }
                        }
                        fmt.Println(v2, v2)
                    }
                    if v3 {
                        if v3 {
                            v36 := []int{-243, fi20, fi20}
                            if (73 >= 42) {
                                v2 = (v1 / (v2 * ((((v1 + v1) * v1) + float32((float64(v2) * float64(float32(1.5))))) / v2)))
                                switch (("hello" + "") + "hello") {
                                case "a":
                                    switch -143 {
                                    case -9223372036854775808:
                                        v3 = v3
                                    case 102:
                                        v37 := GenS1{F0: 42}
                                        _ = v37
                                    case 4294967295:
                                        v1 += v1
                                    }
                                }
                                if ("a" > ((("hello" + "\n\t") + (("hello" + "a") + ("foo bar" + ""))) + "\n\t")) {
                                    var v38 int8 = 1
                                    switch ("" + "foo bar") {
                                    case "":
                                        for fi39 := 0; fi39 < 1; fi39++ {
                                            fmt.Println(fi20)
                                            _ = fi39
                                        }
                                    case "a":
                                        var v40 int64 = -1
                                        _ = v40
                                    case "foo bar":
                                        if v3 {
                                            switch "" {
                                            case "a":
                                                v41 := GenS1{F0: v38}
                                                _ = v41
                                            case "foo bar":
                                                v42 := GenS3{F0: false, F1: 32767}
                                                _ = v42
                                            default:
                                                v43 := GenS3{F0: (0 <= 128), F1: 1000}
                                                _ = v43
                                            }
                                            fmt.Println(v1, v2, v1)
                                        }
                                    }
                                    _ = v38
                                } else {
                                    var v44 int = 4294967295
                                    if len(v36) >= 2 {
                                        v36 = v36[1:len(v36)]
                                    }
                                    for fi45 := 0; fi45 < 19; fi45++ {
                                        var v46 uint64 = 4294967295
                                        fmt.Println(v1)
                                        _ = fi45
                                        _ = v46
                                    }
                                    _ = v44
                                }
                            }
                            var v47 bool = v3
                            _ = v36
                            _ = v47
                        } else {
                            if v3 {
                                switch (fi20 | 2147483647) {
                                case 2147483647:
                                    v3 = v3
                                }
                            }
                            var v48 uint64 = 4294967295
                            var v49 int64 = 986
                            _ = v48
                            _ = v49
                        }
                    }
                    _ = fi20
                }
            }
        case 42:
            var v50 int8 = 127
            _ = v50
        default:
            v1 = (0.0 / v1)
        }
    default:
        v51 := []int{-2147483648}
        _ = v51
    }
    if (1 <= -9223372036854775808) {
        for fi52 := 0; fi52 < 10; fi52++ {
            switch 9223372036854775807 {
            case 0:
                var v53 int16 = 478
                _ = v53
            case -1:
                v54 := GenS4{F0: -1, F1: 2}
                _ = v54
            default:
                v55 := []int{42, (-(fi52))}
                _ = v55
            }
            _ = fi52
        }
    }
    v2 = (v1 / v2)
    fmt.Println(v2, v3, v1)
    if true {
        v2 = (0.1 - 0.0)
    }
    if false {
        if false {
            fmt.Println(v1, v2, v12)
            var v56 uint8 = 127
            v56 += 168
            _ = v56
        }
        v57 := GenS5{F0: 512}
        _ = v57
    }
    var v58 string = (("\n\t" + (("foo bar" + ("hello" + "\n\t")) + "")) + "foo bar")
    fmt.Println(v1, v1, v2)
    if (v12 >= v12) {
        v59 := []int{}
        v2 -= v1
        v60 := GenS2{F0: 939, F1: 112}
        _ = v59
        _ = v60
    }
    switch (v58 + (("a" + v58) + "foo bar")) {
    case "foo bar":
        v61 := GenS2{F0: 348, F1: 1}
        _ = v61
    case "":
        for fi62 := 0; fi62 < 2; fi62++ {
            v63 := GenS6{F0: -431, F1: 4294967295}
            _ = fi62
            _ = v63
        }
    case "\n\t":
        v3 = v3
    default:
        v12 += v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v58
}
