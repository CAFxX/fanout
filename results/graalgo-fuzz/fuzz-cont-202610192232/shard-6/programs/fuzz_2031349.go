package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 string, p1 uint8) uint64 {
    var v4 float64 = 0.1
    fmt.Println(p0)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int = -2147483648
        var v7 uint8 = p1
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return 18446744073709551615
}

func main() {
    v1 := []string{}
    var v2 int64 = 42
    var v3 uint64 = 4294967295
    var v8 uint32 = 760
    if false {
        fmt.Println(v8, v2)
        var v9 float64 = (1e10 + -2.25)
        _ = v9
    }
    if (0 <= 32767) {
        if false {
            v8 -= v8
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v10 uint64 = (v3 >> ((4294967295) & 63))
    v2 -= int64(v8)
    for fi11 := 0; fi11 < 2; fi11++ {
        if false {
            v10 -= 500
            var v12 float64 = (((0.1 - 0.0) - (((0.0 - 0.0) - (((((0.0 * ((((0.0 * -2.25) / -2.25) / 1.0) / 3.14159)) - 1e10) / 1.0) * 0.0) - (1e10 + 1.5))) * (float64(0.1) + 1e10))) - (0.0 - (1e10 / 0.1)))
            _ = v12
        } else {
            v1 = append(v1, ("\n\t" + "hello"))
        }
        _ = fi11
    }
    v8 = v8
    fmt.Println(v8)
    if ((!true) || ((3.14159 == (((-2.25 / 0.1) - 0.0) * 0.0)) && ((float32((((3.14159 * (3.14159 / 1.0)) / 1e10) / 1e10)) * float32((float64(1.5) * 1.5))) == -2.25))) {
        for fi13 := 0; fi13 < 20; fi13++ {
            fmt.Println(v10, v8)
            v14 := GenS1{F0: 782}
            _ = fi13
            _ = v14
        }
    } else {
        if len(v1) > 0 {
            v1[0] = ("hello" + "foo bar")
        }
        v15 := GenS1{F0: 65535}
        v16 := GenS1{F0: 18446744073709551615}
        _ = v15
        _ = v16
    }
    var v17 float64 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
