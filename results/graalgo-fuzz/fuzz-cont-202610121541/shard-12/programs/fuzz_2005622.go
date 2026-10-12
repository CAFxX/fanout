package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint64 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint8; F1 string }

func main() {
    var v1 float32 = (1e10 * (-2.25 / float32(1.5)))
    var v2 uint16 = 788
    var v3 float32 = v1
    var v4 uint16 = 2
    var v5 uint16 = (^(v4))
    var v6 uint = 0
    var v7 uint8 = 127
    if (!(4294967295 == 4294967295)) {
        var v8 int32 = 2147483647
        _ = v8
    }
    if (1 != -128) {
        fmt.Println(v1)
    } else {
        v1 = v3
        v2 = v5
    }
    var v9 string = "\n\t"
    fmt.Println(v2, v1, v9)
    var v10 int16 = 1000
    v11 := GenS1{F0: 516, F1: -1, F2: 0}
    v12 := []int{}
    if (!(!(true && ((127 == 102) && false)))) {
        fmt.Println(v7)
    } else {
        fmt.Println(v9, v7, v1)
        switch v9 {
        case "foo bar":
            for fi13 := 0; fi13 < 14; fi13++ {
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 uint32 = 1
                    _ = fi14
                    _ = v15
                }
                switch ("" + v9) {
                case "\n\t":
                    v16 := GenS1{F0: 9223372036854775807, F1: -9223372036854775808, F2: 1}
                    _ = v16
                case "foo bar":
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                case "hello":
                    var v17 uint64 = uint64(v2)
                    _ = v17
                }
                _ = fi13
            }
        default:
            v18 := GenS2{F0: 836, F1: (true && false)}
            _ = v18
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            v6 += v6
            var v20 uint64 = 42
            var v21 uint = (v6 % ((617) | 1))
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    fmt.Println(v2, v3, v5)
    var v22 uint16 = v2
    var v23 float64 = (0.1 + (3.14159 + 0.0))
    if false {
        var v24 int32 = -2147483648
        var v25 uint16 = v22
        _ = v24
        _ = v25
    } else {
        var v26 uint8 = (v7 + v7)
        for fi27 := 0; fi27 < 11; fi27++ {
            switch (-(fi27)) {
            case 128:
                fmt.Println(v5, v2)
            case 946:
                v28 := GenS3{F0: v26, F1: ((((v9 + (("" + v9) + "hello")) + "a") + ("a" + "\n\t")) + v9)}
                _ = v28
            }
            if false {
                var v29 bool = false
                fmt.Println(v6, v4)
                var v30 int32 = 42
                _ = v29
                _ = v30
            }
            _ = fi27
        }
        _ = v26
    }
    var v31 uint32 = 42
    if false {
        var v32 float64 = v23
        _ = v32
    } else {
        var v33 int = 128
        v34 := GenS2{F0: 827, F1: true}
        _ = v33
        _ = v34
    }
    fmt.Println(v7)
    var v35 int16 = (v10 & (v10 | v10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
