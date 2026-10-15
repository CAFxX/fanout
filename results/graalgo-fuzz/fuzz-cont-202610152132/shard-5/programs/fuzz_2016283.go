package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0() uint8 {
    var v4 string = "foo bar"
    v5 := GenS1{F0: 65535}
    var v6 float64 = 1.5
    var v7 float32 = float32(3.14159)
    v8 := []int{-1, -9223372036854775808, -1, 1}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    var v3 uint64 = 42
    var v9 int16 = -41
    v10 := GenS1{F0: 0}
    if false {
        var v11 int = 905
        _ = v11
    }
    var v12 uint16 = 1
    v9 = 0
    for fi13 := 0; fi13 < 18; fi13++ {
        v3 = 101
        if (0 <= 127) {
            switch ("hello" + ("a" + "")) {
            case "hello":
                switch (fi13 * (fi13 - fi13)) {
                case -761:
                    switch "a" {
                    case "a":
                        v14 := []int{2147483647, fi13, fi13}
                        _ = v14
                    case "hello":
                        switch "\n\t" {
                        case "hello":
                            v12 = v12
                        case "":
                            v3 = (v3 >> ((0) & 63))
                        case "\n\t":
                            v9 -= 42
                        }
                    case "\n\t":
                        if v1 {
                            var v15 uint32 = 4294967295
                            switch "foo bar" {
                            case "foo bar":
                                v2[v15] = ((44 <= 0) || (!v1))
                            }
                            _ = v15
                        }
                    }
                default:
                    v16 := []int{fi13, fi13, fi13, fi13}
                    _ = v16
                }
            }
        } else {
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 string = "\n\t"
                _ = fi17
                _ = v18
            }
        }
        _ = fi13
    }
    v9 = (v9 + (v9 >> ((0) & 63)))
    v19 := []int{32768, -2147483648, -2147483648, -9223372036854775808}
    fmt.Println(v9, v1, v3)
    v20 := GenS2{F0: 1, F1: -9223372036854775808}
    for fi21 := 0; fi21 < 17; fi21++ {
        switch "a" {
        case "a":
            v22 := GenS2{F0: 2, F1: 9223372036854775807}
            _ = v22
        case "\n\t":
            for fi23 := 0; fi23 < 14; fi23++ {
                fmt.Println(fi23, v3)
                v3 -= 18446744073709551615
                var v24 uint16 = (v12 >> ((42) & 63))
                v19 = append(v19, (fi21 ^ fi23))
                _ = fi23
                _ = v24
            }
        }
        _ = fi21
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        v9 += 42
        var v26 float32 = 3.14159
        switch "hello" {
        case "a":
            var v27 uint64 = v3
            _ = v27
        case "\n\t":
            v28 := v2[uint32(v12)]
            _ = v28
        case "hello":
            v9 += v9
        default:
            if len(v19) > 0 {
                v19[0] = fi25
            }
        }
        _ = fi25
        _ = v26
    }
    if false {
        v29 := GenS2{F0: 42, F1: -149}
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
