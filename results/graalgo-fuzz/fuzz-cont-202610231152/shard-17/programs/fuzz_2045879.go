package main

import (
    "fmt"
)


func main() {
    var v1 float32 = ((1.5 - (0.0 - 1.5)) - 0.1)
    var v2 int32 = -1
    var v3 uint32 = 2
    var v4 bool = (2 != 194)
    v5 := []int{374, -2147483648, -2147483648, 4294967295}
    var v6 int16 = -2
    var v7 bool = (v4 || ((-9223372036854775808 <= 292) && (578 != 65535)))
    var v8 uint = 119
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v7, v6, v7)
        _ = fi9
    }
    switch (("a" + "") + "\n\t") {
    case "foo bar":
        for fi10 := 0; fi10 < 11; fi10++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v11 uint32 = v3
            var v12 int32 = (v2 ^ v2)
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v8, v8, v2)
    fmt.Println(v6, v8)
    var v13 int64 = -9223372036854775808
    v8 += (v8 + (v8 << (((v8 * 65535)) & 63)))
    v8 = 18446744073709551615
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
