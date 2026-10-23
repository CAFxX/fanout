package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int8; F1 int32; F2 bool }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 string; F1 int; F2 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int; F1 int8; F2 string }
func genFn0(p0 uint16, p1 string) uint16 {
    v7 := GenS1{F0: 0, F1: 128, F2: 39}
    v8 := []int{32767, 128}
    p1 = ("foo bar" + "a")
    _ = len(v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 964
}
func genFn1(p0 int16) int16 {
    var v9 int32 = 42
    if true {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 bool = false
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 float32 = 0.0
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    } else {
        fmt.Println(p0)
    }
    var v14 bool = (1e10 < 0.1)
    v15 := GenS1{F0: 0, F1: 477, F2: 77}
    _ = p0
    _ = v14
    _ = v15
    _ = v9
    return 1000
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 42
    var v3 string = ("foo bar" + ("" + ("hello" + ("" + ("" + "")))))
    var v4 int64 = 9223372036854775807
    v5 := []int{}
    v6 := make(map[int]uint32)
    switch -2147483648 {
    case 32767:
        switch 32768 {
        case 2147483647:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case -707:
            v16 := GenS2{F0: v2}
            _ = v16
        case 331:
            for fi17 := 0; fi17 < 1; fi17++ {
                v5 = append(v5, 983)
                _ = fi17
            }
        default:
            v4 = (v4 - (v4 + 42))
        }
    case 25:
        var v18 int8 = 42
        _ = v18
    default:
        v3 = ((v3 + v3) + ("hello" + (v3 + (v3 + "hello"))))
    }
    switch v3 {
    case "hello":
        fmt.Println(v2, v3)
    case "":
        if (!true) {
            v3 = "foo bar"
        } else {
            v3 = v3
            fmt.Println(v4)
            var v19 int16 = -1
            _ = v19
        }
    case "foo bar":
        v20 := GenS3{F0: v4, F1: -32768}
        _ = v20
    default:
        v21 := GenS4{F0: 0, F1: 1, F2: (-331 > (v4 - v4))}
        _ = v21
    }
    v3 = (("hello" + (v3 + ((((v3 + v3) + v3) + ((("\n\t" + (v3 + v3)) + (("hello" + v3) + v3)) + "")) + v3))) + (v3 + (((v3 + "a") + "\n\t") + v3)))
    v2 = v2
    v22 := GenS5{F0: 2147483648, F1: 0}
    v4 = (v4 >> ((v2) & 63))
    for fi23 := 0; fi23 < 6; fi23++ {
        if (112 != 42) {
            v24 := GenS6{F0: v3, F1: fi23, F2: 493}
            _ = v24
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS5{F0: 2, F1: (v4 << ((v2) & 63))}
            _ = fi25
            _ = v26
        }
        _ = fi23
    }
    if (true && false) {
        for fi27 := 0; fi27 < 8; fi27++ {
            switch fi27 {
            case 4294967295:
                var v28 int32 = -420
                _ = v28
            case 1:
                v29 := GenS7{F0: 595}
                _ = v29
            case -79:
                var v30 bool = (fi27 < int(v2))
                _ = v30
            default:
                switch "" {
                case "a":
                    v2 += 4294967295
                case "hello":
                    v2 = v2
                }
            }
            _ = fi27
        }
        v31 := GenS8{F0: 9}
        _ = v31
    }
    v32 := GenS9{F0: -480, F1: 100, F2: v3}
    var v33 int16 = genFn1(-1)
    if (42 > v4) {
        if (v2 != 18446744073709551615) {
            _ = len(v5)
            _ = len(v5)
            v4 += v4
        }
        var v34 int64 = (v4 << ((v2) & 63))
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
