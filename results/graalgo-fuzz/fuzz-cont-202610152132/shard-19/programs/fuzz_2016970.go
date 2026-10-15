package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint, p1 uint32) bool {
    if true {
        var v5 uint32 = (p1 & (p1 * (^(p1))))
        _ = v5
    } else {
        if false {
            var v6 float32 = 3.14159
            var v7 uint = 42
            _ = v6
            _ = v7
        } else {
            p1 -= 42
            v8 := []int{int(p1), -2147483648, -807}
            _ = v8
        }
        v9 := []int{4294967295, 4294967295, 32768}
        _ = v9
    }
    v10 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v10
    return false
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 uint8 = 127
    var v4 int16 = -1
    v3 -= 0
    var v11 uint8 = (v3 | 199)
    var v12 uint8 = v3
    var v13 string = ("hello" + "foo bar")
    var v14 int32 = 0
    var v15 int = -117
    v14 += (v14 - v14)
    var v16 int16 = (v4 & (v4 % (((v4 + (-(v4)))) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
