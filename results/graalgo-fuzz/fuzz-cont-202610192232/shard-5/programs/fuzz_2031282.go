package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int8 }
func genFn0() uint32 {
    var v6 int16 = -419
    var v7 int32 = 42
    var v8 uint8 = 0
    switch 0 {
    case 4294967295:
        var v9 string = "\n\t"
        _ = v9
    case -592:
        v10 := GenS1{F0: 1}
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]int)
    var v3 int = 32767
    var v4 uint16 = 255
    var v5 int32 = -728
    var v11 int8 = 0
    v12 := GenS2{F0: (!(32767 >= 575)), F1: ((255 != 1) && false)}
    var v13 uint64 = uint64(v1)
    v1 = 2147483647
    v14 := []int{(v3 >> ((1) & 63)), 0, (-(v3)), (v3 << ((209) & 63))}
    var v15 int = (v3 & (v3 % ((v3) | 1)))
    switch (v15 | (v15 - 127)) {
    case 4294967295:
        var v16 float32 = 0.1
        _ = v16
    }
    v2[""] = (v15 - v3)
    var v17 bool = (true || false)
    switch "hello" {
    case "\n\t":
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int32 = 106
            v20 := GenS3{F0: (-(v11))}
            var v21 int32 = -208
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 8; fi22++ {
            v1 = (v1 * (v1 & (v1 + v1)))
            v23 := v2[""]
            _ = fi22
            _ = v23
        }
    }
    v14 = append(v14, -846)
    fmt.Println(v5)
    switch (v15 * 626) {
    case 32768:
        var v24 uint16 = 171
        _ = v24
    case -9223372036854775808:
        var v25 int16 = 499
        _ = v25
    case 0:
        switch "a" {
        case "hello":
            for fi26 := 0; fi26 < 10; fi26++ {
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 uint32 = v1
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
        case "":
            v5 -= (v5 ^ v5)
        case "foo bar":
            for fi29 := 0; fi29 < 7; fi29++ {
                if true {
                    fmt.Println(v1)
                } else {
                    var v30 int64 = 857
                    v1 = 10
                    _ = v30
                }
                _ = len(v14)
                _ = fi29
            }
        }
    default:
        var v31 int8 = (v11 / (((v11 + v11)) | 1))
        _ = v31
    }
    v14 = append(v14, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
