package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 uint16; F2 int32 }
func genFn0(p0 uint8, p1 uint8) string {
    p0 -= p0
    p0 = 1
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    var v1 float64 = (3.14159 + -2.25)
    var v2 float64 = (v1 - (float64((3.14159 - 1e10)) - 3.14159))
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    switch "a" {
    case "hello":
        fmt.Println(v1, v1)
    case "foo bar":
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: int32(fi6), F1: -2147483648}
            var v8 uint16 = 65535
            for fi9 := 0; fi9 < 15; fi9++ {
                if ((1e10 - 3.14159) > v2) {
                    fmt.Println(v1)
                    v10 := v3[fi9]
                    v11 := GenS2{F0: 1, F1: 100}
                    if false {
                        switch 1 {
                        case 1:
                            v12 := v4["\n\t"]
                            _ = v12
                        default:
                            v1 -= (-2.25 * 3.14159)
                        }
                        switch "hello" {
                        case "\n\t":
                            var v13 bool = (0 > 1000)
                            _ = v13
                        }
                    }
                    _ = v10
                    _ = v11
                }
                v8 = (v8 | (v8 >> ((2) & 63)))
                v14 := GenS3{F0: (^(fi6))}
                _ = fi9
                _ = v14
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    case "":
        var v15 int16 = 931
        _ = v15
    default:
        v16 := GenS4{F0: -1, F1: 4294967295}
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        switch (fi17 * -9223372036854775808) {
        case -91:
            for fi18 := 0; fi18 < 5; fi18++ {
                fmt.Println(v1, fi17)
                v19 := GenS5{F0: 2147483647, F1: 42}
                _ = len(v5)
                if false {
                    var v20 int16 = 1
                    v4[("" + "foo bar")] = 128
                    _ = v20
                } else {
                    var v21 uint16 = 902
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
        case 9223372036854775807:
            fmt.Println(v2, v2)
        case 178:
            if (492 < 279) {
                var v22 uint8 = 2
                _ = v22
            }
        }
        var v23 int16 = 42
        _ = fi17
        _ = v23
    }
    v24 := GenS3{F0: 2147483647}
    fmt.Println(v1, v2)
    for fi25 := 0; fi25 < 1; fi25++ {
        if (!false) {
            if (255 != 32767) {
                for fi26 := 0; fi26 < 9; fi26++ {
                    v2 -= 1e10
                    _ = fi26
                }
            }
        }
        switch fi25 {
        case -9223372036854775808:
            var v27 uint64 = 0
            _ = v27
        case -702:
            for fi28 := 0; fi28 < 1; fi28++ {
                v29 := GenS6{F0: -522}
                v30 := GenS2{F0: -2147483648, F1: 127}
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        _ = fi25
    }
    fmt.Println(v2)
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 int32 = 2147483647
        v33 := GenS2{F0: v32, F1: 127}
        if (42 >= 0) {
            var v34 uint8 = 0
            switch fi31 {
            case 128:
                v35 := v4["foo bar"]
                _ = v35
            case 127:
                v36 := []int{-9223372036854775808}
                _ = v36
            }
            for fi37 := 0; fi37 < 20; fi37++ {
                if ((-(fi37)) < fi31) {
                    v34 = (v34 ^ (v34 + (^(v34))))
                    v38 := v5[(fi37 / ((fi37) | 1))]
                    _ = v38
                } else {
                    fmt.Println(v32, fi31)
                    delete(v3, fi31)
                    fmt.Println(v32, v34, v2)
                }
                v39 := GenS4{F0: v32, F1: 545}
                switch ("\n\t" + ("foo bar" + "")) {
                case "\n\t":
                    v40 := GenS3{F0: -1}
                    _ = v40
                case "a":
                    v1 += (v1 - (v2 + 3.14159))
                }
                var v41 string = "\n\t"
                v32 = (v32 >> ((18446744073709551615) & 63))
                _ = fi37
                _ = v39
                _ = v41
            }
            _ = v34
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    var v42 bool = false
    if ("foo bar" == ("foo bar" + "\n\t")) {
        v2 = (float64(0.0) + (v1 * 3.14159))
    }
    switch (("a" + "") + ("\n\t" + "foo bar")) {
    case "foo bar":
        v43 := GenS1{F0: -252, F1: -2147483648}
        _ = v43
    case "":
        var v44 uint8 = 42
        _ = v44
    case "hello":
        var v45 string = "hello"
        _ = v45
    default:
        v46 := GenS4{F0: -2147483648, F1: 0}
        _ = v46
    }
    v47 := GenS5{F0: -2147483648, F1: 4294967295}
    for fi48 := 0; fi48 < 14; fi48++ {
        delete(v5, fi48)
        _ = fi48
    }
    for fi49 := 0; fi49 < 5; fi49++ {
        var v50 int = fi49
        for fi51 := 0; fi51 < 2; fi51++ {
            switch -438 {
            case 1:
                v42 = ((v42 && (false || false)) && (1000 > -600))
            case -2147483648:
                var v52 int8 = -128
                _ = v52
            default:
                var v53 bool = v42
                _ = v53
            }
            v54 := v4["\n\t"]
            _ = len(v5)
            _ = fi51
            _ = v54
        }
        _ = fi49
        _ = v50
    }
    v55 := GenS3{F0: 32767}
    v56 := []int{}
    for fi57 := 0; fi57 < 3; fi57++ {
        if len(v56) > 0 {
            v56[0] = (fi57 % ((fi57) | 1))
        }
        if (0 < 42) {
            var v58 uint8 = 2
            v59 := GenS7{F0: 127, F1: 2, F2: -1}
            _ = v58
            _ = v59
        }
        switch "\n\t" {
        case "":
            var v60 string = "a"
            _ = v60
        case "hello":
            for fi61 := 0; fi61 < 4; fi61++ {
                var v62 bool = v42
                v63 := GenS6{F0: 32767}
                var v64 string = ((("\n\t" + (("foo bar" + "foo bar") + "foo bar")) + "hello") + (("\n\t" + "a") + "a"))
                _ = fi61
                _ = v62
                _ = v63
                _ = v64
            }
        case "a":
            _ = len(v56)
        default:
            fmt.Println(v2)
        }
        _ = fi57
    }
    var v65 string = ("foo bar" + "hello")
    var v66 uint64 = 623
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v47
    _ = v5
    _ = v55
    _ = v56
    _ = v65
    _ = v66
}
