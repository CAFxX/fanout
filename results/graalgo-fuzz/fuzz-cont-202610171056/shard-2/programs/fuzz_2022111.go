package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint) string {
    v7 := GenS1{F0: 4294967295}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint64 = 688
        fmt.Println(p0)
        var v10 uint32 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint = p1
    var v12 int32 = 1
    v13 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v7
    return ""
}

func main() {
    v1 := []uint32{0, 2147483647, 965, 704}
    v2 := []uint32{2147483647, 2147483647, 356, 2147483648}
    v3 := []string{(("" + ("\n\t" + ("hello" + ""))) + ("foo bar" + "a")), "a", "\n\t", "hello"}
    var v4 bool = true
    v5 := make(map[int]uint32)
    var v6 int64 = 0
    v2 = append(v2, 1)
    for fi14 := 0; fi14 < 13; fi14++ {
        if len(v2) > 0 {
            v2[0] = 6
        }
        var v15 int = fi14
        _ = fi14
        _ = v15
    }
    v4 = (!true)
    switch "" {
    case "\n\t":
        v4 = v4
    }
    var v16 uint32 = 2147483647
    fmt.Println(v4, v6)
    v16 = 4294967295
    var v17 bool = v4
    for fi18 := 0; fi18 < 12; fi18++ {
        v6 += v6
        v4 = (v16 <= v16)
        _ = fi18
    }
    v16 += v16
    v4 = v4
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint8 = 46
        v20 = v20
        _ = fi19
        _ = v20
    }
    var v21 uint16 = 42
    if len(v1) > 0 {
        v1[0] = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
