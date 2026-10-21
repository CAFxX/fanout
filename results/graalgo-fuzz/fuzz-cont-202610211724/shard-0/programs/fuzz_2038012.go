package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint8) int64 {
    v7 := []int{1000}
    switch 127 {
    case -919:
        p0 = p0
    case 32768:
        var v8 float32 = 0.1
        _ = v8
    case -1:
        if true {
            fmt.Println(p0)
            p0 = p0
            fmt.Println(p0)
        } else {
            p0 -= 32
            v9 := GenS1{F0: 0}
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: 0}
            var v12 bool = (!(((858 == 2) || (false || (fi10 >= fi10))) && (!false)))
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    if false {
        var v13 float64 = (-2.25 - (3.14159 - float64(-2.25)))
        _ = v13
    } else {
        _ = len(v7)
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 string, p1 uint) bool {
    p0 = p0
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint32 = 678
        var v16 uint32 = 65535
        _ = fi14
        _ = v15
        _ = v16
    }
    switch ((p0 + p0) + ("a" + ("foo bar" + ("foo bar" + p0)))) {
    case "a":
        v17 := GenS1{F0: 1}
        _ = v17
    case "\n\t":
        v18 := []int{0}
        _ = v18
    default:
        var v19 uint16 = 591
        _ = v19
    }
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 float32 = 1.5
    v5 := make(map[string]bool)
    var v6 uint32 = 2147483648
    var v20 uint64 = 139
    fmt.Println(v6)
    var v21 float32 = (v4 + ((float32((float64(v4) + 1.5)) + 1e10) - 3.14159))
    var v22 uint8 = 1
    v23 := []int{32767, 497, 126, -9223372036854775808}
    var v24 int32 = 0
    v4 -= v21
    var v25 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
