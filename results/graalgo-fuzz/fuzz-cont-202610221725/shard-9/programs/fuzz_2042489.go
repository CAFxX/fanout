package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 string; F1 uint8; F2 int16 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 1
    var v3 int32 = 2147483647
    v4 := []int{-2147483648, -2147483648, 128, 128}
    v5 := make(map[int]uint32)
    var v6 uint32 = 2147483648
    fmt.Println(v3, v1)
    v7 := GenS1{F0: 80}
    var v8 int = -9223372036854775808
    fmt.Println(v8)
    v9 := GenS2{F0: 1, F1: (v8 >> ((4294967295) & 63))}
    v6 = 0
    if (3.14159 >= ((0.0 / (0.1 * 3.14159)) * 0.0)) {
        if (2 == 119) {
            v3 = (^(v3))
            fmt.Println(v2)
        } else {
            v10 := GenS2{F0: 497, F1: (v8 ^ v8)}
            if false {
                if false {
                    if false {
                        for fi11 := 0; fi11 < 17; fi11++ {
                            var v12 uint = 65535
                            v8 -= 2147483647
                            _ = fi11
                            _ = v12
                        }
                    }
                    fmt.Println(v2, v3, v2)
                    v13 := GenS2{F0: 460, F1: v8}
                    _ = v13
                } else {
                    v1 = (v1 | 19)
                }
            } else {
                v14 := GenS3{F0: 32767, F1: "a"}
                _ = v14
            }
            _ = v10
        }
        var v15 float32 = 0.0
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v16 int16 = -1
            _ = v16
        }
        _ = v15
    } else {
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(v3)
            _ = fi17
        }
        if (!true) {
            switch v8 {
            case -9223372036854775808:
                v3 += v3
            }
        } else {
            for fi18 := 0; fi18 < 7; fi18++ {
                v8 = 724
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 int16 = 1000
                    _ = fi19
                    _ = v20
                }
                var v21 int = (v8 - -2147483648)
                _ = fi18
                _ = v21
            }
        }
        if (true || (v3 == -2147483648)) {
            for fi22 := 0; fi22 < 12; fi22++ {
                for fi23 := 0; fi23 < 10; fi23++ {
                    v6 = v6
                    _ = fi23
                }
                for fi24 := 0; fi24 < 10; fi24++ {
                    var v25 int32 = 0
                    v26 := GenS4{F0: 2, F1: true}
                    var v27 int32 = (v3 >> ((255) & 63))
                    var v28 uint16 = 474
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi22
            }
            v5[(v8 & v8)] = 42
        }
    }
    var v29 uint = 65535
    v30 := GenS4{F0: 0, F1: (!(v1 != uint64(v2)))}
    fmt.Println(v2)
    var v31 int16 = 806
    v29 += v29
    v8 = -466
    fmt.Println(v2, v3, v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch -786 {
    case 513:
        v32 := GenS1{F0: -34}
        _ = v32
    default:
        if true {
            v33 := GenS5{F0: "a", F1: 42, F2: 832}
            var v34 uint32 = 2
            v5[(v8 | (v8 ^ -584))] = 4294967295
            _ = v33
            _ = v34
        } else {
            v29 = v29
            var v35 uint16 = v2
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
