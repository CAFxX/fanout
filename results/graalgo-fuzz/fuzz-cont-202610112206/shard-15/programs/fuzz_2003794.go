package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = -1
    v3 := []int{0, 9223372036854775807, 9223372036854775807, -9223372036854775808}
    var v4 uint8 = 2
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(fi5, v4, v4)
        _ = fi5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        if false {
            v2 += (v2 - v2)
            var v7 float32 = 3.14159
            v2 = v2
            _ = v7
        } else {
            if len(v3) > 0 {
                v3[0] = (-(fi6))
            }
        }
        var v8 uint = uint(v4)
        _ = fi6
        _ = v8
    }
    var v9 uint64 = 4294967295
    switch ("\n\t" + "a") {
    case "a":
        switch ("a" + "hello") {
        case "hello":
            if true {
                var v10 int64 = 9223372036854775807
                for fi11 := 0; fi11 < 2; fi11++ {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        var v13 bool = ((-1 == 100) && false)
                        _ = fi12
                        _ = v13
                    }
                    var v14 uint8 = uint8(v9)
                    if true {
                        var v15 float32 = 3.14159
                        _ = v15
                    }
                    _ = fi11
                    _ = v14
                }
                if len(v3) > 0 {
                    v3[0] = 128
                }
                _ = v10
            }
        case "":
            var v16 int16 = 32767
            _ = v16
        case "foo bar":
            v17 := GenS1{F0: (^(v4)), F1: 42}
            _ = v17
        default:
            for fi18 := 0; fi18 < 17; fi18++ {
                _ = len(v1)
                v19 := GenS1{F0: (v4 | (^(v4))), F1: 351}
                for fi20 := 0; fi20 < 17; fi20++ {
                    if ((v4 / (((v4 ^ (v4 << ((42) & 63)))) | 1)) != v4) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    if (!(v9 < v9)) {
                        if true {
                            v21 := GenS1{F0: v4, F1: 2147483648}
                            v4 += 2
                            if (20 != 0) {
                                v4 = (v4 / (((v4 << ((2) & 63))) | 1))
                                if (int16(v9) <= -1) {
                                    v9 = 0
                                    var v22 int32 = -755
                                    var v23 bool = false
                                    _ = v22
                                    _ = v23
                                }
                            }
                            switch 9223372036854775807 {
                            case 128:
                                fmt.Println(v4, v4, fi20)
                            default:
                                var v24 uint32 = 903
                                _ = v24
                            }
                            _ = v21
                        }
                        switch (fi20 / ((fi18) | 1)) {
                        case 128:
                            var v25 string = ""
                            _ = v25
                        default:
                            if (!false) {
                                if (!(fi20 <= (fi18 - (^(fi20))))) {
                                    if (65535 >= 934) {
                                        var v26 bool = (((v4 & (v4 >> ((42) & 63))) == v4) || (0.1 == 0.1))
                                        _ = v26
                                    } else {
                                        var v27 bool = (false && (!true))
                                        _ = v27
                                    }
                                } else {
                                    if ((((-2.25 / 1.5) * (0.1 / ((1e10 / 1.0) * (0.0 - -2.25)))) - float64(-2.25)) != 1e10) {
                                        v3 = append(v3, fi20)
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    fmt.Println(v9)
                                    if (!false) {
                                        switch fi20 {
                                        case 2147483647:
                                            v4 += 107
                                        }
                                        v28 := v1[""]
                                        switch "" {
                                        case "\n\t":
                                            delete(v1, "")
                                        case "a":
                                            fmt.Println(v9)
                                        }
                                        _ = v28
                                    }
                                }
                                v4 = (v4 & 42)
                                var v29 uint8 = (v4 | (v4 - 255))
                                _ = v29
                            }
                        }
                        switch "foo bar" {
                        case "a":
                            var v30 int = fi20
                            _ = v30
                        case "\n\t":
                            var v31 uint = 897
                            _ = v31
                        case "":
                            v32 := GenS2{F0: 65535, F1: 757, F2: -62}
                            v4 = uint8(v9)
                            v4 += 0
                            _ = v32
                        }
                    } else {
                        v3 = append(v3, (fi18 - fi20))
                        v33 := GenS1{F0: v4, F1: 42}
                        _ = v33
                    }
                    var v34 uint64 = 827
                    fmt.Println(fi18, v9)
                    var v35 string = ("foo bar" + "\n\t")
                    var v36 uint64 = (v9 | (v9 | (^(v9))))
                    _ = fi20
                    _ = v34
                    _ = v35
                    _ = v36
                }
                _ = fi18
                _ = v19
            }
        }
    case "\n\t":
        var v37 uint32 = 1
        _ = v37
    default:
        var v38 string = "foo bar"
        _ = v38
    }
    if (3.14159 != ((1.5 * (-2.25 / 1.0)) * -2.25)) {
        var v39 float64 = 1.5
        if true {
            var v40 uint16 = 42
            v2 = (v2 & v2)
            _ = v40
        }
        _ = v39
    }
    v2 += v2
    switch 9223372036854775807 {
    case 42:
        var v41 float32 = 3.14159
        _ = v41
    case 703:
        var v42 string = ""
        _ = v42
    case 128:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        v4 -= (v4 >> ((65535) & 63))
    }
    v43 := GenS3{F0: 255}
    fmt.Println(v9)
    v2 += (v2 % ((0) | 1))
    for fi44 := 0; fi44 < 2; fi44++ {
        v4 -= v4
        v4 = 42
        v4 += 42
        _ = fi44
    }
    v9 = v9
    v1[""] = 1
    if (605 > 32767) {
        for fi45 := 0; fi45 < 17; fi45++ {
            v4 -= 41
            fmt.Println(v2, v4)
            v46 := v1["foo bar"]
            _ = fi45
            _ = v46
        }
        v47 := GenS1{F0: v4, F1: 489}
        switch 128 {
        case 3:
            switch 32768 {
            case -1:
                v48 := GenS2{F0: 693, F1: 887, F2: (v2 % ((-830) | 1))}
                _ = v48
            case -717:
                delete(v1, ("hello" + "a"))
            default:
                v2 -= v2
            }
        }
        _ = v47
    } else {
        for fi49 := 0; fi49 < 2; fi49++ {
            if (!("a" <= "foo bar")) {
                v2 = (v2 & 0)
                if true {
                    if false {
                        v3 = append(v3, (fi49 << ((4294967295) & 63)))
                    }
                    var v50 int16 = 1000
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 1; fi51++ {
                        v9 += v9
                        var v52 int16 = 0
                        var v53 uint32 = 2147483647
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    var v54 uint64 = (^(v9))
                    _ = v54
                }
            }
            _ = fi49
        }
        v55 := GenS1{F0: (v4 ^ v4), F1: 42}
        switch ("\n\t" + ("a" + "foo bar")) {
        case "":
            fmt.Println(v4, v2)
        case "foo bar":
            switch "hello" {
            case "\n\t":
                v4 = (v4 ^ 1)
            case "":
                var v56 bool = (-9223372036854775808 == 42)
                _ = v56
            case "a":
                fmt.Println(v9, v9)
            }
        }
        _ = v55
    }
    fmt.Println(v9)
    var v57 uint64 = 18446744073709551615
    v58 := GenS1{F0: 0, F1: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v57
    _ = v58
    _ = v9
}
