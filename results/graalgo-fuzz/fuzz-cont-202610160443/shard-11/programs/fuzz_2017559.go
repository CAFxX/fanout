package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int64; F1 uint; F2 uint8 }
func genFn0() int64 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint64 = 18446744073709551615
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v10 := GenS1{F0: 42, F1: 0}
        var v11 uint32 = 877
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "foo bar" {
    case "\n\t":
        switch "a" {
        case "hello":
            switch "a" {
            case "\n\t":
                var v12 float64 = (-2.25 / 1.0)
                _ = v12
            case "foo bar":
                var v13 uint8 = 127
                _ = v13
            default:
                if (42 > 1) {
                    var v14 int = 32767
                    v15 := []int{(v14 ^ (v14 >> ((255) & 63))), (v14 | v14), (v14 - (v14 << ((905) & 63))), v14}
                    for fi16 := 0; fi16 < 16; fi16++ {
                        fmt.Println(v14)
                        var v17 int16 = 32767
                        _ = fi16
                        _ = v17
                    }
                    _ = v14
                    _ = v15
                }
            }
        }
    }
    if (65535 <= 596) {
        var v18 uint8 = 0
        v19 := GenS2{F0: 128, F1: 523}
        switch "foo bar" {
        case "":
            fmt.Println(v18)
        case "foo bar":
            v18 = v18
        }
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 int = fi20
            _ = fi20
            _ = v21
        }
        switch -458 {
        case 32768:
            for fi22 := 0; fi22 < 9; fi22++ {
                fmt.Println(fi22)
                var v23 bool = false
                _ = fi22
                _ = v23
            }
        case -420:
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 int64 = -9223372036854775808
                v26 := GenS3{F0: -939}
                for fi27 := 0; fi27 < 8; fi27++ {
                    switch fi27 {
                    case 42:
                        if true {
                            v28 := GenS1{F0: 42, F1: 32767}
                            switch fi24 {
                            case -1:
                                v29 := []int{127, fi27, (fi27 & (fi27 << ((4294967295) & 63))), fi27}
                                switch "" {
                                case "a":
                                    v25 = -1
                                case "\n\t":
                                    if (!(!true)) {
                                        switch 4294967295 {
                                        case 32768:
                                            var v30 uint16 = 656
                                            _ = v30
                                        case 4294967295:
                                            fmt.Println(fi27)
                                        default:
                                            v31 := GenS2{F0: 1, F1: 85}
                                            _ = v31
                                        }
                                    } else {
                                        if false {
                                            fmt.Println(fi24, fi27)
                                        }
                                    }
                                }
                                if (0 < 1000) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v32 uint32 = 212
                                        _ = v32
                                    case "a":
                                        var v33 int64 = -517
                                        _ = v33
                                    default:
                                        if len(v29) >= 2 {
                                            v29 = v29[1:len(v29)]
                                        }
                                    }
                                    var v34 int64 = v25
                                    _ = v34
                                }
                                _ = v29
                            case -529:
                                fmt.Println(fi24, fi24, fi24)
                            default:
                                v25 += (v25 * v25)
                            }
                            if true {
                                fmt.Println(v25)
                                v25 -= v25
                                v35 := []int{(fi24 | (fi27 - fi24)), fi27, fi24, 127}
                                _ = v35
                            } else {
                                v25 += (-(v25))
                                fmt.Println(v25, fi27)
                            }
                            v36 := GenS4{F0: 4294967295, F1: 27, F2: fi24}
                            v37 := GenS1{F0: 545, F1: 1}
                            fmt.Println(fi27, v25)
                            v38 := GenS2{F0: 32767, F1: 69}
                            _ = v28
                            _ = v36
                            _ = v37
                            _ = v38
                        } else {
                            var v39 uint64 = 1
                            _ = v39
                        }
                    }
                    var v40 uint8 = 175
                    _ = fi27
                    _ = v40
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        case -162:
            var v41 uint8 = 181
            _ = v41
        default:
            if true {
                v42 := GenS5{F0: 9223372036854775807, F1: 4294967295}
                for fi43 := 0; fi43 < 18; fi43++ {
                    var v44 float64 = 0.0
                    for fi45 := 0; fi45 < 6; fi45++ {
                        var v46 int = (fi45 - (^(fi45)))
                        _ = fi45
                        _ = v46
                    }
                    v47 := GenS2{F0: fi43, F1: 139}
                    _ = fi43
                    _ = v44
                    _ = v47
                }
                var v48 uint32 = 42
                _ = v42
                _ = v48
            }
        }
        var v49 int32 = 42
        _ = v49
    }
    return 9223372036854775807
}

func main() {
    var v1 int = 629
    var v2 int16 = 134
    var v3 uint32 = 65535
    var v4 uint8 = 2
    v5 := []int{4294967295}
    v6 := make(map[uint32]string)
    _ = len(v6)
    if (1.5 < 3.14159) {
        for fi50 := 0; fi50 < 10; fi50++ {
            switch "hello" {
            case "foo bar":
                var v51 int16 = -657
                _ = v51
            case "\n\t":
                var v52 int8 = int8(v1)
                _ = v52
            case "hello":
                for fi53 := 0; fi53 < 3; fi53++ {
                    delete(v6, v3)
                    v4 = 128
                    _ = fi53
                }
            }
            fmt.Println(v2, v1, v4)
            v54 := v6[v3]
            _ = fi50
            _ = v54
        }
        var v55 uint8 = 216
        _ = v55
    }
    fmt.Println(v2, v3)
    if (3.14159 < 1.5) {
        _ = len(v5)
        v2 = v2
    } else {
        var v56 string = "\n\t"
        for fi57 := 0; fi57 < 1; fi57++ {
            v5 = append(v5, (fi57 % ((-256) | 1)))
            _ = fi57
        }
        _ = v56
    }
    var v58 uint16 = 1
    v4 = (v4 / ((v4) | 1))
    switch 32768 {
    case 127:
        v59 := GenS4{F0: v3, F1: 127, F2: v1}
        _ = v59
    case 330:
        v4 -= v4
    case 32767:
        v60 := GenS6{F0: 42, F1: 933, F2: (v4 | (v4 | (v4 ^ v4)))}
        _ = v60
    }
    var v61 uint32 = v3
    delete(v6, (v3 / (((v3 - (v61 << ((4294967295) & 63)))) | 1)))
    v62 := GenS6{F0: 859, F1: 419, F2: v4}
    switch (("foo bar" + ((("\n\t" + ("\n\t" + "")) + (("a" + "\n\t") + "a")) + "foo bar")) + ("foo bar" + "foo bar")) {
    case "a":
        if (v3 == v3) {
            v58 -= 255
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v61
    _ = v62
}
