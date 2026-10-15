package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int = 1
    var v4 float64 = 3.14159
    var v5 uint64 = 876
    var v6 uint = 255
    if false {
        var v7 float64 = 0.0
        if true {
            var v8 int32 = 0
            _ = v8
        }
        _ = v7
    } else {
        var v9 float32 = (3.14159 - float32((((v4 * 0.0) - (1e10 * (v4 / 1.0))) + (v4 - (v4 - ((v4 / 1.0) - v4))))))
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 int = 125
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(v4, fi12)
                _ = fi12
            }
            v13 := GenS1{F0: -130}
            _ = fi10
            _ = v11
            _ = v13
        }
        _ = v9
    }
    var v14 float32 = 1.5
    delete(v2, 782)
    if false {
        if false {
            v14 += 0.0
        }
    }
    fmt.Println(v5, v5, v4)
    var v15 float64 = 1.5
    var v16 uint64 = v5
    var v17 uint32 = 1
    v18 := []int{}
    switch "\n\t" {
    case "":
        v15 = v15
    case "a":
        for fi19 := 0; fi19 < 1; fi19++ {
            if true {
                if false {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        fmt.Println(v4, v5, fi19)
                        delete(v2, (^(v17)))
                        _ = fi20
                    }
                }
                v21 := GenS2{F0: -1}
                _ = v21
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                if false {
                    var v23 int8 = 43
                    _ = v23
                } else {
                    fmt.Println(v17)
                    v24 := GenS3{F0: false}
                    _ = v24
                }
                _ = len(v1)
                _ = fi22
            }
            var v25 bool = (!(("" + "\n\t") > ((((("" + ("\n\t" + "")) + "\n\t") + "\n\t") + "hello") + ("foo bar" + "hello"))))
            _ = fi19
            _ = v25
        }
    case "foo bar":
        v4 = v4
    default:
        v26 := GenS3{F0: true}
        _ = v26
    }
    fmt.Println(v6, v16, v6)
    v15 = v4
    var v27 uint8 = 25
    if (!((18446744073709551615 >= v6) && (((v16 << ((v6) & 63)) <= 187) && false))) {
        v2[v17] = 424
    } else {
        var v28 string = "\n\t"
        var v29 uint16 = 32767
        _ = v28
        _ = v29
    }
    v30 := GenS4{F0: (v27 & v27), F1: 42}
    v18 = append(v18, v3)
    v17 = (v17 * (v17 / (((v17 % (((v17 ^ v17)) | 1))) | 1)))
    if true {
        v27 += v27
    } else {
        v31 := GenS5{F0: 644, F1: true}
        var v32 int32 = 659
        if ((!(598 < 0)) || (71 >= 2)) {
            v33 := GenS2{F0: 1000}
            var v34 uint64 = v16
            _ = v33
            _ = v34
        } else {
            fmt.Println(v3)
        }
        _ = v31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
