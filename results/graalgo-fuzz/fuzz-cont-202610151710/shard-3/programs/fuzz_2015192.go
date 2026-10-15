package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int8; F1 int32; F2 int64 }
type GenS4 struct { F0 uint64; F1 int8; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 int64; F2 uint16 }
type GenS8 struct { F0 uint64; F1 uint; F2 int32 }
func genFn0(p0 bool) int32 {
    fmt.Println(p0)
    v5 := GenS1{F0: 127, F1: 42}
    v6 := GenS1{F0: 0, F1: -1}
    v7 := GenS1{F0: 128, F1: 42}
    v8 := GenS2{F0: 42, F1: 255}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 int) uint32 {
    p0 += p0
    for fi9 := 0; fi9 < 12; fi9++ {
        if ((!(18446744073709551615 != 872)) || true) {
            var v10 float64 = 3.14159
            _ = v10
        } else {
            p0 = 274
            if false {
                p0 = 127
            }
        }
        fmt.Println(fi9, p0)
        fmt.Println(p0, fi9)
        _ = fi9
    }
    v11 := []int{-406, (-(p0))}
    _ = p0
    _ = v11
    return 4294967295
}

func main() {
    var v1 int = 444
    var v2 float64 = float64(-2.25)
    var v3 string = ("a" + "hello")
    v4 := []uint32{65535, 0, 2147483647, 65535}
    switch (^(v1)) {
    case -134:
        v1 = 9223372036854775807
    }
    _ = len(v4)
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            if false {
                if true {
                    if (42 != 18446744073709551615) {
                        var v14 int16 = -805
                        _ = v14
                    }
                    var v15 string = ("hello" + v3)
                    _ = v15
                } else {
                    switch v3 {
                    case "":
                        v3 = (v3 + "foo bar")
                    case "\n\t":
                        v2 += 1e10
                    }
                    var v16 uint16 = 65535
                    var v17 uint32 = 65535
                    _ = v16
                    _ = v17
                }
                var v18 int16 = -32768
                _ = v18
            }
            var v19 bool = false
            _ = fi13
            _ = v19
        }
        v2 += 3.14159
        _ = fi12
    }
    var v20 uint16 = 1
    switch (v1 & (v1 & 1)) {
    case 32767:
        var v21 int8 = 7
        _ = v21
    case 4294967295:
        fmt.Println(v1)
    default:
        if (!(4294967295 > 0)) {
            if false {
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 uint32 = 0
                    v24 := GenS1{F0: 1, F1: 0}
                    v25 := GenS3{F0: 42, F1: -10, F2: 1}
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            } else {
                fmt.Println(v2, v2)
                fmt.Println(v20, v1, v1)
                v4 = append(v4, 65535)
            }
            var v26 int16 = 32767
            if ((1e10 != 1.5) && (1e10 <= 0.1)) {
                var v27 int16 = 32767
                _ = v27
            }
            _ = v26
        }
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 string = ((v3 + v3) + v3)
            switch ((v30 + v3) + "\n\t") {
            case "foo bar":
                var v31 int64 = -1
                _ = v31
            case "hello":
                v32 := GenS3{F0: 42, F1: -2147483648, F2: -72}
                v1 = (fi28 ^ fi28)
                if false {
                    if (!false) {
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                    } else {
                        var v33 bool = false
                        _ = v33
                    }
                    fmt.Println(v30, v3, v2)
                    switch "\n\t" {
                    case "":
                        v3 = v30
                    case "\n\t":
                        _ = len(v4)
                    }
                }
                _ = v32
            case "\n\t":
                if ((v2 - ((v2 + 0.0) - (v2 * float64(1e10)))) < ((v2 / (0.0 / (v2 - v2))) + (v2 - 1e10))) {
                    switch v30 {
                    case "\n\t":
                        switch "hello" {
                        case "a":
                            switch v1 {
                            case 1:
                                v4 = append(v4, 1)
                            case 236:
                                var v34 uint32 = genFn1(fi28)
                                var v35 uint16 = 820
                                _ = v34
                                _ = v35
                            case -575:
                                switch ((v3 + v30) + v30) {
                                case "hello":
                                    if false {
                                        v3 = v30
                                    } else {
                                        fmt.Println(v30, v20)
                                    }
                                case "":
                                    v36 := GenS4{F0: 1, F1: 23, F2: -405}
                                    _ = v36
                                default:
                                    if ((^(v20)) < 32767) {
                                        if false {
                                            if (0 != -118) {
                                                v2 += v2
                                                var v37 int16 = 42
                                                _ = v37
                                            }
                                            switch v30 {
                                            case "hello":
                                                v3 = v3
                                            case "foo bar":
                                                v3 = (v3 + "")
                                            case "a":
                                                fmt.Println(v3, v2, fi28)
                                            default:
                                                var v38 uint8 = 255
                                                _ = v38
                                            }
                                            _ = len(v4)
                                            v20 = 32767
                                            var v39 int = 1
                                            var v40 int16 = 32767
                                            var v41 uint = 18446744073709551615
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        } else {
                                            if (v2 <= (v2 - v2)) {
                                                switch (v1 + v1) {
                                                case 438:
                                                    fmt.Println(v3, v30)
                                                case 684:
                                                    if false {
                                                        v2 += v2
                                                        var v42 int32 = -2147483648
                                                        v1 = -881
                                                        _ = v42
                                                    }
                                                default:
                                                    v43 := GenS5{F0: 715}
                                                    _ = v43
                                                }
                                                v30 = v3
                                                switch (fi28 & fi29) {
                                                case 386:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v44 := GenS1{F0: 42, F1: 0}
                                                        _ = v44
                                                    }
                                                }
                                                fmt.Println(fi28)
                                                if false {
                                                    fmt.Println(v30, fi29)
                                                    v45 := GenS5{F0: 4294967295}
                                                    _ = v45
                                                } else {
                                                    fmt.Println(fi28)
                                                    var v46 int64 = 42
                                                    _ = v46
                                                }
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = 2147483648
                                            }
                                            if (v20 == v20) {
                                                fmt.Println(v3, v2)
                                                fmt.Println(v20, v30, v3)
                                                v47 := GenS2{F0: (fi28 | v1), F1: 42}
                                                _ = v47
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        v2 -= ((v2 * 3.14159) / 1.5)
                                        fmt.Println(v1)
                                    }
                                }
                            default:
                                var v48 uint8 = 181
                                _ = v48
                            }
                        case "":
                            switch v1 {
                            case 32767:
                                var v49 uint = 1
                                _ = v49
                            case 301:
                                var v50 int16 = 1
                                _ = len(v4)
                                if (!(4294967295 >= 583)) {
                                    switch fi29 {
                                    case 128:
                                        if (42 != 0) {
                                            fmt.Println(v20)
                                            if (true && (!false)) {
                                                v50 += v50
                                                v3 = v3
                                                v51 := GenS6{F0: -3}
                                                _ = v51
                                            }
                                            v52 := GenS1{F0: 255, F1: 42}
                                            _ = v52
                                        }
                                    default:
                                        v53 := GenS7{F0: v3, F1: 9223372036854775807, F2: (^(v20))}
                                        _ = v53
                                    }
                                    v4 = append(v4, genFn1(2147483647))
                                    fmt.Println(v50, v50, v50)
                                } else {
                                    v4 = append(v4, genFn1(391))
                                    v20 += (v20 >> ((uint(v20)) & 63))
                                    v54 := GenS2{F0: int(v50), F1: 211}
                                    _ = len(v4)
                                    _ = v54
                                }
                                _ = v50
                            case 1:
                                var v55 int8 = 127
                                _ = v55
                            }
                        }
                    case "foo bar":
                        v3 = v30
                    case "a":
                        v56 := GenS6{F0: -471}
                        _ = v56
                    default:
                        v57 := GenS8{F0: 1, F1: 42, F2: 1}
                        _ = v57
                    }
                    if (!true) {
                        var v58 int32 = 2147483647
                        v2 -= v2
                        switch v1 {
                        case -1:
                            var v59 string = (v30 + "foo bar")
                            _ = v59
                        case 32768:
                            if len(v4) > 0 {
                                v4[0] = 2147483647
                            }
                        default:
                            var v60 int8 = 0
                            _ = v60
                        }
                        _ = v58
                    } else {
                        _ = len(v4)
                    }
                    var v61 int32 = -319
                    _ = v61
                }
            default:
                switch ("hello" + v30) {
                case "a":
                    _ = len(v4)
                }
            }
            fmt.Println(v3)
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v2 = float64(0.0)
    v2 += float64(-2.25)
    var v62 uint64 = 870
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v62
}
