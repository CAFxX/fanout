package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint8 {
    v4 := []int{-115}
    var v5 uint16 = 437
    var v6 float32 = 1.5
    var v7 int16 = int16(v5)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 string = "hello"
    var v2 int64 = -9223372036854775808
    var v3 string = "hello"
    var v8 string = v3
    var v9 uint = 212
    v10 := GenS1{F0: 100}
    var v11 int16 = -82
    switch "\n\t" {
    case "\n\t":
        if (false && false) {
            v12 := []int{-9223372036854775808, 32767, 0, -9223372036854775808}
            if (!(float32((-2.25 + (1.5 * 1.5))) != 1.5)) {
                var v13 int16 = v11
                fmt.Println(v1, v13, v3)
                v14 := GenS1{F0: 24}
                _ = v13
                _ = v14
            } else {
                v9 -= (v9 + v9)
                if (!(945 > 145)) {
                    fmt.Println(v9, v9, v8)
                }
            }
            _ = v12
        }
    }
    v3 = ((v8 + v8) + "foo bar")
    var v15 uint = 18446744073709551615
    v16 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
