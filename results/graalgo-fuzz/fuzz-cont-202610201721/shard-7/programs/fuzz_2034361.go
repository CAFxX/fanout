package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int32 = -181
    v2 := make(map[string]bool)
    var v3 uint16 = 65535
    var v4 uint16 = 947
    var v5 int64 = 9223372036854775807
    var v6 uint64 = 2
    var v7 uint16 = v4
    if ("foo bar" != "a") {
        v8 := GenS1{F0: (^(v5)), F1: 18446744073709551615, F2: 0}
        v9 := GenS1{F0: v5, F1: 833, F2: 18446744073709551615}
        _ = v8
        _ = v9
    } else {
        var v10 int = -2147483648
        _ = v10
    }
    v11 := []int{-1, 816, 0, -653}
    v12 := GenS2{F0: (v6 << ((1) & 63)), F1: v1, F2: 1}
    if false {
        var v13 bool = ((!(42 < -9223372036854775808)) && false)
        if (0.0 == 3.14159) {
            v13 = v13
        }
        _ = v13
    } else {
        if (float64(-2.25) == (1.5 * 0.0)) {
            _ = len(v2)
        }
    }
    var v14 int16 = 32767
    if (("\n\t" + ("hello" + "foo bar")) < "") {
        var v15 int16 = v14
        v15 -= v14
        var v16 uint = 255
        _ = v15
        _ = v16
    } else {
        var v17 float64 = 0.0
        if ((v3 << ((0) & 63)) == v7) {
            if len(v11) > 0 {
                v11[0] = 4294967295
            }
            v2[("a" + "a")] = (v1 != (^(v1)))
        } else {
            v18 := GenS1{F0: v5, F1: 0, F2: 495}
            var v19 int = 32767
            _ = v18
            _ = v19
        }
        v6 += v6
        _ = v17
    }
    v5 = (v5 / ((v5) | 1))
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v11 = append(v11, 2147483647)
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS1{F0: v5, F1: 368, F2: 2}
        switch "" {
        case "foo bar":
            fmt.Println(v6)
        case "a":
            v1 = (v1 ^ (v1 % (((v1 - 0)) | 1)))
        }
        v2[((("" + "\n\t") + "") + "a")] = (((0 < 1) && ((v3 - v3) > 531)) && true)
        _ = fi20
        _ = v21
    }
    if false {
        v22 := GenS1{F0: v5, F1: 526, F2: 18446744073709551615}
        var v23 int8 = -56
        if (v6 >= (v6 + (v6 & v6))) {
            if ("a" < "\n\t") {
                var v24 int8 = (v23 / ((v23) | 1))
                var v25 string = "a"
                v3 += v4
                _ = v24
                _ = v25
            }
        }
        _ = v22
        _ = v23
    } else {
        v26 := GenS1{F0: v5, F1: 42, F2: 999}
        switch -416 {
        case -9223372036854775808:
            v4 -= v7
        default:
            for fi27 := 0; fi27 < 10; fi27++ {
                for fi28 := 0; fi28 < 5; fi28++ {
                    fmt.Println(v6, v14, v3)
                    if false {
                        v29 := GenS3{F0: 255, F1: 42}
                        var v30 bool = false
                        var v31 uint8 = 42
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    var v32 uint8 = 0
                    v3 = v4
                    v11 = append(v11, fi28)
                    _ = fi28
                    _ = v32
                }
                var v33 int = (fi27 << ((18446744073709551615) & 63))
                switch v33 {
                case 26:
                    var v34 int32 = (v1 & -2147483648)
                    _ = v34
                case 9223372036854775807:
                    v35 := GenS3{F0: 2, F1: 2}
                    _ = v35
                case 1:
                    delete(v2, "\n\t")
                default:
                    v6 += 176
                }
                _ = fi27
                _ = v33
            }
        }
        v11 = append(v11, 1)
        _ = v26
    }
    v36 := GenS4{F0: 620}
    v4 = v3
    var v37 float32 = ((float32((0.1 + 1e10)) + 1.5) + (0.1 - 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
