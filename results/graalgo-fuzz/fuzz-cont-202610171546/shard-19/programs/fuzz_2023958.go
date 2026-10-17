package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int; F2 int64 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int8 }
func genFn0(p0 int64) uint16 {
    p0 += p0
    for fi6 := 0; fi6 < 2; fi6++ {
        if true {
            switch (("\n\t" + ("a" + ("\n\t" + ((("\n\t" + "foo bar") + "") + "hello")))) + "a") {
            case "foo bar":
                var v7 int32 = -2147483648
                _ = v7
            }
            var v8 int64 = (p0 & -1)
            _ = v8
        } else {
            v9 := GenS1{F0: 2}
            var v10 int = 922
            var v11 float32 = (0.1 + 0.1)
            _ = v10
            _ = v11
            _ = v9
        }
        switch "foo bar" {
        case "hello":
            for fi12 := 0; fi12 < 17; fi12++ {
                v13 := GenS2{F0: -2147483648}
                v14 := GenS1{F0: 727}
                var v15 uint16 = 32767
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        case "":
            v16 := GenS3{F0: 42, F1: (fi6 & fi6), F2: (p0 | p0)}
            _ = v16
        case "foo bar":
            if false {
                p0 -= p0
            }
        default:
            var v17 int32 = 0
            _ = v17
        }
        _ = fi6
    }
    p0 += (-(p0))
    _ = p0
    return 32767
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]bool)
    v3 := []uint32{982, 626, 42}
    v4 := make(map[string]bool)
    var v5 float64 = (0.0 + (((1e10 * float64(v1)) - 1.5) * float64(((((v1 + (float32(0.1) + 0.0)) / (v1 / v1)) + float32(3.14159)) * v1))))
    var v18 float64 = -2.25
    var v19 uint32 = 110
    v20 := GenS2{F0: 1}
    var v21 int = -9223372036854775808
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 string = (("a" + "\n\t") + "\n\t")
            v25 := GenS2{F0: -1}
            var v26 string = "foo bar"
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi22
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        if true {
            v28 := GenS3{F0: (^(v19)), F1: fi27, F2: -738}
            _ = v28
        }
        var v29 float64 = (1.5 - (0.0 / v18))
        _ = fi27
        _ = v29
    }
    var v30 uint16 = 32767
    for fi31 := 0; fi31 < 18; fi31++ {
        delete(v2, (v19 & v19))
        if (1 != uint64(v19)) {
            if true {
                v1 -= v1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 int32 = -2147483648
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    v18 += -2.25
    v18 += float64((v1 - 0.1))
    switch v21 {
    case -979:
        v30 = (v30 << ((4294967295) & 63))
    case 9223372036854775807:
        for fi34 := 0; fi34 < 18; fi34++ {
            v1 = v1
            _ = fi34
        }
    case 876:
        switch 929 {
        case 42:
            switch ("" + "\n\t") {
            case "foo bar":
                _ = len(v4)
            }
        case 2147483647:
            v3 = append(v3, v19)
        case 728:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            if (!true) {
                switch "hello" {
                case "\n\t":
                    if (991 != v21) {
                        var v35 int8 = 65
                        if (v18 <= v18) {
                            var v36 int16 = 1000
                            v37 := GenS4{F0: 453, F1: (v21 | 282)}
                            for fi38 := 0; fi38 < 20; fi38++ {
                                for fi39 := 0; fi39 < 8; fi39++ {
                                    delete(v4, "foo bar")
                                    v3 = append(v3, (v19 ^ v19))
                                    if true {
                                        var v40 uint8 = 1
                                        _ = v40
                                    }
                                    _ = fi39
                                }
                                var v41 float32 = -2.25
                                _ = fi38
                                _ = v41
                            }
                            _ = v36
                            _ = v37
                        }
                        v42 := GenS2{F0: 419}
                        _ = v35
                        _ = v42
                    }
                case "hello":
                    var v43 int32 = 0
                    _ = v43
                case "foo bar":
                    for fi44 := 0; fi44 < 19; fi44++ {
                        for fi45 := 0; fi45 < 10; fi45++ {
                            switch "a" {
                            case "":
                                var v46 float32 = v1
                                _ = v46
                            case "hello":
                                v19 += uint32(v30)
                                v47 := GenS5{F0: -128}
                                v48 := v4["a"]
                                _ = v47
                                _ = v48
                            default:
                                v1 += 0.1
                            }
                            _ = fi45
                        }
                        v3 = append(v3, 0)
                        _ = fi44
                    }
                }
            } else {
                fmt.Println(v30, v1)
                var v49 uint = 0
                _ = v49
            }
        }
    default:
        if false {
            _ = len(v3)
            for fi50 := 0; fi50 < 1; fi50++ {
                fmt.Println(v5, v18)
                var v51 uint16 = 706
                _ = fi50
                _ = v51
            }
            v52 := GenS1{F0: 4294967295}
            _ = v52
        }
    }
    var v53 uint64 = 1
    v18 += v5
    var v54 int64 = 0
    for fi55 := 0; fi55 < 12; fi55++ {
        if (496 <= v30) {
            var v56 int32 = -1
            var v57 int16 = int16(v30)
            var v58 bool = false
            _ = v56
            _ = v57
            _ = v58
        }
        _ = fi55
    }
    var v59 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v59
}
