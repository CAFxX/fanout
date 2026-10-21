package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint, p1 uint) int8 {
    var v5 int = 9223372036854775807
    switch 127 {
    case 42:
        for fi6 := 0; fi6 < 7; fi6++ {
            switch "\n\t" {
            case "foo bar":
                for fi7 := 0; fi7 < 2; fi7++ {
                    if (!true) {
                        var v8 uint32 = uint32(p1)
                        var v9 float32 = 3.14159
                        _ = v8
                        _ = v9
                    }
                    v10 := GenS1{F0: 42, F1: 42}
                    _ = fi7
                    _ = v10
                }
            case "":
                var v11 string = "a"
                _ = v11
            case "hello":
                switch "foo bar" {
                case "\n\t":
                    var v12 int8 = -117
                    _ = v12
                }
            }
            for fi13 := 0; fi13 < 16; fi13++ {
                if false {
                    var v14 int32 = 42
                    v15 := []int{v5}
                    v16 := GenS1{F0: (v14 | v14), F1: v14}
                    if len(v15) > 0 {
                        v15[0] = v5
                    }
                    if len(v15) > 0 {
                        v15[0] = -2147483648
                    }
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = fi13
            }
            _ = fi6
        }
    default:
        v17 := GenS2{F0: 42}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v5
    return 127
}

func main() {
    v1 := []string{"", (("" + "hello") + ("hello" + "\n\t")), "foo bar"}
    v2 := []int{-863, -2147483648}
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v18 bool = true
    v19 := GenS2{F0: 754}
    if v18 {
        switch 32767 {
        case 2147483647:
            v4[127] = v18
        case 0:
            switch 42 {
            case 614:
                v4[32768] = v18
            case 128:
                switch "hello" {
                case "hello":
                    v18 = v18
                default:
                    v18 = true
                }
            default:
                v18 = (2147483647 > 373)
            }
        default:
            switch (((("hello" + "") + ("\n\t" + "")) + "hello") + "foo bar") {
            case "\n\t":
                v20 := GenS1{F0: 496, F1: -850}
                _ = v20
            case "foo bar":
                v18 = v18
            default:
                for fi21 := 0; fi21 < 14; fi21++ {
                    fmt.Println(v18, v18)
                    _ = fi21
                }
            }
        }
        switch ("\n\t" + "a") {
        case "foo bar":
            _ = len(v4)
        case "":
            var v22 uint = 42
            _ = v22
        case "hello":
            for fi23 := 0; fi23 < 7; fi23++ {
                for fi24 := 0; fi24 < 14; fi24++ {
                    var v25 uint8 = 255
                    var v26 uint64 = 635
                    v25 -= (v25 & v25)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
        default:
            v27 := GenS2{F0: -2147483648}
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            v18 = false
            switch 4294967295 {
            case 2147483647:
                if ((fi29 >> ((2) & 63)) != fi28) {
                    var v30 bool = v18
                    v30 = v18
                    _ = v30
                }
            default:
                _ = len(v1)
                if false {
                    v18 = (((4294967295 == 2) && false) || false)
                } else {
                    v18 = ((65535 >= 255) && (65535 <= 42))
                    v18 = ((^(fi28)) >= 32767)
                }
            }
            v18 = (255 >= 83)
            _ = fi29
        }
        var v31 string = "a"
        for fi32 := 0; fi32 < 1; fi32++ {
            if (9223372036854775807 != 0) {
                v18 = (1 < 0)
                v33 := GenS3{F0: 2147483647, F1: -908}
                _ = v33
            } else {
                if (59 != 37) {
                    if v18 {
                        var v34 uint8 = 42
                        v34 -= (v34 / (((^(v34))) | 1))
                        var v35 float64 = 1.5
                        _ = v34
                        _ = v35
                    } else {
                        delete(v4, fi32)
                        switch (fi32 >> ((401) & 63)) {
                        case 4294967295:
                            var v36 uint64 = 42
                            _ = v36
                        case 127:
                            var v37 string = v31
                            _ = v37
                        case 9223372036854775807:
                            var v38 uint16 = 42
                            _ = v38
                        }
                    }
                } else {
                    var v39 uint = 42
                    fmt.Println(fi28, fi32, v39)
                    switch fi32 {
                    case -9223372036854775808:
                        v40 := GenS1{F0: 42, F1: 2147483647}
                        _ = v40
                    default:
                        v41 := GenS1{F0: -2147483648, F1: 0}
                        _ = v41
                    }
                    _ = v39
                }
                fmt.Println(v18)
                switch v31 {
                case "a":
                    v31 = ((v31 + "hello") + v31)
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = fi28
                    }
                    var v42 uint8 = 0
                    v42 += 255
                    _ = v42
                }
            }
            v31 = (v31 + "")
            _ = fi32
        }
        _ = fi28
        _ = v31
    }
    switch "foo bar" {
    case "foo bar":
        if false {
            v18 = false
            v43 := GenS4{F0: -732, F1: 128, F2: 42}
            _ = v43
        }
    }
    switch ("hello" + "foo bar") {
    case "hello":
        v18 = v18
    case "":
        if (!(true && (v18 && (1 > 0)))) {
            fmt.Println(v18)
            v44 := GenS3{F0: 2147483647, F1: -690}
            _ = v44
        }
    case "a":
        for fi45 := 0; fi45 < 5; fi45++ {
            var v46 int32 = 51
            fmt.Println(v18)
            v46 = (v46 % ((v46) | 1))
            _ = fi45
            _ = v46
        }
    default:
        var v47 float32 = (0.0 * (3.14159 + 3.14159))
        _ = v47
    }
    v18 = (!v18)
    v48 := GenS1{F0: -2147483648, F1: 42}
    delete(v3, 128)
    for fi49 := 0; fi49 < 19; fi49++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi50 := 0; fi50 < 6; fi50++ {
            v51 := GenS5{F0: 968, F1: 590}
            _ = fi50
            _ = v51
        }
        _ = fi49
    }
    fmt.Println(v18)
    for fi52 := 0; fi52 < 3; fi52++ {
        v18 = (genFn0(741, 4294967295) > genFn0(65535, 4294967295))
        fmt.Println(v18, fi52)
        switch "" {
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "\n\t":
            v18 = (-1 == -78)
        case "a":
            for fi53 := 0; fi53 < 2; fi53++ {
                if (!v18) {
                    var v54 string = ("\n\t" + "\n\t")
                    fmt.Println(fi53)
                    _ = v54
                }
                _ = fi53
            }
        default:
            switch 127 {
            case 496:
                var v55 int32 = 42
                _ = v55
            default:
                var v56 uint16 = 638
                _ = v56
            }
        }
        _ = fi52
    }
    switch 1 {
    case 9223372036854775807:
        var v57 uint16 = 1
        _ = v57
    case -2147483648:
        var v58 int32 = -1
        _ = v58
    case 32768:
        if (73 >= 0) {
            v2 = append(v2, -694)
            var v59 string = "a"
            _ = len(v2)
            _ = v59
        } else {
            var v60 string = "foo bar"
            v18 = false
            var v61 string = "a"
            _ = v60
            _ = v61
        }
    }
    if false {
        v18 = (!v18)
        v18 = v18
        if v18 {
            switch ((("a" + ("hello" + ("foo bar" + "foo bar"))) + "a") + (("a" + "hello") + "foo bar")) {
            case "":
                v18 = v18
            case "a":
                if (!(42 > 1)) {
                    var v62 string = (("\n\t" + (("" + "hello") + "hello")) + "foo bar")
                    fmt.Println(v62, v62)
                    v18 = false
                    _ = v62
                } else {
                    var v63 float64 = (0.0 / 3.14159)
                    v18 = v18
                    v64 := GenS3{F0: -2147483648, F1: 566}
                    _ = v63
                    _ = v64
                }
            case "foo bar":
                for fi65 := 0; fi65 < 16; fi65++ {
                    if v18 {
                        v66 := GenS6{F0: 0}
                        _ = v66
                    } else {
                        for fi67 := 0; fi67 < 16; fi67++ {
                            var v68 int32 = 2147483647
                            fmt.Println(fi65, fi65)
                            _ = fi67
                            _ = v68
                        }
                    }
                    v69 := GenS1{F0: 1, F1: -1}
                    _ = fi65
                    _ = v69
                }
            default:
                var v70 string = ("foo bar" + "")
                _ = v70
            }
        } else {
            var v71 uint64 = 1
            var v72 int8 = 42
            _ = v71
            _ = v72
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v48
}
