package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 string, p1 int, p2 string) uint64 {
    v7 := GenS1{F0: 42}
    v8 := []int{(^(p1)), (p1 % (((p1 | (p1 | 42))) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    var v4 uint8 = 127
    var v5 float32 = 1e10
    var v6 uint32 = 251
    _ = len(v2)
    fmt.Println(v4, v5)
    if (!true) {
        v5 = 1e10
        delete(v1, v6)
    } else {
        switch "hello" {
        case "foo bar":
            var v9 uint16 = 42
            _ = v9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := []int{(fi10 * (fi10 * int(v4))), (fi10 << ((739) & 63)), -780}
            v5 = v5
            if (-32768 != 1) {
                v5 += (((-2.25 + v5) / v5) + v5)
                var v12 int32 = 745
                _ = v12
            }
            _ = fi10
            _ = v11
        }
    }
    v2["\n\t"] = "a"
    if false {
        v13 := []int{771, 9223372036854775807}
        _ = v13
    } else {
        var v14 int8 = 0
        v4 = 2
        _ = v14
    }
    v2[("foo bar" + ("\n\t" + "\n\t"))] = ((("" + "\n\t") + "foo bar") + (("a" + "a") + "\n\t"))
    v5 = (v5 + 0.0)
    if false {
        v6 -= (v6 & (^(v6)))
        var v15 bool = ((-64 == -128) && (!(1 != 1)))
        _ = v15
    }
    var v16 float32 = (v5 + (v5 * float32((3.14159 - 1e10))))
    v17 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
