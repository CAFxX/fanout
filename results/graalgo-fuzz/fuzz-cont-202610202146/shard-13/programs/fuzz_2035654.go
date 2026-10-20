package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
func genFn0(p0 int8) bool {
    var v6 uint64 = 18446744073709551615
    v6 += (v6 << ((1) & 63))
    var v7 float64 = 3.14159
    v8 := []int{-82, 723, 42, 128}
    if len(v8) > 0 {
        v8[0] = 128
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (9223372036854775807 > -9223372036854775808)
}
func genFn1(p0 int, p1 int8, p2 string) uint64 {
    switch (p0 - (p0 | -522)) {
    case -9223372036854775808:
        var v9 uint = 4294967295
        _ = v9
    default:
        if false {
            v10 := GenS1{F0: 79, F1: 4294967295, F2: ((false && true) || (!false))}
            var v11 uint8 = 37
            p2 = ("\n\t" + (p2 + ("" + p2)))
            _ = v10
            _ = v11
        }
    }
    switch 32767 {
    case -2147483648:
        var v12 int64 = 9223372036854775807
        _ = v12
    case -1:
        p2 = (p2 + ("hello" + ((p2 + "hello") + (p2 + ("" + "\n\t")))))
    case 9223372036854775807:
        v13 := []int{(p0 / ((p0) | 1)), int(p1), 0, p0}
        _ = v13
    default:
        var v14 bool = ((!((42 >= 1) || (true || (p2 < (p2 + ((("hello" + "") + (p2 + p2)) + p2)))))) && (true || (42 < 0)))
        _ = v14
    }
    v15 := []int{}
    p1 = (p1 / (((p1 / ((42) | 1))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 1
}

func main() {
    var v1 uint32 = 630
    var v2 string = ""
    var v3 uint8 = 118
    var v4 int8 = 57
    v5 := []uint32{(v1 / ((42) | 1))}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v16 uint = 0
    v17 := GenS1{F0: 2, F1: 42, F2: (-2147483648 == -403)}
    v18 := GenS1{F0: 2, F1: 1, F2: (1.5 < ((0.1 + -2.25) - -2.25))}
    var v19 string = "hello"
    v19 = v19
    switch (v2 + v19) {
    case "\n\t":
        var v20 uint16 = 32767
        _ = v20
    }
    if len(v5) > 0 {
        v5[0] = uint32(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
