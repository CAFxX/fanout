package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint64; F2 int }
type GenS3 struct { F0 int16 }
func genFn0() bool {
    var v5 string = "a"
    for fi6 := 0; fi6 < 19; fi6++ {
        switch (v5 + "hello") {
        case "hello":
            if (-2147483648 >= 0) {
                for fi7 := 0; fi7 < 20; fi7++ {
                    v5 = v5
                    if (true && (true && false)) {
                        var v8 int16 = -1
                        v8 = 32767
                        v9 := GenS1{F0: (970 <= int(v8))}
                        _ = v8
                        _ = v9
                    }
                    var v10 int16 = 42
                    _ = fi7
                    _ = v10
                }
                var v11 int64 = 220
                if false {
                    var v12 uint = 2
                    _ = v12
                }
                _ = v11
            } else {
                v5 = v5
                v5 = "\n\t"
            }
        case "foo bar":
            switch fi6 {
            case 128:
                v13 := []int{0, (fi6 & fi6)}
                _ = v13
            default:
                v14 := GenS1{F0: true}
                _ = v14
            }
        }
        if false {
            v15 := []int{(fi6 % ((fi6) | 1))}
            v5 = v5
            v15 = append(v15, (fi6 * (^(fi6))))
            _ = v15
        } else {
            v5 = (v5 + "hello")
        }
        var v16 int8 = 0
        _ = fi6
        _ = v16
    }
    if (((-76 > 0) || (0 >= 2)) && ((!true) || (1e10 == 0.1))) {
        var v17 float64 = 1e10
        _ = v17
    }
    fmt.Println(v5)
    _ = v5
    return true
}

func main() {
    var v1 float32 = 1.5
    var v2 uint = 4294967295
    var v3 bool = false
    var v4 uint64 = 872
    if (-2147483648 >= -2147483648) {
        v18 := []int{-9223372036854775808, 42, -2147483648, 42}
        var v19 int8 = 0
        v20 := GenS2{F0: (v2 - (v2 + 4294967295)), F1: 0, F2: 0}
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := []int{0}
    v1 = v1
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 uint64 = v4
        _ = fi22
        _ = v23
    }
    _ = len(v21)
    v2 -= v2
    switch 9223372036854775807 {
    case 127:
        if (-2.25 == 1e10) {
            v24 := GenS2{F0: 74, F1: (v4 / (((v4 * (v4 + v4))) | 1)), F2: 4294967295}
            if ("" > "a") {
                fmt.Println(v2, v1)
                fmt.Println(v4, v3, v3)
                v2 -= v2
            } else {
                switch "hello" {
                case "\n\t":
                    var v25 float32 = v1
                    _ = v25
                default:
                    v26 := GenS1{F0: (141 != 32767)}
                    _ = v26
                }
                v1 = 1e10
            }
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            _ = v24
        } else {
            v27 := GenS3{F0: -1}
            _ = v27
        }
    case -1:
        v21 = append(v21, 9223372036854775807)
    case 42:
        var v28 uint8 = 102
        _ = v28
    default:
        v2 += v2
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        v2 = (^(v2))
        var v30 uint = v2
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 uint8 = 0
        _ = fi31
        _ = v32
    }
    if (v3 || (!true)) {
        var v33 int64 = 42
        for fi34 := 0; fi34 < 2; fi34++ {
            if ((!v3) && (743 == 42)) {
                v35 := GenS3{F0: -1}
                _ = v35
            }
            if len(v21) > 0 {
                v21[0] = (fi34 / ((int(v33)) | 1))
            }
            _ = fi34
        }
        v36 := GenS3{F0: 299}
        _ = v33
        _ = v36
    }
    _ = len(v21)
    v4 += v4
    var v37 string = ("" + "foo bar")
    fmt.Println(v2)
    switch 128 {
    case 646:
        for fi38 := 0; fi38 < 11; fi38++ {
            fmt.Println(v1, v1)
            _ = fi38
        }
    case 1:
        fmt.Println(v3, v1, v2)
    default:
        var v39 int16 = -721
        _ = v39
    }
    var v40 int16 = 32767
    v4 = (v4 & (v4 - (v4 % ((v4) | 1))))
    var v41 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
}
