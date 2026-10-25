package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int32; F1 int16; F2 int16 }
func genFn0(p0 bool) bool {
    var v4 int8 = 46
    v4 = 100
    if p0 {
        switch ("hello" + "") {
        case "":
            v5 := GenS1{F0: 1, F1: false}
            _ = v5
        case "a":
            var v6 int32 = -1
            _ = v6
        case "hello":
            v7 := GenS2{F0: p0, F1: 18446744073709551615, F2: 552}
            _ = v7
        }
    }
    _ = p0
    _ = v4
    return p0
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int = 42
    var v3 uint64 = 2
    v8 := GenS2{F0: (379 == -1), F1: 1, F2: 0}
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS3{F0: 967, F1: 362, F2: 127}
        v1 -= (-(v1))
        _ = fi9
        _ = v10
    }
    v3 += (v3 % ((v3) | 1))
    switch ("\n\t" + "foo bar") {
    case "a":
        fmt.Println(v1, v1)
    case "hello":
        var v11 float32 = 0.0
        _ = v11
    case "\n\t":
        if (!true) {
            var v12 int32 = 440
            v2 = v2
            _ = v12
        } else {
            v3 = (v3 & v3)
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: 32767, F1: ("a" == "\n\t")}
                v15 := []int{(fi13 - (fi13 % (((fi13 - fi13)) | 1))), -9223372036854775808, fi13}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    }
    v16 := GenS1{F0: -1, F1: true}
    fmt.Println(v3, v1)
    var v17 float64 = 3.14159
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 int16 = -32768
        if false {
            if true {
                var v20 int8 = -1
                _ = v20
            } else {
                v19 = v19
            }
            v21 := GenS4{F0: 4294967295, F1: 42}
            var v22 uint32 = 0
            _ = v21
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    switch v2 {
    case -1:
        var v23 float32 = 3.14159
        _ = v23
    case 1:
        v17 -= 3.14159
    case 429:
        fmt.Println(v2, v1)
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        v1 += (v1 >> ((473) & 63))
        _ = fi24
    }
    v1 -= v1
    v25 := GenS5{F0: -2147483648, F1: -32768, F2: 1000}
    if (42 != 100) {
        v1 = v1
        v3 -= 42
        switch (("\n\t" + "hello") + "\n\t") {
        case "foo bar":
            switch "foo bar" {
            case "foo bar":
                switch "\n\t" {
                case "":
                    var v26 uint32 = 2
                    _ = v26
                case "hello":
                    fmt.Println(v2, v17, v3)
                default:
                    v27 := []int{v2, v2}
                    _ = v27
                }
            case "a":
                var v28 int = -1
                _ = v28
            }
        case "a":
            v1 = (v1 & (v1 * (v1 % ((v1) | 1))))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v8
}
