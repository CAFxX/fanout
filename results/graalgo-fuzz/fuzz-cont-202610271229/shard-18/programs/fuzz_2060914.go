package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int16; F2 uint }
func genFn0(p0 uint, p1 uint16) int {
    p1 = p1
    p0 = p0
    fmt.Println(p0, p1)
    var v5 int = 127
    _ = p0
    _ = p1
    _ = v5
    return v5
}
func genFn1(p0 string, p1 int8, p2 int8) uint8 {
    v6 := []int{617, 475}
    switch 1 {
    case 32768:
        v6 = append(v6, 323)
    case -9223372036854775808:
        v7 := GenS1{F0: (^(p1))}
        _ = v7
    case 4294967295:
        p0 = (p0 + (p0 + "\n\t"))
    }
    switch p0 {
    case "foo bar":
        var v8 int16 = 1
        _ = v8
    case "\n\t":
        if (127 >= 116) {
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint16 = 32767
                if (true || true) {
                    fmt.Println(v10, v10, p1)
                    switch (fi9 ^ 32768) {
                    case 32768:
                        var v11 uint32 = 176
                        _ = v11
                    case 0:
                        if true {
                            fmt.Println(p2, v10, p1)
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    default:
                        for fi12 := 0; fi12 < 20; fi12++ {
                            var v13 bool = ((-43 == 1000) || (!(!false)))
                            _ = fi12
                            _ = v13
                        }
                    }
                    var v14 uint8 = 128
                    _ = v14
                } else {
                    var v15 bool = (true || (("foo bar" + p0) == p0))
                    var v16 uint8 = 57
                    _ = v15
                    _ = v16
                }
                _ = fi9
                _ = v10
            }
            for fi17 := 0; fi17 < 11; fi17++ {
                for fi18 := 0; fi18 < 14; fi18++ {
                    var v19 int16 = 715
                    var v20 float64 = (float64(1.5) - 0.1)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                p1 = 42
                _ = fi17
            }
        } else {
            switch p0 {
            case "hello":
                fmt.Println(p0)
            default:
                v21 := GenS1{F0: -1}
                _ = v21
            }
            v22 := GenS2{F0: 1, F1: 2}
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    v1 := []int{-705, 1}
    var v2 bool = (0.0 != ((float32(float64(0.0)) / 1e10) + float32(0.0)))
    v3 := []uint32{2147483648, 65535, 421, 4294967295}
    var v4 int = 170
    fmt.Println(v4, v4)
    _ = len(v3)
    v2 = ((v4 + 2147483647) < genFn0(95, 88))
    var v23 string = "\n\t"
    v24 := GenS3{F0: -32768}
    var v25 int = 0
    v1 = append(v1, v25)
    var v26 uint16 = 32767
    var v27 int64 = -1
    v28 := GenS4{F0: int8(v25), F1: 1000, F2: 539}
    v4 = (^(v25))
    var v29 float64 = 1.5
    if v2 {
        v25 += v25
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 int32 = 42
            if ((v25 <= (v4 * (fi30 << ((818) & 63)))) || v2) {
                var v32 int = 42
                for fi33 := 0; fi33 < 18; fi33++ {
                    var v34 uint16 = v26
                    switch -2147483648 {
                    case -9223372036854775808:
                        v35 := GenS3{F0: -32768}
                        _ = v35
                    case -807:
                        var v36 float32 = 1e10
                        _ = v36
                    }
                    _ = fi33
                    _ = v34
                }
                if len(v3) > 0 {
                    v3[0] = 97
                }
                _ = v32
            } else {
                if ((fi30 % (((fi30 % ((387) | 1))) | 1)) < -2147483648) {
                    var v37 bool = (true || true)
                    _ = v37
                } else {
                    if v2 {
                        v26 = 32767
                        fmt.Println(v31, v31)
                        v1 = append(v1, int(v26))
                    }
                    _ = len(v3)
                }
                _ = len(v3)
            }
            _ = fi30
            _ = v31
        }
        for fi38 := 0; fi38 < 20; fi38++ {
            fmt.Println(v4)
            var v39 int16 = 42
            _ = fi38
            _ = v39
        }
    }
    fmt.Println(v26, v23)
    var v40 float64 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v40
}
