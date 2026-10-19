package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int32 = -982
    var v2 bool = (18446744073709551615 > 590)
    var v3 int64 = 9223372036854775807
    var v4 uint = 1
    var v5 bool = (!(!(!(uint64(v4) < 2))))
    fmt.Println(v1)
    switch (((("foo bar" + ("" + "foo bar")) + ("a" + "hello")) + "") + "a") {
    case "":
        fmt.Println(v3)
    case "a":
        v3 -= 9223372036854775807
    case "\n\t":
        v6 := []int{127, 32767, 32768}
        _ = v6
    }
    v7 := []int{2147483647, 2147483647, -655}
    v4 = v4
    var v8 int8 = 0
    var v9 int8 = v8
    v9 = v9
    var v10 uint16 = 42
    var v11 float64 = 1.5
    var v12 string = ("a" + "a")
    v13 := GenS1{F0: 425}
    var v14 float32 = ((3.14159 - 0.1) / (float32(1.5) / float32(v11)))
    var v15 uint8 = 69
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
