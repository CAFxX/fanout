package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint64, p1 string) uint {
    p1 = (((p1 + "\n\t") + (p1 + ("\n\t" + p1))) + (p1 + p1))
    v7 := []int{701, 27, -9223372036854775808, 690}
    v7 = append(v7, 4294967295)
    if (42 <= -273) {
        p1 = "a"
        var v8 uint32 = 42
        fmt.Println(p0, p1, v8)
        _ = v8
    }
    var v9 bool = ((3.14159 - float64(0.0)) > 0.0)
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 65535
}
func genFn1() uint {
    var v10 string = "a"
    v10 = (((((("" + v10) + "a") + "hello") + "foo bar") + ("hello" + "")) + v10)
    fmt.Println(v10)
    var v11 uint64 = 2
    v12 := GenS1{F0: (4294967295 <= v11)}
    _ = v10
    _ = v11
    _ = v12
    return 458
}

func main() {
    var v1 int32 = 217
    var v2 float32 = 1e10
    v3 := make(map[int]uint32)
    var v4 float32 = 1e10
    var v5 bool = false
    var v6 string = "foo bar"
    _ = len(v3)
    fmt.Println(v5)
    v6 = "a"
    for fi13 := 0; fi13 < 4; fi13++ {
        v4 -= v4
        var v14 int32 = 120
        v15 := GenS1{F0: true}
        _ = fi13
        _ = v14
        _ = v15
    }
    switch -2147483648 {
    case -1:
        var v16 float64 = 1.5
        _ = v16
    case 42:
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 string = ((("\n\t" + "") + "a") + "\n\t")
            v2 += ((v2 / (v2 * 3.14159)) + v4)
            _ = fi17
            _ = v18
        }
    case 922:
        fmt.Println(v1, v5)
    default:
        var v19 int = 221
        _ = v19
    }
    var v20 int16 = 32767
    switch (v6 + (v6 + v6)) {
    case "a":
        v21 := []int{969}
        _ = v21
    }
    v22 := v3[0]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
