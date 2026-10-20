package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 string; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -1
    v3 := []uint32{567, 876, 2147483648, 511}
    fmt.Println(v2)
    fmt.Println(v2)
    v4 := GenS1{F0: -32768}
    var v5 uint64 = 0
    if false {
        var v6 int32 = 0
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        if true {
            switch (("\n\t" + ("hello" + ("a" + "\n\t"))) + ("hello" + "foo bar")) {
            case "":
                var v8 int32 = -262
                _ = v8
            case "hello":
                for fi9 := 0; fi9 < 11; fi9++ {
                    switch "foo bar" {
                    case "hello":
                        v2 = (v2 | 42)
                    case "foo bar":
                        var v10 float64 = (0.0 + (((0.1 - (0.1 * 3.14159)) + -2.25) * 1.5))
                        fmt.Println(fi7, v10)
                        v11 := GenS1{F0: 830}
                        _ = v10
                        _ = v11
                    case "a":
                        if false {
                            var v12 uint16 = uint16(v5)
                            var v13 float64 = 3.14159
                            switch "hello" {
                            case "":
                                var v14 uint32 = 2147483648
                                _ = v14
                            case "hello":
                                v15 := GenS1{F0: 1000}
                                _ = v15
                            case "foo bar":
                                if (true || true) {
                                    fmt.Println(fi7, v12)
                                    v16 := GenS1{F0: -1}
                                    _ = v16
                                }
                            }
                            _ = v12
                            _ = v13
                        } else {
                            var v17 uint = 2
                            var v18 uint32 = 2147483648
                            var v19 uint16 = 830
                            switch (fi9 | (fi7 & 802)) {
                            case -37:
                                fmt.Println(fi7)
                            case 2147483647:
                                switch ("\n\t" + "hello") {
                                case "a":
                                    if ((v17 == 4294967295) && true) {
                                        v5 += (v5 % ((v5) | 1))
                                        v18 += v18
                                        delete(v1, fi7)
                                    } else {
                                        v5 += v5
                                        v2 = (v2 ^ v2)
                                    }
                                case "hello":
                                    delete(v1, (-(fi7)))
                                default:
                                    var v20 bool = ((v2 % ((v2) | 1)) < (^(v2)))
                                    _ = v20
                                }
                            case -2147483648:
                                if (1e10 != (0.0 * (float32(3.14159) / (0.1 - 3.14159)))) {
                                    switch "a" {
                                    case "foo bar":
                                        v17 = 505
                                        fmt.Println(v17)
                                        if (!true) {
                                            var v21 uint = (v17 >> ((v17) & 63))
                                            if ((v21 >> ((v21) & 63)) != v21) {
                                                v22 := v1[0]
                                                v17 -= uint(v5)
                                                switch (v22 ^ v22) {
                                                case -109:
                                                    var v23 uint8 = uint8(v18)
                                                    _ = v23
                                                case 4294967295:
                                                    fmt.Println(v21)
                                                }
                                                if true {
                                                    switch (v22 + v22) {
                                                    case 301:
                                                        v17 -= (v17 * (v17 - uint(v5)))
                                                    case 9223372036854775807:
                                                        if false {
                                                            delete(v1, 32767)
                                                            switch fi9 {
                                                            case 127:
                                                                if false {
                                                                    var v24 uint8 = 42
                                                                    var v25 uint32 = (v18 << ((v17) & 63))
                                                                    var v26 int8 = 42
                                                                    v3 = append(v3, v18)
                                                                    v27 := GenS2{F0: 255}
                                                                    _ = v24
                                                                    _ = v25
                                                                    _ = v26
                                                                    _ = v27
                                                                } else {
                                                                    if ("" > "foo bar") {
                                                                        delete(v1, (v22 / (((fi9 % ((fi7) | 1))) | 1)))
                                                                        if (true && (v2 <= (v2 / ((127) | 1)))) {
                                                                            _ = len(v3)
                                                                        } else {
                                                                            if ((((1e10 / (float32((float64(3.14159) * float64(1.5))) - 0.1)) + ((float32(1.5) / 1.5) * (0.1 * float32(0.1)))) - 3.14159) < -2.25) {
                                                                                v28 := GenS3{F0: v18, F1: -1}
                                                                                v29 := GenS3{F0: uint32(v17), F1: 42}
                                                                                _ = v28
                                                                                _ = v29
                                                                            }
                                                                            v19 += v19
                                                                        }
                                                                    } else {
                                                                        var v30 int = fi7
                                                                        _ = v30
                                                                    }
                                                                    switch (((("\n\t" + "foo bar") + (("" + ("\n\t" + "\n\t")) + "hello")) + "a") + "") {
                                                                    case "hello":
                                                                        var v31 int64 = 9223372036854775807
                                                                        _ = v31
                                                                    case "":
                                                                        v32 := GenS1{F0: 0}
                                                                        _ = v32
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (v18 + v18)
                                                                    }
                                                                }
                                                            case 128:
                                                                if true {
                                                                    var v33 int = (fi9 / ((v22) | 1))
                                                                    v34 := GenS1{F0: -32768}
                                                                    _ = v33
                                                                    _ = v34
                                                                } else {
                                                                    if false {
                                                                        v19 = v19
                                                                        var v35 float64 = 1.5
                                                                        _ = v35
                                                                    } else {
                                                                        v5 = v5
                                                                    }
                                                                }
                                                            default:
                                                                var v36 float64 = 0.1
                                                                _ = v36
                                                            }
                                                            v1[v22] = fi7
                                                        } else {
                                                            switch int(v2) {
                                                            case -2147483648:
                                                                v17 -= 4294967295
                                                                fmt.Println(fi9, fi9, v21)
                                                                if (1 >= 817) {
                                                                    var v37 float32 = (1e10 + 1e10)
                                                                    v18 = v18
                                                                    _ = v37
                                                                }
                                                            case -1:
                                                                fmt.Println(v18)
                                                            default:
                                                                v38 := GenS1{F0: 1000}
                                                                _ = v38
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(fi9, v22, v17)
                                                    }
                                                } else {
                                                    if (v18 < 1) {
                                                        v21 -= v21
                                                    } else {
                                                        v5 -= (v5 + v5)
                                                        var v39 uint64 = v5
                                                        _ = v39
                                                    }
                                                    if (1000 < -653) {
                                                        if (!false) {
                                                            v2 -= v2
                                                        }
                                                        switch (fi9 << (((^(v17))) & 63)) {
                                                        case 32767:
                                                            if false {
                                                                v22 = fi9
                                                            }
                                                        }
                                                    }
                                                    var v40 int8 = v2
                                                    _ = v40
                                                }
                                                var v41 int16 = 1
                                                v2 = 127
                                                _ = v22
                                                _ = v41
                                            }
                                            var v42 int = (fi7 / (((^(fi7))) | 1))
                                            var v43 uint8 = uint8(v18)
                                            if ((v18 % (((^(v18))) | 1)) < (^(v18))) {
                                                var v44 int16 = -90
                                                v44 -= v44
                                                _ = v44
                                            } else {
                                                v45 := GenS2{F0: uint8(v19)}
                                                _ = v45
                                            }
                                            var v46 uint64 = (v5 ^ (v5 + (v5 << ((v17) & 63))))
                                            v46 = (^(v46))
                                            switch "a" {
                                            case "\n\t":
                                                v47 := GenS2{F0: (v43 + v43)}
                                                _ = v47
                                            case "":
                                                _ = len(v3)
                                            default:
                                                v3 = append(v3, 0)
                                            }
                                            v43 -= (^(v43))
                                            _ = v21
                                            _ = v42
                                            _ = v43
                                            _ = v46
                                        }
                                    case "":
                                        fmt.Println(fi9, v19)
                                    }
                                    if true {
                                        fmt.Println(v19, v5, v5)
                                        var v48 string = "a"
                                        var v49 int16 = -32768
                                        _ = len(v3)
                                        _ = v48
                                        _ = v49
                                    }
                                }
                            }
                            var v50 int16 = int16(v19)
                            _ = v17
                            _ = v18
                            _ = v19
                            _ = v50
                        }
                    }
                    if (65535 >= 32767) {
                        var v51 uint32 = 42
                        _ = v51
                    } else {
                        v52 := v1[-1]
                        var v53 uint8 = 2
                        if true {
                            if (0 == v53) {
                                v53 = (v53 << ((65535) & 63))
                                switch (fi9 >> ((65535) & 63)) {
                                case 128:
                                    var v54 uint32 = 0
                                    _ = v54
                                case 0:
                                    v55 := GenS4{F0: 535, F1: -2147483648}
                                    _ = v55
                                case 1:
                                    var v56 int64 = 42
                                    _ = v56
                                }
                                v52 -= 507
                            } else {
                                v57 := GenS2{F0: 0}
                                _ = v57
                            }
                            fmt.Println(v53, v53, v2)
                        } else {
                            var v58 int = -1
                            var v59 uint32 = 2
                            fmt.Println(v52, v53, v2)
                            _ = v58
                            _ = v59
                        }
                        _ = v52
                        _ = v53
                    }
                    if (-586 > 327) {
                        fmt.Println(v5, fi9)
                    }
                    _ = fi9
                }
            case "a":
                v5 -= v5
            }
            var v60 int = fi7
            _ = v60
        } else {
            v61 := GenS5{F0: ("\n\t" + "\n\t"), F1: 600}
            if true {
                v62 := GenS3{F0: 60, F1: 1}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v63 uint8 = 42
                _ = v62
                _ = v63
            }
            _ = v61
        }
        _ = fi7
    }
    for fi64 := 0; fi64 < 1; fi64++ {
        fmt.Println(v5)
        var v65 float32 = 0.1
        _ = fi64
        _ = v65
    }
    var v66 bool = (false && false)
    if len(v3) > 0 {
        v3[0] = 42
    }
    switch -1 {
    case 32768:
        fmt.Println(v66, v2, v2)
    case 864:
        switch -2147483648 {
        case -336:
            var v67 string = "foo bar"
            _ = v67
        default:
            fmt.Println(v5, v66)
        }
    }
    fmt.Println(v2, v2, v66)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
}
