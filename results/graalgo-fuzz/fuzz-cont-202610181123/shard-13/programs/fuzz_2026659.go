package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 string }
func genFn0(p0 uint16) uint16 {
    switch 128 {
    case 32767:
        v5 := []int{53, 861, -9223372036854775808}
        _ = v5
    case -607:
        p0 += (^(p0))
    default:
        v6 := GenS1{F0: 197, F1: 2147483648, F2: "a"}
        _ = v6
    }
    v7 := []int{2147483647}
    var v8 int = 569
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v7
    _ = v8
    return p0
}

func main() {
    v1 := []uint32{737, 282, 2}
    var v2 int32 = 42
    var v3 uint = 2
    var v4 uint32 = 0
    var v9 string = (("a" + "a") + "\n\t")
    switch v9 {
    case "":
        var v10 string = (("a" + "hello") + v9)
        _ = v10
    }
    switch "a" {
    case "\n\t":
        var v11 int64 = 1
        _ = v11
    case "a":
        if false {
            v4 -= 0
            v2 = (v2 + v2)
            for fi12 := 0; fi12 < 10; fi12++ {
                _ = len(v1)
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 float32 = (-2.25 / -2.25)
                    v3 += 2
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        } else {
            for fi15 := 0; fi15 < 4; fi15++ {
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 int32 = -2147483648
                    v18 := GenS1{F0: 0, F1: v4, F2: ""}
                    v19 := GenS1{F0: 2, F1: (^(v4)), F2: (v9 + v9)}
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                switch "foo bar" {
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "foo bar":
                    fmt.Println(v3, v3, v2)
                default:
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 uint = v3
                        switch v9 {
                        case "":
                            v21 += (v3 % ((42) | 1))
                        default:
                            v22 := GenS1{F0: 148, F1: v4, F2: v9}
                            _ = v22
                        }
                        _ = fi20
                        _ = v21
                    }
                }
                _ = fi15
            }
            var v23 int32 = v2
            v24 := GenS1{F0: 167, F1: v4, F2: (("hello" + v9) + ((v9 + v9) + "foo bar"))}
            _ = v23
            _ = v24
        }
    case "hello":
        switch 9223372036854775807 {
        case -1:
            var v25 bool = (false && (v2 > (v2 & v2)))
            _ = v25
        default:
            var v26 int16 = -770
            _ = v26
        }
    }
    _ = len(v1)
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v27 bool = true
    fmt.Println(v27, v9)
    fmt.Println(v3, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
