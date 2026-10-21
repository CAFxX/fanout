package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 string; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 bool }

func main() {
    var v1 uint = 4294967295
    var v2 int32 = 0
    v3 := []string{""}
    var v4 string = "a"
    var v5 uint64 = 18446744073709551615
    v6 := make(map[string]string)
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int8 = 100
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        if (true || (2 != 42)) {
            for fi10 := 0; fi10 < 3; fi10++ {
                fmt.Println(v5, fi10, v4)
                _ = fi10
            }
        } else {
            for fi11 := 0; fi11 < 8; fi11++ {
                if true {
                    var v12 float64 = (3.14159 + 1e10)
                    var v13 uint32 = 0
                    var v14 uint32 = v13
                    fmt.Println(fi9)
                    var v15 uint32 = v14
                    if ((((v12 - 1e10) <= 1e10) || true) && (v5 == v5)) {
                        var v16 uint8 = 39
                        if true {
                            var v17 int = fi11
                            switch (fi9 | (fi9 * (-(fi9)))) {
                            case 4294967295:
                                v17 = (fi11 + (fi9 ^ fi11))
                            case -737:
                                var v18 uint32 = (^(v13))
                                _ = v18
                            case 644:
                                fmt.Println(v2, v15, v17)
                            }
                            _ = v17
                        }
                        v15 += (v14 * 682)
                        _ = v16
                    } else {
                        var v19 uint = (v1 ^ v1)
                        delete(v6, ("\n\t" + v4))
                        var v20 bool = true
                        _ = v19
                        _ = v20
                    }
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                v1 -= (v1 << ((v1) & 63))
                switch (fi11 / ((128) | 1)) {
                case 4294967295:
                    var v21 int8 = -128
                    _ = v21
                case 127:
                    if (true && false) {
                        v22 := GenS1{F0: 318}
                        v1 = 255
                        _ = v22
                    }
                    var v23 int32 = (v2 ^ (v2 | v2))
                    v1 = 42
                    var v24 float32 = (-2.25 * float32(1e10))
                    fmt.Println(v24, v2)
                    var v25 float64 = 0.1
                    var v26 int8 = 40
                    _ = v23
                    _ = v24
                    _ = v25
                    _ = v26
                case 258:
                    v5 += v5
                }
                _ = fi11
            }
        }
        var v27 int = (fi9 >> ((v1) & 63))
        switch v4 {
        case "hello":
            var v28 uint32 = 42
            _ = v28
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "a":
            v4 = (v4 + v4)
        default:
            switch 32767 {
            case -2147483648:
                var v29 int16 = 42
                _ = v29
            case -9223372036854775808:
                v30 := GenS1{F0: 42}
                _ = v30
            default:
                if false {
                    fmt.Println(v1, v5, v27)
                }
            }
        }
        _ = fi9
        _ = v27
    }
    v6[(v4 + ((v4 + ("" + v4)) + "\n\t"))] = ""
    var v31 int32 = v2
    v32 := GenS2{F0: 42, F1: "a", F2: 255}
    var v33 int64 = 48
    for fi34 := 0; fi34 < 4; fi34++ {
        v35 := GenS1{F0: 32767}
        if ("hello" > "hello") {
            switch "" {
            case "foo bar":
                var v36 int16 = 1000
                _ = v36
            }
        }
        switch "" {
        case "":
            v1 = v1
        }
        _ = fi34
        _ = v35
    }
    if false {
        switch v4 {
        case "foo bar":
            fmt.Println(v5)
        default:
            var v37 int32 = v31
            _ = v37
        }
        for fi38 := 0; fi38 < 9; fi38++ {
            if len(v3) > 0 {
                v3[0] = (v4 + v4)
            }
            switch (v4 + v4) {
            case "hello":
                v39 := GenS3{F0: 0, F1: 2}
                _ = v39
            case "":
                v4 = v4
            }
            fmt.Println(v31, v1)
            _ = fi38
        }
        v40 := v6[v4]
        _ = v40
    } else {
        v41 := GenS2{F0: 990, F1: "\n\t", F2: 2}
        if (0 > 1000) {
            for fi42 := 0; fi42 < 9; fi42++ {
                var v43 uint32 = 0
                _ = fi42
                _ = v43
            }
            var v44 float32 = 0.0
            _ = v44
        }
        _ = v41
    }
    fmt.Println(v5, v4, v1)
    v31 = v2
    if true {
        if (!(!((^(v33)) >= (v33 / ((v33) | 1))))) {
            var v45 int64 = v33
            _ = v45
        }
        var v46 float32 = float32(((0.1 * (-2.25 - float64(float32(((float64(1.5) - 0.1) / 0.1))))) - (1e10 - 3.14159)))
        for fi47 := 0; fi47 < 15; fi47++ {
            if ((v33 * v33) <= (v33 << ((v1) & 63))) {
                if (!false) {
                    v46 = (v46 - v46)
                    fmt.Println(v5)
                } else {
                    v5 = v5
                }
            }
            v46 = v46
            switch "\n\t" {
            case "":
                if (v46 <= v46) {
                    var v48 bool = false
                    v49 := GenS3{F0: -1, F1: v1}
                    _ = v48
                    _ = v49
                }
            case "a":
                v50 := GenS1{F0: 144}
                _ = v50
            case "foo bar":
                var v51 int = fi47
                _ = v51
            }
            _ = fi47
        }
        _ = v46
    } else {
        fmt.Println(v31)
    }
    if true {
        if false {
            if true {
                v1 = v1
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v3 = append(v3, ("\n\t" + ((("foo bar" + (v4 + v4)) + ((("\n\t" + "foo bar") + "\n\t") + ((v4 + v4) + ("foo bar" + "\n\t")))) + ("" + v4))))
        }
        v52 := GenS3{F0: 42, F1: 255}
        switch 42 {
        case 42:
            for fi53 := 0; fi53 < 20; fi53++ {
                switch v4 {
                case "a":
                    var v54 uint32 = 1
                    _ = v54
                case "":
                    var v55 int32 = 1
                    _ = v55
                case "hello":
                    fmt.Println(v31)
                default:
                    for fi56 := 0; fi56 < 3; fi56++ {
                        v1 -= (v1 ^ (v1 % (((v1 | 18446744073709551615)) | 1)))
                        var v57 int64 = (v33 | (v33 >> (((v1 - (^(v1)))) & 63)))
                        var v58 int = fi56
                        var v59 int = fi53
                        switch (v4 + "hello") {
                        case "\n\t":
                            var v60 int8 = 0
                            _ = v60
                        case "hello":
                            v61 := GenS2{F0: 990, F1: v4, F2: uint16(v1)}
                            var v62 int16 = int16(v31)
                            _ = v61
                            _ = v62
                        case "":
                            if len(v3) > 0 {
                                v3[0] = ("hello" + (v4 + v4))
                            }
                            v31 = (v31 >> ((v1) & 63))
                            v33 -= (v57 + (v57 - (v57 | (v57 ^ int64(v5)))))
                        default:
                            v63 := GenS3{F0: 1000, F1: 4294967295}
                            _ = v63
                        }
                        _ = fi56
                        _ = v57
                        _ = v58
                        _ = v59
                    }
                }
                _ = fi53
            }
        case 0:
            v1 += (v1 + (^(v1)))
        }
        _ = v52
    }
    if true {
        for fi64 := 0; fi64 < 15; fi64++ {
            fmt.Println(v31, v31, fi64)
            for fi65 := 0; fi65 < 11; fi65++ {
                if false {
                    v66 := GenS1{F0: 42}
                    if (false && true) {
                        if (!true) {
                            v67 := GenS1{F0: 32767}
                            _ = v67
                        }
                        var v68 string = ((((v4 + (v4 + v4)) + "foo bar") + v4) + (v4 + v4))
                        var v69 int = 1
                        _ = v68
                        _ = v69
                    }
                    v2 = v31
                    v2 = (^(v2))
                    var v70 int32 = v31
                    _ = v66
                    _ = v70
                } else {
                    v71 := GenS4{F0: -173}
                    _ = v71
                }
                _ = fi65
            }
            v72 := GenS5{F0: -32768, F1: (false && ((false && false) && (3.14159 > 3.14159)))}
            _ = fi64
            _ = v72
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v73 float64 = (((3.14159 / (-2.25 * ((float64(0.1) / 1e10) * ((0.1 - 1e10) + ((float64(1e10) + float64((0.1 - 1.5))) + 3.14159))))) + -2.25) / 1.5)
        _ = v73
    }
    _ = len(v3)
    v33 -= v33
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
