package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 bool }
type GenS2 struct { F0 bool }
func genFn0() int {
    var v6 string = "foo bar"
    v6 = v6
    v7 := []int{32768, 32768, 4294967295, 127}
    v6 = ((((("a" + (((v6 + (v6 + ("a" + ((v6 + (v6 + v6)) + (v6 + "a"))))) + v6) + v6)) + "foo bar") + v6) + (v6 + v6)) + v6)
    _ = v6
    _ = v7
    return -545
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 2147483647
    var v3 int16 = 1
    v4 := make(map[uint32]uint32)
    var v5 string = "hello"
    var v8 uint64 = 2
    fmt.Println(v3)
    _ = len(v4)
    v9 := GenS1{F0: uint(v8), F1: 18446744073709551615, F2: ("" >= v5)}
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v10 uint64 = v8
        _ = v10
    }
    var v11 float64 = 0.0
    var v12 uint32 = 434
    switch v5 {
    case "foo bar":
        fmt.Println(v11)
    case "a":
        v11 = v11
    case "":
        switch "\n\t" {
        case "":
            v2 = (v2 ^ -2147483648)
        default:
            if (true || (v2 > v2)) {
                v13 := GenS2{F0: false}
                var v14 float64 = v11
                switch v5 {
                case "a":
                    var v15 float32 = (1e10 + 0.1)
                    _ = v15
                case "\n\t":
                    switch "foo bar" {
                    case "":
                        for fi16 := 0; fi16 < 15; fi16++ {
                            var v17 string = (("a" + (("a" + (v5 + v5)) + v5)) + "")
                            v8 = (v8 & 18446744073709551615)
                            var v18 uint64 = 688
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    default:
                        v12 -= v12
                    }
                case "":
                    switch genFn0() {
                    case 978:
                        fmt.Println(v2, v11, v3)
                    case 42:
                        var v19 bool = false
                        _ = v19
                    case -2147483648:
                        for fi20 := 0; fi20 < 15; fi20++ {
                            v21 := GenS2{F0: true}
                            v11 -= 0.1
                            var v22 int64 = 42
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    default:
                        v23 := []int{1, -932}
                        _ = v23
                    }
                }
                _ = v13
                _ = v14
            } else {
                if true {
                    delete(v4, 419)
                } else {
                    var v24 uint16 = 2
                    v12 -= v12
                    _ = v24
                }
                v8 = (v8 & (v8 & 2))
                var v25 uint64 = 42
                _ = v25
            }
        }
    }
    fmt.Println(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
