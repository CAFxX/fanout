package main

import (
    "fmt"
)


func main() {
    var v1 string = ""
    var v2 uint = 42
    var v3 float32 = -2.25
    var v4 string = (v1 + (v1 + ((v1 + v1) + (v1 + "hello"))))
    var v5 uint8 = 127
    var v6 bool = (1e10 > 3.14159)
    var v7 uint = (v2 % (((v2 - v2)) | 1))
    var v8 float32 = ((v3 - v3) + v3)
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{(fi9 >> ((v7) & 63)), (fi9 & (fi9 * 2147483647)), 1}
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 42
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int8 = 1
        fmt.Println(v6)
        _ = fi12
        _ = v13
    }
    var v14 int32 = -530
    v11 -= v11
    switch 0 {
    case 32768:
        fmt.Println(v1, v1, v5)
    case -9223372036854775808:
        fmt.Println(v6, v2)
    case 4294967295:
        var v15 float64 = -2.25
        _ = v15
    }
    fmt.Println(v6, v11)
    switch ("hello" + "") {
    case "hello":
        v16 := []int{2147483647}
        _ = v16
    case "\n\t":
        var v17 int = 42
        _ = v17
    default:
        var v18 uint32 = 4294967295
        _ = v18
    }
    var v19 bool = true
    if v19 {
        v7 -= v7
    }
    var v20 bool = (v14 >= (v14 - (v14 << ((v7) & 63))))
    if v6 {
        v21 := []int{-2147483648, -1}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
