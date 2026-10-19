package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int16, p1 uint32, p2 uint32) int8 {
    v7 := []int{9223372036854775807, 128}
    if true {
        var v8 bool = (0 >= 0)
        fmt.Println(p0, p1, p2)
        _ = v8
    }
    p1 = (p1 >> ((658) & 63))
    var v9 uint64 = 42
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int64 = 186
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return -127
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 float64 = 1e10
    v4 := make(map[int]string)
    var v5 int8 = 100
    v6 := make(map[uint32]int)
    var v12 uint32 = 42
    if false {
        fmt.Println(v5, v5)
        for fi13 := 0; fi13 < 3; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                switch (("foo bar" + ("" + "")) + "foo bar") {
                case "a":
                    var v15 bool = true
                    _ = v15
                }
                delete(v6, (v12 << ((255) & 63)))
                var v16 int64 = 1
                _ = fi14
                _ = v16
            }
            _ = fi13
        }
        if (0.1 > -2.25) {
            var v17 int64 = 42
            _ = v17
        }
    } else {
        var v18 string = "hello"
        fmt.Println(v18)
        delete(v2, 2147483648)
        _ = v18
    }
    if (882 > -653) {
        var v19 uint64 = 4294967295
        if false {
            if true {
                v20 := []int{2147483647, 253, 32767, 127}
                _ = v20
            }
        }
        v21 := GenS1{F0: 0, F1: v5, F2: int16(v5)}
        _ = v19
        _ = v21
    }
    var v22 bool = ("\n\t" >= "\n\t")
    v23 := GenS2{F0: v5}
    if v22 {
        v3 += 1e10
        switch (("foo bar" + "foo bar") + ("hello" + "a")) {
        case "":
            v24 := []int{}
            _ = v24
        default:
            var v25 string = ("\n\t" + "a")
            _ = v25
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            fmt.Println(fi26, v12, v5)
            var v27 uint16 = 65535
            fmt.Println(fi26, v22)
            _ = fi26
            _ = v27
        }
    }
    var v28 int = 1
    for fi29 := 0; fi29 < 9; fi29++ {
        v5 = (v5 >> ((255) & 63))
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
