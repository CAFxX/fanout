package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 bool; F2 int32 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int32 }
func genFn0() int16 {
    v7 := GenS1{F0: 1000, F1: 42}
    var v8 int = -2147483648
    v8 = -1
    _ = v7
    _ = v8
    return -785
}
func genFn1(p0 uint64, p1 int32, p2 uint16) int8 {
    var v9 float32 = (1.5 * 1e10)
    var v10 uint = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 100
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 146
    v3 := []uint32{42, 123, 42, 4294967295}
    var v4 int64 = 1
    v5 := make(map[string]string)
    v6 := make(map[int]int)
    switch "foo bar" {
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                v3 = append(v3, 1)
                v2 += (^(v2))
                v2 -= (v2 - 103)
                v4 = -9223372036854775808
                switch (-(fi12)) {
                case -2147483648:
                    if (v2 == 23) {
                        if false {
                            delete(v6, fi12)
                        } else {
                            v4 = v4
                            var v13 uint = 4294967295
                            _ = v13
                        }
                        var v14 uint8 = v2
                        var v15 int8 = 42
                        _ = v14
                        _ = v15
                    } else {
                        var v16 uint = 65535
                        var v17 bool = true
                        if false {
                            v17 = v17
                        } else {
                            if v17 {
                                fmt.Println(v4, v16, v17)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                var v18 float64 = float64(((-2.25 + 3.14159) * 0.0))
                                var v19 int8 = -71
                                v20 := GenS1{F0: 1, F1: 0}
                                v21 := GenS2{F0: 0}
                                if (0 != genFn0()) {
                                    if (v18 > 1e10) {
                                        v18 = v18
                                        switch fi12 {
                                        case 1:
                                            if (false && (100 <= 2)) {
                                                var v22 int = (fi11 % (((fi11 ^ -9223372036854775808)) | 1))
                                                switch "a" {
                                                case "a":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case "":
                                                    delete(v5, "a")
                                                case "\n\t":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                _ = v22
                                            } else {
                                                var v23 int32 = -1
                                                var v24 int8 = (v19 * -1)
                                                _ = v23
                                                _ = v24
                                            }
                                        case -2147483648:
                                            var v25 int32 = 0
                                            _ = v25
                                        }
                                    } else {
                                        var v26 int32 = 42
                                        v27 := GenS3{F0: genFn0(), F1: false, F2: v26}
                                        v28 := GenS3{F0: -1, F1: (v18 >= (v18 + 3.14159)), F2: -107}
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                _ = v18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            fmt.Println(fi12)
                            v29 := GenS4{F0: 65535, F1: (2 < 676)}
                            _ = v29
                        }
                        _ = v16
                        _ = v17
                    }
                case 32767:
                    v4 += (-(v4))
                    if ((^(v2)) > (^(v2))) {
                        var v30 uint = 0
                        if ((v4 | 42) > v4) {
                            var v31 uint8 = (v2 ^ (v2 * v2))
                            _ = v31
                        } else {
                            _ = len(v6)
                            fmt.Println(v2)
                            var v32 float64 = (float64(1.5) * -2.25)
                            if (((587 > (v4 % (((^(v4))) | 1))) || ((!true) || (-2147483648 > 42))) && false) {
                                if true {
                                    var v33 uint16 = 1
                                    var v34 int = fi12
                                    v35 := GenS5{F0: 2147483647}
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                            } else {
                                var v36 float32 = 1e10
                                fmt.Println(v30, v32, v30)
                                switch 4294967295 {
                                case 127:
                                    v3 = append(v3, 241)
                                }
                                _ = v36
                            }
                            _ = v32
                        }
                        _ = len(v3)
                        _ = v30
                    } else {
                        var v37 string = "a"
                        v38 := v5[v37]
                        _ = v37
                        _ = v38
                    }
                case -1:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if len(v3) > 0 {
                        v3[0] = 747
                    }
                }
                _ = fi12
            }
            _ = fi11
        }
    }
    var v39 int64 = (v4 & v4)
    if len(v3) > 0 {
        v3[0] = 1
    }
    switch 4294967295 {
    case -725:
        v2 += 42
    case 33:
        v2 += (v2 ^ (v2 - v2))
    }
    v40 := GenS4{F0: 4294967295, F1: (3.14159 <= (1e10 * 0.1))}
    v41 := GenS4{F0: 1, F1: (((false || false) || true) || true)}
    switch 127 {
    case 0:
        var v42 uint64 = 801
        _ = v42
    case 562:
        fmt.Println(v2, v4, v2)
    default:
        v43 := GenS4{F0: 589, F1: true}
        _ = v43
    }
    switch "foo bar" {
    case "":
        v39 = (v39 / (((v4 | (^(v4)))) | 1))
    case "\n\t":
        if (!(!(42 <= int8(v4)))) {
            var v44 float32 = 0.1
            if true {
                v39 += v39
                v45 := v5[(("a" + "foo bar") + "")]
                _ = v45
            }
            v4 = -9223372036854775808
            _ = v44
        } else {
            var v46 uint32 = 4294967295
            var v47 int = 303
            _ = v46
            _ = v47
        }
    case "hello":
        if true {
            if false {
                if (59 > v2) {
                    fmt.Println(v4)
                    for fi48 := 0; fi48 < 8; fi48++ {
                        fmt.Println(v4)
                        v4 -= (v39 - 768)
                        var v49 int64 = int64(v2)
                        _ = fi48
                        _ = v49
                    }
                } else {
                    var v50 bool = ((float64(3.14159) + 0.1) > -2.25)
                    v1[("hello" + "a")] = ((v2 % (((v2 * v2)) | 1)) >= (v2 % (((v2 >> ((122) & 63))) | 1)))
                    v4 -= (v4 << ((190) & 63))
                    _ = v50
                }
            } else {
                fmt.Println(v4, v2)
            }
            v39 = (v4 / ((v39) | 1))
            v2 = v2
        } else {
            v4 += v4
            switch 0 {
            case -559:
                v2 += (v2 >> ((42) & 63))
            case -2147483648:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 127:
                var v51 uint = 65535
                _ = v51
            default:
                for fi52 := 0; fi52 < 15; fi52++ {
                    fmt.Println(v39)
                    _ = fi52
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
