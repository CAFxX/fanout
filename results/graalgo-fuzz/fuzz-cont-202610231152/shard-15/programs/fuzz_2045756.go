package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 int64) uint64 {
    p0 = p0
    switch 42 {
    case -2147483648:
        p0 += p0
    case 128:
        switch 128 {
        case -267:
            for fi7 := 0; fi7 < 19; fi7++ {
                p0 = p0
                switch "hello" {
                case "foo bar":
                    var v8 float32 = ((-2.25 + float32(((1.5 / -2.25) + 0.1))) * float32(0.1))
                    _ = v8
                case "\n\t":
                    v9 := GenS1{F0: "", F1: (!false), F2: true}
                    _ = v9
                case "":
                    v10 := []int{fi7, fi7}
                    _ = v10
                }
                _ = fi7
            }
        case 128:
            v11 := GenS1{F0: ("\n\t" + "foo bar"), F1: ((p0 * p0) < (^(p0))), F2: (((float32(float64(float32(0.1))) * 1e10) + ((1e10 * (3.14159 * 1e10)) / ((1.5 / -2.25) / 1e10))) > 1.5)}
            _ = v11
        default:
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 bool = true
                _ = fi12
                _ = v13
            }
        }
    }
    var v14 int = 0
    p0 = (p0 << ((0) & 63))
    _ = p0
    _ = v14
    return 1
}

func main() {
    var v1 uint32 = 875
    var v2 float64 = 3.14159
    var v3 uint = 0
    var v4 float32 = 0.1
    var v5 uint8 = 127
    v6 := make(map[int]uint32)
    var v15 uint16 = 231
    switch 419 {
    case 127:
        fmt.Println(v3, v3, v3)
    }
    var v16 uint32 = v1
    v4 -= v4
    if (!("hello" >= ("" + "hello"))) {
        if ((!false) && true) {
            if true {
                switch ("\n\t" + "hello") {
                case "":
                    if (128 < 2) {
                        v5 = v5
                    } else {
                        v17 := GenS2{F0: 9223372036854775807}
                        fmt.Println(v2)
                        _ = v17
                    }
                }
                var v18 int32 = 0
                _ = v18
            }
            v4 = 0.0
            if true {
                var v19 bool = (true || false)
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS3{F0: 32768, F1: 127}
            v5 = v5
            var v22 int64 = 1
            _ = fi20
            _ = v21
            _ = v22
        }
    } else {
        fmt.Println(v3, v4, v1)
    }
    var v23 uint8 = (v5 & (v5 + (v5 % (((^(v5))) | 1))))
    var v24 int = 2147483647
    v25 := v6[(v24 | v24)]
    v26 := []int{(v24 | (v24 & v24)), v24}
    v4 = v4
    delete(v6, 912)
    fmt.Println(v2, v4, v2)
    var v27 int = 2147483647
    for fi28 := 0; fi28 < 6; fi28++ {
        _ = len(v26)
        _ = fi28
    }
    v29 := v6[v24]
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
