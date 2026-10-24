package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
func genFn0(p0 int64, p1 int64, p2 uint32) uint32 {
    var v7 bool = false
    v8 := GenS1{F0: -1, F1: 18446744073709551615}
    var v9 uint16 = 695
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 string) int32 {
    fmt.Println(p0)
    for fi10 := 0; fi10 < 19; fi10++ {
        p0 = "hello"
        if ((!true) && (!(255 > 42))) {
            var v11 uint = 65535
            var v12 int16 = 0
            _ = v11
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    return -1
}

func main() {
    var v1 uint8 = 113
    var v2 float32 = 0.0
    var v3 uint64 = 1
    v4 := make(map[uint32]int)
    var v5 int16 = -32
    v6 := []int{int(v5)}
    v2 -= v2
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS2{F0: ((("a" + ("foo bar" + "foo bar")) + ("\n\t" + ("" + "hello"))) + "foo bar"), F1: ((!("\n\t" != "hello")) || false), F2: 2147483648}
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: ((("\n\t" + "hello") + (("hello" + "\n\t") + ("hello" + "foo bar"))) + "a"), F1: false, F2: 434}
    fmt.Println(v1, v2, v1)
    v3 = v3
    fmt.Println(v3)
    _ = len(v4)
    for fi16 := 0; fi16 < 7; fi16++ {
        if len(v6) > 0 {
            v6[0] = (fi16 >> ((0) & 63))
        }
        var v17 int = (fi16 % (((fi16 << ((2) & 63))) | 1))
        v1 = 14
        _ = fi16
        _ = v17
    }
    v3 = 42
    v18 := GenS2{F0: ("a" + ("" + "")), F1: (0 > genFn1(("" + ((((("hello" + ("foo bar" + "foo bar")) + "\n\t") + "\n\t") + "hello") + ("\n\t" + "foo bar"))))), F2: 65535}
    v19 := GenS1{F0: -1, F1: (^(v3))}
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    var v20 int = 127
    if len(v6) > 0 {
        v6[0] = (v20 & -304)
    }
    var v21 int64 = -898
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
