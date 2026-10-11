package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint16, p1 int8) int {
    p1 = 45
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: p1, F1: 197}
        if true {
            switch "hello" {
            case "hello":
                v9 := []int{}
                _ = v9
            case "a":
                fmt.Println(p0, p1, p0)
            }
        }
        var v10 float64 = 3.14159
        _ = fi7
        _ = v10
        _ = v8
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int64 = -1
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    return -572
}

func main() {
    v1 := []string{"\n\t"}
    var v2 uint16 = 1
    v3 := make(map[string]bool)
    var v4 uint32 = 42
    var v5 int64 = -973
    v6 := make(map[int]uint32)
    var v13 uint16 = (v2 - v2)
    if false {
        for fi14 := 0; fi14 < 18; fi14++ {
            v5 = 0
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (int(v5) >= genFn0((v2 | 535), 100)) {
                if (!true) {
                    v1 = append(v1, "\n\t")
                    v2 += 0
                }
            } else {
                v15 := v6[-793]
                _ = v15
            }
            _ = fi14
        }
    } else {
        var v16 float32 = float32(1.5)
        v17 := GenS1{F0: -1, F1: 0}
        v16 = (0.1 * (((((v16 * v16) / 1.0) - v16) + 1e10) - float32(1e10)))
        _ = v16
        _ = v17
    }
    v4 -= uint32(v13)
    if true {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 uint = 0
            _ = fi18
            _ = v19
        }
        var v20 uint32 = v4
        var v21 int32 = -1
        _ = v20
        _ = v21
    } else {
        var v22 uint = 0
        _ = v22
    }
    delete(v3, "foo bar")
    fmt.Println(v4, v4)
    var v23 int32 = 2147483647
    var v24 float32 = 3.14159
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    switch 956 {
    case 4294967295:
        if len(v1) > 0 {
            v1[0] = "a"
        }
    case -2147483648:
        fmt.Println(v4)
    default:
        var v25 bool = (v4 >= v4)
        _ = v25
    }
    v26 := GenS1{F0: -1, F1: uint8(v2)}
    switch -775 {
    case 127:
        v2 -= 312
    case -18:
        if len(v1) > 0 {
            v1[0] = ""
        }
    }
    var v27 uint32 = (v4 / ((v4) | 1))
    for fi28 := 0; fi28 < 18; fi28++ {
        for fi29 := 0; fi29 < 19; fi29++ {
            v30 := GenS1{F0: 127, F1: 127}
            fmt.Println(v2, fi28, v13)
            _ = fi29
            _ = v30
        }
        if (v2 < 42) {
            for fi31 := 0; fi31 < 11; fi31++ {
                var v32 uint32 = (^(v27))
                v13 = (v2 >> ((18446744073709551615) & 63))
                if true {
                    v2 = (v13 >> ((2) & 63))
                    _ = len(v1)
                }
                _ = fi31
                _ = v32
            }
        }
        switch fi28 {
        case 1:
            v33 := GenS1{F0: 100, F1: 150}
            _ = v33
        }
        _ = fi28
    }
    if (false || (!false)) {
        var v34 bool = false
        var v35 int8 = 1
        _ = v34
        _ = v35
    }
    v36 := GenS1{F0: 0, F1: 1}
    var v37 int8 = 104
    var v38 int = -578
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
