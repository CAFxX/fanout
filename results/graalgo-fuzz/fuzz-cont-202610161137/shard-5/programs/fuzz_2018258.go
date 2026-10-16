package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint; F1 int8; F2 int }
type GenS6 struct { F0 int; F1 int64; F2 uint }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 int16; F1 int64 }
type GenS10 struct { F0 int8 }
func genFn0(p0 bool) int8 {
    v5 := GenS1{F0: -1}
    v6 := GenS1{F0: -50}
    var v7 int8 = 121
    fmt.Println(v7, p0)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 int64) bool {
    if (!(1 >= 930)) {
        p0 = (p0 + 0)
    }
    v8 := GenS2{F0: -503}
    p0 = p0
    fmt.Println(p0)
    var v9 int32 = 42
    _ = p0
    _ = v8
    _ = v9
    return (101 > 127)
}

func main() {
    var v1 bool = false
    var v2 int32 = 42
    v3 := []uint32{475, 2, 925, 65535}
    v4 := []uint32{255, 4294967295}
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v12 uint32 = 658
            if v1 {
                v12 = (v12 / (((v12 >> ((4294967295) & 63))) | 1))
                v13 := GenS3{F0: 140, F1: v12}
                _ = v13
            } else {
                var v14 int32 = int32(fi11)
                fmt.Println(fi11, v14)
                var v15 int8 = 100
                _ = v14
                _ = v15
            }
            _ = fi11
            _ = v12
        }
        v1 = (((0.1 / float64(3.14159)) - 1e10) != float64(1e10))
        _ = fi10
    }
    v16 := GenS2{F0: -9223372036854775808}
    v17 := GenS3{F0: 241, F1: 65535}
    switch -461 {
    case 568:
        v2 -= 1
    case 885:
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := GenS1{F0: -74}
            var v20 uint64 = 2
            _ = fi18
            _ = v19
            _ = v20
        }
    case -2147483648:
        switch "foo bar" {
        case "a":
            var v21 uint8 = 1
            _ = v21
        }
    default:
        if (100 == 1) {
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 uint16 = 416
                var v24 uint32 = 2147483647
                for fi25 := 0; fi25 < 1; fi25++ {
                    var v26 float64 = (-2.25 + (1e10 + 1e10))
                    var v27 uint16 = 42
                    var v28 uint16 = (v27 | (v23 >> ((0) & 63)))
                    v27 = (v27 - v23)
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            v4 = append(v4, 868)
        } else {
            v1 = (!(!true))
            for fi29 := 0; fi29 < 20; fi29++ {
                switch 896 {
                case 0:
                    var v30 uint32 = 4294967295
                    _ = v30
                case -1:
                    switch (fi29 - (fi29 + fi29)) {
                    case 788:
                        if (v1 && v1) {
                            var v31 uint = 564
                            var v32 uint8 = 1
                            v2 = 641
                            _ = v31
                            _ = v32
                        }
                    case 2147483647:
                        v33 := GenS4{F0: 65535, F1: 230}
                        _ = v33
                    default:
                        v34 := GenS5{F0: 2, F1: 127, F2: fi29}
                        _ = v34
                    }
                case -331:
                    for fi35 := 0; fi35 < 8; fi35++ {
                        if v1 {
                            var v36 float32 = 1e10
                            v37 := GenS6{F0: (fi35 & fi35), F1: -545, F2: 4294967295}
                            fmt.Println(v36, v2, v1)
                            _ = v36
                            _ = v37
                        } else {
                            switch 4294967295 {
                            case 1:
                                var v38 float64 = (3.14159 * -2.25)
                                _ = v38
                            }
                            switch (("\n\t" + ("hello" + "\n\t")) + "hello") {
                            case "a":
                                v1 = v1
                            }
                        }
                        v2 = (^(v2))
                        switch 1 {
                        case -2147483648:
                            v2 = (v2 & v2)
                        }
                        _ = fi35
                    }
                default:
                    v2 -= v2
                }
                if (700 <= 2) {
                    var v39 bool = (-55 >= (v2 & -502))
                    v2 += (v2 * 42)
                    for fi40 := 0; fi40 < 19; fi40++ {
                        var v41 uint = 18446744073709551615
                        var v42 uint16 = 775
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    _ = v39
                } else {
                    var v43 int32 = (v2 ^ v2)
                    if v1 {
                        for fi44 := 0; fi44 < 7; fi44++ {
                            v45 := GenS2{F0: fi29}
                            var v46 uint = 255
                            var v47 uint64 = 0
                            v48 := GenS6{F0: -2147483648, F1: 1, F2: v46}
                            _ = fi44
                            _ = v45
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        v3 = append(v3, 0)
                        fmt.Println(v2, v1, v1)
                    } else {
                        switch "hello" {
                        case "":
                            switch fi29 {
                            case -1:
                                switch ("hello" + (("a" + "") + "\n\t")) {
                                case "\n\t":
                                    _ = len(v4)
                                case "foo bar":
                                    v49 := GenS5{F0: 805, F1: genFn0(v1), F2: (fi29 ^ (^(fi29)))}
                                    _ = v49
                                }
                            case -666:
                                var v50 int64 = 42
                                _ = v50
                            }
                        }
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                        fmt.Println(v43, fi29)
                    }
                    var v51 uint16 = 65535
                    _ = v43
                    _ = v51
                }
                _ = fi29
            }
            fmt.Println(v2)
        }
    }
    if true {
        if v1 {
            v2 += 475
        }
    } else {
        for fi52 := 0; fi52 < 3; fi52++ {
            v53 := GenS5{F0: 432, F1: 0, F2: fi52}
            if (!false) {
                for fi54 := 0; fi54 < 8; fi54++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi54
                }
                var v55 float64 = 0.1
                _ = v55
            } else {
                var v56 string = "a"
                v56 = v56
                v57 := GenS3{F0: 64, F1: 432}
                _ = v56
                _ = v57
            }
            _ = fi52
            _ = v53
        }
    }
    switch ((("foo bar" + "\n\t") + "\n\t") + "foo bar") {
    case "":
        switch 877 {
        case -262:
            fmt.Println(v2)
        }
    }
    switch 1 {
    case 4294967295:
        if v1 {
            fmt.Println(v2, v2)
            v2 += (v2 - (v2 | 1))
        } else {
            for fi58 := 0; fi58 < 8; fi58++ {
                if v1 {
                    v59 := GenS7{F0: 32767}
                    _ = v59
                }
                var v60 float64 = 0.0
                if v1 {
                    var v61 int = (fi58 % ((1) | 1))
                    _ = v61
                }
                _ = fi58
                _ = v60
            }
            fmt.Println(v1)
        }
    default:
        if (true && (1 == 2)) {
            var v62 int16 = 32767
            _ = v62
        }
    }
    _ = len(v3)
    for fi63 := 0; fi63 < 2; fi63++ {
        for fi64 := 0; fi64 < 9; fi64++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi64
        }
        for fi65 := 0; fi65 < 20; fi65++ {
            v66 := GenS8{F0: fi65, F1: 702}
            _ = fi65
            _ = v66
        }
        fmt.Println(v2, v1)
        _ = fi63
    }
    var v67 uint64 = 669
    if (v1 && (2 < 255)) {
        switch "" {
        case "hello":
            var v68 uint32 = 4294967295
            _ = v68
        default:
            var v69 bool = (!(2 == 695))
            _ = v69
        }
        var v70 int64 = 42
        _ = v70
    } else {
        v67 = (v67 ^ 1)
        if v1 {
            v71 := GenS9{F0: -1, F1: 42}
            var v72 string = ("\n\t" + ("foo bar" + ("a" + "\n\t")))
            _ = v71
            _ = v72
        } else {
            _ = len(v3)
            if true {
                v67 = v67
                v73 := GenS6{F0: 890, F1: -1, F2: 609}
                _ = v73
            }
            for fi74 := 0; fi74 < 10; fi74++ {
                v1 = v1
                var v75 uint32 = 2147483648
                _ = len(v3)
                _ = fi74
                _ = v75
            }
        }
    }
    for fi76 := 0; fi76 < 4; fi76++ {
        for fi77 := 0; fi77 < 19; fi77++ {
            v78 := GenS10{F0: 1}
            _ = fi77
            _ = v78
        }
        v2 = 42
        _ = fi76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v67
}
