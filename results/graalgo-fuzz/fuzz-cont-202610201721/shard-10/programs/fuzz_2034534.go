package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS2 struct { F0 int }
func genFn0() int16 {
    v5 := []int{-1, 1}
    var v6 int32 = -2147483648
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 uint32, p1 string, p2 uint) uint32 {
    var v7 int32 = -20
    v8 := GenS1{F0: 128, F1: -12, F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 127
    v3 := []string{"hello", "\n\t"}
    v4 := []int{-2147483648, 11}
    var v9 int32 = -726
    fmt.Println(v9, v2)
    v9 -= v9
    var v10 int8 = 42
    fmt.Println(v1, v10)
    v3 = append(v3, "")
    fmt.Println(v9)
    v9 = (^(v9))
    v11 := GenS1{F0: (v2 << ((42) & 63)), F1: 100, F2: 42}
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := GenS1{F0: 1, F1: v10, F2: (^(v2))}
            var v15 int8 = v10
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(v1)
        if false {
            var v16 int8 = (v10 | 1)
            if len(v4) > 0 {
                v4[0] = (fi12 + 2147483647)
            }
            _ = v16
        } else {
            switch ("hello" + "") {
            case "hello":
                var v17 uint = 648
                _ = v17
            case "a":
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case "\n\t":
                if true {
                    if (!false) {
                        v18 := GenS2{F0: 32768}
                        if ((("" + "a") + ("hello" + ("" + "hello"))) <= ("\n\t" + "")) {
                            var v19 int64 = -1
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 6; fi20++ {
                                v1 -= ((((v1 / -2.25) * v1) * float64(float32((v1 + 1e10)))) / v1)
                                _ = fi20
                            }
                        }
                        var v21 int64 = -9223372036854775808
                        _ = v18
                        _ = v21
                    }
                }
            }
            v1 = v1
            for fi22 := 0; fi22 < 4; fi22++ {
                v1 -= (((v1 - (v1 + v1)) * v1) + 1.5)
                switch int(v10) {
                case -1:
                    var v23 bool = (!(2 <= 565))
                    _ = v23
                case 9223372036854775807:
                    var v24 uint = 4294967295
                    _ = v24
                case 689:
                    var v25 uint32 = 2147483647
                    _ = v25
                default:
                    fmt.Println(v2, v9, v9)
                }
                fmt.Println(v1, v9, fi22)
                _ = fi22
            }
        }
        _ = fi12
    }
    v1 += v1
    var v26 uint16 = 8
    var v27 int8 = v10
    switch 2147483647 {
    case 4294967295:
        var v28 uint = 4294967295
        _ = v28
    }
    var v29 string = "foo bar"
    var v30 uint = 18446744073709551615
    var v31 uint = (v30 >> ((v30) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v9
}
