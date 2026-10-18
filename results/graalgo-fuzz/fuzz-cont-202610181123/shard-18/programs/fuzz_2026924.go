package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8, p1 int32, p2 bool) uint8 {
    var v4 bool = p2
    v5 := GenS1{F0: -664}
    v4 = ("foo bar" < ("\n\t" + ("" + "hello")))
    fmt.Println(p0, p1)
    fmt.Println(p1, p0, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1() int32 {
    var v6 bool = false
    if (0 < 42) {
        v7 := GenS2{F0: -128, F1: v6}
        _ = v7
    } else {
        v8 := []int{-1}
        switch 32768 {
        case -14:
            if (((42 >= 2147483647) && (!(!v6))) || (1 == -9223372036854775808)) {
                v6 = v6
                var v9 int32 = 1
                _ = len(v8)
                _ = v9
            }
        default:
            v6 = true
        }
        _ = v8
    }
    _ = v6
    return -2147483648
}

func main() {
    var v1 uint16 = 0
    v2 := []int{}
    v3 := make(map[int]string)
    var v10 int = -2147483648
    var v11 uint32 = 112
    fmt.Println(v10, v11)
    switch ("a" + "foo bar") {
    case "hello":
        var v12 int = 127
        _ = v12
    default:
        v1 = (v1 & (^(v1)))
    }
    if true {
        var v13 float32 = 0.1
        _ = v13
    } else {
        var v14 int32 = genFn1()
        _ = v14
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v15 := GenS2{F0: 0, F1: (-9223372036854775808 <= 0)}
    _ = len(v2)
    fmt.Println(v11)
    fmt.Println(v11, v10)
    if true {
        v2 = append(v2, v10)
    } else {
        v16 := GenS3{F0: v11}
        _ = v16
    }
    v1 += (v1 * 32767)
    v2 = append(v2, 0)
    v11 += v11
    var v17 uint32 = 2147483647
    v18 := GenS1{F0: 625}
    var v19 bool = ("a" >= "foo bar")
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
