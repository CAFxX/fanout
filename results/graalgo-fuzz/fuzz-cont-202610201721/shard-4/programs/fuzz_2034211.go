package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 bool, p1 int8, p2 int64) bool {
    p1 += (^(p1))
    v7 := []int{-9223372036854775808, -1, 32767}
    v7 = append(v7, 42)
    if p0 {
        var v8 string = ("\n\t" + ((("hello" + "\n\t") + "a") + "hello"))
        v9 := GenS1{F0: 10}
        var v10 uint16 = 0
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("hello" < "\n\t")
}
func genFn1(p0 uint64) bool {
    switch "foo bar" {
    case "":
        p0 = p0
    case "hello":
        v11 := GenS2{F0: 207, F1: 0, F2: 42}
        _ = v11
    }
    var v12 int32 = 0
    if true {
        var v13 int = 32767
        v13 -= v13
        var v14 uint64 = p0
        _ = v13
        _ = v14
    } else {
        if (true || false) {
            v15 := []int{-9223372036854775808}
            v12 = (v12 << ((62) & 63))
            _ = v15
        }
        var v16 uint16 = 32767
        p0 -= p0
        _ = v16
    }
    var v17 uint64 = (p0 & 148)
    _ = p0
    _ = v12
    _ = v17
    return false
}

func main() {
    var v1 uint32 = 2
    var v2 uint8 = 127
    v3 := make(map[int]int)
    var v4 uint8 = v2
    var v5 int64 = -9223372036854775808
    var v6 int = -9223372036854775808
    var v18 uint64 = 4294967295
    if true {
        delete(v3, v6)
        v1 = v1
    }
    v19 := v3[(-(v6))]
    var v20 int = v19
    v21 := []int{128, (v19 & v6), 1}
    switch "foo bar" {
    case "\n\t":
        v19 -= 0
    case "foo bar":
        v22 := GenS3{F0: (v6 & (v6 % ((v6) | 1))), F1: v4}
        _ = v22
    }
    var v23 int64 = (v5 >> ((207) & 63))
    var v24 int32 = -978
    if true {
        v2 = (v2 >> ((42) & 63))
        var v25 int32 = v24
        fmt.Println(v19)
        _ = v25
    }
    if true {
        var v26 uint16 = 65535
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
