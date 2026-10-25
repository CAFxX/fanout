package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 0
    v3 := make(map[string]string)
    var v4 int8 = -55
    v5 := make(map[string]uint32)
    v6 := make(map[int]string)
    if (453 == 2) {
        if (458 > 2) {
            for fi7 := 0; fi7 < 1; fi7++ {
                v2 -= (v2 << ((255) & 63))
                v8 := GenS1{F0: v2, F1: "a"}
                var v9 uint16 = v2
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 float32 = 3.14159
            _ = v10
        }
    }
    v2 -= (v2 >> ((388) & 63))
    var v11 uint = 0
    var v12 int = 94
    fmt.Println(v4)
    if true {
        if (false || (true || (!((((("hello" + ((((("hello" + "hello") + ("" + "hello")) + (("a" + "foo bar") + "hello")) + ("\n\t" + "a")) + "")) + "hello") + ("hello" + "foo bar")) + ("\n\t" + "")) <= ("\n\t" + "\n\t"))))) {
            var v13 int64 = 225
            v14 := GenS1{F0: (^(v2)), F1: "foo bar"}
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            fmt.Println(v4)
            v2 = (v2 | (v2 * 0))
            v4 = 121
            _ = fi15
        }
        delete(v5, "foo bar")
    }
    if true {
        var v16 float32 = (1.5 * float32(float64(-2.25)))
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint16 = (v2 >> (((v11 + v11)) & 63))
        _ = fi17
        _ = v18
    }
    fmt.Println(v12, v12, v11)
    var v19 int8 = (v4 & (v4 << ((v11) & 63)))
    var v20 uint = v11
    v19 += 126
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
