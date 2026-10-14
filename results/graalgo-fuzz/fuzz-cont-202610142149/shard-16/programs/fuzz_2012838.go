package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32) string {
    fmt.Println(p0)
    p0 = 1
    _ = p0
    return "\n\t"
}
func genFn1(p0 uint) int64 {
    if true {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 int16 = 1000
            _ = fi6
            _ = v7
        }
        p0 += (^(p0))
    }
    if true {
        v8 := []int{190}
        _ = v8
    } else {
        fmt.Println(p0)
    }
    if (true && false) {
        if true {
            if (!(!(false || (1 < 42)))) {
                switch 287 {
                case 32767:
                    var v9 float64 = 3.14159
                    _ = v9
                case 42:
                    fmt.Println(p0)
                case -930:
                    var v10 string = "\n\t"
                    _ = v10
                }
            } else {
                v11 := GenS1{F0: false}
                for fi12 := 0; fi12 < 3; fi12++ {
                    p0 += p0
                    v13 := GenS2{F0: -1}
                    switch 127 {
                    case 128:
                        p0 = (^(p0))
                    case 286:
                        p0 -= 579
                    }
                    _ = fi12
                    _ = v13
                }
                var v14 float32 = 1e10
                _ = v11
                _ = v14
            }
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 int8 = -1
                fmt.Println(fi15)
                _ = fi15
                _ = v16
            }
        }
        var v17 int64 = 42
        var v18 bool = (float32(float64(((1e10 - 1.5) - (((0.1 / 1.5) * 0.0) + 1.5)))) <= float32(float64(0.0)))
        _ = v17
        _ = v18
    }
    switch "a" {
    case "hello":
        p0 = (p0 % ((p0) | 1))
    case "a":
        var v19 int8 = 57
        _ = v19
    case "":
        v20 := GenS2{F0: 187}
        _ = v20
    }
    var v21 uint64 = 1
    _ = p0
    _ = v21
    return -855
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "hello"
    v3 := []int{}
    v4 := []int{-810, 530}
    var v5 uint64 = 1
    delete(v1, 2147483647)
    v5 += v5
    var v22 uint8 = 0
    v23 := GenS2{F0: 42}
    v5 = (^(v5))
    var v24 int16 = 42
    var v25 int16 = 1
    var v26 float32 = (0.0 / ((float32(0.1) - float32((0.1 - 1e10))) / ((((1e10 * ((1e10 / 1.0) + ((3.14159 + 1.5) - 0.1))) * ((1.5 + 1.5) + 0.0)) + 0.1) - -2.25)))
    var v27 bool = false
    v27 = true
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 float32 = 1.5
        v30 := GenS2{F0: int16(v22)}
        for fi31 := 0; fi31 < 19; fi31++ {
            if true {
                v26 = v29
                var v32 float64 = 0.0
                _ = len(v1)
                _ = v32
            } else {
                var v33 string = v2
                _ = v33
            }
            _ = fi31
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    var v34 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
