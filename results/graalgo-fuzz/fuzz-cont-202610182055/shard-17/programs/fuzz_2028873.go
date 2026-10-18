package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int16 = -659
    var v2 uint64 = 2
    var v3 int16 = 1000
    var v4 float64 = 0.1
    v3 += (v3 / ((v1) | 1))
    var v5 int16 = (v3 & (v1 ^ v3))
    switch ("hello" + "hello") {
    case "a":
        v6 := []int{4294967295}
        _ = v6
    case "foo bar":
        if (!true) {
            v7 := []int{9223372036854775807, 9223372036854775807, 726, 128}
            if true {
                var v8 int = 9223372036854775807
                var v9 int32 = 0
                fmt.Println(v8)
                _ = v8
                _ = v9
            } else {
                var v10 uint64 = v2
                switch "" {
                case "\n\t":
                    var v11 uint64 = v2
                    _ = v11
                case "a":
                    fmt.Println(v5, v1, v2)
                default:
                    v5 += (v3 ^ (v3 | (v5 / ((v1) | 1))))
                }
                _ = v10
            }
            _ = v7
        }
    case "":
        v12 := []int{}
        _ = v12
    default:
        var v13 string = ("a" + ("" + ((("" + "") + "a") + "")))
        _ = v13
    }
    v5 = -1
    v14 := GenS1{F0: 100}
    v15 := []int{4294967295, 0, -9223372036854775808}
    switch "" {
    case "a":
        v1 = v1
    case "foo bar":
        if true {
            v15 = append(v15, -71)
            v16 := GenS2{F0: (((v2 / ((1) | 1)) != v2) && true), F1: -548}
            _ = v16
        }
    default:
        if (false || (-1 < 100)) {
            var v17 string = "\n\t"
            var v18 int8 = 1
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 6; fi19++ {
                fmt.Println(v1, v3, v1)
                v2 += v2
                v1 += 1000
                _ = fi19
            }
        }
    }
    _ = len(v15)
    v20 := GenS1{F0: 100}
    fmt.Println(v4, v3, v3)
    v21 := GenS3{F0: 255, F1: 9223372036854775807}
    var v22 uint8 = 1
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 bool = (217 < 65535)
        _ = fi23
        _ = v24
    }
    v25 := GenS3{F0: v22, F1: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
