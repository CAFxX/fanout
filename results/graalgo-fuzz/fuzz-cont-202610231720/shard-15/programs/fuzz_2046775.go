package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 int32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int16; F1 int64; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 int16; F2 int32 }
type GenS9 struct { F0 int16; F1 int16 }
func genFn0() string {
    var v5 uint32 = 42
    v6 := []int{1, 9223372036854775807, 9223372036854775807}
    if true {
        switch "foo bar" {
        case "hello":
            var v7 int32 = 42
            _ = v7
        }
    } else {
        v8 := GenS1{F0: (^(v5)), F1: 18446744073709551615, F2: -260}
        _ = v8
    }
    var v9 float64 = -2.25
    v6 = append(v6, -9223372036854775808)
    _ = v5
    _ = v6
    _ = v9
    return ("foo bar" + ("" + "hello"))
}

func main() {
    var v1 uint8 = 255
    var v2 float64 = float64(3.14159)
    var v3 int16 = 0
    var v4 int8 = 0
    var v10 int = -290
    var v11 uint8 = (v1 ^ 2)
    var v12 int16 = v3
    if false {
        switch "a" {
        case "a":
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 bool = false
                for fi15 := 0; fi15 < 13; fi15++ {
                    v16 := GenS1{F0: 131, F1: 0, F2: 786}
                    var v17 uint16 = 674
                    switch "hello" {
                    case "":
                        v14 = (!false)
                    case "foo bar":
                        switch (("" + "a") + ("" + "hello")) {
                        case "hello":
                            if (int16(v17) != (v12 / (((v12 & 42)) | 1))) {
                                if v14 {
                                    switch (("a" + ("hello" + "foo bar")) + "\n\t") {
                                    case "a":
                                        var v18 bool = ((v14 && (65535 == 508)) || (42 >= -1))
                                        _ = v18
                                    }
                                    v4 = (-(v4))
                                    v19 := []int{32768, (fi15 << ((255) & 63))}
                                    _ = v19
                                }
                                v20 := []int{(v10 ^ (fi15 ^ 375)), (fi13 % ((v10) | 1)), v10}
                                if true {
                                    v10 = 29
                                } else {
                                    v21 := GenS1{F0: 2147483647, F1: 42, F2: 597}
                                    v22 := GenS2{F0: v3, F1: (true || (1 < -1)), F2: 0}
                                    _ = v21
                                    _ = v22
                                }
                                _ = v20
                            } else {
                                if true {
                                    switch "hello" {
                                    case "":
                                        v14 = (0 != 0)
                                    }
                                    v23 := []int{(fi13 & (v10 & fi13)), (^(fi15)), v10, 2147483647}
                                    if false {
                                        fmt.Println(fi13)
                                    } else {
                                        v24 := GenS3{F0: v4, F1: uint16(v11)}
                                        _ = len(v23)
                                        v17 -= v17
                                        _ = v24
                                    }
                                    _ = v23
                                }
                                var v25 int32 = 1
                                v26 := GenS4{F0: (v3 - 0), F1: -9223372036854775808, F2: v10}
                                _ = v25
                                _ = v26
                            }
                        case "\n\t":
                            var v27 int32 = -80
                            _ = v27
                        case "":
                            if (1 >= v10) {
                                switch fi13 {
                                case 128:
                                    v28 := GenS5{F0: 255}
                                    _ = v28
                                case -1:
                                    var v29 int32 = 843
                                    _ = v29
                                default:
                                    v30 := GenS6{F0: ("hello" + ""), F1: 32767}
                                    _ = v30
                                }
                            } else {
                                switch (("\n\t" + "\n\t") + "\n\t") {
                                case "foo bar":
                                    var v31 int32 = -1
                                    _ = v31
                                case "\n\t":
                                    v10 = v10
                                }
                                if v14 {
                                    fmt.Println(v17)
                                    switch fi15 {
                                    case 500:
                                        var v32 uint64 = 42
                                        _ = v32
                                    }
                                    v10 -= 2147483647
                                } else {
                                    v17 -= 312
                                }
                                v33 := []int{}
                                v11 += v11
                                var v34 uint32 = 2147483648
                                v35 := GenS4{F0: (v3 / ((int16(v11)) | 1)), F1: 1, F2: v10}
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        default:
                            var v36 uint16 = v17
                            _ = v36
                        }
                    case "\n\t":
                        var v37 int16 = 32767
                        _ = v37
                    }
                    fmt.Println(v4)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                for fi38 := 0; fi38 < 18; fi38++ {
                    v39 := []int{42, (fi38 << ((4294967295) & 63)), (fi13 ^ v10)}
                    _ = fi38
                    _ = v39
                }
                _ = fi13
                _ = v14
            }
        case "\n\t":
            fmt.Println(v11, v11)
        }
        var v40 int64 = 9223372036854775807
        v41 := []int{-834, v10, v10}
        _ = v40
        _ = v41
    }
    v42 := GenS7{F0: (65535 != 65535)}
    switch (v10 + -9223372036854775808) {
    case 42:
        v43 := []int{int(v11), 352}
        _ = v43
    default:
        for fi44 := 0; fi44 < 6; fi44++ {
            v2 = (v2 * ((1.5 - ((float64((float32(0.0) / 1e10)) * float64(float32(v2))) + v2)) + v2))
            var v45 uint = 923
            var v46 int = 930
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    v1 += 128
    if true {
        for fi47 := 0; fi47 < 20; fi47++ {
            v48 := []int{}
            var v49 uint8 = v11
            if true {
                v48 = append(v48, (v10 & 2147483647))
                var v50 bool = (((true && true) || true) && (!(!true)))
                _ = v50
            } else {
                v51 := GenS7{F0: true}
                _ = v51
            }
            _ = fi47
            _ = v48
            _ = v49
        }
        v3 += 42
        v52 := GenS3{F0: 42, F1: 42}
        _ = v52
    } else {
        v53 := []int{}
        v4 -= (v4 % (((v4 | 127)) | 1))
        _ = v53
    }
    var v54 bool = false
    fmt.Println(v2, v11, v2)
    fmt.Println(v4)
    v55 := GenS8{F0: 4294967295, F1: v12, F2: 2147483647}
    for fi56 := 0; fi56 < 15; fi56++ {
        v4 -= (v4 + v4)
        v57 := GenS4{F0: (v12 << ((377) & 63)), F1: 1, F2: (v10 + int(v12))}
        _ = fi56
        _ = v57
    }
    for fi58 := 0; fi58 < 3; fi58++ {
        if ("" > "hello") {
            fmt.Println(v10)
            v59 := GenS7{F0: v54}
            var v60 uint64 = 0
            _ = v59
            _ = v60
        }
        v61 := GenS9{F0: (^(v12)), F1: (v12 / ((32767) | 1))}
        v62 := GenS7{F0: (0.1 != ((0.0 / float64(1e10)) - 3.14159))}
        _ = fi58
        _ = v61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v54
    _ = v55
}
