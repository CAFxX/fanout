package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int32; F1 uint64; F2 bool }

func main() {
    v1 := []uint32{2147483647, 4294967295, 42}
    var v2 int64 = -392
    var v3 float64 = (0.0 / -2.25)
    var v4 bool = true
    v5 := GenS1{F0: 18446744073709551615}
    v2 = v2
    switch 528 {
    case 127:
        if v4 {
            fmt.Println(v4, v2, v3)
            var v6 string = "a"
            var v7 int64 = -9223372036854775808
            _ = v6
            _ = v7
        }
    }
    var v8 uint32 = 728
    fmt.Println(v8, v8)
    v8 = v8
    var v9 uint64 = 1
    if false {
        v9 = 67
    }
    v10 := GenS2{F0: 2, F1: 107}
    v11 := GenS3{F0: v4}
    var v12 int32 = -1
    v2 = v2
    if (2147483647 != (v8 * (v8 / ((v8) | 1)))) {
        v4 = (uint8(v8) <= 219)
        switch -1 {
        case -191:
            if v4 {
                var v13 uint16 = 741
                for fi14 := 0; fi14 < 17; fi14++ {
                    if ((float64((-2.25 / 1e10)) < (v3 - v3)) && true) {
                        if len(v1) > 0 {
                            v1[0] = (v8 * 65535)
                        }
                    }
                    if true {
                        fmt.Println(fi14)
                        var v15 int64 = v2
                        _ = v15
                    } else {
                        var v16 uint16 = v13
                        if (!v4) {
                            v8 += (v8 + 4294967295)
                        }
                        var v17 float32 = 3.14159
                        _ = v16
                        _ = v17
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi14
                }
                v18 := GenS2{F0: 42, F1: v12}
                _ = v13
                _ = v18
            }
        case 0:
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(v12)
                _ = fi19
            }
        }
        var v20 uint16 = 255
        _ = v20
    } else {
        v9 = (v9 * v9)
        if len(v1) > 0 {
            v1[0] = 1
        }
        if false {
            fmt.Println(v9, v12, v9)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch 4294967295 {
            case -926:
                v9 += (v9 / ((v9) | 1))
            case -9223372036854775808:
                v4 = v4
            }
        }
    }
    v9 = (v9 % ((2) | 1))
    for fi21 := 0; fi21 < 5; fi21++ {
        if ((true || false) && v4) {
            for fi22 := 0; fi22 < 3; fi22++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 = (22 != 128)
                v23 := GenS1{F0: 205}
                var v24 float32 = 0.0
                _ = fi22
                _ = v23
                _ = v24
            }
            if (1.5 != 3.14159) {
                switch (fi21 | fi21) {
                case 0:
                    v25 := GenS2{F0: 42, F1: (v12 ^ (v12 >> ((0) & 63)))}
                    _ = v25
                }
                v2 = (v2 % ((v2) | 1))
            } else {
                v26 := GenS1{F0: 444}
                if v4 {
                    switch (fi21 * (^(fi21))) {
                    case 127:
                        v27 := GenS4{F0: 255, F1: 100}
                        _ = v27
                    case 42:
                        for fi28 := 0; fi28 < 9; fi28++ {
                            fmt.Println(v2)
                            v29 := GenS3{F0: v4}
                            _ = fi28
                            _ = v29
                        }
                    case 643:
                        v30 := GenS5{F0: (v12 / (((v12 ^ (v12 - v12))) | 1)), F1: (v9 % ((v9) | 1)), F2: ((-32768 < 898) && (128 <= 114))}
                        _ = v30
                    }
                } else {
                    fmt.Println(v3)
                    var v31 int8 = 100
                    v32 := GenS4{F0: 0, F1: (v31 & (v31 | v31))}
                    _ = v31
                    _ = v32
                }
                _ = v26
            }
        } else {
            var v33 float64 = v3
            v12 = v12
            _ = v33
        }
        v4 = (v4 || ((1000 != -329) || v4))
        if true {
            v34 := GenS5{F0: int32(v8), F1: v9, F2: v4}
            _ = v34
        }
        _ = fi21
    }
    v4 = (v4 && v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
