package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 255
    v3 := []int{-965, 127, 0, 32767}
    v4 := make(map[string]bool)
    var v5 bool = (4294967295 == 65535)
    fmt.Println(v5, v5)
    var v6 int64 = 42
    var v7 int8 = 1
    v5 = v5
    var v8 int = -2147483648
    switch (v8 >> ((1) & 63)) {
    case 1:
        v9 := GenS1{F0: 2}
        _ = v9
    case 2147483647:
        fmt.Println(v5, v8)
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if true {
            fmt.Println(v8, v7)
        } else {
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 string = "hello"
                var v13 int8 = v7
                var v14 uint64 = uint64(v2)
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            v5 = (!(2 >= 18446744073709551615))
            var v15 uint16 = 42
            _ = v15
        }
        _ = fi10
    }
    v7 += (^(v7))
    v6 = v6
    v6 = int64(v8)
    v16 := v4["hello"]
    v17 := GenS1{F0: 18446744073709551615}
    var v18 float64 = 0.1
    v19 := GenS1{F0: 1}
    v7 += -1
    v5 = v5
    var v20 uint64 = 786
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
