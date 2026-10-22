package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint8; F2 int8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint8 = 255
    var v2 int = 131
    v3 := []int{v2, (v2 | v2), 435}
    var v4 int64 = -9223372036854775808
    v5 := []int{}
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int = v2
        v4 -= (-(v4))
        var v8 int32 = -1
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint32 = 2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v10 uint = 42
    var v11 int64 = (v4 + 9223372036854775807)
    var v12 uint32 = (v9 >> ((v10) & 63))
    _ = len(v5)
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: 1}
        if ((int32(v1) > -2147483648) && false) {
            fmt.Println(v2, v12)
            v4 = v4
            for fi15 := 0; fi15 < 14; fi15++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (4294967295 >= v9) {
                    var v16 uint = (^(v10))
                    v1 = (v1 >> (((v10 | 1)) & 63))
                    if false {
                        if ((1.5 + 3.14159) == 3.14159) {
                            v17 := GenS1{F0: 127}
                            if ((v11 & (v4 ^ v11)) > (v11 + v4)) {
                                var v18 int = (^(fi13))
                                _ = v18
                            }
                            _ = v17
                        } else {
                            switch "foo bar" {
                            case "":
                                if len(v3) > 0 {
                                    v3[0] = 2147483647
                                }
                            case "\n\t":
                                if false {
                                    switch ("hello" + ("a" + "foo bar")) {
                                    case "hello":
                                        v9 += v9
                                    case "foo bar":
                                        var v19 uint16 = 32767
                                        _ = v19
                                    }
                                }
                            case "hello":
                                if (!(!true)) {
                                    if true {
                                        v16 = v10
                                        var v20 uint32 = v9
                                        var v21 uint = v10
                                        _ = v20
                                        _ = v21
                                    } else {
                                        v22 := GenS2{F0: (^(v11)), F1: v1, F2: 1}
                                        var v23 string = ("hello" + ("hello" + ("foo bar" + "hello")))
                                        switch (fi15 * -9223372036854775808) {
                                        case 523:
                                            v5 = append(v5, 865)
                                        }
                                        _ = v22
                                        _ = v23
                                    }
                                    var v24 uint16 = 2
                                    _ = v24
                                } else {
                                    var v25 int32 = 2147483647
                                    switch "a" {
                                    case "a":
                                        switch ("a" + "\n\t") {
                                        case "foo bar":
                                            v4 = (v4 << ((679) & 63))
                                        case "hello":
                                            v5 = append(v5, fi15)
                                        case "a":
                                            v2 = (v2 ^ fi13)
                                        default:
                                            v9 = v9
                                        }
                                    case "foo bar":
                                        var v26 int64 = (^(v11))
                                        _ = v26
                                    default:
                                        var v27 string = "a"
                                        _ = v27
                                    }
                                    _ = v25
                                }
                            }
                            v2 = fi13
                            var v28 uint16 = 0
                            v9 = v9
                            switch "" {
                            case "foo bar":
                                switch (v2 + (fi15 & (v2 >> ((1) & 63)))) {
                                case -775:
                                    v11 += 42
                                    v9 = 1
                                case 127:
                                    switch 32767 {
                                    case 369:
                                        v4 = (v11 % ((-229) | 1))
                                    default:
                                        switch "" {
                                        case "\n\t":
                                            v29 := GenS1{F0: -110}
                                            _ = v29
                                        case "foo bar":
                                            var v30 bool = true
                                            var v31 float32 = float32(float64(0.0))
                                            _ = v30
                                            _ = v31
                                        case "":
                                            v3 = append(v3, 42)
                                        }
                                    }
                                default:
                                    var v32 uint = 729
                                    _ = v32
                                }
                            case "hello":
                                if false {
                                    var v33 float64 = 0.1
                                    if (true && true) {
                                        v12 = 2147483648
                                        v34 := GenS2{F0: 0, F1: uint8(v16), F2: -128}
                                        _ = v34
                                    }
                                    _ = v33
                                }
                                if ((!false) && (42 != 42)) {
                                    if false {
                                        v1 = 53
                                        v35 := GenS3{F0: -625, F1: 251}
                                        _ = v35
                                    } else {
                                        var v36 bool = (v4 <= (-(v11)))
                                        var v37 string = ""
                                        _ = v36
                                        _ = v37
                                    }
                                    v38 := GenS1{F0: 42}
                                    v1 = v1
                                    _ = v38
                                }
                                fmt.Println(v10, v9)
                            }
                            _ = v28
                        }
                    }
                    _ = v16
                } else {
                    if len(v5) > 0 {
                        v5[0] = (fi13 + v2)
                    }
                    if (float64(1.5) >= -2.25) {
                        if true {
                            v3 = append(v3, (fi13 + (fi13 ^ (fi13 % ((267) | 1)))))
                            if len(v5) > 0 {
                                v5[0] = (v2 % (((fi13 & 2147483647)) | 1))
                            }
                            v39 := GenS3{F0: 128, F1: v12}
                            switch "\n\t" {
                            case "foo bar":
                                v5 = append(v5, fi13)
                            case "":
                                v40 := GenS1{F0: -120}
                                _ = v40
                            }
                            v10 = (v10 >> ((v10) & 63))
                            _ = v39
                        }
                    }
                }
                _ = fi15
            }
        } else {
            v41 := GenS4{F0: 1}
            v42 := GenS2{F0: v11, F1: v1, F2: int8(fi13)}
            _ = v41
            _ = v42
        }
        _ = fi13
        _ = v14
    }
    for fi43 := 0; fi43 < 11; fi43++ {
        if len(v3) > 0 {
            v3[0] = v2
        }
        if (false && false) {
            switch (fi43 | (fi43 * -1)) {
            case 4294967295:
                var v44 string = "\n\t"
                _ = v44
            case 691:
                switch -2147483648 {
                case 4294967295:
                    var v45 string = "\n\t"
                    _ = v45
                case -9223372036854775808:
                    if false {
                        fmt.Println(v10, v12)
                    } else {
                        var v46 uint32 = v12
                        _ = v46
                    }
                default:
                    for fi47 := 0; fi47 < 12; fi47++ {
                        v11 = 42
                        _ = len(v3)
                        v11 = v11
                        _ = fi47
                    }
                }
            case 837:
                v48 := GenS4{F0: int8(v1)}
                _ = v48
            default:
                v10 = v10
            }
        }
        _ = fi43
    }
    if (1 == (^(v11))) {
        var v49 bool = true
        v9 = 1
        _ = v49
    } else {
        if len(v5) > 0 {
            v5[0] = v2
        }
        v50 := GenS5{F0: 100, F1: (v9 - (v12 >> (((v10 << ((1) & 63))) & 63)))}
        _ = v50
    }
    v12 = (v9 % ((v12) | 1))
    fmt.Println(v9, v10, v1)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v3) > 0 {
        v3[0] = v2
    }
    var v51 float64 = 3.14159
    v51 = float64(0.0)
    if false {
        var v52 float64 = (3.14159 * v51)
        for fi53 := 0; fi53 < 18; fi53++ {
            var v54 uint8 = v1
            _ = fi53
            _ = v54
        }
        v3 = append(v3, (v2 | (v2 | int(v4))))
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v9
}
