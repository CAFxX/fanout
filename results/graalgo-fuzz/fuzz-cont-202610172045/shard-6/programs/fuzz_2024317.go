package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int8 }
func genFn0() int {
    var v6 int16 = -1
    v6 -= v6
    var v7 uint = 4294967295
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 int = -826
    var v2 uint16 = 1
    var v3 int8 = -2
    var v4 int8 = (v3 % ((v3) | 1))
    v5 := make(map[uint32]uint32)
    var v8 float64 = 0.0
    v9 := GenS1{F0: "a", F1: 250}
    v2 += v2
    if (true || (!(2147483648 == 1))) {
        fmt.Println(v1, v1)
    }
    v4 += 83
    v4 = int8(v2)
    v1 += v1
    var v10 string = "\n\t"
    switch (v1 << ((2) & 63)) {
    case 42:
        v11 := GenS2{F0: 2147483648, F1: -2147483648, F2: (v2 * (v2 ^ 65535))}
        _ = v11
    case 9223372036854775807:
        v1 = v1
    }
    var v12 uint64 = 2
    switch genFn0() {
    case 42:
        switch genFn0() {
        case -929:
            fmt.Println(v3)
        case 127:
            if false {
                switch "\n\t" {
                case "hello":
                    v4 += v3
                default:
                    if (902 < genFn0()) {
                        var v13 uint64 = v12
                        switch genFn0() {
                        case 42:
                            v14 := []int{}
                            _ = v14
                        default:
                            v15 := GenS3{F0: v2}
                            _ = v15
                        }
                        _ = v13
                    } else {
                        fmt.Println(v3, v12, v2)
                    }
                }
                v1 = v1
                var v16 bool = (-93 <= (v4 % ((0) | 1)))
                _ = v16
            } else {
                if (255 >= uint8(v2)) {
                    for fi17 := 0; fi17 < 6; fi17++ {
                        for fi18 := 0; fi18 < 5; fi18++ {
                            var v19 int8 = (v3 & (v3 ^ 0))
                            var v20 uint = 2
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = fi17
                    }
                    v21 := []int{v1, 127, genFn0()}
                    _ = v21
                }
                var v22 float64 = v8
                for fi23 := 0; fi23 < 20; fi23++ {
                    v2 -= 42
                    _ = fi23
                }
                _ = v22
            }
        }
    case -709:
        v24 := GenS4{F0: -9223372036854775808, F1: v4}
        _ = v24
    }
    if (0 == 1) {
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 uint8 = 127
            var v27 float32 = 0.0
            switch fi25 {
            case -642:
                var v28 int64 = 1
                _ = v28
            case 128:
                v26 = (v26 | (v26 + v26))
            case 0:
                v27 = 0.0
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = len(v5)
        switch (v10 + v10) {
        case "foo bar":
            if true {
                if true {
                    var v29 float32 = float32(v8)
                    fmt.Println(v3, v29, v29)
                    _ = v29
                } else {
                    v8 = (0.0 - float64(float32(v8)))
                }
            } else {
                var v30 uint16 = v2
                _ = len(v5)
                v31 := GenS4{F0: 0, F1: 1}
                _ = v30
                _ = v31
            }
        }
    } else {
        var v32 uint16 = 1
        switch v10 {
        case "\n\t":
            var v33 uint64 = 1
            _ = v33
        case "":
            for fi34 := 0; fi34 < 7; fi34++ {
                for fi35 := 0; fi35 < 11; fi35++ {
                    switch v10 {
                    case "a":
                        v4 = (^(v3))
                        switch v10 {
                        case "":
                            var v36 int32 = int32(v4)
                            _ = v36
                        }
                        v32 -= 238
                        var v37 bool = (0 != 127)
                        _ = v37
                    case "hello":
                        var v38 uint = 890
                        _ = v38
                    }
                    _ = fi35
                }
                var v39 uint64 = 0
                _ = fi34
                _ = v39
            }
        case "a":
            var v40 uint64 = (^(v12))
            _ = v40
        }
        _ = v32
    }
    switch -202 {
    case -9223372036854775808:
        v41 := GenS2{F0: 2147483648, F1: 42, F2: (v2 << ((675) & 63))}
        _ = v41
    case 282:
        var v42 float32 = (((float32((1e10 + v8)) * (1.5 / (((3.14159 - (((3.14159 * ((-2.25 * float32(v8)) + (0.1 * 3.14159))) + float32(float64(((float32(float64((0.1 + 1.5))) + (1e10 + float32(v8))) - float32(((3.14159 - 3.14159) / v8)))))) - 1e10)) / 1e10) * ((1e10 + 0.0) * 0.1)))) + float32((float64(1.5) + (v8 - v8)))) + (((((1e10 * 0.1) + 0.1) - 1e10) + 3.14159) * 3.14159))
        _ = v42
    }
    var v43 float64 = v8
    for fi44 := 0; fi44 < 3; fi44++ {
        var v45 float64 = v43
        var v46 int64 = -1
        v47 := []int{genFn0(), (v1 | fi44), 492}
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v8
    _ = v9
}
