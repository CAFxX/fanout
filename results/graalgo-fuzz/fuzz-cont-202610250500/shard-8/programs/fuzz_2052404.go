package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint32; F2 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int32; F2 int }
type GenS5 struct { F0 string; F1 uint }
func genFn0(p0 uint8, p1 int16, p2 uint64) int64 {
    var v5 int64 = -589
    fmt.Println(p1, v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return v5
}
func genFn1(p0 int64, p1 uint) int32 {
    var v6 uint8 = 0
    v7 := GenS1{F0: -128}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 255
    v2 := []int{128}
    v3 := []string{"foo bar", "foo bar", ("a" + "foo bar"), "hello"}
    var v4 uint32 = 2147483647
    if (false || true) {
        v1 = (v1 % (((^(v1))) | 1))
        v4 += (^(v4))
    }
    v1 = (v1 % ((128) | 1))
    var v8 float64 = 3.14159
    v4 = v4
    v9 := GenS1{F0: 100}
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 int32 = 255
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 uint = 255
            v4 = 470
            var v14 uint64 = 18446744073709551615
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    switch (("\n\t" + "\n\t") + "hello") {
    case "foo bar":
        fmt.Println(v4, v8)
    case "":
        v4 -= (^(v4))
    case "\n\t":
        switch ("foo bar" + "") {
        case "a":
            v15 := GenS2{F0: 42, F1: 0, F2: -1}
            _ = v15
        default:
            fmt.Println(v1, v4, v8)
        }
    }
    v16 := GenS3{F0: (4294967295 > 1), F1: v1}
    var v17 uint8 = 255
    fmt.Println(v8, v17, v17)
    for fi18 := 0; fi18 < 1; fi18++ {
        fmt.Println(fi18, v8, v8)
        _ = len(v2)
        for fi19 := 0; fi19 < 11; fi19++ {
            if true {
                v4 = (v4 + (v4 * (v4 ^ v4)))
                if false {
                    v20 := GenS3{F0: false, F1: 170}
                    var v21 uint8 = (v17 ^ v1)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v20
                    _ = v21
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch "a" {
                    case "a":
                        var v22 uint32 = (v4 * v4)
                        _ = v22
                    case "hello":
                        v23 := GenS3{F0: (32767 == 42), F1: v17}
                        _ = v23
                    }
                    var v24 float32 = (1e10 - ((1e10 / 1.5) + -2.25))
                    _ = v24
                }
                v4 -= 65535
                if true {
                    v25 := GenS4{F0: 322, F1: -1, F2: -2147483648}
                    _ = v25
                }
                v26 := GenS5{F0: "a", F1: 116}
                v8 -= 3.14159
                _ = v26
            }
            var v27 float64 = v8
            switch 165 {
            case 128:
                var v28 uint16 = 878
                _ = v28
            }
            _ = fi19
            _ = v27
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
