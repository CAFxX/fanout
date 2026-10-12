package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 int; F1 int8; F2 bool }
func genFn0() bool {
    switch "hello" {
    case "a":
        v7 := GenS1{F0: 2147483647}
        _ = v7
    case "\n\t":
        var v8 int64 = 9223372036854775807
        _ = v8
    }
    if (-1 >= -1) {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 string = ("a" + "\n\t")
            _ = fi9
            _ = v10
        }
    }
    if true {
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := []int{fi11}
            var v13 int = (fi11 & (fi11 * fi11))
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := []int{0, 32768, 9223372036854775807, 520}
        if true {
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 uint8 = 255
                var v17 int = 1
                var v18 bool = true
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        _ = v14
    } else {
        var v19 uint32 = 2
        v20 := GenS1{F0: 714}
        _ = v19
        _ = v20
    }
    var v21 float32 = 0.0
    _ = v21
    return true
}
func genFn1(p0 bool) uint {
    var v22 uint16 = 255
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := GenS1{F0: 714}
        _ = fi23
        _ = v24
    }
    if p0 {
        v25 := []int{42, 0, 32768, 1}
        v25 = append(v25, -1)
        v22 = v22
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        if true {
            var v27 float32 = ((1.5 + float32((0.0 * 3.14159))) + (-2.25 + 1e10))
            var v28 uint8 = 42
            _ = v27
            _ = v28
        }
        _ = fi26
    }
    v22 = 1
    _ = p0
    _ = v22
    return 2
}

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 4294967295
    var v3 int64 = 1
    v4 := []string{"foo bar", ""}
    var v5 uint64 = v2
    var v6 float64 = v1
    var v29 float32 = 1e10
    var v30 uint8 = 172
    switch "foo bar" {
    case "\n\t":
        switch 32767 {
        case -907:
            if true {
                v31 := GenS2{F0: (0 < 689), F1: (^(v30))}
                for fi32 := 0; fi32 < 11; fi32++ {
                    v29 -= (v29 - 3.14159)
                    if true {
                        fmt.Println(v29)
                        fmt.Println(fi32)
                    }
                    _ = fi32
                }
                _ = v31
            }
        case 565:
            v1 += v6
        case 42:
            fmt.Println(v1, v30, v2)
        }
    case "foo bar":
        v5 = (v2 / (((v2 | v2)) | 1))
    }
    v33 := GenS3{F0: v2}
    if true {
        v4 = append(v4, "hello")
        switch "\n\t" {
        case "hello":
            v6 = 3.14159
        case "a":
            var v34 int64 = v3
            _ = v34
        case "":
            for fi35 := 0; fi35 < 12; fi35++ {
                var v36 uint8 = 1
                _ = fi35
                _ = v36
            }
        }
    } else {
        v4 = append(v4, "")
        if (100 <= -64) {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        } else {
            if true {
                fmt.Println(v6, v5, v3)
                var v37 float64 = (0.0 * float64((v29 - v29)))
                _ = v37
            } else {
                v29 = v29
                v2 = (v2 % ((v2) | 1))
                var v38 int16 = -308
                _ = v38
            }
        }
        v1 = float64(0.0)
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        for fi40 := 0; fi40 < 3; fi40++ {
            v41 := GenS3{F0: v2}
            v42 := GenS1{F0: (fi39 - (fi39 * 1))}
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = fi39
    }
    var v43 bool = true
    fmt.Println(v5)
    var v44 uint16 = 80
    for fi45 := 0; fi45 < 13; fi45++ {
        for fi46 := 0; fi46 < 8; fi46++ {
            var v47 uint = 600
            switch (fi45 << (((^(v47))) & 63)) {
            case -514:
                fmt.Println(v5)
            case -163:
                if (v3 == v3) {
                    v48 := GenS1{F0: (^(fi46))}
                    v4 = append(v4, ("\n\t" + "\n\t"))
                    _ = v48
                } else {
                    switch ("" + "a") {
                    case "":
                        if len(v4) > 0 {
                            v4[0] = ("hello" + (("a" + "") + ""))
                        }
                    case "foo bar":
                        if v43 {
                            switch "hello" {
                            case "":
                                v43 = ((225 == 2147483647) || (((v29 > (v29 * 0.0)) || false) && (((!true) || (42 > v30)) || v43)))
                            case "\n\t":
                                var v49 string = ("foo bar" + (("" + "") + "a"))
                                _ = v49
                            case "a":
                                var v50 int64 = 42
                                _ = v50
                            default:
                                var v51 int64 = (v3 + (v3 % (((v3 - v3)) | 1)))
                                _ = v51
                            }
                            v52 := GenS1{F0: 128}
                            _ = v52
                        }
                    case "a":
                        v29 = v29
                    }
                    v53 := GenS4{F0: "foo bar", F1: 0}
                    _ = v53
                }
            default:
                fmt.Println(v2, fi45)
                switch "foo bar" {
                case "":
                    fmt.Println(v2, v30, fi46)
                }
            }
            _ = fi46
            _ = v47
        }
        var v54 float32 = v29
        _ = fi45
        _ = v54
    }
    _ = len(v4)
    v2 += 2
    var v55 uint64 = (v5 << ((genFn1((1 <= 0))) & 63))
    switch "hello" {
    case "":
        if v43 {
            var v56 int = 2147483647
            _ = v56
        } else {
            var v57 string = "foo bar"
            v58 := GenS3{F0: (v2 * v55)}
            _ = v57
            _ = v58
        }
    case "foo bar":
        _ = len(v4)
    default:
        var v59 float32 = v29
        _ = v59
    }
    v3 = 42
    fmt.Println(v3, v5, v30)
    for fi60 := 0; fi60 < 9; fi60++ {
        if v43 {
            var v61 int8 = -2
            _ = v61
        } else {
            if (true || v43) {
                if (true || ("" != ("a" + (((("a" + "\n\t") + "a") + ("\n\t" + "\n\t")) + "foo bar")))) {
                    _ = len(v4)
                } else {
                    if (!false) {
                        var v62 int32 = -2147483648
                        _ = v62
                    }
                    var v63 bool = true
                    if false {
                        v43 = v43
                        if len(v4) > 0 {
                            v4[0] = ("a" + "\n\t")
                        }
                        v64 := GenS5{F0: 42, F1: 92, F2: v30}
                        _ = v64
                    }
                    _ = v63
                }
                v65 := GenS6{F0: uint16(v2), F1: -1}
                _ = v65
            } else {
                v44 += 1
                _ = len(v4)
                if len(v4) > 0 {
                    v4[0] = ""
                }
            }
            v66 := GenS7{F0: (fi60 / ((fi60) | 1)), F1: 50, F2: (fi60 <= (fi60 >> ((18446744073709551615) & 63)))}
            _ = v66
        }
        _ = len(v4)
        _ = fi60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v55
    _ = v6
}
