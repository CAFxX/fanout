package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint; F1 uint64; F2 int64 }
type GenS9 struct { F0 int64; F1 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint32; F1 int16 }
type GenS12 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = -2.25
    v3 := make(map[int]bool)
    var v4 float64 = (3.14159 - ((0.0 * 0.1) * 1e10))
    var v5 int8 = 100
    var v6 int32 = 1
    fmt.Println(v2, v2, v4)
    switch 32767 {
    case 0:
        v6 += v6
    case 2147483647:
        v7 := []int{}
        _ = v7
    default:
        fmt.Println(v2, v6)
    }
    v8 := GenS1{F0: v5, F1: 596}
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    fmt.Println(v4, v2, v4)
    v2 += (v2 - v2)
    if (!true) {
        if (!true) {
            if true {
                for fi11 := 0; fi11 < 3; fi11++ {
                    if (((0 <= 255) && (!(0.0 > (0.1 - -2.25)))) || false) {
                        if false {
                            if ((true || (!false)) && false) {
                                if (((-(v5)) == (v5 / (((v5 % ((v5) | 1))) | 1))) || false) {
                                    var v12 bool = false
                                    if false {
                                        switch ("" + (("\n\t" + "foo bar") + "\n\t")) {
                                        case "\n\t":
                                            var v13 uint = 2
                                            _ = v13
                                        case "foo bar":
                                            switch (fi11 >> ((2) & 63)) {
                                            case -168:
                                                switch fi11 {
                                                case -663:
                                                    switch (("foo bar" + ("" + "")) + "foo bar") {
                                                    case "a":
                                                        fmt.Println(v2, fi11, v2)
                                                    case "\n\t":
                                                        v14 := GenS2{F0: 23}
                                                        _ = v14
                                                    default:
                                                        v15 := GenS3{F0: v12, F1: 374}
                                                        _ = v15
                                                    }
                                                default:
                                                    if v12 {
                                                        for fi16 := 0; fi16 < 5; fi16++ {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v6 += 42
                                                            case "":
                                                                if (1 != 0) {
                                                                    v2 = 0.0
                                                                    if (867 != 115) {
                                                                        var v17 int = 0
                                                                        _ = v17
                                                                    }
                                                                }
                                                            default:
                                                                v18 := GenS1{F0: v5, F1: 201}
                                                                _ = v18
                                                            }
                                                            _ = fi16
                                                        }
                                                        v6 -= v6
                                                    } else {
                                                        if (255 != 129) {
                                                            for fi19 := 0; fi19 < 20; fi19++ {
                                                                var v20 float32 = ((v2 * 1e10) - (v2 * float32(1e10)))
                                                                v2 += float32(v4)
                                                                _ = fi19
                                                                _ = v20
                                                            }
                                                        }
                                                    }
                                                }
                                            case 127:
                                                v21 := GenS4{F0: 2147483647, F1: 255, F2: 32767}
                                                _ = v21
                                            case 4294967295:
                                                if (-6 == v5) {
                                                    v1[1] = -9223372036854775808
                                                } else {
                                                    fmt.Println(fi11, v5, v12)
                                                    switch "a" {
                                                    case "\n\t":
                                                        v2 -= v2
                                                    default:
                                                        for fi22 := 0; fi22 < 10; fi22++ {
                                                            var v23 uint64 = 2
                                                            _ = fi22
                                                            _ = v23
                                                        }
                                                    }
                                                    fmt.Println(v4, v5, v4)
                                                }
                                            }
                                        }
                                    }
                                    for fi24 := 0; fi24 < 5; fi24++ {
                                        v4 -= 0.1
                                        _ = fi24
                                    }
                                    _ = v12
                                } else {
                                    for fi25 := 0; fi25 < 16; fi25++ {
                                        var v26 uint = 18446744073709551615
                                        _ = fi25
                                        _ = v26
                                    }
                                    var v27 uint32 = 65535
                                    var v28 bool = true
                                    _ = v27
                                    _ = v28
                                }
                                fmt.Println(v6, v6)
                            }
                            var v29 uint64 = 42
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 1; fi30++ {
                            var v31 uint8 = 42
                            v32 := GenS5{F0: 0, F1: -1, F2: 241}
                            var v33 uint64 = 0
                            fmt.Println(v4)
                            _ = fi30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS6{F0: 365, F1: -9223372036854775808}
                        _ = v34
                    }
                    if (true || (false || false)) {
                        for fi35 := 0; fi35 < 17; fi35++ {
                            v36 := GenS7{F0: 42}
                            var v37 uint = 329
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                        delete(v1, 221)
                        v2 = v2
                    }
                    v2 = float32(float64((3.14159 - float32(v4))))
                    _ = fi11
                }
                v2 = 0.1
            }
        }
        if false {
            for fi38 := 0; fi38 < 16; fi38++ {
                var v39 uint32 = 2
                _ = fi38
                _ = v39
            }
            switch -1 {
            case 1:
                var v40 int8 = (-(v5))
                _ = v40
            case 127:
                if true {
                    fmt.Println(v5, v6)
                    for fi41 := 0; fi41 < 14; fi41++ {
                        v42 := GenS8{F0: 65535, F1: 935, F2: -878}
                        _ = fi41
                        _ = v42
                    }
                    for fi43 := 0; fi43 < 10; fi43++ {
                        switch fi43 {
                        case 0:
                            if false {
                                for fi44 := 0; fi44 < 12; fi44++ {
                                    var v45 uint64 = 42
                                    var v46 uint32 = 2147483648
                                    fmt.Println(v2, v45, v45)
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                                var v47 uint = 2
                                _ = len(v1)
                                _ = v47
                            }
                        }
                        _ = fi43
                    }
                }
            default:
                if true {
                    v4 += 3.14159
                }
            }
            if ((true && (2 >= 0)) && (-2.25 <= ((3.14159 - 0.1) * 3.14159))) {
                v4 += v4
                var v48 uint64 = 0
                if false {
                    v49 := GenS9{F0: 705, F1: 42}
                    for fi50 := 0; fi50 < 3; fi50++ {
                        var v51 int32 = (v6 % ((v6) | 1))
                        v6 = v51
                        v52 := GenS6{F0: 65535, F1: (fi50 / (((fi50 + (fi50 << ((853) & 63)))) | 1))}
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    if (49 != 18446744073709551615) {
                        for fi53 := 0; fi53 < 10; fi53++ {
                            v54 := GenS10{F0: (float32(v4) == v2)}
                            _ = fi53
                            _ = v54
                        }
                        v4 = v4
                        var v55 int = 42
                        _ = v55
                    }
                    _ = v49
                } else {
                    v6 -= (v6 / ((2147483647) | 1))
                    if (v4 == float64(v2)) {
                        v56 := GenS11{F0: 2, F1: -1}
                        _ = v56
                    } else {
                        var v57 int = 0
                        var v58 int = -802
                        _ = v57
                        _ = v58
                    }
                    var v59 float32 = v2
                    _ = v59
                }
                _ = v48
            } else {
                var v60 uint16 = 401
                delete(v3, 240)
                _ = v60
            }
        } else {
            var v61 string = ("a" + "hello")
            _ = v61
        }
        v5 -= 100
    } else {
        v62 := GenS12{F0: 128}
        fmt.Println(v4, v5)
        fmt.Println(v4)
        _ = v62
    }
    var v63 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v8
}
