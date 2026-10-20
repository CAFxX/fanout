package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint64, p1 int8) uint16 {
    v7 := GenS1{F0: 0, F1: 42}
    if false {
        v8 := []int{}
        fmt.Println(p0, p1)
        _ = v8
    } else {
        p0 = (p0 & 0)
        if (2147483647 != -629) {
            if true {
                var v9 int16 = -726
                _ = v9
            }
        } else {
            var v10 float32 = (float32((0.1 * (0.1 + 1e10))) + 0.0)
            _ = v10
        }
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(p1, p0)
            var v12 int = fi11
            p1 = (p1 | p1)
            _ = fi11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 405
}

func main() {
    v1 := []int{-255}
    v2 := make(map[uint32]bool)
    var v3 uint32 = 115
    v4 := make(map[uint32]string)
    v5 := []string{}
    var v6 bool = true
    fmt.Println(v3)
    v3 = (v3 & v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v3 += v3
    for fi13 := 0; fi13 < 9; fi13++ {
        switch "a" {
        case "foo bar":
            v14 := GenS1{F0: 1000, F1: 42}
            _ = v14
        case "hello":
            v6 = v6
        default:
            if (32767 == 0) {
                var v15 uint32 = (v3 - (v3 / ((821) | 1)))
                _ = v15
            }
        }
        switch (^(fi13)) {
        case 4294967295:
            var v16 int32 = 0
            _ = v16
        case -9223372036854775808:
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint64 = 4294967295
                var v19 int8 = 0
                var v20 uint32 = v3
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        default:
            switch "hello" {
            case "a":
                v6 = v6
            case "\n\t":
                _ = len(v5)
            case "hello":
                if false {
                    var v21 uint = 65535
                    var v22 int = fi13
                    _ = v21
                    _ = v22
                } else {
                    var v23 uint64 = 111
                    _ = v23
                }
            }
        }
        _ = fi13
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v24 := GenS2{F0: -467}
    for fi25 := 0; fi25 < 2; fi25++ {
        fmt.Println(v3, fi25, fi25)
        delete(v4, v3)
        _ = fi25
    }
    var v26 float64 = 1.5
    var v27 float64 = v26
    switch "a" {
    case "a":
        v1 = append(v1, 9)
    case "\n\t":
        if ("foo bar" <= "hello") {
            fmt.Println(v3)
            v3 -= (^(v3))
            v27 += 1.5
        } else {
            var v28 uint32 = v3
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := GenS2{F0: 825}
                var v31 int = fi29
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
                _ = fi29
                _ = v30
                _ = v31
            }
            v28 = (v28 << ((83) & 63))
            _ = v28
        }
    case "":
        var v32 int8 = 127
        _ = v32
    }
    v6 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
