package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 string }
func genFn0(p0 bool, p1 int8, p2 bool) int16 {
    v4 := GenS1{F0: 2, F1: 65535, F2: ((("" + "") + ("\n\t" + ("a" + "\n\t"))) + ("hello" + "hello"))}
    switch ("hello" + ("a" + "hello")) {
    case "foo bar":
        v5 := []int{1}
        _ = v5
    default:
        fmt.Println(p2, p1, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -356
}

func main() {
    var v1 uint16 = 541
    var v2 uint32 = 0
    v3 := []uint32{(v2 * v2), 42}
    v1 = 0
    v6 := GenS1{F0: 170, F1: 18446744073709551615, F2: ""}
    var v7 uint64 = 500
    switch -9223372036854775808 {
    case -2147483648:
        fmt.Println(v7, v2)
    }
    fmt.Println(v7)
    v2 = (v2 | (v2 - (v2 & (^(v2)))))
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int8 = -56
            _ = len(v3)
            if (-32768 < 32767) {
                var v11 bool = false
                v12 := GenS1{F0: 127, F1: 1, F2: ""}
                _ = v11
                _ = v12
            } else {
                fmt.Println(fi9)
                if (1 >= int32(fi9)) {
                    var v13 uint64 = 1
                    _ = v13
                } else {
                    var v14 float64 = (1e10 + ((0.0 * -2.25) - ((1.5 + ((float64(1.5) / 1.0) * 1e10)) / 1.0)))
                    _ = v14
                }
            }
            var v15 uint16 = 255
            _ = fi9
            _ = v10
            _ = v15
        }
        _ = fi8
    }
    v16 := GenS1{F0: 127, F1: 814, F2: ""}
    fmt.Println(v7)
    var v17 string = "hello"
    v7 += (v7 << ((80) & 63))
    v18 := GenS1{F0: uint8(v7), F1: 42, F2: v17}
    v3 = append(v3, 4294967295)
    fmt.Println(v7)
    v3 = append(v3, uint32(v7))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
