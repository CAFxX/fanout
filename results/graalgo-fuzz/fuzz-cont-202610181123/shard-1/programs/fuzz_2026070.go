package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
func genFn0(p0 int16) uint {
    var v6 int = 128
    var v7 int32 = 1
    _ = p0
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int16 = 32767
    var v2 int = 247
    v3 := []int{v2}
    v4 := make(map[uint32]int)
    v5 := []string{"foo bar", ("\n\t" + ("a" + "")), "\n\t", "a"}
    var v8 uint = genFn0(-1)
    var v9 int16 = (v1 % ((v1) | 1))
    v9 = 32767
    v10 := GenS1{F0: "a"}
    delete(v4, 278)
    var v11 int64 = 847
    for fi12 := 0; fi12 < 6; fi12++ {
        switch "" {
        case "hello":
            for fi13 := 0; fi13 < 17; fi13++ {
                v8 -= 65535
                var v14 int16 = (v9 & (v9 & (v9 % ((v1) | 1))))
                v8 += 65535
                _ = fi13
                _ = v14
            }
        }
        fmt.Println(v9, v11)
        _ = fi12
    }
    v11 -= v11
    var v15 int64 = v11
    for fi16 := 0; fi16 < 6; fi16++ {
        delete(v4, 821)
        _ = fi16
    }
    v17 := GenS2{F0: 1, F1: 32767, F2: "a"}
    v18 := GenS1{F0: "hello"}
    v19 := GenS1{F0: "hello"}
    fmt.Println(v15)
    for fi20 := 0; fi20 < 3; fi20++ {
        _ = len(v4)
        v21 := GenS1{F0: "a"}
        v8 = genFn0(1)
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
