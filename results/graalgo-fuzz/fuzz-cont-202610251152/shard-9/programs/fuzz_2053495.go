package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint32; F2 int8 }
type GenS4 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = (-2.25 / (1.5 * float32(((-2.25 / float64((1e10 * 3.14159))) * 1.5))))
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    var v5 int8 = 100
    var v6 int = 127
    v7 := GenS1{F0: -293}
    var v8 int = v6
    fmt.Println(v2)
    var v9 uint8 = 1
    v8 -= -9223372036854775808
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(v2, v8)
        v9 = (v9 | v9)
        _ = fi10
    }
    switch ("a" + "\n\t") {
    case "\n\t":
        delete(v1, 954)
    case "a":
        v5 = v5
    }
    fmt.Println(v8, v6)
    v9 = 127
    var v11 uint = 1
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := []int{v6, v8, int(v11), 32768}
        switch 32768 {
        case -9223372036854775808:
            var v14 uint = v11
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v11 = (v11 - v11)
        if true {
            var v16 int32 = 873
            for fi17 := 0; fi17 < 14; fi17++ {
                v18 := v1[1]
                var v19 int8 = (^(v5))
                v8 = (v8 | -129)
                v20 := GenS1{F0: -2147483648}
                v5 += v5
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = v16
        } else {
            v21 := v4[-66]
            for fi22 := 0; fi22 < 10; fi22++ {
                v11 = (^(v11))
                var v23 int64 = int64(v5)
                v24 := []int{759, v8, (^(v8))}
                v25 := GenS2{F0: 4294967295, F1: 1, F2: 65535}
                v24 = append(v24, (fi22 % (((fi15 << (((v11 + v11)) & 63))) | 1)))
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 16; fi26++ {
                v21 = ("hello" + v21)
                var v27 int8 = 100
                _ = fi26
                _ = v27
            }
            _ = v21
        }
        _ = fi15
    }
    var v28 int32 = 0
    v29 := GenS2{F0: 2, F1: 65535, F2: 0}
    for fi30 := 0; fi30 < 15; fi30++ {
        fmt.Println(v8, v5)
        for fi31 := 0; fi31 < 6; fi31++ {
            if true {
                var v32 string = (("foo bar" + ("foo bar" + "hello")) + ("foo bar" + "\n\t"))
                v33 := GenS3{F0: -120, F1: 42, F2: (^(v5))}
                _ = v32
                _ = v33
            }
            v34 := GenS4{F0: (v5 / ((42) | 1)), F1: v5}
            var v35 uint16 = 0
            _ = fi31
            _ = v34
            _ = v35
        }
        fmt.Println(v11, v6)
        _ = fi30
    }
    var v36 uint32 = 4294967295
    if true {
        if (2 < v11) {
            v2 -= (v2 + 1.5)
        } else {
            switch "" {
            case "":
                v5 += (v5 / ((int8(v9)) | 1))
            case "hello":
                fmt.Println(v8)
            case "\n\t":
                var v37 float32 = ((0.1 - (float32(1e10) - v2)) / v2)
                _ = v37
            default:
                if false {
                    if true {
                        fmt.Println(v9, v6, v9)
                    }
                }
            }
            var v38 uint32 = (v36 / (((v36 >> (((^(v11))) & 63))) | 1))
            _ = v38
        }
        var v39 int = v8
        if ((v5 * (-(v5))) <= v5) {
            var v40 int64 = 1
            var v41 bool = true
            _ = v40
            _ = v41
        } else {
            var v42 int = (v8 & (v8 >> (((v11 / ((v11) | 1))) & 63)))
            _ = v42
        }
        _ = v39
    } else {
        for fi43 := 0; fi43 < 19; fi43++ {
            for fi44 := 0; fi44 < 12; fi44++ {
                var v45 string = "\n\t"
                var v46 uint32 = v36
                _ = fi44
                _ = v45
                _ = v46
            }
            v36 = (^(v36))
            _ = fi43
        }
    }
    var v47 float64 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
