package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int16; F2 string }
type GenS3 struct { F0 uint; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 uint16; F2 bool }
func genFn0(p0 uint64) string {
    p0 = (^(p0))
    var v4 string = "a"
    v5 := GenS1{F0: 438}
    v6 := []int{}
    v4 = "hello"
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return v4
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = -2.25
    var v3 int8 = 45
    if (((42 != 42) && (42 != 1)) || ((true && (0 != 9223372036854775807)) && (!true))) {
        var v7 int16 = 0
        if ("\n\t" >= "a") {
            if (-626 > 1) {
                v3 = v3
                switch "foo bar" {
                case "hello":
                    v8 := GenS2{F0: 0, F1: (v7 | 1000), F2: "\n\t"}
                    _ = v8
                }
            }
        }
        v7 = (v7 << ((4294967295) & 63))
        _ = v7
    }
    var v9 int8 = (v3 * v3)
    var v10 float32 = ((1.5 + 1.5) / ((0.1 + 0.1) / 0.1))
    v3 += v3
    var v11 uint8 = 165
    fmt.Println(v9)
    v12 := []int{285, 42, 9223372036854775807, 1}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v10 -= float32(v2)
    var v13 uint16 = 42
    var v14 float32 = v10
    for fi15 := 0; fi15 < 4; fi15++ {
        v9 = v9
        var v16 uint16 = (v13 >> ((uint(v13)) & 63))
        if false {
            delete(v1, 2147483647)
            if true {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                v17 := v1[2147483647]
                _ = v17
            } else {
                v11 = 255
                v13 -= (v13 * v16)
            }
            var v18 int64 = 755
            _ = v18
        } else {
            if (667 >= -197) {
                for fi19 := 0; fi19 < 9; fi19++ {
                    v11 = v11
                    v11 = (^(v11))
                    if (275 <= 1) {
                        var v20 float64 = v2
                        switch int(v3) {
                        case 0:
                            v10 += float32(((v2 * (v2 * -2.25)) * (v2 * v20)))
                        case 9223372036854775807:
                            v21 := GenS3{F0: 255, F1: 4294967295, F2: 65535}
                            _ = v21
                        case -828:
                            var v22 float64 = ((((v20 - 0.0) / 3.14159) / float64(v14)) - ((float64(0.1) - 0.1) * v20))
                            _ = v22
                        default:
                            v10 = (v10 / (v10 / (v14 / (0.1 - -2.25))))
                        }
                        switch int(v13) {
                        case -611:
                            var v23 int16 = -1
                            _ = v23
                        case 32768:
                            if false {
                                if false {
                                    var v24 uint8 = 2
                                    if len(v12) > 0 {
                                        v12[0] = (fi15 / (((fi15 | (fi19 & (-(fi15))))) | 1))
                                    }
                                    _ = v24
                                }
                            } else {
                                if true {
                                    fmt.Println(v11)
                                    var v25 uint16 = 616
                                    v10 = v10
                                    _ = v25
                                } else {
                                    delete(v1, uint32(v16))
                                }
                                _ = len(v12)
                                var v26 uint = 255
                                _ = v26
                            }
                        }
                        _ = v20
                    } else {
                        switch ("" + "") {
                        case "\n\t":
                            v13 = 418
                        case "foo bar":
                            switch fi15 {
                            case -899:
                                v27 := GenS3{F0: 1, F1: 65535, F2: 429}
                                _ = v27
                            }
                        case "":
                            if (-2147483648 >= 456) {
                                if true {
                                    fmt.Println(v14, v14, v2)
                                    var v28 bool = true
                                    _ = v28
                                }
                                var v29 int32 = 732
                                _ = v29
                            }
                        default:
                            v16 += v16
                        }
                        if (v3 >= (-(v9))) {
                            var v30 int32 = -2147483648
                            v10 -= (-2.25 * 0.1)
                            switch "" {
                            case "hello":
                                v9 = v3
                            default:
                                switch 4294967295 {
                                case 782:
                                    var v31 string = (("foo bar" + ("\n\t" + "\n\t")) + "a")
                                    _ = v31
                                case 444:
                                    var v32 int8 = (v9 & v9)
                                    _ = v32
                                case -2147483648:
                                    v9 = (v9 << ((1) & 63))
                                }
                            }
                            if (!((fi19 >> ((0) & 63)) == int(v30))) {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                v16 += v16
                            } else {
                                fmt.Println(fi19, v14)
                                v33 := v1[2147483647]
                                _ = v33
                            }
                            _ = v30
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v14)
                            default:
                                delete(v1, 1)
                            }
                        }
                        if false {
                            v10 = 0.0
                        } else {
                            if true {
                                v34 := GenS4{F0: 1, F1: 567, F2: 1}
                                v35 := GenS5{F0: 1}
                                _ = v34
                                _ = v35
                            }
                            fmt.Println(v16, v13)
                        }
                        switch "" {
                        case "\n\t":
                            if ((v11 / ((v11) | 1)) != 42) {
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(v16, v10, v3)
                                case "foo bar":
                                    if (!((1 < 42) && (42 < 524))) {
                                        if (!((9223372036854775807 >= -3) && true)) {
                                            var v36 int32 = -1
                                            v37 := GenS6{F0: true, F1: (^(v13)), F2: true}
                                            _ = v36
                                            _ = v37
                                        }
                                        var v38 int16 = 42
                                        _ = v38
                                    }
                                default:
                                    v12 = append(v12, fi15)
                                }
                                switch (^(fi15)) {
                                case 122:
                                    var v39 float64 = (((float64((v10 + v14)) * v2) / v2) * v2)
                                    _ = v39
                                case 2147483647:
                                    v2 -= (v2 * (float64(v14) + v2))
                                case 9223372036854775807:
                                    var v40 uint16 = (v13 & 2)
                                    _ = v40
                                default:
                                    fmt.Println(v14)
                                }
                            }
                        case "a":
                            delete(v1, 146)
                        }
                    }
                    _ = fi19
                }
                delete(v1, 2147483647)
                var v41 int64 = 42
                _ = v41
            }
        }
        _ = fi15
        _ = v16
    }
    v11 -= 87
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
