package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint8) uint32 {
    v6 := []int{802, 32767}
    var v7 int16 = 0
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint16 = 202
    v2 := make(map[int]uint32)
    v3 := []int{4294967295, 1, -2147483648}
    var v4 uint64 = 0
    v5 := make(map[int]bool)
    var v8 uint = 18446744073709551615
    v9 := GenS1{F0: (v8 / ((65535) | 1))}
    v10 := GenS1{F0: (v8 & v8)}
    switch "" {
    case "a":
        v11 := GenS1{F0: (v8 >> (((^(v8))) & 63))}
        _ = v11
    case "":
        for fi12 := 0; fi12 < 8; fi12++ {
            switch fi12 {
            case 9223372036854775807:
                v4 = uint64(v1)
            case 861:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 4294967295:
                switch (fi12 / ((fi12) | 1)) {
                case 9223372036854775807:
                    v13 := GenS2{F0: 42}
                    _ = v13
                case 42:
                    var v14 bool = true
                    _ = v14
                case 757:
                    var v15 int = fi12
                    _ = v15
                }
            default:
                v16 := GenS3{F0: -1}
                _ = v16
            }
            var v17 uint64 = v4
            var v18 uint16 = (v1 ^ v1)
            _ = fi12
            _ = v17
            _ = v18
        }
    case "\n\t":
        var v19 int64 = -9223372036854775808
        _ = v19
    default:
        v20 := GenS3{F0: -1}
        _ = v20
    }
    fmt.Println(v8, v4, v1)
    v1 -= v1
    if true {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 int16 = int16(fi21)
            var v23 float32 = (-2.25 * 0.0)
            _ = fi21
            _ = v22
            _ = v23
        }
        if len(v3) > 0 {
            v3[0] = 32768
        }
    }
    if true {
        var v24 uint = v8
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 uint8 = 128
            v1 = 32767
            var v27 float64 = -2.25
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = v24
    }
    if len(v3) > 0 {
        v3[0] = 127
    }
    switch "hello" {
    case "\n\t":
        v8 = (v8 * 0)
    case "a":
        v1 += (v1 >> (((v8 >> ((v8) & 63))) & 63))
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        switch 32768 {
        case 0:
            var v29 uint = 0
            _ = v29
        default:
            if (genFn0(0) >= 2147483648) {
                for fi30 := 0; fi30 < 7; fi30++ {
                    v31 := GenS2{F0: int8(v1)}
                    v1 = (v1 / ((v1) | 1))
                    v32 := GenS4{F0: (4294967295 < (fi28 & 128))}
                    if (2 != 0) {
                        v4 = v4
                        var v33 int16 = int16(v1)
                        delete(v2, 32767)
                        _ = v33
                    } else {
                        var v34 int16 = 958
                        v8 = 669
                        _ = v34
                    }
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                var v35 uint = (^(v8))
                _ = v35
            } else {
                var v36 uint8 = 188
                _ = v36
            }
        }
        var v37 float64 = 0.0
        _ = fi28
        _ = v37
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        for fi39 := 0; fi39 < 1; fi39++ {
            var v40 int64 = -9223372036854775808
            v4 += 86
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
