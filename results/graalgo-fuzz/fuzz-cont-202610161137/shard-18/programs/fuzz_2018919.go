package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8) uint8 {
    v6 := []int{-476, 128, 127, -9223372036854775808}
    var v7 string = "a"
    p0 = (^(p0))
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1() int16 {
    var v8 uint16 = 1
    if (-273 != 0) {
        if true {
            v8 += 65535
        } else {
            v8 = 42
            v9 := []int{}
            v8 = (v8 >> ((4294967295) & 63))
            _ = v9
        }
        v8 = (^(v8))
        switch "hello" {
        case "foo bar":
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := []int{fi10, fi10, (-(fi10)), fi10}
                v8 -= 32767
                _ = fi10
                _ = v11
            }
        case "a":
            v8 += v8
        default:
            v8 += (^(v8))
        }
    }
    _ = v8
    return 32767
}

func main() {
    var v1 string = "a"
    var v2 int64 = 1
    var v3 int64 = v2
    var v4 string = v1
    var v5 uint16 = 754
    v1 = v1
    v12 := []int{128, 768, 848, 42}
    fmt.Println(v5, v2, v3)
    if false {
        v4 = v1
        v13 := GenS1{F0: genFn1(), F1: 4294967295, F2: 9223372036854775807}
        if true {
            var v14 uint32 = 718
            _ = v14
        }
        _ = v13
    }
    fmt.Println(v3, v3, v2)
    v15 := GenS1{F0: genFn1(), F1: 32768, F2: 127}
    v3 += (v2 / (((-(v2))) | 1))
    v16 := GenS2{F0: 65535, F1: (v5 ^ v5)}
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(fi17, fi17)
        v5 += v5
        _ = fi17
    }
    var v18 int8 = 42
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v19 := GenS3{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
