package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 string = (("a" + "") + "\n\t")
    var v2 int64 = 156
    var v3 float32 = (((0.0 - (3.14159 / 0.1)) - 3.14159) * 0.1)
    var v4 string = (("foo bar" + "") + v1)
    var v5 int32 = -730
    var v6 int = 2147483647
    var v7 uint32 = 42
    v5 = v5
    v2 += (v2 << ((2) & 63))
    var v8 uint64 = 18446744073709551615
    if false {
        var v9 float32 = -2.25
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                if ((true && false) && (-32768 == 1000)) {
                    v3 = v9
                    var v12 string = v4
                    _ = v12
                }
                v6 = fi10
                _ = fi11
            }
            fmt.Println(v6, v9, v3)
            _ = fi10
        }
        var v13 int64 = v2
        _ = v13
        _ = v9
    } else {
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS1{F0: 360, F1: -592}
            _ = fi14
            _ = v15
        }
        v16 := []int{v6}
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            if false {
                if ((true && (65535 >= 0)) && true) {
                    v19 := []int{}
                    _ = v19
                } else {
                    v20 := GenS2{F0: true, F1: v5}
                    _ = v20
                }
                v5 = (v5 ^ v5)
                if true {
                    var v21 uint8 = 255
                    fmt.Println(v8, fi17)
                    v22 := GenS2{F0: true, F1: 305}
                    _ = v21
                    _ = v22
                }
            } else {
                v23 := []int{4294967295, fi17}
                _ = v23
            }
            var v24 string = ("\n\t" + v1)
            var v25 string = (v24 + (v1 + ""))
            _ = fi18
            _ = v24
            _ = v25
        }
        _ = fi17
    }
    v1 = "foo bar"
    v26 := []int{v6}
    v27 := GenS3{F0: -892, F1: 128}
    var v28 int16 = -1
    v29 := GenS4{F0: v2}
    for fi30 := 0; fi30 < 11; fi30++ {
        if len(v26) > 0 {
            v26[0] = 1
        }
        _ = fi30
    }
    var v31 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
