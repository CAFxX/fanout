package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int = 649
    var v2 uint8 = 127
    var v3 bool = true
    var v4 int64 = 0
    v5 := GenS1{F0: 2147483647, F1: 0, F2: v1}
    if (-2.25 > 1.5) {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int32 = 0
            switch "\n\t" {
            case "":
                v8 := GenS1{F0: (fi6 / ((v1) | 1)), F1: (v4 & v4), F2: 725}
                _ = v8
            case "a":
                v9 := GenS2{F0: uint16(v2), F1: v2}
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        v10 := GenS2{F0: 2, F1: v2}
        _ = v10
    } else {
        var v11 int64 = v4
        _ = v11
    }
    fmt.Println(v4)
    switch (v1 & v1) {
    case 347:
        v12 := []int{-255}
        _ = v12
    }
    fmt.Println(v4)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int8 = -104
        _ = fi13
        _ = v14
    }
    v2 += (v2 << ((246) & 63))
    switch 128 {
    case 1:
        if true {
            if false {
                fmt.Println(v2, v3)
                if (1 >= 2147483647) {
                    v15 := GenS2{F0: 2, F1: 128}
                    v16 := []int{}
                    _ = v15
                    _ = v16
                } else {
                    if (!v3) {
                        if ("\n\t" != ("hello" + "\n\t")) {
                            v17 := []int{(v1 | 1), v1, 0}
                            _ = v17
                        } else {
                            v4 = -1
                        }
                    }
                }
                var v18 int = -79
                _ = v18
            } else {
                var v19 int16 = 0
                _ = v19
            }
            v2 += (v2 - (v2 % ((20) | 1)))
            if true {
                var v20 uint8 = 0
                _ = v20
            } else {
                v21 := GenS3{F0: 32767, F1: 2147483648}
                var v22 uint32 = 1
                _ = v21
                _ = v22
            }
        } else {
            if v3 {
                v23 := []int{(-(v1))}
                var v24 uint8 = 0
                fmt.Println(v4)
                _ = v23
                _ = v24
            }
            var v25 uint8 = (v2 | 127)
            _ = v25
        }
    case 42:
        for fi26 := 0; fi26 < 10; fi26++ {
            for fi27 := 0; fi27 < 15; fi27++ {
                v28 := GenS4{F0: 18446744073709551615, F1: 42}
                _ = fi27
                _ = v28
            }
            fmt.Println(v1, fi26, fi26)
            var v29 int16 = 804
            _ = fi26
            _ = v29
        }
    case -1:
        fmt.Println(v2)
    default:
        fmt.Println(v2)
    }
    fmt.Println(v4, v4, v1)
    var v30 uint8 = (v2 & (v2 % ((2) | 1)))
    var v31 uint8 = 128
    switch (v1 + (-(v1))) {
    case 0:
        var v32 int16 = 0
        _ = v32
    case -2147483648:
        var v33 string = ("" + "")
        _ = v33
    default:
        v1 = 985
    }
    if v3 {
        var v34 bool = v3
        fmt.Println(v34)
        _ = v34
    }
    var v35 uint64 = 18446744073709551615
    if false {
        fmt.Println(v1)
    }
    v36 := GenS3{F0: 539, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
