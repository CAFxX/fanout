package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
func genFn0() uint64 {
    var v5 uint64 = 400
    v6 := GenS1{F0: 127}
    v7 := GenS1{F0: 127}
    fmt.Println(v5)
    _ = v5
    _ = v6
    _ = v7
    return (^(v5))
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -890
    var v3 uint = 65535
    var v4 uint = (v3 >> (((v3 / ((v3) | 1))) & 63))
    delete(v1, 9223372036854775807)
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: fi8}
        switch "a" {
        case "":
            v1[(fi8 % ((2147483647) | 1))] = "\n\t"
        }
        var v10 float32 = 0.0
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v2, v4, v2)
    var v11 uint64 = 42
    fmt.Println(v2, v2)
    var v12 float64 = -2.25
    switch 42 {
    case 42:
        switch "hello" {
        case "hello":
            if true {
                fmt.Println(v4, v3)
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 int = 618
                    _ = fi13
                    _ = v14
                }
            }
        case "\n\t":
            v2 -= v2
        case "foo bar":
            for fi15 := 0; fi15 < 3; fi15++ {
                fmt.Println(v2)
                fmt.Println(fi15, v12, v11)
                _ = fi15
            }
        default:
            var v16 uint = v4
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int32 = 0
            v18 -= v18
            _ = fi17
            _ = v18
        }
    }
    if (!false) {
        switch "\n\t" {
        case "\n\t":
            v2 += v2
        default:
            v19 := GenS1{F0: 32767}
            _ = v19
        }
        v20 := GenS2{F0: v4}
        delete(v1, 9223372036854775807)
        _ = v20
    } else {
        v11 += (v11 % ((983) | 1))
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 bool = (true && false)
        if v22 {
            v4 += 797
            v23 := v1[fi21]
            v24 := GenS1{F0: (fi21 / ((fi21) | 1))}
            _ = v23
            _ = v24
        } else {
            v22 = v22
            v1[fi21] = "a"
            for fi25 := 0; fi25 < 10; fi25++ {
                switch "foo bar" {
                case "":
                    v11 = 1
                    var v26 int16 = 883
                    v27 := GenS2{F0: 0}
                    _ = v26
                    _ = v27
                default:
                    v2 = 9223372036854775807
                }
                var v28 int16 = 0
                var v29 int = (^(fi21))
                _ = fi25
                _ = v28
                _ = v29
            }
        }
        _ = fi21
        _ = v22
    }
    v30 := []int{-50, 51, 1, -9223372036854775808}
    if false {
        var v31 float64 = ((v12 / 0.1) - float64((0.1 + (1e10 / 1.0))))
        _ = v31
    } else {
        v4 += v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
