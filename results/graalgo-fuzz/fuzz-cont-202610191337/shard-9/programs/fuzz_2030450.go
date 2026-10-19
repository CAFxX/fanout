package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int32 }

func main() {
    var v1 int64 = 517
    v2 := []string{("foo bar" + ("" + "")), "hello", "hello", (("a" + "foo bar") + "")}
    v3 := make(map[uint32]bool)
    var v4 string = "a"
    var v5 int8 = -43
    v4 = "hello"
    v2 = append(v2, v4)
    var v6 uint16 = 662
    v6 -= 987
    var v7 uint = 65535
    fmt.Println(v4, v7, v5)
    if false {
        v4 = v4
        v7 = (v7 + v7)
    } else {
        if (v5 >= int8(v7)) {
            if false {
                for fi8 := 0; fi8 < 6; fi8++ {
                    _ = len(v2)
                    switch v4 {
                    case "\n\t":
                        switch (fi8 | fi8) {
                        case 127:
                            switch fi8 {
                            case 32768:
                                if ((1 <= 1) || (fi8 < (fi8 >> ((18446744073709551615) & 63)))) {
                                    switch v4 {
                                    case "":
                                        if false {
                                            v7 = v7
                                            delete(v3, 87)
                                            v9 := GenS1{F0: (fi8 / ((fi8) | 1)), F1: (v6 | (v6 ^ v6)), F2: 1000}
                                            _ = v9
                                        } else {
                                            if (3.14159 != 0.1) {
                                                v10 := GenS2{F0: 42}
                                                _ = v10
                                            }
                                            v7 -= v7
                                            if true {
                                                for fi11 := 0; fi11 < 18; fi11++ {
                                                    v1 += (v1 - v1)
                                                    v6 = 2
                                                    var v12 uint8 = 239
                                                    var v13 uint = (v7 + v7)
                                                    fmt.Println(v5, v6)
                                                    _ = fi11
                                                    _ = v12
                                                    _ = v13
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                            } else {
                                                v1 = 0
                                                var v14 int8 = v5
                                                _ = v14
                                            }
                                        }
                                    case "foo bar":
                                        v2 = append(v2, ((("foo bar" + "hello") + (v4 + v4)) + ((v4 + v4) + v4)))
                                    }
                                }
                            case -1:
                                var v15 float64 = ((float64(1e10) * 3.14159) - 0.0)
                                _ = v15
                            }
                        case 0:
                            v1 = (v1 | v1)
                        case -1:
                            var v16 int8 = 127
                            _ = v16
                        }
                    case "foo bar":
                        if (((("" != (("\n\t" + ("hello" + ((("\n\t" + "") + ("foo bar" + v4)) + "a"))) + v4)) && false) || (4294967295 != 1)) || ((((0.1 - (1.5 + float64(float32(float64((0.1 - 3.14159)))))) / 1e10) + (0.1 / 1.0)) <= 1e10)) {
                            var v17 string = "foo bar"
                            for fi18 := 0; fi18 < 18; fi18++ {
                                fmt.Println(fi8, v17)
                                v19 := GenS3{F0: 32767, F1: (v6 + v6), F2: (v6 / ((uint16(v7)) | 1))}
                                var v20 uint8 = 42
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            v6 += (v6 % ((uint16(v7)) | 1))
                            _ = v17
                        } else {
                            v21 := GenS4{F0: (v6 / ((255) | 1)), F1: 2}
                            v7 -= v7
                            v22 := GenS5{F0: 144, F1: 1}
                            _ = v21
                            _ = v22
                        }
                    case "":
                        v4 = (v4 + (v4 + v4))
                    }
                    _ = fi8
                }
            }
        } else {
            if len(v2) > 0 {
                v2[0] = (v4 + "\n\t")
            }
            v7 = 18446744073709551615
        }
        v23 := GenS3{F0: int16(v5), F1: v6, F2: 255}
        v4 = v4
        _ = v23
    }
    v24 := GenS5{F0: 1, F1: 42}
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 uint64 = 1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v27 bool = true
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
