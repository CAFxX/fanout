package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0(p0 int64, p1 uint64) int32 {
    v7 := []int{}
    if (!(255 >= 127)) {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 uint16 = 42
    v3 := make(map[int]int)
    v4 := []uint32{539, 2147483648}
    v5 := []int{-165, 4294967295, 42}
    v6 := []string{"hello", "", ""}
    var v8 string = "a"
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint8 = 128
        _ = fi9
        _ = v10
    }
    var v11 float64 = 0.1
    v11 -= (v11 + (v11 + 0.0))
    var v12 int64 = -9223372036854775808
    if len(v4) > 0 {
        v4[0] = 42
    }
    v13 := GenS1{F0: v8, F1: (v2 ^ v2)}
    switch v8 {
    case "":
        switch ((v8 + v8) + "") {
        case "\n\t":
            var v14 int32 = genFn0((v1 << ((uint(v2)) & 63)), 785)
            _ = v14
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    default:
        delete(v3, -590)
    }
    var v15 int16 = 42
    var v16 string = ("\n\t" + ((v8 + v8) + v8))
    v16 = (v16 + v8)
    if true {
        _ = len(v4)
        var v17 string = ("a" + (v8 + v16))
        if (1 >= 4294967295) {
            if false {
                for fi18 := 0; fi18 < 12; fi18++ {
                    delete(v3, 4294967295)
                    v8 = (v8 + v17)
                    _ = fi18
                }
                v8 = ""
            }
            if true {
                var v19 bool = false
                v20 := GenS1{F0: "", F1: (^(v2))}
                _ = v19
                _ = v20
            } else {
                if false {
                    if false {
                        fmt.Println(v17)
                    } else {
                        v8 = "a"
                        v21 := GenS2{F0: 18446744073709551615, F1: 2, F2: 0}
                        switch 32768 {
                        case 128:
                            v11 = v11
                        default:
                            if (v2 >= (^(v2))) {
                                var v22 uint16 = 448
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v22
                            } else {
                                if true {
                                    fmt.Println(v15)
                                    v23 := GenS1{F0: "a", F1: 42}
                                    _ = v23
                                }
                                v24 := GenS3{F0: (v2 * (v2 & (v2 + (v2 * 65535)))), F1: (v12 / ((int64(v15)) | 1)), F2: -80}
                                var v25 uint32 = 1
                                _ = v24
                                _ = v25
                            }
                        }
                        _ = v21
                    }
                }
                var v26 uint8 = 170
                _ = v26
            }
            _ = len(v3)
        } else {
            v27 := GenS2{F0: 42, F1: 93, F2: 1}
            _ = v27
        }
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
