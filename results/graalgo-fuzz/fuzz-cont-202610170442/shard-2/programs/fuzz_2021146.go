package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 int64) int8 {
    var v7 uint64 = 4294967295
    v8 := GenS1{F0: -1, F1: 943}
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint32) bool {
    var v9 uint = 4294967295
    fmt.Println(v9, v9)
    v9 = v9
    for fi10 := 0; fi10 < 19; fi10++ {
        p0 = (p0 % (((p0 / (((p0 ^ p0)) | 1))) | 1))
        _ = fi10
    }
    _ = p0
    _ = v9
    return false
}

func main() {
    var v1 uint64 = 829
    v2 := []int{32767, -246, 307}
    var v3 int32 = 42
    var v4 int8 = 1
    var v5 float32 = (3.14159 + (3.14159 - 0.1))
    var v6 string = ("hello" + ("foo bar" + ""))
    var v11 uint32 = 79
    var v12 bool = (int8(v3) <= -115)
    v6 = "a"
    fmt.Println(v5, v6, v12)
    v13 := GenS2{F0: 42, F1: v4}
    v5 += 0.0
    fmt.Println(v6, v1, v11)
    v3 = (v3 | -615)
    v2 = append(v2, 9223372036854775807)
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 uint64 = v1
        _ = len(v2)
        _ = fi14
        _ = v15
    }
    if true {
        v4 = v4
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int32 = int32(v11)
            for fi18 := 0; fi18 < 5; fi18++ {
                fmt.Println(v3, v12)
                v19 := GenS1{F0: -1, F1: -756}
                _ = fi18
                _ = v19
            }
            if len(v2) > 0 {
                v2[0] = fi16
            }
            _ = fi16
            _ = v17
        }
        var v20 int16 = -651
        _ = v20
    } else {
        var v21 uint32 = 0
        v21 += (v11 ^ v11)
        for fi22 := 0; fi22 < 20; fi22++ {
            v21 = v21
            var v23 int = (fi22 >> ((4294967295) & 63))
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    v24 := GenS1{F0: (v3 - 42), F1: -432}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
