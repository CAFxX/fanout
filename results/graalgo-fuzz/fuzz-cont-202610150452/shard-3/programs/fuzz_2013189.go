package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int32 }
type GenS2 struct { F0 int8; F1 int16; F2 string }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint) uint16 {
    v5 := []int{4, -2147483648, 42}
    for fi6 := 0; fi6 < 8; fi6++ {
        _ = len(v5)
        _ = fi6
    }
    if (true || ((true && (!false)) || false)) {
        v7 := GenS1{F0: (0 > 2147483647), F1: 1, F2: -1}
        if (32768 <= int(p0)) {
            var v8 int32 = 1
            _ = v8
        } else {
            var v9 uint32 = 1
            _ = v9
        }
        var v10 uint8 = 42
        _ = v10
        _ = v7
    } else {
        var v11 int32 = -2147483648
        p0 -= p0
        if (("foo bar" + "a") >= "\n\t") {
            var v12 float32 = (1.5 - ((3.14159 * (-2.25 - 3.14159)) / 1.0))
            _ = v12
        }
        _ = v11
    }
    if false {
        if (!(!false)) {
            p0 += (^(p0))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        if ((4294967295 >= p0) && true) {
            p0 += (^(p0))
            fmt.Println(p0)
            _ = len(v5)
        }
    }
    var v13 int64 = -1
    _ = p0
    _ = v13
    _ = v5
    return 487
}

func main() {
    var v1 float32 = ((0.0 / 1.0) - 1.5)
    var v2 uint8 = 127
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    var v14 int = 935
    fmt.Println(v2, v1, v14)
    v15 := []int{4294967295}
    v16 := GenS1{F0: (127 < (v14 << ((1) & 63))), F1: -558, F2: -216}
    v17 := GenS2{F0: -1, F1: 1000, F2: ("hello" + "a")}
    var v18 float32 = (v1 / v1)
    var v19 string = "a"
    if (-1 < 100) {
        var v20 int = 2147483647
        _ = v20
    } else {
        v21 := GenS1{F0: false, F1: 42, F2: -2147483648}
        _ = v21
    }
    if (true || true) {
        if (((v2 - v2) <= v2) && (1.5 <= ((((v1 - v18) * 3.14159) + ((v18 / v1) * -2.25)) + v18))) {
            fmt.Println(v2)
            _ = len(v15)
        }
    }
    switch ("a" + "\n\t") {
    case "a":
        if true {
            v19 = v19
            fmt.Println(v18, v14, v14)
        } else {
            var v22 string = ""
            v23 := GenS3{F0: 42}
            _ = v22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v18 = (v1 - v18)
        switch 127 {
        case 2147483647:
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        }
        _ = fi24
    }
    var v25 int16 = 0
    var v26 int = v14
    switch v19 {
    case "":
        if true {
            var v27 uint32 = 442
            _ = v27
        }
    case "hello":
        v26 = 9223372036854775807
    default:
        var v28 uint32 = 2147483648
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
