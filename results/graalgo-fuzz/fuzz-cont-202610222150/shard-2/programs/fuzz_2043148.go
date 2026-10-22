package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint }
type GenS3 struct { F0 uint64 }
func genFn0() int64 {
    if true {
        v6 := []int{-2147483648, 127}
        if true {
            if true {
                var v7 int = 128
                v7 = 128
                switch v7 {
                case -385:
                    v8 := GenS1{F0: 1000, F1: 42}
                    _ = v8
                default:
                    for fi9 := 0; fi9 < 18; fi9++ {
                        v10 := GenS1{F0: -1, F1: 1}
                        fmt.Println(fi9, fi9)
                        _ = fi9
                        _ = v10
                    }
                }
                _ = v7
            } else {
                _ = len(v6)
            }
        } else {
            var v11 uint32 = 262
            _ = v11
        }
        switch -9223372036854775808 {
        case 1:
            var v12 int16 = 996
            _ = v12
        case -1:
            if (true || true) {
                v13 := GenS2{F0: false, F1: 42, F2: 2}
                for fi14 := 0; fi14 < 9; fi14++ {
                    fmt.Println(fi14)
                    for fi15 := 0; fi15 < 2; fi15++ {
                        var v16 int8 = 11
                        var v17 uint64 = 419
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
                _ = v13
            } else {
                var v18 int16 = 1
                _ = v18
            }
        }
        _ = v6
    } else {
        v19 := GenS3{F0: 324}
        _ = v19
    }
    var v20 uint32 = 275
    if (!true) {
        v21 := GenS3{F0: 4294967295}
        _ = v21
    } else {
        var v22 bool = true
        switch 1 {
        case -2147483648:
            v22 = ("hello" < "")
        }
        v23 := []int{282, -9223372036854775808, 128, 0}
        _ = v22
        _ = v23
    }
    _ = v20
    return 843
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[uint32]bool)
    var v3 int16 = 42
    var v4 float32 = float32(0.0)
    var v5 int16 = 141
    switch 4294967295 {
    case 2147483647:
        if (397 <= 2147483647) {
            var v24 uint16 = 1
            _ = v24
        } else {
            v4 = 0.0
            var v25 uint16 = 42
            if false {
                for fi26 := 0; fi26 < 7; fi26++ {
                    var v27 bool = (false && true)
                    var v28 float64 = 0.0
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v29 := v2[42]
                if false {
                    fmt.Println(v5, v5, v5)
                } else {
                    if (((("\n\t" + "") + ("a" + "hello")) + "\n\t") != ("a" + ("foo bar" + "hello"))) {
                        delete(v2, 1)
                    }
                }
                _ = v29
            }
            _ = v25
        }
    case 1:
        var v30 float64 = -2.25
        _ = v30
    case 42:
        v1 -= v1
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        switch "a" {
        case "\n\t":
            if false {
                v3 = 233
            }
        case "hello":
            var v32 float32 = v4
            _ = v32
        case "a":
            _ = len(v2)
        }
        _ = fi31
    }
    var v33 float32 = v4
    v4 = 0.0
    delete(v2, 403)
    v3 = v3
    v1 -= (-(v1))
    switch v1 {
    case 32768:
        v3 = -32768
    }
    v1 -= 32767
    fmt.Println(v4)
    if false {
        var v34 bool = (!(false && (0 < 127)))
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
