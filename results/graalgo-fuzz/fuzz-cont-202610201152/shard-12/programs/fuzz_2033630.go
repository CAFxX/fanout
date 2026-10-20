package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := []int{32768}
    var v4 float64 = ((0.1 + 1.5) - 3.14159)
    var v5 uint32 = 2147483647
    var v6 uint16 = 785
    if (int8(v5) != 42) {
        fmt.Println(v6, v6, v6)
    } else {
        var v7 int16 = -545
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: v5, F1: (^(v6))}
            _ = fi8
            _ = v9
        }
        if true {
            v4 -= 3.14159
            if false {
                for fi10 := 0; fi10 < 3; fi10++ {
                    if ((^(v5)) >= 2147483647) {
                        fmt.Println(fi10, v6, v5)
                        v4 = (float64(1.5) / 1e10)
                    }
                    _ = len(v2)
                    if true {
                        delete(v2, (v5 << ((115) & 63)))
                    } else {
                        var v11 int64 = 0
                        _ = v11
                    }
                    _ = fi10
                }
                fmt.Println(v4, v4, v7)
                if (false && false) {
                    _ = len(v1)
                    for fi12 := 0; fi12 < 2; fi12++ {
                        switch (fi12 >> ((0) & 63)) {
                        case -91:
                            for fi13 := 0; fi13 < 2; fi13++ {
                                v14 := GenS2{F0: (v5 | (v5 * 42)), F1: 221}
                                var v15 bool = true
                                var v16 int = fi12
                                _ = fi13
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                        }
                        v4 += v4
                        v1[v5] = v5
                        _ = fi12
                    }
                }
            }
        } else {
            var v17 bool = true
            _ = v17
        }
        _ = v7
    }
    v1[v5] = (v5 * 762)
    if true {
        if true {
            var v18 int64 = -9223372036854775808
            _ = v18
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch ("a" + "a") {
        case "hello":
            if len(v3) > 0 {
                v3[0] = -2147483648
            }
        }
    } else {
        delete(v1, (v5 & 1))
        fmt.Println(v4, v5)
    }
    var v19 uint8 = 127
    for fi20 := 0; fi20 < 1; fi20++ {
        switch "\n\t" {
        case "a":
            if true {
                fmt.Println(fi20, v6)
            }
        case "foo bar":
            v19 = (v19 ^ v19)
        default:
            v4 += v4
        }
        if len(v3) > 0 {
            v3[0] = fi20
        }
        fmt.Println(v4)
        _ = fi20
    }
    v21 := GenS3{F0: 2, F1: 4294967295}
    if (!(true && true)) {
        var v22 int16 = 1
        if true {
            for fi23 := 0; fi23 < 18; fi23++ {
                fmt.Println(v22, v19)
                _ = fi23
            }
            var v24 int = -2147483648
            _ = v24
        } else {
            var v25 bool = (v6 <= 32767)
            fmt.Println(v22, v22, v5)
            _ = v25
        }
        _ = v22
    }
    var v26 int8 = 0
    v27 := GenS3{F0: 4294967295, F1: 1}
    var v28 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
