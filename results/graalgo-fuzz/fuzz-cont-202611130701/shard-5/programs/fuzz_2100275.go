package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int16; F1 int16; F2 int16 }

func main() {
    var v1 string = "hello"
    var v2 uint = 819
    v3 := make(map[uint32]bool)
    v1 = (("hello" + v1) + "foo bar")
    v1 = v1
    var v4 float32 = 0.0
    v5 := []int{}
    fmt.Println(v2, v2)
    for fi6 := 0; fi6 < 11; fi6++ {
        if (628 < uint32(v2)) {
            var v7 int = 127
            if false {
                if true {
                    for fi8 := 0; fi8 < 7; fi8++ {
                        delete(v3, 60)
                        delete(v3, 222)
                        _ = fi8
                    }
                } else {
                    fmt.Println(v2)
                    var v9 float64 = 1.5
                    for fi10 := 0; fi10 < 17; fi10++ {
                        var v11 int = fi6
                        var v12 string = "hello"
                        fmt.Println(fi6)
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = v9
                }
                v13 := GenS1{F0: -1}
                fmt.Println(v1, v4, v4)
                _ = v13
            } else {
                if (v1 > "a") {
                    v14 := GenS1{F0: int16(v2)}
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 uint8 = 1
                        _ = fi15
                        _ = v16
                    }
                }
                v1 = ((v1 + v1) + (v1 + "hello"))
                fmt.Println(v7, v7, v1)
            }
            _ = v7
        }
        _ = fi6
    }
    v2 += v2
    if (!(!(18446744073709551615 != 4294967295))) {
        v5 = append(v5, 9223372036854775807)
        for fi17 := 0; fi17 < 4; fi17++ {
            if (true || (true && true)) {
                var v18 bool = true
                for fi19 := 0; fi19 < 16; fi19++ {
                    fmt.Println(v1)
                    switch fi17 {
                    case 946:
                        fmt.Println(v2, v2, fi19)
                    default:
                        var v20 int = fi19
                        _ = v20
                    }
                    var v21 int = -9223372036854775808
                    _ = fi19
                    _ = v21
                }
                var v22 string = v1
                _ = v18
                _ = v22
            } else {
                fmt.Println(fi17)
            }
            v2 = (v2 % (((v2 ^ 558)) | 1))
            _ = fi17
        }
    } else {
        v23 := GenS2{F0: (^(v2)), F1: 330}
        if (!true) {
            var v24 uint64 = 314
            for fi25 := 0; fi25 < 19; fi25++ {
                v3[2147483648] = (true && false)
                _ = fi25
            }
            var v26 int16 = 0
            _ = v24
            _ = v26
        }
        _ = v23
    }
    v5 = append(v5, 516)
    var v27 float64 = ((((3.14159 - (0.1 + 1e10)) * float64((float32(float64(v4)) + v4))) / 3.14159) * 3.14159)
    v4 -= float32(v27)
    var v28 uint32 = 2
    if false {
        var v29 int8 = -81
        if (1 > uint64(v28)) {
            for fi30 := 0; fi30 < 9; fi30++ {
                v5 = append(v5, (fi30 << (((v2 >> ((4294967295) & 63))) & 63)))
                _ = fi30
            }
        }
        _ = v29
    } else {
        v2 = (v2 + (v2 & 4294967295))
        v31 := GenS1{F0: 32767}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v31
    }
    if false {
        v32 := GenS3{F0: -32768, F1: -833, F2: -32768}
        var v33 float32 = ((v4 + v4) * 1e10)
        _ = v32
        _ = v33
    } else {
        delete(v3, v28)
        if (42 > -930) {
            switch 2147483647 {
            case 0:
                v2 += (v2 ^ (^(v2)))
            case -853:
                var v34 int = 128
                _ = v34
            default:
                var v35 uint16 = 255
                _ = v35
            }
        } else {
            var v36 int = -9223372036854775808
            v36 += v36
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
