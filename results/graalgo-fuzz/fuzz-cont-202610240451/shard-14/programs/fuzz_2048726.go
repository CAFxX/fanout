package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
func genFn0() string {
    var v6 int8 = 127
    var v7 int8 = -114
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 1
    var v3 int = v2
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 4294967295
    v3 = (v3 >> ((42) & 63))
    var v8 int64 = 464
    v9 := []int{v3}
    v10 := GenS1{F0: v5, F1: -128, F2: 0}
    if (false && false) {
        v4[v5] = (v5 + v5)
    }
    if (496 < 149) {
        fmt.Println(v3, v5, v8)
    } else {
        var v11 uint16 = 32767
        var v12 int8 = 42
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            switch (fi14 >> ((0) & 63)) {
            case 9223372036854775807:
                v9 = append(v9, -2147483648)
            }
            _ = fi14
        }
        v3 = 32767
        if true {
            v2 += (^(fi13))
        } else {
            switch "\n\t" {
            case "foo bar":
                if (-39 >= 42) {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    v3 += fi13
                } else {
                    if false {
                        switch fi13 {
                        case 1:
                            v8 += 0
                        case 128:
                            var v15 int32 = 2147483647
                            _ = v15
                        case 9223372036854775807:
                            v5 = v5
                        }
                    }
                    fmt.Println(v8, fi13, fi13)
                    for fi16 := 0; fi16 < 15; fi16++ {
                        if true {
                            var v17 int16 = 32767
                            var v18 uint = 18446744073709551615
                            var v19 uint8 = 42
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        switch v2 {
                        case -697:
                            v20 := GenS2{F0: 2, F1: 1, F2: 1}
                            _ = v20
                        case 813:
                            if (true && true) {
                                var v21 int32 = 1
                                fmt.Println(v8, v2, v3)
                                fmt.Println(fi16)
                                _ = v21
                            }
                        default:
                            v22 := GenS1{F0: v5, F1: 0, F2: 725}
                            _ = v22
                        }
                        _ = fi16
                    }
                }
            case "hello":
                v2 -= fi13
            }
            fmt.Println(v3, v8)
            fmt.Println(fi13, fi13, v8)
        }
        _ = fi13
    }
    if true {
        if true {
            for fi23 := 0; fi23 < 20; fi23++ {
                if (3.14159 != 0.0) {
                    v3 += -315
                    var v24 uint64 = 42
                    if len(v9) > 0 {
                        v9[0] = 2147483647
                    }
                    _ = v24
                }
                for fi25 := 0; fi25 < 1; fi25++ {
                    var v26 bool = true
                    v8 += (v8 << ((18446744073709551615) & 63))
                    v27 := GenS1{F0: v5, F1: 11, F2: 255}
                    v3 += (^(v2))
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = fi23
            }
        } else {
            var v28 int16 = -32768
            v8 -= (v8 - int64(v3))
            _ = v28
        }
    }
    fmt.Println(v8, v3)
    var v29 int16 = 32767
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 string = ("a" + "")
        for fi32 := 0; fi32 < 4; fi32++ {
            v2 = fi30
            _ = fi32
        }
        v33 := GenS2{F0: 255, F1: 127, F2: 42}
        _ = fi30
        _ = v31
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
