package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0(p0 string) int16 {
    v7 := []int{128}
    p0 = "\n\t"
    _ = p0
    _ = v7
    return -32768
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{0}
    var v3 bool = (!false)
    v4 := make(map[string]uint32)
    v5 := []string{}
    var v6 string = v1
    v3 = (-22 != 278)
    var v8 int64 = -1
    fmt.Println(v3, v8, v6)
    switch 4294967295 {
    case -9223372036854775808:
        var v9 float64 = ((((0.0 / 3.14159) + 3.14159) - 1e10) * float64((1e10 + 0.1)))
        _ = v9
    case 0:
        var v10 int = -1
        _ = v10
    default:
        fmt.Println(v3, v3)
    }
    switch 251 {
    case -291:
        v1 = "\n\t"
    case 32768:
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS1{F0: 0, F1: -1}
            _ = fi11
            _ = v12
        }
    case 42:
        fmt.Println(v3)
    }
    if ((("foo bar" + "a") != (v1 + v6)) || (0 != 65535)) {
        var v13 uint8 = 128
        _ = v13
    }
    switch 2147483647 {
    case 9223372036854775807:
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint = 18446744073709551615
            _ = fi14
            _ = v15
        }
    default:
        fmt.Println(v8, v3, v1)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 = (v6 + "a")
    v6 = v1
    v16 := GenS1{F0: 4294967295, F1: -1}
    if false {
        v17 := GenS2{F0: 2, F1: 0}
        v4[v1] = 1
        _ = v17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS2{F0: 4294967295, F1: 371}
        _ = fi18
        _ = v19
    }
    fmt.Println(v8, v1, v8)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
