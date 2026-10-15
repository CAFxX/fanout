package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 bool }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int64, p1 bool) int16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 bool = p1
        var v9 int16 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{4294967295, -335, 611}
    _ = p0
    _ = p1
    _ = v10
    return -32768
}
func genFn1(p0 uint, p1 uint64, p2 uint64) uint {
    v11 := GenS1{F0: 0, F1: 128, F2: true}
    if true {
        if false {
            for fi12 := 0; fi12 < 2; fi12++ {
                if (false && ((0 >= 0) && true)) {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 bool = false
                        _ = fi13
                        _ = v14
                    }
                } else {
                    p0 += (p0 & 18446744073709551615)
                }
                var v15 int64 = 9223372036854775807
                for fi16 := 0; fi16 < 9; fi16++ {
                    p1 -= p2
                    p0 = (p0 >> ((p0) & 63))
                    fmt.Println(p2, p1, p1)
                    _ = fi16
                }
                _ = fi12
                _ = v15
            }
            v17 := GenS2{F0: 32767, F1: 362}
            _ = v17
        } else {
            v18 := GenS3{F0: 56, F1: -1, F2: 127}
            var v19 uint8 = 255
            var v20 int32 = 0
            _ = v18
            _ = v19
            _ = v20
        }
        if ((!(!false)) && true) {
            v21 := []int{}
            fmt.Println(p2, p2)
            v22 := GenS4{F0: 0, F1: true, F2: 65535}
            _ = v21
            _ = v22
        }
    }
    if false {
        v23 := GenS3{F0: 100, F1: 1, F2: 255}
        _ = v23
    }
    var v24 uint16 = 65535
    v25 := GenS1{F0: -1, F1: -550, F2: (false || (!false))}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v24
    _ = v25
    return 117
}

func main() {
    var v1 int8 = 1
    var v2 uint64 = 18446744073709551615
    var v3 int64 = 564
    v4 := make(map[int]int)
    var v5 uint64 = v2
    var v6 float32 = -2.25
    var v26 int8 = v1
    v1 = (v26 + v1)
    switch ("\n\t" + "a") {
    case "foo bar":
        switch "" {
        case "foo bar":
            var v27 uint32 = 2147483648
            _ = v27
        case "\n\t":
            v28 := []int{}
            _ = v28
        case "hello":
            fmt.Println(v1, v5, v5)
        }
    case "":
        v3 = 318
    case "\n\t":
        v29 := v4[32767]
        _ = v29
    }
    var v30 int32 = 748
    switch "foo bar" {
    case "foo bar":
        v31 := GenS5{F0: (v1 - v1)}
        _ = v31
    case "a":
        var v32 uint = 741
        _ = v32
    }
    v1 = v1
    var v33 uint32 = 2147483648
    fmt.Println(v26)
    v4[32767] = 0
    var v34 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
