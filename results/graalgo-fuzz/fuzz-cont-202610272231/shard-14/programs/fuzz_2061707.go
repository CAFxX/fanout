package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint32 }

func main() {
    var v1 int32 = 0
    var v2 string = "\n\t"
    var v3 int64 = 42
    v4 := []int{int(v3), -582, -1, 127}
    v5 := make(map[string]string)
    for fi6 := 0; fi6 < 13; fi6++ {
        v3 -= 0
        v7 := GenS1{F0: 246, F1: 0, F2: 42}
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 1
    for fi9 := 0; fi9 < 19; fi9++ {
        switch (v2 + "foo bar") {
        case "foo bar":
            v10 := GenS1{F0: 632, F1: v8, F2: 18446744073709551615}
            _ = v10
        case "\n\t":
            var v11 int64 = (v3 % ((int64(fi9)) | 1))
            _ = v11
        default:
            if ((true && (!(148 > 42))) || true) {
                v2 = (v2 + (("hello" + ("foo bar" + (v2 + v2))) + (v2 + (v2 + v2))))
                v1 = (^(v1))
            }
        }
        var v12 int64 = 42
        if true {
            if false {
                switch 127 {
                case 2147483647:
                    if true {
                        switch (fi9 >> ((255) & 63)) {
                        case -913:
                            if true {
                                if ((!(255 != 128)) && true) {
                                    if true {
                                        fmt.Println(fi9)
                                        var v13 uint8 = 127
                                        _ = v13
                                    } else {
                                        var v14 float64 = 1.5
                                        _ = v14
                                    }
                                    var v15 uint32 = (v8 * (v8 / (((v8 + 2147483648)) | 1)))
                                    _ = v15
                                }
                                v12 = v12
                            }
                        case 32767:
                            v16 := GenS1{F0: 496, F1: (^(v8)), F2: 2}
                            _ = v16
                        default:
                            fmt.Println(v1)
                        }
                    }
                case 43:
                    v17 := GenS1{F0: 18446744073709551615, F1: (v8 << ((438) & 63)), F2: 2}
                    _ = v17
                case 524:
                    var v18 bool = false
                    _ = v18
                }
                switch "" {
                case "foo bar":
                    for fi19 := 0; fi19 < 3; fi19++ {
                        v2 = v2
                        fmt.Println(v3, v3, v2)
                        var v20 float32 = float32(1.5)
                        v20 = (v20 * v20)
                        var v21 float64 = (3.14159 / -2.25)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                case "\n\t":
                    for fi22 := 0; fi22 < 8; fi22++ {
                        var v23 uint16 = 2
                        var v24 uint16 = 0
                        if true {
                            var v25 uint32 = 152
                            var v26 uint32 = 468
                            var v27 uint8 = 42
                            v3 = 9223372036854775807
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                case "hello":
                    fmt.Println(v2)
                default:
                    fmt.Println(fi9, v2)
                }
                v28 := GenS1{F0: 2, F1: 911, F2: 1}
                _ = v28
            } else {
                var v29 uint16 = 65535
                v30 := GenS1{F0: 18446744073709551615, F1: v8, F2: 0}
                var v31 int16 = -32768
                _ = v29
                _ = v30
                _ = v31
            }
            var v32 bool = true
            _ = v32
        } else {
            v33 := v5[(v2 + ("hello" + "hello"))]
            _ = v33
        }
        _ = fi9
        _ = v12
    }
    if (true || false) {
        fmt.Println(v3, v1)
        for fi34 := 0; fi34 < 15; fi34++ {
            v3 = 364
            _ = fi34
        }
    }
    v35 := GenS1{F0: 301, F1: (v8 ^ v8), F2: 2}
    v36 := GenS1{F0: 211, F1: v8, F2: 255}
    for fi37 := 0; fi37 < 1; fi37++ {
        for fi38 := 0; fi38 < 20; fi38++ {
            v2 = ((v2 + v2) + v2)
            v8 += v8
            _ = fi38
        }
        v39 := GenS2{F0: 18446744073709551615, F1: (v3 % (((v3 ^ v3)) | 1)), F2: 42}
        _ = fi37
        _ = v39
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v41 uint16 = 0
        v42 := GenS3{F0: 42, F1: -32768, F2: (v8 << ((255) & 63))}
        _ = fi40
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 13; fi43++ {
        if (0 <= 5) {
            fmt.Println(v3, v1)
            var v44 uint = 4294967295
            var v45 uint32 = (v8 << ((0) & 63))
            _ = v44
            _ = v45
        } else {
            v2 = (((v2 + v2) + "a") + v2)
            v5[v2] = ("foo bar" + v2)
            v2 = ("a" + (("hello" + v2) + (v2 + v2)))
        }
        var v46 int32 = v1
        _ = fi43
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v8
}
