package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 string }
func genFn0() int {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 string = "hello"
        _ = fi6
        _ = v7
    }
    var v8 int8 = 66
    var v9 uint16 = 65535
    fmt.Println(v9)
    switch 127 {
    case 0:
        v10 := []int{431, 762}
        _ = v10
    case -9223372036854775808:
        v9 = 264
    case 581:
        if (v9 == (^(v9))) {
            v11 := []int{2147483647, 9223372036854775807, 4294967295}
            if ((^(v9)) == 703) {
                var v12 uint8 = 243
                var v13 int16 = 1000
                _ = v12
                _ = v13
            }
            var v14 uint16 = (v9 / (((v9 / ((v9) | 1))) | 1))
            _ = v11
            _ = v14
        }
    default:
        var v15 bool = (v8 != v8)
        _ = v15
    }
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 2147483647
    v3 := make(map[uint32]bool)
    var v4 uint16 = 2
    var v5 int32 = -748
    v16 := GenS1{F0: false}
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 float32 = -2.25
        v2 = v5
        _ = fi17
        _ = v18
    }
    var v19 int32 = v2
    switch "a" {
    case "":
        v20 := GenS1{F0: (!false)}
        _ = v20
    }
    fmt.Println(v19)
    v21 := GenS2{F0: "hello", F1: 986, F2: 85}
    if true {
        v19 -= v2
    } else {
        v22 := GenS3{F0: v4, F1: 1, F2: 42}
        fmt.Println(v4, v19, v4)
        _ = v22
    }
    var v23 float64 = (float64(-2.25) - 1.5)
    if true {
        v2 = (^(v19))
        var v24 int8 = 127
        v25 := GenS3{F0: (v4 << ((738) & 63)), F1: 2, F2: 128}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 uint = 214
        _ = fi26
        _ = v27
    }
    var v28 int16 = int16(v5)
    v29 := GenS3{F0: v4, F1: 966, F2: 75}
    v30 := GenS1{F0: true}
    switch 569 {
    case -489:
        v31 := GenS3{F0: v4, F1: 860, F2: 128}
        _ = v31
    default:
        fmt.Println(v4)
    }
    v32 := GenS4{F0: (v4 >> ((uint(v4)) & 63)), F1: true}
    v23 += v23
    switch "" {
    case "hello":
        var v33 uint64 = 0
        _ = v33
    case "\n\t":
        v28 += (v28 | v28)
    case "foo bar":
        for fi34 := 0; fi34 < 6; fi34++ {
            for fi35 := 0; fi35 < 6; fi35++ {
                v36 := GenS5{F0: (("a" + "a") + "")}
                _ = fi35
                _ = v36
            }
            var v37 int8 = -1
            _ = fi34
            _ = v37
        }
    default:
        for fi38 := 0; fi38 < 7; fi38++ {
            v19 = (v5 - v19)
            _ = fi38
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v21
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
}
