package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint8, p1 int32, p2 string) int64 {
    var v7 string = p2
    p2 = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 956
}

func main() {
    v1 := []uint32{83, 591, 975}
    v2 := make(map[string]int)
    var v3 uint32 = 4294967295
    var v4 uint64 = 42
    var v5 int64 = 0
    var v6 int16 = 32767
    var v8 uint32 = v3
    v3 = (v8 >> ((255) & 63))
    v3 = v3
    v6 += (v6 - v6)
    v3 += 65535
    v3 -= v3
    v4 += (v4 * (v4 >> ((557) & 63)))
    if len(v1) > 0 {
        v1[0] = v3
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            v2[("" + "a")] = (fi10 | fi9)
            var v11 int32 = -257
            _ = fi10
            _ = v11
        }
        fmt.Println(fi9)
        _ = fi9
    }
    var v12 int16 = v6
    var v13 uint8 = 42
    v14 := GenS1{F0: "hello", F1: (1e10 >= 1e10), F2: v8}
    if false {
        v15 := GenS2{F0: 0}
        _ = len(v2)
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int64 = genFn0(v13, 0, ((("" + "a") + "\n\t") + (("foo bar" + ("foo bar" + ("a" + "foo bar"))) + "a")))
        _ = fi16
        _ = v17
    }
    if false {
        var v18 uint16 = 65535
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 float64 = (0.1 * 0.1)
            v3 += v3
            _ = fi19
            _ = v20
        }
        fmt.Println(v8)
        _ = v18
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        v1 = append(v1, uint32(v4))
        v22 := GenS2{F0: 42}
        fmt.Println(v13, v5, v13)
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: ("hello" + "hello"), F1: true, F2: 377}
    for fi24 := 0; fi24 < 12; fi24++ {
        v12 += v12
        v4 -= v4
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
