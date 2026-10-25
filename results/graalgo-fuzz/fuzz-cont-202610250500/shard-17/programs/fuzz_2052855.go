package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0() string {
    var v5 uint64 = 0
    v5 += (v5 & v5)
    _ = v5
    return "a"
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    v2 += 0
    fmt.Println(v2)
    v2 += v2
    v6 := GenS1{F0: 2}
    v7 := GenS1{F0: 458}
    switch -338 {
    case 42:
        v2 = 2
    case 277:
        var v8 int8 = -119
        _ = v8
    default:
        v2 += 42
    }
    v9 := GenS1{F0: 18446744073709551615}
    var v10 bool = true
    var v11 uint64 = 569
    v12 := GenS1{F0: 523}
    if v10 {
        if (!((!(!(265 == 255))) || (v10 && v10))) {
            if (v10 || (v2 >= 81)) {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v10 = true
                    _ = fi13
                }
            } else {
                fmt.Println(v2, v11)
                v11 = v11
            }
        } else {
            v11 = 1
        }
        v10 = false
        switch -1 {
        case -2147483648:
            switch "foo bar" {
            case "":
                v11 = v11
            case "a":
                var v14 uint8 = v2
                _ = v14
            case "hello":
                var v15 uint = 834
                _ = v15
            default:
                fmt.Println(v11)
            }
        case 4294967295:
            switch 42 {
            case 127:
                if false {
                    switch "" {
                    case "hello":
                        var v16 string = ""
                        _ = v16
                    case "a":
                        var v17 uint16 = 1
                        _ = v17
                    default:
                        if (v10 || (v10 && v10)) {
                            for fi18 := 0; fi18 < 7; fi18++ {
                                var v19 uint16 = 528
                                fmt.Println(fi18, v10, v19)
                                _ = fi18
                                _ = v19
                            }
                            v10 = ((1 == 2147483647) && (!true))
                            for fi20 := 0; fi20 < 12; fi20++ {
                                v21 := GenS2{F0: (998 == int16(fi20)), F1: v11}
                                fmt.Println(v2, fi20, v2)
                                fmt.Println(v11)
                                _ = fi20
                                _ = v21
                            }
                        }
                    }
                } else {
                    var v22 float32 = 1e10
                    var v23 uint = 0
                    _ = v22
                    _ = v23
                }
            case -1:
                var v24 bool = v10
                _ = v24
            }
        default:
            v2 = 127
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
