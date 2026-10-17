package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint32, p2 int) uint16 {
    switch ("hello" + ("\n\t" + "foo bar")) {
    case "":
        var v6 int16 = 194
        _ = v6
    case "a":
        p0 = 2
    case "foo bar":
        fmt.Println(p1, p0)
    }
    if true {
        var v7 float64 = 0.0
        if true {
            p0 = 18446744073709551615
            v8 := []int{p2, (p2 - p2), p2}
            for fi9 := 0; fi9 < 19; fi9++ {
                switch ("\n\t" + ((("foo bar" + "\n\t") + "a") + ((("foo bar" + ("foo bar" + "foo bar")) + "foo bar") + "a"))) {
                case "":
                    p2 -= p2
                case "foo bar":
                    p0 = 2
                default:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
                _ = fi9
            }
            _ = v8
        }
        fmt.Println(p2)
        _ = v7
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int32 = 42
            v12 := []int{p2}
            p1 = 65535
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 uint = 255
        var v14 int32 = 1
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return 820
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{2147483647}
    v3 := []string{"foo bar", "a", "hello", "\n\t"}
    var v4 uint16 = 1
    v5 := make(map[string]uint32)
    switch "a" {
    case "\n\t":
        v15 := GenS1{F0: "foo bar", F1: 2147483647, F2: 0}
        _ = v15
    case "a":
        switch "\n\t" {
        case "hello":
            var v16 int16 = 1
            _ = v16
        }
    case "":
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    if false {
        fmt.Println(v4)
        var v18 float32 = float32(1.5)
        v4 = 255
        _ = v18
    } else {
        switch "\n\t" {
        case "\n\t":
            var v19 float64 = (0.0 * float64(((0.0 - (1.5 * (3.14159 - ((1e10 / (1.5 - (float32((0.0 / -2.25)) + 0.1))) / 1.5)))) - float32(0.1))))
            _ = v19
        case "a":
            var v20 int32 = -564
            _ = v20
        default:
            if ("" == ((("" + (((((((("a" + "hello") + (((("hello" + "") + ("hello" + "hello")) + "\n\t") + "a")) + "hello") + ("foo bar" + ("a" + "hello"))) + ((("" + "hello") + "a") + "a")) + "hello") + "hello") + "a")) + "hello") + (("" + "") + (("foo bar" + ("\n\t" + "a")) + "\n\t")))) {
                v4 -= (v4 & (v4 + v4))
                for fi21 := 0; fi21 < 8; fi21++ {
                    v22 := GenS1{F0: "\n\t", F1: fi21, F2: 32767}
                    var v23 uint = 513
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            } else {
                v24 := GenS2{F0: 0}
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 int32 = -2147483648
            _ = fi25
            _ = v26
        }
    }
    v27 := v5["foo bar"]
    for fi28 := 0; fi28 < 14; fi28++ {
        for fi29 := 0; fi29 < 2; fi29++ {
            switch fi28 {
            case -679:
                v30 := GenS3{F0: (3.14159 > 0.0), F1: v4}
                _ = v30
            case -907:
                if (545 <= (-(fi28))) {
                    switch fi28 {
                    case -776:
                        v31 := GenS1{F0: "foo bar", F1: (fi28 - 32767), F2: 1}
                        _ = v31
                    case -9223372036854775808:
                        var v32 uint32 = (v27 & v27)
                        _ = v32
                    }
                    v33 := GenS3{F0: false, F1: (v4 & v4)}
                    var v34 uint16 = (v4 / ((v4) | 1))
                    _ = v33
                    _ = v34
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        var v35 int = fi29
                        _ = v35
                    case "foo bar":
                        switch "" {
                        case "a":
                            v36 := GenS1{F0: "\n\t", F1: 128, F2: 1000}
                            _ = v36
                        }
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            v4 += v4
                        case "foo bar":
                            var v37 uint = 4294967295
                            _ = v37
                        }
                    }
                }
            case 4294967295:
                fmt.Println(fi29)
            default:
                fmt.Println(v27, v4)
            }
            var v38 int64 = 684
            var v39 int8 = 42
            _ = fi29
            _ = v38
            _ = v39
        }
        _ = fi28
    }
    fmt.Println(v27, v27)
    _ = len(v1)
    v40 := GenS4{F0: 1}
    if false {
        v41 := v5[("hello" + "")]
        v42 := GenS3{F0: (!(-1 >= -483)), F1: v4}
        _ = v41
        _ = v42
    } else {
        var v43 uint8 = 1
        if (v43 == v43) {
            switch int(v27) {
            case 764:
                v44 := GenS5{F0: -113}
                _ = v44
            case 9223372036854775807:
                for fi45 := 0; fi45 < 5; fi45++ {
                    v43 -= v43
                    v27 += (v27 - 2)
                    _ = fi45
                }
            case 4294967295:
                _ = len(v3)
            default:
                var v46 string = "foo bar"
                _ = v46
            }
            v43 += uint8(v27)
        } else {
            for fi47 := 0; fi47 < 11; fi47++ {
                var v48 int64 = 133
                fmt.Println(v4, fi47)
                switch 2147483647 {
                case 32768:
                    var v49 uint64 = 0
                    _ = v49
                }
                _ = fi47
                _ = v48
            }
            var v50 int32 = -1
            _ = v50
        }
        _ = v43
    }
    v51 := GenS2{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v51
}
