package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int8; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 string }
type GenS4 struct { F0 string; F1 int8; F2 bool }
type GenS5 struct { F0 bool; F1 bool }

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    var v3 int = 514
    var v4 uint32 = 65535
    v3 = int(v1)
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: 873, F1: 255}
        v7 := GenS1{F0: 18446744073709551615, F1: 1}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 string = ("foo bar" + "\n\t")
    v9 := []int{(-(v3))}
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := GenS2{F0: 1, F1: 0, F2: 657}
            _ = fi11
            _ = v12
        }
        var v13 float64 = (3.14159 * float64(-2.25))
        _ = fi10
        _ = v13
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS2{F0: 1, F1: 91, F2: 2147483647}
        var v16 uint8 = v1
        v17 := GenS2{F0: 970, F1: 100, F2: -763}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    v1 -= (v1 & v1)
    var v18 float64 = 1e10
    v4 = v4
    switch (v8 + "hello") {
    case "\n\t":
        v19 := v2[(v3 & (v3 + (^(v3))))]
        _ = v19
    default:
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS3{F0: 103, F1: 1, F2: v8}
            _ = fi20
            _ = v21
        }
    }
    switch v8 {
    case "":
        v4 -= (v4 * 1)
    case "hello":
        var v22 int64 = -324
        _ = v22
    case "\n\t":
        var v23 uint16 = 997
        _ = v23
    default:
        v3 = -2147483648
    }
    if (-9223372036854775808 > 9223372036854775807) {
        if true {
            v4 = (v4 * v4)
        } else {
            for fi24 := 0; fi24 < 9; fi24++ {
                v3 = 9223372036854775807
                _ = fi24
            }
            for fi25 := 0; fi25 < 11; fi25++ {
                switch (v3 | v3) {
                case 1:
                    var v26 float64 = v18
                    _ = v26
                }
                v27 := GenS4{F0: v8, F1: -128, F2: (2147483647 < 232)}
                v18 += float64(((3.14159 - 0.1) - ((-2.25 * (1.5 + ((float32((0.0 - float64(1.5))) / float32(((1e10 - -2.25) * v18))) + (3.14159 * -2.25)))) / 0.1)))
                _ = fi25
                _ = v27
            }
        }
        switch (v8 + v8) {
        case "a":
            for fi28 := 0; fi28 < 18; fi28++ {
                for fi29 := 0; fi29 < 5; fi29++ {
                    v30 := GenS3{F0: 849, F1: 2, F2: ("\n\t" + v8)}
                    var v31 uint8 = 0
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                for fi32 := 0; fi32 < 12; fi32++ {
                    var v33 int16 = -1
                    switch fi28 {
                    case 339:
                        fmt.Println(v33, v33, v1)
                    default:
                        var v34 int64 = -9223372036854775808
                        _ = v34
                    }
                    v35 := GenS1{F0: 4294967295, F1: 722}
                    _ = fi32
                    _ = v33
                    _ = v35
                }
                _ = fi28
            }
        case "foo bar":
            switch v3 {
            case 5:
                var v36 int64 = 9223372036854775807
                _ = v36
            case -9223372036854775808:
                var v37 uint32 = v4
                _ = v37
            default:
                v38 := GenS5{F0: (0 >= 795), F1: (!true)}
                _ = v38
            }
        default:
            v39 := GenS1{F0: 85, F1: 255}
            _ = v39
        }
    }
    v3 += (v3 - v3)
    var v40 string = "a"
    v1 -= (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v8
    _ = v9
}
