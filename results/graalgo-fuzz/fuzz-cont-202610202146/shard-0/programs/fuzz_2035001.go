package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 15; fi4++ {
        v1 = v1
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 int32 = -2147483648
            var v7 string = "foo bar"
            if true {
                switch v7 {
                case "\n\t":
                    if true {
                        fmt.Println(v6, v7)
                    }
                case "foo bar":
                    var v8 int = (fi5 % ((fi5) | 1))
                    _ = v8
                }
            } else {
                var v9 uint32 = 65535
                _ = v9
            }
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    v1 = (v1 * float32(3.14159))
    v1 -= (v1 - (v1 * (((v1 + 0.0) + (v1 * v1)) + ((float32(1.5) / (1e10 + (v1 - ((v1 * (v1 - 0.1)) - 1e10)))) / float32(3.14159)))))
    v10 := GenS1{F0: 42, F1: -86, F2: 65535}
    v1 -= v1
    v11 := GenS1{F0: 849, F1: -1, F2: 189}
    if (-94 >= -128) {
        v1 = float32(1.5)
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 bool = false
        v14 := GenS1{F0: 435, F1: 79, F2: 42}
        _ = fi12
        _ = v13
        _ = v14
    }
    delete(v3, 2147483648)
    fmt.Println(v1)
    _ = len(v2)
    v1 = v1
    var v15 int = 408
    fmt.Println(v1, v15)
    var v16 string = "hello"
    v1 += (v1 * -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
