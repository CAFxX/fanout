package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 bool }
type GenS2 struct { F0 bool; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int; F1 bool; F2 int }
func genFn0(p0 uint32, p1 string, p2 string) uint64 {
    switch p1 {
    case "hello":
        p1 = (p1 + ((p1 + (((((p2 + ((p2 + "a") + (p1 + ((p1 + ((p1 + ("a" + "foo bar")) + ("foo bar" + (p2 + p1)))) + ("foo bar" + "\n\t"))))) + (p1 + p1)) + p1) + "") + ((p1 + ("a" + (((p2 + (p2 + p2)) + (((((p1 + "a") + p2) + (("hello" + p1) + "foo bar")) + "foo bar") + ("\n\t" + p1))) + (p2 + "\n\t")))) + (p1 + p1)))) + ("\n\t" + p1)))
    case "\n\t":
        v6 := []int{}
        _ = v6
    case "foo bar":
        v7 := []int{1}
        _ = v7
    }
    v8 := []int{}
    p2 = "\n\t"
    if (true || ((1.5 * (float64(-2.25) + (0.1 * -2.25))) < (0.0 - 3.14159))) {
        var v9 uint32 = (^(p0))
        _ = v9
    } else {
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 845
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    var v4 int = 772
    v5 := []uint32{0, 2147483647, 856}
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            fmt.Println(fi10, v4, fi11)
            v2[("foo bar" + "\n\t")] = "a"
            if true {
                var v12 uint16 = 344
                fmt.Println(fi11, v12, v1)
                v1 = (v1 / ((2) | 1))
                _ = v12
            }
            fmt.Println(fi11)
            var v13 float64 = 0.0
            var v14 int16 = -32768
            v13 = 0.0
            v13 -= 0.1
            v4 = (fi11 * 0)
            var v15 uint64 = 1
            _ = fi11
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi10
    }
    if true {
        switch "" {
        case "\n\t":
            for fi16 := 0; fi16 < 13; fi16++ {
                v1 += v1
                _ = fi16
            }
        case "":
            v4 -= v4
        default:
            var v17 float32 = float32(0.1)
            _ = v17
        }
        v4 += (v4 << ((442) & 63))
    }
    v18 := GenS1{F0: 32767, F1: 0, F2: true}
    var v19 int8 = 108
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 int32 = 1
        _ = fi20
        _ = v21
    }
    var v22 uint8 = (v1 % (((^(v1))) | 1))
    if (!true) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 = v1
        if false {
            v5 = append(v5, uint32(v1))
        } else {
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 int8 = 127
                v25 := GenS2{F0: false, F1: v19, F2: 255}
                if false {
                    switch 127 {
                    case 2147483647:
                        for fi26 := 0; fi26 < 11; fi26++ {
                            if ((v19 ^ v24) < 1) {
                                var v27 uint16 = 32767
                                _ = v27
                            }
                            v24 = v19
                            var v28 bool = true
                            v19 = v24
                            v29 := GenS2{F0: (795 > 198), F1: (v19 / ((-1) | 1)), F2: 2}
                            _ = fi26
                            _ = v28
                            _ = v29
                        }
                    case 32768:
                        switch "" {
                        case "foo bar":
                            v30 := GenS2{F0: false, F1: int8(fi23), F2: 32767}
                            _ = v30
                        case "hello":
                            if ("" <= ("" + "")) {
                                v31 := GenS1{F0: 733, F1: (v19 | v19), F2: (true && false)}
                                switch 9223372036854775807 {
                                case 2147483647:
                                    var v32 uint16 = 32767
                                    _ = v32
                                case -2147483648:
                                    v19 -= (v24 >> ((65535) & 63))
                                default:
                                    var v33 string = "hello"
                                    _ = v33
                                }
                                switch ("a" + ("a" + "a")) {
                                case "hello":
                                    for fi34 := 0; fi34 < 9; fi34++ {
                                        fmt.Println(v1, v24, v22)
                                        _ = fi34
                                    }
                                case "foo bar":
                                    fmt.Println(v24, v19, v22)
                                }
                                _ = v31
                            } else {
                                var v35 uint16 = 65535
                                var v36 int16 = -1
                                _ = len(v5)
                                _ = v35
                                _ = v36
                            }
                        case "a":
                            var v37 int32 = 0
                            _ = v37
                        default:
                            v38 := GenS3{F0: 876, F1: -312}
                            _ = v38
                        }
                    }
                    var v39 uint64 = 1
                    if len(v5) > 0 {
                        v5[0] = 2147483647
                    }
                    _ = v39
                } else {
                    var v40 int = (^(fi23))
                    for fi41 := 0; fi41 < 9; fi41++ {
                        v40 -= 128
                        var v42 int64 = 0
                        var v43 float64 = -2.25
                        if (697 <= 42) {
                            var v44 float32 = (1e10 - 0.0)
                            _ = v44
                        } else {
                            fmt.Println(v1, v22)
                            var v45 int16 = 8
                            _ = v45
                        }
                        var v46 uint = 0
                        _ = fi41
                        _ = v42
                        _ = v43
                        _ = v46
                    }
                    var v47 int64 = 0
                    _ = v40
                    _ = v47
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            var v48 uint64 = 18446744073709551615
            v49 := GenS1{F0: 2, F1: v19, F2: ((1.5 / ((0.1 * 1e10) / 1.0)) != (((((0.0 / 0.1) / 1.0) * 3.14159) * -2.25) - (0.1 + float32(-2.25))))}
            _ = v48
            _ = v49
        }
    }
    v4 = v4
    for fi50 := 0; fi50 < 14; fi50++ {
        fmt.Println(fi50)
        _ = fi50
    }
    v19 = v19
    switch "a" {
    case "a":
        v51 := GenS4{F0: 0, F1: 527}
        _ = v51
    case "":
        switch "a" {
        case "a":
            var v52 uint = 18446744073709551615
            _ = v52
        }
    case "\n\t":
        switch "hello" {
        case "foo bar":
            switch "hello" {
            case "foo bar":
                fmt.Println(v19)
            default:
                fmt.Println(v19, v4)
            }
        case "a":
            v5 = append(v5, 4294967295)
        case "":
            v53 := GenS2{F0: ("foo bar" > ""), F1: (v19 - v19), F2: 403}
            _ = v53
        default:
            v54 := GenS3{F0: -1, F1: -9223372036854775808}
            _ = v54
        }
    default:
        var v55 uint = 872
        _ = v55
    }
    switch (^(v4)) {
    case 127:
        switch "a" {
        case "\n\t":
            fmt.Println(v19, v19)
        case "foo bar":
            v22 = 2
        }
    case 128:
        var v56 int = v4
        _ = v56
    case 987:
        var v57 float64 = (float64(float32(0.1)) - -2.25)
        _ = v57
    default:
        fmt.Println(v22, v19)
    }
    v22 += 128
    for fi58 := 0; fi58 < 10; fi58++ {
        if true {
            switch ((("a" + ("foo bar" + "\n\t")) + "foo bar") + "a") {
            case "hello":
                v4 -= fi58
            case "a":
                v59 := GenS2{F0: (!(0 > 4294967295)), F1: 65, F2: 32767}
                _ = v59
            case "":
                v60 := GenS1{F0: 1, F1: v19, F2: false}
                _ = v60
            }
        }
        var v61 int = 32767
        v22 -= v1
        _ = fi58
        _ = v61
    }
    v62 := GenS5{F0: v4, F1: false, F2: (v4 & (v4 / ((-847) | 1)))}
    for fi63 := 0; fi63 < 7; fi63++ {
        for fi64 := 0; fi64 < 12; fi64++ {
            var v65 float64 = (float64(0.1) - 1e10)
            switch (v4 | fi64) {
            case 32768:
                fmt.Println(v1, v65, fi63)
            case -849:
                var v66 int64 = -579
                _ = v66
            case 171:
                switch ("\n\t" + "a") {
                case "foo bar":
                    fmt.Println(fi63, v65, v4)
                case "":
                    switch (fi64 >> ((2) & 63)) {
                    case 9223372036854775807:
                        if (!false) {
                            delete(v2, "\n\t")
                            v4 = (fi63 << ((4294967295) & 63))
                        }
                    default:
                        if len(v5) > 0 {
                            v5[0] = 2
                        }
                    }
                default:
                    v19 = v19
                }
            }
            var v67 int16 = 0
            _ = fi64
            _ = v65
            _ = v67
        }
        var v68 int64 = -1
        _ = fi63
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v62
}
