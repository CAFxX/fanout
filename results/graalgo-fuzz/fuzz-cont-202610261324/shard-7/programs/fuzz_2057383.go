package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]bool)
    var v2 string = "hello"
    var v3 int64 = -9223372036854775808
    v3 = v3
    var v4 uint = 255
    switch v2 {
    case "\n\t":
        v2 = (v2 + v2)
    case "":
        if false {
            v3 = v3
            if (42 >= 0) {
                fmt.Println(v2)
                v4 -= (v4 << ((v4) & 63))
            } else {
                var v5 int32 = 2147483647
                if (true || true) {
                    v1[v2] = false
                } else {
                    v4 = (v4 << (((^(v4))) & 63))
                }
                _ = v5
            }
            var v6 int32 = 1
            _ = v6
        } else {
            var v7 float64 = (float64((3.14159 + 1.5)) / 1e10)
            _ = v7
        }
    }
    var v8 uint64 = 4294967295
    v1[(("" + (v2 + "foo bar")) + (("" + v2) + ((v2 + ("" + (((((("hello" + ("a" + "")) + (v2 + v2)) + "\n\t") + v2) + v2) + (((("hello" + "a") + ((v2 + v2) + (v2 + (v2 + (v2 + ""))))) + "a") + v2)))) + "\n\t")))] = (!(18446744073709551615 < v4))
    v9 := v1["foo bar"]
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 string = "a"
        _ = fi10
        _ = v11
    }
    var v12 float64 = (0.0 - 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
