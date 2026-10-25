package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
func genFn0(p0 int, p1 int16) uint32 {
    fmt.Println(p1, p0)
    var v7 float32 = (-2.25 / 1.0)
    p1 += 224
    var v8 int64 = -1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := []string{(("a" + "hello") + "hello"), ("hello" + "")}
    v2 := []int{-2147483648, 183, -9223372036854775808, 760}
    var v3 int32 = -608
    var v4 uint32 = 2
    var v5 float32 = 1.5
    var v6 string = ("a" + "foo bar")
    var v9 uint8 = 255
    v6 = "hello"
    v10 := GenS1{F0: 255, F1: 42, F2: -123}
    v4 = 922
    var v11 int32 = 42
    fmt.Println(v6, v6, v5)
    switch v6 {
    case "\n\t":
        fmt.Println(v9, v3)
    case "a":
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 float64 = 3.14159
                switch ("\n\t" + (v6 + "a")) {
                case "hello":
                    if len(v1) > 0 {
                        v1[0] = v6
                    }
                case "a":
                    switch v6 {
                    case "foo bar":
                        v15 := GenS2{F0: 911, F1: -128, F2: 42}
                        _ = v15
                    case "hello":
                        fmt.Println(v11)
                    case "\n\t":
                        _ = len(v1)
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case "foo bar":
                    var v16 string = "hello"
                    _ = v16
                default:
                    v2 = append(v2, (fi13 % (((fi12 << ((994) & 63))) | 1)))
                }
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    default:
        if true {
            v6 = v6
        } else {
            var v17 uint8 = v9
            var v18 uint8 = 236
            switch "hello" {
            case "a":
                for fi19 := 0; fi19 < 14; fi19++ {
                    var v20 uint64 = 333
                    var v21 int8 = int8(fi19)
                    v11 = 2147483647
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            case "":
                v2 = append(v2, 128)
            default:
                switch 127 {
                case 769:
                    v17 = v17
                default:
                    v22 := GenS1{F0: 4294967295, F1: 1, F2: 100}
                    _ = v22
                }
            }
            _ = v17
            _ = v18
        }
    }
    var v23 bool = (!(!false))
    var v24 string = "a"
    v25 := GenS3{F0: (0 <= 1)}
    var v26 uint16 = 126
    var v27 string = v6
    var v28 uint64 = 104
    if (false || true) {
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 uint64 = (^(v28))
            v6 = ("hello" + v6)
            v26 -= v26
            _ = fi29
            _ = v30
        }
    } else {
        v31 := GenS4{F0: 32768}
        if len(v1) > 0 {
            v1[0] = v6
        }
        v1 = append(v1, ((v24 + (("\n\t" + (v6 + (v27 + v24))) + v6)) + ("a" + v6)))
        _ = v31
    }
    var v32 uint8 = (v9 * v9)
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 uint64 = v28
        fmt.Println(v11, v5, v24)
        var v35 string = "\n\t"
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
