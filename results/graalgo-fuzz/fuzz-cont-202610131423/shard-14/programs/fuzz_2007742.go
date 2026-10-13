package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    var v1 uint8 = 127
    var v2 int64 = -9223372036854775808
    var v3 string = (("\n\t" + "") + "foo bar")
    v4 := make(map[int]bool)
    var v5 uint8 = v1
    v6 := make(map[string]uint32)
    v7 := []int{-1, 4294967295, int(v5), 127}
    v3 = (("a" + v3) + v3)
    for fi8 := 0; fi8 < 15; fi8++ {
        switch fi8 {
        case 128:
            var v9 float64 = (1e10 / 3.14159)
            _ = v9
        case 1:
            var v10 uint = 872
            _ = v10
        case -1:
            fmt.Println(v3, fi8, v3)
        }
        var v11 uint16 = 1
        _ = fi8
        _ = v11
    }
    v12 := GenS1{F0: (698 == 0)}
    switch 423 {
    case 127:
        if false {
            switch "\n\t" {
            case "a":
                v1 = (v1 >> ((893) & 63))
            }
            if true {
                v5 = (v5 << ((18446744073709551615) & 63))
                if (-1 == 773) {
                    var v13 int8 = 0
                    _ = v13
                } else {
                    fmt.Println(v2, v2)
                }
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 uint32 = 65535
                    v15 += (^(v15))
                    _ = fi14
                    _ = v15
                }
            } else {
                switch 127 {
                case 805:
                    var v16 bool = (127 > -128)
                    _ = v16
                case -9223372036854775808:
                    fmt.Println(v5, v5, v5)
                }
                var v17 uint16 = 395
                _ = v17
            }
        } else {
            v5 = v1
            v18 := GenS2{F0: (v2 * v2), F1: 42}
            _ = v18
        }
    case -2147483648:
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS2{F0: 9223372036854775807, F1: 590}
            switch v3 {
            case "a":
                var v21 int32 = 42
                _ = v21
            case "foo bar":
                fmt.Println(v2)
            }
            var v22 float32 = 0.1
            _ = fi19
            _ = v20
            _ = v22
        }
    }
    var v23 uint32 = 956
    var v24 bool = false
    var v25 string = v3
    v3 = v3
    v26 := GenS2{F0: (v2 >> ((401) & 63)), F1: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
