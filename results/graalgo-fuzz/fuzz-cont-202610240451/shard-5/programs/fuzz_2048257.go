package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 string = "a"
    v4 := []uint32{65535}
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint8 = 42
            var v8 int8 = 127
            v3 = "foo bar"
            _ = fi6
            _ = v7
            _ = v8
        }
        v3 = v3
        var v9 uint = 65535
        _ = fi5
        _ = v9
    }
    v3 = v3
    var v10 int = 32767
    v3 = "\n\t"
    fmt.Println(v10, v3)
    fmt.Println(v10)
    var v11 string = "hello"
    for fi12 := 0; fi12 < 13; fi12++ {
        v11 = v3
        _ = fi12
    }
    if ((float64(((-2.25 - 0.0) / 1.0)) + 1e10) < -2.25) {
        var v13 float64 = 1.5
        _ = v13
    }
    v10 = (v10 % ((-852) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
