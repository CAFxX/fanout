package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 string }
type GenS4 struct { F0 uint16; F1 int8 }

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 42
    v3 := []int{127, 32768}
    var v4 float32 = (v1 + v1)
    var v5 uint32 = 108
    v6 := GenS1{F0: 127, F1: (v5 * 0)}
    var v7 uint = (v2 / ((v2) | 1))
    var v8 bool = (0.1 != 1.5)
    for fi9 := 0; fi9 < 10; fi9++ {
        if true {
            v4 += 1e10
            if (635 != (v5 & v5)) {
                v10 := GenS1{F0: 127, F1: (^(v5))}
                var v11 int16 = 32767
                _ = v10
                _ = v11
            }
        } else {
            var v12 int16 = 32767
            _ = v12
        }
        v13 := GenS1{F0: 42, F1: (^(v5))}
        _ = fi9
        _ = v13
    }
    v3 = append(v3, 934)
    v4 = v1
    _ = len(v3)
    if v8 {
        switch 919 {
        case 127:
            v3 = append(v3, -9223372036854775808)
        case 32768:
            v7 += (v2 << (((v7 - v2)) & 63))
        default:
            if len(v3) > 0 {
                v3[0] = -631
            }
        }
        v14 := GenS2{F0: 2}
        var v15 int64 = 9223372036854775807
        _ = v14
        _ = v15
    }
    if true {
        var v16 uint32 = 42
        switch ("foo bar" + "\n\t") {
        case "":
            v4 = v1
        case "hello":
            v3 = append(v3, 9223372036854775807)
        case "foo bar":
            v17 := GenS3{F0: v2, F1: 831, F2: ("hello" + (("hello" + ("hello" + "\n\t")) + "foo bar"))}
            _ = v17
        default:
            var v18 int16 = 78
            _ = v18
        }
        _ = v16
    }
    v19 := GenS3{F0: (v7 - 1), F1: 418, F2: ("a" + "")}
    for fi20 := 0; fi20 < 7; fi20++ {
        switch (^(fi20)) {
        case -9223372036854775808:
            _ = len(v3)
        case 32768:
            var v21 int64 = 816
            _ = v21
        case 0:
            v8 = ((!(((^(v5)) < (^(v5))) && v8)) && true)
        }
        v22 := GenS4{F0: uint16(v2), F1: 0}
        var v23 string = ("hello" + "foo bar")
        _ = fi20
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: uint8(v7), F1: (v5 + v5)}
    v4 += 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
