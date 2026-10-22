package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 string }
func genFn0(p0 uint, p1 bool, p2 bool) string {
    p1 = (!false)
    for fi7 := 0; fi7 < 8; fi7++ {
        p2 = (!(!((!p2) || p1)))
        var v8 uint16 = 2
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    var v1 int8 = 40
    var v2 int8 = (v1 | v1)
    v3 := []uint32{65535}
    v4 := make(map[int]bool)
    v5 := []uint32{65535, 35, 153, 2147483648}
    v6 := []string{"", ("\n\t" + "hello")}
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(v1)
        v10 := GenS1{F0: "a", F1: -122, F2: "hello"}
        _ = fi9
        _ = v10
    }
    v1 -= (v1 ^ v2)
    for fi11 := 0; fi11 < 3; fi11++ {
        v2 = v1
        fmt.Println(fi11)
        v1 = 100
        _ = fi11
    }
    var v12 float64 = 1.5
    var v13 int32 = -2147483648
    fmt.Println(v2)
    v13 -= 42
    fmt.Println(v2, v1, v1)
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 uint = 4294967295
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            delete(v4, (-(fi15)))
            _ = fi15
            _ = v16
        }
        if true {
            fmt.Println(v13, v2)
        }
        _ = fi14
    }
    var v17 int16 = int16(v2)
    v2 += (v1 - v2)
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 int64 = -533
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
