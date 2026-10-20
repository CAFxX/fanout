package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8, p1 int, p2 int) uint {
    p2 = (^(p2))
    fmt.Println(p2, p1)
    switch ((("a" + "") + ("foo bar" + ("a" + ("" + (("foo bar" + "a") + "a"))))) + "a") {
    case "foo bar":
        for fi4 := 0; fi4 < 9; fi4++ {
            fmt.Println(p0, fi4, fi4)
            fmt.Println(p1, p0, p0)
            p0 += -38
            _ = fi4
        }
    case "a":
        var v5 float32 = 1.5
        _ = v5
    }
    if (747 >= 156) {
        for fi6 := 0; fi6 < 3; fi6++ {
            p0 += 127
            _ = fi6
        }
        if false {
            v7 := []int{(p2 ^ p2)}
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(p1, p0)
                _ = fi8
            }
            _ = v7
        }
        fmt.Println(p0)
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint16 = 868
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    for fi9 := 0; fi9 < 17; fi9++ {
        if false {
            v1 = (^(v1))
            v10 := []int{(fi9 - (fi9 * (^(fi9))))}
            v11 := GenS1{F0: "foo bar", F1: -2147483648}
            _ = v10
            _ = v11
        }
        v1 += v1
        _ = fi9
    }
    var v12 uint16 = v1
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS1{F0: ("a" + ""), F1: 294}
        v15 := []int{fi13, (fi13 >> ((65535) & 63))}
        v16 := GenS2{F0: -384}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint = 4294967295
    _ = len(v3)
    var v18 int32 = 42
    delete(v2, 32767)
    if (!((0 >= 25) || (!("\n\t" >= ("foo bar" + "\n\t"))))) {
        var v19 string = ((("a" + "a") + "\n\t") + ("\n\t" + (("hello" + "a") + "hello")))
        switch v19 {
        case "\n\t":
            fmt.Println(v12)
        }
        _ = v19
    } else {
        var v20 int = 32768
        v20 += -843
        _ = v20
    }
    v12 = (v1 << (((v17 ^ (v17 - v17))) & 63))
    switch "" {
    case "a":
        fmt.Println(v1)
    case "foo bar":
        delete(v2, -542)
    case "\n\t":
        var v21 int = -835
        _ = v21
    default:
        for fi22 := 0; fi22 < 11; fi22++ {
            _ = len(v2)
            _ = fi22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
