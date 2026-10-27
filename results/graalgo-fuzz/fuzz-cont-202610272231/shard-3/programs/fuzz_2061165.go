package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint; F1 int32; F2 int8 }

func main() {
    var v1 float64 = 3.14159
    var v2 int32 = 42
    v3 := make(map[string]uint32)
    var v4 float64 = float64(0.0)
    var v5 uint16 = 37
    if false {
        v6 := []int{2147483647}
        fmt.Println(v5, v4)
        var v7 uint32 = 981
        _ = v6
        _ = v7
    }
    fmt.Println(v1)
    v1 = (v4 + v4)
    switch -9223372036854775808 {
    case 2147483647:
        fmt.Println(v5, v4, v4)
    case -1:
        for fi8 := 0; fi8 < 11; fi8++ {
            v4 = v4
            _ = fi8
        }
    case 0:
        v1 = v1
    default:
        v9 := GenS1{F0: 0, F1: -96}
        _ = v9
    }
    v3[""] = 0
    switch ("hello" + "") {
    case "":
        if (18446744073709551615 <= 805) {
            switch -2147483648 {
            case 32768:
                var v10 int = 2147483647
                _ = v10
            case 42:
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 uint64 = 1
                    fmt.Println(v5, v4)
                    _ = fi11
                    _ = v12
                }
            default:
                fmt.Println(v4, v2, v4)
            }
            _ = len(v3)
        } else {
            for fi13 := 0; fi13 < 8; fi13++ {
                v1 = (v4 - v4)
                _ = fi13
            }
            if false {
                for fi14 := 0; fi14 < 4; fi14++ {
                    for fi15 := 0; fi15 < 9; fi15++ {
                        var v16 bool = (false && true)
                        var v17 float32 = -2.25
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
                switch ("a" + ("foo bar" + "foo bar")) {
                case "hello":
                    var v18 uint = 23
                    _ = v18
                case "\n\t":
                    for fi19 := 0; fi19 < 16; fi19++ {
                        v20 := GenS1{F0: 9223372036854775807, F1: 127}
                        fmt.Println(v5, v5, v4)
                        _ = fi19
                        _ = v20
                    }
                }
                v4 = (v1 * 1e10)
            }
        }
    case "a":
        var v21 uint64 = 944
        _ = v21
    default:
        v22 := []int{4294967295}
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        v4 -= 1e10
        _ = fi23
    }
    if false {
        var v24 float64 = v1
        for fi25 := 0; fi25 < 8; fi25++ {
            v26 := GenS1{F0: -952, F1: -112}
            _ = fi25
            _ = v26
        }
        v1 += v24
        _ = v24
    } else {
        if true {
            var v27 int16 = -665
            _ = v27
        }
    }
    v5 = v5
    v28 := []int{127, 344, 128, -2147483648}
    v5 = v5
    v29 := GenS2{F0: 1, F1: (v2 & -1), F2: 0}
    var v30 string = (("" + "foo bar") + "hello")
    _ = len(v28)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
