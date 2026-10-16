package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int; F1 int64; F2 int32 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int64; F1 int32; F2 string }

func main() {
    v1 := []int{-658, 322}
    v2 := []int{}
    var v3 uint8 = 61
    v4 := []uint32{596, 283, 2147483648}
    v5 := []uint32{2147483648, 2147483647, 42, 4294967295}
    v6 := []uint32{}
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: false}
        for fi9 := 0; fi9 < 2; fi9++ {
            if len(v5) > 0 {
                v5[0] = 2
            }
            var v10 int64 = -395
            var v11 uint = 4294967295
            if true {
                switch ("\n\t" + "a") {
                case "":
                    v10 -= 42
                    var v12 uint = (v11 - (^(v11)))
                    v3 += v3
                    _ = v12
                default:
                    _ = len(v1)
                }
                if true {
                    var v13 int32 = -1
                    var v14 uint = 827
                    _ = v13
                    _ = v14
                } else {
                    v11 = v11
                    v15 := GenS2{F0: (v11 >> ((v11) & 63)), F1: 127}
                    _ = v15
                }
            }
            var v16 float32 = 3.14159
            switch -9223372036854775808 {
            case 42:
                fmt.Println(v10, v10, v11)
            default:
                v3 = (v3 / ((1) | 1))
            }
            _ = fi9
            _ = v10
            _ = v11
            _ = v16
        }
        v17 := GenS3{F0: fi7, F1: -9223372036854775808, F2: 716}
        _ = fi7
        _ = v17
        _ = v8
    }
    var v18 string = "foo bar"
    var v19 int64 = 0
    var v20 int = 9223372036854775807
    fmt.Println(v19, v20, v19)
    switch v20 {
    case 42:
        v2 = append(v2, (v20 >> ((1) & 63)))
    case 19:
        v21 := GenS2{F0: 682, F1: -128}
        _ = v21
    default:
        v22 := GenS2{F0: 4294967295, F1: 1}
        _ = v22
    }
    if (646 <= uint(v3)) {
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS1{F0: (42 >= 42)}
            _ = fi23
            _ = v24
        }
        v19 = -300
        v3 -= v3
    }
    var v25 int = (v20 % ((v20) | 1))
    var v26 int8 = -58
    fmt.Println(v3, v18, v26)
    v3 = v3
    fmt.Println(v3, v19)
    if (0 != -175) {
        v26 -= (v26 | 127)
        for fi27 := 0; fi27 < 7; fi27++ {
            fmt.Println(v3, v26, v26)
            _ = fi27
        }
        if true {
            fmt.Println(v25)
        } else {
            fmt.Println(v20, v19, v19)
            if ((v20 >= 9223372036854775807) || ((v19 != -9223372036854775808) && (!((1.5 != 1e10) && (!true))))) {
                v28 := GenS2{F0: 243, F1: -128}
                v19 += 0
                _ = v28
            } else {
                if true {
                    if false {
                        v29 := GenS1{F0: true}
                        _ = v29
                    } else {
                        var v30 bool = true
                        switch v18 {
                        case "\n\t":
                            if len(v1) > 0 {
                                v1[0] = v20
                            }
                        case "":
                            v19 -= v19
                        }
                        v25 += v25
                        _ = v30
                    }
                } else {
                    fmt.Println(v3, v18)
                    fmt.Println(v18, v26)
                    for fi31 := 0; fi31 < 15; fi31++ {
                        var v32 uint64 = 482
                        var v33 int16 = 1000
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                }
                fmt.Println(v3)
                if true {
                    fmt.Println(v3, v20)
                    fmt.Println(v3)
                } else {
                    v34 := GenS3{F0: 128, F1: v19, F2: 1}
                    v35 := GenS4{F0: 0, F1: -2147483648}
                    _ = v34
                    _ = v35
                }
            }
        }
    } else {
        v26 += v26
        if (!(!true)) {
            switch ((v18 + "") + ((v18 + v18) + "foo bar")) {
            case "hello":
                if len(v6) > 0 {
                    v6[0] = 4294967295
                }
            default:
                v1 = append(v1, v20)
            }
            v20 = (-(v20))
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        if true {
            var v36 int = v25
            v37 := GenS5{F0: (v19 >> ((18446744073709551615) & 63)), F1: 1, F2: v18}
            for fi38 := 0; fi38 < 2; fi38++ {
                v18 = (v18 + "")
                _ = fi38
            }
            _ = v36
            _ = v37
        }
    }
    var v39 uint64 = 42
    var v40 int16 = 326
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        fmt.Println(v19)
        var v41 string = v18
        _ = v41
    } else {
        for fi42 := 0; fi42 < 13; fi42++ {
            v26 = (v26 | v26)
            _ = fi42
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
