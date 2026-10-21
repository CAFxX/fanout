package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 string, p1 uint) uint32 {
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 0}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        if false {
            v10 := GenS1{F0: 79}
            _ = v10
        }
        _ = fi9
    }
    if true {
        v11 := []int{-687, -1, 1, -9223372036854775808}
        _ = v11
    } else {
        var v12 int16 = 0
        switch p0 {
        case "foo bar":
            if false {
                v12 += (v12 << (((^(p1))) & 63))
                for fi13 := 0; fi13 < 19; fi13++ {
                    p1 += p1
                    _ = fi13
                }
                for fi14 := 0; fi14 < 11; fi14++ {
                    p0 = p0
                    _ = fi14
                }
            }
        case "hello":
            p0 = ("a" + "foo bar")
        default:
            var v15 float64 = (((3.14159 / float64(float32(((3.14159 + (3.14159 / 1.0)) - 3.14159)))) - 1e10) - (float64(1e10) - float64(1e10)))
            _ = v15
        }
        var v16 uint32 = 0
        _ = v12
        _ = v16
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v4 uint8 = 207
    var v5 int64 = 42
    v6 := make(map[string]int)
    var v17 int8 = -1
    if false {
        delete(v1, 0)
        for fi18 := 0; fi18 < 7; fi18++ {
            v4 -= v4
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 int32 = 1
                v17 = (v17 >> ((1) & 63))
                v21 := GenS1{F0: v4}
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 int16 = 42
        var v24 int16 = v23
        var v25 bool = false
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        for fi27 := 0; fi27 < 16; fi27++ {
            v5 -= (v5 * v5)
            _ = fi27
        }
        _ = fi26
    }
    var v28 uint = 42
    fmt.Println(v5, v28)
    if true {
        for fi29 := 0; fi29 < 1; fi29++ {
            v17 = (v17 % ((-30) | 1))
            fmt.Println(v28)
            _ = fi29
        }
        var v30 float64 = (1.5 / (0.0 + (((0.1 - -2.25) + 0.1) + (0.1 / (0.1 * 1e10)))))
        fmt.Println(v17)
        _ = v30
    } else {
        var v31 int32 = 0
        _ = v31
    }
    var v32 int64 = v5
    fmt.Println(v28, v5)
    var v33 float64 = float64(-2.25)
    if true {
        var v34 uint64 = 936
        _ = v34
    }
    var v35 int64 = v5
    v17 += (v17 & (v17 / ((42) | 1)))
    var v36 int8 = (v17 & (-(v17)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
