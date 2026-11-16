package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0() uint32 {
    var v6 string = "a"
    fmt.Println(v6)
    var v7 int = 127
    v7 += v7
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 0
    v3 := make(map[string]string)
    v4 := []string{"foo bar", "", "hello"}
    var v5 string = ""
    if (!false) {
        v5 = (v5 + v5)
        var v8 uint = 4294967295
        _ = v8
    }
    if true {
        var v9 uint8 = 2
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int8 = -1
            v12 := GenS1{F0: 2, F1: 4294967295}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    v3["foo bar"] = "foo bar"
    v2 -= v2
    var v13 int64 = 9223372036854775807
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            v13 = v13
            _ = fi15
        }
        if false {
            var v16 uint64 = 0
            _ = v16
        } else {
            var v17 int32 = 450
            _ = v17
        }
        _ = fi14
    }
    v5 = v5
    v18 := GenS2{F0: 843, F1: 1}
    fmt.Println(v13, v13, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
