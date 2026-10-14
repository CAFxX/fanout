package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (1e10 + float64(float32(3.14159)))
    var v3 uint32 = 854
    v4 := []uint32{(v3 - (v3 * 0)), v3, 65535, 818}
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint = 121
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 0
    if (int64(v7) >= 9223372036854775807) {
        var v8 float64 = (v2 - v2)
        _ = v8
    } else {
        fmt.Println(v7)
    }
    if true {
        v3 -= (v3 / ((v3) | 1))
        var v9 uint8 = 42
        fmt.Println(v7)
        _ = v9
    }
    if true {
        switch 32767 {
        case -218:
            v10 := GenS1{F0: 42, F1: 1}
            _ = v10
        case -9223372036854775808:
            v3 -= 36
        case 32767:
            var v11 int8 = 42
            _ = v11
        }
    } else {
        v12 := GenS1{F0: 1, F1: -245}
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 uint32 = (v3 / (((v3 - (v3 / ((v3) | 1)))) | 1))
            v14 -= (v3 << ((886) & 63))
            _ = fi13
            _ = v14
        }
        v3 = v3
        _ = v12
    }
    switch "foo bar" {
    case "a":
        v15 := GenS1{F0: -1, F1: -811}
        _ = v15
    case "hello":
        if true {
            v3 += v3
        } else {
            for fi16 := 0; fi16 < 16; fi16++ {
                switch 127 {
                case -2147483648:
                    v17 := GenS1{F0: 2147483647, F1: -447}
                    _ = v17
                case 1:
                    v3 = v3
                case -128:
                    var v18 int = (fi16 & (fi16 % ((32767) | 1)))
                    _ = v18
                default:
                    v19 := GenS1{F0: 2147483647, F1: -32768}
                    _ = v19
                }
                _ = fi16
            }
        }
    case "":
        fmt.Println(v7)
    }
    var v20 uint8 = 255
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint8 = v20
        _ = fi21
        _ = v22
    }
    v2 += v2
    for fi23 := 0; fi23 < 1; fi23++ {
        _ = len(v4)
        fmt.Println(v7)
        for fi24 := 0; fi24 < 16; fi24++ {
            if true {
                fmt.Println(v7, fi24, v2)
                var v25 string = ""
                if true {
                    var v26 float32 = ((((0.1 / 1.0) * 3.14159) + -2.25) + -2.25)
                    if true {
                        v27 := GenS2{F0: 2147483647, F1: 276}
                        _ = v27
                    } else {
                        fmt.Println(v26)
                        fmt.Println(v25, v2)
                        fmt.Println(fi23, fi24)
                    }
                    _ = v26
                }
                fmt.Println(fi23)
                v28 := v1[-2147483648]
                _ = v25
                _ = v28
            } else {
                if (!(!false)) {
                    var v29 int32 = -732
                    var v30 int32 = v29
                    v31 := GenS1{F0: (v30 >> ((459) & 63)), F1: 32767}
                    if (fi24 != -650) {
                        fmt.Println(v29, v2, fi24)
                    } else {
                        _ = len(v4)
                        var v32 uint64 = v7
                        _ = v32
                    }
                    v30 -= (v30 | (v29 & v29))
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    var v33 int64 = 42
                    v20 = (v20 / ((v20) | 1))
                    delete(v1, fi23)
                    _ = v33
                }
                var v34 int16 = 32767
                _ = v34
            }
            _ = fi24
        }
        _ = fi23
    }
    v7 = uint64(v20)
    v35 := GenS2{F0: -1, F1: 255}
    var v36 string = (("" + "\n\t") + "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v7
}
