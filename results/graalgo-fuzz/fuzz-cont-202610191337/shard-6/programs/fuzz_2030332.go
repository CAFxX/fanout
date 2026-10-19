package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int64) string {
    p0 -= p0
    v6 := GenS1{F0: 333}
    var v7 int = 32767
    for fi8 := 0; fi8 < 3; fi8++ {
        p0 = (p0 / ((p0) | 1))
        v9 := []int{(-(v7))}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return (("hello" + ("foo bar" + "hello")) + "hello")
}

func main() {
    var v1 int8 = -128
    v2 := []int{249}
    var v3 uint16 = 808
    v4 := make(map[uint32]string)
    var v5 string = ""
    var v10 int64 = 42
    var v11 bool = true
    v1 = v1
    v12 := GenS1{F0: 42}
    var v13 uint16 = (v3 % ((65535) | 1))
    if (1 >= 493) {
        v1 = (^(v1))
    } else {
        _ = len(v4)
        for fi14 := 0; fi14 < 4; fi14++ {
            if (v1 < v1) {
                var v15 int64 = (v10 & -1)
                _ = len(v4)
                v5 = (("a" + (v5 + ((((v5 + ((v5 + (("a" + ("" + (v5 + "hello"))) + ((v5 + v5) + "hello"))) + ((("\n\t" + "") + v5) + v5))) + v5) + (v5 + v5)) + v5))) + v5)
                _ = v15
            }
            fmt.Println(v11, v5, v13)
            fmt.Println(v5, fi14)
            _ = fi14
        }
    }
    v16 := GenS1{F0: 42}
    v4[4294967295] = v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
