package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 446
    var v3 uint32 = 2
    v4 := make(map[uint32]bool)
    fmt.Println(v3, v2)
    if (false && ((0.1 < ((0.0 / 1.0) * (0.1 - (1e10 - 1e10)))) && (65535 < (v2 << ((255) & 63))))) {
        v2 += v2
    }
    v5 := []int{32767, -1, 128, 128}
    switch ("hello" + "\n\t") {
    case "":
        switch ("foo bar" + "") {
        case "hello":
            var v6 int64 = 1
            _ = v6
        }
    case "\n\t":
        fmt.Println(v3)
    case "hello":
        if true {
            fmt.Println(v3, v3)
            var v7 string = (("hello" + ((("\n\t" + "\n\t") + "hello") + ("" + ("hello" + "\n\t")))) + ("hello" + "hello"))
            _ = v7
        }
    }
    delete(v4, (v3 % ((696) | 1)))
    for fi8 := 0; fi8 < 14; fi8++ {
        v3 = (v3 + (v3 * v3))
        _ = fi8
    }
    var v9 int = 4294967295
    v2 = v2
    v2 = 65535
    var v10 float64 = float64(1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
