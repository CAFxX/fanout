package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
func genFn0(p0 bool, p1 uint16) int64 {
    v5 := []int{1, 127, 32767, 4294967295}
    var v6 uint32 = 4294967295
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 uint32 = (v6 & 2147483647)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 464
}
func genFn1(p0 string, p1 int64) string {
    var v8 uint = 42
    v9 := GenS1{F0: 29, F1: 715}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 uint16 = 65535
    v4 := make(map[int]string)
    if (!(!(4294967295 == 65535))) {
        v3 = v3
        v3 = (v3 % ((1) | 1))
    }
    var v10 bool = false
    fmt.Println(v10)
    if (!v10) {
        fmt.Println(v10, v10)
        fmt.Println(v10, v10)
    } else {
        if v10 {
            v11 := GenS1{F0: 42, F1: -794}
            _ = v11
        }
        var v12 int = 4294967295
        _ = v12
    }
    if ((((((float32(((3.14159 - 1.5) - 1.5)) / 1.0) / 3.14159) + float32((-2.25 - 1e10))) - 0.1) - ((3.14159 / float32(((3.14159 - 1e10) * 3.14159))) - 0.0)) <= 1.5) {
        for fi13 := 0; fi13 < 1; fi13++ {
            v10 = (true && false)
            _ = fi13
        }
        var v14 uint32 = 42
        v15 := []int{0, -653, 142}
        _ = v14
        _ = v15
    }
    var v16 uint16 = v3
    var v17 int = 32767
    for fi18 := 0; fi18 < 7; fi18++ {
        v16 += (v16 - (v16 - 234))
        _ = fi18
    }
    if ((-(v17)) != (-(v17))) {
        v16 -= 65535
    } else {
        var v19 string = (("a" + ("hello" + "foo bar")) + "foo bar")
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 int16 = 454
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
