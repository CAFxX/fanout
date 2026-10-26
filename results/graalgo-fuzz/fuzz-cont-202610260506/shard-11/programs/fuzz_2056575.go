package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
func genFn0(p0 int64, p1 uint, p2 uint16) uint64 {
    v4 := []int{723, 128, -594}
    v5 := GenS1{F0: -9223372036854775808, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int) int32 {
    var v6 int32 = 1
    v6 -= v6
    _ = p0
    _ = v6
    return (v6 + v6)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{32767, -1}
    var v3 int = 32768
    for fi7 := 0; fi7 < 8; fi7++ {
        switch "a" {
        case "\n\t":
            _ = len(v1)
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            if (-206 < 1000) {
                v9 := GenS1{F0: -9223372036854775808, F1: 2147483647}
                var v10 string = ("a" + (("hello" + "foo bar") + ("foo bar" + "hello")))
                _ = v10
                _ = v9
            }
            _ = fi8
        }
        _ = fi7
    }
    v11 := GenS1{F0: -1, F1: -2147483648}
    var v12 uint32 = 0
    switch (("\n\t" + ("hello" + "hello")) + (("foo bar" + ((("" + ("hello" + "a")) + "foo bar") + "\n\t")) + "a")) {
    case "a":
        var v13 uint = 255
        _ = v13
    case "foo bar":
        if (!(2 > 42)) {
            v14 := GenS1{F0: -9223372036854775808, F1: 961}
            var v15 bool = ((65535 < 0) && true)
            v12 += v12
            _ = v14
            _ = v15
        }
    }
    var v16 int16 = -1
    _ = len(v1)
    var v17 uint32 = (^(v12))
    v12 -= 2147483648
    v3 = (v3 << ((2) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
