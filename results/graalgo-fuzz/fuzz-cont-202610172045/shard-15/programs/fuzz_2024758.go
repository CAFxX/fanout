package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }

func main() {
    var v1 int32 = -783
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    for fi5 := 0; fi5 < 18; fi5++ {
        switch "hello" {
        case "hello":
            v6 := []int{fi5, 4294967295, 913, (fi5 / ((fi5) | 1))}
            _ = v6
        case "":
            v2[("foo bar" + "")] = (fi5 / ((int(v1)) | 1))
        case "foo bar":
            v7 := []int{(^(fi5)), fi5, fi5}
            _ = v7
        }
        v4[int(v1)] = 2
        _ = fi5
    }
    var v8 uint16 = 656
    v8 = 86
    v9 := GenS1{F0: 128, F1: 0}
    switch 91 {
    case 2147483647:
        v10 := GenS1{F0: 1, F1: 42}
        _ = v10
    default:
        var v11 float32 = ((float32((((0.1 - ((((0.1 * -2.25) + 1.5) - 0.0) - 1.5)) - ((3.14159 / (float64(float32(1e10)) + 0.1)) - -2.25)) - ((1e10 * -2.25) * 0.1))) / 1e10) - 1.5)
        _ = v11
    }
    switch -497 {
    case 127:
        v12 := GenS1{F0: 127, F1: 42}
        _ = v12
    case 0:
        if (42 < 82) {
            v8 = v8
            if true {
                var v13 uint32 = 0
                _ = v13
            } else {
                if (!true) {
                    var v14 uint16 = v8
                    v1 = -914
                    v1 = (^(v1))
                    _ = v14
                } else {
                    var v15 int64 = 80
                    _ = v15
                }
                var v16 int8 = 45
                _ = v16
            }
            switch "" {
            case "a":
                v17 := []int{9223372036854775807, 1, -2147483648, 32768}
                _ = v17
            case "foo bar":
                if (false && (v1 >= (v1 - 0))) {
                    v18 := []int{1, -598, 2147483647, 32768}
                    switch "a" {
                    case "a":
                        v8 += (v8 % ((344) | 1))
                    case "hello":
                        v8 = (v8 << ((42) & 63))
                    case "\n\t":
                        fmt.Println(v1, v1)
                    default:
                        var v19 uint8 = 128
                        _ = v19
                    }
                    _ = v18
                } else {
                    var v20 float64 = 3.14159
                    var v21 int32 = (v1 ^ v1)
                    _ = v20
                    _ = v21
                }
            case "hello":
                _ = len(v3)
            default:
                v8 += 45
            }
        } else {
            for fi22 := 0; fi22 < 19; fi22++ {
                for fi23 := 0; fi23 < 2; fi23++ {
                    v1 = (^(v1))
                    _ = fi23
                }
                var v24 uint8 = 1
                _ = fi22
                _ = v24
            }
        }
    }
    v1 = v1
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 int16 = -414
            _ = fi26
            _ = v27
        }
        var v28 uint64 = 182
        v29 := GenS1{F0: int(v1), F1: 0}
        _ = fi25
        _ = v28
        _ = v29
    }
    var v30 uint16 = 2
    var v31 uint16 = v30
    v31 = 1
    v30 = (v31 ^ v30)
    v8 -= 271
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v8
    _ = v9
}
