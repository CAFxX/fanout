package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
func genFn0(p0 bool) uint8 {
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := GenS1{F0: (fi4 >> ((2) & 63)), F1: 18446744073709551615}
        _ = fi4
        _ = v5
    }
    v6 := []int{32768, 32767, 9223372036854775807}
    v6 = append(v6, -881)
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 bool, p1 uint16, p2 int8) int64 {
    fmt.Println(p2, p1, p2)
    p0 = (p0 && (!(-1 < 42)))
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 bool = (754 == 42)
    v2 := make(map[uint32]uint32)
    v3 := []string{"a", ("hello" + "foo bar"), "a", ("\n\t" + "hello")}
    v1 = false
    var v7 string = "\n\t"
    if v1 {
        var v8 string = v7
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := v2[2147483647]
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    var v11 uint8 = 54
    var v12 int64 = genFn1(true, 913, 127)
    v2[42] = 42
    if (1 != 42) {
        if v1 {
            var v13 uint8 = genFn0(true)
            for fi14 := 0; fi14 < 19; fi14++ {
                v13 -= v13
                _ = len(v3)
                _ = fi14
            }
            _ = v13
        }
        if v1 {
            var v15 int16 = int16(v12)
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(fi16)
        _ = fi16
    }
    var v17 string = v7
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
