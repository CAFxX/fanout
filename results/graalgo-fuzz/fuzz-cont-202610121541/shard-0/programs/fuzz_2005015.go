package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int8) bool {
    var v6 int16 = -1
    var v7 int8 = (p0 | p0)
    _ = p0
    _ = v6
    _ = v7
    return (-115 <= -28)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 1.5
    var v3 string = ("\n\t" + "hello")
    v4 := []int{}
    var v5 float64 = 0.1
    var v8 uint8 = 48
    fmt.Println(v5, v5)
    v2 = ((v5 - (v2 * ((v5 * float64(3.14159)) + v5))) / 1.0)
    switch "foo bar" {
    case "\n\t":
        v3 = v3
    case "foo bar":
        v2 += 3.14159
    case "a":
        if true {
            v3 = v3
            v9 := GenS1{F0: -319}
            v10 := GenS1{F0: -216}
            _ = v10
            _ = v9
        } else {
            if true {
                v3 = (v3 + ("" + (("a" + ("a" + "hello")) + (v3 + v3))))
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    var v11 int64 = -576
    switch -673 {
    case 127:
        var v12 bool = (-903 <= 2147483647)
        _ = v12
    case 128:
        v2 = (0.1 - v5)
    default:
        v13 := GenS1{F0: -509}
        _ = v13
    }
    v14 := v1[32768]
    fmt.Println(v14, v14, v2)
    var v15 string = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
