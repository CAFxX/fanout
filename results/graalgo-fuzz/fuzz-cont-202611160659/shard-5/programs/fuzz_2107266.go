package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint16 }
type GenS2 struct { F0 string; F1 int16 }
func genFn0(p0 int16) int16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        p0 = -1
        switch (("\n\t" + "a") + "\n\t") {
        case "a":
            for fi6 := 0; fi6 < 8; fi6++ {
                v7 := GenS1{F0: 0, F1: -1, F2: 621}
                var v8 int16 = (p0 / ((p0) | 1))
                p0 -= (v8 / ((1) | 1))
                _ = fi6
                _ = v7
                _ = v8
            }
        case "hello":
            p0 = -1
        case "foo bar":
            v9 := GenS1{F0: 42, F1: 69, F2: 896}
            _ = v9
        }
        _ = fi5
    }
    p0 = p0
    var v10 int32 = -206
    fmt.Println(v10, v10)
    v10 = 0
    _ = p0
    _ = v10
    return p0
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint = 966
    v3 := []uint32{65535, 4294967295}
    v4 := []uint32{}
    fmt.Println(v2, v2)
    _ = len(v4)
    v1 = -2147483648
    v11 := GenS2{F0: "foo bar", F1: -816}
    v1 -= 1
    v12 := GenS2{F0: (((("a" + "\n\t") + "") + ("foo bar" + "foo bar")) + (("\n\t" + ("foo bar" + "")) + "")), F1: genFn0(1)}
    var v13 uint64 = 65
    var v14 int32 = (v1 << (((v2 | v2)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
