package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint) int8 {
    fmt.Println(p0)
    var v4 int64 = 611
    _ = p0
    _ = v4
    return 42
}
func genFn1(p0 uint) int16 {
    v5 := GenS1{F0: true}
    if true {
        if true {
            var v6 float32 = ((0.1 - 1.5) / ((0.1 - -2.25) / 1.0))
            p0 -= p0
            p0 -= 4294967295
            _ = v6
        }
        if true {
            v7 := []int{}
            var v8 uint64 = 0
            p0 += p0
            _ = v7
            _ = v8
        } else {
            fmt.Println(p0)
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS1{F0: ((fi9 > (-(fi9))) || false)}
                _ = fi9
                _ = v10
            }
        }
    }
    p0 -= (p0 >> (((p0 - (p0 ^ 2))) & 63))
    var v11 uint64 = 2
    _ = p0
    _ = v11
    _ = v5
    return -32768
}

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    v1 = (v1 & (v1 >> ((1) & 63)))
    fmt.Println(v1)
    v12 := GenS2{F0: 22}
    switch ("\n\t" + "hello") {
    case "hello":
        v13 := GenS2{F0: -94}
        _ = v13
    default:
        switch "" {
        case "foo bar":
            var v14 float32 = 0.0
            _ = v14
        case "":
            fmt.Println(v1)
        default:
            if true {
                _ = len(v2)
            } else {
                var v15 float32 = (1e10 - float32(float64(-2.25)))
                var v16 uint8 = 127
                _ = v15
                _ = v16
            }
        }
    }
    v1 -= v1
    delete(v2, 392)
    v1 = v1
    var v17 string = ("" + "\n\t")
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int8 = genFn0(65535)
        if (255 > 2) {
            v1 -= v1
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 float32 = 1e10
                if (v17 < v17) {
                    delete(v3, (fi20 + -766))
                }
                _ = fi20
                _ = v21
            }
            var v22 float32 = (-2.25 * 1.5)
            _ = v22
        }
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS1{F0: (30 == 349)}
            _ = fi23
            _ = v24
        }
        _ = fi18
        _ = v19
    }
    v17 = "foo bar"
    fmt.Println(v1)
    if false {
        var v25 int8 = 42
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 uint = 4294967295
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        v1 += (v1 + 899)
        v1 = v1
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 10; fi29++ {
                v30 := []int{(fi29 & (fi29 >> ((0) & 63)))}
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    }
    var v31 uint = 44
    var v32 int64 = (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
}
