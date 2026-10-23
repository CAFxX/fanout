package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int64 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 bool; F1 int8 }
func genFn0(p0 int16) string {
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            switch "\n\t" {
            case "a":
                v9 := []int{(fi8 << ((1) & 63)), fi8}
                _ = v9
            case "foo bar":
                fmt.Println(fi8, fi7)
                var v10 uint16 = 2
                p0 -= p0
                _ = v10
            default:
                v11 := []int{(fi7 & fi7), fi7, fi7}
                _ = v11
            }
            v12 := GenS1{F0: true, F1: 2, F2: -1}
            var v13 int64 = -9223372036854775808
            _ = fi8
            _ = v12
            _ = v13
        }
        v14 := GenS1{F0: false, F1: 0, F2: -1}
        _ = fi7
        _ = v14
    }
    var v15 string = ("foo bar" + ("foo bar" + "\n\t"))
    _ = p0
    _ = v15
    return (("foo bar" + ((v15 + ((("a" + v15) + ((("a" + (("foo bar" + "") + "hello")) + ((v15 + v15) + v15)) + v15)) + "")) + ("hello" + "a"))) + ("\n\t" + "\n\t"))
}

func main() {
    var v1 float32 = 0.0
    var v2 bool = (-1 != -116)
    v3 := make(map[uint32]int)
    var v4 uint64 = 2
    var v5 int64 = 138
    var v6 bool = v2
    fmt.Println(v4)
    v16 := GenS1{F0: (-586 == -194), F1: 861, F2: v5}
    v1 += float32(0.1)
    switch (("\n\t" + "foo bar") + "hello") {
    case "\n\t":
        var v17 int = 32767
        _ = v17
    }
    fmt.Println(v2, v6, v6)
    _ = len(v3)
    v18 := GenS2{F0: int32(v5), F1: v2}
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := []int{fi19, fi19}
        v5 = 922
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: false, F1: 42}
    v5 = (^(v5))
    v2 = (!(882 != 492))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
