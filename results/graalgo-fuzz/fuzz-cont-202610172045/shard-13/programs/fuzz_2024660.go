package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }
type GenS2 struct { F0 uint32; F1 bool; F2 bool }
func genFn0(p0 int32, p1 uint64) uint {
    var v7 int8 = -49
    v8 := []int{-259}
    v7 = 127
    var v9 int16 = 42
    var v10 uint16 = uint16(p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 267
}

func main() {
    v1 := []int{-185, 42, 463, 0}
    var v2 uint64 = 546
    var v3 float64 = (3.14159 / float64((((1e10 * 3.14159) - 0.0) * float32(0.1))))
    var v4 uint16 = 92
    v5 := make(map[int]string)
    var v6 int16 = 450
    var v11 uint32 = 65535
    v12 := GenS1{F0: "\n\t", F1: 830, F2: -944}
    if true {
        delete(v5, 1)
        fmt.Println(v3, v11, v2)
    } else {
        var v13 int32 = -1
        _ = v13
    }
    if false {
        v4 = (v4 + (v4 + v4))
        if true {
            v14 := GenS2{F0: v11, F1: (32767 >= -2147483648), F2: true}
            _ = v14
        } else {
            v2 += v2
        }
    }
    v6 -= 32767
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 float32 = 0.0
        switch -2147483648 {
        case 32767:
            if true {
                v16 = v16
                var v17 int = 646
                _ = v17
            } else {
                if ((1 <= (^(v6))) && (v6 <= v6)) {
                    var v18 uint64 = v2
                    if true {
                        if true {
                            for fi19 := 0; fi19 < 5; fi19++ {
                                var v20 int = (fi19 << ((0) & 63))
                                _ = fi19
                                _ = v20
                            }
                            v21 := GenS2{F0: 65535, F1: false, F2: (476 <= 4294967295)}
                            v4 = v4
                            _ = v21
                        }
                        fmt.Println(v11, v11, v16)
                        fmt.Println(v18, v2)
                    }
                    if (334 != -271) {
                        fmt.Println(v2, v3)
                        var v22 bool = true
                        _ = v22
                    }
                    _ = v18
                } else {
                    var v23 uint32 = v11
                    _ = v23
                }
                v16 += (3.14159 * v16)
            }
        }
        v24 := GenS2{F0: (v11 + (v11 / (((v11 - v11)) | 1))), F1: false, F2: true}
        _ = fi15
        _ = v16
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v1 = append(v1, (fi25 ^ int(v11)))
        if false {
            fmt.Println(v3, v6)
            for fi26 := 0; fi26 < 5; fi26++ {
                v1 = append(v1, fi26)
                v27 := GenS2{F0: uint32(v4), F1: (true && false), F2: (2147483648 >= 0)}
                _ = fi26
                _ = v27
            }
        }
        _ = fi25
    }
    v11 = v11
    v11 = v11
    _ = len(v1)
    fmt.Println(v4, v3, v4)
    v28 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: 4294967295}
    for fi29 := 0; fi29 < 18; fi29++ {
        v4 = 42
        _ = fi29
    }
    if false {
        v5[-2147483648] = "a"
    }
    fmt.Println(v4, v4)
    v11 = (v11 / ((15) | 1))
    var v30 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
