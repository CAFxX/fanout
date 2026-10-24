package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0() uint {
    v6 := []int{4294967295, -960, -943, 42}
    var v7 int16 = 1
    v7 = (v7 << ((2) & 63))
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 int = 128
    var v2 float64 = float64(0.0)
    v3 := make(map[uint32]bool)
    var v4 bool = (!false)
    v5 := []uint32{}
    v8 := v3[300]
    var v9 float64 = v2
    v10 := GenS1{F0: 2, F1: 0, F2: -462}
    v2 += 0.0
    fmt.Println(v8, v2, v9)
    var v11 float64 = (v9 + (v9 / v9))
    var v12 uint32 = 2
    if v8 {
        v8 = (v1 < 929)
        v4 = false
    } else {
        var v13 string = ((("a" + (("\n\t" + ((("a" + (("foo bar" + "\n\t") + "foo bar")) + "foo bar") + "")) + "hello")) + "a") + "foo bar")
        var v14 uint16 = 676
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 18446744073709551615, F1: 2, F2: -1}
    v16 := GenS1{F0: 255, F1: 255, F2: -9223372036854775808}
    if v8 {
        for fi17 := 0; fi17 < 17; fi17++ {
            for fi18 := 0; fi18 < 9; fi18++ {
                if ((0.0 / (-2.25 / -2.25)) >= 0.0) {
                    var v19 uint = 4294967295
                    v3[426] = (!v8)
                    v3[v12] = ((v9 * (v9 + (v11 + (3.14159 + v2)))) == 1e10)
                    var v20 uint64 = 18446744073709551615
                    fmt.Println(fi17, v19, fi17)
                    _ = v19
                    _ = v20
                } else {
                    if ((^(fi17)) > 2147483647) {
                        var v21 uint16 = 1
                        _ = v21
                    } else {
                        v8 = (9223372036854775807 > 0)
                    }
                    var v22 uint64 = 0
                    _ = v22
                }
                v11 = (1e10 * v11)
                _ = fi18
            }
            v23 := GenS1{F0: 255, F1: 32767, F2: -1}
            _ = fi17
            _ = v23
        }
        var v24 uint64 = 1
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 float64 = v2
            if len(v5) > 0 {
                v5[0] = v12
            }
            var v27 uint64 = v24
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = v24
    }
    v9 = -2.25
    if v8 {
        switch "foo bar" {
        case "hello":
            var v28 uint8 = 2
            _ = v28
        case "a":
            v29 := GenS1{F0: genFn0(), F1: 981, F2: 0}
            _ = v29
        default:
            if ((240 == 2) && (1 > 5)) {
                for fi30 := 0; fi30 < 11; fi30++ {
                    var v31 uint = 42
                    if v4 {
                        v32 := GenS1{F0: v31, F1: 1, F2: -9223372036854775808}
                        _ = v32
                    }
                    var v33 bool = (!v8)
                    _ = fi30
                    _ = v31
                    _ = v33
                }
                v34 := GenS1{F0: 197, F1: 2, F2: 158}
                _ = v34
            }
        }
        var v35 uint32 = v12
        v36 := GenS2{F0: 42, F1: 1000}
        _ = v35
        _ = v36
    } else {
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
