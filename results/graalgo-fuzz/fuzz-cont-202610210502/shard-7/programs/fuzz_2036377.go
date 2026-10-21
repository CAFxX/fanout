package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint = (v1 >> (((v1 % (((^(v1))) | 1))) & 63))
    v3 := []string{("a" + (("a" + "hello") + "a"))}
    if (!(true && true)) {
        v3 = append(v3, ("foo bar" + "foo bar"))
    }
    v1 -= 255
    fmt.Println(v2, v1)
    var v4 uint64 = 325
    v5 := GenS1{F0: 42}
    v4 -= (v4 & v4)
    if true {
        if ((-530 >= 32767) && (548 >= 745)) {
            for fi6 := 0; fi6 < 13; fi6++ {
                v4 -= (v4 + (^(v4)))
                _ = fi6
            }
        }
        v1 = 18446744073709551615
        v3 = append(v3, "a")
    }
    fmt.Println(v2, v2)
    if ((((-2.25 == 0.0) || (1e10 <= (3.14159 * 0.0))) || (true || (-2.25 >= (-2.25 * 0.1)))) && true) {
        fmt.Println(v4, v2, v2)
    }
    v7 := GenS1{F0: 55}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
