package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() string {
    var v7 string = ""
    v8 := []int{415}
    var v9 int16 = 26
    fmt.Println(v7, v7)
    v7 = "foo bar"
    _ = v7
    _ = v8
    _ = v9
    return ((v7 + ("" + ("" + v7))) + "foo bar")
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    var v4 float64 = -2.25
    v5 := make(map[string]bool)
    var v6 uint16 = 432
    v10 := GenS1{F0: 42}
    var v11 int16 = 42
    var v12 int8 = 100
    var v13 bool = (v6 < v6)
    var v14 float32 = (float32(v4) - 0.0)
    v11 = 32767
    v12 = v12
    var v15 uint8 = 155
    var v16 int64 = -601
    v15 = (v15 & (^(v15)))
    if true {
        var v17 int32 = 1
        _ = v17
    }
    var v18 int16 = (v11 >> ((18446744073709551615) & 63))
    v19 := v2["foo bar"]
    var v20 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
