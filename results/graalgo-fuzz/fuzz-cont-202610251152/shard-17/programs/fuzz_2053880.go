package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    var v7 bool = (!false)
    v7 = (9223372036854775807 >= 527)
    if (1 == 0) {
        var v8 bool = (!false)
        _ = v8
    }
    _ = v7
    return -1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 941
    v3 := make(map[string]bool)
    var v4 int = 42
    var v9 bool = true
    v10 := v3["foo bar"]
    var v11 string = "hello"
    v12 := []int{v4, 128, 9223372036854775807, v4}
    for fi13 := 0; fi13 < 7; fi13++ {
        switch "\n\t" {
        case "a":
            v10 = false
        default:
            var v14 int64 = -9223372036854775808
            _ = v14
        }
        _ = fi13
    }
    v15 := GenS1{F0: v4, F1: 4294967295}
    v9 = v9
    v10 = v9
    switch (("a" + "foo bar") + "a") {
    case "a":
        v9 = v9
    case "foo bar":
        if true {
            v16 := GenS2{F0: 18446744073709551615, F1: 758, F2: 2}
            v17 := GenS3{F0: 42, F1: int16(v4)}
            v10 = true
            _ = v16
            _ = v17
        } else {
            v10 = (1000 >= 1000)
        }
    default:
        v4 += v4
    }
    v4 = v4
    v18 := GenS3{F0: (v2 % ((v2) | 1)), F1: 32767}
    v9 = false
    switch "" {
    case "a":
        v11 = v11
    }
    var v19 int32 = 1
    v11 = "a"
    var v20 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
