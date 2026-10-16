package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int8; F2 int16 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 bool; F1 bool }

func main() {
    var v1 uint16 = 79
    var v2 int64 = 9223372036854775807
    var v3 int32 = 1
    v4 := make(map[int]int)
    switch 663 {
    case 32767:
        fmt.Println(v3)
    case 1:
        var v5 uint32 = 2147483647
        _ = v5
    case 42:
        v3 -= (v3 & (v3 ^ v3))
    }
    if false {
        var v6 uint32 = 42
        _ = v6
    }
    v1 = (^(v1))
    fmt.Println(v3, v1, v1)
    v7 := GenS1{F0: (v3 * v3), F1: 4294967295, F2: 2147483648}
    switch ("hello" + "hello") {
    case "foo bar":
        var v8 int = 32767
        _ = v8
    }
    v2 = v2
    fmt.Println(v2, v3, v1)
    v1 = v1
    if ((-551 >= 326) && (!true)) {
        if (0 > 18446744073709551615) {
            fmt.Println(v2)
            if ((-2.25 / 0.1) != 0.1) {
                for fi9 := 0; fi9 < 7; fi9++ {
                    v10 := GenS1{F0: 2147483647, F1: 127, F2: 196}
                    _ = fi9
                    _ = v10
                }
                v11 := []int{2147483647, -1, 4294967295, 606}
                v3 = v3
                _ = v11
            }
            var v12 int64 = (v2 + v2)
            _ = v12
        } else {
            fmt.Println(v2, v2)
        }
        if ((234 > 218) && (!false)) {
            var v13 int16 = -1
            v2 = v2
            for fi14 := 0; fi14 < 5; fi14++ {
                v2 = v2
                _ = fi14
            }
            _ = v13
        } else {
            fmt.Println(v2, v3)
        }
        v3 = v3
    } else {
        v2 = (^(v2))
        var v15 float32 = 0.0
        _ = v15
    }
    switch "\n\t" {
    case "foo bar":
        if (true || (v1 >= v1)) {
            v16 := []int{-128, 9223372036854775807, 2147483647}
            _ = v16
        } else {
            v2 = -318
            v2 += (v2 & 9223372036854775807)
            v17 := GenS2{F0: 127}
            _ = v17
        }
    default:
        v3 += (v3 / ((2147483647) | 1))
    }
    v4[32768] = 32768
    v2 = (^(v2))
    if false {
        v18 := GenS1{F0: 2147483647, F1: 18446744073709551615, F2: 0}
        _ = v18
    }
    v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
    for fi19 := 0; fi19 < 7; fi19++ {
        if false {
            v3 = -1
            v20 := GenS1{F0: (v3 + 0), F1: 18446744073709551615, F2: 2}
            for fi21 := 0; fi21 < 2; fi21++ {
                if (false || false) {
                    if (831 > (v1 >> ((42) & 63))) {
                        if (!true) {
                            switch (fi21 >> ((847) & 63)) {
                            case 823:
                                var v22 string = ("a" + "")
                                _ = v22
                            default:
                                fmt.Println(v1, fi21)
                            }
                            switch "foo bar" {
                            case "":
                                v3 = (v3 % (((v3 ^ v3)) | 1))
                            case "hello":
                                var v23 string = (("" + ("foo bar" + "")) + "a")
                                _ = v23
                            default:
                                fmt.Println(v3, fi21, v3)
                            }
                            v2 += (v2 % (((-(v2))) | 1))
                        } else {
                            var v24 bool = false
                            _ = v24
                        }
                        v25 := []int{(fi19 ^ 32767), 32767, fi21}
                        v3 -= (-(v3))
                        v26 := GenS1{F0: (v3 | v3), F1: 42, F2: uint32(v1)}
                        v1 = (v1 % ((255) | 1))
                        var v27 uint16 = v1
                        v1 = 255
                        fmt.Println(v2, fi19, v27)
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    var v28 int = 42
                    v29 := []int{1, v28}
                    v30 := GenS3{F0: (v2 - (v2 - (v2 * (v2 >> ((0) & 63)))))}
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    v31 := GenS4{F0: 207}
                    _ = v31
                }
                var v32 string = "hello"
                if (true || true) {
                    v1 = (^(v1))
                    var v33 int32 = v3
                    v34 := GenS4{F0: (v2 - v2)}
                    _ = v33
                    _ = v34
                } else {
                    var v35 int32 = v3
                    if (float32((1.5 * -2.25)) < 0.1) {
                        v36 := GenS5{F0: 127, F1: 127, F2: 32767}
                        switch 2147483647 {
                        case 1:
                            fmt.Println(fi21)
                        case 9223372036854775807:
                            v37 := GenS5{F0: 75, F1: 13, F2: -74}
                            _ = v37
                        default:
                            switch (fi19 >> ((2) & 63)) {
                            case 0:
                                fmt.Println(v2, v3)
                            case 127:
                                var v38 uint32 = 2147483648
                                _ = v38
                            default:
                                v39 := GenS6{F0: 0, F1: 42}
                                var v40 int64 = (v2 / (((-(v2))) | 1))
                                _ = v39
                                _ = v40
                            }
                        }
                        switch (fi21 << ((371) & 63)) {
                        case 296:
                            v41 := []int{(fi19 / ((-9223372036854775808) | 1)), int(v2)}
                            _ = v41
                        case -634:
                            v42 := GenS7{F0: (float64(-2.25) != 0.0), F1: (4294967295 < 18446744073709551615)}
                            _ = v42
                        case 4294967295:
                            v3 = -2147483648
                        default:
                            if (128 == 128) {
                                var v43 uint8 = 42
                                v44 := GenS5{F0: 100, F1: 127, F2: 1000}
                                if false {
                                    v1 = (v1 % ((v1) | 1))
                                }
                                v4[9223372036854775807] = fi21
                                _ = v43
                                _ = v44
                            }
                        }
                        fmt.Println(v1, v32)
                        switch fi21 {
                        case -9223372036854775808:
                            if (!true) {
                                v45 := GenS3{F0: -1}
                                var v46 int = (fi21 / ((fi19) | 1))
                                var v47 string = v32
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                        case 32768:
                            v48 := GenS7{F0: ((32767 < -833) || (3.14159 == 0.1)), F1: true}
                            _ = v48
                        case 0:
                            var v49 uint16 = (v1 ^ v1)
                            _ = v49
                        }
                        _ = v36
                    } else {
                        v2 -= v2
                        var v50 string = "hello"
                        v51 := []int{fi19}
                        _ = v50
                        _ = v51
                    }
                    _ = v35
                }
                _ = fi21
                _ = v32
            }
            _ = v20
        } else {
            v2 = v2
        }
        var v52 int16 = -1
        _ = fi19
        _ = v52
    }
    v1 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
