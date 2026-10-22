package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{}
    var v3 int8 = 127
    v4 := []int{32768, -612, 669, 2147483647}
    v5 := make(map[uint32]int)
    var v6 string = "a"
    var v7 float64 = 1e10
    v1[2147483647] = -9223372036854775808
    var v8 uint64 = 1
    switch v6 {
    case "a":
        v9 := GenS1{F0: (v6 + v6), F1: (v8 << ((1) & 63)), F2: (^(v8))}
        _ = v9
    }
    v7 = 0.0
    if len(v4) > 0 {
        v4[0] = 32767
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 uint32 = 65535
        var v12 int = (fi10 % ((fi10) | 1))
        switch v12 {
        case 0:
            var v13 uint16 = uint16(v11)
            _ = v13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v7 = (0.0 * v7)
    v14 := GenS2{F0: 32767, F1: 1, F2: 42}
    if true {
        var v15 int16 = 1
        _ = v15
    } else {
        switch 544 {
        case -1:
            _ = len(v4)
        case 32767:
            for fi16 := 0; fi16 < 16; fi16++ {
                for fi17 := 0; fi17 < 16; fi17++ {
                    if (-124 == 42) {
                        if false {
                            if true {
                                v8 -= v8
                                var v18 int = (-(fi16))
                                _ = v18
                            }
                        }
                        switch fi16 {
                        case -193:
                            v3 -= v3
                        }
                    } else {
                        var v19 uint32 = 725
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v6 = "a"
                        v20 := v1[-402]
                        _ = v19
                        _ = v20
                    }
                    _ = fi17
                }
                var v21 uint16 = 42
                _ = fi16
                _ = v21
            }
        case 32768:
            v22 := GenS3{F0: -9223372036854775808}
            _ = v22
        }
        fmt.Println(v3, v3)
    }
    if ((!false) || true) {
        v5[334] = 42
    } else {
        fmt.Println(v8, v7, v6)
    }
    if false {
        v3 -= 102
        for fi23 := 0; fi23 < 11; fi23++ {
            fmt.Println(v7, v6)
            _ = fi23
        }
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        switch ("a" + v6) {
        case "":
            var v25 int16 = -32768
            _ = v25
        }
        var v26 string = v6
        var v27 int = -933
        _ = fi24
        _ = v26
        _ = v27
    }
    if (0 == 0) {
        v5[4294967295] = -9223372036854775808
        for fi28 := 0; fi28 < 18; fi28++ {
            v3 = v3
            _ = fi28
        }
    } else {
        v29 := GenS4{F0: v3, F1: 9}
        var v30 uint = 971
        _ = v29
        _ = v30
    }
    var v31 int64 = -9223372036854775808
    v1[310] = 9223372036854775807
    var v32 uint16 = 830
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
