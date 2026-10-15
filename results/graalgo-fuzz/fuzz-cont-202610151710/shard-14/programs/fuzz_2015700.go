package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 bool }

func main() {
    v1 := []string{(("hello" + (((("foo bar" + (("a" + "foo bar") + "\n\t")) + "a") + "foo bar") + ("" + ""))) + ("foo bar" + "hello"))}
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 36
    var v4 int = -9223372036854775808
    v5 := make(map[int]string)
    var v6 string = "a"
    var v7 int16 = int16(v4)
    v4 += v4
    if (!((536 > 11) || (!false))) {
        fmt.Println(v3, v6, v6)
        if true {
            if (!(!true)) {
                v8 := GenS1{F0: 4294967295, F1: (v7 & (v7 >> ((753) & 63))), F2: 255}
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 int8 = 1
                    delete(v5, (fi9 | fi9))
                    if ((false || ((!true) || true)) || (18446744073709551615 >= 18446744073709551615)) {
                        v6 = ("a" + v6)
                    }
                    _ = fi9
                    _ = v10
                }
                v11 := GenS1{F0: v3, F1: -32768, F2: 65535}
                _ = v11
                _ = v8
            }
            fmt.Println(v7, v6)
            var v12 uint64 = 4294967295
            _ = v12
        } else {
            v6 = ((v6 + "") + v6)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v13 := GenS2{F0: -9223372036854775808}
            _ = v13
        }
        fmt.Println(v6, v3)
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        if false {
            v7 = (v7 & (v7 - (v7 % (((v7 << ((957) & 63))) | 1))))
        } else {
            var v15 string = v6
            _ = v15
        }
        v3 = v3
        _ = fi14
    }
    var v16 uint32 = (v3 * (v3 ^ v3))
    if ((1.5 != 3.14159) && true) {
        _ = len(v5)
        var v17 int32 = -2147483648
        _ = v17
    }
    if (!true) {
        var v18 int = v4
        _ = v18
    }
    if (v3 < 4294967295) {
        v3 += (v16 << ((613) & 63))
        if true {
            if false {
                _ = len(v1)
                if (-110 > -1) {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        fmt.Println(v3)
                        _ = fi19
                    }
                    v20 := GenS1{F0: 2147483648, F1: v7, F2: 0}
                    v21 := GenS2{F0: 42}
                    _ = v20
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 10; fi22++ {
                v6 = ("a" + "a")
                v23 := GenS3{F0: -1, F1: 212}
                _ = fi22
                _ = v23
            }
        }
        var v24 uint = 1
        _ = v24
    }
    if false {
        fmt.Println(v3, v7)
        var v25 int16 = v7
        if (127 <= 42) {
            v26 := v2[(v3 / ((v16) | 1))]
            if (!true) {
                v7 -= v7
                if (true || false) {
                    v27 := GenS4{F0: false}
                    var v28 float64 = ((1e10 * float64((float32(-2.25) - 1e10))) * 0.0)
                    _ = v27
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 20; fi29++ {
                        if true {
                            fmt.Println(v25, v6, fi29)
                        }
                        if (fi29 == 980) {
                            var v30 float64 = (float64((1.5 + (float32(1.5) / float32(float64(1e10))))) + 1e10)
                            if true {
                                var v31 string = "\n\t"
                                v32 := GenS3{F0: 100, F1: v3}
                                _ = v31
                                _ = v32
                            } else {
                                v33 := GenS1{F0: (v3 / (((v16 + v3)) | 1)), F1: (v7 / ((v7) | 1)), F2: 255}
                                var v34 uint16 = 0
                                _ = v33
                                _ = v34
                            }
                            _ = v30
                        } else {
                            for fi35 := 0; fi35 < 11; fi35++ {
                                v36 := GenS3{F0: 1, F1: v3}
                                v37 := v5[(fi35 >> ((891) & 63))]
                                fmt.Println(v6)
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        }
                        fmt.Println(v7, v25)
                        _ = fi29
                    }
                    fmt.Println(v4)
                    var v38 string = v6
                    _ = v38
                }
                fmt.Println(v16, v16)
            }
            if false {
                for fi39 := 0; fi39 < 19; fi39++ {
                    v3 -= v26
                    v40 := GenS3{F0: 127, F1: (v3 / (((v16 / ((0) | 1))) | 1))}
                    var v41 int64 = 498
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                var v42 string = "a"
                for fi43 := 0; fi43 < 7; fi43++ {
                    v4 = (v4 / ((fi43) | 1))
                    v1 = append(v1, v6)
                    var v44 uint = 18446744073709551615
                    _ = fi43
                    _ = v44
                }
                _ = v42
            }
            _ = v26
        }
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
