package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float32 = (0.1 / float32(1.5))
    var v2 int64 = 9223372036854775807
    var v3 int8 = 127
    var v4 uint = 42
    var v5 bool = (true || (int8(v2) == (v3 ^ (v3 & v3))))
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 float32 = v1
            _ = fi7
            _ = v8
        }
        var v9 string = "hello"
        var v10 int64 = -304
        _ = fi6
        _ = v10
        _ = v9
    }
    v2 += v2
    switch "a" {
    case "":
        v11 := []int{-999, 42}
        _ = v11
    case "a":
        v1 += ((v1 - (v1 + v1)) / 1e10)
    default:
        v3 += (-(v3))
    }
    fmt.Println(v4, v4)
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint32 = 760
        fmt.Println(v4)
        v14 := GenS1{F0: 73}
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 -= (v3 / ((v3) | 1))
    var v15 uint16 = 762
    switch 32767 {
    case 9223372036854775807:
        for fi16 := 0; fi16 < 7; fi16++ {
            fmt.Println(v15)
            _ = fi16
        }
    default:
        fmt.Println(v5, v3)
    }
    v17 := GenS2{F0: 186}
    switch 153 {
    case 572:
        v18 := GenS3{F0: 899}
        _ = v18
    case 128:
        if v5 {
            switch "a" {
            case "hello":
                v19 := []int{519, 128}
                _ = v19
            case "foo bar":
                if true {
                    if (true && false) {
                        v20 := []int{}
                        v5 = true
                        _ = v20
                    } else {
                        var v21 bool = (!false)
                        v22 := GenS2{F0: 1}
                        _ = v21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 16; fi23++ {
                        var v24 int32 = 1
                        v25 := []int{}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    v26 := []int{2147483647, -2147483648, 127}
                    _ = v26
                } else {
                    v4 += 0
                    var v27 uint = (v4 + (v4 << ((v4) & 63)))
                    for fi28 := 0; fi28 < 15; fi28++ {
                        var v29 int16 = 32767
                        v30 := GenS3{F0: 42}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = v27
                }
            default:
                switch "hello" {
                case "hello":
                    v3 -= (v3 % ((-128) | 1))
                case "a":
                    if false {
                        v31 := []int{}
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        var v32 bool = (v1 >= (1e10 * (float32((0.1 / 1.0)) * (v1 + v1))))
                        _ = v31
                        _ = v32
                    }
                case "\n\t":
                    v33 := GenS2{F0: 41}
                    _ = v33
                }
            }
            var v34 int16 = 1
            fmt.Println(v34, v15)
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
