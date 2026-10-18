package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int16 }

func main() {
    var v1 bool = (42 >= 701)
    var v2 uint = 65535
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    v1 = ((v2 - (v2 % (((v2 | 65535)) | 1))) < (^(v2)))
    v1 = (-264 == 0)
    v1 = (false || (1 != 42))
    delete(v3, 637)
    v5 := []int{127, -1, 32767, 2147483647}
    v6 := GenS1{F0: -2147483648, F1: 1, F2: 531}
    if ((-2.25 / 3.14159) != -2.25) {
        v7 := GenS1{F0: 42, F1: 1, F2: 283}
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(fi8, v1, fi8)
            _ = fi8
        }
        v2 = v2
        _ = v7
    }
    v1 = v1
    v9 := GenS1{F0: 649, F1: -143, F2: -32768}
    if (v1 && (1.5 >= 3.14159)) {
        if v1 {
            switch (("hello" + (((((("" + "hello") + ("a" + "foo bar")) + "") + "a") + "hello") + "")) + "hello") {
            case "hello":
                for fi10 := 0; fi10 < 2; fi10++ {
                    var v11 uint8 = 166
                    _ = len(v3)
                    switch (fi10 * fi10) {
                    case 42:
                        if v1 {
                            fmt.Println(fi10, v11, v2)
                        } else {
                            if true {
                                if (v1 || false) {
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(v2)
                                    default:
                                        v12 := v3[128]
                                        _ = v12
                                    }
                                }
                                fmt.Println(v11, fi10)
                                v11 = v11
                            } else {
                                if (1 <= 255) {
                                    _ = len(v4)
                                    var v13 uint64 = uint64(v2)
                                    v2 = 1
                                    _ = v13
                                } else {
                                    if v1 {
                                        var v14 uint64 = 601
                                        var v15 float32 = 1.5
                                        _ = v14
                                        _ = v15
                                    } else {
                                        v2 += (^(v2))
                                        v11 = uint8(v2)
                                        v4[(fi10 | 127)] = (!(!(((!(-47 != 5)) && v1) || v1)))
                                    }
                                }
                                v2 = (v2 ^ (v2 % ((0) | 1)))
                                var v16 float64 = 0.1
                                _ = v16
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case 128:
                        if v1 {
                            var v17 int = (-(fi10))
                            v17 += (fi10 * (v17 >> ((v2) & 63)))
                            _ = v17
                        }
                    case -2147483648:
                        var v18 uint16 = 2
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 3; fi19++ {
                            if true {
                                v2 += v2
                            }
                            v2 = uint(v11)
                            _ = fi19
                        }
                    }
                    _ = fi10
                    _ = v11
                }
            case "":
                v3[766] = (478 != 88)
            }
        }
        fmt.Println(v2)
    }
    v1 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
