package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    var v4 uint64 = 2
    if (127 == 974) {
        var v5 int = 0
        _ = v5
    }
    if false {
        var v6 int8 = 127
        _ = v6
    } else {
        v4 += (v4 << ((18446744073709551615) & 63))
        if (!false) {
            if ((!true) || (0 < 0)) {
                switch "hello" {
                case "hello":
                    fmt.Println(v4)
                case "a":
                    v1[319] = 561
                case "":
                    for fi7 := 0; fi7 < 14; fi7++ {
                        v4 += (v4 + v4)
                        var v8 float32 = 3.14159
                        v9 := GenS1{F0: 42, F1: (fi7 / (((fi7 * (fi7 & 32767))) | 1)), F2: 2}
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                }
                v4 = (v4 * (v4 / ((v4) | 1)))
                var v10 uint64 = (^(v4))
                _ = v10
            } else {
                for fi11 := 0; fi11 < 1; fi11++ {
                    v4 += 42
                    var v12 int64 = 9223372036854775807
                    _ = fi11
                    _ = v12
                }
                v4 -= (v4 ^ 528)
                v13 := GenS1{F0: 18446744073709551615, F1: 880, F2: 340}
                _ = v13
            }
            for fi14 := 0; fi14 < 16; fi14++ {
                v4 += v4
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := GenS1{F0: 2, F1: -756, F2: 42}
                    _ = fi15
                    _ = v16
                }
                var v17 uint = 4294967295
                _ = fi14
                _ = v17
            }
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS1{F0: 526, F1: fi18, F2: 607}
                v20 := v1[uint32(v4)]
                v21 := v1[(v20 & (v20 % ((v20) | 1)))]
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
    }
    v22 := GenS2{F0: 77}
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(fi23, fi23)
        v24 := GenS2{F0: 2147483647}
        _ = fi23
        _ = v24
    }
    var v25 int16 = 1
    if true {
        var v26 bool = (42 <= 2)
        v26 = v26
        if true {
            for fi27 := 0; fi27 < 8; fi27++ {
                if (42 <= 42) {
                    var v28 int32 = 1
                    v29 := GenS2{F0: 2}
                    _ = v28
                    _ = v29
                }
                if (!(v4 <= v4)) {
                    v26 = (!(v26 && (v4 != v4)))
                    if ((42 <= 40) && v26) {
                        fmt.Println(v26, v26, fi27)
                        var v30 uint8 = 1
                        _ = v30
                    }
                }
                v31 := GenS2{F0: 0}
                _ = fi27
                _ = v31
            }
            v4 -= (v4 * 18446744073709551615)
            v32 := GenS2{F0: 2147483648}
            _ = v32
        } else {
            var v33 uint16 = 530
            if len(v2) > 0 {
                v2[0] = 0
            }
            _ = len(v2)
            _ = v33
        }
        _ = v26
    } else {
        v25 -= v25
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "" {
        case "\n\t":
            fmt.Println(v4)
        }
    }
    v4 = (v4 << ((1) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
