package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint; F1 int16; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint }
type GenS4 struct { F0 int8; F1 uint8 }
func genFn0() uint {
    switch "foo bar" {
    case "":
        for fi4 := 0; fi4 < 9; fi4++ {
            fmt.Println(fi4)
            if true {
                switch "\n\t" {
                case "a":
                    var v5 uint64 = 2
                    _ = v5
                }
                var v6 int8 = -117
                v6 = v6
                _ = v6
            } else {
                var v7 float32 = 1e10
                v7 -= (3.14159 / (v7 - v7))
                var v8 int32 = -421
                _ = v7
                _ = v8
            }
            _ = fi4
        }
    case "hello":
        var v9 int64 = 42
        _ = v9
    }
    var v10 uint = 42
    fmt.Println(v10)
    v10 = v10
    _ = v10
    return 117
}

func main() {
    var v1 int8 = 100
    var v2 uint32 = 2147483648
    var v3 float64 = 3.14159
    for fi11 := 0; fi11 < 3; fi11++ {
        switch (("\n\t" + ("" + "hello")) + ("a" + "\n\t")) {
        case "\n\t":
            if (!true) {
                var v12 string = ""
                _ = v12
            } else {
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 int64 = 42
                    v3 -= ((v3 / v3) * (v3 + v3))
                    v14 += v14
                    v15 := []int{}
                    _ = len(v15)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
        case "":
            v16 := GenS1{F0: 784, F1: false}
            _ = v16
        }
        v17 := GenS2{F0: 0, F1: 1000, F2: -1}
        _ = fi11
        _ = v17
    }
    v18 := []int{-9223372036854775808, -9223372036854775808}
    v19 := GenS3{F0: 42, F1: 320, F2: genFn0()}
    fmt.Println(v3, v1, v2)
    if len(v18) > 0 {
        v18[0] = 4294967295
    }
    v18 = append(v18, 128)
    var v20 int8 = (v1 >> ((65535) & 63))
    if len(v18) > 0 {
        v18[0] = int(v2)
    }
    var v21 int16 = -814
    var v22 float32 = 1e10
    switch -71 {
    case 588:
        if (864 <= 951) {
            if false {
                for fi23 := 0; fi23 < 19; fi23++ {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        v3 = v3
                        fmt.Println(v22, v20)
                        var v25 string = ("" + "a")
                        _ = fi24
                        _ = v25
                    }
                    fmt.Println(v3, v2, v2)
                    _ = fi23
                }
            }
            _ = len(v18)
            switch 9223372036854775807 {
            case -2147483648:
                var v26 uint64 = 496
                _ = v26
            case 0:
                var v27 uint16 = 255
                _ = v27
            default:
                v18 = append(v18, 9223372036854775807)
            }
        } else {
            v22 = v22
        }
    case 0:
        for fi28 := 0; fi28 < 12; fi28++ {
            if false {
                var v29 float64 = v3
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                _ = v29
            } else {
                var v30 int8 = (-(v20))
                v2 = v2
                if true {
                    v31 := GenS4{F0: (v1 >> ((244) & 63)), F1: 0}
                    fmt.Println(v2)
                    _ = v31
                }
                _ = v30
            }
            v20 -= v1
            _ = fi28
        }
    case -1:
        v3 = (v3 * v3)
    }
    v2 = (v2 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
