package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 bool) string {
    v5 := GenS1{F0: 29}
    p0 = false
    _ = p0
    _ = v5
    return ""
}
func genFn1(p0 uint16, p1 bool) bool {
    v6 := []int{-486, 1, -582}
    var v7 float32 = 1.5
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{128}
    var v3 uint16 = 423
    v4 := []string{((("\n\t" + "a") + ("foo bar" + ("" + "foo bar"))) + "\n\t"), "\n\t", "", ("\n\t" + "")}
    v8 := v1[65535]
    v9 := GenS1{F0: 2}
    v10 := GenS1{F0: 2147483647}
    v11 := GenS1{F0: 238}
    if true {
        v8 += (v8 >> ((18446744073709551615) & 63))
        if (true || ((42 != 1) && false)) {
            v8 = (v8 * (v8 + v8))
            fmt.Println(v8)
            var v12 uint64 = 1
            _ = v12
        }
        var v13 float64 = float64(3.14159)
        _ = v13
    }
    fmt.Println(v8, v8)
    v8 = (v8 ^ (v8 + v8))
    delete(v1, 752)
    var v14 int32 = -921
    delete(v1, 1)
    var v15 float64 = 0.1
    var v16 bool = (true || (true && (42 < -1)))
    v8 = -751
    v17 := GenS1{F0: 1}
    v18 := GenS1{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
