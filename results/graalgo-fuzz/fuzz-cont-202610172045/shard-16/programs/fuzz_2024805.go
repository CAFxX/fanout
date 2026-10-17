package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int16; F2 uint8 }
func genFn0(p0 uint64, p1 int8, p2 uint16) int {
    var v6 bool = true
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32767
}

func main() {
    var v1 uint8 = 211
    var v2 int8 = -1
    v3 := make(map[string]string)
    v4 := []string{("\n\t" + "foo bar")}
    v5 := make(map[string]bool)
    v2 = v2
    delete(v3, (((("" + "\n\t") + (("foo bar" + ("foo bar" + "\n\t")) + "foo bar")) + "a") + ("" + ("foo bar" + "a"))))
    v1 -= (^(v1))
    v9 := GenS1{F0: -32768}
    v1 = (v1 << ((177) & 63))
    v2 = v2
    switch genFn0(130, 1, 0) {
    case 1:
        v10 := GenS1{F0: -37}
        _ = v10
    default:
        var v11 uint64 = 2
        _ = v11
    }
    var v12 uint16 = 255
    _ = len(v4)
    if (true && false) {
        v13 := GenS2{F0: 1, F1: 0, F2: 17}
        v14 := GenS2{F0: 146, F1: -32768, F2: 2}
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
