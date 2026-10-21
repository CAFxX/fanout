package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint16; F1 int16; F2 string }
type GenS5 struct { F0 int16; F1 bool; F2 int64 }
func genFn0(p0 int, p1 int8) int8 {
    v7 := GenS1{F0: true, F1: ("\n\t" + "a")}
    v8 := []int{(p0 & (p0 >> ((2) & 63))), p0}
    var v9 int32 = -446
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 int8, p1 int) uint8 {
    if true {
        switch (p1 ^ 32768) {
        case 624:
            if true {
                p0 = (^(p0))
                var v10 uint64 = 1
                _ = v10
            } else {
                v11 := GenS1{F0: (p1 == p1), F1: ""}
                var v12 uint16 = 42
                _ = v11
                _ = v12
            }
        case 4294967295:
            if true {
                if (42 == 546) {
                    v13 := GenS1{F0: (("" + (("\n\t" + "\n\t") + "\n\t")) < ("foo bar" + "hello")), F1: ("" + "hello")}
                    _ = v13
                }
            }
        case 32768:
            switch "hello" {
            case "a":
                v14 := []int{p1, p1, p1, p1}
                _ = v14
            case "foo bar":
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 float32 = 0.0
                    for fi17 := 0; fi17 < 6; fi17++ {
                        if true {
                            var v18 bool = (!false)
                            _ = v18
                        }
                        var v19 uint = 42
                        v20 := GenS2{F0: 1, F1: 2147483647, F2: 16}
                        _ = fi17
                        _ = v19
                        _ = v20
                    }
                    _ = fi15
                    _ = v16
                }
            case "":
                var v21 uint64 = 935
                _ = v21
            }
        }
        v22 := []int{}
        _ = v22
    } else {
        var v23 int32 = 2147483647
        fmt.Println(p1, p0, v23)
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        switch "\n\t" {
        case "foo bar":
            if true {
                switch "a" {
                case "foo bar":
                    p0 += 1
                case "":
                    fmt.Println(p0, fi24, fi24)
                default:
                    for fi25 := 0; fi25 < 11; fi25++ {
                        p1 = p1
                        if (255 < 2) {
                            var v26 float32 = float32(0.0)
                            p0 += p0
                            switch fi25 {
                            case 42:
                                p0 = p0
                            case -570:
                                v27 := []int{2147483647, fi25}
                                _ = v27
                            default:
                                fmt.Println(p0, p0, v26)
                            }
                            _ = v26
                        }
                        _ = fi25
                    }
                }
                v28 := GenS3{F0: "foo bar", F1: "hello"}
                _ = v28
            }
        case "hello":
            var v29 string = (((((((("a" + "\n\t") + (("\n\t" + "a") + (("a" + ("\n\t" + ("foo bar" + "a"))) + ""))) + ("" + (("a" + "foo bar") + "hello"))) + "hello") + ("a" + "hello")) + "foo bar") + ("" + "\n\t")) + "")
            _ = v29
        }
        _ = fi24
    }
    v30 := GenS3{F0: "\n\t", F1: "foo bar"}
    var v31 uint8 = 251
    _ = p0
    _ = p1
    _ = v30
    _ = v31
    return v31
}

func main() {
    var v1 float32 = 1e10
    var v2 int64 = -9223372036854775808
    var v3 int = 128
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]bool)
    var v6 float32 = 0.0
    _ = len(v4)
    for fi32 := 0; fi32 < 1; fi32++ {
        switch (fi32 * fi32) {
        case 0:
            v33 := GenS2{F0: 0, F1: 2147483647, F2: 1}
            _ = v33
        case -9223372036854775808:
            v34 := GenS2{F0: 869, F1: 65535, F2: 374}
            _ = v34
        case 4294967295:
            var v35 float32 = float32(3.14159)
            _ = v35
        default:
            fmt.Println(fi32)
        }
        _ = fi32
    }
    v36 := []int{}
    switch (v3 >> ((18446744073709551615) & 63)) {
    case -830:
        fmt.Println(v3, v2)
    case -672:
        if (-1 == 1) {
            v37 := GenS4{F0: 244, F1: 341, F2: "a"}
            var v38 bool = true
            _ = v37
            _ = v38
        }
    case 127:
        delete(v5, ("a" + "foo bar"))
    default:
        for fi39 := 0; fi39 < 1; fi39++ {
            for fi40 := 0; fi40 < 2; fi40++ {
                fmt.Println(v1, fi40, fi39)
                var v41 uint32 = 1
                var v42 uint = 2
                _ = fi40
                _ = v41
                _ = v42
            }
            var v43 uint64 = 0
            var v44 uint8 = 255
            _ = fi39
            _ = v43
            _ = v44
        }
    }
    v45 := v5["hello"]
    var v46 float64 = 3.14159
    var v47 bool = (!false)
    if (385 >= 0) {
        for fi48 := 0; fi48 < 2; fi48++ {
            fmt.Println(fi48)
            fmt.Println(v1, v2, v45)
            _ = fi48
        }
        if v47 {
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
        }
    }
    for fi49 := 0; fi49 < 8; fi49++ {
        fmt.Println(v2, v46)
        var v50 uint32 = 2147483648
        _ = fi49
        _ = v50
    }
    var v51 int16 = -240
    var v52 uint16 = 228
    var v53 uint = 0
    var v54 uint16 = (^(v52))
    v55 := GenS5{F0: v51, F1: v45, F2: v2}
    if (0 == 18446744073709551615) {
        _ = len(v36)
        var v56 bool = (v54 >= v54)
        _ = v56
    } else {
        if v45 {
            var v57 bool = v45
            _ = v57
        } else {
            for fi58 := 0; fi58 < 3; fi58++ {
                _ = len(v36)
                _ = fi58
            }
        }
    }
    v53 += v53
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
