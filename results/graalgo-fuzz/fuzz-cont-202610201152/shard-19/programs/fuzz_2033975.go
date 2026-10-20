package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }

func main() {
    var v1 int8 = -3
    var v2 uint8 = 1
    var v3 bool = false
    var v4 uint16 = 1
    var v5 int8 = v1
    var v6 uint16 = (^(v4))
    v7 := GenS1{F0: 1, F1: 255}
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                v5 = (v1 ^ 100)
                v3 = (("" + "hello") < "foo bar")
            } else {
                if true {
                    v4 = (v4 << ((94) & 63))
                    v10 := GenS1{F0: -9223372036854775808, F1: 42}
                    _ = v10
                }
                switch "foo bar" {
                case "\n\t":
                    switch "a" {
                    case "hello":
                        v3 = v3
                    }
                    v4 -= (v4 * 869)
                case "foo bar":
                    var v11 uint8 = (v2 * (v2 / ((173) | 1)))
                    v3 = ((v1 / (((^(v5))) | 1)) <= (v5 / (((v1 - v5)) | 1)))
                    _ = v11
                default:
                    if true {
                        v12 := []int{(fi8 | fi9), (fi8 % ((fi9) | 1))}
                        _ = v12
                    }
                }
                v4 = 1
            }
            var v13 uint64 = 882
            v2 = 1
            _ = fi9
            _ = v13
        }
        _ = fi8
    }
    v6 += 1
    v3 = v3
    v2 -= 42
    var v14 string = "hello"
    v15 := GenS1{F0: 127, F1: 0}
    switch ("foo bar" + ((v14 + (v14 + (v14 + "a"))) + v14)) {
    case "":
        for fi16 := 0; fi16 < 5; fi16++ {
            v4 = (v4 >> ((2) & 63))
            v17 := GenS1{F0: 262, F1: 550}
            _ = fi16
            _ = v17
        }
    case "a":
        v18 := []int{9223372036854775807, 32768}
        _ = v18
    }
    v19 := []int{-95, 1, 165, 4294967295}
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    var v20 int16 = 0
    fmt.Println(v1)
    switch "\n\t" {
    case "":
        var v21 bool = (-9223372036854775808 >= -1)
        _ = v21
    case "\n\t":
        v22 := GenS1{F0: 42, F1: 934}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
