package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint64; F2 string }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint32; F1 string; F2 int }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 int8 = -1
    var v2 int64 = 1
    v3 := []int{1, 128, 4294967295, -2147483648}
    v4 := []uint32{1, 900, 2147483648}
    v5 := []int{32767, 127}
    var v6 string = "\n\t"
    v7 := GenS1{F0: false, F1: 0}
    fmt.Println(v1)
    v1 -= v1
    fmt.Println(v6, v1, v2)
    v6 = (((((v6 + v6) + ((v6 + (("foo bar" + v6) + v6)) + ("" + v6))) + ("foo bar" + ((("foo bar" + "foo bar") + "hello") + "foo bar"))) + v6) + "\n\t")
    if false {
        v2 += v2
        if true {
            var v8 int32 = 1
            _ = v8
        }
    } else {
        v9 := GenS2{F0: 42, F1: 236, F2: ""}
        _ = v9
    }
    v10 := GenS3{F0: 32767, F1: 2}
    var v11 uint8 = 255
    v6 = "hello"
    _ = len(v3)
    var v12 int64 = v2
    v1 = v1
    if (false && true) {
        if true {
            fmt.Println(v1)
        } else {
            if (!(false || ((-2.25 == 1.5) && (v6 != v6)))) {
                var v13 uint64 = 42
                var v14 int16 = 32767
                _ = v13
                _ = v14
            }
            switch 127 {
            case -622:
                var v15 uint16 = 0
                _ = v15
            default:
                var v16 uint8 = (^(v11))
                _ = v16
            }
            v1 += v1
        }
        if true {
            v17 := GenS4{F0: 42, F1: "", F2: 776}
            _ = v17
        }
        switch 1 {
        case 1:
            var v18 uint = 1
            _ = v18
        default:
            var v19 int64 = 0
            _ = v19
        }
    } else {
        v20 := GenS5{F0: 32767}
        _ = v20
    }
    v11 = v11
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v11, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
