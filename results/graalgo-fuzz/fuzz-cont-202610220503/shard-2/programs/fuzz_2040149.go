package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint64; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 uint16 }

func main() {
    var v1 uint8 = 128
    var v2 int32 = 2147483647
    var v3 uint8 = 1
    var v4 float64 = (float64(3.14159) / 1.5)
    v5 := []int{31, -171, 889, 4294967295}
    fmt.Println(v2)
    var v6 float32 = float32((float64(0.0) / float64(((0.1 * float32(1e10)) / (0.1 + (((0.1 / -2.25) / 3.14159) / 3.14159))))))
    fmt.Println(v3, v1, v2)
    if true {
        v6 -= ((v6 / float32(v4)) - ((v6 + v6) / (float32(float64(1.5)) / -2.25)))
    }
    fmt.Println(v1, v6)
    v7 := GenS1{F0: 380}
    if false {
        for fi8 := 0; fi8 < 19; fi8++ {
            if (false && (!(false || true))) {
                var v9 uint8 = (v3 << ((255) & 63))
                _ = v9
            }
            for fi10 := 0; fi10 < 18; fi10++ {
                if true {
                    var v11 int = (fi8 % (((fi8 % ((9223372036854775807) | 1))) | 1))
                    if (930 <= 255) {
                        if (-1 != 513) {
                            var v12 int32 = -437
                            var v13 uint = 1
                            _ = v12
                            _ = v13
                        } else {
                            fmt.Println(v11, v11)
                        }
                        if (v3 <= 128) {
                            v1 += v1
                        }
                    } else {
                        var v14 float32 = v6
                        switch "\n\t" {
                        case "\n\t":
                            v11 -= fi8
                        case "hello":
                            v15 := GenS1{F0: 42}
                            _ = v15
                        default:
                            fmt.Println(v6)
                        }
                        v2 = (v2 & 0)
                        _ = v14
                    }
                    fmt.Println(v2, v6)
                    v16 := GenS1{F0: 935}
                    _ = v11
                    _ = v16
                } else {
                    v17 := GenS2{F0: (^(v1)), F1: 297}
                    fmt.Println(v3)
                    if true {
                        v3 -= 139
                        v3 = v3
                    }
                    _ = v17
                }
                _ = fi10
            }
            if (-24 <= -128) {
                _ = len(v5)
                v1 = (v1 >> ((0) & 63))
                var v18 int64 = 1
                _ = v18
            }
            _ = fi8
        }
        v6 = v6
    }
    v19 := GenS3{F0: 42, F1: 477, F2: -32768}
    if false {
        var v20 int = -2147483648
        var v21 int = v20
        v5 = append(v5, v21)
        _ = v20
        _ = v21
    }
    if true {
        for fi22 := 0; fi22 < 12; fi22++ {
            for fi23 := 0; fi23 < 2; fi23++ {
                if (true || true) {
                    var v24 uint64 = 37
                    _ = v24
                }
                v25 := GenS4{F0: v2}
                _ = fi23
                _ = v25
            }
            switch (fi22 | fi22) {
            case -693:
                _ = len(v5)
            case 127:
                var v26 string = "foo bar"
                _ = v26
            case 42:
                for fi27 := 0; fi27 < 11; fi27++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(v4, v4)
                    _ = fi27
                }
            default:
                for fi28 := 0; fi28 < 11; fi28++ {
                    v2 = v2
                    if false {
                        var v29 string = "hello"
                        _ = v29
                    }
                    _ = fi28
                }
            }
            v2 -= v2
            _ = fi22
        }
        v30 := GenS5{F0: 212}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v30
    }
    v1 -= 1
    v6 = 0.1
    if (-9223372036854775808 >= -223) {
        if (217 != 0) {
            var v31 uint16 = 1
            _ = v31
        } else {
            var v32 int = 32768
            if true {
                v33 := GenS1{F0: 2}
                v34 := GenS1{F0: 0}
                _ = v33
                _ = v34
            }
            _ = v32
        }
        v35 := GenS6{F0: 1, F1: 301}
        _ = v35
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v36 := GenS3{F0: -1, F1: 998, F2: -1}
        var v37 bool = (!false)
        _ = v36
        _ = v37
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
