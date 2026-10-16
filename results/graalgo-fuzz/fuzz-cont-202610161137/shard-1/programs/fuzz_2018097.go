package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int; F2 int }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint32 }
func genFn0(p0 string) uint8 {
    if (!true) {
        var v7 float64 = float64((-2.25 - float32(1e10)))
        _ = v7
    }
    if false {
        v8 := []int{}
        var v9 int = 127
        _ = v8
        _ = v9
    } else {
        switch 42 {
        case -806:
            var v10 float32 = (1e10 - (((-2.25 - (0.1 * 3.14159)) / ((1.5 + ((1.5 - (1e10 - (1e10 / 3.14159))) + 1e10)) + (3.14159 - float32(3.14159)))) / 3.14159))
            _ = v10
        }
        var v11 bool = (86 <= 255)
        var v12 int16 = 42
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: (48 == 42), F1: 36}
    _ = p0
    _ = v13
    return 2
}
func genFn1(p0 uint16, p1 int16, p2 string) int {
    p1 = p1
    if true {
        v14 := []int{127, 32767}
        fmt.Println(p1, p2, p0)
        _ = v14
    }
    p2 = "hello"
    p2 = "foo bar"
    switch -9223372036854775808 {
    case -1:
        p2 = p2
    case 128:
        switch (p2 + (p2 + p2)) {
        case "":
            var v15 int16 = (^(p1))
            _ = v15
        case "foo bar":
            switch "foo bar" {
            case "a":
                v16 := []int{}
                _ = v16
            case "hello":
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 uint32 = 809
                    v19 := GenS2{F0: 0}
                    fmt.Println(fi17, p2)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case "\n\t":
                var v20 uint = 4294967295
                _ = v20
            }
        default:
            v21 := GenS3{F0: 127, F1: 128, F2: -227}
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 32767
    var v3 float64 = (0.1 + -2.25)
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    v6 := []int{9223372036854775807, -681, 32768, -2147483648}
    switch genFn1(v2, -811, ((((v1 + v1) + "a") + "hello") + "")) {
    case -710:
        v2 -= 395
    case 32768:
        v22 := GenS2{F0: 0}
        _ = v22
    case 42:
        var v23 bool = (!false)
        _ = v23
    }
    if ((32767 != 236) || (42 > 42)) {
        var v24 uint32 = 916
        fmt.Println(v24, v1)
        _ = v24
    }
    v1 = v1
    var v25 int16 = 0
    v1 = v1
    v1 = ((v1 + v1) + v1)
    if false {
        var v26 uint64 = 0
        var v27 string = v1
        _ = v26
        _ = v27
    }
    v28 := GenS4{F0: 0, F1: 0, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
