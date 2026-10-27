package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int {
    var v5 float32 = ((-2.25 / 0.1) / 1.0)
    v5 = (float32(0.1) - ((((v5 * 1e10) + (v5 / v5)) * float32(3.14159)) - v5))
    var v6 uint = 1
    var v7 uint64 = 42
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    var v3 uint = 255
    var v4 string = ("hello" + "\n\t")
    var v8 uint64 = 998
    v9 := GenS1{F0: 42}
    fmt.Println(v3, v8, v1)
    switch 128 {
    case -9223372036854775808:
        if v1 {
            var v10 string = ("foo bar" + (("hello" + "foo bar") + "foo bar"))
            _ = v10
        } else {
            v11 := []int{-618, 9223372036854775807, 42}
            _ = v11
        }
    case 32768:
        v8 = v8
    case 443:
        if (((((((3.14159 * ((1.5 + (3.14159 + (0.1 * -2.25))) + 3.14159)) * 1e10) / (1.5 + 0.0)) - 0.0) + ((1e10 / 0.1) * 0.1)) - 0.1) > (1.5 / (0.1 + 1.5))) {
            var v12 uint32 = 2147483647
            v13 := []int{32767, 127, -164}
            v12 = (v12 | v12)
            _ = v12
            _ = v13
        }
    }
    fmt.Println(v1, v4)
    var v14 uint64 = (v8 & (v8 * (v8 & v8)))
    var v15 int16 = 1
    switch v4 {
    case "\n\t":
        v4 = (("foo bar" + "\n\t") + v4)
    case "a":
        v1 = v1
    }
    v16 := GenS1{F0: 127}
    var v17 uint16 = 338
    v1 = (!v1)
    fmt.Println(v17)
    fmt.Println(v14)
    v18 := v2["hello"]
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
