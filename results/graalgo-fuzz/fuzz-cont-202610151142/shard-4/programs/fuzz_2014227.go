package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
func genFn0(p0 int16) int8 {
    fmt.Println(p0)
    var v4 int16 = -32768
    _ = p0
    _ = v4
    return 34
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{1, 173, 618}
    var v3 int64 = -21
    v5 := v1[711]
    v6 := GenS1{F0: genFn0(-245), F1: false}
    var v7 bool = (-892 > -968)
    var v8 int8 = -49
    var v9 bool = (!true)
    v3 += -23
    delete(v1, 2147483648)
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            if (v7 && (!((1e10 * (3.14159 * 1.5)) != -2.25))) {
                v9 = v9
                v1[42] = (v5 + ((("foo bar" + "hello") + (((("\n\t" + v5) + "") + "\n\t") + "hello")) + v5))
                v5 = v5
            }
            fmt.Println(v9, fi10, fi11)
            _ = fi11
        }
        v2 = append(v2, 146)
        var v12 bool = v9
        _ = fi10
        _ = v12
    }
    fmt.Println(v5)
    v3 = 42
    for fi13 := 0; fi13 < 12; fi13++ {
        fmt.Println(v5)
        _ = fi13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 string = v5
        _ = fi14
        _ = v15
    }
    fmt.Println(v8, v8, v5)
    v5 = v5
    v3 = (v3 << ((65535) & 63))
    var v16 float64 = 3.14159
    if (true || v9) {
        v5 = ("hello" + "")
    }
    fmt.Println(v9, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
