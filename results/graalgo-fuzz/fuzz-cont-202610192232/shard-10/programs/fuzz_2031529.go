package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
func genFn0(p0 int32, p1 int, p2 string) string {
    p1 = (-(p1))
    var v5 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    var v1 uint32 = 4294967295
    v2 := []string{("\n\t" + "foo bar"), (("hello" + "a") + "\n\t")}
    v3 := make(map[string]string)
    var v4 uint32 = (^(v1))
    var v6 uint64 = 113
    var v7 int = 127
    fmt.Println(v6, v6, v7)
    delete(v3, ("hello" + ""))
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 bool = (false || (960 == 482))
        v10 := v3[("hello" + "foo bar")]
        if (255 != 255) {
            _ = len(v3)
            var v11 uint64 = 4294967295
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    delete(v3, (((((("\n\t" + ("" + "\n\t")) + (("foo bar" + "foo bar") + (("\n\t" + "\n\t") + (("\n\t" + (("" + "\n\t") + "")) + "a")))) + ("hello" + "")) + "a") + "foo bar") + "hello"))
    v12 := GenS1{F0: 42, F1: (-104 <= 42), F2: false}
    if (1 > -1) {
        var v13 int32 = -2147483648
        var v14 uint8 = 42
        _ = v13
        _ = v14
    } else {
        if false {
            var v15 float32 = 3.14159
            v4 += uint32(v6)
            _ = v15
        }
        var v16 uint32 = (^(v1))
        _ = v16
    }
    delete(v3, "\n\t")
    v4 -= (v1 >> ((252) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
