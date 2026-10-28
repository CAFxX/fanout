package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint, p1 int64, p2 bool) string {
    p2 = p2
    v6 := []int{-1, 2147483647, -761, 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "a"
}
func genFn1(p0 int32, p1 uint16, p2 int32) uint16 {
    var v7 uint8 = 5
    var v8 float32 = ((1.5 / 0.1) * -2.25)
    v9 := GenS1{F0: 1}
    if true {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: 402}
            var v12 bool = true
            if v12 {
                p1 = 197
                var v13 uint8 = (v7 ^ (^(v7)))
                var v14 string = "\n\t"
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v15 := GenS1{F0: 255}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p1 * p1)
}

func main() {
    var v1 int16 = -32768
    var v2 int8 = 31
    var v3 bool = true
    var v4 float32 = 1e10
    v5 := make(map[uint32]uint32)
    var v16 float64 = 1.5
    if false {
        for fi17 := 0; fi17 < 7; fi17++ {
            v3 = false
            v18 := GenS1{F0: 42}
            _ = fi17
            _ = v18
        }
        var v19 uint32 = 42
        _ = v19
    } else {
        fmt.Println(v1, v3, v1)
        v20 := []int{-1, 9223372036854775807}
        v21 := GenS1{F0: 1}
        _ = v20
        _ = v21
    }
    fmt.Println(v1)
    v22 := []int{1, 127, 9223372036854775807, 0}
    fmt.Println(v16)
    var v23 bool = false
    v2 = (v2 - (v2 * (v2 >> ((42) & 63))))
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    if (1 <= 42) {
        var v24 string = "\n\t"
        v23 = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 8; fi25++ {
            switch "foo bar" {
            case "hello":
                fmt.Println(v2, v1)
            case "a":
                var v26 uint16 = 255
                _ = v26
            case "foo bar":
                if v3 {
                    var v27 int32 = 2147483647
                    for fi28 := 0; fi28 < 2; fi28++ {
                        v1 = v1
                        _ = fi28
                    }
                    _ = v27
                }
            }
            _ = fi25
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            v30 := GenS1{F0: 42}
            _ = fi29
            _ = v30
        }
        v31 := GenS2{F0: v2}
        _ = v31
    }
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
