package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
func genFn0(p0 string) int {
    p0 = "foo bar"
    v5 := GenS1{F0: 1, F1: false, F2: 823}
    p0 = p0
    if (p0 >= p0) {
        if (!(!(!(true || (-570 >= 0))))) {
            var v6 int8 = -77
            var v7 uint = 2
            if (2 >= 702) {
                fmt.Println(v6, p0, v6)
            }
            _ = v6
            _ = v7
        }
    }
    _ = p0
    _ = v5
    return -120
}

func main() {
    var v1 int16 = -1
    var v2 int64 = 1
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 42
    v8 := GenS2{F0: genFn0("\n\t"), F1: 77}
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            v11 := GenS3{F0: 4294967295, F1: -1}
            var v12 int8 = 1
            if (!((v12 & (v12 >> ((353) & 63))) >= v12)) {
                v4 += v4
                var v13 int = (fi9 / (((fi10 & -827)) | 1))
                var v14 uint32 = 147
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    fmt.Println(v1, v1)
    v15 := GenS4{F0: uint8(v4)}
    v16 := GenS4{F0: 207}
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := []int{}
        v1 += (v1 >> ((2) & 63))
        if true {
            for fi19 := 0; fi19 < 16; fi19++ {
                v20 := GenS5{F0: 0}
                switch "" {
                case "\n\t":
                    if true {
                        v1 = v1
                        var v21 float64 = -2.25
                        var v22 uint16 = (v4 | v4)
                        _ = v21
                        _ = v22
                    }
                case "hello":
                    fmt.Println(fi19, v1)
                case "foo bar":
                    v23 := GenS5{F0: -984}
                    _ = v23
                }
                _ = fi19
                _ = v20
            }
            v3[4294967295] = 2147483648
        } else {
            fmt.Println(v2, fi17)
            fmt.Println(fi17, v4)
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v1, v2)
    v4 = 65535
    switch 0 {
    case 303:
        v2 += v2
    case 700:
        var v24 uint = 65535
        _ = v24
    default:
        fmt.Println(v2, v2, v4)
    }
    v4 -= v4
    v25 := []int{int(v1), 0, 9223372036854775807}
    fmt.Println(v2, v1)
    switch int(v1) {
    case 2147483647:
        for fi26 := 0; fi26 < 5; fi26++ {
            v4 += 255
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v8
}
