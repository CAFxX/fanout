package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string; F2 int }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 string; F1 bool; F2 bool }
type GenS6 struct { F0 uint8; F1 uint; F2 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 uint16 }
func genFn0(p0 uint32) uint64 {
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            if (-9223372036854775808 >= fi7) {
                v9 := GenS1{F0: -128}
                var v10 uint64 = 1
                if (-306 >= 0) {
                    p0 -= p0
                    p0 = p0
                    fmt.Println(v10, v10, v10)
                } else {
                    v11 := []int{fi8, (fi7 >> ((797) & 63)), 4294967295, fi8}
                    v12 := GenS2{F0: -32768, F1: "foo bar", F2: fi7}
                    _ = v11
                    _ = v12
                }
                _ = v10
                _ = v9
            } else {
                var v13 uint64 = 1
                v14 := GenS3{F0: false, F1: (("foo bar" + "") + "hello")}
                var v15 int8 = -1
                _ = v13
                _ = v14
                _ = v15
            }
            switch (fi8 * fi7) {
            case 968:
                var v16 uint = 890
                _ = v16
            case 2:
                var v17 uint = 65535
                _ = v17
            default:
                p0 -= (p0 - p0)
            }
            _ = fi8
        }
        v18 := GenS2{F0: 32767, F1: ("" + ""), F2: fi7}
        _ = fi7
        _ = v18
    }
    switch -1 {
    case 4294967295:
        var v19 float32 = 3.14159
        _ = v19
    case -257:
        switch "foo bar" {
        case "":
            v20 := GenS1{F0: 100}
            _ = v20
        case "a":
            v21 := []int{-2147483648, 42, 246}
            _ = v21
        case "hello":
            v22 := GenS1{F0: 42}
            _ = v22
        }
    case -9223372036854775808:
        switch "\n\t" {
        case "":
            var v23 int8 = 0
            _ = v23
        case "a":
            if (!(-659 < -1)) {
                v24 := GenS2{F0: 32767, F1: "a", F2: 0}
                v25 := []int{32767, 127, 32767, -454}
                _ = v24
                _ = v25
            }
        }
    default:
        v26 := GenS2{F0: 851, F1: "a", F2: -2147483648}
        _ = v26
    }
    var v27 int8 = int8(p0)
    _ = p0
    _ = v27
    return 42
}

func main() {
    var v1 int16 = 1
    var v2 uint32 = 1
    var v3 int = 1
    v4 := make(map[string]bool)
    v5 := []string{("a" + (("\n\t" + "") + (("" + "\n\t") + ("foo bar" + ("" + ((("hello" + "\n\t") + ("foo bar" + "a")) + (("a" + "") + "a")))))))}
    var v6 int8 = 127
    switch ("a" + "") {
    case "a":
        v28 := GenS4{F0: v1, F1: "hello"}
        _ = v28
    }
    v3 += v3
    v3 = v3
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    if false {
        var v29 int8 = (v6 * (v6 / (((^(v6))) | 1)))
        v30 := GenS5{F0: "", F1: true, F2: true}
        var v31 int16 = v1
        _ = v29
        _ = v30
        _ = v31
    }
    switch "a" {
    case "":
        for fi32 := 0; fi32 < 4; fi32++ {
            v33 := GenS6{F0: 42, F1: 4294967295, F2: v3}
            var v34 float32 = 0.0
            for fi35 := 0; fi35 < 8; fi35++ {
                v2 -= v2
                v2 -= 1
                _ = len(v4)
                _ = fi35
            }
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    switch "hello" {
    case "a":
        for fi36 := 0; fi36 < 1; fi36++ {
            v6 = (v6 & v6)
            switch "" {
            case "":
                v1 -= -423
            case "hello":
                if (!false) {
                    var v37 int32 = 1
                    var v38 int32 = 741
                    _ = v37
                    _ = v38
                }
            default:
                if (((1e10 + (0.1 - float32(0.1))) + (0.1 / 1.5)) > 0.0) {
                    if true {
                        switch -286 {
                        case 0:
                            if ((v1 - v1) <= (v1 + (v1 * (v1 + (v1 + 42))))) {
                                v39 := GenS2{F0: (v1 % ((v1) | 1)), F1: "", F2: (fi36 << ((2) & 63))}
                                _ = len(v5)
                                if false {
                                    fmt.Println(v6, v1)
                                    for fi40 := 0; fi40 < 5; fi40++ {
                                        var v41 uint64 = genFn0(968)
                                        _ = fi40
                                        _ = v41
                                    }
                                }
                                _ = v39
                            }
                        case 128:
                            var v42 bool = (255 > 198)
                            _ = v42
                        case 1:
                            var v43 string = ("" + "foo bar")
                            _ = v43
                        }
                    } else {
                        v2 = 1
                        v44 := GenS7{F0: 609}
                        var v45 uint32 = (v2 - 1)
                        _ = v44
                        _ = v45
                    }
                }
            }
            _ = fi36
        }
    default:
        v1 += 477
    }
    v46 := GenS8{F0: v3, F1: 42}
    v2 -= 228
    v6 = v6
    if ((v3 * 127) <= 2147483647) {
        var v47 float32 = 0.0
        if true {
            if false {
                v48 := GenS2{F0: v1, F1: ("foo bar" + ""), F2: (v3 - (v3 ^ v3))}
                _ = v48
            }
            v3 = (v3 & v3)
            switch (("" + "hello") + "\n\t") {
            case "hello":
                v1 -= v1
            case "foo bar":
                v2 = 517
            }
        } else {
            var v49 int32 = -584
            fmt.Println(v3, v3, v3)
            _ = v49
        }
        for fi50 := 0; fi50 < 12; fi50++ {
            var v51 uint16 = 2
            fmt.Println(v1)
            _ = fi50
            _ = v51
        }
        _ = v47
    } else {
        var v52 string = "a"
        for fi53 := 0; fi53 < 3; fi53++ {
            v54 := GenS2{F0: -956, F1: v52, F2: -1}
            v4[v52] = (3.14159 > 0.1)
            if false {
                v6 += (v6 | 1)
            } else {
                var v55 float64 = 0.1
                _ = v55
            }
            _ = fi53
            _ = v54
        }
        _ = v52
    }
    fmt.Println(v1, v6, v6)
    var v56 string = ""
    for fi57 := 0; fi57 < 19; fi57++ {
        fmt.Println(v1, v56, fi57)
        v56 = (v56 + (v56 + v56))
        switch (fi57 << ((2) & 63)) {
        case 42:
            var v58 int8 = (^(v6))
            _ = v58
        case -1:
            fmt.Println(v2, v56)
        default:
            v59 := GenS7{F0: v1}
            _ = v59
        }
        _ = fi57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v56
    _ = v6
}
