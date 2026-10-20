package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint; F2 bool }
type GenS5 struct { F0 uint64; F1 bool; F2 uint }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 int }
func genFn0(p0 int64, p1 uint8) uint16 {
    if (-1 < 533) {
        fmt.Println(p0, p1)
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: p1}
            _ = fi6
            _ = v7
        }
    }
    p0 = (-(p0))
    v8 := GenS1{F0: (^(p1))}
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(fi9, p1, p0)
        switch (fi9 / ((-1) | 1)) {
        case -82:
            p1 += (^(p1))
        }
        if (!(!(p0 == (-(p0))))) {
            var v10 uint32 = 907
            switch (fi9 * (^(fi9))) {
            case -970:
                for fi11 := 0; fi11 < 4; fi11++ {
                    v12 := []int{fi11, 128, 1}
                    var v13 int32 = -2147483648
                    v10 += 608
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 103:
                v14 := []int{0, (fi9 - fi9)}
                _ = v14
            }
            if true {
                fmt.Println(p0, fi9, p1)
            }
            _ = v10
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return 601
}
func genFn1() uint16 {
    var v15 uint8 = 2
    v16 := GenS1{F0: (^(v15))}
    fmt.Println(v15)
    v17 := GenS1{F0: v15}
    var v18 uint32 = 187
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    return 1
}

func main() {
    var v1 bool = (-1 <= -69)
    var v2 uint16 = 217
    v3 := make(map[string]int)
    var v4 int32 = int32(v2)
    var v5 int64 = -19
    fmt.Println(v1, v5, v2)
    v19 := GenS1{F0: 1}
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := GenS1{F0: 127}
        v1 = true
        if true {
            for fi22 := 0; fi22 < 18; fi22++ {
                if (v5 == (v5 | (^(v5)))) {
                    var v23 float32 = 0.1
                    _ = v23
                } else {
                    v24 := GenS1{F0: 172}
                    v25 := GenS2{F0: 113}
                    _ = v24
                    _ = v25
                }
                v26 := GenS3{F0: 0}
                _ = fi22
                _ = v26
            }
            switch (^(fi20)) {
            case -1:
                var v27 float64 = (-2.25 - 3.14159)
                _ = v27
            default:
                switch (fi20 >> ((2) & 63)) {
                case 128:
                    var v28 int64 = v5
                    _ = v28
                }
            }
            var v29 string = "hello"
            _ = v29
        } else {
            v30 := GenS4{F0: ((("foo bar" + "foo bar") + "\n\t") + ("a" + "hello")), F1: 1, F2: (genFn0(v5, 106) < (^(v2)))}
            if v1 {
                if false {
                    var v31 float64 = (0.1 + -2.25)
                    var v32 float32 = 1e10
                    var v33 uint32 = 65535
                    _ = v31
                    _ = v32
                    _ = v33
                }
            }
            switch "\n\t" {
            case "foo bar":
                v34 := GenS5{F0: 0, F1: (v1 || v1), F2: 42}
                _ = v34
            }
            _ = v30
        }
        _ = fi20
        _ = v21
    }
    switch -9223372036854775808 {
    case -2147483648:
        v5 = v5
    case 0:
        v35 := GenS4{F0: "foo bar", F1: 255, F2: false}
        _ = v35
    case -403:
        fmt.Println(v1, v1, v2)
    }
    var v36 int = 4294967295
    if v1 {
        if (0 < 147) {
            v4 = (v4 - (v4 + (v4 * (v4 | v4))))
        }
        for fi37 := 0; fi37 < 8; fi37++ {
            switch ("\n\t" + ("hello" + ((("hello" + ("foo bar" + "")) + "") + "\n\t"))) {
            case "hello":
                switch ("" + ("\n\t" + "a")) {
                case "foo bar":
                    var v38 float32 = float32(3.14159)
                    _ = v38
                case "hello":
                    for fi39 := 0; fi39 < 12; fi39++ {
                        var v40 bool = (!v1)
                        v2 = 478
                        _ = fi39
                        _ = v40
                    }
                default:
                    switch 42 {
                    case -2147483648:
                        v4 = -939
                    case 0:
                        var v41 uint64 = 18446744073709551615
                        _ = v41
                    }
                }
            case "":
                if (((!v1) && v1) && (!v1)) {
                    v1 = false
                    v42 := GenS6{F0: v36, F1: -1}
                    _ = v42
                }
            case "a":
                var v43 uint64 = 18446744073709551615
                _ = v43
            default:
                for fi44 := 0; fi44 < 3; fi44++ {
                    if v1 {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(fi44, v2)
                        default:
                            switch -504 {
                            case 2147483647:
                                switch "\n\t" {
                                case "":
                                    if (!(!true)) {
                                        var v45 int = fi37
                                        v46 := GenS4{F0: "foo bar", F1: 65535, F2: ((v5 / ((42) | 1)) <= v5)}
                                        v4 = 330
                                        _ = v45
                                        _ = v46
                                    }
                                case "\n\t":
                                    if ((-(v4)) <= v4) {
                                        var v47 int = 9223372036854775807
                                        v5 += (v5 >> ((2) & 63))
                                        delete(v3, ((("foo bar" + "\n\t") + "") + "hello"))
                                        _ = v47
                                    } else {
                                        if v1 {
                                            v48 := GenS5{F0: 1, F1: ((fi37 % ((127) | 1)) == v36), F2: 42}
                                            _ = v48
                                        }
                                        v49 := []int{fi37, 1}
                                        _ = v49
                                    }
                                default:
                                    v36 = fi37
                                }
                            case -2147483648:
                                v1 = v1
                            default:
                                v5 -= v5
                            }
                        }
                    } else {
                        var v50 uint16 = (v2 | v2)
                        switch "foo bar" {
                        case "\n\t":
                            v2 = (v50 + v50)
                            fmt.Println(v36)
                            var v51 int8 = -57
                            fmt.Println(fi44, v1, v50)
                            var v52 uint = 255
                            v2 = v50
                            _ = v51
                            _ = v52
                        case "hello":
                            var v53 float32 = (1e10 + 1.5)
                            v54 := GenS4{F0: "a", F1: 0, F2: (1.5 == v53)}
                            _ = v53
                            _ = v54
                        case "foo bar":
                            v55 := []int{}
                            _ = v55
                        }
                        var v56 uint = 65535
                        var v57 uint16 = (v2 << (((v56 % ((uint(v2)) | 1))) & 63))
                        _ = v50
                        _ = v56
                        _ = v57
                    }
                    _ = fi44
                }
            }
            _ = fi37
        }
    } else {
        for fi58 := 0; fi58 < 17; fi58++ {
            v1 = v1
            switch ("" + ("a" + "a")) {
            case "a":
                v59 := GenS7{F0: 2}
                _ = v59
            case "":
                var v60 uint8 = 255
                _ = v60
            default:
                var v61 string = "foo bar"
                _ = v61
            }
            v62 := GenS8{F0: (v4 - (v4 % ((v4) | 1)))}
            _ = fi58
            _ = v62
        }
    }
    v63 := GenS9{F0: 218, F1: v4}
    fmt.Println(v4, v36)
    for fi64 := 0; fi64 < 1; fi64++ {
        v65 := GenS10{F0: fi64}
        _ = fi64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v63
}
