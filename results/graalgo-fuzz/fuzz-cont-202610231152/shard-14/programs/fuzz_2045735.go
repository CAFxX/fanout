package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint64; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    v3 := []int{-2147483648, -9223372036854775808, 482}
    v4 := make(map[int]bool)
    v5 := []uint32{0}
    switch "foo bar" {
    case "foo bar":
        var v6 float64 = (-2.25 - 0.1)
        _ = v6
    case "hello":
        v7 := GenS1{F0: false, F1: 190, F2: 969}
        _ = v7
    default:
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 string = (("\n\t" + ("\n\t" + "\n\t")) + "hello")
            if (1 == 1) {
                var v10 uint32 = 823
                v11 := v1[42]
                v11 = v10
                _ = v10
                _ = v11
            } else {
                if (v9 > (v9 + ((("" + v9) + v9) + (v9 + "a")))) {
                    if ((3.14159 * 0.1) < ((-2.25 * (1.5 - float64(float32((-2.25 * 0.1))))) / float64(3.14159))) {
                        switch (^(fi8)) {
                        case -2147483648:
                            v9 = "a"
                        case 1:
                            if (v9 != (((v9 + ("hello" + "\n\t")) + v9) + "a")) {
                                if (0 <= 4294967295) {
                                    v12 := GenS1{F0: false, F1: 1, F2: 4294967295}
                                    var v13 int8 = 18
                                    _ = v12
                                    _ = v13
                                }
                            }
                        }
                        for fi14 := 0; fi14 < 16; fi14++ {
                            var v15 int32 = 37
                            _ = fi14
                            _ = v15
                        }
                    } else {
                        v9 = v9
                        v16 := GenS1{F0: true, F1: -1, F2: 4294967295}
                        fmt.Println(v9, v9)
                        _ = v16
                    }
                }
                var v17 int = 32768
                switch (v9 + "a") {
                case "":
                    var v18 int = (v17 / ((fi8) | 1))
                    _ = v18
                case "hello":
                    switch 4294967295 {
                    case -959:
                        if (!true) {
                            switch (fi8 / ((42) | 1)) {
                            case 464:
                                if false {
                                    var v19 int64 = 42
                                    _ = v19
                                }
                            }
                            if true {
                                var v20 float64 = (1e10 * 0.1)
                                v21 := GenS1{F0: true, F1: -2147483648, F2: 2147483647}
                                _ = v20
                                _ = v21
                            } else {
                                v22 := GenS2{F0: true, F1: -1}
                                v23 := GenS3{F0: 337, F1: -128}
                                _ = v22
                                _ = v23
                            }
                        } else {
                            v17 = (fi8 % (((fi8 >> ((330) & 63))) | 1))
                        }
                    case 1:
                        v17 = (fi8 & fi8)
                    case 2147483647:
                        v9 = "\n\t"
                    }
                case "\n\t":
                    v24 := GenS2{F0: false, F1: -128}
                    _ = v24
                default:
                    fmt.Println(fi8)
                }
                _ = v17
            }
            var v25 int8 = 84
            _ = fi8
            _ = v25
            _ = v9
        }
    }
    var v26 uint32 = 1
    var v27 int = 9223372036854775807
    v28 := GenS3{F0: 0, F1: -1}
    if false {
        fmt.Println(v26)
        var v29 int = -640
        var v30 int32 = -1
        _ = v29
        _ = v30
    }
    if ((!(1 != 42)) && (-9223372036854775808 > (v27 - v27))) {
        v31 := GenS3{F0: 2, F1: 1}
        var v32 int16 = 1000
        var v33 float32 = (3.14159 + 1e10)
        _ = v31
        _ = v32
        _ = v33
    }
    v27 = (v27 * v27)
    if (-113 >= -683) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v27 = 1
    v34 := v4[(v27 * -122)]
    v27 -= (^(v27))
    v34 = v34
    for fi35 := 0; fi35 < 3; fi35++ {
        v26 -= 890
        var v36 uint64 = 0
        v27 = 42
        _ = fi35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
