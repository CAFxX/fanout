package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 string; F1 int8 }
func genFn0(p0 int64) uint16 {
    switch 32768 {
    case -1:
        v5 := []int{-2147483648, -318}
        _ = v5
    case 32768:
        v6 := GenS1{F0: 18446744073709551615, F1: 139}
        _ = v6
    }
    p0 = p0
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{9223372036854775807, -9223372036854775808, fi7}
        _ = fi7
        _ = v8
    }
    p0 = 0
    _ = p0
    return 42
}

func main() {
    v1 := []uint32{842, 4294967295, 0, 4294967295}
    v2 := make(map[uint32]uint32)
    v3 := []uint32{}
    var v4 int8 = 1
    if true {
        var v9 bool = (!(1 < 127))
        _ = v9
    }
    _ = len(v2)
    var v10 float32 = 1e10
    var v11 uint64 = 4294967295
    var v12 uint8 = 127
    var v13 int32 = -1
    _ = len(v2)
    var v14 int16 = -1
    v15 := GenS1{F0: v11, F1: -448}
    _ = len(v1)
    v16 := GenS1{F0: (v11 / ((42) | 1)), F1: 32767}
    var v17 uint32 = 4294967295
    v11 = 661
    var v18 int16 = 1
    if true {
        if true {
            v19 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808}
            var v20 uint8 = v12
            _ = v19
            _ = v20
        } else {
            if ((false && (false && (65535 < 1))) || true) {
                if true {
                    for fi21 := 0; fi21 < 18; fi21++ {
                        fmt.Println(v12, v12)
                        _ = fi21
                    }
                    v22 := GenS1{F0: v11, F1: 1}
                    var v23 uint8 = (^(v12))
                    _ = v22
                    _ = v23
                }
                v13 = (v13 & v13)
                _ = len(v1)
            } else {
                var v24 int32 = v13
                v4 -= 0
                v11 += (^(v11))
                _ = v24
            }
        }
    } else {
        v14 -= (v14 % ((v14) | 1))
        var v25 float32 = ((float32(0.1) - v10) + float32(((1.5 + 1.5) * -2.25)))
        v26 := v2[(v17 % (((v17 | 803)) | 1))]
        _ = v25
        _ = v26
    }
    switch "" {
    case "hello":
        for fi27 := 0; fi27 < 3; fi27++ {
            fmt.Println(v17, v18, fi27)
            var v28 int = (fi27 + fi27)
            v29 := GenS1{F0: (^(v11)), F1: v28}
            _ = fi27
            _ = v28
            _ = v29
        }
    case "\n\t":
        switch "foo bar" {
        case "":
            v30 := GenS2{F0: false, F1: int64(v14), F2: 2}
            _ = v30
        case "\n\t":
            var v31 uint16 = 2
            _ = v31
        case "foo bar":
            fmt.Println(v13, v11, v10)
        default:
            var v32 int16 = v18
            _ = v32
        }
    case "foo bar":
        var v33 uint32 = (v17 * (v17 + v17))
        _ = v33
    default:
        if false {
            v18 += (v14 - v18)
            v34 := v2[2147483648]
            var v35 uint32 = 4294967295
            _ = v34
            _ = v35
        } else {
            if false {
                if false {
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                }
                for fi36 := 0; fi36 < 18; fi36++ {
                    if (!(611 >= 1)) {
                        var v37 uint32 = v17
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 10; fi38++ {
                        v12 = (v12 * v12)
                        _ = fi38
                    }
                    v13 = (^(v13))
                    _ = fi36
                }
            } else {
                var v39 int = 0
                v40 := GenS3{F0: 482, F1: v14}
                switch 128 {
                case -2147483648:
                    var v41 string = ("foo bar" + "a")
                    _ = v41
                case 572:
                    if false {
                        v42 := GenS4{F0: (^(v12)), F1: true}
                        v10 = (((((v10 - (v10 - v10)) * (v10 * ((v10 + v10) * 0.0))) / float32(3.14159)) / -2.25) * 1e10)
                        _ = v42
                    } else {
                        switch "a" {
                        case "\n\t":
                            var v43 float32 = (0.1 / -2.25)
                            _ = v43
                        default:
                            switch (v39 & v39) {
                            case 128:
                                if false {
                                    v44 := GenS5{F0: "foo bar", F1: 15}
                                    _ = v44
                                } else {
                                    for fi45 := 0; fi45 < 7; fi45++ {
                                        v46 := GenS5{F0: "foo bar", F1: v4}
                                        _ = fi45
                                        _ = v46
                                    }
                                    v11 = (v11 << ((189) & 63))
                                    var v47 float64 = (0.0 * (1e10 - -2.25))
                                    _ = v47
                                }
                            case 4294967295:
                                if (false || ((!(3.14159 < 0.0)) && false)) {
                                    v1 = append(v1, v17)
                                } else {
                                    var v48 uint8 = v12
                                    _ = v48
                                }
                            default:
                                delete(v2, (v17 * (v17 ^ v17)))
                            }
                        }
                        var v49 int64 = -1
                        _ = v49
                    }
                case 4294967295:
                    v4 -= v4
                default:
                    var v50 int8 = (v4 * 42)
                    _ = v50
                }
                _ = v39
                _ = v40
            }
            var v51 uint32 = 2147483648
            _ = v51
        }
    }
    switch int(v17) {
    case 267:
        var v52 uint32 = v17
        _ = v52
    case -2147483648:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -1:
        v10 -= 0.0
    default:
        if false {
            var v53 uint16 = 0
            var v54 bool = (2147483647 <= 0)
            _ = v53
            _ = v54
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
