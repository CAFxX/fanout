package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 string) int8 {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{}
        for fi9 := 0; fi9 < 16; fi9++ {
            p0 = "hello"
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    p0 = (("hello" + p0) + p0)
    _ = p0
    return -29
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    var v3 uint32 = 2147483647
    v4 := []uint32{}
    var v5 uint32 = 0
    var v6 int64 = -231
    fmt.Println(v3, v3, v3)
    v10 := GenS1{F0: 0}
    v3 += (v3 >> ((42) & 63))
    var v11 int = 73
    if true {
        v1[("" + "foo bar")] = v11
        fmt.Println(v11, v3, v6)
    }
    var v12 int64 = (v6 & (v6 * (v6 & v6)))
    var v13 int64 = v12
    for fi14 := 0; fi14 < 5; fi14++ {
        _ = len(v4)
        fmt.Println(v6)
        if ((!(!((!((int64(v11) > (v13 - (v6 | (v12 * -9223372036854775808)))) && false)) || (711 <= -172)))) || (true && false)) {
            var v15 uint16 = 0
            var v16 uint16 = (^(v15))
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    var v17 string = "hello"
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS1{F0: genFn0("a")}
            _ = fi19
            _ = v20
        }
        delete(v1, ((v17 + (v17 + v17)) + "\n\t"))
        _ = fi18
    }
    if true {
        var v21 uint64 = 1
        var v22 bool = false
        fmt.Println(v17, v17)
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 12; fi23++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi23
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
