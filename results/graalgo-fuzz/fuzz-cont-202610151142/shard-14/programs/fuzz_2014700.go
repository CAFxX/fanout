package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint }
func genFn0(p0 bool) uint32 {
    v5 := GenS1{F0: 127, F1: p0}
    p0 = (!p0)
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 bool, p1 int8, p2 int32) string {
    var v6 int64 = -1
    p2 = (^(p2))
    p2 += (p2 | p2)
    v7 := []int{0, -841, -654, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 217
    var v3 bool = (249 <= 2)
    var v4 uint8 = 42
    for fi8 := 0; fi8 < 9; fi8++ {
        if false {
            if true {
                fmt.Println(v3, v2, v3)
                if v3 {
                    if (v3 || (!(v3 && true))) {
                        var v9 uint64 = 18446744073709551615
                        _ = v9
                    } else {
                        v10 := []int{4294967295}
                        var v11 uint8 = 80
                        _ = v10
                        _ = v11
                    }
                }
                v12 := GenS2{F0: 1, F1: false}
                _ = v12
            } else {
                if false {
                    if v3 {
                        v3 = v3
                        v13 := GenS3{F0: v4, F1: (v2 << ((v2) & 63))}
                        v14 := GenS4{F0: ("foo bar" + "a")}
                        _ = v13
                        _ = v14
                    } else {
                        fmt.Println(v2)
                    }
                    v15 := GenS3{F0: (v4 >> (((^(v2))) & 63)), F1: (v2 + (v2 * (v2 - v2)))}
                    var v16 int32 = -399
                    _ = v15
                    _ = v16
                }
            }
            fmt.Println(v4, v2)
        } else {
            v2 += (^(v2))
            var v17 uint8 = uint8(v2)
            if v3 {
                var v18 uint64 = 4294967295
                fmt.Println(v2)
                _ = v18
            }
            _ = v17
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := []int{(fi8 + (fi8 & -377)), (fi8 | (fi19 - fi19))}
            _ = fi19
            _ = v20
        }
        v21 := GenS3{F0: (v4 ^ (v4 + v4)), F1: 42}
        _ = fi8
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int = (fi22 | fi22)
        if (!false) {
            if true {
                if (3.14159 != 1.5) {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        switch ("foo bar" + "\n\t") {
                        case "\n\t":
                            if (0 == 1) {
                                var v25 string = "foo bar"
                                v23 += (-(fi24))
                                var v26 int = v23
                                _ = v25
                                _ = v26
                            } else {
                                if v3 {
                                    switch (("foo bar" + "foo bar") + "\n\t") {
                                    case "hello":
                                        switch v23 {
                                        case 2147483647:
                                            v2 += (v2 ^ v2)
                                        }
                                    case "foo bar":
                                        if (!(("foo bar" + "foo bar") == "\n\t")) {
                                            switch fi24 {
                                            case 4294967295:
                                                fmt.Println(v23)
                                            case 32768:
                                                fmt.Println(v23, fi24)
                                            case -541:
                                                v27 := GenS5{F0: 0, F1: v4, F2: 127}
                                                _ = v27
                                            }
                                        }
                                    }
                                    v2 = v2
                                    switch fi22 {
                                    case 829:
                                        var v28 uint8 = 1
                                        _ = v28
                                    case 128:
                                        var v29 bool = (((1e10 + 3.14159) + (0.0 - (float64(1.5) / (float64(float32(0.0)) + 1.5)))) >= float64(0.1))
                                        _ = v29
                                    case -443:
                                        var v30 bool = v3
                                        v30 = v30
                                        _ = v30
                                    default:
                                        if true {
                                            if true {
                                                if (!(!("hello" <= (("hello" + "foo bar") + "foo bar")))) {
                                                    var v31 uint = (v2 - 1)
                                                    _ = v31
                                                }
                                                v3 = (!v3)
                                                var v32 uint8 = v4
                                                _ = v32
                                            }
                                        }
                                    }
                                    fmt.Println(fi22, v23)
                                    if v3 {
                                        var v33 int8 = 100
                                        var v34 uint8 = (^(v4))
                                        if (v33 == (v33 >> ((v2) & 63))) {
                                            delete(v1, (("\n\t" + "a") + "\n\t"))
                                            var v35 uint16 = 2
                                            var v36 int32 = 42
                                            _ = v35
                                            _ = v36
                                        } else {
                                            fmt.Println(v2)
                                            var v37 int16 = -1
                                            v4 += (v4 * v4)
                                            _ = v37
                                        }
                                        var v38 uint8 = (^(v34))
                                        var v39 float32 = (-2.25 / 1.5)
                                        _ = v33
                                        _ = v34
                                        _ = v38
                                        _ = v39
                                    }
                                }
                                var v40 float64 = float64(0.1)
                                var v41 uint = 526
                                _ = v40
                                _ = v41
                            }
                        case "hello":
                            v4 = v4
                        case "":
                            if (v3 && (2147483647 < genFn0(v3))) {
                                v42 := GenS2{F0: 0, F1: true}
                                fmt.Println(v3, v3)
                                v3 = v3
                                _ = v42
                            }
                        }
                        _ = fi24
                    }
                    var v43 uint64 = 394
                    v3 = (true && v3)
                    _ = v43
                }
            } else {
                var v44 int32 = 2147483647
                if (2147483647 <= 1) {
                    v45 := GenS1{F0: -126, F1: (1 <= -472)}
                    var v46 int32 = v44
                    _ = v45
                    _ = v46
                } else {
                    switch 32767 {
                    case 2147483647:
                        for fi47 := 0; fi47 < 2; fi47++ {
                            switch "" {
                            case "":
                                v48 := []int{}
                                _ = v48
                            default:
                                _ = len(v1)
                            }
                            v2 = 862
                            _ = fi47
                        }
                    case -830:
                        if (!false) {
                            fmt.Println(v4, v3, fi22)
                        }
                    case 128:
                        v3 = (!false)
                    default:
                        v23 = int(v44)
                    }
                    var v49 int = -381
                    _ = v49
                }
                _ = v44
            }
            for fi50 := 0; fi50 < 14; fi50++ {
                var v51 uint = 65535
                _ = fi50
                _ = v51
            }
            var v52 int16 = 32767
            _ = v52
        } else {
            v53 := []int{v23, v23, v23, v23}
            v54 := GenS2{F0: 0, F1: false}
            v23 = fi22
            _ = v53
            _ = v54
        }
        if v3 {
            var v55 bool = (-273 != -1)
            _ = v55
        }
        _ = fi22
        _ = v23
    }
    v3 = true
    var v56 uint = 851
    v3 = (v3 || v3)
    v4 = v4
    fmt.Println(v2)
    if false {
        var v57 int16 = -1
        var v58 int = 128
        _ = v57
        _ = v58
    } else {
        var v59 int64 = -1
        _ = v59
    }
    if true {
        v4 = (v4 + v4)
        v60 := []int{419, 513, -48}
        _ = v60
    }
    v61 := []int{9223372036854775807, 2147483647, 157}
    v62 := v1["a"]
    v61 = append(v61, -2147483648)
    var v63 bool = ((42 >= 0) || v3)
    v56 -= (v2 >> (((v56 - (v2 - v2))) & 63))
    v4 += v4
    v64 := GenS1{F0: 1, F1: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
