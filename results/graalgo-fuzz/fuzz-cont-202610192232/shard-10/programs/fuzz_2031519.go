package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int32 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 1.5
    var v3 uint = 255
    v4 := []int{30, 44, 42, 42}
    v5 := []string{"a", ("\n\t" + ("a" + "hello")), (("hello" + "\n\t") + ("foo bar" + "a")), "\n\t"}
    fmt.Println(v3, v3)
    var v6 string = ""
    if (1e10 != (v2 - float32(3.14159))) {
        switch 42 {
        case 1:
            for fi7 := 0; fi7 < 8; fi7++ {
                v6 = v6
                _ = len(v1)
                _ = fi7
            }
        case 42:
            v8 := GenS1{F0: v3, F1: v6, F2: -1}
            _ = v8
        case -1:
            v9 := GenS1{F0: v3, F1: (v6 + v6), F2: 1}
            _ = v9
        }
        v3 -= v3
        delete(v1, "")
    }
    var v10 uint = (v3 / ((529) | 1))
    var v11 uint32 = 2147483648
    v12 := GenS2{F0: 73}
    var v13 int32 = 752
    var v14 int32 = v13
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
