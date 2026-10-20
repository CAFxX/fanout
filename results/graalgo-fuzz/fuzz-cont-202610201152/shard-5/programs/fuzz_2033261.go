package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0(p0 int32, p1 uint16) int8 {
    if true {
        if (!(!(!false))) {
            p0 = (p0 << ((65535) & 63))
            v4 := []int{2147483647, 1}
            _ = len(v4)
            _ = v4
        } else {
            v5 := []int{1, 3}
            _ = v5
        }
    } else {
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := GenS1{F0: (156 != -32768)}
            p1 -= p1
            _ = fi6
            _ = v7
        }
        v8 := []int{}
        _ = v8
    }
    p1 = p1
    v9 := GenS1{F0: (173 <= -856)}
    _ = p0
    _ = p1
    _ = v9
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -1
    v3 := make(map[int]int)
    v10 := v3[4294967295]
    v11 := GenS1{F0: (42 >= v2)}
    v12 := []int{(v10 >> ((255) & 63)), 1}
    for fi13 := 0; fi13 < 15; fi13++ {
        switch fi13 {
        case -45:
            fmt.Println(v10, v10)
        case 4294967295:
            v2 = (v2 & 493)
        default:
            if false {
                v14 := GenS2{F0: 234, F1: 1, F2: 279}
                if len(v12) > 0 {
                    v12[0] = v10
                }
                switch "hello" {
                case "\n\t":
                    var v15 int16 = -490
                    _ = v15
                case "hello":
                    for fi16 := 0; fi16 < 15; fi16++ {
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = fi16
                    }
                case "foo bar":
                    fmt.Println(fi13)
                default:
                    var v17 uint16 = 42
                    _ = v17
                }
                _ = v14
            } else {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                v2 -= (v2 | v2)
                var v18 int16 = 723
                _ = v18
            }
        }
        var v19 int32 = 0
        fmt.Println(v2, fi13, v10)
        _ = fi13
        _ = v19
    }
    v20 := v1[(("" + "a") + ("a" + (("a" + ("hello" + (("hello" + (("foo bar" + "") + "\n\t")) + "a"))) + "\n\t")))]
    for fi21 := 0; fi21 < 17; fi21++ {
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 uint = 1
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    v20 = false
    v2 = v2
    switch (v10 & (v10 / ((v10) | 1))) {
    case -2147483648:
        v10 = (v10 >> ((1) & 63))
    case -661:
        v24 := GenS3{F0: 986}
        _ = v24
    default:
        var v25 bool = (v20 && (-2.25 > 0.0))
        _ = v25
    }
    var v26 int16 = 1
    switch "a" {
    case "hello":
        var v27 float32 = 1e10
        _ = v27
    case "\n\t":
        delete(v1, ("hello" + "foo bar"))
    case "":
        v3[v10] = (v10 | v10)
    }
    if (!(("a" + "foo bar") <= "\n\t")) {
        var v28 uint32 = 65535
        var v29 bool = v20
        fmt.Println(v26, v28)
        _ = v28
        _ = v29
    } else {
        v30 := GenS4{F0: v10, F1: 2}
        v10 = (v10 ^ 128)
        fmt.Println(v20)
        _ = v30
    }
    v2 = (v2 << ((672) & 63))
    fmt.Println(v10, v2, v26)
    v31 := v1["a"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v31
}
