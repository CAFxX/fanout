package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{4294967295}
    v3 := make(map[int]int)
    var v4 uint = 4294967295
    var v5 bool = (0 == 772)
    var v6 string = "a"
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint64 = 23
        if v5 {
            v6 = (v6 + v6)
            for fi9 := 0; fi9 < 10; fi9++ {
                v3[fi9] = fi9
                _ = fi9
            }
        }
        var v10 uint16 = 255
        _ = fi7
        _ = v10
        _ = v8
    }
    var v11 float32 = 1.5
    var v12 int64 = 42
    var v13 uint = (v4 >> ((v4) & 63))
    v11 = v11
    var v14 bool = v5
    _ = len(v1)
    switch 9223372036854775807 {
    case -9223372036854775808:
        if v5 {
            var v15 float32 = (1.5 * (v11 + v11))
            var v16 uint64 = 561
            _ = v15
            _ = v16
        } else {
            var v17 int64 = (v12 & v12)
            switch 32768 {
            case 9223372036854775807:
                switch -597 {
                case 42:
                    var v18 bool = (255 > 0)
                    _ = v18
                default:
                    var v19 int8 = -128
                    _ = v19
                }
            case 483:
                v4 = v4
            case -363:
                if false {
                    var v20 float64 = 0.1
                    var v21 string = v6
                    v13 += 65535
                    _ = v20
                    _ = v21
                }
            default:
                v22 := GenS1{F0: 0}
                _ = v22
            }
            _ = v17
        }
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        if true {
            for fi24 := 0; fi24 < 6; fi24++ {
                v5 = v14
                fmt.Println(v5, v6)
                _ = len(v1)
                _ = fi24
            }
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 float32 = (-2.25 + float32(0.1))
                _ = fi25
                _ = v26
            }
            for fi27 := 0; fi27 < 12; fi27++ {
                v14 = false
                var v28 int32 = 635
                v29 := GenS1{F0: 743}
                _ = fi27
                _ = v28
                _ = v29
            }
        } else {
            if (!v5) {
                v30 := GenS1{F0: 4294967295}
                _ = v30
            }
            v4 += v13
        }
        for fi31 := 0; fi31 < 20; fi31++ {
            fmt.Println(v13, v13, v13)
            v32 := v1[185]
            _ = fi31
            _ = v32
        }
        var v33 int8 = -32
        _ = fi23
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
