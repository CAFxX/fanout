package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint16; F2 int64 }
type GenS3 struct { F0 bool; F1 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 1
    v4 := make(map[string]bool)
    var v5 int64 = 984
    var v6 uint32 = 2
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int16 = 42
        fmt.Println(v3, v6)
        for fi9 := 0; fi9 < 10; fi9++ {
            v6 -= 845
            if false {
                v10 := GenS1{F0: (int8(fi7) >= 42)}
                fmt.Println(v3, v3, v3)
                v8 = (v8 | (v8 * (^(v8))))
                _ = v10
            }
            if (1 < 1) {
                v11 := v1[317]
                _ = v11
            }
            v12 := []int{fi9, 4294967295}
            if (false || ((v5 >> ((17) & 63)) > -1)) {
                v13 := v1[(v6 | v6)]
                _ = v13
            }
            _ = fi9
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    v5 = v5
    fmt.Println(v6, v6)
    switch "" {
    case "hello":
        var v14 bool = true
        _ = v14
    case "":
        for fi15 := 0; fi15 < 15; fi15++ {
            v6 = 1
            _ = fi15
        }
    default:
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 bool = true
            var v18 int = fi16
            v6 = uint32(v3)
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v5 -= int64(v6)
        v20 := []int{fi19, (fi19 & fi19)}
        var v21 uint32 = (v6 * (v6 << ((42) & 63)))
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v3, v6)
    v6 += (v6 >> ((4294967295) & 63))
    var v22 uint8 = 59
    v22 = (^(v22))
    fmt.Println(v3, v22)
    for fi23 := 0; fi23 < 1; fi23++ {
        fmt.Println(fi23)
        var v24 uint = 255
        var v25 float64 = 3.14159
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := GenS2{F0: 100, F1: 32767, F2: v5}
    for fi27 := 0; fi27 < 14; fi27++ {
        for fi28 := 0; fi28 < 19; fi28++ {
            v5 -= (v5 ^ (-(v5)))
            delete(v4, "foo bar")
            _ = fi28
        }
        var v29 uint32 = 2
        _ = fi27
        _ = v29
    }
    v30 := GenS3{F0: (true || (-2.25 > 0.1)), F1: 42}
    var v31 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
