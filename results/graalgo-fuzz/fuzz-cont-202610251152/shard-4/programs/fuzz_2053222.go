package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint32 }
type GenS2 struct { F0 string }

func main() {
    var v1 float32 = (0.1 * ((3.14159 / 1.0) - -2.25))
    var v2 int8 = 100
    var v3 string = ("a" + (("foo bar" + "hello") + "\n\t"))
    if true {
        fmt.Println(v2, v2, v3)
        v4 := []int{128, -9223372036854775808}
        _ = v4
    } else {
        if (!true) {
            fmt.Println(v3)
        }
        var v5 float64 = float64(v1)
        _ = v5
    }
    var v6 uint8 = 1
    if (((v6 * (v6 + 127)) < (v6 / (((v6 ^ 2)) | 1))) || (uint16(v6) <= 1)) {
        v6 += (v6 / ((v6) | 1))
        fmt.Println(v6, v6)
        v2 -= (v2 / ((v2) | 1))
    }
    v6 = v6
    var v7 float64 = 3.14159
    switch -2147483648 {
    case 495:
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint64 = 784
            switch fi8 {
            case 27:
                v6 -= 127
            default:
                if (uint16(v9) < 854) {
                    var v10 float32 = (-2.25 + 1.5)
                    fmt.Println(fi8)
                    _ = v10
                }
            }
            if false {
                v1 = 0.0
                switch (v3 + v3) {
                case "hello":
                    var v11 uint = 804
                    _ = v11
                case "":
                    v2 += v2
                case "a":
                    switch fi8 {
                    case -403:
                        for fi12 := 0; fi12 < 6; fi12++ {
                            var v13 uint16 = 313
                            var v14 uint32 = 42
                            v15 := []int{}
                            var v16 float32 = v1
                            _ = fi12
                            _ = v13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    case 42:
                        v9 += (v9 ^ (v9 | v9))
                    }
                default:
                    switch (fi8 - (-(fi8))) {
                    case -2147483648:
                        for fi17 := 0; fi17 < 3; fi17++ {
                            if true {
                                var v18 int64 = -1
                                var v19 uint16 = 65535
                                if (-1 != 844) {
                                    var v20 uint = 65535
                                    _ = v20
                                } else {
                                    fmt.Println(v9, v18)
                                    v21 := GenS1{F0: 1000, F1: (v18 & (v18 + (v18 * v18))), F2: 2147483647}
                                    _ = v21
                                }
                                _ = v18
                                _ = v19
                            } else {
                                v22 := GenS1{F0: -1, F1: 0, F2: 2147483648}
                                var v23 int = fi8
                                v2 = 100
                                _ = v22
                                _ = v23
                            }
                            var v24 float32 = ((v1 + 1.5) - 1e10)
                            v6 -= v6
                            _ = fi17
                            _ = v24
                        }
                    case 32768:
                        fmt.Println(v3, v1)
                    }
                }
                if ((((v2 | (v2 + v2)) < (v2 << ((42) & 63))) || (879 >= -961)) || ((v9 >> ((808) & 63)) != 1)) {
                    v2 = (^(v2))
                }
            }
            _ = fi8
            _ = v9
        }
    case -671:
        v25 := []int{-1}
        _ = v25
    default:
        v26 := []int{785}
        _ = v26
    }
    if (646 < 1) {
        v27 := GenS2{F0: v3}
        _ = v27
    }
    if false {
        var v28 uint32 = 324
        _ = v28
    }
    v29 := []int{1, 9223372036854775807, -733, 32767}
    var v30 int8 = 1
    v2 += (v30 % (((v30 * v30)) | 1))
    switch v3 {
    case "a":
        _ = len(v29)
    case "hello":
        var v31 int32 = 42
        _ = v31
    default:
        fmt.Println(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v6
    _ = v7
}
