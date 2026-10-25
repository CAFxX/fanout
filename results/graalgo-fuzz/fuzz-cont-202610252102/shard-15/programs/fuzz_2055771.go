package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int64 = 10
    var v2 uint16 = 41
    v3 := make(map[string]string)
    var v4 int32 = 2147483647
    v2 = (^(v2))
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int8 = 127
        v6 -= -128
        _ = fi5
        _ = v6
    }
    v4 -= 0
    if true {
        v7 := []int{32767, 42}
        if (65535 == 2) {
            v7 = append(v7, -131)
        } else {
            v4 -= -2147483648
        }
        var v8 int32 = (-(v4))
        _ = v7
        _ = v8
    }
    v1 += v1
    var v9 float64 = 0.0
    v10 := GenS1{F0: 115}
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            if false {
                var v13 string = "hello"
                v3[""] = v13
                fmt.Println(v2, v1)
                _ = v13
            }
            _ = fi12
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 int8 = 127
        _ = fi14
        _ = v15
    }
    v9 -= 0.1
    var v16 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
