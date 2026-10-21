package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int }
func genFn0(p0 uint64, p1 bool, p2 uint32) uint16 {
    for fi6 := 0; fi6 < 8; fi6++ {
        switch "hello" {
        case "a":
            fmt.Println(p2, fi6)
        }
        _ = fi6
    }
    v7 := GenS1{F0: (p2 + 2147483647), F1: (1000 == 1000), F2: 128}
    p2 = 65535
    switch -112 {
    case 32768:
        p0 -= (p0 >> ((4294967295) & 63))
    }
    var v8 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 783
}
func genFn1() uint8 {
    var v9 float64 = 0.0
    for fi10 := 0; fi10 < 15; fi10++ {
        switch 9223372036854775807 {
        case 4294967295:
            var v11 uint64 = 412
            _ = v11
        }
        fmt.Println(v9, fi10)
        v9 = (float64((1.5 - 3.14159)) * float64((1e10 / (((1.5 / 3.14159) * 3.14159) * 0.1))))
        _ = fi10
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := []int{(fi13 | fi12), fi12, fi12}
            v9 -= (float64(1e10) - (float64(1.5) + 0.1))
            _ = fi13
            _ = v14
        }
        switch (^(fi12)) {
        case 127:
            fmt.Println(v9, fi12)
        case 0:
            var v15 int = (fi12 & (fi12 + (-(fi12))))
            _ = v15
        case -333:
            if false {
                v9 = 1.5
                v16 := []int{}
                _ = v16
            }
        }
        var v17 float64 = (0.0 * (v9 / (float64((0.1 / 3.14159)) - float64(0.1))))
        _ = fi12
        _ = v17
    }
    v18 := []int{2147483647, 0, 1}
    v18 = append(v18, -1)
    _ = v18
    _ = v9
    return 127
}

func main() {
    var v1 int = -880
    var v2 float32 = -2.25
    var v3 uint16 = 3
    var v4 int32 = -1
    v5 := make(map[uint32]uint32)
    var v19 float32 = v2
    var v20 string = "foo bar"
    v21 := GenS1{F0: 4294967295, F1: false, F2: (v1 | 32768)}
    if true {
        var v22 bool = false
        var v23 bool = v22
        var v24 int16 = 42
        _ = v22
        _ = v23
        _ = v24
    } else {
        var v25 uint8 = 128
        v26 := GenS1{F0: 65535, F1: false, F2: 1}
        if (1 != -1) {
            v20 = "foo bar"
        }
        _ = v25
        _ = v26
    }
    v27 := []int{v1}
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int64 = 0
        for fi30 := 0; fi30 < 1; fi30++ {
            v31 := GenS1{F0: 2147483648, F1: true, F2: fi28}
            v32 := v5[65535]
            v3 -= 0
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    v19 -= (0.1 + (float32(-2.25) - v2))
    if false {
        v3 = (^(v3))
    } else {
        for fi33 := 0; fi33 < 6; fi33++ {
            if ((127 != 128) || (true || true)) {
                var v34 uint16 = (^(v3))
                _ = v34
            }
            _ = len(v27)
            _ = fi33
        }
        v1 = 0
        var v35 int16 = -879
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
