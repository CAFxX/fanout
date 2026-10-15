package main

import (
    "fmt"
)


func main() {
    var v1 uint16 = 215
    v2 := make(map[string]uint32)
    var v3 bool = (2 != 0)
    var v4 uint16 = (^(v1))
    var v5 float32 = 3.14159
    _ = len(v2)
    switch -304 {
    case -147:
        for fi6 := 0; fi6 < 18; fi6++ {
            fmt.Println(fi6, v4, fi6)
            _ = fi6
        }
    case -539:
        if (!(true && ((("a" == "") && v3) && true))) {
            var v7 bool = false
            _ = v7
        } else {
            var v8 int32 = -2147483648
            var v9 uint64 = 1
            v5 += ((v5 + 3.14159) / v5)
            _ = v8
            _ = v9
        }
    case 4294967295:
        v3 = (!v3)
    }
    fmt.Println(v1)
    fmt.Println(v4, v4)
    var v10 uint8 = 24
    var v11 string = "foo bar"
    _ = len(v2)
    v2[((("hello" + v11) + v11) + "foo bar")] = 575
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
