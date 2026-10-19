package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 string; F1 uint; F2 bool }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = (((-2.25 * -2.25) * (0.1 - (3.14159 - 3.14159))) / 1.0)
    var v3 float64 = 3.14159
    v4 := []string{"a"}
    var v5 float32 = v2
    var v6 string = ""
    v6 = "a"
    switch 4294967295 {
    case -1:
        fmt.Println(v6)
    default:
        delete(v1, 2147483648)
    }
    fmt.Println(v6)
    delete(v1, 835)
    if false {
        v6 = v6
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: -1, F1: 141}
            var v9 uint16 = 341
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        var v10 uint16 = 42
        _ = v10
    }
    var v11 bool = (-32768 != 1000)
    switch -793 {
    case 2147483647:
        var v12 bool = v11
        _ = v12
    default:
        var v13 int8 = 1
        _ = v13
    }
    var v14 uint8 = 1
    var v15 int64 = -1
    for fi16 := 0; fi16 < 16; fi16++ {
        if (v14 > v14) {
            switch fi16 {
            case 32767:
                switch -2147483648 {
                case 32768:
                    switch "a" {
                    case "\n\t":
                        v17 := GenS2{F0: 931, F1: (v14 ^ (v14 >> ((0) & 63)))}
                        _ = v17
                    }
                case 141:
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v19 := GenS1{F0: (fi18 - fi16), F1: 1}
                        _ = fi18
                        _ = v19
                    }
                default:
                    switch 32768 {
                    case 526:
                        v14 += v14
                    case -1:
                        v20 := v1[4294967295]
                        _ = v20
                    case 4294967295:
                        v21 := GenS3{F0: v6, F1: 759, F2: false}
                        _ = v21
                    }
                }
            default:
                if false {
                    switch (v6 + ((v6 + "a") + ((v6 + (v6 + v6)) + v6))) {
                    case "a":
                        v4 = append(v4, v6)
                    case "":
                        v15 += v15
                    default:
                        fmt.Println(v14, v11)
                    }
                } else {
                    v3 = v3
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 string = "\n\t"
                        fmt.Println(v11, v11, fi22)
                        _ = fi22
                        _ = v23
                    }
                }
            }
            var v24 int8 = 127
            var v25 int64 = v15
            _ = v24
            _ = v25
        }
        _ = fi16
    }
    if (-17 >= -1) {
        for fi26 := 0; fi26 < 2; fi26++ {
            for fi27 := 0; fi27 < 11; fi27++ {
                v14 = v14
                _ = fi27
            }
            for fi28 := 0; fi28 < 3; fi28++ {
                var v29 int32 = 30
                fmt.Println(v5)
                _ = len(v4)
                _ = fi28
                _ = v29
            }
            _ = fi26
        }
    }
    v30 := GenS2{F0: uint32(v14), F1: (v14 - v14)}
    v5 += (float32((float64(0.0) + v3)) + (v2 / (v2 * v5)))
    var v31 string = "a"
    for fi32 := 0; fi32 < 15; fi32++ {
        _ = len(v1)
        _ = fi32
    }
    if (!(((v2 + v2) < v5) && (128 > 1))) {
        switch -661 {
        case -798:
            for fi33 := 0; fi33 < 14; fi33++ {
                _ = len(v4)
                if true {
                    v5 -= ((v5 - v2) * v2)
                    if false {
                        for fi34 := 0; fi34 < 16; fi34++ {
                            fmt.Println(v5, fi34)
                            _ = fi34
                        }
                    }
                }
                var v35 int8 = -128
                _ = fi33
                _ = v35
            }
        }
        var v36 string = v6
        _ = v36
    } else {
        switch 0 {
        case -1:
            fmt.Println(v15)
        case 2147483647:
            fmt.Println(v14, v31)
        }
        for fi37 := 0; fi37 < 18; fi37++ {
            v38 := GenS1{F0: (fi37 ^ (fi37 >> ((186) & 63))), F1: 560}
            _ = fi37
            _ = v38
        }
        switch -895 {
        case 496:
            v15 = v15
        case 42:
            var v39 int = -2147483648
            _ = v39
        }
    }
    v40 := GenS1{F0: -192, F1: 186}
    v3 += 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
