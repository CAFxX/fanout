package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 string }
type GenS2 struct { F0 int8; F1 bool; F2 uint }
type GenS3 struct { F0 bool; F1 int32; F2 int64 }

func main() {
    v1 := []int{}
    var v2 int32 = 42
    v3 := []uint32{}
    for fi4 := 0; fi4 < 18; fi4++ {
        switch "\n\t" {
        case "":
            if (616 != -468) {
                fmt.Println(fi4, v2)
                for fi5 := 0; fi5 < 1; fi5++ {
                    fmt.Println(fi4, fi5, fi4)
                    _ = fi5
                }
            } else {
                var v6 float32 = 1.5
                v7 := GenS1{F0: 1, F1: v2, F2: "\n\t"}
                _ = v6
                _ = v7
            }
        }
        var v8 int = fi4
        var v9 bool = true
        _ = fi4
        _ = v8
        _ = v9
    }
    fmt.Println(v2)
    _ = len(v3)
    v2 -= (v2 >> ((1) & 63))
    var v10 int16 = -655
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 10; fi12++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v10)
            if false {
                var v13 uint32 = 65535
                if false {
                    if true {
                        fmt.Println(v10, v2)
                        if true {
                            v10 = (v10 << ((255) & 63))
                            var v14 uint32 = (^(v13))
                            _ = v14
                        }
                        v10 += v10
                    } else {
                        v2 = v2
                    }
                    var v15 uint32 = (v13 - 2)
                    _ = v15
                }
                fmt.Println(fi12, v2)
                _ = v13
            }
            _ = fi12
        }
        _ = fi11
    }
    v16 := GenS2{F0: -128, F1: true, F2: 0}
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            v10 = (v10 | 1)
            _ = fi18
        }
        _ = fi17
    }
    if true {
        v19 := GenS2{F0: 93, F1: (0 <= 4294967295), F2: 1}
        fmt.Println(v10, v2)
        _ = v19
    } else {
        fmt.Println(v2)
        fmt.Println(v2, v10)
        v20 := GenS3{F0: true, F1: (v2 >> ((18446744073709551615) & 63)), F2: -9223372036854775808}
        _ = v20
    }
    v10 -= v10
    var v21 int32 = v2
    var v22 int64 = -9223372036854775808
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 uint8 = 229
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v25 int32 = int32(fi23)
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println(v10, v22)
    fmt.Println(v2, v21, v10)
    var v26 int64 = (^(v22))
    var v27 int16 = int16(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
}
