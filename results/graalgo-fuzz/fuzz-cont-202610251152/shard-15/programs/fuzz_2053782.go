package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int16; F2 int }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 int64; F1 int }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = 3.14159
    v3 := make(map[string]string)
    if (!(3.14159 < 0.0)) {
        var v4 int8 = 1
        v5 := []int{-20, 42, 127}
        _ = v4
        _ = v5
    } else {
        if false {
            var v6 float64 = 0.0
            for fi7 := 0; fi7 < 18; fi7++ {
                var v8 int = (fi7 & (fi7 % ((fi7) | 1)))
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 5; fi9++ {
                switch "hello" {
                case "":
                    if true {
                        v6 = 1e10
                        fmt.Println(v2, v6)
                    }
                case "a":
                    v1 = v1
                case "hello":
                    if (false || (!false)) {
                        if (!(1e10 != -2.25)) {
                            fmt.Println(v1, fi9, v1)
                            for fi10 := 0; fi10 < 18; fi10++ {
                                if false {
                                    switch "foo bar" {
                                    case "\n\t":
                                        _ = len(v3)
                                    }
                                    v6 -= (v6 * float64(v2))
                                }
                                v2 = 0.0
                                v6 = (-2.25 / v6)
                                _ = fi10
                            }
                            var v11 int16 = -591
                            _ = v11
                        }
                        switch fi9 {
                        case 1:
                            if (true || true) {
                                v12 := GenS1{F0: "hello"}
                                var v13 uint8 = 0
                                _ = v12
                                _ = v13
                            }
                        default:
                            fmt.Println(v1, v1)
                        }
                    } else {
                        v3[""] = ("hello" + ("" + ("" + "foo bar")))
                    }
                default:
                    var v14 uint16 = 321
                    _ = v14
                }
                if (551 < 65535) {
                    var v15 float32 = v2
                    _ = v15
                }
                if ((-360 <= 593) && true) {
                    var v16 int64 = 42
                    _ = v16
                } else {
                    var v17 uint16 = 0
                    v18 := GenS1{F0: (("" + (("" + (("\n\t" + "hello") + "foo bar")) + "")) + "\n\t")}
                    v19 := []int{(fi9 * fi9), fi9, 32768, 128}
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = fi9
            }
            _ = v6
        }
        v20 := []int{}
        if false {
            v2 += v2
        } else {
            if ((42 < 476) || false) {
                var v21 string = ("foo bar" + "\n\t")
                for fi22 := 0; fi22 < 20; fi22++ {
                    v1 -= (^(v1))
                    _ = fi22
                }
                switch v21 {
                case "":
                    switch (v21 + v21) {
                    case "hello":
                        if (false || true) {
                            delete(v3, v21)
                        }
                    case "foo bar":
                        var v23 string = v21
                        _ = v23
                    }
                case "a":
                    v1 = (v1 | (v1 ^ v1))
                case "hello":
                    if false {
                        switch v21 {
                        case "a":
                            var v24 uint64 = v1
                            _ = v24
                        default:
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                        }
                    } else {
                        fmt.Println(v1)
                    }
                }
                _ = v21
            }
            var v25 int = -1
            v1 -= 4294967295
            _ = v25
        }
        _ = v20
    }
    var v26 bool = ((1e10 * 3.14159) >= v2)
    if true {
        delete(v3, "")
        v27 := GenS2{F0: v1, F1: 93}
        for fi28 := 0; fi28 < 12; fi28++ {
            switch "foo bar" {
            case "a":
                var v29 float32 = v2
                _ = v29
            }
            var v30 uint = 42
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 string = ("foo bar" + (("\n\t" + "a") + ""))
                v33 := []int{fi28, (fi31 * -921)}
                switch 128 {
                case 0:
                    var v34 uint = uint(v1)
                    _ = v34
                case -529:
                    var v35 int = fi31
                    _ = v35
                case 354:
                    v2 -= v2
                default:
                    v30 = (^(v30))
                }
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = fi28
            _ = v30
        }
        _ = v27
    } else {
        fmt.Println(v1, v26, v2)
        v36 := GenS2{F0: (v1 * (v1 + (^(v1)))), F1: 1}
        if (!v26) {
            if ((!false) || (!v26)) {
                v37 := v3[("\n\t" + "foo bar")]
                if (v26 && (((v1 >= (v1 & (^(v1)))) || false) || v26)) {
                    var v38 int64 = -9223372036854775808
                    for fi39 := 0; fi39 < 16; fi39++ {
                        v37 = (((v37 + v37) + "foo bar") + v37)
                        _ = fi39
                    }
                    if true {
                        if v26 {
                            for fi40 := 0; fi40 < 3; fi40++ {
                                v2 = v2
                                var v41 uint64 = (v1 << ((255) & 63))
                                switch fi40 {
                                case -9223372036854775808:
                                    var v42 uint = 1
                                    _ = v42
                                case -1:
                                    var v43 uint32 = 2147483647
                                    _ = v43
                                }
                                _ = fi40
                                _ = v41
                            }
                            v44 := GenS3{F0: v1}
                            v37 = ((v37 + v37) + (v37 + "hello"))
                            _ = v44
                        }
                        var v45 bool = (4294967295 < 690)
                        var v46 uint64 = (v1 % ((v1) | 1))
                        _ = v45
                        _ = v46
                    } else {
                        var v47 uint32 = 0
                        v48 := GenS3{F0: v1}
                        if v26 {
                            v37 = v37
                            delete(v3, v37)
                        } else {
                            v26 = ((((v37 + "\n\t") + v37) > (v37 + v37)) || (0.1 <= float64(v2)))
                            switch v37 {
                            case "\n\t":
                                fmt.Println(v1)
                            case "a":
                                v3[(((v37 + v37) + v37) + "a")] = "foo bar"
                            }
                        }
                        _ = v47
                        _ = v48
                    }
                    _ = v38
                }
                fmt.Println(v37, v2, v1)
                _ = v37
            } else {
                v49 := GenS1{F0: ""}
                switch 9223372036854775807 {
                case 4294967295:
                    _ = len(v3)
                case 2147483647:
                    v2 = v2
                default:
                    var v50 uint32 = 1
                    _ = v50
                }
                _ = v49
            }
            if (!(255 <= 20)) {
                var v51 int16 = -1
                _ = v51
            } else {
                v1 = v1
                v3["hello"] = "a"
            }
        } else {
            v52 := GenS2{F0: (v1 ^ (v1 + v1)), F1: -128}
            v53 := []int{-2147483648}
            var v54 int16 = 1000
            _ = v52
            _ = v53
            _ = v54
        }
        _ = v36
    }
    if (2 <= 32767) {
        v55 := GenS4{F0: 18446744073709551615, F1: 0, F2: 9223372036854775807}
        fmt.Println(v2, v1)
        v1 = v1
        _ = v55
    } else {
        v26 = v26
        switch "\n\t" {
        case "a":
            var v56 uint32 = 367
            _ = v56
        case "":
            var v57 int32 = -2147483648
            _ = v57
        case "hello":
            fmt.Println(v1)
        }
    }
    var v58 uint64 = (^(v1))
    var v59 int16 = -710
    if (!(!v26)) {
        for fi60 := 0; fi60 < 11; fi60++ {
            var v61 float32 = -2.25
            var v62 int = (fi60 - fi60)
            _ = fi60
            _ = v61
            _ = v62
        }
    }
    switch "\n\t" {
    case "":
        var v63 int16 = v59
        _ = v63
    case "\n\t":
        v58 = 1
    default:
        fmt.Println(v2, v58, v26)
    }
    v59 += v59
    var v64 int32 = 2147483647
    switch "\n\t" {
    case "\n\t":
        for fi65 := 0; fi65 < 17; fi65++ {
            v66 := GenS3{F0: (v1 & (v58 - 42))}
            var v67 int64 = -9223372036854775808
            switch (fi65 >> ((25) & 63)) {
            case 0:
                _ = len(v3)
            case 128:
                v58 = (^(v1))
            case 127:
                v68 := GenS5{F0: v59, F1: 2}
                _ = v68
            }
            _ = fi65
            _ = v66
            _ = v67
        }
    }
    fmt.Println(v59, v59, v2)
    v69 := GenS6{F0: 343, F1: int(v58)}
    v70 := GenS3{F0: (v58 ^ (v1 - 42))}
    if v26 {
        var v71 int32 = 0
        var v72 uint = 194
        v73 := GenS2{F0: 0, F1: 0}
        _ = v71
        _ = v72
        _ = v73
    }
    switch 32768 {
    case -1:
        var v74 bool = (980 > 42)
        _ = v74
    }
    var v75 int = 0
    var v76 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v64:", v64)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v58
    _ = v59
    _ = v64
    _ = v69
    _ = v70
    _ = v75
    _ = v76
}
