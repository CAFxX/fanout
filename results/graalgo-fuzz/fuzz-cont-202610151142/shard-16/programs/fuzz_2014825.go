package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint8, p1 int) int8 {
    v6 := []int{p1, (-(p1))}
    p1 += (p1 | p1)
    _ = p0
    _ = p1
    _ = v6
    return -28
}

func main() {
    var v1 string = ("" + ("\n\t" + "hello"))
    var v2 uint = 537
    v3 := []uint32{2147483648}
    v4 := make(map[string]bool)
    var v5 uint8 = 1
    v2 = v2
    var v7 uint16 = 81
    v1 = "\n\t"
    _ = len(v4)
    for fi8 := 0; fi8 < 20; fi8++ {
        v7 -= 325
        var v9 int64 = 9223372036854775807
        switch (((v1 + ((v1 + "\n\t") + (v1 + v1))) + ((v1 + v1) + ("hello" + v1))) + "foo bar") {
        case "hello":
            v10 := GenS1{F0: v9}
            _ = v10
        case "":
            v1 = "a"
        default:
            if false {
                var v11 uint64 = 4294967295
                _ = v11
            } else {
                v7 = 255
            }
        }
        _ = fi8
        _ = v9
    }
    if false {
        var v12 int8 = -1
        _ = v12
    }
    var v13 float64 = ((-2.25 - 3.14159) + 1e10)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v1
    var v14 int32 = 560
    fmt.Println(v13, v1)
    fmt.Println(v5, v13, v5)
    var v15 bool = (-117 < -28)
    v15 = (!v15)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
