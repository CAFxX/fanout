package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint; F1 uint32; F2 int32 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 string; F2 int }
type GenS7 struct { F0 int; F1 uint16 }
func genFn0() int8 {
    if false {
        switch 513 {
        case -9223372036854775808:
            if false {
                var v4 int32 = 2147483647
                v5 := GenS1{F0: (v4 ^ (v4 * 723)), F1: 1}
                var v6 uint16 = 42
                _ = v4
                _ = v5
                _ = v6
            }
        case 42:
            var v7 uint8 = 18
            _ = v7
        case -364:
            var v8 string = ("a" + ("a" + ("hello" + "")))
            _ = v8
        }
        v9 := GenS1{F0: 575, F1: 454}
        if (0 <= 552) {
            v10 := GenS1{F0: 0, F1: 42}
            if (true && true) {
                var v11 uint32 = 2147483648
                if (((1.5 / 1.5) - ((1.5 * -2.25) / float32((0.1 / 1.0)))) >= 3.14159) {
                    if false {
                        v11 -= 308
                        v12 := GenS2{F0: 65535, F1: 955, F2: 2147483647}
                        switch 485 {
                        case 32768:
                            if (-1 != -1) {
                                switch 808 {
                                case -1:
                                    v13 := []int{-9223372036854775808, 786}
                                    _ = v13
                                case 120:
                                    v11 += (^(v11))
                                case 0:
                                    var v14 uint32 = v11
                                    _ = v14
                                }
                                var v15 uint = 18446744073709551615
                                _ = v15
                            }
                        case 4294967295:
                            v16 := []int{-456}
                            _ = v16
                        default:
                            fmt.Println(v11)
                        }
                        _ = v12
                    } else {
                        v17 := GenS2{F0: 42, F1: (v11 << ((42) & 63)), F2: 0}
                        _ = v17
                    }
                }
                _ = v11
            } else {
                var v18 int16 = 623
                _ = v18
            }
            var v19 int32 = 1
            _ = v10
            _ = v19
        } else {
            var v20 int32 = -1
            v21 := GenS1{F0: v20, F1: -1}
            for fi22 := 0; fi22 < 20; fi22++ {
                fmt.Println(fi22, v20)
                _ = fi22
            }
            _ = v20
            _ = v21
        }
        _ = v9
    } else {
        var v23 float32 = 0.1
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 int8 = -1
            v23 += v23
            v25 = (v25 % ((-128) | 1))
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    v26 := GenS3{F0: ((677 <= 9223372036854775807) || false), F1: -32768}
    v27 := GenS3{F0: true, F1: 0}
    for fi28 := 0; fi28 < 15; fi28++ {
        for fi29 := 0; fi29 < 8; fi29++ {
            if false {
                var v30 int32 = 304
                fmt.Println(fi29, fi28, fi29)
                var v31 uint64 = 42
                var v32 int32 = 1
                if true {
                    fmt.Println(fi29, v31, fi29)
                } else {
                    v33 := GenS4{F0: false}
                    var v34 float64 = -2.25
                    v35 := GenS3{F0: (!(1e10 <= v34)), F1: -32768}
                    _ = v33
                    _ = v34
                    _ = v35
                }
                _ = v30
                _ = v31
                _ = v32
            }
            fmt.Println(fi28, fi28)
            var v36 int16 = 1
            fmt.Println(fi28, fi28)
            if false {
                switch ("foo bar" + ("" + "hello")) {
                case "hello":
                    v37 := GenS3{F0: false, F1: (v36 + (v36 ^ v36))}
                    _ = v37
                case "foo bar":
                    var v38 uint8 = 42
                    _ = v38
                default:
                    var v39 int8 = 0
                    _ = v39
                }
                var v40 int64 = 821
                _ = v40
            }
            v41 := GenS5{F0: 1}
            v36 -= -32768
            _ = fi29
            _ = v36
            _ = v41
        }
        _ = fi28
    }
    _ = v26
    _ = v27
    return -122
}
func genFn1(p0 uint32, p1 int32) uint {
    fmt.Println(p1, p0)
    var v42 float32 = (float32(-2.25) + 3.14159)
    _ = p0
    _ = p1
    _ = v42
    return 1
}

func main() {
    var v1 float64 = 1.5
    var v2 string = ("hello" + "hello")
    var v3 float64 = 3.14159
    var v43 uint8 = 27
    var v44 uint32 = 491
    v45 := []int{}
    _ = len(v45)
    v46 := GenS1{F0: 42, F1: -265}
    var v47 float32 = float32(float64(1.5))
    for fi48 := 0; fi48 < 15; fi48++ {
        v49 := GenS6{F0: 1, F1: "", F2: fi48}
        switch (v2 + ((v2 + "a") + v2)) {
        case "\n\t":
            v50 := GenS7{F0: fi48, F1: 1}
            _ = v50
        case "":
            fmt.Println(fi48, v44)
        case "hello":
            var v51 int8 = 42
            _ = v51
        }
        _ = fi48
        _ = v49
    }
    if false {
        v44 -= v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
