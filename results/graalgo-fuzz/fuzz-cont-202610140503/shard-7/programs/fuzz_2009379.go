package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int64; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 int32; F2 uint }
func genFn0(p0 int8) string {
    v6 := []int{}
    if (((0.1 / (0.1 - 1e10)) < 0.0) && (!true)) {
        var v7 int64 = 0
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v7
    }
    if (378 <= 42) {
        v8 := GenS1{F0: 0, F1: 172, F2: 611}
        if true {
            switch 4294967295 {
            case 127:
                if false {
                    v9 := GenS2{F0: 388}
                    p0 = p0
                    v10 := GenS3{F0: 2, F1: (21 < (p0 >> ((42) & 63)))}
                    _ = v10
                    _ = v9
                }
            }
            var v11 uint8 = 2
            for fi12 := 0; fi12 < 7; fi12++ {
                if true {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        if len(v6) > 0 {
                            v6[0] = (fi13 << ((255) & 63))
                        }
                        if (!true) {
                            p0 -= (p0 / (((p0 - (p0 / ((p0) | 1)))) | 1))
                        }
                        var v14 int = 518
                        _ = fi13
                        _ = v14
                    }
                } else {
                    if true {
                        fmt.Println(v11, p0)
                    }
                    var v15 int16 = 1000
                    if true {
                        p0 = p0
                    }
                    _ = v15
                }
                if false {
                    switch (("hello" + "\n\t") + (("\n\t" + ("a" + ("\n\t" + "a"))) + ("\n\t" + "foo bar"))) {
                    case "\n\t":
                        var v16 uint8 = (v11 / ((127) | 1))
                        _ = v16
                    }
                } else {
                    if len(v6) > 0 {
                        v6[0] = 123
                    }
                    switch (fi12 + (fi12 - (fi12 + fi12))) {
                    case 42:
                        v6 = append(v6, (fi12 ^ fi12))
                    case 65:
                        v6 = append(v6, int(v11))
                    default:
                        _ = len(v6)
                    }
                    v11 -= (v11 - v11)
                }
                p0 -= p0
                _ = fi12
            }
            _ = v11
        }
        _ = v8
    }
    fmt.Println(p0)
    v17 := GenS3{F0: 131, F1: (!false)}
    _ = p0
    _ = v17
    _ = v6
    return ""
}
func genFn1(p0 int64, p1 uint8) int64 {
    p0 = (p0 / (((p0 & p0)) | 1))
    p0 += (p0 | 0)
    var v18 float64 = 1e10
    v19 := GenS1{F0: (^(p1)), F1: 65535, F2: 1}
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return p0
}

func main() {
    v1 := []int{}
    var v2 uint = 1
    v3 := []string{}
    v4 := []string{("\n\t" + ""), ("\n\t" + (("foo bar" + ("" + "foo bar")) + "a")), (((("foo bar" + "") + "\n\t") + ("foo bar" + "foo bar")) + "\n\t"), "\n\t"}
    var v5 uint = v2
    var v20 string = "a"
    switch "a" {
    case "hello":
        v2 += (v2 / ((v5) | 1))
    }
    var v21 int32 = -29
    v21 -= v21
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 uint16 = 0
            fmt.Println(v5)
            _ = fi23
            _ = v24
        }
        if true {
            for fi25 := 0; fi25 < 14; fi25++ {
                if len(v1) > 0 {
                    v1[0] = (fi22 | fi22)
                }
                _ = fi25
            }
            for fi26 := 0; fi26 < 9; fi26++ {
                var v27 float32 = -2.25
                _ = fi26
                _ = v27
            }
            switch "\n\t" {
            case "":
                if true {
                    v28 := GenS1{F0: 255, F1: (^(v5)), F2: v21}
                    v29 := GenS4{F0: -481, F1: 42, F2: -9223372036854775808}
                    _ = v28
                    _ = v29
                } else {
                    fmt.Println(v5, v20, v20)
                    fmt.Println(v20, v5, v20)
                }
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = 42
                }
            case "\n\t":
                v3 = append(v3, "hello")
            }
        }
        var v30 int = fi22
        _ = fi22
        _ = v30
    }
    if (v20 == v20) {
        v31 := GenS2{F0: 971}
        var v32 int = -820
        switch (v32 - (v32 ^ (v32 & (v32 >> ((255) & 63))))) {
        case 0:
            if false {
                if (v20 == (("hello" + ("" + (v20 + v20))) + ("a" + ((("hello" + v20) + ((v20 + "foo bar") + v20)) + "\n\t")))) {
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + v20)
                    }
                    for fi33 := 0; fi33 < 20; fi33++ {
                        _ = len(v4)
                        v34 := GenS5{F0: 0}
                        v35 := GenS6{F0: 4294967295, F1: (v21 * (v21 % ((v21) | 1))), F2: (v5 - 4294967295)}
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
            } else {
                fmt.Println(v2)
                fmt.Println(v5, v21)
                v32 += (v32 / ((v32) | 1))
            }
        }
        _ = v31
        _ = v32
    } else {
        var v36 uint32 = 42
        for fi37 := 0; fi37 < 3; fi37++ {
            fmt.Println(v2, v21, v20)
            _ = fi37
        }
        _ = len(v4)
        _ = v36
    }
    if ((((3.14159 / 3.14159) * 1.5) + 0.0) > (3.14159 * 1e10)) {
        if false {
            var v38 int64 = 1
            switch (v20 + "hello") {
            case "a":
                if (float32(0.1) <= (1e10 - -2.25)) {
                    v39 := GenS2{F0: 32767}
                    v3 = append(v3, "\n\t")
                    _ = v39
                } else {
                    var v40 bool = true
                    v21 += int32(v5)
                    _ = v40
                }
            case "foo bar":
                v5 = (v2 - v5)
            }
            _ = v38
        }
    }
    if true {
        if len(v3) > 0 {
            v3[0] = v20
        }
        v3 = append(v3, v20)
    } else {
        v5 += (v2 % ((0) | 1))
        if (51 != 65535) {
            for fi41 := 0; fi41 < 4; fi41++ {
                var v42 int32 = v21
                _ = len(v4)
                var v43 int32 = (v42 * v21)
                _ = fi41
                _ = v42
                _ = v43
            }
        } else {
            if len(v4) > 0 {
                v4[0] = ((("" + ("hello" + ((v20 + v20) + "foo bar"))) + (("foo bar" + v20) + "a")) + v20)
            }
        }
        v3 = append(v3, v20)
    }
    if len(v1) > 0 {
        v1[0] = -398
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        var v45 uint8 = 0
        _ = fi44
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
