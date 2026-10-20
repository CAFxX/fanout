package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 bool; F2 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := []uint32{530}
    var v2 int8 = 127
    v3 := []int{}
    v2 -= v2
    for fi4 := 0; fi4 < 4; fi4++ {
        if (false && false) {
            switch (fi4 ^ (fi4 ^ 2147483647)) {
            case 128:
                if (true || (2147483647 <= 1)) {
                    var v5 uint = 0
                    _ = v5
                }
            case -2147483648:
                if false {
                    if (true && true) {
                        if false {
                            if (0 > 2147483647) {
                                v2 -= v2
                                if true {
                                    v2 = v2
                                    v6 := GenS1{F0: 32767, F1: 805, F2: -744}
                                    fmt.Println(v2, fi4)
                                    _ = v6
                                }
                                for fi7 := 0; fi7 < 13; fi7++ {
                                    if false {
                                        v8 := GenS1{F0: 1, F1: 4294967295, F2: -1}
                                        if (!true) {
                                            var v9 bool = (65535 >= 1)
                                            _ = v9
                                        } else {
                                            if (-189 >= 2147483647) {
                                                v2 = (v2 & v2)
                                                v3 = append(v3, 42)
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            fmt.Println(fi4)
                                        }
                                        _ = v8
                                    }
                                    var v10 uint = 4294967295
                                    _ = fi7
                                    _ = v10
                                }
                            }
                            var v11 int32 = -1
                            _ = v11
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v12 := GenS2{F0: ((v2 >> ((18446744073709551615) & 63)) >= v2)}
                    _ = v12
                } else {
                    var v13 string = "hello"
                    fmt.Println(fi4, fi4)
                    _ = v13
                }
            case 32767:
                fmt.Println(fi4, v2)
            }
        } else {
            for fi14 := 0; fi14 < 19; fi14++ {
                if true {
                    fmt.Println(fi4, fi4)
                } else {
                    if true {
                        v2 = 100
                        var v15 uint8 = 0
                        var v16 float32 = 0.0
                        _ = v15
                        _ = v16
                    }
                }
                _ = fi14
            }
            switch "foo bar" {
            case "\n\t":
                v17 := GenS3{F0: fi4, F1: true, F2: 2147483648}
                _ = v17
            }
            v2 += (v2 | v2)
        }
        switch fi4 {
        case 9223372036854775807:
            fmt.Println(v2, v2)
        case 42:
            var v18 bool = true
            _ = v18
        case 32768:
            if ((!false) && (127 > 1)) {
                v3 = append(v3, 128)
            }
        default:
            var v19 uint32 = 399
            _ = v19
        }
        var v20 float32 = (0.0 + (0.0 + 3.14159))
        _ = fi4
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 int16 = 0
        fmt.Println(v22, v2)
        _ = fi21
        _ = v22
    }
    fmt.Println(v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v23 := GenS4{F0: v2}
    fmt.Println(v2)
    var v24 int8 = v2
    var v25 int16 = -930
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (false && true) {
        var v26 uint64 = 2
        fmt.Println(v26)
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 int16 = -179
            _ = fi27
            _ = v28
        }
        _ = v26
    } else {
        v25 = v25
    }
    v25 = (^(v25))
    v29 := GenS5{F0: 0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 -= (-(v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
}
