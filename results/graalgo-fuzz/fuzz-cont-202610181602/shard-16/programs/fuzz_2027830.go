package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint64; F2 int }
type GenS3 struct { F0 int64 }

func main() {
    var v1 string = ("foo bar" + ("hello" + "a"))
    v2 := []uint32{755, 267, 1}
    var v3 int64 = -1
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: 0}
        _ = fi4
        _ = v5
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        v3 += (v3 - v3)
        fmt.Println(fi6)
        _ = fi6
    }
    v3 = 42
    var v7 uint8 = 127
    var v8 float64 = (((((0.0 + 1e10) - 1e10) + 0.1) + 0.1) - 1e10)
    fmt.Println(v8, v3, v3)
    switch 447 {
    case 32768:
        if len(v2) > 0 {
            v2[0] = 96
        }
    case -854:
        v1 = ((v1 + "") + ("foo bar" + (v1 + ("\n\t" + (v1 + (v1 + v1))))))
    case -2147483648:
        var v9 uint = 255
        _ = v9
    }
    switch "" {
    case "foo bar":
        v10 := GenS1{F0: 0}
        _ = v10
    default:
        var v11 uint8 = v7
        _ = v11
    }
    switch "" {
    case "":
        if (2 != v7) {
            v8 += v8
            for fi12 := 0; fi12 < 6; fi12++ {
                if true {
                    switch ("" + "a") {
                    case "\n\t":
                        var v13 int16 = 1000
                        _ = v13
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v14 int32 = 2147483647
                    v1 = "hello"
                    _ = v14
                } else {
                    var v15 uint64 = 42
                    fmt.Println(v15, fi12, v15)
                    fmt.Println(v7, v15)
                    _ = v15
                }
                _ = fi12
            }
            if (882 > 484) {
                v8 = v8
            }
        }
    case "hello":
        switch 42 {
        case -9223372036854775808:
            fmt.Println(v7, v8)
        case 127:
            var v16 float32 = (float32((v8 - (-2.25 - v8))) * 1e10)
            _ = v16
        case -1:
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 float64 = v8
                if false {
                    if false {
                        var v19 string = v1
                        fmt.Println(v8, v1, v18)
                        _ = v19
                    }
                    var v20 uint8 = (v7 % ((v7) | 1))
                    switch ((((v1 + "hello") + v1) + (v1 + v1)) + ("\n\t" + "foo bar")) {
                    case "hello":
                        var v21 uint64 = 1
                        _ = v21
                    case "\n\t":
                        switch v1 {
                        case "hello":
                            v22 := GenS1{F0: 323}
                            _ = v22
                        }
                    default:
                        if false {
                            var v23 string = ""
                            _ = v23
                        }
                    }
                    _ = v20
                }
                if ((v18 * (float64(0.0) + (v8 + (0.1 * ((((((1.5 - v18) * 1.5) + (((float64(1e10) / float64(float32(1e10))) / float64(1.5)) + 1e10)) + (float64(1.5) - v8)) * v8) - v18))))) > (float64(3.14159) - 0.1)) {
                    var v24 int16 = 1000
                    var v25 int = 553
                    _ = v24
                    _ = v25
                }
                _ = fi17
                _ = v18
            }
        default:
            if (false || (true && true)) {
                fmt.Println(v7, v1)
                if true {
                    for fi26 := 0; fi26 < 15; fi26++ {
                        v8 = 1e10
                        v27 := GenS1{F0: 0}
                        v8 += float64(1.5)
                        _ = fi26
                        _ = v27
                    }
                    fmt.Println(v3, v1)
                }
            } else {
                fmt.Println(v8, v8)
                var v28 string = ""
                if true {
                    var v29 int8 = 64
                    var v30 uint64 = 0
                    _ = v29
                    _ = v30
                } else {
                    _ = len(v2)
                    switch 0 {
                    case 4294967295:
                        v28 = v1
                    case 9223372036854775807:
                        v31 := GenS2{F0: 500, F1: 1, F2: -50}
                        _ = v31
                    }
                }
                _ = v28
            }
        }
    case "a":
        var v32 int64 = -1
        _ = v32
    }
    v33 := GenS3{F0: v3}
    var v34 float32 = 3.14159
    var v35 int = 788
    var v36 int32 = -411
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v7
    _ = v8
}
