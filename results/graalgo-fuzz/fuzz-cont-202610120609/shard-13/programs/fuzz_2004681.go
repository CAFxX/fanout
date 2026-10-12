package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint32; F1 int16; F2 int64 }
type GenS3 struct { F0 uint8; F1 int; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 uint8; F1 string; F2 uint64 }
type GenS9 struct { F0 int }
func genFn0(p0 string) int32 {
    var v4 int16 = 1
    v5 := GenS1{F0: 1, F1: p0}
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint64 = 1
        for fi8 := 0; fi8 < 17; fi8++ {
            p0 = p0
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    return 306
}
func genFn1(p0 uint16, p1 uint8, p2 int32) int32 {
    fmt.Println(p1, p2)
    var v9 bool = (2 > uint32(p1))
    if v9 {
        for fi10 := 0; fi10 < 12; fi10++ {
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 uint16 = 2
                fmt.Println(fi11, fi10)
                p1 -= p1
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p2 + p2)
}

func main() {
    var v1 float32 = (float32(1.5) / 1.5)
    var v2 int32 = 873
    v3 := make(map[uint32]int)
    v13 := GenS1{F0: 42, F1: ("\n\t" + ((("" + "\n\t") + ("hello" + "")) + "foo bar"))}
    for fi14 := 0; fi14 < 10; fi14++ {
        if true {
            switch "hello" {
            case "foo bar":
                v2 = (v2 << ((0) & 63))
            }
        } else {
            var v15 string = ""
            v16 := GenS1{F0: 480, F1: ""}
            v2 = v2
            _ = v15
            _ = v16
        }
        v17 := []int{(fi14 - (^(fi14))), (^(fi14)), fi14}
        var v18 bool = true
        _ = fi14
        _ = v17
        _ = v18
    }
    switch "\n\t" {
    case "foo bar":
        v19 := v3[42]
        _ = v19
    case "\n\t":
        v20 := GenS2{F0: 0, F1: 32767, F2: 847}
        _ = v20
    case "a":
        if true {
            var v21 int = 4294967295
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 bool = true
                _ = fi22
                _ = v23
            }
            _ = v21
        } else {
            v2 -= -2147483648
            v24 := []int{-218, -934, 4294967295, 0}
            if ((!false) && true) {
                var v25 float32 = float32(((0.0 / 1.0) * (float64(v1) + 0.1)))
                _ = len(v24)
                _ = v25
            } else {
                v26 := GenS3{F0: 128, F1: -797, F2: "a"}
                switch -1 {
                case -1:
                    v2 -= (^(v2))
                case 9223372036854775807:
                    delete(v3, 2147483647)
                case -423:
                    for fi27 := 0; fi27 < 5; fi27++ {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            fmt.Println(fi27, v2, v1)
                            _ = fi28
                        }
                        v29 := GenS4{F0: 291}
                        v1 = (3.14159 + (3.14159 * (1.5 + (v1 * v1))))
                        _ = fi27
                        _ = v29
                    }
                }
                if false {
                    if true {
                        for fi30 := 0; fi30 < 12; fi30++ {
                            var v31 uint8 = 27
                            _ = fi30
                            _ = v31
                        }
                    } else {
                        switch (("hello" + ("foo bar" + "")) + ((("\n\t" + "foo bar") + "") + (("\n\t" + "\n\t") + "a"))) {
                        case "hello":
                            v2 += v2
                        case "\n\t":
                            switch "\n\t" {
                            case "":
                                var v32 int64 = 951
                                _ = v32
                            }
                        case "foo bar":
                            v33 := v3[65535]
                            _ = v33
                        }
                        switch 32768 {
                        case 418:
                            var v34 float32 = 1.5
                            _ = v34
                        }
                    }
                    v35 := GenS2{F0: 800, F1: 764, F2: -587}
                    switch "\n\t" {
                    case "":
                        v2 = -32
                    case "foo bar":
                        v1 = 1e10
                    }
                    _ = v35
                } else {
                    var v36 float64 = 1.5
                    if (!(-119 >= 100)) {
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        _ = len(v24)
                    } else {
                        switch ((("a" + ("" + "")) + "foo bar") + ("hello" + "foo bar")) {
                        case "hello":
                            fmt.Println(v36, v36)
                        case "foo bar":
                            _ = len(v3)
                        default:
                            v1 += v1
                        }
                        fmt.Println(v2)
                    }
                    _ = v36
                }
                _ = v26
            }
            _ = v24
        }
    default:
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 uint32 = 42
            _ = fi37
            _ = v38
        }
    }
    v2 = 42
    v39 := []int{-2147483648, -1}
    for fi40 := 0; fi40 < 9; fi40++ {
        if true {
            fmt.Println(v2, fi40, fi40)
        } else {
            fmt.Println(v1)
            v41 := GenS5{F0: 336, F1: "a"}
            _ = v41
        }
        v2 = 0
        for fi42 := 0; fi42 < 19; fi42++ {
            switch "\n\t" {
            case "":
                v43 := GenS5{F0: 2, F1: ""}
                _ = v43
            case "\n\t":
                var v44 bool = (20 == 7)
                var v45 uint64 = 4294967295
                if len(v39) >= 2 {
                    v39 = v39[1:len(v39)]
                }
                _ = v44
                _ = v45
            case "foo bar":
                v2 += (v2 & (v2 << ((322) & 63)))
            }
            fmt.Println(v2, v1, fi42)
            _ = fi42
        }
        _ = fi40
    }
    fmt.Println(v1, v1)
    switch ("a" + ("\n\t" + "")) {
    case "hello":
        for fi46 := 0; fi46 < 12; fi46++ {
            v2 += v2
            fmt.Println(fi46, v2)
            for fi47 := 0; fi47 < 11; fi47++ {
                switch "\n\t" {
                case "\n\t":
                    switch (fi47 ^ -96) {
                    case -2147483648:
                        var v48 uint32 = 2147483648
                        _ = v48
                    case -9223372036854775808:
                        var v49 int32 = genFn0(("" + ("a" + "\n\t")))
                        _ = v49
                    case 164:
                        v1 = v1
                    }
                case "foo bar":
                    v50 := GenS1{F0: 65535, F1: "hello"}
                    _ = v50
                }
                if false {
                    v39 = append(v39, fi47)
                } else {
                    if false {
                        v51 := GenS4{F0: 447}
                        _ = v51
                    } else {
                        var v52 int32 = v2
                        _ = len(v3)
                        var v53 int64 = 42
                        _ = v52
                        _ = v53
                    }
                }
                if false {
                    switch 32767 {
                    case 2147483647:
                        _ = len(v39)
                    default:
                        v54 := GenS2{F0: 2147483648, F1: 32767, F2: 42}
                        _ = v54
                    }
                    var v55 uint8 = 128
                    if len(v39) > 0 {
                        v39[0] = fi47
                    }
                    _ = v55
                } else {
                    v1 -= v1
                    v56 := GenS6{F0: 42}
                    _ = v56
                }
                v57 := GenS7{F0: (^(fi46)), F1: 882}
                _ = fi47
                _ = v57
            }
            _ = fi46
        }
    case "a":
        var v58 uint16 = 65535
        _ = v58
    default:
        if (255 != 0) {
            var v59 float64 = 0.0
            _ = v59
        } else {
            if (0 <= 251) {
                fmt.Println(v1, v1)
                v2 += (v2 * 0)
            } else {
                fmt.Println(v1)
                var v60 uint64 = 673
                _ = v60
            }
            if true {
                v3[1] = -2147483648
            }
            switch ("hello" + ("hello" + "\n\t")) {
            case "foo bar":
                switch "\n\t" {
                case "":
                    fmt.Println(v2, v1)
                }
            case "hello":
                v2 -= (v2 / (((-(v2))) | 1))
            case "":
                var v61 uint32 = 1
                _ = v61
            }
        }
    }
    fmt.Println(v2, v1)
    var v62 int8 = 42
    for fi63 := 0; fi63 < 11; fi63++ {
        for fi64 := 0; fi64 < 17; fi64++ {
            v1 = 1e10
            _ = fi64
        }
        if false {
            v2 -= v2
        } else {
            v65 := GenS8{F0: 195, F1: "\n\t", F2: 2}
            _ = v65
        }
        _ = fi63
    }
    v66 := GenS9{F0: 4294967295}
    delete(v3, 4294967295)
    var v67 float32 = v1
    fmt.Println(v1, v2)
    if len(v39) >= 2 {
        v39 = v39[1:len(v39)]
    }
    switch 840 {
    case 2147483647:
        switch "foo bar" {
        case "foo bar":
            v68 := GenS8{F0: 1, F1: ("foo bar" + "foo bar"), F2: 42}
            _ = v68
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v62:", v62)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v39
    _ = v62
    _ = v66
    _ = v67
}
