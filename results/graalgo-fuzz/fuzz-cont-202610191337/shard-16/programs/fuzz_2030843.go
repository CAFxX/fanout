package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint32) string {
    p0 -= 2
    v5 := GenS1{F0: 127}
    _ = p0
    _ = v5
    return ""
}
func genFn1() uint32 {
    var v6 float32 = 0.1
    var v7 float32 = (v6 * v6)
    _ = v6
    _ = v7
    return 734
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]bool)
    v3 := []int{}
    v4 := []int{-2147483648, -224}
    var v8 string = "hello"
    var v9 float32 = 1.5
    switch (v8 + v8) {
    case "":
        for fi10 := 0; fi10 < 8; fi10++ {
            switch v8 {
            case "a":
                v1 -= 0.1
            case "foo bar":
                if false {
                    fmt.Println(v8, v1)
                    for fi11 := 0; fi11 < 10; fi11++ {
                        var v12 string = (v8 + v8)
                        v9 = float32(0.1)
                        _ = fi11
                        _ = v12
                    }
                }
            }
            _ = fi10
        }
    }
    var v13 bool = false
    if v13 {
        v14 := GenS1{F0: 128}
        if len(v4) > 0 {
            v4[0] = 32768
        }
        _ = v14
    }
    fmt.Println(v8, v1, v1)
    v8 = v8
    v1 -= 1e10
    var v15 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
