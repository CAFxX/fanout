package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int64; F1 int16; F2 int }
func genFn0(p0 uint, p1 int64) string {
    v7 := []int{127, 32768, 4294967295}
    v7 = append(v7, 127)
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    return ("a" + ("\n\t" + ""))
}

func main() {
    var v1 bool = (float32(0.1) <= (-2.25 * ((-2.25 / (((-2.25 - 0.0) / -2.25) * (0.0 + float32(0.1)))) - 0.1)))
    v2 := make(map[string]int)
    var v3 int16 = 0
    v4 := []uint32{4294967295, 2147483648, 2}
    v5 := []string{("\n\t" + (("\n\t" + "foo bar") + ("" + ("foo bar" + ("a" + ("" + (("a" + (("hello" + "\n\t") + "hello")) + ((("foo bar" + "\n\t") + ("" + "\n\t")) + "foo bar")))))))), "a", ""}
    var v6 int = -129
    var v8 uint64 = 1
    fmt.Println(v3, v6, v6)
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := v2["a"]
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v11 uint = 18446744073709551615
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        fmt.Println(v8, fi12)
        var v13 uint8 = 2
        _ = fi12
        _ = v13
    }
    switch ("hello" + "\n\t") {
    case "":
        for fi14 := 0; fi14 < 10; fi14++ {
            v4 = append(v4, 391)
            _ = fi14
        }
    case "a":
        v15 := GenS1{F0: ("" + (("foo bar" + "hello") + "a")), F1: 948}
        _ = v15
    case "\n\t":
        switch "\n\t" {
        case "":
            var v16 int64 = 0
            _ = v16
        default:
            v3 = v3
        }
    }
    v3 += 194
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v17 uint16 = 65535
    switch ((("foo bar" + "foo bar") + (((("\n\t" + "") + "\n\t") + "\n\t") + (("" + "foo bar") + "\n\t"))) + ((("" + ((("" + ("\n\t" + "")) + (("a" + "") + "\n\t")) + "foo bar")) + "") + "hello")) {
    case "\n\t":
        switch ("foo bar" + "foo bar") {
        case "a":
            var v18 float32 = -2.25
            _ = v18
        case "foo bar":
            var v19 float32 = 0.1
            _ = v19
        default:
            switch "\n\t" {
            case "foo bar":
                for fi20 := 0; fi20 < 20; fi20++ {
                    v5 = append(v5, "foo bar")
                    _ = fi20
                }
            case "a":
                v2[("" + "foo bar")] = (v6 ^ -811)
            }
        }
    case "foo bar":
        switch "foo bar" {
        case "hello":
            var v21 int = (v6 & (v6 | (v6 % (((v6 * 32767)) | 1))))
            _ = v21
        default:
            if ((2 < v17) && v1) {
                var v22 int32 = 1
                _ = v22
            }
        }
    case "":
        var v23 uint8 = 127
        _ = v23
    default:
        v24 := GenS2{F0: 32768}
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        v2["a"] = fi25
        _ = fi25
    }
    v26 := v2[(("hello" + ("hello" + "a")) + "foo bar")]
    if v1 {
        if (!(!v1)) {
            v8 += 2
        } else {
            var v27 uint8 = 2
            v5 = append(v5, "a")
            _ = v27
        }
        for fi28 := 0; fi28 < 17; fi28++ {
            v29 := GenS3{F0: 836, F1: 0, F2: 2}
            for fi30 := 0; fi30 < 11; fi30++ {
                v31 := GenS3{F0: 0, F1: 128, F2: 1}
                v8 = (^(v8))
                switch (v6 | fi30) {
                case 2147483647:
                    var v32 uint16 = v17
                    fmt.Println(v26, v1, v8)
                    _ = v32
                case 401:
                    switch "a" {
                    case "":
                        var v33 int32 = -819
                        _ = v33
                    case "foo bar":
                        var v34 int64 = 499
                        _ = v34
                    case "a":
                        var v35 int16 = (v3 | v3)
                        _ = v35
                    default:
                        v26 = 203
                        if len(v5) > 0 {
                            v5[0] = (("hello" + ("foo bar" + "foo bar")) + "")
                        }
                    }
                }
                _ = fi30
                _ = v31
            }
            _ = fi28
            _ = v29
        }
        fmt.Println(v26, v1, v8)
    }
    fmt.Println(v1, v1, v3)
    if (!false) {
        var v36 uint16 = (v17 - 32767)
        fmt.Println(v26, v3, v36)
        _ = v36
    } else {
        v37 := GenS4{F0: (v8 * (v8 / ((v8) | 1))), F1: ((v17 * (v17 & (v17 >> ((259) & 63)))) <= (^(v17)))}
        if (!(v8 != v8)) {
            v38 := GenS2{F0: (^(v26))}
            _ = v38
        } else {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                var v39 int64 = -1
                _ = v39
            case "foo bar":
                for fi40 := 0; fi40 < 5; fi40++ {
                    switch "foo bar" {
                    case "foo bar":
                        var v41 int64 = -913
                        _ = v41
                    case "\n\t":
                        v8 += (v8 % ((629) | 1))
                    default:
                        fmt.Println(v26)
                    }
                    var v42 int = 1
                    _ = fi40
                    _ = v42
                }
            case "hello":
                if (!((42 <= 607) || (0.1 > 3.14159))) {
                    fmt.Println(v3, v1)
                    if false {
                        var v43 int64 = -1
                        if (-128 > 52) {
                            if v1 {
                                fmt.Println(v26, v43)
                            }
                            v1 = v1
                            if (!(-83 != int8(v17))) {
                                v8 += (v8 * (v8 % ((2) | 1)))
                                switch "hello" {
                                case "hello":
                                    v44 := GenS4{F0: v8, F1: (105 <= 42)}
                                    _ = v44
                                case "foo bar":
                                    v45 := GenS5{F0: v43, F1: v3, F2: v6}
                                    _ = v45
                                }
                                if ((!v1) && (531 > 18446744073709551615)) {
                                    for fi46 := 0; fi46 < 12; fi46++ {
                                        v17 = 1
                                        _ = fi46
                                    }
                                    switch (v6 >> ((2) & 63)) {
                                    case 1:
                                        v1 = (255 >= 0)
                                    default:
                                        var v47 bool = (v26 >= (v26 + 42))
                                        _ = v47
                                    }
                                    var v48 int16 = 1000
                                    _ = v48
                                } else {
                                    v3 = (v3 >> ((4294967295) & 63))
                                    var v49 uint16 = 547
                                    for fi50 := 0; fi50 < 1; fi50++ {
                                        var v51 int16 = (v3 * v3)
                                        if (v1 || v1) {
                                            var v52 uint8 = 0
                                            for fi53 := 0; fi53 < 5; fi53++ {
                                                fmt.Println(v26, v49)
                                                v3 = -1
                                                _ = fi53
                                            }
                                            _ = len(v2)
                                            _ = v52
                                        }
                                        var v54 int64 = v43
                                        _ = fi50
                                        _ = v51
                                        _ = v54
                                    }
                                    _ = v49
                                }
                            }
                        }
                        if len(v4) > 0 {
                            v4[0] = 655
                        }
                        _ = v43
                    }
                }
            }
            var v55 string = "\n\t"
            for fi56 := 0; fi56 < 10; fi56++ {
                v3 = (v3 & (v3 ^ v3))
                for fi57 := 0; fi57 < 18; fi57++ {
                    switch v55 {
                    case "a":
                        v6 -= 127
                    }
                    _ = fi57
                }
                _ = fi56
            }
            _ = v55
        }
        _ = v37
    }
    var v58 float32 = ((0.0 - 1e10) + 1e10)
    v17 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v8
}
