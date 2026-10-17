package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int; F2 int8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int32; F1 int; F2 int32 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 uint16; F1 uint64; F2 bool }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (37 <= 149)
    v3 := make(map[int]uint32)
    var v4 bool = (!v2)
    if (!(2147483647 <= 862)) {
        var v5 int16 = 504
        _ = v5
    } else {
        v4 = v2
        fmt.Println(v2)
        for fi6 := 0; fi6 < 2; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := v3[fi6]
                if (v4 && v2) {
                    v9 := []int{(fi7 | fi7), (fi6 & fi7), 4294967295}
                    var v10 int = fi7
                    var v11 float64 = ((0.1 * -2.25) - float64(-2.25))
                    _ = v10
                    _ = v11
                    _ = v9
                } else {
                    v2 = false
                    var v12 int16 = 441
                    _ = v12
                }
                _ = fi7
                _ = v8
            }
            for fi13 := 0; fi13 < 9; fi13++ {
                if v4 {
                    var v14 uint = 255
                    _ = v14
                }
                var v15 uint8 = 1
                v16 := GenS1{F0: v2, F1: 1}
                _ = fi13
                _ = v15
                _ = v16
            }
            _ = fi6
        }
    }
    v1[-2147483648] = false
    var v17 float64 = 0.0
    if v4 {
        v18 := GenS1{F0: (v2 || (656 != 0)), F1: 422}
        _ = v18
    } else {
        if v2 {
            v19 := GenS1{F0: (v2 || true), F1: 453}
            v17 = 1.5
            fmt.Println(v17)
            _ = v19
        } else {
            v20 := GenS2{F0: 2}
            _ = v20
        }
    }
    if ((v4 || (127 == 0)) || v4) {
        fmt.Println(v2, v4)
    }
    v21 := GenS3{F0: 42, F1: 849, F2: 127}
    if (127 == 255) {
        if (!(!v2)) {
            switch 9223372036854775807 {
            case 9223372036854775807:
                var v22 float64 = (float64(3.14159) - (v17 * v17))
                _ = v22
            case 127:
                if (true || (-605 <= -1)) {
                    if false {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            v4 = true
                            _ = fi23
                        }
                        v24 := GenS3{F0: 4294967295, F1: -9223372036854775808, F2: 42}
                        var v25 int8 = -1
                        _ = v24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 14; fi26++ {
                        var v27 string = ((("hello" + "foo bar") + "") + "hello")
                        _ = fi26
                        _ = v27
                    }
                } else {
                    var v28 uint16 = 125
                    for fi29 := 0; fi29 < 8; fi29++ {
                        if v4 {
                            if v2 {
                                if true {
                                    v30 := GenS3{F0: 4294967295, F1: (fi29 ^ 2147483647), F2: -26}
                                    switch (("hello" + (("foo bar" + "hello") + "foo bar")) + "foo bar") {
                                    case "foo bar":
                                        if false {
                                            var v31 uint32 = 0
                                            _ = v31
                                        }
                                    case "a":
                                        v32 := GenS4{F0: 9223372036854775807, F1: 127}
                                        _ = v32
                                    }
                                    v33 := []int{fi29, fi29, (fi29 % ((42) | 1)), (fi29 / (((^(fi29))) | 1))}
                                    _ = v30
                                    _ = v33
                                } else {
                                    v34 := GenS5{F0: -1, F1: (fi29 << ((4294967295) & 63)), F2: -1}
                                    _ = v34
                                }
                                v35 := GenS6{F0: 2147483647, F1: 42}
                                _ = v35
                            } else {
                                v36 := []int{fi29, fi29, fi29}
                                if false {
                                    var v37 int64 = -1
                                    _ = v37
                                }
                                v28 = (v28 / ((v28) | 1))
                                _ = v36
                            }
                            v38 := GenS3{F0: 1, F1: (fi29 >> ((816) & 63)), F2: 100}
                            _ = v38
                        }
                        _ = fi29
                    }
                    v39 := GenS7{F0: 42, F1: 563, F2: true}
                    _ = v28
                    _ = v39
                }
            case -853:
                var v40 uint8 = 104
                _ = v40
            }
        }
    }
    fmt.Println(v17, v2, v2)
    var v41 uint = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v41
}
