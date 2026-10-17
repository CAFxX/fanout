package main

import (
    "fmt"
)

func genFn0(p0 uint16) bool {
    var v7 int16 = 0
    var v8 uint16 = p0
    var v9 float64 = 0.1
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 int8, p1 uint32) int64 {
    var v10 int64 = -1
    for fi11 := 0; fi11 < 17; fi11++ {
        p1 += 957
        for fi12 := 0; fi12 < 15; fi12++ {
            p1 = (p1 ^ 4294967295)
            _ = fi12
        }
        var v13 int = (fi11 % ((32767) | 1))
        _ = fi11
        _ = v13
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v10
    return (^(v10))
}

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]bool)
    v3 := []string{("hello" + ("a" + "hello")), (("\n\t" + "\n\t") + "hello"), ("hello" + "a")}
    var v4 uint64 = 4294967295
    var v5 int16 = 0
    v6 := make(map[uint32]bool)
    var v14 int16 = v5
    var v15 string = ("\n\t" + (("" + ("\n\t" + "")) + ("foo bar" + "hello")))
    fmt.Println(v4, v4)
    v15 = (("a" + (v15 + v15)) + v15)
    v5 -= (v14 % ((700) | 1))
    v16 := v2[351]
    v5 = (v14 << (((v1 + (v1 - (v1 << (((v1 ^ v1)) & 63))))) & 63))
    v14 -= (v5 ^ (v14 & v14))
    var v17 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
