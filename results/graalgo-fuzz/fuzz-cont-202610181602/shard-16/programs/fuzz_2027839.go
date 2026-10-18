package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint32 }
func genFn0() uint32 {
    v6 := []int{}
    var v7 uint8 = 142
    var v8 bool = true
    v9 := GenS1{F0: int(v7), F1: 616}
    var v10 float32 = 1.5
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 127
    var v3 uint16 = 255
    var v4 uint8 = 10
    var v5 uint8 = v4
    v11 := GenS1{F0: int(v2), F1: -716}
    if false {
        var v12 int64 = 9223372036854775807
        _ = v12
    } else {
        v13 := GenS2{F0: (false || (true && (!(false && ((1.5 * float64(((3.14159 / ((0.1 * 1e10) + 3.14159)) * (0.1 + float32(float64(1.5)))))) != 1.5))))), F1: 128}
        var v14 bool = false
        _ = v13
        _ = v14
    }
    v2 = v2
    v4 = (v5 | (^(v4)))
    var v15 int8 = v2
    _ = len(v1)
    var v16 float64 = (3.14159 - 0.0)
    _ = len(v1)
    v17 := GenS3{F0: 646}
    var v18 string = "foo bar"
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS2{F0: (false && false), F1: 1}
        if false {
            var v21 uint16 = 32767
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS1{F0: fi22, F1: fi22}
        switch (fi22 << ((4294967295) & 63)) {
        case -1:
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 string = (v18 + "\n\t")
                _ = fi24
                _ = v25
            }
        case -842:
            v5 += (v5 >> ((18446744073709551615) & 63))
        case 128:
            if true {
                var v26 uint32 = 170
                fmt.Println(v16, v15)
                _ = v26
            }
        }
        var v27 int16 = 0
        _ = fi22
        _ = v23
        _ = v27
    }
    _ = len(v1)
    if true {
        if true {
            var v28 string = "\n\t"
            _ = v28
        } else {
            v29 := v1[636]
            _ = v29
        }
        v30 := GenS4{F0: (v2 << ((uint(v5)) & 63)), F1: uint32(v3)}
        _ = v30
    } else {
        switch v18 {
        case "":
            var v31 int8 = 127
            _ = v31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
