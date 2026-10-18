package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0() string {
    v6 := GenS1{F0: 128}
    var v7 int8 = 127
    switch 272 {
    case -678:
        v7 += (v7 - (v7 * -128))
    case 759:
        if true {
            var v8 int64 = -1
            v9 := []int{1}
            v10 := GenS1{F0: 1}
            _ = v10
            _ = v8
            _ = v9
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                v7 -= (v7 - v7)
                _ = fi11
            }
            if true {
                var v12 uint = 2
                fmt.Println(v7, v7)
                _ = v12
            } else {
                v13 := GenS2{F0: -449, F1: 42}
                var v14 uint8 = 128
                for fi15 := 0; fi15 < 7; fi15++ {
                    fmt.Println(fi15, v14, v7)
                    _ = fi15
                }
                _ = v13
                _ = v14
            }
        }
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS1{F0: (fi16 * (fi16 ^ fi16))}
        _ = fi16
        _ = v17
    }
    fmt.Println(v7)
    _ = v6
    _ = v7
    return ("hello" + "foo bar")
}
func genFn1(p0 uint32) uint16 {
    v18 := []int{}
    switch "\n\t" {
    case "a":
        p0 = p0
    }
    _ = p0
    _ = v18
    return 515
}

func main() {
    var v1 string = ""
    var v2 int16 = 0
    var v3 int64 = -619
    var v4 uint = 174
    var v5 uint = v4
    v3 = (^(v3))
    v5 = v5
    if true {
        fmt.Println(v4, v4)
        var v19 bool = false
        var v20 uint64 = 893
        _ = v19
        _ = v20
    }
    v21 := []int{128}
    v2 += (v2 >> (((v4 ^ v5)) & 63))
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 bool = false
        var v24 float32 = -2.25
        var v25 float64 = ((1.5 / -2.25) - 0.0)
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    v26 := GenS3{F0: 201, F1: 42}
    v3 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
