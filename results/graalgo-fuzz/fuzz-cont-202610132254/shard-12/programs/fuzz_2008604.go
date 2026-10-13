package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }
type GenS3 struct { F0 int8; F1 uint64 }

func main() {
    var v1 string = (("\n\t" + "") + "hello")
    v2 := make(map[string]bool)
    var v3 bool = (160 >= 2)
    fmt.Println(v1, v1)
    var v4 uint64 = 2
    switch -204 {
    case 32767:
        v5 := []int{686, -2147483648, -2147483648, 707}
        _ = v5
    case -598:
        v6 := GenS1{F0: 4294967295}
        _ = v6
    case 460:
        var v7 int64 = -9223372036854775808
        _ = v7
    }
    var v8 string = v1
    for fi9 := 0; fi9 < 7; fi9++ {
        v1 = ("" + ((v8 + "\n\t") + "\n\t"))
        var v10 uint = 65535
        v1 = "a"
        _ = fi9
        _ = v10
    }
    var v11 uint = 18446744073709551615
    fmt.Println(v1)
    for fi12 := 0; fi12 < 1; fi12++ {
        if v3 {
            v13 := GenS2{F0: 4294967295, F1: 1, F2: int32(v4)}
            var v14 float32 = (1.5 / -2.25)
            _ = v13
            _ = v14
        }
        _ = fi12
    }
    v15 := GenS1{F0: 4294967295}
    v2[(((("\n\t" + (("" + (((v1 + "a") + (v1 + v1)) + (v8 + "\n\t"))) + "a")) + v8) + (v1 + v8)) + "")] = v3
    fmt.Println(v11, v1, v4)
    v16 := []int{941, 4294967295, 32767, 2147483647}
    v17 := GenS3{F0: 100, F1: uint64(v11)}
    switch ((("foo bar" + (((((("\n\t" + (((v8 + (v1 + v8)) + ("\n\t" + (v8 + v1))) + v8)) + v8) + v1) + v1) + "\n\t") + "hello")) + v8) + v8) {
    case "hello":
        v18 := GenS2{F0: 2, F1: 113, F2: -1}
        _ = v18
    }
    if true {
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(v3, v4, v1)
            _ = fi19
        }
    } else {
        var v20 uint16 = 32767
        _ = v20
    }
    var v21 uint16 = 816
    v22 := v2[v1]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v8
}
