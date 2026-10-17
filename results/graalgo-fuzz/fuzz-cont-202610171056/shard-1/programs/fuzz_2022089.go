package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 uint32 = 751
    v2 := make(map[string]string)
    var v3 int64 = -902
    v4 := GenS1{F0: v1, F1: 1}
    _ = len(v2)
    v5 := GenS1{F0: 45, F1: 127}
    v6 := GenS2{F0: 0, F1: 1, F2: 2}
    if ((float32(0.1) / ((0.0 + 1e10) - float32((((1e10 * 1e10) / float64(1.5)) * 0.0)))) <= 1.5) {
        v1 = 2
        fmt.Println(v1, v1)
    }
    v7 := GenS1{F0: (^(v1)), F1: 59}
    v8 := []int{4294967295}
    var v9 uint32 = 2147483648
    if len(v8) > 0 {
        v8[0] = 4294967295
    }
    if false {
        switch -482 {
        case -917:
            fmt.Println(v9, v1, v1)
        }
        var v10 int8 = 100
        switch 0 {
        case 32767:
            v9 -= (v1 - v9)
        case 682:
            switch 565 {
            case 705:
                for fi11 := 0; fi11 < 10; fi11++ {
                    v9 += 2147483647
                    _ = fi11
                }
            }
        case 4294967295:
            v9 = 65535
        }
        _ = v10
    } else {
        var v12 float32 = -2.25
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        v9 = (v9 | 1)
        _ = fi13
    }
    switch -1 {
    case 460:
        var v14 uint16 = 0
        _ = v14
    case 42:
        if true {
            if true {
                if false {
                    fmt.Println(v3)
                } else {
                    var v15 bool = (137 != 42)
                    if len(v8) > 0 {
                        v8[0] = -1
                    }
                    _ = v15
                }
                if ((597 >= -1) || (-1 <= 32767)) {
                    var v16 int8 = -115
                    v17 := GenS3{F0: (v9 ^ v1), F1: v1}
                    fmt.Println(v3, v1)
                    _ = v16
                    _ = v17
                }
            } else {
                if true {
                    if false {
                        v9 += (v1 & (v1 | v9))
                        var v18 string = "a"
                        for fi19 := 0; fi19 < 7; fi19++ {
                            v9 = (v1 << ((18446744073709551615) & 63))
                            if (v1 == (v9 + (^(v1)))) {
                                v9 = v1
                            }
                            _ = fi19
                        }
                        _ = v18
                    } else {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            for fi21 := 0; fi21 < 10; fi21++ {
                                var v22 uint = 0
                                if len(v8) > 0 {
                                    v8[0] = (-(fi20))
                                }
                                _ = fi21
                                _ = v22
                            }
                            v1 = v9
                            _ = fi20
                        }
                    }
                } else {
                    for fi23 := 0; fi23 < 15; fi23++ {
                        fmt.Println(v1, v1)
                        _ = fi23
                    }
                    switch 963 {
                    case 4294967295:
                        v1 -= 2147483648
                    }
                    v9 += 2
                }
                _ = len(v2)
                v24 := GenS1{F0: (v1 ^ v9), F1: 1}
                _ = v24
            }
            _ = len(v8)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
