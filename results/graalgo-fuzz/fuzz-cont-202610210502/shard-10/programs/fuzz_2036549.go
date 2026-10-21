package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint16 = 0
    var v2 string = ("hello" + (("" + (("foo bar" + "") + "")) + ""))
    var v3 uint16 = (v1 / (((v1 + 218)) | 1))
    var v4 uint = 42
    var v5 float32 = float32(3.14159)
    fmt.Println(v4, v1, v1)
    var v6 int = 32768
    var v7 int32 = -2147483648
    if ((18446744073709551615 < 2) || true) {
        v1 += 42
    } else {
        fmt.Println(v6, v1, v6)
    }
    var v8 int = (v6 / ((v6) | 1))
    var v9 float32 = v5
    v1 = 42
    v10 := []int{v8, (v6 + v8), (^(v6))}
    var v11 string = "\n\t"
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS1{F0: v4}
        _ = fi12
        _ = v13
    }
    v8 = (-(v8))
    fmt.Println(v6, v7, v11)
    v5 = (v5 / 1.0)
    v9 -= float32((3.14159 - (-2.25 / float64(v9))))
    fmt.Println(v8)
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if false {
        fmt.Println(v4, v2, v9)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
