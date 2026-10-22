package main

import (
    "fmt"
)


func main() {
    v1 := []int{2147483647, 234, -9223372036854775808}
    var v2 int8 = 122
    var v3 uint = 4294967295
    for fi4 := 0; fi4 < 15; fi4++ {
        v3 = (v3 - 255)
        v3 = v3
        _ = fi4
    }
    for fi5 := 0; fi5 < 6; fi5++ {
        v3 -= (v3 - (v3 + 18446744073709551615))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi5
    }
    v3 -= (v3 & (v3 + 150))
    v1 = append(v1, 32767)
    var v6 float32 = 0.1
    switch 32768 {
    case 4294967295:
        v6 = v6
    default:
        fmt.Println(v2, v6, v2)
    }
    fmt.Println(v2)
    v2 = (v2 / (((v2 - 42)) | 1))
    var v7 uint64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
