package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint8 = 255
    var v2 int = 32767
    var v3 uint16 = 1
    var v4 uint64 = 2
    v5 := make(map[uint32]bool)
    v6 := v5[65535]
    fmt.Println(v2, v4, v4)
    switch 874 {
    case -508:
        v7 := GenS1{F0: 32767, F1: 957, F2: v3}
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v6, v6)
        _ = fi8
    }
    if true {
        v9 := v5[42]
        _ = v9
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            v1 -= (v1 | (v1 & v1))
            if (v6 || (!(-2.25 >= 1e10))) {
                v11 := []int{int(v4)}
                var v12 uint16 = uint16(v1)
                _ = v11
                _ = v12
            } else {
                v3 -= 255
                var v13 uint = 43
                _ = v13
            }
            _ = fi10
        }
        fmt.Println(v3)
    }
    v14 := []int{v2}
    v15 := v5[65535]
    v3 = 1
    fmt.Println(v2, v3, v3)
    if true {
        fmt.Println(v1, v3, v6)
        switch (v2 / (((^(v2))) | 1)) {
        case -2147483648:
            switch "" {
            case "":
                if true {
                    v16 := GenS1{F0: 32767, F1: 65535, F2: (v3 - 1)}
                    fmt.Println(v3, v2, v3)
                    _ = v16
                } else {
                    v17 := GenS1{F0: -1, F1: uint(v4), F2: (v3 % ((43) | 1))}
                    var v18 int = v2
                    v19 := GenS1{F0: -1, F1: 2, F2: v3}
                    _ = v17
                    _ = v18
                    _ = v19
                }
            case "a":
                if ((!(false || (226 != -9223372036854775808))) || v15) {
                    v1 += 0
                }
            }
        case 42:
            var v20 int = 0
            _ = v20
        }
        var v21 uint8 = v1
        _ = v21
    }
    var v22 bool = v15
    switch 127 {
    case -620:
        v4 = v4
    case 128:
        v3 -= v3
    case 32768:
        fmt.Println(v4)
    default:
        var v23 uint64 = (^(v4))
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        switch "a" {
        case "hello":
            var v25 uint32 = 4294967295
            _ = v25
        case "a":
            switch (v2 + v2) {
            case 489:
                v6 = (v6 || (42 <= -1))
            }
        case "foo bar":
            v26 := GenS2{F0: v4, F1: 42}
            _ = v26
        default:
            for fi27 := 0; fi27 < 17; fi27++ {
                switch "" {
                case "":
                    v28 := v5[718]
                    _ = v28
                case "hello":
                    switch fi24 {
                    case -9223372036854775808:
                        v29 := GenS1{F0: 1000, F1: uint(v3), F2: v3}
                        _ = v29
                    default:
                        v14 = append(v14, fi27)
                        var v30 uint = 1
                        _ = v30
                    }
                default:
                    switch 1 {
                    case 9223372036854775807:
                        v1 -= (v1 / ((v1) | 1))
                    case 210:
                        var v31 string = ""
                        _ = v31
                    case -895:
                        var v32 bool = true
                        _ = v32
                    }
                }
                v33 := GenS3{F0: (^(v1))}
                _ = fi27
                _ = v33
            }
        }
        var v34 int16 = 1000
        _ = fi24
        _ = v34
    }
    v1 = (v1 / (((v1 % ((v1) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
