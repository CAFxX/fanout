package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int16; F1 uint32; F2 int }
type GenS4 struct { F0 int64 }
func genFn0(p0 int64) uint32 {
    var v5 int8 = 0
    var v6 int32 = 1
    v7 := GenS1{F0: 65535}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1() uint8 {
    if (true || true) {
        switch (("hello" + "hello") + (("\n\t" + "hello") + "hello")) {
        case "a":
            var v8 string = ("foo bar" + "")
            _ = v8
        case "":
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 int32 = 2147483647
                _ = fi9
                _ = v10
            }
        case "hello":
            for fi11 := 0; fi11 < 3; fi11++ {
                switch "a" {
                case "\n\t":
                    for fi12 := 0; fi12 < 9; fi12++ {
                        var v13 bool = true
                        var v14 int16 = 32767
                        fmt.Println(v14, v13, fi12)
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                v15 := GenS2{F0: -82, F1: 285}
                _ = fi11
                _ = v15
            }
        default:
            var v16 uint = 516
            _ = v16
        }
        var v17 uint = 4294967295
        _ = v17
    } else {
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 int32 = 0
            v20 := []int{}
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 int64 = 42
        _ = v21
    }
    if (127 != 8) {
        if (1 > 4294967295) {
            v22 := GenS1{F0: 2}
            if (false || (42 <= 713)) {
                v23 := GenS3{F0: 42, F1: genFn0(743), F2: 2147483647}
                _ = v23
            } else {
                var v24 int16 = -322
                var v25 int64 = 42
                v26 := GenS1{F0: 4294967295}
                _ = v24
                _ = v25
                _ = v26
            }
            _ = v22
        } else {
            var v27 string = "hello"
            v28 := GenS1{F0: 2147483647}
            var v29 uint8 = 134
            _ = v27
            _ = v28
            _ = v29
        }
        var v30 string = "a"
        _ = v30
    } else {
        var v31 int32 = -2147483648
        _ = v31
    }
    var v32 float32 = 1.5
    v33 := GenS3{F0: -32768, F1: 42, F2: -460}
    v32 = v32
    _ = v32
    _ = v33
    return 111
}

func main() {
    v1 := []string{"hello", ("foo bar" + "hello"), "foo bar"}
    v2 := make(map[string]string)
    v3 := []int{-478}
    var v4 bool = false
    for fi34 := 0; fi34 < 1; fi34++ {
        _ = len(v2)
        if v4 {
            for fi35 := 0; fi35 < 14; fi35++ {
                fmt.Println(v4, fi35)
                var v36 int64 = -9223372036854775808
                _ = fi35
                _ = v36
            }
            var v37 string = (("a" + (("" + ("hello" + (("hello" + ("" + "\n\t")) + "hello"))) + "a")) + "foo bar")
            fmt.Println(fi34, fi34)
            _ = v37
        }
        switch "" {
        case "a":
            switch "foo bar" {
            case "hello":
                var v38 string = ""
                _ = v38
            case "":
                for fi39 := 0; fi39 < 7; fi39++ {
                    v40 := GenS1{F0: 122}
                    var v41 int16 = -1
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            case "a":
                _ = len(v1)
            }
        case "":
            if false {
                v4 = (2 <= 42)
                for fi42 := 0; fi42 < 17; fi42++ {
                    v4 = true
                    _ = fi42
                }
                var v43 float32 = float32(3.14159)
                _ = v43
            }
        }
        _ = fi34
    }
    var v44 uint8 = 128
    fmt.Println(v44, v4)
    fmt.Println(v44)
    _ = len(v1)
    switch "a" {
    case "":
        var v45 float32 = 3.14159
        _ = v45
    case "\n\t":
        v44 = (v44 - v44)
    case "hello":
        v4 = ((v4 && true) && (4294967295 != 65535))
    default:
        v44 = 251
    }
    fmt.Println(v44, v44)
    v44 -= v44
    var v46 float32 = -2.25
    for fi47 := 0; fi47 < 12; fi47++ {
        if len(v3) > 0 {
            v3[0] = 32767
        }
        for fi48 := 0; fi48 < 18; fi48++ {
            var v49 int32 = -752
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch (-(fi47)) {
            case 4294967295:
                if ((4294967295 != 725) || true) {
                    fmt.Println(v44)
                }
            case -1:
                v44 = (v44 >> ((257) & 63))
            }
            _ = fi48
            _ = v49
        }
        var v50 int8 = 1
        _ = fi47
        _ = v50
    }
    if (!(v4 || ((!v4) || (uint64(v44) == 42)))) {
        switch -452 {
        case 32767:
            v4 = (625 > -264)
        case -681:
            for fi51 := 0; fi51 < 5; fi51++ {
                var v52 int32 = 1
                switch "\n\t" {
                case "":
                    if ((-128 <= 49) || v4) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch 845 {
                        case -1:
                            v46 = v46
                        case -821:
                            for fi53 := 0; fi53 < 16; fi53++ {
                                _ = len(v3)
                                var v54 float32 = -2.25
                                v55 := GenS1{F0: genFn0(697)}
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                        default:
                            switch (fi51 << ((0) & 63)) {
                            case -9223372036854775808:
                                var v56 uint16 = 32767
                                _ = v56
                            }
                        }
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi51
                _ = v52
            }
        }
        for fi57 := 0; fi57 < 7; fi57++ {
            switch 32768 {
            case 128:
                var v58 int16 = 42
                _ = v58
            case 0:
                if len(v1) > 0 {
                    v1[0] = ""
                }
            default:
                if false {
                    fmt.Println(v46, fi57, v44)
                    v4 = v4
                    _ = len(v3)
                }
            }
            switch "a" {
            case "foo bar":
                fmt.Println(v44)
            case "a":
                if v4 {
                    v4 = (-2147483648 <= 0)
                    fmt.Println(fi57, fi57)
                }
            case "hello":
                v59 := GenS2{F0: -1, F1: 849}
                _ = v59
            }
            v60 := GenS4{F0: 0}
            _ = fi57
            _ = v60
        }
        if (265 == 2) {
            for fi61 := 0; fi61 < 17; fi61++ {
                v3 = append(v3, (fi61 * fi61))
                var v62 uint16 = 65535
                _ = fi61
                _ = v62
            }
            _ = len(v1)
            _ = len(v1)
        } else {
            v4 = (true && (-614 >= 4294967295))
        }
    }
    v44 = v44
    var v63 int64 = -417
    v44 += 204
    var v64 string = ((("hello" + ("foo bar" + ("foo bar" + "foo bar"))) + "") + "hello")
    switch v64 {
    case "hello":
        v4 = (620 == 42)
    case "\n\t":
        v65 := v2["hello"]
        _ = v65
    case "foo bar":
        fmt.Println(v46, v44, v63)
    default:
        switch "foo bar" {
        case "\n\t":
            var v66 int = -967
            _ = v66
        case "":
            for fi67 := 0; fi67 < 20; fi67++ {
                var v68 int8 = -1
                _ = fi67
                _ = v68
            }
        }
    }
    if v4 {
        var v69 int = 42
        var v70 uint64 = 18446744073709551615
        _ = v69
        _ = v70
    }
    fmt.Println(v4, v4, v44)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v63
    _ = v64
}
