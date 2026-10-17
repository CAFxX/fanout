package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint64; F1 int16; F2 int32 }
func genFn0(p0 uint, p1 bool) int8 {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 float32 = 0.0
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int64 = -9223372036854775808
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v9 := GenS1{F0: 128, F1: 42, F2: 2}
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    var v11 int16 = 1000
    v11 += v11
    var v12 uint16 = 255
    if ((2 < 87) || false) {
        if false {
            v13 := GenS1{F0: 255, F1: 1, F2: 2147483648}
            fmt.Println(v11, v12)
            fmt.Println(v11, v11)
            _ = v13
        } else {
            v14 := GenS1{F0: 1, F1: 596, F2: 954}
            v11 = v11
            if true {
                v11 = v11
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 float32 = (1e10 - float32(3.14159))
                    var v17 int32 = -2147483648
                    if false {
                        var v18 uint32 = 65535
                        _ = v18
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                for fi19 := 0; fi19 < 5; fi19++ {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        if ((v11 % ((-1) | 1)) <= v11) {
                            var v21 string = "\n\t"
                            _ = v21
                        } else {
                            var v22 int16 = v11
                            switch ("" + ((("\n\t" + "a") + "a") + "a")) {
                            case "foo bar":
                                if (1 == 4294967295) {
                                    v23 := GenS2{F0: 100, F1: -418}
                                    _ = v23
                                }
                            }
                            var v24 float32 = (0.1 * 0.0)
                            var v25 int = fi19
                            _ = v22
                            _ = v24
                            _ = v25
                        }
                        v11 = v11
                        _ = fi20
                    }
                    v26 := GenS3{F0: 1, F1: (v11 / (((v11 << ((uint(v12)) & 63))) | 1)), F2: 2147483647}
                    v11 += (v11 ^ v11)
                    _ = fi19
                    _ = v26
                }
            }
            _ = v14
        }
        v11 = v11
    } else {
        if false {
            v27 := []int{706, 128, -9223372036854775808, 32767}
            _ = v27
        }
    }
    if (956 > 4294967295) {
        var v28 uint32 = 42
        _ = v28
    }
    v29 := GenS1{F0: 0, F1: 65535, F2: 337}
    v30 := GenS3{F0: 2, F1: (-(v11)), F2: -928}
    v11 -= 0
    v12 -= 114
    v11 += v11
    var v31 int8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
