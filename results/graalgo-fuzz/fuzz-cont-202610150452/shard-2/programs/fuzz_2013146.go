package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint8; F2 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 string; F2 string }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 uint16; F1 uint; F2 uint32 }
func genFn0(p0 uint64, p1 bool, p2 int) uint16 {
    v7 := []int{(^(p2)), -858}
    if true {
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                p0 -= (p0 * (p0 * p0))
                if false {
                    v10 := GenS1{F0: 635, F1: fi8, F2: 2}
                    if p1 {
                        var v11 int32 = 2147483647
                        _ = v11
                    }
                    _ = v10
                }
                p2 = (fi9 / (((p2 / ((9223372036854775807) | 1))) | 1))
                v12 := GenS2{F0: 21}
                var v13 string = ""
                var v14 string = (((("\n\t" + v13) + ("\n\t" + ("hello" + v13))) + ("a" + v13)) + "a")
                _ = fi9
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi8
        }
        var v15 float32 = -2.25
        _ = v15
    }
    v16 := GenS3{F0: 1000}
    for fi17 := 0; fi17 < 9; fi17++ {
        if (true && true) {
            p0 = (p0 ^ (p0 & (^(p0))))
            switch ("\n\t" + ("a" + "a")) {
            case "\n\t":
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 int32 = 964
                    var v20 uint8 = 14
                    v21 := GenS2{F0: 0}
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
            case "hello":
                switch "a" {
                case "\n\t":
                    _ = len(v7)
                case "":
                    fmt.Println(p0)
                case "foo bar":
                    fmt.Println(p0, p1)
                }
            case "foo bar":
                switch ("a" + "foo bar") {
                case "hello":
                    for fi22 := 0; fi22 < 7; fi22++ {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        v23 := GenS4{F0: ""}
                        _ = fi22
                        _ = v23
                    }
                case "\n\t":
                    if len(v7) > 0 {
                        v7[0] = (-(fi17))
                    }
                case "":
                    var v24 int64 = 42
                    _ = v24
                }
            }
        }
        var v25 int64 = 9223372036854775807
        p1 = (p1 && ((p2 % ((int(p0)) | 1)) > p2))
        _ = fi17
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 742
    v3 := []string{}
    v4 := []string{"a", ("a" + ""), "\n\t"}
    v5 := []uint32{(v2 * v2)}
    var v6 uint64 = 147
    if (false && false) {
        for fi26 := 0; fi26 < 6; fi26++ {
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := v1[v2]
                switch 0 {
                case 274:
                    switch "\n\t" {
                    case "":
                        var v29 uint16 = 65535
                        _ = v29
                    case "\n\t":
                        fmt.Println(v28, v2)
                    case "foo bar":
                        v6 += (^(v6))
                    default:
                        switch "\n\t" {
                        case "":
                            switch (fi26 + (fi27 - 860)) {
                            case 4294967295:
                                var v30 float64 = float64(-2.25)
                                _ = v30
                            default:
                                v31 := GenS2{F0: int8(fi26)}
                                _ = v31
                            }
                            delete(v1, (v2 | 475))
                        case "a":
                            v28 = v28
                        case "hello":
                            var v32 int16 = -569
                            _ = v32
                        }
                        v1[439] = (274 >= 0)
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                        var v33 int32 = 2147483647
                        _ = v33
                    }
                case 128:
                    v34 := GenS5{F0: 32768, F1: 0, F2: 1}
                    _ = v34
                case -9223372036854775808:
                    if v28 {
                        var v35 float32 = 1.5
                        if len(v3) > 0 {
                            v3[0] = (("" + "hello") + ("a" + "a"))
                        }
                        v36 := GenS5{F0: fi27, F1: 0, F2: 42}
                        _ = v35
                        _ = v36
                    }
                default:
                    switch (fi26 % ((fi27) | 1)) {
                    case 9:
                        var v37 float32 = 3.14159
                        _ = v37
                    case -2147483648:
                        var v38 bool = (42 <= 0)
                        _ = v38
                    case 32767:
                        v2 = v2
                    }
                }
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    } else {
        for fi39 := 0; fi39 < 3; fi39++ {
            for fi40 := 0; fi40 < 3; fi40++ {
                var v41 int64 = -9223372036854775808
                fmt.Println(fi40)
                _ = fi40
                _ = v41
            }
            v6 += (v6 - 4294967295)
            var v42 int = fi39
            _ = fi39
            _ = v42
        }
        v6 -= (v6 + (v6 ^ (v6 & v6)))
    }
    v43 := GenS6{F0: true}
    for fi44 := 0; fi44 < 3; fi44++ {
        if true {
            if ((1 > -770) && ((0.0 + float64(0.0)) <= 3.14159)) {
                switch (^(fi44)) {
                case 996:
                    if len(v5) > 0 {
                        v5[0] = 421
                    }
                }
                v2 = (v2 / (((v2 * 1)) | 1))
                switch "\n\t" {
                case "hello":
                    var v45 int64 = 808
                    _ = v45
                case "":
                    switch fi44 {
                    case -2147483648:
                        for fi46 := 0; fi46 < 8; fi46++ {
                            fmt.Println(fi44)
                            if false {
                                v2 -= 2
                                _ = len(v5)
                            } else {
                                var v47 int32 = 1
                                v48 := v1[(v2 >> ((255) & 63))]
                                if ((-183 != 1) || (!true)) {
                                    if len(v5) > 0 {
                                        v5[0] = v2
                                    }
                                }
                                _ = v47
                                _ = v48
                            }
                            if len(v4) > 0 {
                                v4[0] = ("" + ((("\n\t" + ("\n\t" + (("foo bar" + ("hello" + "")) + "\n\t"))) + (("foo bar" + "hello") + "a")) + "a"))
                            }
                            _ = fi46
                        }
                    default:
                        for fi49 := 0; fi49 < 19; fi49++ {
                            if (!(-2.25 > ((0.0 + (((0.1 - -2.25) * 1.5) * -2.25)) / ((3.14159 + 1.5) + (float64(0.1) + 0.1))))) {
                                fmt.Println(v2)
                                v6 -= v6
                            } else {
                                v50 := GenS7{F0: 246, F1: "\n\t", F2: "foo bar"}
                                _ = v50
                            }
                            v51 := GenS8{F0: 2, F1: 700}
                            fmt.Println(v2)
                            _ = fi49
                            _ = v51
                        }
                    }
                default:
                    var v52 string = "hello"
                    _ = v52
                }
            }
            v2 = 0
            fmt.Println(v6, v6)
        }
        _ = fi44
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v6, v2)
    case "":
        fmt.Println(v6, v6)
    case "\n\t":
        fmt.Println(v6)
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v53 bool = false
    v54 := GenS9{F0: 1, F1: 255, F2: v2}
    var v55 uint32 = v2
    v56 := GenS4{F0: "\n\t"}
    v53 = v53
    v53 = false
    var v57 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v6
}
