package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int8, p1 uint8, p2 string) string {
    if (true && true) {
        var v6 uint32 = 4294967295
        p1 += 255
        _ = v6
    } else {
        v7 := GenS1{F0: 0}
        _ = v7
    }
    v8 := []int{9223372036854775807, 127, 32768}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v9 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return p2
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1.5
    v3 := make(map[string]int)
    v4 := []uint32{}
    v5 := []int{32767}
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v12 int = fi10
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 32767}
        _ = fi10
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 string = (("\n\t" + "") + "\n\t")
        _ = fi14
        _ = v15
    }
    fmt.Println(v2)
    v2 -= 1.5
    v2 += 1.5
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int64 = 507
        v18 := GenS1{F0: 123}
        switch 482 {
        case 729:
            v2 += v2
        case 32768:
            fmt.Println(v17)
        case 32767:
            var v19 int16 = -940
            _ = v19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3[""] = -9223372036854775808
    v2 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
