package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int64, p1 int16, p2 string) uint16 {
    var v4 uint = 65535
    fmt.Println(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    var v5 uint16 = genFn0(v3, -1, (("\n\t" + (("a" + ("" + "\n\t")) + ("\n\t" + ("hello" + "foo bar")))) + "\n\t"))
    v5 = (v5 & genFn0((v3 - (v3 ^ -187)), 1000, "a"))
    fmt.Println(v5)
    v3 = (v3 - -9223372036854775808)
    var v6 int16 = -1
    var v7 uint64 = 1
    v7 = v7
    switch (("" + "foo bar") + "foo bar") {
    case "a":
        v8 := GenS1{F0: 867, F1: (v5 * 443), F2: 0}
        _ = v8
    case "\n\t":
        v9 := []int{-1, -785}
        _ = v9
    default:
        switch "hello" {
        case "a":
            v10 := GenS2{F0: 164}
            _ = v10
        case "":
            for fi11 := 0; fi11 < 14; fi11++ {
                if true {
                    v12 := []int{}
                    _ = v12
                }
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 uint16 = v5
                    var v15 int8 = 0
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi11
            }
        case "\n\t":
            v7 += 1
        }
    }
    fmt.Println(v5)
    var v16 int32 = -805
    v5 = v5
    fmt.Println(v16, v5, v5)
    fmt.Println(v3)
    v17 := GenS2{F0: 127}
    var v18 int16 = (^(v6))
    if true {
        if (!(!true)) {
            v5 += (^(v5))
            v19 := GenS2{F0: 2}
            v20 := v2[4294967295]
            _ = v19
            _ = v20
        }
        var v21 uint16 = genFn0((v3 >> ((255) & 63)), int16(v7), "")
        v22 := v2[0]
        _ = v21
        _ = v22
    }
    if (!(((1.5 - 1e10) / 0.1) > 1e10)) {
        v3 -= 340
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 uint16 = v5
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
