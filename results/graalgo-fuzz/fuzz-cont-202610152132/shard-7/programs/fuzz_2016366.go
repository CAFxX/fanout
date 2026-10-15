package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int) bool {
    v6 := []int{(p1 | p1), p1, p1}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int32 = -515
        var v9 float32 = 0.0
        v8 += (v8 / ((v8) | 1))
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            p0 -= (p0 >> ((0) & 63))
            if false {
                if len(v6) > 0 {
                    v6[0] = (fi11 * 4294967295)
                }
                var v12 uint64 = 911
                _ = v12
            } else {
                var v13 bool = false
                _ = v13
            }
            var v14 int16 = -1
            _ = fi11
            _ = v14
        }
        p1 = 32767
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v6
    return (-1 > 1)
}
func genFn1(p0 int8, p1 int32, p2 string) uint8 {
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := []int{fi15, 9223372036854775807}
        _ = fi15
        _ = v16
    }
    var v17 int64 = -93
    var v18 int = 0
    v19 := GenS1{F0: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v19
    return 34
}

func main() {
    v1 := []int{959, 315, 1}
    var v2 uint8 = 128
    v3 := make(map[uint32]uint32)
    var v4 int16 = -707
    var v5 uint64 = 1
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 bool = true
        _ = fi20
        _ = v21
    }
    v2 -= 141
    v4 = v4
    var v22 int32 = 964
    var v23 int16 = v4
    if true {
        v5 -= v5
        var v24 float32 = (-2.25 - (1.5 - (3.14159 + 3.14159)))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v24
    } else {
        fmt.Println(v23)
        for fi25 := 0; fi25 < 18; fi25++ {
            if ((127 != (v2 / ((v2) | 1))) && false) {
                if false {
                    var v26 uint16 = 0
                    v5 = 4294967295
                    _ = v26
                }
                switch "hello" {
                case "hello":
                    v2 = (^(v2))
                default:
                    var v27 uint16 = 2
                    _ = v27
                }
                v23 = (v4 - int16(v2))
            }
            var v28 string = "hello"
            _ = fi25
            _ = v28
        }
        v2 -= (^(v2))
    }
    v29 := GenS1{F0: 191}
    v30 := GenS1{F0: v2}
    fmt.Println(v4)
    fmt.Println(v22, v2, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
