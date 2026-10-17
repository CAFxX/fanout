package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int16; F1 uint64; F2 int }
type GenS3 struct { F0 int64; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    v3 := []string{("a" + "hello"), "foo bar"}
    v4 := []string{"hello", "", ("\n\t" + (("hello" + ("hello" + "hello")) + "hello"))}
    var v5 int = 127
    switch "" {
    case "hello":
        var v6 int16 = 1000
        _ = v6
    case "a":
        fmt.Println(v5)
    case "":
        var v7 int16 = 1000
        _ = v7
    }
    if true {
        v2["foo bar"] = (true && ((2147483647 == 0) || false))
        v8 := GenS1{F0: 865, F1: (v5 & (v5 / ((v5) | 1)))}
        v3 = append(v3, ("foo bar" + "foo bar"))
        _ = v8
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    v3 = append(v3, "a")
    for fi9 := 0; fi9 < 1; fi9++ {
        v5 = (v5 & (-(fi9)))
        switch ("\n\t" + "") {
        case "a":
            v10 := v1[2147483647]
            _ = v10
        case "\n\t":
            for fi11 := 0; fi11 < 6; fi11++ {
                v5 -= fi9
                _ = fi11
            }
        case "foo bar":
            fmt.Println(fi9, v5)
        default:
            if (!(!false)) {
                var v12 float64 = 1e10
                _ = v12
            } else {
                if (fi9 < -46) {
                    var v13 string = "hello"
                    _ = v13
                }
                v14 := GenS2{F0: 1, F1: 1, F2: (-(v5))}
                if (((!false) || ((4294967295 == 1) || (255 <= 255))) || true) {
                    var v15 int64 = 1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v16 int = (fi9 << ((255) & 63))
                    _ = v15
                    _ = v16
                } else {
                    if true {
                        var v17 int64 = 0
                        _ = v17
                    } else {
                        _ = len(v1)
                        if (true && false) {
                            var v18 int8 = 1
                            fmt.Println(v18, v5)
                            _ = v18
                        }
                    }
                    v5 = fi9
                    v5 = (fi9 * (v5 >> ((18446744073709551615) & 63)))
                }
                _ = v14
            }
        }
        _ = fi9
    }
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "foo bar")
    }
    fmt.Println(v5)
    if true {
        if false {
            v19 := GenS1{F0: 151, F1: (v5 >> ((577) & 63))}
            _ = v19
        } else {
            v20 := GenS1{F0: -2147483648, F1: v5}
            var v21 uint32 = 88
            _ = v20
            _ = v21
        }
        fmt.Println(v5)
        for fi22 := 0; fi22 < 9; fi22++ {
            if (255 == 18446744073709551615) {
                var v23 int16 = 32767
                v23 += 427
                v23 += v23
                _ = v23
            }
            fmt.Println(fi22)
            var v24 bool = (true && (2147483647 <= 1))
            _ = fi22
            _ = v24
        }
    } else {
        if false {
            if false {
                v5 += v5
                v3 = append(v3, "\n\t")
            } else {
                if false {
                    v25 := GenS2{F0: -1, F1: 418, F2: (-(v5))}
                    _ = v25
                }
                for fi26 := 0; fi26 < 19; fi26++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (1 < -1) {
                        var v27 string = "\n\t"
                        for fi28 := 0; fi28 < 16; fi28++ {
                            v29 := GenS2{F0: -1, F1: 2, F2: v5}
                            _ = fi28
                            _ = v29
                        }
                        _ = len(v2)
                        _ = v27
                    }
                    for fi30 := 0; fi30 < 4; fi30++ {
                        var v31 int64 = 0
                        v32 := GenS2{F0: 0, F1: 4294967295, F2: 415}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    _ = fi26
                }
            }
            v5 = (v5 - v5)
            v1[595] = v5
        } else {
            var v33 int32 = 0
            v5 = v5
            _ = v33
        }
        v34 := GenS2{F0: -72, F1: 42, F2: 32767}
        for fi35 := 0; fi35 < 9; fi35++ {
            v36 := v2[("foo bar" + (("foo bar" + ("foo bar" + (((((("" + ((("a" + (("" + (("hello" + (("hello" + "foo bar") + "a")) + "\n\t")) + "hello")) + "foo bar") + "\n\t")) + ("hello" + "a")) + "") + "\n\t") + "a") + "a"))) + ""))]
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    var v37 int64 = 1
    for fi38 := 0; fi38 < 6; fi38++ {
        v39 := GenS2{F0: 1000, F1: 18446744073709551615, F2: v5}
        v37 += v37
        _ = fi38
        _ = v39
    }
    _ = len(v1)
    switch v5 {
    case 0:
        var v40 uint32 = 0
        _ = v40
    case -80:
        v41 := GenS3{F0: v37, F1: 73, F2: 1}
        _ = v41
    case 32767:
        if false {
            var v42 string = (("foo bar" + "a") + "a")
            v43 := GenS1{F0: -1, F1: (v5 - (^(v5)))}
            _ = v42
            _ = v43
        } else {
            for fi44 := 0; fi44 < 4; fi44++ {
                v4 = append(v4, ((("a" + "\n\t") + "hello") + "hello"))
                _ = fi44
            }
            v5 = (v5 >> ((1) & 63))
            var v45 float64 = (1e10 * 1.5)
            _ = v45
        }
    }
    v46 := GenS4{F0: 4294967295, F1: (404 >= 255)}
    fmt.Println(v5, v5)
    delete(v1, 2147483647)
    switch 32768 {
    case 2147483647:
        var v47 float32 = ((float32((float64(3.14159) - -2.25)) / ((0.1 - -2.25) / 1.0)) / (0.1 * 0.1))
        _ = v47
    case -9223372036854775808:
        for fi48 := 0; fi48 < 18; fi48++ {
            v5 -= fi48
            v37 += v37
            _ = fi48
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v46
    _ = v5
}
