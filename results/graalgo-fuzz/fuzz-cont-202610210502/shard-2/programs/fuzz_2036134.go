package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 int16; F1 int; F2 uint64 }
type GenS3 struct { F0 int32 }
func genFn0() uint32 {
    if (true || true) {
        var v7 string = (("\n\t" + ("" + "hello")) + "foo bar")
        _ = v7
    } else {
        var v8 uint16 = 1
        _ = v8
    }
    v9 := GenS1{F0: 2, F1: -2147483648, F2: 48}
    var v10 uint8 = 201
    var v11 uint8 = (v10 / ((127) | 1))
    _ = v10
    _ = v11
    _ = v9
    return 290
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ""
    var v3 float64 = 0.0
    v4 := make(map[string]int)
    var v5 uint16 = 248
    var v6 int16 = 1
    if (3.14159 >= (0.0 / 3.14159)) {
        v5 = 947
        v5 += (v5 * v5)
    }
    v12 := []int{364, -322, -9223372036854775808, 128}
    for fi13 := 0; fi13 < 6; fi13++ {
        v3 -= 0.0
        delete(v4, v2)
        if true {
            var v14 int = (fi13 | fi13)
            _ = v14
        } else {
            var v15 int8 = -89
            var v16 int64 = 1
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    var v17 float64 = (((3.14159 / v3) + (1.5 * 3.14159)) - 1e10)
    switch v2 {
    case "a":
        var v18 uint32 = 2
        _ = v18
    }
    if (!false) {
        v17 = 0.1
    } else {
        switch -510 {
        case -425:
            fmt.Println(v6, v2)
        case 32768:
            _ = len(v12)
        default:
            v6 = -32768
        }
        fmt.Println(v2, v5, v5)
        var v19 bool = false
        _ = v19
    }
    v20 := GenS1{F0: v5, F1: 0, F2: 0}
    v2 = v2
    var v21 float32 = 1.5
    for fi22 := 0; fi22 < 13; fi22++ {
        v21 = 0.0
        fmt.Println(v5, v5)
        delete(v4, "foo bar")
        _ = fi22
    }
    if (2147483647 > genFn0()) {
        if false {
            for fi23 := 0; fi23 < 12; fi23++ {
                for fi24 := 0; fi24 < 18; fi24++ {
                    var v25 uint16 = v5
                    var v26 int64 = 0
                    var v27 uint = 2
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = fi23
            }
            if (v2 > ("hello" + "\n\t")) {
                v2 = v2
                var v28 uint16 = (v5 << ((65535) & 63))
                _ = v28
            }
        }
        var v29 int = 32768
        fmt.Println(v17, v3)
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        switch -968 {
        case 336:
            switch 819 {
            case 873:
                _ = len(v12)
            default:
                v31 := GenS2{F0: (^(v6)), F1: (fi30 / ((fi30) | 1)), F2: 18446744073709551615}
                _ = v31
            }
        case 128:
            v32 := GenS3{F0: -649}
            _ = v32
        default:
            v17 += v17
        }
        var v33 int16 = 32767
        var v34 uint64 = 2
        _ = fi30
        _ = v33
        _ = v34
    }
    var v35 float64 = (float64((1e10 + v21)) * v3)
    v35 = v35
    v6 -= (v6 + v6)
    v5 = (v5 % ((42) | 1))
    v36 := GenS3{F0: -2147483648}
    var v37 int16 = 233
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
