package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }
func genFn0() string {
    var v6 uint8 = 99
    v7 := []int{469, 1, 2147483647, 127}
    fmt.Println(v6)
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1(p0 uint, p1 int16) int32 {
    v8 := GenS1{F0: 0, F1: (p1 >> ((255) & 63)), F2: true}
    p1 = 131
    fmt.Println(p0)
    var v9 int16 = p1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint64 = 171
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    var v5 uint32 = 2
    v10 := []int{760, -1, 1}
    fmt.Println(v1, v5)
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v11 string = "hello"
    var v12 int64 = -895
    switch v11 {
    case "hello":
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(v1)
            for fi14 := 0; fi14 < 3; fi14++ {
                if (false && (990 == 521)) {
                    switch ((v11 + v11) + v11) {
                    case "foo bar":
                        v5 = 2
                    case "hello":
                        switch "foo bar" {
                        case "\n\t":
                            switch ("a" + "a") {
                            case "":
                                v15 := GenS1{F0: 0, F1: -780, F2: false}
                                var v16 uint16 = 32767
                                switch fi13 {
                                case 2147483647:
                                    fmt.Println(v1, v16)
                                case 4294967295:
                                    v17 := GenS2{F0: 0, F1: uint8(v5), F2: (v16 % ((v16) | 1))}
                                    _ = v17
                                case -600:
                                    if true {
                                        v12 = 0
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        var v18 int = fi14
                                        var v19 uint8 = 128
                                        fmt.Println(fi13)
                                        _ = v18
                                        _ = v19
                                    }
                                }
                                _ = v15
                                _ = v16
                            }
                        case "hello":
                            v11 = "foo bar"
                        }
                    }
                }
                v20 := v3[""]
                var v21 uint16 = 1
                _ = fi14
                _ = v20
                _ = v21
            }
            _ = fi13
        }
    case "\n\t":
        v5 += 4294967295
    default:
        v5 = v5
    }
    var v22 string = v11
    v23 := GenS1{F0: -128, F1: -32768, F2: true}
    v5 = v5
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 int64 = (v12 % (((v12 * 9223372036854775807)) | 1))
        v5 += (v5 / ((v5) | 1))
        _ = fi24
        _ = v25
    }
    v11 = v11
    v11 = ("foo bar" + "foo bar")
    switch 1 {
    case 9223372036854775807:
        v26 := GenS2{F0: 206, F1: 255, F2: 32767}
        _ = v26
    case -782:
        var v27 uint32 = 0
        _ = v27
    case 0:
        v11 = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
