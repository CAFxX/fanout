package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int64 }
func genFn0() bool {
    var v7 float32 = 1e10
    if (!true) {
        v8 := GenS1{F0: 102, F1: 0, F2: 97}
        var v9 uint32 = 0
        v10 := GenS1{F0: 65535, F1: 65535, F2: 1}
        _ = v10
        _ = v8
        _ = v9
    } else {
        v7 -= 1.5
        var v11 int16 = 179
        var v12 bool = false
        _ = v11
        _ = v12
    }
    var v13 uint32 = 4294967295
    _ = v13
    _ = v7
    return false
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v4 uint8 = 61
    var v5 uint64 = 18446744073709551615
    var v6 int32 = -586
    v6 += v6
    var v14 int64 = 0
    var v15 uint32 = 42
    if true {
        v14 += (^(v14))
        var v16 int8 = 100
        v4 -= v4
        _ = v16
    } else {
        switch int(v14) {
        case 32768:
            if (!(v6 != v6)) {
                v14 = (v14 + 1)
                v2[(^(v15))] = (("" + "a") + "a")
                var v17 string = ""
                _ = v17
            } else {
                if false {
                    v4 += (v4 * v4)
                }
                if (!(!true)) {
                    var v18 int16 = -32768
                    _ = v18
                }
            }
        default:
            var v19 int8 = -128
            _ = v19
        }
    }
    var v20 int8 = 42
    v6 = (v6 ^ v6)
    var v21 uint = 762
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 int16 = 42
        if false {
            var v24 bool = true
            for fi25 := 0; fi25 < 12; fi25++ {
                v26 := v2[v15]
                v15 += (v15 >> (((v21 >> (((v21 | (v21 / ((255) | 1)))) & 63))) & 63))
                v1 = 9223372036854775807
                _ = fi25
                _ = v26
            }
            _ = v24
        }
        var v27 string = "hello"
        _ = fi22
        _ = v23
        _ = v27
    }
    var v28 int32 = v6
    fmt.Println(v20)
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 15; fi30++ {
            if (2 <= uint64(v4)) {
                v1 = fi29
                if (((0.1 + float64(((((0.0 + 1e10) - 1.5) + ((1.5 / 0.1) - 0.0)) / 1.0))) < 0.1) && true) {
                    v31 := []int{}
                    _ = v31
                } else {
                    v32 := []int{v1}
                    if false {
                        v4 = (v4 >> (((v21 | 123)) & 63))
                        v1 -= v1
                    } else {
                        var v33 float64 = -2.25
                        _ = v33
                    }
                    var v34 bool = (892 >= (v15 & v15))
                    _ = v32
                    _ = v34
                }
            } else {
                v5 = (v5 ^ 380)
            }
            var v35 int32 = v6
            _ = fi30
            _ = v35
        }
        v20 += (v20 + v20)
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
