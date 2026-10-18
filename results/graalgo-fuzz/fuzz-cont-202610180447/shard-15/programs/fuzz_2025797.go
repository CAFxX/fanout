package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0(p0 string, p1 uint32) bool {
    var v4 uint8 = uint8(p1)
    p1 += (p1 << ((uint(v4)) & 63))
    _ = p0
    _ = p1
    _ = v4
    return ((!((p1 * (p1 & p1)) > (p1 / ((111) | 1)))) && ("hello" <= p0))
}
func genFn1(p0 uint8) int8 {
    p0 += p0
    v5 := GenS1{F0: -2147483648, F1: 255}
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]bool)
    v3 := []string{("a" + "a"), "hello", (("" + "hello") + ("foo bar" + "\n\t"))}
    var v6 int32 = -947
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v7 int16 = -428
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v8 int = 9223372036854775807
    var v9 int64 = int64(v7)
    var v10 float32 = 1.5
    if true {
        v7 = v7
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 int32 = (v6 * int32(v8))
            v9 = 333
            _ = fi11
            _ = v12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
