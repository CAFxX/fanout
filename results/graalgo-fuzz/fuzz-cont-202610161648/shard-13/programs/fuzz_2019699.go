package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 0
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 uint8 = 2
    v6 := make(map[uint32]string)
    v5 = (^(v5))
    for fi7 := 0; fi7 < 6; fi7++ {
        switch (((("hello" + (("a" + ("a" + "")) + "")) + "foo bar") + "a") + "\n\t") {
        case "foo bar":
            v2 = v2
        case "a":
            var v8 string = "hello"
            _ = v8
        }
        switch ("a" + "") {
        case "a":
            var v9 float32 = 1e10
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{4294967295}
    fmt.Println(v5, v2)
    fmt.Println(v2, v2)
    v2 = (v2 >> ((18446744073709551615) & 63))
    var v11 int16 = 1
    var v12 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
