package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0() int16 {
    var v4 int64 = -9223372036854775808
    if (!(!(true && false))) {
        v5 := []int{-2147483648, 32767, -9223372036854775808, -563}
        if (32767 <= 32767) {
            for fi6 := 0; fi6 < 7; fi6++ {
                var v7 uint8 = 23
                var v8 uint64 = 42
                v7 = 2
                _ = fi6
                _ = v7
                _ = v8
            }
        } else {
            var v9 float32 = 3.14159
            for fi10 := 0; fi10 < 19; fi10++ {
                for fi11 := 0; fi11 < 8; fi11++ {
                    if true {
                        v4 = v4
                        if true {
                            if len(v5) > 0 {
                                v5[0] = fi10
                            }
                            fmt.Println(v4, v9)
                            v9 = (v9 * 0.1)
                            v5 = append(v5, 4294967295)
                        }
                    }
                    v12 := GenS1{F0: 778, F1: 2}
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
            fmt.Println(v4)
            _ = v9
        }
        var v13 int32 = -1
        _ = v13
        _ = v5
    } else {
        fmt.Println(v4)
    }
    fmt.Println(v4)
    _ = v4
    return 1000
}

func main() {
    var v1 float32 = (0.0 + (3.14159 * float32(1e10)))
    v2 := []int{2147483647, -57}
    var v3 string = ("foo bar" + "")
    var v14 int64 = 1
    if len(v2) > 0 {
        v2[0] = 127
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        if ((-587 > 0) && (true || (((-(v14)) < 42) && true))) {
            var v16 uint64 = 42
            var v17 uint64 = (v16 << ((4294967295) & 63))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v16
            _ = v17
        } else {
            v18 := GenS1{F0: 288, F1: 65535}
            _ = v18
        }
        if len(v2) > 0 {
            v2[0] = (fi15 / (((fi15 | fi15)) | 1))
        }
        _ = len(v2)
        _ = fi15
    }
    v3 = v3
    fmt.Println(v1)
    var v19 float64 = 0.1
    fmt.Println(v3, v19, v3)
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := GenS1{F0: (fi20 & (fi20 + fi20)), F1: 65535}
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: 42, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v22
    _ = v3
}
