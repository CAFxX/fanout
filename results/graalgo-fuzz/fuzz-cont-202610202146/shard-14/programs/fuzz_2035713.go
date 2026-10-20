package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int8; F1 bool; F2 int32 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 uint; F1 int16; F2 int8 }
func genFn0() bool {
    v4 := []int{-716, -642, -9223372036854775808, -9223372036854775808}
    var v5 int = 499
    _ = v4
    _ = v5
    return ((127 == 255) && false)
}
func genFn1() string {
    var v6 uint = 42
    v6 -= 65535
    var v7 uint8 = 255
    var v8 uint16 = 696
    _ = v6
    _ = v7
    _ = v8
    return ((("foo bar" + ("" + ("\n\t" + "\n\t"))) + "hello") + "a")
}

func main() {
    var v1 uint16 = 42
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]bool)
    for fi9 := 0; fi9 < 18; fi9++ {
        if true {
            if true {
                if (-128 >= 42) {
                    v10 := GenS1{F0: 255}
                    var v11 uint8 = 249
                    var v12 int16 = 32767
                    _ = v10
                    _ = v11
                    _ = v12
                }
                v13 := []int{fi9, fi9}
                if (!((1 >= 0) || (4294967295 != 408))) {
                    fmt.Println(v2, v1, v1)
                    fmt.Println(fi9)
                    for fi14 := 0; fi14 < 18; fi14++ {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi9)
                        case "":
                            switch (^(fi9)) {
                            case -590:
                                if (true || false) {
                                    var v15 uint64 = 42
                                    var v16 float64 = 1e10
                                    v16 += 3.14159
                                    _ = v15
                                    _ = v16
                                }
                            case 128:
                                var v17 string = (("" + "foo bar") + "foo bar")
                                _ = v17
                            default:
                                if true {
                                    fmt.Println(fi14, fi9, v1)
                                    if ((42 >= -9223372036854775808) || (!false)) {
                                        fmt.Println(fi14, fi9)
                                        if ((("a" + "a") + "a") <= "\n\t") {
                                            if true {
                                                var v18 uint32 = 65535
                                                v19 := GenS2{F0: 141, F1: (v2 | (v2 >> ((2) & 63))), F2: uint16(v18)}
                                                var v20 int8 = -128
                                                _ = v18
                                                _ = v19
                                                _ = v20
                                            } else {
                                                if false {
                                                    v1 = v1
                                                    v21 := GenS2{F0: 42, F1: v2, F2: 32767}
                                                    fmt.Println(v2)
                                                    _ = v21
                                                }
                                                fmt.Println(fi9, fi14, v1)
                                                v1 -= (^(v1))
                                                fmt.Println(v1, fi9, fi9)
                                            }
                                            v13 = append(v13, 9223372036854775807)
                                        } else {
                                            var v22 float64 = 1e10
                                            _ = v22
                                        }
                                        v23 := GenS3{F0: 42, F1: true, F2: -907}
                                        _ = v23
                                    }
                                    fmt.Println(v2)
                                } else {
                                    var v24 bool = false
                                    v25 := GenS3{F0: -1, F1: ((v1 * 896) <= (v1 & 65535)), F2: int32(v2)}
                                    _ = v24
                                    _ = v25
                                }
                            }
                        }
                        _ = fi14
                    }
                } else {
                    switch fi9 {
                    case -874:
                        var v26 uint16 = uint16(v2)
                        _ = v26
                    default:
                        var v27 uint32 = 1
                        _ = v27
                    }
                }
                _ = v13
            } else {
                var v28 string = "\n\t"
                _ = v28
            }
        } else {
            v1 = (^(v1))
            v1 = v1
        }
        if false {
            var v29 uint = 908
            _ = v29
        }
        var v30 int16 = 0
        _ = fi9
        _ = v30
    }
    fmt.Println(v1, v2)
    var v31 uint8 = 128
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := GenS4{F0: v2, F1: 0}
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        for fi35 := 0; fi35 < 8; fi35++ {
            if (!((v2 << ((696) & 63)) == 0)) {
                v36 := GenS3{F0: 0, F1: true, F2: 2147483647}
                v1 += 42
                _ = v36
            }
            if true {
                v37 := GenS5{F0: 4294967295, F1: 1}
                var v38 bool = (true && true)
                var v39 int64 = 1
                _ = v37
                _ = v38
                _ = v39
            }
            switch fi35 {
            case -111:
                v31 -= v31
            case -1:
                var v40 int64 = -32
                _ = v40
            }
            _ = fi35
        }
        var v41 int32 = -23
        v42 := GenS6{F0: 326, F1: -1, F2: 42}
        _ = fi34
        _ = v41
        _ = v42
    }
    var v43 float32 = 3.14159
    var v44 uint8 = (v31 | v31)
    var v45 string = ("\n\t" + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v43
    _ = v44
    _ = v45
}
