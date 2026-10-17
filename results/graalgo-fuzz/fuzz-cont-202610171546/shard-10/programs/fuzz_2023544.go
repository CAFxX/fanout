package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 bool; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 string; F2 int32 }
type GenS7 struct { F0 uint32; F1 int16; F2 int }
type GenS8 struct { F0 bool; F1 uint16 }
func genFn0() int32 {
    var v4 int = 32767
    switch ("foo bar" + (("a" + "a") + "foo bar")) {
    case "":
        v4 = v4
    case "foo bar":
        var v5 float64 = 1.5
        _ = v5
    }
    var v6 uint64 = 2
    for fi7 := 0; fi7 < 5; fi7++ {
        v6 = 4294967295
        _ = fi7
    }
    v8 := GenS1{F0: 0, F1: 2147483647}
    _ = v4
    _ = v6
    _ = v8
    return int32(v6)
}
func genFn1(p0 int8, p1 uint, p2 int64) uint {
    v9 := []int{-233, 128, 9223372036854775807}
    p1 -= p1
    var v10 int16 = 42
    var v11 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return (p1 | p1)
}

func main() {
    v1 := []uint32{1, 4294967295}
    var v2 int = -854
    v3 := make(map[string]int)
    switch "\n\t" {
    case "foo bar":
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int = (fi12 ^ fi12)
            for fi14 := 0; fi14 < 18; fi14++ {
                fmt.Println(fi12, fi12, v2)
                v15 := v3[("hello" + ("" + "\n\t"))]
                v2 += (-(v2))
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    case "hello":
        if (!(true && (false && true))) {
            v2 = v2
            v2 += (v2 | v2)
            fmt.Println(v2)
        } else {
            var v16 int32 = 42
            v17 := v3["foo bar"]
            var v18 uint64 = 1
            _ = v16
            _ = v17
            _ = v18
        }
    case "\n\t":
        fmt.Println(v2)
    default:
        fmt.Println(v2)
    }
    v19 := GenS1{F0: 127, F1: 42}
    var v20 uint32 = 4294967295
    var v21 int = v2
    if true {
        if (((genFn1(-128, 0, -9223372036854775808) == 0) && (100 >= 1)) && false) {
            fmt.Println(v20)
            if true {
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 int8 = 100
                    _ = fi22
                    _ = v23
                }
            } else {
                v24 := v3["foo bar"]
                v25 := GenS2{F0: (3.14159 > 0.0)}
                _ = v24
                _ = v25
            }
            switch v2 {
            case -256:
                v26 := v3["foo bar"]
                _ = v26
            case 945:
                fmt.Println(v20, v20, v2)
            }
        }
    }
    var v27 float32 = float32((3.14159 / 1e10))
    v27 -= 0.0
    v20 = 0
    if false {
        if false {
            var v28 uint16 = 964
            _ = v28
        } else {
            var v29 int64 = -692
            var v30 uint = 18446744073709551615
            _ = v29
            _ = v30
        }
        v31 := GenS3{F0: ""}
        _ = v31
    } else {
        fmt.Println(v20, v21)
        for fi32 := 0; fi32 < 3; fi32++ {
            fmt.Println(v2, fi32, v20)
            if false {
                if true {
                    var v33 uint64 = 18446744073709551615
                    var v34 int32 = -2147483648
                    _ = v33
                    _ = v34
                }
                v1 = append(v1, v20)
                for fi35 := 0; fi35 < 1; fi35++ {
                    if false {
                        if false {
                            if ((!true) || ((v2 & 128) < v2)) {
                                if true {
                                    fmt.Println(v2)
                                    var v36 float64 = 0.1
                                    _ = v36
                                }
                                v27 = (v27 + (v27 / 1.5))
                            } else {
                                var v37 float32 = 3.14159
                                v27 = v37
                                _ = v37
                            }
                            var v38 uint32 = 65535
                            _ = v38
                        }
                        v20 = v20
                        switch "" {
                        case "foo bar":
                            fmt.Println(fi32)
                        case "a":
                            fmt.Println(fi35, fi32, fi32)
                        default:
                            fmt.Println(fi35, fi32, v2)
                        }
                        fmt.Println(fi32, v2)
                    } else {
                        var v39 uint8 = 2
                        if (v39 == v39) {
                            var v40 int8 = -96
                            v40 += v40
                            _ = v40
                        } else {
                            fmt.Println(fi35, v21, fi32)
                        }
                        v41 := GenS4{F0: 9223372036854775807, F1: (v39 > 139), F2: 0}
                        _ = v39
                        _ = v41
                    }
                    var v42 uint32 = v20
                    _ = fi35
                    _ = v42
                }
            } else {
                if true {
                    switch v21 {
                    case 64:
                        fmt.Println(v20, v27, v20)
                    case -286:
                        for fi43 := 0; fi43 < 16; fi43++ {
                            v44 := GenS5{F0: "\n\t"}
                            v45 := GenS6{F0: false, F1: ("" + "a"), F2: 0}
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                    }
                    fmt.Println(v20, fi32, v27)
                } else {
                    if false {
                        var v46 int8 = -116
                        v47 := GenS7{F0: v20, F1: -1, F2: (fi32 + (v21 ^ (fi32 % (((-(v21))) | 1))))}
                        _ = v46
                        _ = v47
                    } else {
                        var v48 uint = 505
                        for fi49 := 0; fi49 < 8; fi49++ {
                            var v50 uint64 = 18446744073709551615
                            v51 := GenS8{F0: ((v50 + v50) != (^(v50))), F1: 65535}
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                        var v52 uint = (^(v48))
                        _ = v48
                        _ = v52
                    }
                }
            }
            var v53 bool = false
            _ = fi32
            _ = v53
        }
    }
    for fi54 := 0; fi54 < 8; fi54++ {
        var v55 uint = 936
        _ = fi54
        _ = v55
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        if true {
            var v57 bool = true
            fmt.Println(v20)
            _ = len(v1)
            _ = v57
        }
        var v58 int = 1
        _ = fi56
        _ = v58
    }
    var v59 uint = 4294967295
    fmt.Println(v59)
    v60 := GenS6{F0: (!((v20 | v20) < 0)), F1: ("hello" + ("hello" + (("" + "hello") + ("a" + "")))), F2: -643}
    var v61 uint16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v59
    _ = v60
    _ = v61
}
