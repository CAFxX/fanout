package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    v1 := make(map[int]int)
    var v2 int = 128
    var v3 int16 = -32768
    v4 := []int{v2, 169, (v2 - (v2 >> ((0) & 63))), (v2 | v2)}
    v5 := make(map[int]string)
    var v6 uint8 = 0
    v2 -= (v2 | (v2 / ((v2) | 1)))
    var v7 uint16 = uint16(v6)
    var v8 int8 = 127
    if true {
        var v9 uint = 755
        if true {
            switch "foo bar" {
            case "":
                var v10 string = (("\n\t" + (("" + (("" + "a") + "")) + "\n\t")) + "foo bar")
                _ = v10
            }
            delete(v1, (v2 & v2))
            if true {
                if (!true) {
                    fmt.Println(v2, v7)
                    v4 = append(v4, (-(v2)))
                    _ = len(v1)
                }
            }
        }
        if len(v4) > 0 {
            v4[0] = -1
        }
        _ = v9
    }
    var v11 int8 = int8(v2)
    var v12 uint32 = 964
    if false {
        if true {
            v6 += v6
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            if (("hello" + "foo bar") > "") {
                for fi14 := 0; fi14 < 5; fi14++ {
                    v6 += 127
                    _ = fi14
                }
                v8 -= v11
                var v15 uint64 = 302
                _ = v15
            } else {
                var v16 int32 = int32(v12)
                var v17 uint32 = (^(v12))
                _ = v16
                _ = v17
            }
            _ = fi13
        }
    }
    v6 = v6
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v6 -= v6
    v2 -= (v2 & int(v11))
    fmt.Println(v8, v6)
    _ = len(v4)
    var v18 string = ("" + "hello")
    var v19 int8 = v11
    v20 := GenS1{F0: v18, F1: "a"}
    var v21 float64 = ((((1e10 + 1e10) * 3.14159) - 3.14159) * (0.1 - 1.5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
