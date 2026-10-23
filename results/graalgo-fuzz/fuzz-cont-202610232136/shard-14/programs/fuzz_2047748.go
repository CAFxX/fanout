package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 string; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 uint64 = 747
    var v2 bool = (("hello" + "\n\t") >= (("a" + "a") + ("foo bar" + "foo bar")))
    var v3 uint32 = 1
    for fi4 := 0; fi4 < 20; fi4++ {
        switch 2147483647 {
        case 32768:
            switch "a" {
            case "\n\t":
                v5 := []int{}
                _ = v5
            default:
                v1 = (v1 << ((4294967295) & 63))
            }
        default:
            if false {
                for fi6 := 0; fi6 < 8; fi6++ {
                    v7 := GenS1{F0: -253}
                    if (!(4294967295 == uint(v1))) {
                        var v8 uint = 42
                        v3 += 497
                        v9 := []int{fi6}
                        _ = v8
                        _ = v9
                    } else {
                        var v10 string = "foo bar"
                        if true {
                            v10 = ((v10 + v10) + "\n\t")
                            v10 = ((v10 + "hello") + v10)
                        }
                        fmt.Println(v3, v2)
                        _ = v10
                    }
                    _ = fi6
                    _ = v7
                }
            } else {
                if ((fi4 / ((fi4) | 1)) > fi4) {
                    v11 := GenS1{F0: 9223372036854775807}
                    var v12 int = -2147483648
                    _ = v11
                    _ = v12
                } else {
                    switch (fi4 | fi4) {
                    case 42:
                        v13 := GenS2{F0: (v3 | (^(v3))), F1: 15}
                        _ = v13
                    case 9223372036854775807:
                        v2 = (!v2)
                    default:
                        v2 = v2
                    }
                }
                fmt.Println(fi4)
                fmt.Println(v3, v1, fi4)
            }
        }
        v14 := GenS3{F0: fi4, F1: ""}
        v15 := []int{fi4}
        _ = fi4
        _ = v14
        _ = v15
    }
    v3 += (^(v3))
    var v16 bool = (-32768 != 903)
    var v17 float64 = -2.25
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 int32 = 2147483647
        _ = fi18
        _ = v19
    }
    var v20 int8 = 100
    v21 := GenS4{F0: "a", F1: uint8(v1), F2: 887}
    v22 := GenS1{F0: 9223372036854775807}
    fmt.Println(v1)
    var v23 int32 = -426
    v24 := GenS5{F0: -32768}
    if (!false) {
        switch ("" + ("" + "a")) {
        case "hello":
            var v25 int16 = -32768
            _ = v25
        }
    } else {
        var v26 int8 = 127
        v27 := []int{32768, 127, -132}
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v17 -= (v17 * 1e10)
        _ = fi28
    }
    switch "" {
    case "foo bar":
        fmt.Println(v2)
    case "hello":
        var v29 uint = 440
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
