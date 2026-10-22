package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int16; F1 uint16 }

func main() {
    v1 := []string{"hello", "\n\t"}
    v2 := make(map[int]bool)
    var v3 uint64 = 255
    for fi4 := 0; fi4 < 4; fi4++ {
        _ = len(v1)
        var v5 int16 = -522
        switch "a" {
        case "hello":
            _ = len(v2)
        case "foo bar":
            switch "hello" {
            case "":
                switch "" {
                case "hello":
                    v5 -= v5
                case "":
                    var v6 int32 = 0
                    _ = v6
                default:
                    switch fi4 {
                    case -2147483648:
                        var v7 float64 = 0.0
                        _ = v7
                    case 206:
                        if (!(true && true)) {
                            var v8 uint8 = 60
                            var v9 int64 = -1
                            v8 += (v8 & v8)
                            _ = v8
                            _ = v9
                        } else {
                            for fi10 := 0; fi10 < 6; fi10++ {
                                fmt.Println(fi4, v3)
                                var v11 int = (fi10 - 983)
                                v12 := GenS1{F0: 0, F1: (-(v5))}
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                        }
                    }
                }
            }
        case "\n\t":
            delete(v2, fi4)
        default:
            v13 := GenS2{F0: (v5 + v5), F1: 0}
            _ = v13
        }
        _ = fi4
        _ = v5
    }
    if true {
        var v14 uint64 = 697
        _ = v14
    }
    var v15 uint32 = uint32(v3)
    var v16 uint64 = v3
    var v17 uint = 761
    fmt.Println(v3, v15, v3)
    v2[99] = true
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(v16, v15)
        v15 = (v15 - v15)
        fmt.Println(fi18, v17)
        _ = fi18
    }
    var v19 uint64 = v3
    switch 2147483647 {
    case -2147483648:
        switch ("foo bar" + "hello") {
        case "":
            v3 = v16
        case "foo bar":
            if ((v16 & 2) == v19) {
                v20 := GenS1{F0: 100, F1: 1000}
                var v21 float32 = 0.1
                _ = v20
                _ = v21
            }
        default:
            v22 := GenS1{F0: 64, F1: -1}
            _ = v22
        }
    case -1:
        for fi23 := 0; fi23 < 20; fi23++ {
            v3 = v16
            var v24 float64 = float64(0.0)
            v1 = append(v1, "\n\t")
            _ = fi23
            _ = v24
        }
    }
    switch 926 {
    case 1:
        switch 1 {
        case -1:
            v3 += 4294967295
        }
    }
    v1 = append(v1, ("hello" + ""))
    v25 := GenS2{F0: -1, F1: 273}
    if true {
        var v26 float64 = 3.14159
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := v2[fi27]
        _ = fi27
        _ = v28
    }
    var v29 int = 9223372036854775807
    var v30 uint32 = (v15 ^ 2147483647)
    for fi31 := 0; fi31 < 18; fi31++ {
        if ((v19 & v19) <= (v16 * (v16 | (v16 / (((v16 * (v16 - (v3 | v3)))) | 1))))) {
            fmt.Println(v30, v15)
            v19 += 1
        }
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
}
