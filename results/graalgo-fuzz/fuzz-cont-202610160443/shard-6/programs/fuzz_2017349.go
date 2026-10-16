package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 float32 = ((0.0 - 0.1) + float32(0.0))
    var v2 string = "hello"
    var v3 int = 32768
    var v4 uint32 = 468
    var v5 uint = uint(v4)
    v6 := make(map[int]string)
    if (true && (222 >= 797)) {
        switch -9223372036854775808 {
        case 1:
            var v7 bool = false
            _ = v7
        }
    }
    fmt.Println(v5, v4)
    var v8 float64 = 3.14159
    var v9 uint64 = 0
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(fi10)
        v11 := GenS1{F0: 42}
        _ = fi10
        _ = v11
    }
    if (71 > -1) {
        switch v2 {
        case "":
            fmt.Println(v3)
        }
        v3 = v3
    }
    if (float64(v1) >= 3.14159) {
        var v12 int = -600
        fmt.Println(v4)
        var v13 int = int(v4)
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint = 65535
        _ = fi14
        _ = v15
    }
    var v16 float64 = v8
    if false {
        v17 := []int{(v3 * v3), (v3 + v3), v3}
        v17 = append(v17, v3)
        var v18 uint8 = 255
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
