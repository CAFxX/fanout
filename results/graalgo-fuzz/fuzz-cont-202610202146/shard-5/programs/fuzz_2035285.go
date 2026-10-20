package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() string {
    if true {
        v7 := GenS1{F0: true}
        if ((0 < 1) && true) {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := GenS1{F0: true}
                var v10 uint8 = 127
                v11 := []int{-2147483648, (fi8 >> ((740) & 63)), fi8}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        }
        _ = v7
    }
    var v12 string = "\n\t"
    _ = v12
    return v12
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = -2.25
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    var v5 uint64 = 2
    var v6 uint64 = v5
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 int = -865
            v6 = 18446744073709551615
            v16 := GenS1{F0: false}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    var v17 int = 9223372036854775807
    fmt.Println(v6, v17)
    if false {
        _ = len(v4)
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v2, v6)
        var v19 string = "foo bar"
        _ = fi18
        _ = v19
    }
    var v20 int16 = -32768
    if false {
        delete(v1, 4294967295)
    } else {
        for fi21 := 0; fi21 < 15; fi21++ {
            v5 = (^(v5))
            _ = fi21
        }
        v22 := []int{32767}
        _ = v22
    }
    v23 := GenS1{F0: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
