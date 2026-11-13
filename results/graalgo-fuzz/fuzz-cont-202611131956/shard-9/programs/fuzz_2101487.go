package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint32; F1 int; F2 bool }
type GenS4 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS5 struct { F0 int32; F1 bool }
func genFn0(p0 int8, p1 int, p2 int8) uint64 {
    var v7 bool = (549 > -2147483648)
    var v8 uint16 = 32767
    switch "hello" {
    case "hello":
        p1 -= p1
    case "":
        fmt.Println(v7, p1)
    case "a":
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 328
}

func main() {
    var v1 int16 = -600
    v2 := []int{-1, -129, 0}
    var v3 float64 = -2.25
    v4 := make(map[string]string)
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]uint32)
    v9 := GenS1{F0: ((("a" + ("foo bar" + "hello")) + "hello") + "a"), F1: 1}
    var v10 int64 = -9223372036854775808
    if len(v2) > 0 {
        v2[0] = 498
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 float64 = 1.5
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 float64 = 1.5
            _ = fi14
            _ = v15
        }
        if (true && false) {
            var v16 int32 = 1
            v17 := GenS2{F0: v10, F1: 0}
            _ = v16
            _ = v17
        }
        _ = fi11
    }
    fmt.Println(v10)
    v1 = v1
    switch -1 {
    case 1:
        v3 = (v3 / v3)
    case 128:
        switch ("" + "hello") {
        case "a":
            switch "hello" {
            case "foo bar":
                v1 = v1
            default:
                var v18 float32 = (0.1 - 0.0)
                _ = v18
            }
        case "hello":
            if true {
                if (v10 < v10) {
                    v3 -= v3
                    for fi19 := 0; fi19 < 1; fi19++ {
                        var v20 uint16 = 2
                        _ = fi19
                        _ = v20
                    }
                }
            } else {
                switch "\n\t" {
                case "":
                    fmt.Println(v10, v10, v3)
                }
                var v21 int8 = -1
                _ = v21
            }
        }
    case 741:
        delete(v4, ("hello" + "hello"))
    default:
        if (false && (178 > 2)) {
            v22 := GenS3{F0: 29, F1: 2147483647, F2: (false || true)}
            _ = v22
        } else {
            if false {
                fmt.Println(v10, v10, v10)
                v3 += (((v3 - 3.14159) * v3) + 0.0)
            } else {
                if false {
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        var v23 int8 = 115
                        _ = v23
                    case "a":
                        var v24 bool = true
                        _ = v24
                    case "hello":
                        v1 = v1
                    }
                }
            }
        }
    }
    var v25 uint8 = 128
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 float32 = ((0.1 / 1.5) + (1.5 - 0.1))
        v2 = append(v2, (fi26 | fi26))
        _ = fi26
        _ = v27
    }
    v28 := GenS3{F0: 101, F1: 42, F2: (true && false)}
    delete(v6, 722)
    if ((!((false || (0 != 1)) || true)) && true) {
        var v29 int64 = (v10 & v10)
        v30 := GenS3{F0: 546, F1: 127, F2: true}
        _ = v29
        _ = v30
    }
    if ("hello" == (("hello" + (("a" + ("foo bar" + "foo bar")) + "foo bar")) + "hello")) {
        if (!false) {
            var v31 int8 = -128
            var v32 int = int(v31)
            v33 := v4[((("foo bar" + "") + "foo bar") + "a")]
            _ = v31
            _ = v32
            _ = v33
        }
        v34 := GenS4{F0: v25, F1: 394, F2: 2147483647}
        _ = v34
    } else {
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 uint64 = 2
            switch fi35 {
            case -9223372036854775808:
                v1 = (v1 + 42)
            case 720:
                v37 := GenS4{F0: 42, F1: 0, F2: 42}
                _ = v37
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            switch "\n\t" {
            case "foo bar":
                for fi38 := 0; fi38 < 4; fi38++ {
                    var v39 string = "foo bar"
                    _ = fi38
                    _ = v39
                }
            case "a":
                if (true || (744 <= 312)) {
                    fmt.Println(v25, v36, v3)
                    var v40 uint = 42
                    if true {
                        switch 32767 {
                        case 9223372036854775807:
                            v41 := GenS3{F0: 42, F1: (fi35 / ((128) | 1)), F2: false}
                            _ = v41
                        default:
                            var v42 float32 = 1.5
                            _ = v42
                        }
                        if (0 >= -767) {
                            var v43 int32 = -1
                            if true {
                                switch fi35 {
                                case -198:
                                    switch (^(fi35)) {
                                    case 162:
                                        delete(v6, 220)
                                    case 323:
                                        fmt.Println(v10)
                                    }
                                case 578:
                                    v44 := GenS5{F0: v43, F1: (!false)}
                                    _ = v44
                                }
                                for fi45 := 0; fi45 < 17; fi45++ {
                                    _ = len(v2)
                                    fmt.Println(v40)
                                    var v46 int8 = -128
                                    _ = fi45
                                    _ = v46
                                }
                            }
                            _ = v43
                        }
                        var v47 uint64 = v36
                        _ = v47
                    }
                    _ = v40
                }
            case "hello":
                v36 -= uint64(v25)
            }
            _ = fi35
            _ = v36
        }
        v3 = v3
    }
    var v48 bool = true
    fmt.Println(v25, v25, v10)
    if (v48 || (1 != 2)) {
        switch -2147483648 {
        case 4294967295:
            v25 = (v25 << ((uint(v25)) & 63))
        }
    }
    if v48 {
        v3 += 0.1
        if (32767 <= v1) {
            v25 -= (v25 - (v25 | v25))
            v25 = (v25 << ((473) & 63))
            var v49 int64 = v10
            _ = v49
        }
    }
    switch 42 {
    case 0:
        var v50 uint32 = uint32(v25)
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v9
}
