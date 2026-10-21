package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 string; F1 int8; F2 bool }

func main() {
    var v1 int8 = 1
    var v2 int = -9223372036854775808
    var v3 uint8 = 2
    var v4 float32 = 1e10
    var v5 bool = ((0.0 - 0.0) == 3.14159)
    var v6 float32 = (v4 / (v4 - (v4 * v4)))
    v7 := GenS1{F0: "", F1: 775, F2: 42}
    for fi8 := 0; fi8 < 7; fi8++ {
        if true {
            v9 := []int{v2, v2, fi8, v2}
            var v10 int = (fi8 ^ v2)
            v11 := GenS1{F0: "\n\t", F1: 2147483648, F2: uint64(v3)}
            _ = v10
            _ = v11
            _ = v9
        } else {
            v12 := GenS1{F0: "a", F1: 42, F2: 2}
            _ = v12
        }
        _ = fi8
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int32 = 300
        if v5 {
            v15 := []int{(v2 / (((fi13 * v2)) | 1)), fi13}
            if v5 {
                var v16 int8 = (v1 << ((223) & 63))
                _ = v16
            }
            for fi17 := 0; fi17 < 16; fi17++ {
                switch -560 {
                case 32768:
                    fmt.Println(v3)
                case 157:
                    fmt.Println(v3, v5, v4)
                case 42:
                    if false {
                        v18 := GenS2{F0: 42}
                        _ = v18
                    }
                }
                _ = fi17
            }
            _ = v15
        } else {
            v19 := GenS3{F0: 1, F1: v5}
            if v5 {
                v3 -= 105
            } else {
                var v20 uint64 = 2
                switch -2147483648 {
                case -1:
                    var v21 uint8 = v3
                    _ = v21
                case 42:
                    switch 127 {
                    case 32767:
                        v22 := []int{4294967295}
                        _ = v22
                    case 4294967295:
                        v23 := []int{int(v3), 127, -232}
                        _ = v23
                    }
                default:
                    v3 -= 2
                }
                switch "hello" {
                case "\n\t":
                    v4 -= v6
                case "hello":
                    var v24 uint64 = v20
                    _ = v24
                default:
                    v2 = v2
                }
                _ = v20
            }
            var v25 uint = 18446744073709551615
            _ = v19
            _ = v25
        }
        _ = fi13
        _ = v14
    }
    switch v2 {
    case 128:
        v4 = v6
    case 286:
        v26 := []int{(v2 * 0), v2, 1}
        _ = v26
    case -9223372036854775808:
        v4 += 0.1
    default:
        v1 -= -1
    }
    var v27 uint8 = 165
    v5 = v5
    var v28 uint16 = 0
    switch "hello" {
    case "":
        if v5 {
            v29 := []int{}
            _ = v29
        } else {
            v30 := []int{-1, v2}
            v3 -= 2
            _ = v30
        }
    case "foo bar":
        v31 := []int{}
        _ = v31
    default:
        if v5 {
            v3 -= (v27 ^ 42)
        } else {
            switch (^(v2)) {
            case -1:
                for fi32 := 0; fi32 < 7; fi32++ {
                    var v33 uint8 = (v3 >> ((18446744073709551615) & 63))
                    _ = fi32
                    _ = v33
                }
            case -2147483648:
                fmt.Println(v4, v2)
            default:
                var v34 uint = 830
                _ = v34
            }
            v4 = 0.0
        }
    }
    fmt.Println(v3, v28)
    v35 := GenS4{F0: "a", F1: 127, F2: false}
    var v36 uint8 = (^(v27))
    var v37 uint8 = v27
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
