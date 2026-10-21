package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint64; F1 int; F2 uint16 }

func main() {
    var v1 int16 = -740
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    var v4 float64 = 0.0
    var v5 int8 = 100
    var v6 int = -2147483648
    var v7 bool = (((v5 + (v5 ^ (v5 ^ (v5 / (((v5 ^ v5)) | 1))))) <= v5) || false)
    if v7 {
        var v8 string = "\n\t"
        v6 = v6
        _ = v8
    }
    v9 := GenS1{F0: 2147483648}
    fmt.Println(v7, v6)
    v10 := []int{v6, (v6 << ((1) & 63)), int(v1)}
    if false {
        fmt.Println(v5, v5)
        if len(v10) > 0 {
            v10[0] = 4294967295
        }
    } else {
        fmt.Println(v5)
        for fi11 := 0; fi11 < 14; fi11++ {
            for fi12 := 0; fi12 < 16; fi12++ {
                if v7 {
                    v1 -= v1
                } else {
                    var v13 uint = 18446744073709551615
                    switch (v6 & 852) {
                    case 697:
                        v14 := GenS2{F0: 18446744073709551615, F1: 1, F2: 244}
                        v1 = (^(v1))
                        _ = v14
                    }
                    _ = v13
                }
                v15 := GenS3{F0: 42, F1: v5}
                fmt.Println(v1, v1)
                if true {
                    fmt.Println(v1, v4, v6)
                    var v16 int64 = int64(fi11)
                    fmt.Println(v6, v16)
                    _ = v16
                } else {
                    v17 := v3[42]
                    switch "" {
                    case "foo bar":
                        var v18 int64 = -1
                        _ = v18
                    default:
                        if false {
                            v4 = 0.0
                            switch 2147483647 {
                            case 128:
                                v19 := GenS3{F0: -736, F1: (v5 & v5)}
                                _ = v19
                            case 127:
                                var v20 string = "foo bar"
                                _ = v20
                            }
                            if (255 != 32767) {
                                v21 := GenS2{F0: 1, F1: 9223372036854775807, F2: 42}
                                var v22 bool = v17
                                _ = v21
                                _ = v22
                            }
                        } else {
                            if v17 {
                                v4 = float64((1.5 * 1e10))
                                switch "hello" {
                                case "a":
                                    v23 := GenS4{F0: -808, F1: 767}
                                    _ = v23
                                case "":
                                    var v24 uint = 65535
                                    _ = v24
                                }
                            } else {
                                if v7 {
                                    var v25 uint = 2
                                    if v17 {
                                        fmt.Println(v1)
                                        fmt.Println(v6, v7, v5)
                                    }
                                    fmt.Println(v25, v1)
                                    _ = v25
                                }
                                v26 := GenS1{F0: 4294967295}
                                _ = v26
                            }
                        }
                    }
                    _ = v17
                }
                _ = fi12
                _ = v15
            }
            _ = fi11
        }
    }
    v4 += (1e10 / float64(((0.1 / (1e10 - float32(-2.25))) - 1e10)))
    v1 += v1
    if (v7 && (9223372036854775807 < -764)) {
        var v27 int8 = v5
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int32 = -2147483648
        var v30 int32 = v29
        _ = fi28
        _ = v29
        _ = v30
    }
    if v7 {
        v7 = v7
    } else {
        v10 = append(v10, v6)
        v31 := GenS5{F0: 18446744073709551615, F1: (v6 / ((-877) | 1)), F2: 832}
        for fi32 := 0; fi32 < 10; fi32++ {
            if v7 {
                for fi33 := 0; fi33 < 16; fi33++ {
                    v7 = v7
                    _ = fi33
                }
            }
            _ = fi32
        }
        _ = v31
    }
    if v7 {
        _ = len(v10)
        switch v6 {
        case 128:
            var v34 int32 = -2147483648
            _ = v34
        case -300:
            switch v6 {
            case 2147483647:
                v6 = v6
            case -1:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            default:
                fmt.Println(v1, v5, v4)
            }
        case 32768:
            var v35 int16 = (v1 ^ v1)
            _ = v35
        default:
            _ = len(v10)
        }
    }
    fmt.Println(v7, v1)
    var v36 string = "foo bar"
    _ = len(v10)
    var v37 int16 = -249
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
