package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int32 }
func genFn0(p0 bool) uint8 {
    fmt.Println(p0)
    switch -337 {
    case 1:
        p0 = false
    }
    _ = p0
    return 2
}

func main() {
    v1 := []string{"foo bar"}
    var v2 float64 = -2.25
    v3 := []int{-907, -681, 4294967295}
    v4 := []string{}
    v2 -= float64((3.14159 / float32(1e10)))
    var v5 int = -2147483648
    v5 = 128
    v5 = (v5 << ((255) & 63))
    if len(v3) > 0 {
        v3[0] = 127
    }
    _ = len(v3)
    for fi6 := 0; fi6 < 4; fi6++ {
        if (0 >= 9223372036854775807) {
            if len(v1) > 0 {
                v1[0] = ""
            }
        }
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 uint32 = 0
            var v9 uint = 18446744073709551615
            var v10 uint8 = 26
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v11 := GenS1{F0: 1, F1: -1, F2: 1}
    v1 = append(v1, "hello")
    var v12 int16 = -446
    v3 = append(v3, 42)
    switch ((("foo bar" + "hello") + ("" + ("hello" + "a"))) + ("hello" + "foo bar")) {
    case "a":
        v12 = (v12 << ((255) & 63))
    }
    fmt.Println(v2, v12)
    var v13 int8 = 52
    v12 = 42
    v12 -= -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
