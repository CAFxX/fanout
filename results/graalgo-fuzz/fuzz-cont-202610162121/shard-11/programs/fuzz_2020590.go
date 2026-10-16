package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -1
    v3 := make(map[string]string)
    var v4 string = ("\n\t" + ("hello" + "\n\t"))
    if ((true || true) && true) {
        v5 := GenS1{F0: -50, F1: 32767, F2: (false || (!((v4 + ("\n\t" + v4)) < ("" + v4))))}
        var v6 uint = 2
        if false {
            var v7 int32 = 2147483647
            v4 = v4
            if (!((!(v4 < ((v4 + v4) + ((v4 + "a") + v4)))) && true)) {
                v7 -= (v7 & v7)
            } else {
                fmt.Println(v2, v2, v6)
                var v8 uint = (v6 << ((v6) & 63))
                v9 := GenS1{F0: 102, F1: 656, F2: (0 > 1)}
                _ = v8
                _ = v9
            }
            _ = v7
        }
        _ = v5
        _ = v6
    }
    v10 := GenS1{F0: 0, F1: 0, F2: (v4 >= ((v4 + ("hello" + v4)) + v4))}
    v4 = (v4 + "a")
    v4 = v4
    var v11 int64 = 9223372036854775807
    var v12 int = 128
    v4 = v4
    fmt.Println(v11)
    var v13 int = (v12 ^ (v12 - 127))
    v14 := []int{v12, (v12 - v12), 868, 203}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
