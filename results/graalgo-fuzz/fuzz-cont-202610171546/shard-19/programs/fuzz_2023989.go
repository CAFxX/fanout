package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 uint64, p1 uint8) uint {
    var v7 bool = true
    v8 := []int{-527, -394}
    if (-128 > 1) {
        v8 = append(v8, 975)
        v7 = ((!true) || ("" < "\n\t"))
    } else {
        fmt.Println(p1)
        var v9 uint8 = 42
        v8 = append(v8, 42)
        _ = v9
    }
    if (4294967295 < p0) {
        _ = len(v8)
        var v10 bool = v7
        _ = v10
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 uint = 123
    var v2 uint32 = 753
    var v3 uint = (v1 % (((v1 >> ((v1) & 63))) | 1))
    var v4 float32 = 1e10
    var v5 int8 = 42
    var v6 int8 = (v5 << ((v3) & 63))
    v11 := []int{32768, 128, 128}
    var v12 int16 = 799
    switch -829 {
    case -9223372036854775808:
        var v13 int8 = 31
        _ = v13
    case 0:
        if (false && false) {
            fmt.Println(v2, v3)
        } else {
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 uint = 907
                if ((!true) && true) {
                    var v16 string = ""
                    v17 := GenS1{F0: (!false), F1: v15, F2: 4294967295}
                    _ = v16
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
        }
    }
    if true {
        v12 -= (v12 - v12)
        fmt.Println(v1)
    } else {
        v11 = append(v11, 1)
    }
    if (false && false) {
        v18 := GenS1{F0: (137 != 229), F1: (v1 << ((v3) & 63)), F2: 926}
        switch 128 {
        case 9223372036854775807:
            v4 = v4
        default:
            var v19 int16 = (v12 << ((genFn0(uint64(v2), 128)) & 63))
            _ = v19
        }
        _ = v18
    }
    _ = len(v11)
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 uint64 = uint64(v2)
            v23 := GenS2{F0: -9223372036854775808}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    v4 = (v4 / 1.0)
    for fi24 := 0; fi24 < 7; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            if true {
                var v26 uint32 = (^(v2))
                if ((true || ((v6 & v5) < v5)) || (!true)) {
                    switch fi25 {
                    case -538:
                        if false {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            if false {
                                var v27 float64 = 0.0
                                fmt.Println(v6, v2, v6)
                                _ = v27
                            } else {
                                switch "hello" {
                                case "hello":
                                    v28 := GenS2{F0: 1}
                                    _ = v28
                                default:
                                    v11 = append(v11, (fi24 - fi25))
                                }
                                var v29 uint32 = (v2 << (((v1 + 65535)) & 63))
                                var v30 int = 32768
                                _ = v29
                                _ = v30
                            }
                        }
                    case 624:
                        v2 = (v2 ^ (^(v2)))
                    case -794:
                        v31 := GenS1{F0: false, F1: (v3 ^ (v1 % (((v3 / ((v3) | 1))) | 1))), F2: 2147483648}
                        _ = v31
                    }
                }
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                fmt.Println(fi25, v4)
                _ = v26
            } else {
                if len(v11) > 0 {
                    v11[0] = fi24
                }
            }
            _ = fi25
        }
        v6 += 42
        v32 := GenS2{F0: int64(v12)}
        _ = fi24
        _ = v32
    }
    var v33 float32 = (v4 / v4)
    var v34 string = ("" + "foo bar")
    for fi35 := 0; fi35 < 11; fi35++ {
        switch ((v34 + "\n\t") + ("foo bar" + (("\n\t" + v34) + v34))) {
        case "\n\t":
            if false {
                v36 := GenS3{F0: 0, F1: v12}
                _ = v36
            }
        case "hello":
            v4 = 1.5
        }
        _ = fi35
    }
    var v37 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
