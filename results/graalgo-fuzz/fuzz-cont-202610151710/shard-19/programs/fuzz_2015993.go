package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint32, p1 int) int16 {
    v4 := GenS1{F0: -1}
    var v5 int32 = 325
    v5 = (v5 - (v5 * (-(v5))))
    if (!true) {
        for fi6 := 0; fi6 < 13; fi6++ {
            p1 = p1
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -32768
}

func main() {
    var v1 uint = 65535
    var v2 float64 = ((0.0 * 3.14159) + 0.1)
    var v3 float32 = 1.5
    v1 += v1
    var v7 int = -529
    v3 -= 3.14159
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS2{F0: (fi8 % (((fi8 - v7)) | 1))}
        if false {
            v10 := []int{-1, 32768}
            switch -876 {
            case 32767:
                fmt.Println(v2)
            }
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = v10
        } else {
            for fi11 := 0; fi11 < 14; fi11++ {
                switch "\n\t" {
                case "hello":
                    v12 := GenS2{F0: fi8}
                    v3 += (v3 + v3)
                    _ = v12
                case "\n\t":
                    if false {
                        var v13 uint64 = 0
                        if true {
                            v1 = (v1 << ((v1) & 63))
                            v3 -= float32((v2 + v2))
                            if false {
                                var v14 int8 = -100
                                _ = v14
                            } else {
                                var v15 string = "\n\t"
                                if (127 == -128) {
                                    v16 := []int{v7, (^(fi11)), (fi8 ^ 128)}
                                    v3 -= (v3 / 0.1)
                                    v17 := GenS3{F0: 0}
                                    var v18 float64 = ((v2 / -2.25) / v2)
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                                _ = v15
                            }
                            v2 += float64((v3 / v3))
                        } else {
                            v13 = uint64(v1)
                        }
                        _ = v13
                    } else {
                        var v19 int16 = 808
                        var v20 string = ""
                        var v21 uint8 = 1
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
                var v22 int16 = genFn0(65535, -1)
                _ = fi11
                _ = v22
            }
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            fmt.Println(v1, v1, v7)
            v7 = fi8
            v7 = -1
            _ = fi23
        }
        _ = fi8
        _ = v9
    }
    v24 := []int{(v7 ^ 128)}
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 float64 = v2
        _ = fi25
        _ = v26
    }
    fmt.Println(v1, v2)
    _ = len(v24)
    if (v3 > v3) {
        v27 := GenS2{F0: v7}
        fmt.Println(v7, v2)
        switch "" {
        case "a":
            if false {
                v7 = (v7 ^ (v7 >> (((^(v1))) & 63)))
            } else {
                if len(v24) >= 2 {
                    v24 = v24[1:len(v24)]
                }
            }
        case "\n\t":
            v28 := GenS3{F0: 58}
            _ = v28
        case "foo bar":
            var v29 int64 = 9223372036854775807
            _ = v29
        default:
            if len(v24) > 0 {
                v24[0] = 541
            }
        }
        _ = v27
    }
    var v30 string = ""
    v2 -= (0.0 - (v2 - (float64((v3 * 0.0)) / float64(v3))))
    var v31 float64 = (v2 / (v2 / 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v7
}
