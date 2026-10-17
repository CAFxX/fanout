package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0(p0 bool, p1 int64, p2 uint) int64 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint16 = 42
        if ((p0 && true) && p0) {
            if true {
                for fi7 := 0; fi7 < 12; fi7++ {
                    var v8 uint64 = 2
                    var v9 int64 = (p1 % ((p1) | 1))
                    v10 := []int{(fi5 + fi5), int(v6), fi5, (fi5 | (fi7 % ((615) | 1)))}
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
                var v11 int = fi5
                _ = v11
            }
        }
        _ = fi5
        _ = v6
    }
    var v12 uint32 = uint32(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p1
}
func genFn1(p0 int, p1 uint16) uint8 {
    var v13 uint8 = 2
    switch (((("\n\t" + "a") + (("foo bar" + ("hello" + "foo bar")) + ("hello" + ""))) + "") + "a") {
    case "hello":
        var v14 float64 = float64(-2.25)
        _ = v14
    case "":
        var v15 bool = false
        _ = v15
    case "foo bar":
        fmt.Println(p0, v13, p1)
    }
    _ = p0
    _ = p1
    _ = v13
    return (v13 * (v13 & 42))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 int32 = 42
    fmt.Println(v4)
    _ = len(v1)
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(fi16, v4)
        _ = fi16
    }
    v4 = -1
    fmt.Println(v4)
    var v17 float64 = 1e10
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 uint32 = 105
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 1, F1: 2147483647}
    if (566 > 65535) {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 bool = true
            _ = fi21
            _ = v22
        }
    }
    v4 -= 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
