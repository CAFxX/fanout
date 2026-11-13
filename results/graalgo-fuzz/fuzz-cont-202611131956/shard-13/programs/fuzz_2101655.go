package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int; F1 int8; F2 uint }
func genFn0(p0 uint, p1 int8, p2 int64) int32 {
    for fi5 := 0; fi5 < 12; fi5++ {
        if false {
            var v6 int = (fi5 & fi5)
            _ = v6
        }
        _ = fi5
    }
    if false {
        p2 -= (p2 | (^(p2)))
    } else {
        p1 -= (p1 / ((0) | 1))
        p2 -= (p2 / ((-9223372036854775808) | 1))
        v7 := GenS1{F0: 201, F1: 2, F2: 1}
        _ = v7
    }
    switch "hello" {
    case "":
        if ((((0.0 * 0.1) - ((0.0 + float64(1.5)) / float64(3.14159))) < 0.0) || false) {
            var v8 string = "hello"
            v9 := GenS1{F0: 42, F1: 4294967295, F2: -32768}
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 uint32 = 2147483648
                var v12 float32 = ((0.0 * (((((3.14159 * 0.1) * (1e10 + 3.14159)) / (1e10 + 0.0)) * 0.1) / (1.5 + -2.25))) + (float32(0.0) / 1e10))
                switch "hello" {
                case "foo bar":
                    for fi13 := 0; fi13 < 6; fi13++ {
                        switch ("hello" + ("\n\t" + "foo bar")) {
                        case "":
                            v14 := GenS2{F0: false}
                            _ = v14
                        default:
                            var v15 uint16 = 922
                            _ = v15
                        }
                        var v16 float32 = float32(3.14159)
                        _ = fi13
                        _ = v16
                    }
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            p2 += (p2 % ((-9223372036854775808) | 1))
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := GenS3{F0: 1, F1: 2}
                switch "" {
                case "a":
                    fmt.Println(p0, p0, fi17)
                case "":
                    var v19 float32 = float32((float64(0.1) - -2.25))
                    _ = v19
                case "foo bar":
                    var v20 int32 = 659
                    _ = v20
                default:
                    var v21 float64 = float64(0.1)
                    _ = v21
                }
                _ = fi17
                _ = v18
            }
        }
    }
    if ((1 < 42) && true) {
        v22 := []int{int(p2)}
        _ = v22
    } else {
        var v23 int16 = 1000
        for fi24 := 0; fi24 < 6; fi24++ {
            for fi25 := 0; fi25 < 7; fi25++ {
                p2 -= p2
                _ = fi25
            }
            var v26 float32 = 0.0
            switch "foo bar" {
            case "hello":
                if ((0.1 - 1e10) > 3.14159) {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        var v28 uint32 = 42
                        if (fi27 > fi24) {
                            fmt.Println(p1)
                            v29 := GenS2{F0: (fi24 > fi24)}
                            _ = v29
                        }
                        _ = fi27
                        _ = v28
                    }
                    var v30 float64 = (1.5 + ((((float64(v26) - float64(v26)) * (0.1 - (0.1 * (((3.14159 / -2.25) - float64(v26)) / 0.1)))) + (((3.14159 - 0.1) + (0.1 + ((0.1 / (-2.25 / 1.5)) * ((0.0 * float64((v26 / v26))) - 3.14159)))) * (0.1 + float64(0.0)))) + 0.1))
                    fmt.Println(fi24, v30, v23)
                    _ = v30
                } else {
                    switch (fi24 * (fi24 % ((-334) | 1))) {
                    case 1:
                        for fi31 := 0; fi31 < 2; fi31++ {
                            var v32 int64 = 572
                            _ = fi31
                            _ = v32
                        }
                    case -617:
                        switch 0 {
                        case 128:
                            v33 := GenS3{F0: 58, F1: 4294967295}
                            _ = v33
                        case -834:
                            fmt.Println(v26)
                        case -1:
                            for fi34 := 0; fi34 < 12; fi34++ {
                                v35 := GenS1{F0: 4294967295, F1: 4294967295, F2: v23}
                                v36 := GenS3{F0: 0, F1: 4294967295}
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                        }
                    default:
                        var v37 float32 = v26
                        _ = v37
                    }
                    if false {
                        for fi38 := 0; fi38 < 11; fi38++ {
                            if false {
                                p2 += -215
                                var v39 float64 = 1.5
                                _ = v39
                            }
                            switch fi24 {
                            case 9223372036854775807:
                                fmt.Println(fi38, v26)
                            default:
                                var v40 uint8 = 42
                                _ = v40
                            }
                            p0 -= (^(p0))
                            p0 += p0
                            _ = fi38
                        }
                        var v41 uint = (p0 - p0)
                        _ = v41
                    }
                    v42 := []int{(fi24 - (-(fi24))), fi24, (fi24 & (fi24 >> (((p0 / ((p0) | 1))) & 63))), fi24}
                    _ = v42
                }
            }
            _ = fi24
            _ = v26
        }
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 string = "foo bar"
    var v2 float64 = ((0.0 + 0.1) * (1e10 - float64(-2.25)))
    var v3 int32 = 2147483647
    v4 := []string{("\n\t" + "a"), "", v1, "\n\t"}
    for fi43 := 0; fi43 < 5; fi43++ {
        v2 = float64(0.1)
        for fi44 := 0; fi44 < 18; fi44++ {
            var v45 string = v1
            _ = fi44
            _ = v45
        }
        _ = fi43
    }
    switch 9223372036854775807 {
    case 42:
        v1 = "a"
    case 4294967295:
        for fi46 := 0; fi46 < 20; fi46++ {
            switch (fi46 << ((339) & 63)) {
            case -257:
                v4 = append(v4, "")
            default:
                var v47 int32 = -739
                _ = v47
            }
            v3 -= genFn0(4294967295, -1, -9223372036854775808)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi46
        }
    }
    var v48 uint32 = 42
    for fi49 := 0; fi49 < 12; fi49++ {
        var v50 float64 = v2
        v3 -= genFn0(65535, 42, -1)
        _ = fi49
        _ = v50
    }
    v3 = 42
    v1 = (((v1 + v1) + "\n\t") + (v1 + ((v1 + v1) + "")))
    var v51 float32 = 0.1
    var v52 int = -633
    for fi53 := 0; fi53 < 10; fi53++ {
        v51 = v51
        var v54 uint16 = 65535
        _ = fi53
        _ = v54
    }
    switch v52 {
    case 128:
        v55 := GenS4{F0: 2, F1: 100}
        _ = v55
    case -9223372036854775808:
        fmt.Println(v52, v3)
    case -2147483648:
        v4 = append(v4, "\n\t")
    }
    v56 := GenS5{F0: (v52 + (v52 * (v52 + -392))), F1: 100, F2: 65535}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v57 float64 = ((v2 + v2) / v2)
    var v58 int8 = 42
    var v59 int = v52
    v59 = (-(v59))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v51
    _ = v52
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
