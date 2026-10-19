package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint16 = 255
    v2 := []string{"hello", "", "\n\t"}
    var v3 uint32 = 2
    v4 := make(map[uint32]int)
    var v5 uint64 = 84
    v6 := []int{32768, -1}
    var v7 uint32 = (v3 - v3)
    var v8 float64 = -2.25
    fmt.Println(v1)
    var v9 int32 = -2147483648
    v10 := v4[42]
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: (-(v9))}
        _ = fi11
        _ = v12
    }
    v10 = v10
    for fi13 := 0; fi13 < 6; fi13++ {
        if true {
            _ = len(v4)
        }
        if (true || ((v5 <= v5) || true)) {
            var v14 string = "\n\t"
            fmt.Println(v5)
            var v15 bool = (true || true)
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    v8 -= v8
    fmt.Println(v5)
    var v16 uint = 0
    switch (v10 ^ v10) {
    case 32767:
        var v17 bool = false
        _ = v17
    case 95:
        switch ((("\n\t" + "foo bar") + "foo bar") + "") {
        case "a":
            v4[(v7 >> ((65535) & 63))] = 9223372036854775807
        }
    case -2147483648:
        v18 := GenS1{F0: v9}
        _ = v18
    default:
        _ = len(v6)
    }
    v19 := GenS1{F0: v9}
    v16 = (v16 * (v16 >> ((42) & 63)))
    var v20 uint64 = v5
    for fi21 := 0; fi21 < 14; fi21++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v22 uint = (v16 | v16)
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
