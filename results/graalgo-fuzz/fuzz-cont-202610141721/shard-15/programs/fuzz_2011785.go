package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 bool }
type GenS4 struct { F0 uint; F1 int; F2 uint32 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{2, 0}
    var v3 string = ("a" + ("foo bar" + "a"))
    var v4 int64 = 42
    var v5 uint8 = 126
    var v6 int8 = 100
    v6 = (v6 << ((215) & 63))
    var v7 uint8 = (^(v5))
    if false {
        for fi8 := 0; fi8 < 13; fi8++ {
            if true {
                fmt.Println(v3)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v9 := GenS1{F0: 1, F1: -1, F2: (v5 >> ((2) & 63))}
            var v10 int = (fi8 % ((fi8) | 1))
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        v5 += v5
    }
    v11 := GenS2{F0: 565, F1: 0, F2: 0}
    v12 := GenS2{F0: 0, F1: 2147483648, F2: 18446744073709551615}
    v13 := GenS3{F0: 18446744073709551615, F1: (v5 << ((4294967295) & 63)), F2: (false || false)}
    var v14 uint = 2
    v15 := GenS2{F0: 1, F1: 450, F2: 42}
    var v16 int64 = (v4 >> ((991) & 63))
    for fi17 := 0; fi17 < 7; fi17++ {
        if (2 <= 65535) {
            var v18 uint = (v14 << (((v14 / (((v14 ^ v14)) | 1))) & 63))
            v2 = append(v2, 2147483647)
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS1{F0: 249, F1: (-(v6)), F2: (^(v7))}
                _ = len(v2)
                if (42 > 4294967295) {
                    v21 := GenS4{F0: 2, F1: 32768, F2: 2}
                    if true {
                        v1[-281] = (v3 >= "hello")
                        v5 = 1
                    } else {
                        v22 := GenS5{F0: (v7 << ((v14) & 63)), F1: 800}
                        v14 = v18
                        _ = v22
                    }
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            if false {
                v23 := GenS2{F0: 1, F1: 2147483647, F2: 42}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v4 = v16
                _ = v23
            } else {
                if false {
                    var v24 bool = true
                    _ = v24
                }
                for fi25 := 0; fi25 < 17; fi25++ {
                    if true {
                        if false {
                            var v26 int8 = 42
                            v4 = -9223372036854775808
                            switch (fi25 ^ (fi25 + fi17)) {
                            case -902:
                                var v27 string = "foo bar"
                                _ = v27
                            case 128:
                                v2 = append(v2, 409)
                            default:
                                var v28 int32 = 0
                                _ = v28
                            }
                            var v29 int64 = v4
                            var v30 uint16 = 0
                            _ = v26
                            _ = v29
                            _ = v30
                        } else {
                            _ = len(v2)
                        }
                        if true {
                            if true {
                                v31 := GenS6{F0: 0, F1: int16(fi25)}
                                var v32 uint8 = 127
                                _ = v31
                                _ = v32
                            }
                            if ((-2.25 - (1.5 * 3.14159)) > 3.14159) {
                                v33 := GenS2{F0: 523, F1: uint32(v7), F2: 18446744073709551615}
                                _ = v33
                            } else {
                                var v34 uint8 = (v7 ^ v7)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v34
                            }
                        }
                        if false {
                            v35 := v1[fi25]
                            _ = v35
                        } else {
                            v14 += (v14 * 0)
                            if true {
                                if ((v4 << ((v14) & 63)) < v4) {
                                    var v36 int16 = 1
                                    switch fi25 {
                                    case 128:
                                        v3 = v3
                                        v14 = (v14 / ((v14) | 1))
                                    }
                                    _ = v36
                                } else {
                                    v5 += 120
                                }
                            }
                        }
                    } else {
                        v5 = 42
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v14 += 255
                    _ = fi25
                }
                var v37 float32 = float32(float64((float32(-2.25) - (3.14159 - float32(0.1)))))
                _ = v37
            }
        }
        v38 := GenS1{F0: 139, F1: (v6 | v6), F2: (v5 << (((v14 << ((v14) & 63))) & 63))}
        _ = fi17
        _ = v38
    }
    v14 = (v14 % (((^(v14))) | 1))
    v2 = append(v2, 4294967295)
    v39 := GenS5{F0: 128, F1: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
