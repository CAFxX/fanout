package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 string; F2 uint16 }
type GenS7 struct { F0 uint64; F1 bool }

func main() {
    var v1 bool = (278 < 2147483648)
    var v2 int64 = -1
    var v3 int64 = -9223372036854775808
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]int)
    fmt.Println(v3, v3, v1)
    if (v1 && true) {
        var v6 uint64 = 403
        var v7 uint8 = 112
        delete(v5, 1)
        _ = v6
        _ = v7
    } else {
        if (42 == 164) {
            v5[4294967295] = 781
            fmt.Println(v1, v3, v3)
            _ = len(v5)
        }
        switch ((("foo bar" + "a") + "foo bar") + "a") {
        case "\n\t":
            v8 := []int{42, 418, -1, -9223372036854775808}
            _ = v8
        default:
            v9 := GenS1{F0: 4294967295, F1: v2}
            _ = v9
        }
        if (0 < 1) {
            var v10 int8 = -128
            v2 += v3
            switch "foo bar" {
            case "\n\t":
                if (!((3.14159 - -2.25) != 0.1)) {
                    if ((v1 && ((v1 || v1) && v1)) || ((1e10 - (((1.5 * -2.25) / 1e10) / (3.14159 - -2.25))) != (3.14159 * ((-2.25 / (-2.25 / 1.0)) * (0.0 - (1e10 - float32(0.0))))))) {
                        var v11 uint16 = 32767
                        v12 := []int{}
                        v13 := GenS2{F0: (!(1 <= 1))}
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        v1 = ((!(-2147483648 == 1)) || v1)
                        v1 = v1
                        for fi14 := 0; fi14 < 6; fi14++ {
                            if (4294967295 != 342) {
                                for fi15 := 0; fi15 < 16; fi15++ {
                                    v1 = true
                                    v1 = v1
                                    v10 += (v10 | (v10 % (((v10 + (-(v10)))) | 1)))
                                    if v1 {
                                        var v16 uint8 = 98
                                        v2 += 758
                                        switch (fi15 + 128) {
                                        case -2147483648:
                                            v10 = (v10 - 1)
                                        case -777:
                                            var v17 uint64 = 578
                                            _ = v17
                                        case 42:
                                            if v1 {
                                                v1 = (v1 && (v2 <= v3))
                                                v3 = (v2 - (v3 >> ((4294967295) & 63)))
                                                fmt.Println(v16, v3, v3)
                                            } else {
                                                v18 := v5[4294967295]
                                                _ = v18
                                            }
                                        }
                                        fmt.Println(v2, v3, fi15)
                                        _ = v16
                                    } else {
                                        fmt.Println(fi14)
                                    }
                                    fmt.Println(v1, v2, fi15)
                                    v19 := GenS3{F0: -919, F1: v2}
                                    fmt.Println(fi15, fi15)
                                    _ = fi15
                                    _ = v19
                                }
                                var v20 uint16 = 65535
                                v21 := []int{int(v20), -1, fi14}
                                _ = v20
                                _ = v21
                            } else {
                                v2 = (-(v3))
                            }
                            if ((true && v1) || v1) {
                                var v22 string = ("" + "")
                                if v1 {
                                    v1 = (4294967295 <= 42)
                                    if (274 < 4294967295) {
                                        v10 = (-(v10))
                                        v22 = ("\n\t" + "\n\t")
                                        fmt.Println(fi14, v1, v10)
                                    }
                                    for fi23 := 0; fi23 < 8; fi23++ {
                                        v24 := GenS1{F0: 4294967295, F1: -9223372036854775808}
                                        if ((2 == 42) && v1) {
                                            if v1 {
                                                var v25 float64 = -2.25
                                                _ = v25
                                            }
                                            if ((834 >= 42) || (!v1)) {
                                                switch "hello" {
                                                case "\n\t":
                                                    v26 := []int{(^(fi14))}
                                                    _ = v26
                                                case "":
                                                    v27 := GenS4{F0: v22, F1: 413}
                                                    _ = v27
                                                }
                                                if v1 {
                                                    if (true || true) {
                                                        v10 = (-(v10))
                                                        v22 = (v22 + v22)
                                                        var v28 int8 = (v10 + v10)
                                                        _ = v28
                                                    }
                                                }
                                            }
                                            var v29 int = fi14
                                            v30 := GenS3{F0: -1, F1: (v2 % ((v3) | 1))}
                                            if false {
                                                switch -2147483648 {
                                                case 127:
                                                    fmt.Println(fi14, v10, v10)
                                                default:
                                                    v3 = -9223372036854775808
                                                }
                                                v31 := GenS2{F0: false}
                                                v32 := GenS5{F0: v3}
                                                _ = v31
                                                _ = v32
                                            }
                                            _ = v29
                                            _ = v30
                                        }
                                        _ = fi23
                                        _ = v24
                                    }
                                }
                                _ = v22
                            } else {
                                v33 := GenS3{F0: 1, F1: 42}
                                v3 = v2
                                _ = v33
                            }
                            v10 = (v10 - v10)
                            _ = fi14
                        }
                    }
                    fmt.Println(v3, v10, v2)
                    if (16 >= 170) {
                        var v34 float32 = 1.5
                        for fi35 := 0; fi35 < 2; fi35++ {
                            v34 -= (float32(0.1) - ((0.0 * v34) / 1e10))
                            v10 += 127
                            _ = fi35
                        }
                        _ = v34
                    } else {
                        var v36 int64 = (v2 << ((65535) & 63))
                        _ = len(v5)
                        _ = v36
                    }
                }
            case "":
                var v37 int = -917
                _ = v37
            }
            _ = v10
        }
    }
    switch "hello" {
    case "hello":
        for fi38 := 0; fi38 < 8; fi38++ {
            v3 = 43
            var v39 uint32 = 2147483648
            _ = fi38
            _ = v39
        }
    }
    var v40 uint64 = 345
    v1 = (0 == 250)
    if (!(42 >= -42)) {
        var v41 uint16 = 255
        if false {
            v42 := GenS6{F0: 214, F1: (("" + "") + ("hello" + "\n\t")), F2: 808}
            _ = v42
        } else {
            v43 := GenS2{F0: (v1 && v1)}
            var v44 float32 = ((3.14159 - 0.1) + 0.1)
            fmt.Println(v1, v2, v1)
            _ = v43
            _ = v44
        }
        fmt.Println(v1, v41, v1)
        _ = v41
    }
    v45 := GenS4{F0: "foo bar", F1: 0}
    v3 = (v2 << ((65535) & 63))
    if true {
        if v1 {
            v46 := GenS7{F0: 0, F1: v1}
            var v47 bool = (12 >= -94)
            if true {
                for fi48 := 0; fi48 < 12; fi48++ {
                    var v49 string = "hello"
                    _ = fi48
                    _ = v49
                }
            } else {
                var v50 uint64 = v40
                _ = v50
            }
            _ = v46
            _ = v47
        }
        v2 += v2
        v1 = v1
    } else {
        var v51 uint64 = 42
        _ = v51
    }
    var v52 int8 = 1
    for fi53 := 0; fi53 < 20; fi53++ {
        if v1 {
            v52 -= (v52 / ((v52) | 1))
            var v54 float32 = float32(0.1)
            v3 += (v2 - v3)
            _ = v54
        }
        var v55 int16 = 1000
        var v56 uint = 326
        _ = fi53
        _ = v55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v45
    _ = v5
    _ = v52
}
