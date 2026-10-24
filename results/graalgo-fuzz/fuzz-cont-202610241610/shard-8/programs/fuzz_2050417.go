package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 bool = (true || true)
    var v2 uint16 = 2
    var v3 uint32 = 4294967295
    v4 := []int{}
    fmt.Println(v3, v2)
    if (0.1 >= float64(0.1)) {
        v1 = ((v1 || v1) && (!(((v2 | 1) == (^(v2))) || (!v1))))
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    }
    for fi5 := 0; fi5 < 2; fi5++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (42 != 4294967295) {
            var v6 bool = (true || false)
            v7 := GenS1{F0: 0}
            _ = v6
            _ = v7
        } else {
            var v8 uint16 = (v2 + 510)
            v3 += (v3 % ((248) | 1))
            fmt.Println(v3)
            _ = v8
        }
        _ = fi5
    }
    var v9 int = -1
    v1 = true
    fmt.Println(v1)
    v2 = (v2 << ((4294967295) & 63))
    if (v1 && true) {
        if (1 != -1) {
            v3 = (v3 + (^(v3)))
        } else {
            _ = len(v4)
            v3 = v3
            v9 = int(v3)
        }
        if ((135 != 128) && v1) {
            v3 = 785
        } else {
            fmt.Println(v1)
            v3 = (v3 + (v3 >> ((2) & 63)))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
