package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int32 = -509
    v2 := []uint32{2, 65535}
    v3 := make(map[int]int)
    v4 := GenS1{F0: 4294967295}
    v5 := GenS1{F0: 791}
    switch ("\n\t" + "a") {
    case "hello":
        var v6 uint32 = 290
        _ = v6
    case "":
        var v7 uint = 2
        _ = v7
    case "\n\t":
        v1 -= v1
    default:
        var v8 bool = (42 > 0)
        _ = v8
    }
    v1 = v1
    if len(v2) > 0 {
        v2[0] = 450
    }
    if (-2147483648 == v1) {
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 14; fi10++ {
                v1 += (v1 - -11)
                _ = fi10
            }
            var v11 int64 = -685
            var v12 int8 = int8(fi9)
            _ = fi9
            _ = v11
            _ = v12
        }
        if len(v2) > 0 {
            v2[0] = 986
        }
        if true {
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := GenS1{F0: 1}
                switch (fi13 >> ((255) & 63)) {
                case -524:
                    if (!((!true) || (1 <= 42))) {
                        var v15 uint16 = 65535
                        _ = v15
                    }
                case -1:
                    var v16 int32 = 2147483647
                    _ = v16
                default:
                    if false {
                        var v17 string = (("" + "foo bar") + "foo bar")
                        v18 := GenS2{F0: false, F1: (fi13 * 32767), F2: 300}
                        _ = v17
                        _ = v18
                    }
                }
                switch ("foo bar" + "a") {
                case "hello":
                    v1 = (v1 * v1)
                case "foo bar":
                    var v19 uint32 = 0
                    _ = v19
                }
                _ = fi13
                _ = v14
            }
        }
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v21 := GenS3{F0: 65535, F1: 42}
        v1 -= (v1 + (v1 ^ (v1 | -592)))
        switch (("foo bar" + "foo bar") + "hello") {
        case "\n\t":
            if true {
                if (fi20 == (fi20 | 1)) {
                    v1 = v1
                } else {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v1 -= v1
                        _ = fi22
                    }
                }
                v23 := GenS4{F0: 1}
                v1 -= v1
                _ = v23
            }
        case "a":
            v1 += 42
        default:
            if (1 < 42) {
                if false {
                    fmt.Println(v1, fi20)
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v25 := GenS2{F0: (0 >= 9223372036854775807), F1: fi24, F2: 42}
                        _ = fi24
                        _ = v25
                    }
                    v1 += v1
                } else {
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v1 -= 2147483647
                        _ = fi26
                    }
                }
                var v27 uint64 = 1
                _ = v27
            } else {
                if true {
                    v1 += (v1 << ((1) & 63))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (652 < -1) {
                        switch "" {
                        case "foo bar":
                            var v28 float32 = (1.5 + 0.1)
                            _ = v28
                        case "a":
                            v3[fi20] = -232
                        case "":
                            var v29 string = (("" + ("a" + "hello")) + "a")
                            _ = v29
                        }
                        switch "\n\t" {
                        case "\n\t":
                            var v30 uint32 = 805
                            _ = v30
                        case "hello":
                            for fi31 := 0; fi31 < 9; fi31++ {
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = 6
                                    }
                                }
                                var v32 int = (fi31 ^ fi31)
                                var v33 float32 = 1.5
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        default:
                            var v34 bool = false
                            _ = v34
                        }
                    }
                } else {
                    for fi35 := 0; fi35 < 9; fi35++ {
                        fmt.Println(fi20)
                        var v36 uint8 = 127
                        switch (fi35 & (fi20 - (-(fi35)))) {
                        case -1:
                            var v37 uint32 = 890
                            _ = v37
                        }
                        _ = fi35
                        _ = v36
                    }
                }
                var v38 uint = 4294967295
                _ = v38
            }
        }
        _ = fi20
        _ = v21
    }
    var v39 string = ("\n\t" + "")
    switch v39 {
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        var v40 float64 = (float64(0.0) * (-2.25 - 0.0))
        _ = v40
    case "a":
        v3[9223372036854775807] = 4294967295
    }
    var v41 bool = (!true)
    fmt.Println(v41)
    switch v39 {
    case "foo bar":
        if false {
            if false {
                v1 = 0
            }
        }
    case "\n\t":
        v42 := v3[4294967295]
        _ = v42
    case "a":
        v43 := GenS4{F0: 2}
        _ = v43
    default:
        v2 = append(v2, 818)
    }
    v44 := GenS2{F0: v41, F1: 127, F2: 42}
    v45 := GenS3{F0: 255, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v44
    _ = v45
    _ = v5
}
