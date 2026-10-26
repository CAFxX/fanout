package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 0
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]int)
    var v4 string = ("\n\t" + ("foo bar" + ("" + "hello")))
    var v5 int8 = v1
    var v6 int = -1
    for fi7 := 0; fi7 < 6; fi7++ {
        v6 = v6
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 string = (v4 + v4)
            switch (fi8 ^ 2147483647) {
            case 32767:
                fmt.Println(fi7)
            default:
                v1 -= 42
            }
            var v10 int = (v6 * 4294967295)
            var v11 int32 = 242
            v4 = ("hello" + v4)
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    _ = len(v3)
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 uint16 = 0
        _ = fi12
        _ = v13
    }
    if ((0 >= -1) && false) {
        v6 = (v6 * (v6 & 1))
    } else {
        switch (v6 * v6) {
        case -931:
            v3[2147483647] = (v6 & (^(v6)))
        case 42:
            var v14 int8 = (v5 & v5)
            _ = v14
        case -1:
            _ = len(v3)
        }
    }
    fmt.Println(v5, v1, v6)
    v2 -= (v2 << ((643) & 63))
    v15 := []int{v6}
    v4 = ""
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
