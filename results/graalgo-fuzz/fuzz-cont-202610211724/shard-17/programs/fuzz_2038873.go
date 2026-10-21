package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int64) uint32 {
    p0 -= p0
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(p0, fi7)
        if true {
            v8 := []int{1, -435, (fi7 + fi7)}
            _ = len(v8)
            _ = v8
        }
        _ = fi7
    }
    var v9 bool = true
    for fi10 := 0; fi10 < 17; fi10++ {
        fmt.Println(p0)
        var v11 bool = false
        v12 := []int{fi10, 128}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint64 = 2
    _ = p0
    _ = v13
    _ = v9
    return 635
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 uint16 = 32767
    var v4 uint8 = 1
    var v5 int = 0
    v6 := []string{"foo bar"}
    v4 += (v4 & 124)
    for fi14 := 0; fi14 < 16; fi14++ {
        if (203 == 996) {
            var v15 float64 = 0.0
            _ = v15
        }
        v16 := GenS1{F0: 42}
        _ = fi14
        _ = v16
    }
    _ = len(v1)
    if false {
        fmt.Println(v3, v5)
        var v17 int = -1
        _ = v17
    } else {
        fmt.Println(v3)
    }
    fmt.Println(v3, v3, v5)
    v1["hello"] = (!false)
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(v4)
        _ = fi18
    }
    fmt.Println(v5)
    var v19 int32 = -83
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 int16 = 622
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v21)
        case "foo bar":
            v22 := GenS1{F0: 18446744073709551615}
            _ = v22
        case "":
            v3 = uint16(v4)
        default:
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 int64 = 1
                v25 := GenS1{F0: 769}
                var v26 uint16 = (v3 + v3)
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
        }
        _ = fi20
        _ = v21
    }
    delete(v2, 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
