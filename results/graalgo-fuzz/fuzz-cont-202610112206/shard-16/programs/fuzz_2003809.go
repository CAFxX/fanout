package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int32 = 56
    var v2 string = "\n\t"
    v3 := make(map[int]int)
    var v4 uint16 = 0
    var v5 int32 = (v1 << ((65535) & 63))
    var v6 float32 = ((0.1 - float32(0.0)) - (1e10 + (0.1 * (0.0 - (float32(3.14159) + 1.5)))))
    v7 := GenS1{F0: (("foo bar" + v2) + v2)}
    for fi8 := 0; fi8 < 4; fi8++ {
        v6 -= float32((((float64(((1.5 + v6) - (v6 + v6))) / (0.1 / 0.1)) + 0.0) + (-2.25 / 0.1)))
        v3[42] = (fi8 + (fi8 >> ((42) & 63)))
        _ = fi8
    }
    if (!false) {
        var v9 int16 = 218
        _ = v9
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                if (255 <= 160) {
                    if (127 < 161) {
                        v2 = v2
                    } else {
                        v12 := []int{fi10, (fi10 / ((fi11) | 1)), (fi11 * (fi10 ^ (fi11 >> ((683) & 63))))}
                        _ = len(v12)
                        _ = v12
                    }
                    if (!("foo bar" == "")) {
                        if false {
                            fmt.Println(fi11, fi10, fi10)
                            if true {
                                var v13 string = v2
                                _ = v13
                            }
                        }
                        var v14 string = "a"
                        _ = v14
                    }
                    if true {
                        fmt.Println(v1, v4)
                    } else {
                        v5 -= (-(v5))
                        v15 := GenS1{F0: (((v2 + "a") + (v2 + v2)) + v2)}
                        delete(v3, (^(fi10)))
                        _ = v15
                    }
                } else {
                    var v16 uint = 2
                    v5 += v5
                    v6 = v6
                    _ = v16
                }
                fmt.Println(fi11, fi10)
                v17 := GenS2{F0: 42}
                _ = fi11
                _ = v17
            }
            fmt.Println(v6)
            switch "" {
            case "\n\t":
                var v18 int32 = 2147483647
                _ = v18
            }
            _ = fi10
        }
    }
    var v19 uint8 = 42
    v20 := []int{-586}
    if (32767 != (v4 & v4)) {
        switch 4294967295 {
        case -973:
            fmt.Println(v4, v4, v1)
        default:
            if (v6 != v6) {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
            }
        }
        v4 -= v4
    } else {
        v6 = v6
        switch ("a" + v2) {
        case "":
            var v21 string = v2
            _ = v21
        case "\n\t":
            v4 = v4
        case "foo bar":
            _ = len(v3)
        default:
            fmt.Println(v6)
        }
    }
    v22 := GenS2{F0: -83}
    v2 = "foo bar"
    if (4294967295 >= 2) {
        var v23 int8 = 127
        v24 := GenS3{F0: 2}
        switch 923 {
        case 457:
            v25 := GenS3{F0: uint(v4)}
            _ = v25
        case 1:
            if len(v20) > 0 {
                v20[0] = -106
            }
        case -9223372036854775808:
            var v26 uint16 = (v4 * (v4 % (((v4 * 255)) | 1)))
            _ = v26
        default:
            var v27 int64 = 714
            _ = v27
        }
        _ = v23
        _ = v24
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        v19 -= v19
        if false {
            for fi29 := 0; fi29 < 8; fi29++ {
                v30 := GenS3{F0: 0}
                v19 = (v19 | (v19 - (v19 + (^(v19)))))
                _ = fi29
                _ = v30
            }
            v1 += 42
            fmt.Println(v1)
        }
        v5 = v1
        _ = fi28
    }
    var v31 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
