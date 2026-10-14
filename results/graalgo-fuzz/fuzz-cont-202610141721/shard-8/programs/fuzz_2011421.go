package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint }

func main() {
    var v1 int32 = 976
    var v2 bool = (false && ((true && true) && (0 > 1)))
    v3 := make(map[int]bool)
    var v4 int8 = 127
    var v5 int8 = (v4 ^ (v4 + -1))
    v6 := make(map[int]bool)
    v2 = v2
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v2, v2, v4)
        _ = fi7
    }
    if false {
        switch "a" {
        case "":
            delete(v3, -2147483648)
        case "\n\t":
            v8 := GenS1{F0: v1, F1: v4}
            _ = v8
        case "foo bar":
            if false {
                fmt.Println(v1, v4, v5)
            }
        default:
            v1 = 471
        }
    } else {
        var v9 int8 = -128
        v10 := GenS1{F0: v1, F1: (-(v9))}
        for fi11 := 0; fi11 < 16; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                switch (("a" + "\n\t") + (("foo bar" + ("\n\t" + "hello")) + "foo bar")) {
                case "hello":
                    if ((v9 >> ((1) & 63)) != 100) {
                        switch (fi11 >> ((65535) & 63)) {
                        case -1:
                            if (0 > 1) {
                                switch ("a" + "") {
                                case "foo bar":
                                    v5 = v5
                                case "a":
                                    if (-32768 == 1000) {
                                        v1 = (v1 % ((v1) | 1))
                                        var v13 uint = 1
                                        v5 = int8(v1)
                                        _ = v13
                                    }
                                    fmt.Println(v2, v4)
                                    if v2 {
                                        v14 := []int{(fi12 - (fi11 * 128)), 4294967295, (^(fi12))}
                                        v4 += 0
                                        var v15 int = 32767
                                        _ = v14
                                        _ = v15
                                    } else {
                                        switch "a" {
                                        case "a":
                                            v5 -= v9
                                        }
                                        if (v2 && (1e10 != 1e10)) {
                                            fmt.Println(v2, fi11, v5)
                                        } else {
                                            v1 = (v1 + (v1 | v1))
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v16 int = 113
                                                _ = v16
                                            case "foo bar":
                                                v2 = (true || (float32(0.1) >= (((float32(0.1) - 1.5) * (0.1 / 0.1)) + 0.0)))
                                            case "a":
                                                if (1 <= -9223372036854775808) {
                                                    var v17 int32 = 1
                                                    _ = v17
                                                }
                                                var v18 int16 = -1
                                                _ = v18
                                            default:
                                                var v19 int = -244
                                                _ = v19
                                            }
                                        }
                                    }
                                case "":
                                    v20 := GenS2{F0: 1}
                                    _ = v20
                                }
                                switch fi11 {
                                case 2147483647:
                                    if (42 <= 255) {
                                        v4 = (v4 - v9)
                                        var v21 int32 = (v1 | -958)
                                        v4 -= (v4 >> ((4294967295) & 63))
                                        _ = v21
                                    }
                                case 127:
                                    fmt.Println(fi11, v2)
                                }
                                v4 = 1
                            } else {
                                var v22 uint8 = 1
                                switch (fi11 / ((fi11) | 1)) {
                                case -2147483648:
                                    switch "\n\t" {
                                    case "a":
                                        v23 := GenS1{F0: 0, F1: v5}
                                        _ = v23
                                    case "":
                                        var v24 int32 = (v1 ^ v1)
                                        _ = v24
                                    case "hello":
                                        var v25 float32 = 1.5
                                        _ = v25
                                    default:
                                        v9 += 86
                                    }
                                case 42:
                                    if (v1 != v1) {
                                        v9 += v9
                                        switch (fi11 + -534) {
                                        case -1:
                                            var v26 uint = 854
                                            _ = v26
                                        }
                                        v27 := []int{int(v1), fi12, fi12, 1}
                                        _ = v27
                                    } else {
                                        v28 := []int{fi12}
                                        v1 = 0
                                        var v29 uint8 = v22
                                        _ = v28
                                        _ = v29
                                    }
                                default:
                                    _ = len(v3)
                                }
                                _ = v22
                            }
                        case 658:
                            var v30 int64 = 42
                            _ = v30
                        }
                        switch "\n\t" {
                        case "a":
                            if false {
                                v31 := GenS3{F0: (-(v9)), F1: false}
                                v5 = v9
                                _ = v31
                            } else {
                                if true {
                                    var v32 float64 = ((1.5 * 0.1) / (0.1 - 3.14159))
                                    switch "hello" {
                                    case "foo bar":
                                        var v33 float64 = v32
                                        _ = v33
                                    case "hello":
                                        v34 := GenS2{F0: 1}
                                        _ = v34
                                    case "a":
                                        var v35 float64 = v32
                                        _ = v35
                                    }
                                    v9 -= v5
                                    _ = v32
                                }
                                v5 = v9
                                switch (^(fi12)) {
                                case 32768:
                                    if (2 >= 0) {
                                        fmt.Println(v4, v4, v2)
                                    }
                                case -9223372036854775808:
                                    v4 += v4
                                }
                            }
                        case "\n\t":
                            v1 += v1
                        case "":
                            var v36 uint64 = 2
                            _ = v36
                        }
                    } else {
                        v5 = (^(v4))
                        v5 = v4
                        v1 += (v1 - (v1 % ((v1) | 1)))
                        var v37 float64 = 3.14159
                        if true {
                            if v2 {
                                if v2 {
                                    var v38 int = 127
                                    v39 := []int{(^(v38)), v38, (fi12 & (^(fi11))), fi12}
                                    _ = v38
                                    _ = v39
                                } else {
                                    v1 -= v1
                                    v9 += (v9 + v4)
                                }
                                v37 -= v37
                                v40 := GenS3{F0: v5, F1: (!v2)}
                                v9 -= -126
                                fmt.Println(v37, fi12, v4)
                                _ = v40
                            }
                            if ((600 <= 4294967295) && false) {
                                fmt.Println(fi11, v5)
                            } else {
                                var v41 int16 = 1
                                _ = v41
                            }
                        } else {
                            var v42 bool = (!((!v2) || (v2 && (!(!false)))))
                            if v42 {
                                var v43 int8 = -81
                                v44 := GenS4{F0: 1}
                                _ = v43
                                _ = v44
                            }
                            _ = v42
                        }
                        _ = v37
                    }
                    v5 = v9
                    var v45 int = -703
                    _ = v45
                default:
                    v5 = 100
                }
                v4 -= (v4 / (((v5 << ((876) & 63))) | 1))
                _ = fi12
            }
            v46 := v3[fi11]
            switch fi11 {
            case 785:
                v47 := GenS2{F0: 4294967295}
                _ = v47
            case -212:
                var v48 int16 = 1000
                _ = v48
            default:
                var v49 uint8 = 255
                _ = v49
            }
            _ = fi11
            _ = v46
        }
        _ = v10
        _ = v9
    }
    fmt.Println(v4, v2)
    var v50 float32 = 1.5
    var v51 int16 = 42
    _ = len(v6)
    v1 -= 63
    delete(v6, 4294967295)
    for fi52 := 0; fi52 < 19; fi52++ {
        _ = len(v3)
        delete(v6, (fi52 ^ (fi52 * 2147483647)))
        var v53 uint16 = 1
        _ = fi52
        _ = v53
    }
    v4 = (v5 >> ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
