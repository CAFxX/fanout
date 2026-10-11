package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int }
func genFn0(p0 uint64, p1 uint16) uint8 {
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{fi7, (fi7 >> ((0) & 63))}
        p1 = p1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: int8(p1), F1: (1 >= 0), F2: 557}
    if false {
        v10 := []int{263}
        p1 = 65535
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = v10
    }
    var v11 int16 = -511
    var v12 uint32 = 1
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v9
    return 42
}
func genFn1(p0 uint) uint16 {
    p0 -= (p0 >> (((p0 ^ p0)) & 63))
    p0 -= (p0 >> ((p0) & 63))
    _ = p0
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 9223372036854775807
    var v3 string = ""
    v4 := make(map[string]int)
    var v5 string = (v3 + v3)
    var v6 string = v3
    var v13 uint32 = 1
    if true {
        delete(v1, v2)
        if (18446744073709551615 <= 18446744073709551615) {
            delete(v4, (((("hello" + (v3 + (v5 + "\n\t"))) + "") + (((("foo bar" + v3) + "hello") + (v6 + (v6 + ((("a" + v3) + v6) + v3)))) + "\n\t")) + (v3 + "hello")))
        }
        v14 := GenS2{F0: v2}
        _ = v14
    } else {
        v15 := GenS3{F0: v13}
        fmt.Println(v2)
        v16 := []int{}
        _ = v15
        _ = v16
    }
    v2 = (v2 << ((65535) & 63))
    v17 := GenS4{F0: 1, F1: v2}
    v18 := GenS3{F0: 2147483648}
    v19 := GenS3{F0: (v13 - 2147483647)}
    switch "a" {
    case "foo bar":
        if true {
            var v20 float64 = 3.14159
            _ = v20
        }
    case "a":
        v21 := []int{v2, v2, (v2 - (^(v2)))}
        _ = v21
    case "":
        v3 = "\n\t"
    }
    fmt.Println(v6)
    v6 = v3
    v3 = ("hello" + "hello")
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
