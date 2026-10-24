package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 uint64; F2 int8 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint8; F1 uint64 }
func genFn0() int64 {
    v6 := GenS1{F0: -9223372036854775808}
    var v7 string = "a"
    switch 128 {
    case 42:
        v7 = v7
    case -1:
        var v8 uint = 18446744073709551615
        _ = v8
    default:
        var v9 uint = 4294967295
        _ = v9
    }
    _ = v6
    _ = v7
    return -252
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 771
    var v3 uint64 = 346
    v4 := make(map[int]string)
    var v5 uint64 = v3
    var v10 int32 = -2147483648
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS2{F0: 0}
        var v13 string = (("\n\t" + "") + "hello")
        _ = fi11
        _ = v12
        _ = v13
    }
    switch "foo bar" {
    case "a":
        var v14 uint8 = 128
        _ = v14
    case "":
        v10 -= v10
    case "foo bar":
        var v15 uint16 = 2
        _ = v15
    default:
        v16 := GenS3{F0: 2}
        _ = v16
    }
    if true {
        if ((2147483647 == (v10 * 1)) && (v10 < (v10 | (v10 / ((42) | 1))))) {
            v10 = (v10 >> ((265) & 63))
        } else {
            fmt.Println(v3, v3, v5)
        }
        if false {
            v10 += 679
            var v17 int32 = v10
            v18 := GenS4{F0: false, F1: (v3 & (v3 & v3)), F2: 42}
            _ = v17
            _ = v18
        } else {
            v19 := []int{-2147483648, -256}
            for fi20 := 0; fi20 < 1; fi20++ {
                v3 = 40
                for fi21 := 0; fi21 < 15; fi21++ {
                    fmt.Println(v2, v5)
                    if true {
                        fmt.Println(fi21, fi20, v3)
                        if len(v19) > 0 {
                            v19[0] = fi21
                        }
                        switch (fi21 | (fi20 + fi21)) {
                        case -9223372036854775808:
                            var v22 string = ""
                            _ = v22
                        case -699:
                            var v23 string = "hello"
                            _ = v23
                        case 127:
                            switch (fi20 ^ (fi21 % (((fi21 << ((4294967295) & 63))) | 1))) {
                            case 127:
                                v5 = v3
                            case 1:
                                var v24 uint16 = 42
                                _ = v24
                            }
                        default:
                            switch ("hello" + "a") {
                            case "a":
                                switch fi21 {
                                case -2147483648:
                                    var v25 int8 = 0
                                    _ = v25
                                case 42:
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                default:
                                    switch (-(fi20)) {
                                    case 400:
                                        var v26 uint32 = 4294967295
                                        _ = v26
                                    case 9223372036854775807:
                                        v2 += v2
                                    case 32767:
                                        if (647 < v10) {
                                            if ((false || ((v10 << ((834) & 63)) != v10)) || (128 == (-(fi21)))) {
                                                switch "hello" {
                                                case "foo bar":
                                                    fmt.Println(fi21)
                                                case "a":
                                                    fmt.Println(v2)
                                                default:
                                                    v27 := GenS5{F0: ((42 == (fi20 ^ fi20)) || true), F1: -31}
                                                    _ = v27
                                                }
                                            }
                                            fmt.Println(v5, v3, v10)
                                        }
                                    default:
                                        fmt.Println(v10, v2, fi20)
                                    }
                                }
                            case "hello":
                                fmt.Println(fi20, fi21)
                            case "":
                                switch ("\n\t" + ("\n\t" + "a")) {
                                case "\n\t":
                                    delete(v1, 1)
                                default:
                                    var v28 int16 = -173
                                    _ = v28
                                }
                            default:
                                _ = len(v4)
                            }
                        }
                    } else {
                        switch ("" + "hello") {
                        case "\n\t":
                            var v29 int32 = (v10 - v10)
                            _ = v29
                        default:
                            switch fi20 {
                            case 128:
                                if ((1.5 - float64(0.0)) != 1.5) {
                                    _ = len(v4)
                                    var v30 uint64 = v5
                                    var v31 int8 = 97
                                    _ = v30
                                    _ = v31
                                } else {
                                    var v32 int8 = 0
                                    v33 := GenS1{F0: (-(fi20))}
                                    _ = v32
                                    _ = v33
                                }
                            }
                        }
                        v3 = v3
                        v3 += v5
                    }
                    _ = fi21
                }
                var v34 int8 = 42
                _ = fi20
                _ = v34
            }
            _ = v19
        }
        for fi35 := 0; fi35 < 11; fi35++ {
            v36 := GenS5{F0: false, F1: -32768}
            if false {
                v37 := GenS6{F0: 245, F1: (v3 | 2)}
                switch (^(fi35)) {
                case -894:
                    if (int32(fi35) > v10) {
                        v2 += (-(v2))
                    } else {
                        switch ("a" + "hello") {
                        case "":
                            for fi38 := 0; fi38 < 12; fi38++ {
                                v3 = (v5 << ((211) & 63))
                                if ((!(v2 >= (v2 >> ((0) & 63)))) && (v5 == v3)) {
                                    var v39 uint8 = 1
                                    _ = v39
                                }
                                _ = fi38
                            }
                        default:
                            var v40 int8 = 42
                            _ = v40
                        }
                        fmt.Println(v2, v3)
                    }
                default:
                    fmt.Println(v5, v2, v3)
                }
                var v41 uint32 = 2147483647
                _ = v37
                _ = v41
            } else {
                var v42 int8 = 0
                _ = v42
            }
            var v43 float32 = (1e10 - 1e10)
            _ = fi35
            _ = v36
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        switch 127 {
        case 128:
            fmt.Println(v10, v2)
        case 2147483647:
            var v45 string = ("foo bar" + "foo bar")
            _ = v45
        default:
            var v46 int64 = 0
            _ = v46
        }
        _ = fi44
    }
    var v47 string = ""
    if true {
        v5 += (v5 / (((^(v3))) | 1))
    }
    v48 := GenS2{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
