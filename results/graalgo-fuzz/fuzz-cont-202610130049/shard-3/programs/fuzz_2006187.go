package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int16; F1 uint; F2 int }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint16) string {
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := []int{}
        _ = fi4
        _ = v5
    }
    if false {
        p0 = p0
        for fi6 := 0; fi6 < 11; fi6++ {
            p0 += (p0 + (p0 & 1))
            fmt.Println(fi6)
            v7 := GenS1{F0: false}
            _ = fi6
            _ = v7
        }
        v8 := GenS2{F0: 492, F1: 1}
        _ = v8
    }
    if true {
        v9 := GenS1{F0: true}
        _ = v9
    } else {
        var v10 uint16 = p0
        switch ("foo bar" + ("" + "\n\t")) {
        case "\n\t":
            switch ("a" + "hello") {
            case "":
                v11 := GenS1{F0: (v10 <= (p0 / (((p0 ^ (p0 / ((v10) | 1)))) | 1)))}
                _ = v11
            case "hello":
                for fi12 := 0; fi12 < 8; fi12++ {
                    v10 = (v10 ^ (p0 + v10))
                    v10 = p0
                    for fi13 := 0; fi13 < 10; fi13++ {
                        fmt.Println(v10, fi13, fi12)
                        _ = fi13
                    }
                    _ = fi12
                }
            default:
                var v14 int64 = 0
                _ = v14
            }
        case "foo bar":
            v15 := []int{-2147483648, 9223372036854775807, -233}
            _ = v15
        case "a":
            if (4294967295 != 2147483647) {
                fmt.Println(v10, p0)
            }
        default:
            v16 := GenS1{F0: (!(!(-667 < 127)))}
            _ = v16
        }
        _ = v10
    }
    switch "hello" {
    case "\n\t":
        switch -1 {
        case 569:
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := []int{-2147483648}
                for fi19 := 0; fi19 < 18; fi19++ {
                    switch "a" {
                    case "":
                        if (!false) {
                            switch 877 {
                            case 555:
                                fmt.Println(fi17, fi19)
                            case 32768:
                                p0 = (^(p0))
                            case 9223372036854775807:
                                switch fi17 {
                                case 510:
                                    p0 = p0
                                default:
                                    v18 = append(v18, 2147483647)
                                    if (!(!(690 >= 1))) {
                                        var v20 int64 = 1
                                        var v21 uint16 = (p0 >> ((677) & 63))
                                        switch (^(fi19)) {
                                        case 32768:
                                            var v22 uint8 = 2
                                            _ = v22
                                        default:
                                            v20 -= v20
                                        }
                                        _ = v20
                                        _ = v21
                                    }
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                }
                            }
                            switch "a" {
                            case "foo bar":
                                _ = len(v18)
                            case "hello":
                                var v23 uint = 2
                                _ = v23
                            }
                            if true {
                                var v24 uint8 = 203
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                _ = v24
                            }
                        }
                    case "hello":
                        var v25 float64 = 0.1
                        _ = v25
                    case "\n\t":
                        v26 := GenS2{F0: (^(fi17)), F1: 42}
                        _ = v26
                    }
                    switch "foo bar" {
                    case "foo bar":
                        if (!false) {
                            if true {
                                if (-2.25 == (float32(-2.25) - float32(-2.25))) {
                                    v27 := GenS1{F0: ("a" <= ("hello" + (("foo bar" + "\n\t") + "foo bar")))}
                                    _ = len(v18)
                                    _ = len(v18)
                                    _ = v27
                                }
                                p0 = (p0 | p0)
                                v18 = append(v18, (fi19 / ((fi17) | 1)))
                                var v28 int16 = -32768
                                if true {
                                    if (!false) {
                                        p0 -= p0
                                    } else {
                                        fmt.Println(fi19, v28, fi19)
                                        v29 := GenS3{F0: 255}
                                        _ = v29
                                    }
                                    v30 := GenS4{F0: (false && false), F1: 128}
                                    _ = v30
                                }
                                _ = v28
                            }
                        } else {
                            v31 := GenS5{F0: 42, F1: 0, F2: fi19}
                            _ = v31
                        }
                    default:
                        switch "" {
                        case "":
                            p0 = p0
                        case "foo bar":
                            var v32 int32 = -2147483648
                            _ = v32
                        default:
                            fmt.Println(p0, p0, fi19)
                        }
                    }
                    var v33 int8 = -128
                    _ = fi19
                    _ = v33
                }
                _ = fi17
                _ = v18
            }
        case 0:
            p0 = p0
        case 490:
            for fi34 := 0; fi34 < 3; fi34++ {
                v35 := []int{fi34}
                p0 += (p0 | (p0 << ((855) & 63)))
                _ = fi34
                _ = v35
            }
        }
    }
    fmt.Println(p0)
    _ = p0
    return "foo bar"
}
func genFn1() int32 {
    var v36 int8 = 72
    v36 = 100
    for fi37 := 0; fi37 < 17; fi37++ {
        v36 = (v36 * (v36 - v36))
        fmt.Println(fi37, fi37)
        var v38 float32 = (0.1 + 1e10)
        _ = fi37
        _ = v38
    }
    _ = v36
    return 42
}

func main() {
    v1 := []uint32{225, 15, 0}
    var v2 float64 = (1.5 / (((3.14159 + (1e10 * (1e10 - 3.14159))) / 0.1) / 1.0))
    v3 := make(map[string]string)
    if (127 != 127) {
        var v39 int16 = 1
        if true {
            v39 = (v39 << ((4294967295) & 63))
            fmt.Println(v2, v2)
            var v40 uint8 = 122
            _ = v40
        }
        v2 = v2
        _ = v39
    } else {
        for fi41 := 0; fi41 < 13; fi41++ {
            for fi42 := 0; fi42 < 2; fi42++ {
                var v43 bool = false
                v44 := GenS6{F0: 42}
                _ = fi42
                _ = v43
                _ = v44
            }
            v2 = 0.1
            _ = fi41
        }
        if true {
            var v45 int32 = 1
            _ = v45
        }
        if (true || (1.5 >= 3.14159)) {
            for fi46 := 0; fi46 < 13; fi46++ {
                if false {
                    v2 = -2.25
                    v2 -= v2
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = len(v1)
                    if (true || (!((true && (1000 == 526)) && false))) {
                        switch (("\n\t" + ("foo bar" + ("" + "foo bar"))) + "\n\t") {
                        case "a":
                            v47 := GenS5{F0: -1, F1: 75, F2: 32767}
                            _ = v47
                        case "":
                            var v48 float64 = v2
                            _ = v48
                        case "foo bar":
                            var v49 uint64 = 4294967295
                            _ = v49
                        }
                    }
                }
                v2 -= v2
                _ = fi46
            }
            switch "\n\t" {
            case "hello":
                v2 += v2
            }
            var v50 int64 = 1
            _ = v50
        } else {
            v2 = v2
        }
    }
    switch "hello" {
    case "":
        var v51 uint = 255
        _ = v51
    case "a":
        v52 := v3["\n\t"]
        _ = v52
    default:
        var v53 bool = (false && (138 < 116))
        _ = v53
    }
    for fi54 := 0; fi54 < 3; fi54++ {
        if true {
            var v55 int = 102
            var v56 int32 = 1
            v2 -= 0.0
            _ = v55
            _ = v56
        }
        v2 -= 1.5
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi54
    }
    v2 = v2
    switch -811 {
    case 42:
        fmt.Println(v2)
    }
    v57 := GenS4{F0: false, F1: 128}
    if (!true) {
        v2 -= (v2 - (v2 * 1.5))
        v3["\n\t"] = "\n\t"
        var v58 float32 = ((1.5 / 1.0) - 1e10)
        _ = v58
    } else {
        v2 = v2
    }
    if false {
        v2 -= (0.0 * float64(0.0))
    }
    for fi59 := 0; fi59 < 17; fi59++ {
        var v60 uint16 = 255
        _ = fi59
        _ = v60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v57
}
