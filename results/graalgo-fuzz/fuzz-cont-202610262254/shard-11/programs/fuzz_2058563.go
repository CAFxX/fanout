package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int; F1 int64; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint64 }
func genFn0(p0 uint16) int32 {
    p0 -= p0
    v5 := []int{-138, -1, -320, -687}
    if false {
        for fi6 := 0; fi6 < 8; fi6++ {
            if (3.14159 != float32((3.14159 - ((float64(1.5) * ((3.14159 - 3.14159) / 1e10)) * 0.0)))) {
                if true {
                    if (304 >= 0) {
                        fmt.Println(p0, fi6)
                        v7 := GenS1{F0: fi6, F1: 410, F2: 4294967295}
                        _ = v7
                    } else {
                        var v8 uint16 = 382
                        _ = v8
                    }
                    for fi9 := 0; fi9 < 12; fi9++ {
                        if true {
                            var v10 int32 = 1
                            p0 -= (^(p0))
                            v11 := GenS2{F0: 42, F1: false}
                            p0 += p0
                            _ = v10
                            _ = v11
                        } else {
                            p0 = 255
                            if false {
                                var v12 int64 = 951
                                p0 -= 65535
                                var v13 int32 = 459
                                _ = v12
                                _ = v13
                            }
                            _ = len(v5)
                        }
                        _ = fi9
                    }
                    var v14 int = (fi6 | fi6)
                    _ = v14
                }
            }
            for fi15 := 0; fi15 < 13; fi15++ {
                p0 += p0
                if true {
                    var v16 int = (fi15 / ((32768) | 1))
                    fmt.Println(fi6, fi15, fi6)
                    var v17 int64 = -9223372036854775808
                    _ = v16
                    _ = v17
                } else {
                    _ = len(v5)
                    if true {
                        p0 += (p0 >> ((0) & 63))
                        p0 += 32767
                    } else {
                        fmt.Println(fi15, fi15, p0)
                        var v18 uint32 = 2147483648
                        if false {
                            fmt.Println(p0, fi15)
                            if (1 <= 2147483647) {
                                v18 -= (v18 % ((v18) | 1))
                                p0 -= (p0 + p0)
                                v18 = v18
                            }
                        } else {
                            p0 = (p0 & p0)
                            fmt.Println(fi15, v18)
                            var v19 int16 = 818
                            _ = v19
                        }
                        _ = v18
                    }
                }
                p0 = (p0 << ((1) & 63))
                _ = fi15
            }
            if false {
                switch "\n\t" {
                case "foo bar":
                    v5 = append(v5, fi6)
                case "hello":
                    for fi20 := 0; fi20 < 1; fi20++ {
                        var v21 int8 = 103
                        v21 = v21
                        switch (^(fi20)) {
                        case -38:
                            var v22 uint32 = 65535
                            _ = v22
                        }
                        var v23 int8 = 91
                        var v24 int32 = 716
                        _ = fi20
                        _ = v21
                        _ = v23
                        _ = v24
                    }
                }
                p0 = p0
                v25 := GenS3{F0: (fi6 & (fi6 + fi6)), F1: -1, F2: 2147483647}
                _ = v25
            }
            _ = fi6
        }
        p0 += 65535
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    _ = p0
    _ = v5
    return 776
}

func main() {
    var v1 float32 = 0.1
    var v2 bool = ((!false) || false)
    var v3 int = -814
    v4 := []int{}
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := GenS3{F0: 2147483647, F1: 42, F2: 0}
        fmt.Println(v1, v3, v2)
        _ = fi26
        _ = v27
    }
    v3 = (v3 - (v3 / ((v3) | 1)))
    v28 := GenS4{F0: genFn0(583), F1: 65535, F2: 610}
    var v29 bool = false
    v1 -= 1.5
    var v30 bool = v2
    v3 += v3
    var v31 uint = 255
    switch "foo bar" {
    case "\n\t":
        for fi32 := 0; fi32 < 9; fi32++ {
            if ((9223372036854775807 != 9223372036854775807) || v2) {
                v31 = 42
            }
            var v33 string = ""
            _ = fi32
            _ = v33
        }
    case "a":
        fmt.Println(v1)
    }
    var v34 int8 = 100
    var v35 bool = v29
    for fi36 := 0; fi36 < 4; fi36++ {
        if ((!(-1 != 588)) || true) {
            var v37 uint8 = 0
            _ = v37
        } else {
            var v38 string = ("" + "")
            _ = v38
        }
        _ = fi36
    }
    fmt.Println(v29, v29, v34)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
}
