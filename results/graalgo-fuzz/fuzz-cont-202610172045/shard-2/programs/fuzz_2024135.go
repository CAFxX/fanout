package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0() int8 {
    var v7 int16 = -221
    if (!false) {
        var v8 float64 = (0.0 + 3.14159)
        var v9 uint = 18446744073709551615
        if false {
            for fi10 := 0; fi10 < 15; fi10++ {
                if true {
                    v11 := GenS1{F0: v7, F1: (-(v7)), F2: (v7 - 1000)}
                    _ = v11
                }
                v12 := GenS1{F0: v7, F1: -101, F2: (v7 - (v7 / ((1000) | 1)))}
                v13 := GenS1{F0: (v7 >> ((255) & 63)), F1: v7, F2: 1000}
                _ = fi10
                _ = v12
                _ = v13
            }
            switch "hello" {
            case "\n\t":
                var v14 int8 = 42
                _ = v14
            }
            switch 1 {
            case 9223372036854775807:
                if true {
                    v15 := GenS1{F0: 0, F1: (-(v7)), F2: (v7 * (v7 << (((^(v9))) & 63)))}
                    _ = v15
                } else {
                    if (v9 >= v9) {
                        for fi16 := 0; fi16 < 12; fi16++ {
                            var v17 uint64 = 0
                            var v18 int = fi16
                            v9 -= 65535
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        var v19 string = "foo bar"
                        var v20 int64 = 42
                        _ = v19
                        _ = v20
                    } else {
                        v21 := GenS1{F0: v7, F1: v7, F2: 1000}
                        var v22 uint64 = 18446744073709551615
                        var v23 float32 = 1e10
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS1{F0: v7, F1: v7, F2: int16(v9)}
                    var v25 bool = false
                    _ = v24
                    _ = v25
                }
            case 0:
                var v26 int64 = 0
                _ = v26
            case 2147483647:
                v9 += (v9 << ((65535) & 63))
            }
        } else {
            v27 := GenS1{F0: v7, F1: (v7 * (v7 ^ v7)), F2: (-(v7))}
            _ = v27
        }
        _ = v8
        _ = v9
    }
    var v28 int8 = -128
    v28 = -24
    _ = v28
    _ = v7
    return v28
}
func genFn1(p0 int32, p1 int8) int {
    p1 = (p1 + genFn0())
    fmt.Println(p0, p1)
    p0 -= p0
    if ((p1 & p1) == genFn0()) {
        fmt.Println(p0)
        switch (("\n\t" + "a") + ("hello" + ((("foo bar" + "a") + ("a" + "\n\t")) + "hello"))) {
        case "foo bar":
            if true {
                v29 := []int{-809, -2147483648, 9223372036854775807}
                p1 += (^(p1))
                _ = v29
            } else {
                fmt.Println(p1, p0)
                var v30 int64 = -994
                _ = v30
            }
        case "hello":
            var v31 string = "hello"
            _ = v31
        }
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint = 42
    v2 := make(map[int]string)
    var v3 uint8 = 1
    v4 := make(map[string]string)
    var v5 uint64 = 0
    var v6 int16 = -1
    v32 := v2[32767]
    switch (v32 + "\n\t") {
    case "\n\t":
        if (100 != -1) {
            v33 := []int{}
            var v34 uint16 = 368
            if false {
                switch 536 {
                case 128:
                    var v35 string = (("foo bar" + ("hello" + ("hello" + ("\n\t" + (((v32 + v32) + v32) + v32))))) + "foo bar")
                    _ = v35
                case 445:
                    v6 = v6
                case -941:
                    switch -182 {
                    case 32768:
                        switch (v32 + (v32 + v32)) {
                        case "\n\t":
                            var v36 float32 = ((0.0 + 3.14159) * ((-2.25 + float32(1.5)) / -2.25))
                            _ = v36
                        }
                    default:
                        var v37 string = ("\n\t" + "")
                        _ = v37
                    }
                default:
                    if false {
                        for fi38 := 0; fi38 < 18; fi38++ {
                            v32 = (v32 + (v32 + v32))
                            v39 := GenS1{F0: v6, F1: (-(v6)), F2: v6}
                            var v40 float64 = (0.1 - -2.25)
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        fmt.Println(v32, v5, v3)
                        switch genFn1(42, -128) {
                        case 128:
                            v41 := GenS1{F0: v6, F1: (v6 * (^(v6))), F2: 1}
                            _ = v41
                        default:
                            var v42 float32 = 1e10
                            _ = v42
                        }
                    }
                }
                var v43 int16 = v6
                var v44 int = genFn1(-1, 1)
                _ = v43
                _ = v44
            }
            _ = v33
            _ = v34
        } else {
            v45 := GenS2{F0: true, F1: (v3 / ((v3) | 1))}
            if (v5 <= (v5 % (((v5 | v5)) | 1))) {
                var v46 bool = false
                _ = v46
            }
            fmt.Println(v32, v3, v5)
            _ = v45
        }
    case "":
        var v47 int = -9223372036854775808
        _ = v47
    default:
        if true {
            var v48 int = 9223372036854775807
            v6 += v6
            v49 := GenS1{F0: v6, F1: (v6 / (((^(v6))) | 1)), F2: v6}
            _ = v48
            _ = v49
        } else {
            if (true || true) {
                v50 := []int{4294967295, -533, -908, 9223372036854775807}
                for fi51 := 0; fi51 < 20; fi51++ {
                    var v52 string = ("foo bar" + "")
                    var v53 int = (fi51 & fi51)
                    fmt.Println(fi51)
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                _ = v50
            }
        }
    }
    v6 -= v6
    v54 := GenS2{F0: (uint32(v1) < uint32(v1)), F1: (v3 & (v3 + (v3 * 127)))}
    var v55 string = "hello"
    var v56 uint = (v1 / ((v1) | 1))
    var v57 uint = (v1 - v1)
    for fi58 := 0; fi58 < 6; fi58++ {
        delete(v4, v55)
        _ = fi58
    }
    var v59 uint32 = 1
    for fi60 := 0; fi60 < 5; fi60++ {
        var v61 uint16 = 1
        _ = fi60
        _ = v61
    }
    v59 -= (v59 & 310)
    for fi62 := 0; fi62 < 6; fi62++ {
        for fi63 := 0; fi63 < 16; fi63++ {
            fmt.Println(fi63, v57)
            v59 += 2
            if false {
                fmt.Println(fi62, fi62)
                v64 := GenS3{F0: v59}
                var v65 float64 = (0.1 - -2.25)
                _ = v64
                _ = v65
            }
            v1 -= (^(v57))
            v59 -= (v59 - (v59 >> ((4294967295) & 63)))
            var v66 bool = true
            fmt.Println(fi63, v56, v66)
            _ = fi63
            _ = v66
        }
        _ = fi62
    }
    var v67 uint32 = 577
    v68 := GenS3{F0: uint32(v3)}
    switch 32767 {
    case -770:
        for fi69 := 0; fi69 < 10; fi69++ {
            var v70 float32 = 1.5
            _ = fi69
            _ = v70
        }
    case 9223372036854775807:
        var v71 uint32 = 416
        _ = v71
    default:
        if (true || true) {
            v72 := []int{32768, 32768}
            v72 = append(v72, genFn1(-423, int8(v1)))
            _ = v72
        } else {
            v59 = 2
            switch 0 {
            case -465:
                v73 := GenS3{F0: 2}
                _ = v73
            case 584:
                v6 -= (v6 - int16(v59))
            }
        }
    }
    delete(v2, 42)
    var v74 int16 = v6
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v59
    _ = v6
    _ = v67
    _ = v68
    _ = v74
}
