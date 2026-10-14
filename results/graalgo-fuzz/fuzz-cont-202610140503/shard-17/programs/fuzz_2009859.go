package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int; F1 int32; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 int32; F2 uint64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 uint; F2 string }
func genFn0(p0 bool, p1 int16) uint16 {
    p0 = p0
    p1 = p1
    v5 := GenS1{F0: 436, F1: 796}
    _ = p0
    _ = p1
    _ = v5
    return 932
}

func main() {
    var v1 int8 = -26
    var v2 float64 = 1.5
    var v3 int32 = -2147483648
    var v4 int8 = v1
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int32 = v3
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(v4, v2, fi6)
            fmt.Println(v7, fi6, v2)
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        switch "foo bar" {
        case "a":
            v10 := []int{fi9}
            _ = v10
        case "hello":
            for fi11 := 0; fi11 < 11; fi11++ {
                if false {
                    switch (("" + "\n\t") + "foo bar") {
                    case "hello":
                        v12 := []int{(fi9 + fi11)}
                        _ = v12
                    case "\n\t":
                        v3 -= v3
                    case "foo bar":
                        fmt.Println(v1, v2)
                        v2 -= float64((1.5 / (-2.25 - 1.5)))
                        var v13 uint64 = 234
                        _ = v13
                    }
                } else {
                    switch 2147483647 {
                    case 325:
                        v2 = 0.0
                    default:
                        if (((v3 % ((v3) | 1)) == (v3 & v3)) || false) {
                            if ("hello" == "\n\t") {
                                v14 := GenS2{F0: v3, F1: 42}
                                if (v3 >= v3) {
                                    var v15 bool = (!false)
                                    switch (("hello" + "") + ("a" + "foo bar")) {
                                    case "":
                                        v15 = v15
                                    default:
                                        var v16 float64 = v2
                                        if (574 < -68) {
                                            v17 := []int{(fi11 ^ (^(fi11))), 1, int(v3), 42}
                                            if len(v17) > 0 {
                                                v17[0] = int(v3)
                                            }
                                            _ = v17
                                        } else {
                                            var v18 string = ""
                                            fmt.Println(v15, v16)
                                            _ = v18
                                        }
                                        var v19 uint = 1
                                        v20 := GenS3{F0: fi9, F1: 1, F2: (v3 % ((2147483647) | 1))}
                                        _ = v16
                                        _ = v19
                                        _ = v20
                                    }
                                    var v21 int = (fi11 & fi9)
                                    _ = v15
                                    _ = v21
                                } else {
                                    var v22 int64 = -9223372036854775808
                                    v4 = (v4 - -128)
                                    _ = v22
                                }
                                v1 -= (v1 ^ v1)
                                var v23 uint16 = 2
                                _ = v14
                                _ = v23
                            }
                            fmt.Println(v2)
                        }
                    }
                    v24 := GenS4{F0: (v3 ^ (^(v3)))}
                    fmt.Println(v3, v1)
                    _ = v24
                }
                v25 := []int{fi9}
                if true {
                    if (genFn0(true, -32768) <= 255) {
                        v26 := GenS2{F0: 1, F1: 65535}
                        fmt.Println(fi11, v3, v2)
                        fmt.Println(v1, v3)
                        _ = v26
                    } else {
                        var v27 uint = 255
                        if false {
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                        } else {
                            var v28 uint32 = 4294967295
                            var v29 float64 = v2
                            var v30 int32 = (v3 >> (((v27 << ((1) & 63))) & 63))
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        switch -329 {
                        case -907:
                            if len(v25) > 0 {
                                v25[0] = (fi11 % ((fi11) | 1))
                            }
                        case 127:
                            if false {
                                var v31 float32 = 1.5
                                var v32 float64 = -2.25
                                v1 -= (v1 & (v4 & v1))
                                _ = v31
                                _ = v32
                            } else {
                                fmt.Println(fi11)
                            }
                        }
                        _ = v27
                    }
                    if true {
                        v33 := GenS5{F0: fi9, F1: false}
                        _ = v33
                    }
                    fmt.Println(v4, fi9, v4)
                }
                _ = fi11
                _ = v25
            }
        default:
            v3 = v3
        }
        switch ("hello" + ("" + ("" + "foo bar"))) {
        case "\n\t":
            if false {
                v4 = v1
                v34 := []int{(fi9 >> ((4294967295) & 63)), 0, fi9, -91}
                v3 -= 0
                _ = v34
            }
        case "":
            switch (("" + "\n\t") + "a") {
            case "":
                v35 := GenS6{F0: 491}
                _ = v35
            case "hello":
                switch (^(fi9)) {
                case 2147483647:
                    v36 := []int{(fi9 ^ (fi9 * fi9)), -2147483648, (fi9 | fi9), int(v1)}
                    _ = v36
                default:
                    switch (fi9 & 1) {
                    case 280:
                        v1 -= (v1 / ((v1) | 1))
                    case -184:
                        v37 := []int{4294967295, 2147483647, (-(fi9))}
                        _ = v37
                    case -9223372036854775808:
                        var v38 uint64 = 154
                        _ = v38
                    default:
                        var v39 uint8 = 184
                        _ = v39
                    }
                }
            case "a":
                var v40 uint64 = 791
                _ = v40
            default:
                for fi41 := 0; fi41 < 12; fi41++ {
                    var v42 uint32 = 65535
                    v43 := []int{}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            }
        case "hello":
            for fi44 := 0; fi44 < 1; fi44++ {
                if true {
                    v4 = (v4 >> ((856) & 63))
                    fmt.Println(v3, fi9)
                    fmt.Println(v2, fi9, v1)
                }
                fmt.Println(fi44)
                var v45 int32 = (^(v3))
                _ = fi44
                _ = v45
            }
        }
        _ = fi9
    }
    var v46 bool = ((((v2 - 0.0) + v2) + 1.5) > 3.14159)
    if (!v46) {
        if v46 {
            for fi47 := 0; fi47 < 18; fi47++ {
                fmt.Println(v46, v4)
                v48 := GenS7{F0: -233}
                _ = fi47
                _ = v48
            }
        } else {
            v46 = (!false)
            if true {
                switch "\n\t" {
                case "foo bar":
                    for fi49 := 0; fi49 < 3; fi49++ {
                        v50 := GenS8{F0: (((!v46) || ("foo bar" < "a")) || (128 > 128)), F1: v3, F2: 701}
                        _ = fi49
                        _ = v50
                    }
                case "":
                    for fi51 := 0; fi51 < 20; fi51++ {
                        fmt.Println(v2)
                        _ = fi51
                    }
                default:
                    v3 = v3
                }
            }
        }
        var v52 uint64 = 1
        switch -1 {
        case 42:
            var v53 float32 = 1e10
            _ = v53
        case 9223372036854775807:
            for fi54 := 0; fi54 < 1; fi54++ {
                v3 = v3
                v1 += (^(v4))
                _ = fi54
            }
        case 32767:
            v3 += v3
        }
        _ = v52
    } else {
        v2 += 0.0
        v55 := GenS9{F0: 476}
        _ = v55
    }
    v1 += (v4 & -57)
    var v56 int64 = 1
    switch 680 {
    case 42:
        v57 := []int{-23}
        _ = v57
    }
    v4 = (v1 << ((1) & 63))
    v2 -= 0.1
    v1 = (v1 & (v4 ^ -1))
    if v46 {
        v58 := []int{128, -9223372036854775808, 914}
        _ = v58
    } else {
        if v46 {
            var v59 int8 = -65
            _ = v59
        }
        if (734 != 891) {
            v56 = v56
            v46 = false
        }
        v60 := []int{9223372036854775807}
        _ = v60
    }
    v2 = v2
    for fi61 := 0; fi61 < 20; fi61++ {
        var v62 uint64 = 18446744073709551615
        v4 = v1
        v63 := GenS10{F0: int8(v62), F1: 255, F2: (((("a" + ("a" + "a")) + "a") + "\n\t") + "")}
        _ = fi61
        _ = v62
        _ = v63
    }
    var v64 int64 = (v56 + 0)
    var v65 int8 = (v1 | v4)
    var v66 int64 = 9223372036854775807
    var v67 int32 = int32(v64)
    var v68 bool = v46
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v56:", v56)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v56
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
