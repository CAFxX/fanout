package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 int64; F1 int16; F2 uint8 }
type GenS8 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "a"
    v3 := []string{(v2 + "a"), "", (v2 + ""), v2}
    v4 := make(map[uint32]bool)
    var v5 float32 = 0.1
    var v6 int64 = 1
    if true {
        fmt.Println(v5)
        v5 -= (v5 + v5)
    } else {
        for fi7 := 0; fi7 < 15; fi7++ {
            if (false && true) {
                var v8 uint8 = 42
                v5 -= v5
                var v9 float32 = v5
                _ = v8
                _ = v9
            }
            if (false || (false || (78 <= 2))) {
                var v10 float64 = (float64(-2.25) - 3.14159)
                if true {
                    var v11 uint8 = 1
                    v12 := GenS1{F0: 1, F1: 1}
                    _ = v11
                    _ = v12
                }
                v10 = v10
                _ = v10
            } else {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v14 := GenS1{F0: 2, F1: 329}
                    fmt.Println(fi7, v2, fi7)
                    var v15 int8 = 127
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 int64 = (^(v6))
                _ = v16
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi7
        }
        switch v2 {
        case "foo bar":
            v1[65535] = 128
        case "":
            var v17 uint16 = 552
            _ = v17
        case "\n\t":
            _ = len(v1)
        }
        v2 = v2
    }
    fmt.Println(v6, v5, v5)
    if false {
        fmt.Println(v6, v6)
    } else {
        v5 = (1e10 * 1.5)
        if (true && (2 != 2)) {
            fmt.Println(v6)
        }
        if len(v3) > 0 {
            v3[0] = v2
        }
    }
    var v18 int16 = 32767
    if (!(!(v2 != v2))) {
        if ((v18 < v18) || (42 > 0)) {
            var v19 string = "hello"
            for fi20 := 0; fi20 < 1; fi20++ {
                fmt.Println(v5, v6, v19)
                for fi21 := 0; fi21 < 13; fi21++ {
                    var v22 int16 = v18
                    v6 = v6
                    var v23 uint = 42
                    if false {
                        v6 = (v6 & (v6 ^ v6))
                        var v24 uint = v23
                        _ = len(v4)
                        _ = v24
                    } else {
                        v19 = v19
                        var v25 uint = (v23 + v23)
                        var v26 int = (fi21 << ((v25) & 63))
                        var v27 float64 = 1.5
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    switch -496 {
                    case 42:
                        if true {
                            v28 := v1[2147483647]
                            var v29 string = "hello"
                            v30 := GenS2{F0: false}
                            if (!(1 <= 1)) {
                                v31 := GenS2{F0: false}
                                _ = v31
                            }
                            v19 = v29
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    case 0:
                        var v32 uint = (^(v23))
                        _ = v32
                    case 347:
                        fmt.Println(v5)
                    default:
                        if (((((v19 + "hello") + ("" + (v2 + (v2 + "\n\t")))) != v2) || (!("hello" == "a"))) || false) {
                            fmt.Println(v23)
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                        }
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                if false {
                    v33 := GenS1{F0: 722, F1: 0}
                    if false {
                        var v34 int32 = 42
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                        _ = v34
                    }
                    var v35 float32 = v5
                    _ = v33
                    _ = v35
                } else {
                    if true {
                        fmt.Println(fi20, v5)
                        fmt.Println(v6, v2)
                    }
                    v19 = v2
                }
                _ = fi20
            }
            _ = v19
        } else {
            v36 := GenS3{F0: 2}
            _ = v36
        }
        v18 = v18
        if (v5 >= 1e10) {
            v18 = v18
            _ = len(v3)
        }
    } else {
        for fi37 := 0; fi37 < 11; fi37++ {
            v38 := GenS4{F0: (false && (100 < 1)), F1: 123}
            if false {
                v5 -= (v5 / v5)
                v39 := GenS5{F0: v18}
                if (false && true) {
                    v40 := GenS3{F0: 42}
                    _ = v40
                }
                _ = v39
            }
            _ = fi37
            _ = v38
        }
        switch 0 {
        case 127:
            switch -1 {
            case -2147483648:
                fmt.Println(v5)
            case 1:
                var v41 uint16 = 2
                _ = v41
            default:
                fmt.Println(v18, v6)
            }
        case 440:
            v42 := GenS6{F0: 2147483647, F1: 245}
            _ = v42
        }
        fmt.Println(v6, v2)
    }
    var v43 float64 = float64((float32(0.0) * v5))
    var v44 uint64 = 2
    v45 := GenS7{F0: v6, F1: 0, F2: 178}
    v43 = ((v43 + 3.14159) - v43)
    v5 = ((v5 / float32((v43 / v43))) * float32(v43))
    switch "" {
    case "\n\t":
        var v46 int32 = -1
        _ = v46
    case "hello":
        v47 := GenS7{F0: (v6 << ((42) & 63)), F1: v18, F2: uint8(v44)}
        _ = v47
    }
    v6 = (v6 + (v6 % ((v6) | 1)))
    var v48 int16 = 1
    v49 := GenS8{F0: -647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
