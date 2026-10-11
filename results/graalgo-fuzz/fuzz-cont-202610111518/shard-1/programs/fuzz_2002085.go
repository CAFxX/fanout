package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 0
    var v3 string = "\n\t"
    var v4 uint8 = 255
    v2 = 42
    var v5 int8 = v2
    fmt.Println(v5, v5, v4)
    v5 += (v5 >> ((18446744073709551615) & 63))
    v2 = -109
    if false {
        v2 = v5
        v2 = 51
        v6 := []int{}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            if false {
                var v8 int8 = v5
                v5 = (^(v2))
                _ = v8
            }
            _ = fi7
        }
        var v9 uint16 = 42
        _ = v9
    }
    v10 := v1[32767]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
