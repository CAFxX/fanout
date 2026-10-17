package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 int; F2 uint64 }
type GenS6 struct { F0 int; F1 uint32; F2 int16 }
func genFn0(p0 int8, p1 int, p2 string) uint {
    v4 := GenS1{F0: -9223372036854775808, F1: (p1 & (^(p1)))}
    for fi5 := 0; fi5 < 16; fi5++ {
        if (42 != 424) {
            switch fi5 {
            case -2147483648:
                for fi6 := 0; fi6 < 5; fi6++ {
                    var v7 uint = 65535
                    _ = fi6
                    _ = v7
                }
            case -389:
                v8 := GenS1{F0: 9223372036854775807, F1: p1}
                _ = v8
            default:
                for fi9 := 0; fi9 < 8; fi9++ {
                    p0 = (p0 & 103)
                    v10 := GenS1{F0: 0, F1: 1}
                    _ = fi9
                    _ = v10
                }
            }
            if true {
                fmt.Println(p0)
            }
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := []int{4294967295, p1, p1}
                p1 += p1
                _ = len(v12)
                _ = fi11
                _ = v12
            }
        } else {
            var v13 uint = 2
            var v14 int = -111
            _ = v13
            _ = v14
        }
        if false {
            p2 = ((p2 + "\n\t") + p2)
            var v15 int8 = p0
            _ = v15
        } else {
            switch (p2 + ("a" + "foo bar")) {
            case "foo bar":
                var v16 float64 = -2.25
                _ = v16
            default:
                fmt.Println(p1)
            }
            var v17 uint = 65535
            _ = v17
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := GenS2{F0: 65535}
            p0 -= p0
            var v20 int32 = 201
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi5
    }
    var v21 uint8 = 245
    fmt.Println(p0, p0)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v4
    return 0
}
func genFn1() int16 {
    v22 := []int{128, -292, 128, 4294967295}
    var v23 int = -1
    v24 := GenS3{F0: 1, F1: 898, F2: 2}
    v25 := GenS4{F0: 32768}
    for fi26 := 0; fi26 < 7; fi26++ {
        switch ("" + "hello") {
        case "\n\t":
            if (true || ((0 != 261) && true)) {
                v23 = -9223372036854775808
                var v27 float64 = (((0.0 - 1.5) / 3.14159) - (1.5 - 1e10))
                _ = v27
            }
        case "hello":
            _ = len(v22)
        case "":
            for fi28 := 0; fi28 < 2; fi28++ {
                var v29 uint16 = 1
                v30 := GenS1{F0: 42, F1: (^(fi28))}
                v29 += v29
                v31 := GenS5{F0: 127, F1: (fi28 ^ (fi26 % (((-(fi28))) | 1))), F2: 4294967295}
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
        default:
            fmt.Println(fi26)
        }
        for fi32 := 0; fi32 < 5; fi32++ {
            v23 = fi26
            fmt.Println(v23, fi32)
            if (0 > 2) {
                if len(v22) > 0 {
                    v22[0] = 127
                }
            } else {
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                var v33 float32 = -2.25
                _ = v33
            }
            _ = fi32
        }
        _ = fi26
    }
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    return -625
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    v3 := []int{}
    if false {
        var v34 int64 = 1
        _ = v34
    }
    if true {
        var v35 uint = 65535
        fmt.Println(v35)
        v36 := GenS6{F0: 32768, F1: 4294967295, F2: 171}
        _ = v35
        _ = v36
    } else {
        var v37 int = -486
        _ = v37
    }
    if false {
        var v38 bool = false
        _ = v38
    }
    switch 9223372036854775807 {
    case 2147483647:
        if (!true) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    if ((("" + ("\n\t" + "a")) < "") && (434 < 2147483648)) {
        var v39 uint8 = 128
        var v40 int16 = 0
        v41 := v1[65535]
        _ = v39
        _ = v40
        _ = v41
    }
    var v42 string = ""
    switch 9223372036854775807 {
    case 4294967295:
        switch v42 {
        case "\n\t":
            delete(v1, 65535)
        case "":
            fmt.Println(v42)
        case "foo bar":
            switch "\n\t" {
            case "a":
                v42 = (v42 + (v42 + "foo bar"))
            default:
                if true {
                    var v43 int64 = -154
                    v43 = v43
                    _ = v43
                } else {
                    v42 = v42
                    switch (("" + "hello") + v42) {
                    case "foo bar":
                        fmt.Println(v42)
                    case "":
                        var v44 float32 = (0.1 - 0.1)
                        _ = v44
                    }
                }
            }
        default:
            switch v42 {
            case "foo bar":
                v42 = v42
            case "":
                var v45 int64 = 344
                _ = v45
            default:
                fmt.Println(v42)
            }
        }
    case 42:
        if (1 <= 255) {
            v46 := v2[v42]
            v42 = "\n\t"
            _ = v46
        }
    default:
        v47 := GenS6{F0: -9223372036854775808, F1: 243, F2: 913}
        _ = v47
    }
    fmt.Println(v42)
    var v48 uint = 520
    v3 = append(v3, -104)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v48
}
