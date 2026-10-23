package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int32; F1 uint; F2 uint }

func main() {
    var v1 float64 = (3.14159 + (3.14159 / 3.14159))
    var v2 int8 = -1
    var v3 float32 = 1.5
    v2 -= v2
    fmt.Println(v2, v2, v1)
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 float64 = 0.1
        var v6 uint = 656
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 string = "foo bar"
    if (42 == 789) {
        var v8 int32 = -2147483648
        _ = v8
    } else {
        v2 -= (v2 | 77)
        switch v7 {
        case "foo bar":
            for fi9 := 0; fi9 < 13; fi9++ {
                switch 4294967295 {
                case 32768:
                    if false {
                        var v10 float32 = v3
                        v11 := GenS1{F0: false, F1: -32768, F2: 1}
                        for fi12 := 0; fi12 < 20; fi12++ {
                            if false {
                                switch (v7 + (("hello" + "foo bar") + (v7 + (((v7 + (v7 + v7)) + "hello") + ("hello" + v7))))) {
                                case "hello":
                                    v13 := GenS1{F0: (v2 != v2), F1: 32767, F2: 0}
                                    _ = v13
                                default:
                                    v14 := []int{0, (fi12 / (((^(fi12))) | 1))}
                                    _ = v14
                                }
                            }
                            var v15 int16 = -104
                            _ = fi12
                            _ = v15
                        }
                        _ = v10
                        _ = v11
                    }
                case 2147483647:
                    if (559 != 2147483648) {
                        v2 = (v2 | v2)
                        v16 := GenS2{F0: ((-1 <= 990) && true), F1: 2147483648}
                        var v17 uint32 = 0
                        _ = v16
                        _ = v17
                    }
                default:
                    v18 := GenS3{F0: 198, F1: int64(fi9)}
                    _ = v18
                }
                v1 = (v1 * ((1.5 * v1) + float64(v3)))
                _ = fi9
            }
        }
    }
    var v19 int32 = -1
    var v20 int32 = (v19 >> ((42) & 63))
    switch -2147483648 {
    case 42:
        v20 -= v20
    case -163:
        v21 := GenS1{F0: (true || ((!("foo bar" == "foo bar")) || true)), F1: -1, F2: 1}
        _ = v21
    case 977:
        switch 32767 {
        case 4294967295:
            v22 := GenS3{F0: v20, F1: 40}
            _ = v22
        case 404:
            v19 -= 2147483647
        default:
            v19 += v19
        }
    default:
        var v23 bool = (!(false || (false || false)))
        _ = v23
    }
    var v24 uint16 = 90
    var v25 uint = 255
    v3 -= ((v3 * v3) / float32(((v1 / 1e10) - ((v1 / v1) / v1))))
    fmt.Println(v20)
    fmt.Println(v3, v1, v25)
    v24 = v24
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 uint32 = 2
        _ = fi26
        _ = v27
    }
    if (40 != 2) {
        v28 := GenS4{F0: v19, F1: (^(v25)), F2: v25}
        if true {
            fmt.Println(v1, v24, v7)
            fmt.Println(v7, v2)
            var v29 uint64 = 42
            _ = v29
        }
        _ = v28
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        var v31 uint64 = 4294967295
        var v32 string = ("a" + (v7 + ("hello" + v7)))
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v7
}
