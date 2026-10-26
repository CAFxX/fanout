package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint16) uint32 {
    var v4 bool = (2 > 0)
    v4 = false
    v4 = (int32(p0) == -1)
    _ = p0
    _ = v4
    return 969
}
func genFn1(p0 int32, p1 uint32, p2 int8) uint8 {
    var v5 int32 = p0
    p0 -= v5
    var v6 int8 = (p2 + p2)
    fmt.Println(p0, v5, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 uint64 = 42
    v2 := []int{1, 32768, 4294967295}
    v3 := []uint32{2}
    v1 -= v1
    var v7 int8 = 0
    var v8 int16 = -300
    var v9 uint16 = 716
    v7 -= (v7 + 104)
    switch -1 {
    case 0:
        if len(v3) > 0 {
            v3[0] = 47
        }
    case 128:
        if false {
            fmt.Println(v7, v7)
            switch (("" + (("foo bar" + ("foo bar" + ("hello" + "a"))) + ("a" + ("hello" + (("a" + (("hello" + ("\n\t" + ("\n\t" + "a"))) + "a")) + "foo bar"))))) + ("hello" + ("foo bar" + ""))) {
            case "hello":
                v2 = append(v2, 1)
            case "foo bar":
                v10 := GenS1{F0: -448}
                _ = v10
            }
        }
    case 767:
        switch 32767 {
        case 4294967295:
            v9 -= (v9 << ((0) & 63))
        case -1:
            v2 = append(v2, int(v8))
        case 336:
            if len(v2) > 0 {
                v2[0] = 32767
            }
        }
    default:
        var v11 float32 = 1e10
        _ = v11
    }
    _ = len(v3)
    v9 = (v9 % (((^(v9))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
