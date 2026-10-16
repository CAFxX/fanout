package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 int16, p2 uint8) uint64 {
    v7 := GenS1{F0: 557, F1: (p1 / ((42) | 1))}
    v8 := GenS2{F0: -1}
    var v9 bool = true
    var v10 uint64 = 2
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 string = "foo bar"
        v13 := GenS1{F0: 1, F1: p1}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (v10 << ((18446744073709551615) & 63))
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = -2.25
    var v3 uint32 = 268
    var v4 int64 = 1
    v5 := make(map[uint32]uint32)
    var v6 uint16 = (v1 / (((v1 - (v1 / ((2) | 1)))) | 1))
    for fi14 := 0; fi14 < 15; fi14++ {
        v4 = (v4 * (v4 - (v4 | (v4 << ((4294967295) & 63)))))
        var v15 uint = 255
        _ = fi14
        _ = v15
    }
    var v16 int8 = -128
    fmt.Println(v1, v16)
    var v17 uint64 = 0
    var v18 float32 = float32(v2)
    v2 += (v2 + -2.25)
    fmt.Println(v4)
    fmt.Println(v3, v4)
    v16 += (v16 - v16)
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 int64 = int64(v17)
        v21 := v5[744]
        _ = fi19
        _ = v20
        _ = v21
    }
    if true {
        if true {
            v3 = (^(v3))
            v3 -= 1
        } else {
            switch -9223372036854775808 {
            case -590:
                var v22 int8 = 42
                _ = v22
            case 1:
                v3 -= (^(v3))
            }
            fmt.Println(v18, v1)
        }
        var v23 uint = 580
        var v24 float32 = v18
        _ = v23
        _ = v24
    } else {
        if false {
            delete(v5, (v3 >> ((0) & 63)))
        } else {
            v1 = 1
            fmt.Println(v16)
            var v25 string = ((((("\n\t" + (("\n\t" + "hello") + "a")) + "\n\t") + ("" + "a")) + "a") + "\n\t")
            _ = v25
        }
        var v26 uint32 = v3
        _ = v26
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 uint32 = v3
        v16 = (v16 | 5)
        v29 := GenS3{F0: int32(v1)}
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        v31 := GenS3{F0: 1}
        v6 = v6
        v32 := GenS1{F0: 4294967295, F1: -1}
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
