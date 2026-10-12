package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint8, p1 bool) uint32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: ("foo bar" + "foo bar")}
        _ = fi7
        _ = v8
    }
    var v9 string = "\n\t"
    v10 := []int{}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}

func main() {
    v1 := []uint32{4294967295, 184, 608}
    v2 := []string{}
    var v3 uint64 = 213
    var v4 int32 = 42
    v5 := make(map[string]bool)
    v6 := []int{1}
    delete(v5, "foo bar")
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    if (true && (!false)) {
        v4 -= -1
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v11 bool = (472 >= 42)
        _ = v11
    } else {
        if false {
            v12 := GenS2{F0: 0, F1: -792}
            var v13 int = 919
            _ = v12
            _ = v13
        }
        v3 = v3
    }
    if ((("a" + "foo bar") + "hello") == ("hello" + ("\n\t" + "foo bar"))) {
        var v14 uint16 = 828
        v2 = append(v2, "a")
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v4, v3)
        _ = fi15
    }
    v16 := v5["foo bar"]
    var v17 bool = false
    var v18 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
