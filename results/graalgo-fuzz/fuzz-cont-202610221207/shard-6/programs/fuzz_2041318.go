package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 1
    var v2 float64 = 1e10
    v3 := []string{"", "\n\t", "\n\t", "hello"}
    v4 := make(map[string]int)
    v5 := make(map[string]int)
    v2 = v2
    fmt.Println(v1, v2)
    v2 = v2
    fmt.Println(v1)
    if true {
        if ((v1 - 32767) < 42) {
            if (127 == -128) {
                v1 += v1
                if (8 <= 4294967295) {
                    switch ("hello" + "foo bar") {
                    case "":
                        var v6 uint16 = (v1 >> ((18446744073709551615) & 63))
                        _ = v6
                    }
                }
                if false {
                    fmt.Println(v1, v2)
                }
            }
            var v7 uint16 = (^(v1))
            _ = v7
        } else {
            var v8 int16 = -32768
            v8 = 1000
            _ = v8
        }
    }
    switch 0 {
    case 682:
        var v9 float64 = 1e10
        _ = v9
    case 32768:
        fmt.Println(v2, v2)
    }
    if (false || (127 >= 37)) {
        if ((!false) || false) {
            var v10 int64 = 1
            _ = v10
        } else {
            fmt.Println(v1)
            v1 += (v1 / ((v1) | 1))
            var v11 string = ""
            _ = v11
        }
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
