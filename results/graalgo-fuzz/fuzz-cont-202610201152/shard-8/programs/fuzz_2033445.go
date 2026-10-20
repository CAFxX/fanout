package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 string, p1 bool) bool {
    if (-1 > -1) {
        p0 = ("\n\t" + "hello")
    } else {
        p0 = ("foo bar" + p0)
    }
    v7 := GenS1{F0: 341}
    var v8 float32 = (3.14159 / -2.25)
    switch p0 {
    case "a":
        v9 := GenS1{F0: 65535}
        _ = v9
    case "foo bar":
        v10 := GenS1{F0: 4294967295}
        _ = v10
    case "":
        v11 := GenS1{F0: 4294967295}
        _ = v11
    default:
        v12 := []int{1, 0}
        _ = v12
    }
    v13 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    _ = v8
    return ((42 > 0) && false)
}

func main() {
    var v1 int = -250
    v2 := make(map[int]int)
    var v3 string = ("\n\t" + "")
    v4 := []string{}
    var v5 bool = false
    var v6 int16 = 1000
    var v14 uint64 = 1
    _ = len(v4)
    var v15 uint32 = 156
    var v16 uint = 966
    var v17 int64 = 9223372036854775807
    var v18 int16 = v6
    var v19 int64 = -582
    var v20 int16 = (v18 << (((v16 % ((18446744073709551615) | 1))) & 63))
    v21 := GenS1{F0: (v16 * v16)}
    v5 = (v17 >= (v19 / ((int64(v16)) | 1)))
    var v22 int = (v1 >> (((v16 % ((v16) | 1))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
