package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    var v3 int = 42
    var v4 uint64 = 18446744073709551615
    var v5 float64 = 0.1
    fmt.Println(v3, v5)
    v3 = v3
    var v6 uint32 = 65535
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int64 = 9223372036854775807
        var v9 int64 = (v8 * (v8 << ((414) & 63)))
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = len(v2)
    switch ("a" + "") {
    case "a":
        for fi10 := 0; fi10 < 7; fi10++ {
            v3 = 32768
            _ = fi10
        }
    case "":
        var v11 uint32 = 42
        _ = v11
    case "foo bar":
        v5 += v5
    default:
        var v12 int64 = -9223372036854775808
        _ = v12
    }
    var v13 string = ""
    var v14 float64 = (3.14159 + ((v5 * ((((1e10 * v5) - float64(3.14159)) * v5) + (v5 / -2.25))) * v5))
    switch (v3 + v3) {
    case 4294967295:
        var v15 int8 = -8
        _ = v15
    default:
        switch -2147483648 {
        case 128:
            v16 := []int{}
            _ = v16
        case 127:
            v13 = ""
        case 42:
            v13 = "a"
        default:
            var v17 string = ("foo bar" + "foo bar")
            _ = v17
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
