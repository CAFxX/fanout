package main

import (
    "fmt"
)


func main() {
    var v1 uint = 331
    v2 := make(map[uint32]bool)
    var v3 uint32 = 1
    var v4 bool = (!(!(true || true)))
    var v5 int16 = 32767
    var v6 float64 = 3.14159
    v7 := []int{}
    var v8 int32 = 779
    var v9 uint = v1
    v2[(^(v3))] = (127 < 0)
    var v10 int16 = v5
    var v11 float32 = 1.5
    for fi12 := 0; fi12 < 4; fi12++ {
        v9 = (v9 | (v9 >> (((v1 + uint(v3))) & 63)))
        switch "foo bar" {
        case "foo bar":
            if v4 {
                v2[359] = (42 < -90)
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
        }
        _ = fi12
    }
    var v13 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
