package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int64; F2 uint8 }
type GenS5 struct { F0 int64 }
func genFn0() int64 {
    if (((!(255 > 255)) && (!true)) && false) {
        v7 := []int{4294967295, -982, 799, 349}
        var v8 float32 = -2.25
        _ = v7
        _ = v8
    } else {
        var v9 int16 = -1
        var v10 int = -9223372036854775808
        v11 := GenS1{F0: 255, F1: false}
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := GenS1{F0: 600, F1: false}
    if (323 >= 255) {
        switch 42 {
        case 723:
            var v13 float32 = 0.0
            _ = v13
        case -1:
            if (false && false) {
                var v14 float64 = 1e10
                _ = v14
            } else {
                v15 := []int{0}
                var v16 int16 = -32768
                _ = v15
                _ = v16
            }
        }
        var v17 int16 = -597
        if (2147483648 > 561) {
            var v18 float32 = 3.14159
            for fi19 := 0; fi19 < 8; fi19++ {
                if false {
                    v20 := []int{(fi19 | (fi19 / ((fi19) | 1))), (fi19 / (((fi19 ^ 127)) | 1)), 128}
                    fmt.Println(v18, fi19, v18)
                    if false {
                        if false {
                            for fi21 := 0; fi21 < 13; fi21++ {
                                v17 += -637
                                var v22 uint = 660
                                var v23 uint8 = 36
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        } else {
                            v24 := GenS2{F0: 0}
                            _ = v24
                        }
                    }
                    _ = v20
                }
                _ = fi19
            }
            _ = v18
        } else {
            v25 := GenS3{F0: "hello"}
            _ = v25
        }
        _ = v17
    } else {
        v26 := []int{-1, -9223372036854775808}
        var v27 string = "\n\t"
        _ = v26
        _ = v27
    }
    switch 127 {
    case 60:
        var v28 int = -2147483648
        _ = v28
    case -2147483648:
        var v29 uint64 = 4294967295
        _ = v29
    case -690:
        if false {
            var v30 uint64 = 317
            if false {
                var v31 int8 = 68
                if false {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v30, v30)
                    case "hello":
                        v31 = v31
                    default:
                        v32 := []int{}
                        _ = v32
                    }
                    var v33 int16 = 183
                    _ = v33
                } else {
                    v34 := GenS3{F0: "\n\t"}
                    _ = v34
                }
                _ = v31
            } else {
                switch ("" + "hello") {
                case "":
                    for fi35 := 0; fi35 < 19; fi35++ {
                        v36 := GenS3{F0: ((("foo bar" + "") + ("foo bar" + "foo bar")) + ("hello" + "\n\t"))}
                        _ = fi35
                        _ = v36
                    }
                case "a":
                    v37 := GenS1{F0: 211, F1: (true || (true || (!false)))}
                    _ = v37
                case "hello":
                    var v38 bool = false
                    _ = v38
                }
                v30 = (^(v30))
            }
            _ = v30
        }
    }
    _ = v12
    return 1
}
func genFn1() uint32 {
    v39 := GenS4{F0: false, F1: 1, F2: 74}
    var v40 float64 = 3.14159
    v40 = (v40 / (float64(((float32((v40 * v40)) * 1e10) - 0.1)) * float64(0.1)))
    var v41 uint32 = 2147483648
    switch "" {
    case "\n\t":
        var v42 uint8 = 128
        _ = v42
    }
    _ = v39
    _ = v40
    _ = v41
    return (^(v41))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = -2.25
    v3 := make(map[uint32]uint32)
    var v4 float64 = 0.1
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]bool)
    if false {
        var v43 uint64 = 255
        _ = v43
    }
    var v44 uint = 7
    var v45 uint8 = 128
    v45 -= (v45 * v45)
    fmt.Println(v44, v44)
    fmt.Println(v45)
    v4 = 1e10
    switch ("foo bar" + "hello") {
    case "":
        var v46 bool = (genFn1() == 4294967295)
        _ = v46
    default:
        v47 := GenS2{F0: (v44 | (v44 | (v44 % (((v44 ^ uint(v45))) | 1))))}
        _ = v47
    }
    fmt.Println(v45, v44, v45)
    v45 -= (^(v45))
    v2 = 0.1
    v48 := v6[4294967295]
    if (228 < 1) {
        v49 := GenS5{F0: 0}
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v48
    _ = v5
    _ = v6
}
