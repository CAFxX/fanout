package main

import (
    "fmt"
)


func main() {
    var v1 string = "a"
    var v2 uint8 = 2
    var v3 uint8 = v2
    v4 := make(map[uint32]uint32)
    v5 := []string{((v1 + "\n\t") + (v1 + v1)), ((v1 + "foo bar") + ("" + v1)), v1, "\n\t"}
    var v6 int32 = -2147483648
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        _ = len(v5)
        _ = fi7
    }
    var v8 uint = 1
    v3 = 127
    for fi9 := 0; fi9 < 2; fi9++ {
        v1 = v1
        _ = fi9
    }
    var v10 uint16 = 0
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            _ = len(v4)
            _ = fi12
        }
        v5 = append(v5, "")
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
