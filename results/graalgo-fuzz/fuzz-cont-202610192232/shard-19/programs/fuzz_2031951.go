package main

import (
    "fmt"
)

func genFn0() int8 {
    var v6 uint16 = 32767
    var v7 uint = 306
    var v8 uint = (v7 % ((1) | 1))
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 99
    v3 := []string{"", ("foo bar" + (("" + ("foo bar" + ("\n\t" + "foo bar"))) + "foo bar"))}
    var v4 float64 = 3.14159
    v5 := []uint32{2147483648, 65535, 116}
    var v9 uint16 = 283
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = v2
    if (!(v4 != v4)) {
        v2 += v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            fmt.Println(v9, v9)
            v2 -= (v2 << ((18446744073709551615) & 63))
            var v11 bool = (!true)
            _ = fi10
            _ = v11
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if ((v9 >> ((42) & 63)) > v9) {
        v1[1] = (-2.25 > 0.1)
        fmt.Println(v4, v9, v4)
    }
    fmt.Println(v9, v4, v2)
    var v12 int16 = -32768
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
