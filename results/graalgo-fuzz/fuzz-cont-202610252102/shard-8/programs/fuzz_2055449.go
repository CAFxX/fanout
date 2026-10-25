package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint8, p1 uint32, p2 uint16) int16 {
    p1 -= 4294967295
    v6 := GenS1{F0: 880}
    v7 := []int{0, 42}
    v8 := GenS1{F0: 420}
    v9 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint64 = 247
    v2 := []int{42, -9223372036854775808}
    var v3 bool = true
    var v4 uint64 = v1
    v5 := []uint32{213, 4294967295, 4294967295}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v10 string = "a"
    var v11 int8 = -1
    var v12 float64 = 0.1
    if false {
        v13 := GenS1{F0: -1}
        fmt.Println(v3, v11, v12)
        _ = v13
    } else {
        v14 := GenS1{F0: -802}
        if v3 {
            if len(v5) > 0 {
                v5[0] = 259
            }
        } else {
            v1 = v4
        }
        v15 := GenS1{F0: 9223372036854775807}
        _ = v14
        _ = v15
    }
    v3 = ((998 <= -392) && (0 < 18446744073709551615))
    v3 = v3
    var v16 uint = 807
    if false {
        var v17 uint64 = 483
        if v3 {
            v16 = 937
            var v18 uint64 = (v17 * (v17 | (^(v1))))
            _ = v18
        } else {
            v19 := GenS1{F0: -9}
            _ = v19
        }
        fmt.Println(v10, v10, v17)
        _ = v17
    } else {
        var v20 int32 = 1
        _ = v20
    }
    switch (v10 + (v10 + v10)) {
    case "hello":
        var v21 int32 = 204
        _ = v21
    case "a":
        var v22 string = (v10 + ((("a" + (("a" + "a") + (("a" + v10) + v10))) + ((v10 + "hello") + ((((v10 + ("" + "hello")) + v10) + (v10 + "")) + "\n\t"))) + "a"))
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
