package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 int8; F2 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 uint16; F2 int }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 int8; F1 uint32; F2 string }
func genFn0(p0 int, p1 int64, p2 bool) uint {
    var v6 int = p0
    fmt.Println(v6, p0)
    v6 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 891
}
func genFn1(p0 uint8, p1 int64, p2 uint) bool {
    switch "hello" {
    case "hello":
        fmt.Println(p0)
    case "foo bar":
        fmt.Println(p0)
    }
    v7 := GenS1{F0: 4294967295, F1: 468, F2: 0}
    if false {
        v8 := GenS2{F0: -760}
        _ = v8
    } else {
        fmt.Println(p1, p0, p0)
        v9 := []int{-2147483648, 127, 701}
        _ = v9
    }
    if false {
        fmt.Println(p2)
    }
    if (!true) {
        var v10 string = "a"
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS3{F0: (1e10 >= 1.5), F1: (p2 << ((p2) & 63)), F2: 144}
            var v13 int8 = 29
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
    } else {
        fmt.Println(p0)
        var v14 bool = false
        switch 42 {
        case 846:
            var v15 int8 = 0
            _ = v15
        case -995:
            switch ("\n\t" + ("" + (("a" + "foo bar") + "hello"))) {
            case "foo bar":
                var v16 int8 = 0
                _ = v16
            default:
                switch ("a" + "\n\t") {
                case "\n\t":
                    var v17 int8 = 21
                    _ = v17
                case "foo bar":
                    v14 = false
                default:
                    if v14 {
                        v18 := GenS4{F0: 1}
                        var v19 float32 = 0.0
                        v19 = v19
                        _ = v18
                        _ = v19
                    } else {
                        var v20 float32 = 0.1
                        v21 := GenS4{F0: (p2 << ((p2) & 63))}
                        _ = v20
                        _ = v21
                    }
                }
            }
        case 128:
            switch "" {
            case "":
                var v22 bool = false
                _ = v22
            case "hello":
                if false {
                    switch 399 {
                    case 432:
                        switch -2147483648 {
                        case -9223372036854775808:
                            var v23 uint16 = 836
                            _ = v23
                        case -1:
                            var v24 uint = p2
                            _ = v24
                        }
                    }
                    var v25 bool = true
                    var v26 uint = 937
                    _ = v25
                    _ = v26
                }
            case "\n\t":
                switch ("" + (("\n\t" + "foo bar") + "\n\t")) {
                case "foo bar":
                    for fi27 := 0; fi27 < 16; fi27++ {
                        if ((3.14159 + 3.14159) <= float32((((float64(-2.25) + 0.1) * 1.5) - ((((3.14159 + 0.1) + 0.0) - 3.14159) / (-2.25 + -2.25))))) {
                            p0 = p0
                            p0 -= uint8(p2)
                        }
                        if true {
                            v28 := GenS5{F0: 0, F1: 100, F2: -32768}
                            _ = v28
                        } else {
                            v29 := GenS1{F0: 2, F1: (p2 ^ 4294967295), F2: 4294967295}
                            for fi30 := 0; fi30 < 4; fi30++ {
                                if ("a" > "a") {
                                    v31 := GenS6{F0: "a"}
                                    _ = v31
                                }
                                var v32 float64 = 0.1
                                var v33 uint16 = 32767
                                _ = fi30
                                _ = v32
                                _ = v33
                            }
                            v34 := []int{fi27, (fi27 / ((128) | 1)), 127}
                            _ = v29
                            _ = v34
                        }
                        _ = fi27
                    }
                }
            default:
                v35 := GenS3{F0: (("hello" + "\n\t") >= "hello"), F1: uint(p0), F2: (p0 % ((p0) | 1))}
                _ = v35
            }
        default:
            p1 -= int64(p2)
        }
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return false
}

func main() {
    var v1 int = 42
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 4294967295
    var v4 uint16 = 2
    var v5 uint32 = 2
    v36 := []int{int(v4), 9223372036854775807, 4294967295}
    v4 += 32767
    v5 = (v5 + v5)
    switch v1 {
    case 32768:
        if (false || (-1 < -2147483648)) {
            var v37 int16 = -211
            switch ("hello" + "\n\t") {
            case "foo bar":
                var v38 float32 = 1.5
                _ = v38
            case "a":
                var v39 uint = 18446744073709551615
                _ = v39
            case "":
                if true {
                    switch (v1 ^ (v1 ^ (-(v1)))) {
                    case -2147483648:
                        v4 += v4
                    default:
                        v2 = v2
                    }
                    v1 -= v1
                } else {
                    if ((v5 > (v5 - v5)) || ((float32((0.0 * 1.5)) != 1e10) && true)) {
                        var v40 bool = true
                        _ = v40
                    }
                    if (-1 == (v1 - (v1 >> ((255) & 63)))) {
                        v5 = (v5 + v5)
                        _ = len(v36)
                    }
                    v41 := GenS7{F0: genFn0(2147483647, v2, ((!(2 != 0)) || true)), F1: (v4 | 65535), F2: int(v4)}
                    _ = v41
                }
            }
            _ = v37
        } else {
            var v42 int32 = 687
            _ = v42
        }
    case 9223372036854775807:
        _ = len(v36)
    case 32767:
        v43 := GenS8{F0: (v3 & uint64(v4)), F1: (-20 > 0)}
        _ = v43
    default:
        for fi44 := 0; fi44 < 19; fi44++ {
            switch "hello" {
            case "":
                var v45 int64 = v2
                _ = v45
            default:
                v46 := GenS9{F0: 127, F1: (^(v5)), F2: "hello"}
                _ = v46
            }
            var v47 int16 = 32767
            _ = fi44
            _ = v47
        }
    }
    v3 = 875
    var v48 uint8 = 255
    for fi49 := 0; fi49 < 4; fi49++ {
        var v50 bool = false
        _ = fi49
        _ = v50
    }
    for fi51 := 0; fi51 < 1; fi51++ {
        for fi52 := 0; fi52 < 12; fi52++ {
            v53 := GenS8{F0: v3, F1: false}
            var v54 int = 32768
            _ = fi52
            _ = v53
            _ = v54
        }
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v48
    _ = v5
}
