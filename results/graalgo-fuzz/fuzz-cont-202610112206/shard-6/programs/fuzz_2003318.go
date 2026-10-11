package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]int)
    var v3 int16 = 782
    for fi4 := 0; fi4 < 5; fi4++ {
        if (2147483648 == 511) {
            var v5 bool = true
            v6 := []int{fi4}
            _ = v5
            _ = v6
        }
        switch (v1 | 32767) {
        case -2147483648:
            if false {
                var v7 uint8 = 1
                v2[2147483647] = (fi4 | int(v7))
                switch 969 {
                case -971:
                    switch "a" {
                    case "a":
                        if (2147483647 != 955) {
                            v8 := GenS1{F0: 18446744073709551615, F1: (v7 >> ((2) & 63))}
                            if false {
                                var v9 uint64 = 4294967295
                                _ = v9
                            } else {
                                v2[65535] = fi4
                                fmt.Println(v7, fi4)
                            }
                            _ = v8
                        }
                    default:
                        v1 = (v1 / ((fi4) | 1))
                    }
                }
                _ = v7
            } else {
                if false {
                    v1 -= (fi4 * (-(v1)))
                    _ = len(v2)
                }
            }
        case 42:
            v10 := GenS1{F0: 42, F1: 42}
            _ = v10
        case 32768:
            var v11 uint = 838
            _ = v11
        }
        _ = fi4
    }
    var v12 string = "hello"
    v12 = v12
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := []int{}
        _ = fi13
        _ = v14
    }
    _ = len(v2)
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 string = v12
        switch (("hello" + "hello") + "") {
        case "a":
            fmt.Println(v3, fi15, v12)
        case "foo bar":
            var v17 int64 = 686
            _ = v17
        default:
            var v18 uint16 = 255
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    var v19 int32 = 0
    v20 := GenS1{F0: 4294967295, F1: 79}
    v21 := []int{v1}
    v22 := v2[334]
    var v23 int32 = -2147483648
    if (2 <= 1) {
        var v24 string = (("foo bar" + v12) + ("a" + v12))
        _ = v24
    } else {
        var v25 int32 = (v23 ^ (v19 & v23))
        _ = v25
    }
    delete(v2, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
