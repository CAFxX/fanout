package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 bool; F2 uint32 }
type GenS4 struct { F0 int32; F1 string; F2 int8 }
type GenS5 struct { F0 uint; F1 int64; F2 uint8 }
type GenS6 struct { F0 int64; F1 uint; F2 uint32 }

func main() {
    var v1 int8 = -1
    var v2 int16 = -32768
    var v3 uint64 = 0
    v4 := []uint32{}
    if (1 > 599) {
        var v5 bool = (true || false)
        switch 42 {
        case -770:
            var v6 float64 = (1.5 * 0.1)
            _ = v6
        case 32767:
            var v7 string = "hello"
            _ = v7
        }
        fmt.Println(v1, v5)
        _ = v5
    } else {
        if (42 > 32767) {
            v3 = (v3 * v3)
            var v8 uint16 = 143
            _ = v8
        }
    }
    v3 = v3
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: -2147483648}
        fmt.Println(v3, fi9, fi9)
        _ = fi9
        _ = v10
    }
    fmt.Println(v2)
    for fi11 := 0; fi11 < 19; fi11++ {
        v1 -= (v1 - v1)
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint16 = 0
            _ = fi12
            _ = v13
        }
        var v14 bool = false
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        if false {
            switch (fi15 ^ (fi15 * (fi15 >> ((42) & 63)))) {
            case 399:
                for fi16 := 0; fi16 < 10; fi16++ {
                    v1 = v1
                    var v17 uint = 4294967295
                    v18 := GenS2{F0: 42}
                    if false {
                        if ((((1.5 - 1.5) * (-2.25 + 1.5)) - 3.14159) >= 1e10) {
                            switch int(v1) {
                            case 75:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case -9223372036854775808:
                                var v19 int64 = 0
                                _ = v19
                            }
                        }
                        if (v3 <= 2) {
                            var v20 int8 = 100
                            var v21 int16 = v2
                            _ = v20
                            _ = v21
                        }
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            case 127:
                v2 = 197
            case 1:
                if true {
                    v3 = (v3 & v3)
                    _ = len(v4)
                    v22 := GenS3{F0: int64(v2), F1: (((((float32(0.0) - (1.5 - (3.14159 - 1.5))) - float32(3.14159)) / 0.1) * (3.14159 * float32(-2.25))) != -2.25), F2: 1}
                    _ = v22
                }
            }
        }
        if (!false) {
            if true {
                if (v3 <= v3) {
                    var v23 uint32 = 4294967295
                    fmt.Println(fi15)
                    _ = len(v4)
                    _ = v23
                } else {
                    v24 := GenS3{F0: 42, F1: (65535 != 336), F2: 42}
                    switch (fi15 ^ (-(fi15))) {
                    case 32768:
                        var v25 float64 = (0.1 / 1.5)
                        _ = v25
                    case 660:
                        v26 := GenS4{F0: -47, F1: ("hello" + "hello"), F2: (v1 ^ 127)}
                        _ = v26
                    }
                    _ = len(v4)
                    _ = v24
                }
                v2 = (v2 | v2)
            }
        } else {
            var v27 uint32 = 546
            _ = v27
        }
        v3 -= (^(v3))
        _ = fi15
    }
    if true {
        v2 -= v2
        v1 = 100
        if (!true) {
            var v28 uint8 = 255
            _ = v28
        }
    }
    v29 := GenS5{F0: 81, F1: -699, F2: 128}
    fmt.Println(v2, v1, v3)
    v30 := GenS6{F0: 0, F1: 510, F2: 2147483648}
    var v31 int = -950
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
