package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint; F1 bool; F2 int32 }
type GenS4 struct { F0 int8; F1 int64; F2 uint16 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int32; F1 uint64; F2 uint64 }

func main() {
    var v1 int = -1
    var v2 uint8 = 60
    var v3 string = "foo bar"
    switch v3 {
    case "foo bar":
        switch 32768 {
        case 32768:
            v2 = (^(v2))
        case 813:
            for fi4 := 0; fi4 < 13; fi4++ {
                if false {
                    switch (v3 + v3) {
                    case "":
                        var v5 uint8 = (^(v2))
                        _ = v5
                    case "foo bar":
                        for fi6 := 0; fi6 < 8; fi6++ {
                            var v7 int32 = 2147483647
                            _ = fi6
                            _ = v7
                        }
                    default:
                        fmt.Println(fi4)
                    }
                }
                _ = fi4
            }
        default:
            var v8 uint8 = 255
            _ = v8
        }
    case "hello":
        if (1 >= 817) {
            v9 := GenS1{F0: v3, F1: (v2 << ((42) & 63)), F2: true}
            _ = v9
        }
    }
    var v10 float32 = 1e10
    v3 = "\n\t"
    fmt.Println(v10, v2)
    var v11 float64 = float64((0.1 / -2.25))
    for fi12 := 0; fi12 < 17; fi12++ {
        switch 892 {
        case 1:
            switch ((v3 + "\n\t") + v3) {
            case "hello":
                v13 := GenS1{F0: v3, F1: (v2 & (v2 >> ((517) & 63))), F2: ((v11 - float64(v10)) < -2.25)}
                _ = v13
            }
        case -721:
            v10 = 1e10
        case 32768:
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 string = v3
                v16 := GenS2{F0: 2147483647, F1: 2147483647, F2: 18446744073709551615}
                v17 := []int{(v1 | fi12), (v1 ^ fi14), (fi12 % ((int(v2)) | 1)), fi14}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        }
        _ = fi12
    }
    if (v3 < v3) {
        for fi18 := 0; fi18 < 10; fi18++ {
            v2 -= (v2 + v2)
            _ = fi18
        }
    } else {
        var v19 float32 = v10
        if true {
            v20 := []int{}
            if len(v20) > 0 {
                v20[0] = v1
            }
            _ = v20
        }
        var v21 uint64 = 42
        _ = v19
        _ = v21
    }
    fmt.Println(v2, v1)
    if (false && (((!true) && true) && false)) {
        v10 = 0.0
        v22 := []int{-681, v1}
        var v23 bool = false
        _ = v22
        _ = v23
    } else {
        v24 := GenS2{F0: -601, F1: 475, F2: 769}
        _ = v24
    }
    var v25 int32 = 256
    v26 := []int{int(v2), v1}
    v11 -= (v11 + (0.0 * (3.14159 * 3.14159)))
    v11 = v11
    v27 := GenS3{F0: 4294967295, F1: true, F2: (v25 - (v25 - v25))}
    if false {
        v26 = append(v26, 0)
    } else {
        v28 := GenS4{F0: 20, F1: -801, F2: 1}
        switch v3 {
        case "a":
            var v29 int32 = -479
            _ = v29
        case "\n\t":
            if (((255 > 0) && true) || ((int16(v25) == 42) && (v25 > v25))) {
                v11 = 1.5
                var v30 uint64 = 42
                _ = v30
            } else {
                v10 = 0.0
            }
        default:
            if ((("a" + "") + "foo bar") == "") {
                var v31 uint16 = 389
                fmt.Println(v11, v31, v31)
                v32 := GenS5{F0: 651, F1: false}
                _ = v31
                _ = v32
            } else {
                v33 := GenS6{F0: 2147483647, F1: 18446744073709551615, F2: 2}
                v3 = v3
                v1 += v1
                _ = v33
            }
        }
        _ = v28
    }
    var v34 float64 = (0.1 - (((float64(v10) - 1e10) + v11) + float64(1.5)))
    var v35 uint8 = (v2 ^ (v2 >> ((42) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
}
