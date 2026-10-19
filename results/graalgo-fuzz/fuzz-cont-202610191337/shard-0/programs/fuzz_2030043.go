package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int64 }
func genFn0(p0 string, p1 uint64) int32 {
    var v4 float32 = float32((1.5 - 0.0))
    v5 := GenS1{F0: 255, F1: 4294967295, F2: 1}
    v6 := []int{-575, 9223372036854775807}
    var v7 string = p0
    for fi8 := 0; fi8 < 18; fi8++ {
        if (128 != 1) {
            var v9 int32 = 1
            if len(v6) > 0 {
                v6[0] = (fi8 & fi8)
            }
            _ = v9
        }
        var v10 uint16 = 65535
        var v11 uint64 = (p1 / ((p1) | 1))
        _ = fi8
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := []string{("a" + "")}
    v2 := make(map[uint32]bool)
    v3 := []string{("" + (("\n\t" + ("foo bar" + "hello")) + "")), "a", "a"}
    var v12 int = 32768
    for fi13 := 0; fi13 < 19; fi13++ {
        switch (fi13 / ((fi13) | 1)) {
        case -1:
            v14 := GenS1{F0: 4294967295, F1: 610, F2: 837}
            _ = v14
        case -103:
            var v15 string = (("" + (("hello" + "foo bar") + ("\n\t" + ("\n\t" + "a")))) + "foo bar")
            _ = v15
        case -180:
            var v16 int64 = 9223372036854775807
            _ = v16
        default:
            fmt.Println(fi13)
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            fmt.Println(v12)
            var v18 int8 = 42
            _ = fi17
            _ = v18
        }
        v1 = append(v1, "")
        _ = fi13
    }
    var v19 bool = ((genFn0("a", 18446744073709551615) <= 42) && true)
    var v20 bool = v19
    switch v12 {
    case 823:
        if true {
            var v21 int8 = 127
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(v12)
                v12 = fi22
                switch "hello" {
                case "hello":
                    _ = len(v1)
                case "\n\t":
                    v2[4294967295] = (9223372036854775807 >= -822)
                }
                _ = fi22
            }
            _ = v21
        }
    }
    fmt.Println(v20)
    v3 = append(v3, "foo bar")
    if (false && (-128 != -83)) {
        v12 = 9223372036854775807
        v12 += (v12 * (v12 % ((v12) | 1)))
    } else {
        var v23 int = v12
        v2[65535] = false
        _ = v23
    }
    switch "\n\t" {
    case "a":
        var v24 int32 = genFn0("hello", 18446744073709551615)
        _ = v24
    case "":
        var v25 uint16 = 641
        _ = v25
    case "hello":
        var v26 int = v12
        _ = v26
    default:
        switch 32767 {
        case 9223372036854775807:
            fmt.Println(v20, v19, v20)
        default:
            var v27 int16 = 32767
            _ = v27
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
