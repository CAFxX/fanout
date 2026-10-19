package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int64 = 750
    var v2 float32 = -2.25
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    var v5 int16 = 0
    v2 -= 1e10
    var v6 float64 = (float64(v2) + 1e10)
    var v7 bool = true
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{32767}
        for fi10 := 0; fi10 < 15; fi10++ {
            v2 -= 0.1
            v9 = append(v9, (-(fi10)))
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    v11 := []int{-57}
    var v12 uint = 1
    switch "\n\t" {
    case "":
        for fi13 := 0; fi13 < 16; fi13++ {
            v5 = v5
            var v14 string = "hello"
            _ = fi13
            _ = v14
        }
    case "\n\t":
        v12 += 0
    default:
        v5 += (v5 - v5)
    }
    v15 := GenS1{F0: uint16(v12)}
    var v16 float32 = v2
    fmt.Println(v1)
    for fi17 := 0; fi17 < 3; fi17++ {
        v2 = 1e10
        if (65535 > v12) {
            v11 = append(v11, fi17)
            var v18 uint32 = 2147483647
            _ = v18
        } else {
            if len(v11) > 0 {
                v11[0] = fi17
            }
        }
        var v19 string = ("a" + (("foo bar" + ("hello" + "foo bar")) + ((((("a" + "\n\t") + "") + "\n\t") + ("foo bar" + "")) + (("\n\t" + ("hello" + ("\n\t" + "foo bar"))) + "hello"))))
        _ = fi17
        _ = v19
    }
    fmt.Println(v5)
    if true {
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
