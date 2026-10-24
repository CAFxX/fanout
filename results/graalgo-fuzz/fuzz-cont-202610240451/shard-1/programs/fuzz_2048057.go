package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int16; F1 int16; F2 int64 }
func genFn0(p0 bool, p1 uint32, p2 int32) uint {
    var v7 uint8 = 238
    for fi8 := 0; fi8 < 5; fi8++ {
        if ("hello" == ("" + "hello")) {
            v9 := []int{(fi8 * (fi8 % ((-1) | 1)))}
            v9 = append(v9, (fi8 | fi8))
            _ = v9
        } else {
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := []int{-639}
                _ = fi10
                _ = v11
            }
            p2 = p2
            var v12 float32 = 1e10
            _ = v12
        }
        var v13 int = (fi8 - (fi8 ^ 128))
        _ = fi8
        _ = v13
    }
    v14 := GenS1{F0: 127, F1: (v7 | 1), F2: 2}
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint8 = 42
        v17 := GenS1{F0: fi15, F1: v7, F2: (v7 - 62)}
        var v18 uint32 = (^(p1))
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := []int{fi19, fi19, fi19, (fi19 & 2147483647)}
        p0 = (p2 <= 0)
        switch "foo bar" {
        case "hello":
            fmt.Println(p0)
        case "":
            var v21 uint64 = 2
            _ = v21
        default:
            switch "foo bar" {
            case "\n\t":
                if p0 {
                    v22 := GenS1{F0: (fi19 % ((fi19) | 1)), F1: v7, F2: (v7 | v7)}
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v20 = append(v20, (fi23 << ((2) & 63)))
                        v20 = append(v20, (fi19 - fi23))
                        _ = fi23
                    }
                    _ = len(v20)
                    _ = v22
                } else {
                    var v24 int16 = 1000
                    _ = v24
                }
            case "foo bar":
                var v25 bool = true
                _ = v25
            default:
                var v26 uint16 = 32767
                _ = v26
            }
        }
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v7
    return 478
}
func genFn1(p0 uint) uint8 {
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 uint32 = 240
        _ = fi27
        _ = v28
    }
    v29 := []int{}
    v30 := GenS1{F0: 650, F1: 0, F2: 255}
    var v31 bool = ("\n\t" <= "foo bar")
    v32 := GenS1{F0: -2147483648, F1: 1, F2: 127}
    _ = p0
    _ = v29
    _ = v30
    _ = v31
    _ = v32
    return uint8(p0)
}

func main() {
    var v1 uint = 255
    var v2 int32 = 335
    v3 := make(map[int]int)
    var v4 int8 = 42
    var v5 int64 = -64
    var v6 int32 = (v2 >> (((v1 / (((v1 << ((v1) & 63))) | 1))) & 63))
    if false {
        for fi33 := 0; fi33 < 6; fi33++ {
            v1 -= 262
            _ = fi33
        }
    }
    v3[-2147483648] = 9223372036854775807
    delete(v3, 128)
    var v34 float32 = (float32((1e10 / 0.1)) + (-2.25 / 1.0))
    v35 := GenS2{F0: 492, F1: 0, F2: v5}
    var v36 uint32 = 85
    fmt.Println(v5, v34, v4)
    if (2 > 2) {
        if true {
            v6 = -1
        }
        v37 := GenS2{F0: 0, F1: -32768, F2: v5}
        _ = v37
    } else {
        if false {
            v2 += v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
