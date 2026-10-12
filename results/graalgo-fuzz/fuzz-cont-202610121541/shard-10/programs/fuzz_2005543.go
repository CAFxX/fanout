package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 bool; F2 uint }

func main() {
    var v1 int8 = 100
    v2 := []int{0, 895, -2147483648}
    var v3 string = ""
    v4 := []uint32{42, 0, 1, 2}
    _ = len(v4)
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int64 = 9223372036854775807
        _ = fi5
        _ = v6
    }
    fmt.Println(v1)
    v7 := GenS1{F0: 65535, F1: 65535, F2: 52}
    if true {
        v3 = "foo bar"
        v8 := GenS1{F0: 2147483647, F1: 2, F2: 0}
        _ = v8
    } else {
        v3 = v3
        var v9 int64 = 42
        var v10 int = 0
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        switch (fi11 >> ((559) & 63)) {
        case 0:
            fmt.Println(v3)
        }
        if true {
            switch v3 {
            case "a":
                v12 := GenS1{F0: 197, F1: 2, F2: 128}
                _ = v12
            case "\n\t":
                v13 := GenS1{F0: 2, F1: 542, F2: 141}
                _ = v13
            case "":
                v3 = (v3 + ("" + ((v3 + "") + "\n\t")))
            }
            v3 = v3
            v3 = v3
        } else {
            for fi14 := 0; fi14 < 6; fi14++ {
                if (float64(0.0) >= 0.1) {
                    var v15 int64 = 1
                    _ = v15
                }
                v16 := GenS1{F0: 19, F1: 42, F2: 128}
                v3 = (v3 + ("hello" + (v3 + (v3 + "a"))))
                _ = fi14
                _ = v16
            }
        }
        _ = fi11
    }
    v1 += 1
    var v17 int16 = 32767
    fmt.Println(v17)
    if ((true || ((0.1 - (-2.25 + 1e10)) != (float64(((1.5 - 1.5) * 3.14159)) * (-2.25 / 1e10)))) || ((v17 <= 42) || true)) {
        var v18 string = v3
        _ = v18
    } else {
        v19 := GenS1{F0: 42, F1: 42, F2: 128}
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 int = fi20
            v17 = (v17 << ((42) & 63))
            if (false && true) {
                switch (^(fi20)) {
                case 2147483647:
                    var v22 int16 = v17
                    _ = v22
                case 4294967295:
                    var v23 int = (fi20 % ((32768) | 1))
                    _ = v23
                }
                var v24 uint16 = 500
                for fi25 := 0; fi25 < 20; fi25++ {
                    if true {
                        if false {
                            _ = len(v2)
                            var v26 uint16 = (v24 * v24)
                            _ = v26
                        }
                        fmt.Println(v1, fi20)
                        if (!false) {
                            v17 += v17
                        }
                    }
                    fmt.Println(fi20, v17, fi25)
                    _ = fi25
                }
                _ = v24
            } else {
                _ = len(v4)
                for fi27 := 0; fi27 < 13; fi27++ {
                    if false {
                        v21 += fi20
                        switch (v21 | 9223372036854775807) {
                        case 32767:
                            fmt.Println(v3, v1)
                        default:
                            if (false || true) {
                                if ((("a" + v3) + (v3 + v3)) > "foo bar") {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v1 = 0
                                    switch -1 {
                                    case 952:
                                        var v28 bool = false
                                        _ = v28
                                    case 2147483647:
                                        v17 += (v17 ^ (v17 | v17))
                                    }
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = 709
                                    }
                                }
                                if len(v4) > 0 {
                                    v4[0] = 0
                                }
                                switch fi27 {
                                case 127:
                                    fmt.Println(fi20)
                                default:
                                    var v29 string = (v3 + v3)
                                    _ = v29
                                }
                            }
                        }
                    } else {
                        var v30 uint64 = 2
                        _ = v30
                    }
                    v31 := GenS2{F0: 42, F1: 42, F2: 363}
                    _ = fi27
                    _ = v31
                }
                switch "hello" {
                case "\n\t":
                    var v32 int8 = 49
                    _ = v32
                case "":
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                case "hello":
                    v33 := GenS3{F0: 766}
                    _ = v33
                default:
                    v34 := GenS4{F0: ((v3 + (v3 + "a")) + ("\n\t" + v3))}
                    _ = v34
                }
            }
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    v35 := GenS5{F0: 4294967295}
    switch v3 {
    case "foo bar":
        v36 := GenS6{F0: 935, F1: (1 < 2147483647), F2: 2}
        _ = v36
    default:
        fmt.Println(v17)
    }
    fmt.Println(v17, v1, v1)
    v17 += v17
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v7
}
