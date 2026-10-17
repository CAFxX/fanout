package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint }
type GenS2 struct { F0 int8; F1 uint8; F2 uint }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint8; F1 int16 }

func main() {
    var v1 float32 = (-2.25 - ((3.14159 / -2.25) + (1.5 + float32(1e10))))
    v2 := make(map[int]bool)
    var v3 uint8 = 255
    v3 += 255
    if false {
        var v4 float32 = v1
        v5 := GenS1{F0: 687, F1: -128, F2: 42}
        fmt.Println(v1, v1)
        _ = v4
        _ = v5
    } else {
        for fi6 := 0; fi6 < 15; fi6++ {
            v1 = ((((v1 - (-2.25 - 0.0)) + ((1.5 - 0.0) / v1)) * (v1 - v1)) / 1.5)
            if (4294967295 > 1) {
                v7 := GenS1{F0: 9223372036854775807, F1: 42, F2: 65535}
                v8 := []int{(fi6 | 128), fi6}
                _ = v7
                _ = v8
            }
            v2[fi6] = true
            _ = fi6
        }
    }
    if true {
        v9 := []int{2147483647, 0, -2147483648, 127}
        var v10 string = ("a" + "")
        fmt.Println(v1, v3)
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 42, F1: (v3 ^ v3), F2: 725}
    var v12 int16 = 1000
    fmt.Println(v1, v3, v12)
    if (2 >= v3) {
        v12 += 1
    }
    var v13 bool = (!true)
    switch "\n\t" {
    case "foo bar":
        v14 := GenS1{F0: -9223372036854775808, F1: 42, F2: 18446744073709551615}
        _ = v14
    case "\n\t":
        var v15 bool = (127 != 913)
        _ = v15
    case "":
        var v16 uint = 338
        _ = v16
    }
    var v17 int64 = 1
    for fi18 := 0; fi18 < 10; fi18++ {
        switch fi18 {
        case 66:
            switch fi18 {
            case -1:
                var v19 float32 = 0.1
                _ = v19
            case 845:
                var v20 bool = v13
                _ = v20
            default:
                v1 = (v1 - 1e10)
            }
        case 127:
            if true {
                v21 := GenS3{F0: v12, F1: -935}
                _ = v21
            } else {
                if (115 >= v3) {
                    fmt.Println(fi18, fi18, v12)
                }
                if (3.14159 != (v1 * v1)) {
                    fmt.Println(v13, v17)
                    if (128 >= (v3 / ((v3) | 1))) {
                        for fi22 := 0; fi22 < 3; fi22++ {
                            switch (fi18 ^ (fi18 % ((-827) | 1))) {
                            case 9223372036854775807:
                                v23 := GenS4{F0: (v3 % ((v3) | 1)), F1: v12}
                                _ = v23
                            case -1:
                                if (v13 || false) {
                                    var v24 int32 = 1
                                    _ = v24
                                }
                            }
                            v13 = ((v12 / (((v12 + (v12 / (((v12 + 0)) | 1)))) | 1)) >= v12)
                            fmt.Println(fi22)
                            _ = fi22
                        }
                    } else {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            if ((255 != (v3 << ((255) & 63))) && v13) {
                                if ((true || (!v13)) || v13) {
                                    var v26 int64 = (-(v17))
                                    v27 := GenS1{F0: v17, F1: -128, F2: 4294967295}
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 float32 = v1
                                    if (-2.25 < 1.5) {
                                        v29 := []int{}
                                        v30 := GenS2{F0: -1, F1: v3, F2: 174}
                                        _ = v29
                                        _ = v30
                                    }
                                    v3 = (v3 + 42)
                                    _ = v28
                                }
                                if v13 {
                                    v13 = v13
                                    if (0 <= 1) {
                                        v1 = 1.5
                                        var v31 uint = 0
                                        _ = v31
                                    }
                                } else {
                                    v32 := []int{(-(fi25)), fi18, fi25}
                                    delete(v2, (fi18 ^ (fi25 % ((fi18) | 1))))
                                    _ = v32
                                }
                            }
                            var v33 bool = (v13 || v13)
                            var v34 uint64 = 0
                            _ = fi25
                            _ = v33
                            _ = v34
                        }
                        v35 := GenS2{F0: 127, F1: 0, F2: 881}
                        _ = v35
                    }
                } else {
                    var v36 int64 = (^(v17))
                    _ = v36
                }
                v37 := GenS3{F0: 1, F1: v17}
                _ = v37
            }
        case 691:
            for fi38 := 0; fi38 < 15; fi38++ {
                var v39 uint = 334
                v40 := []int{fi18}
                if len(v40) >= 2 {
                    v40 = v40[1:len(v40)]
                }
                _ = fi38
                _ = v39
                _ = v40
            }
        default:
            var v41 bool = (v1 == 3.14159)
            _ = v41
        }
        _ = fi18
    }
    var v42 uint64 = 0
    fmt.Println(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v42
}
