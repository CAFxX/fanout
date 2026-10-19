package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint32) uint {
    v6 := GenS1{F0: 509, F1: 42, F2: 751}
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint8 = 68
        p0 += (^(p0))
        _ = fi7
        _ = v8
    }
    var v9 int = 1
    v9 = (v9 % ((v9) | 1))
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(v9, fi10)
        var v11 uint = 0
        var v12 uint32 = p0
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v6
    _ = v9
    return 467
}
func genFn1() uint8 {
    v13 := []int{}
    var v14 float64 = (1e10 - ((0.0 - (1e10 - (0.1 / 1e10))) / 0.1))
    _ = v13
    _ = v14
    return 29
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (false && (((1e10 + (((float64(3.14159) + 3.14159) - (1.5 * float64(float32(1e10)))) - (1e10 + (0.1 * 3.14159)))) / (0.1 * (3.14159 + 1e10))) > ((1.5 - (1.5 - (3.14159 + ((0.0 * (-2.25 * ((-2.25 / 1.5) / (1e10 + ((0.0 + (3.14159 - 1.5)) - (0.0 - 0.1)))))) + 0.1)))) / 1.0)))
    v3 := make(map[int]uint32)
    v4 := []uint32{18}
    v5 := []string{"\n\t"}
    if (v2 && ((-903 >= -130) && v2)) {
        switch "foo bar" {
        case "":
            var v15 string = "a"
            _ = v15
        }
        var v16 uint = 18446744073709551615
        _ = v16
    } else {
        if v2 {
            v2 = v2
        }
        var v17 uint8 = 2
        fmt.Println(v17)
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v2 = (42 == -1)
        _ = fi18
    }
    var v19 bool = ((false && (((!false) || v2) || v2)) || (0 > 2147483648))
    var v20 bool = ((0 == 174) && (!(512 < 1)))
    fmt.Println(v19)
    var v21 string = ""
    if false {
        switch 32768 {
        case 397:
            if true {
                var v22 int8 = 127
                if (!true) {
                    v23 := GenS2{F0: 65535}
                    _ = v23
                }
                v22 += (v22 / (((v22 << ((437) & 63))) | 1))
                _ = v22
            } else {
                fmt.Println(v20, v20)
                var v24 int16 = -32768
                v21 = ""
                _ = v24
            }
        default:
            fmt.Println(v19, v21)
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            for fi26 := 0; fi26 < 8; fi26++ {
                if v2 {
                    v27 := GenS1{F0: 0, F1: 0, F2: 0}
                    var v28 uint = 4294967295
                    var v29 int64 = 0
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    var v30 uint16 = 32767
                    _ = len(v5)
                    _ = v30
                }
                v2 = (true || (v21 < v21))
                _ = len(v5)
                _ = fi26
            }
            switch v21 {
            case "a":
                var v31 uint8 = 2
                _ = v31
            case "\n\t":
                v19 = (255 > 0)
            }
            _ = fi25
        }
    } else {
        var v32 uint64 = 966
        var v33 int8 = 127
        var v34 int = 52
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 int8 = -94
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
