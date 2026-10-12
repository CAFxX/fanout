package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 0
    v2 := []uint32{42}
    var v3 uint32 = 65535
    v3 -= v3
    var v4 uint = 255
    var v5 string = ""
    switch -1 {
    case -400:
        v3 = 4294967295
    case 9223372036854775807:
        fmt.Println(v3)
    }
    v2 = append(v2, (v3 & 2147483647))
    fmt.Println(v5)
    _ = len(v2)
    if (4294967295 == 1) {
        var v6 int32 = 1
        _ = v6
    } else {
        if len(v2) > 0 {
            v2[0] = 0
        }
        v4 += 973
        v4 -= 541
    }
    var v7 uint16 = 2
    fmt.Println(v5, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
