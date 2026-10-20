package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int8; F2 bool }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int16; F1 int64; F2 int32 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    v4 := []int{-9223372036854775808, -223, 9223372036854775807}
    for fi5 := 0; fi5 < 11; fi5++ {
        v1 = v1
        switch 32767 {
        case -9223372036854775808:
            var v6 float32 = 1e10
            _ = v6
        }
        if (!((!true) || (4294967295 > 65535))) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi7 := 0; fi7 < 19; fi7++ {
                v1 = "hello"
                if true {
                    fmt.Println(fi7)
                } else {
                    _ = len(v4)
                }
                var v8 int32 = -2147483648
                _ = fi7
                _ = v8
            }
        } else {
            if len(v4) > 0 {
                v4[0] = (fi5 << ((4294967295) & 63))
            }
            v9 := GenS1{F0: (fi5 % ((487) | 1))}
            fmt.Println(v1, fi5)
            _ = v9
        }
        _ = fi5
    }
    v1 = (v1 + "hello")
    fmt.Println(v1)
    var v10 uint32 = 0
    for fi11 := 0; fi11 < 2; fi11++ {
        if true {
            v1 = ""
            var v12 float32 = 3.14159
            _ = v12
        } else {
            v4 = append(v4, -127)
            _ = len(v2)
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 uint = 2
                switch "" {
                case "a":
                    v15 := GenS1{F0: (fi13 ^ 42)}
                    _ = v15
                }
                v10 = v10
                _ = fi13
                _ = v14
            }
        }
        if true {
            fmt.Println(fi11)
            v16 := GenS1{F0: -1}
            v10 = (v10 | v10)
            _ = v16
        } else {
            for fi17 := 0; fi17 < 8; fi17++ {
                if len(v4) > 0 {
                    v4[0] = fi17
                }
                v2[(v10 | 4294967295)] = fi11
                v2[v10] = fi11
                _ = fi17
            }
            var v18 int16 = 1000
            _ = v18
        }
        switch (fi11 % ((fi11) | 1)) {
        case 2147483647:
            var v19 string = (v1 + "\n\t")
            _ = v19
        case 128:
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 int16 = 32767
                v22 := GenS1{F0: fi11}
                _ = fi20
                _ = v21
                _ = v22
            }
        default:
            if (1 >= 458) {
                fmt.Println(fi11)
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 int32 = 467
                    if true {
                        switch "hello" {
                        case "":
                            var v25 float32 = ((1.5 * ((0.0 * (1.5 + float32(1e10))) / -2.25)) + 0.1)
                            _ = v25
                        default:
                            if (v1 >= "foo bar") {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v26 := GenS2{F0: 1, F1: -128, F2: (fi11 == fi11)}
                                _ = v26
                            } else {
                                v27 := GenS3{F0: 506, F1: fi11}
                                fmt.Println(v24, v1)
                                _ = v27
                            }
                        }
                    }
                    _ = fi23
                    _ = v24
                }
                var v28 uint64 = 2
                _ = v28
            } else {
                v10 += 365
            }
        }
        _ = fi11
    }
    var v29 int32 = 0
    delete(v3, -2147483648)
    if (-720 != 42) {
        var v30 uint = 2
        fmt.Println(v1, v10)
        var v31 uint16 = 42
        _ = v30
        _ = v31
    } else {
        if true {
            v32 := GenS4{F0: 1, F1: 2, F2: (v29 | (v29 + (v29 << ((2) & 63))))}
            _ = v32
        } else {
            v1 = v1
        }
    }
    v1 = ((v1 + (((v1 + v1) + v1) + v1)) + v1)
    fmt.Println(v29, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
