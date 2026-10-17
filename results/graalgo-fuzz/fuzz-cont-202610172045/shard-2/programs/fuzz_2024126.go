package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 string; F2 bool }

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("hello" + "")
    var v3 uint16 = 809
    var v4 int8 = -128
    for fi5 := 0; fi5 < 6; fi5++ {
        v4 = (v4 << ((911) & 63))
        _ = fi5
    }
    var v6 int8 = v4
    v2 = v2
    v4 = (v4 >> ((4294967295) & 63))
    var v7 uint = 65535
    v7 = (v7 * 18446744073709551615)
    v8 := []int{9223372036854775807, 127, 1, int(v7)}
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            if false {
                v11 := v1[fi10]
                _ = v11
            }
            if false {
                v3 -= v3
                if (!true) {
                    _ = len(v8)
                } else {
                    _ = len(v1)
                    if len(v8) > 0 {
                        v8[0] = 127
                    }
                }
            } else {
                var v12 uint32 = 4294967295
                if false {
                    var v13 bool = (-9223372036854775808 >= -1)
                    var v14 uint = (v7 >> (((v7 >> (((v7 * v7)) & 63))) & 63))
                    _ = v13
                    _ = v14
                }
                v15 := GenS1{F0: fi9}
                _ = v12
                _ = v15
            }
            _ = fi10
        }
        v2 = (("\n\t" + v2) + v2)
        v16 := GenS2{F0: (v6 << ((18446744073709551615) & 63)), F1: "a", F2: true}
        _ = fi9
        _ = v16
    }
    if (v3 >= (^(v3))) {
        v17 := GenS2{F0: (v6 - (v4 << (((v7 >> (((v7 - 855)) & 63))) & 63))), F1: (v2 + v2), F2: true}
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
