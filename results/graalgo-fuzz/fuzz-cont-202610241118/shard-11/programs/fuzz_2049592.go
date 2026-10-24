package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint64 }
func genFn0(p0 int32, p1 bool) string {
    v6 := []int{32767, 32768, 4294967295}
    switch "" {
    case "":
        p0 = 2147483647
    case "a":
        for fi7 := 0; fi7 < 13; fi7++ {
            if p1 {
                var v8 float32 = -2.25
                _ = v8
            }
            _ = fi7
        }
    default:
        var v9 int32 = p0
        _ = v9
    }
    var v10 uint8 = 255
    var v11 uint32 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    return ("hello" + ("a" + "a"))
}

func main() {
    var v1 uint16 = 42
    var v2 float32 = ((3.14159 * 1.5) * 1.5)
    var v3 string = "hello"
    v4 := []string{(v3 + ("\n\t" + "\n\t"))}
    var v5 int32 = -283
    if ((38 >= 42) || ((v2 > 0.1) && (!(721 > 0)))) {
        v12 := GenS1{F0: 0, F1: "foo bar", F2: 255}
        v13 := GenS1{F0: 0, F1: "foo bar", F2: 65535}
        _ = v12
        _ = v13
    } else {
        var v14 uint32 = 935
        v15 := GenS1{F0: -1, F1: "foo bar", F2: 18446744073709551615}
        _ = v14
        _ = v15
    }
    if true {
        v16 := GenS1{F0: -1, F1: "", F2: 42}
        _ = v16
    }
    if false {
        switch v3 {
        case "":
            if (32767 <= -381) {
                v17 := GenS1{F0: 100, F1: "\n\t", F2: 899}
                v18 := GenS2{F0: 1, F1: 0}
                _ = v17
                _ = v18
            }
        case "foo bar":
            v4 = append(v4, v3)
        case "a":
            for fi19 := 0; fi19 < 9; fi19++ {
                fmt.Println(v2, v2)
                switch v3 {
                case "foo bar":
                    v20 := GenS1{F0: 1, F1: v3, F2: 4294967295}
                    _ = v20
                case "\n\t":
                    for fi21 := 0; fi21 < 9; fi21++ {
                        var v22 uint64 = 42
                        _ = fi21
                        _ = v22
                    }
                default:
                    var v23 int32 = v5
                    _ = v23
                }
                _ = fi19
            }
        }
        v5 = v5
        v2 = 0.1
    }
    if ((!(false && true)) || true) {
        for fi24 := 0; fi24 < 1; fi24++ {
            v2 = (v2 / (v2 + float32(1.5)))
            _ = fi24
        }
        v2 = v2
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v5 = (v5 >> ((42) & 63))
    v1 += (v1 ^ 42)
    for fi25 := 0; fi25 < 4; fi25++ {
        v4 = append(v4, "foo bar")
        _ = fi25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 int64 = 865
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: 0, F1: "\n\t", F2: 18446744073709551615}
    for fi29 := 0; fi29 < 15; fi29++ {
        v30 := GenS3{F0: 2, F1: 1}
        _ = fi29
        _ = v30
    }
    fmt.Println(v3)
    switch v3 {
    case "hello":
        if false {
            v5 -= -2147483648
        } else {
            fmt.Println(v2, v2)
            fmt.Println(v1)
            v5 += v5
        }
    default:
        if false {
            var v31 int32 = v5
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        v1 = 32767
        v33 := GenS4{F0: false, F1: (v1 - 441), F2: 4294967295}
        switch (fi32 / ((fi32) | 1)) {
        case -132:
            switch ("a" + (("a" + ((((v3 + v3) + ("\n\t" + "hello")) + "") + "")) + "a")) {
            case "\n\t":
                switch -507 {
                case 1:
                    for fi34 := 0; fi34 < 18; fi34++ {
                        fmt.Println(v2)
                        v5 = (v5 << ((4294967295) & 63))
                        var v35 int64 = -570
                        _ = fi34
                        _ = v35
                    }
                }
            case "hello":
                _ = len(v4)
            case "":
                v36 := GenS3{F0: 241, F1: 18446744073709551615}
                _ = v36
            }
        case 0:
            var v37 float64 = float64(float32(((0.1 / 0.1) * 3.14159)))
            _ = v37
        case 334:
            var v38 int32 = (v5 >> ((65535) & 63))
            _ = v38
        }
        _ = fi32
        _ = v33
    }
    for fi39 := 0; fi39 < 7; fi39++ {
        for fi40 := 0; fi40 < 16; fi40++ {
            if len(v4) > 0 {
                v4[0] = v3
            }
            v2 += v2
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if ((v2 * v2) > ((v2 / v2) / v2)) {
                    var v41 int64 = -898
                    v4 = append(v4, v3)
                    if (uint(v1) > 702) {
                        var v42 uint64 = 0
                        v3 = (v3 + ("hello" + v3))
                        _ = v42
                    } else {
                        var v43 int8 = -1
                        _ = v43
                    }
                    _ = v41
                } else {
                    fmt.Println(fi40, v5, v3)
                    fmt.Println(v1)
                    fmt.Println(fi40, v1, fi39)
                    v2 -= v2
                }
                _ = len(v4)
            } else {
                if true {
                    v3 = ((v3 + v3) + ("\n\t" + (v3 + v3)))
                }
                if false {
                    if len(v4) > 0 {
                        v4[0] = (("\n\t" + v3) + ("\n\t" + ""))
                    }
                }
                v1 = (v1 ^ 65535)
            }
            _ = fi40
        }
        var v44 float64 = 0.0
        switch ("foo bar" + v3) {
        case "a":
            var v45 int = (fi39 | (fi39 / (((fi39 / (((fi39 << ((835) & 63))) | 1))) | 1)))
            _ = v45
        case "foo bar":
            var v46 float64 = (1e10 * float64(float32((float64(v2) - (v44 / 3.14159)))))
            _ = v46
        case "\n\t":
            v2 = (((0.0 - v2) / ((0.1 / v2) - v2)) * (v2 * -2.25))
        }
        _ = fi39
        _ = v44
    }
    if false {
        var v47 int16 = 42
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
