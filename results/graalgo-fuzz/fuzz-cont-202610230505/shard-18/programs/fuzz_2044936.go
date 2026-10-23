package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int8 = 127
    var v2 uint16 = 498
    var v3 uint64 = 1
    fmt.Println(v3)
    switch int(v3) {
    case 32767:
        var v4 int64 = -1
        _ = v4
    case 4294967295:
        v1 += 1
    default:
        if (84 < 255) {
            v5 := []int{int(v3)}
            if true {
                v1 -= (v1 & v1)
                if ((^(v2)) > (v2 | (^(v2)))) {
                    switch 128 {
                    case 127:
                        v5 = append(v5, 128)
                    default:
                        switch "a" {
                        case "\n\t":
                            if (1 < (v3 - (^(v3)))) {
                                var v6 bool = false
                                v3 = v3
                                _ = v6
                            }
                        default:
                            if len(v5) > 0 {
                                v5[0] = 42
                            }
                        }
                    }
                    if (712 >= -903) {
                        for fi7 := 0; fi7 < 15; fi7++ {
                            v8 := GenS1{F0: fi7}
                            _ = fi7
                            _ = v8
                        }
                    }
                    v9 := GenS1{F0: -1}
                    _ = v9
                }
            }
            for fi10 := 0; fi10 < 9; fi10++ {
                fmt.Println(v2)
                _ = fi10
            }
            _ = v5
        }
    }
    v3 -= 42
    v1 += 103
    v3 -= 42
    v11 := []int{4294967295, -9223372036854775808, 32768}
    v11 = append(v11, 4294967295)
    switch 1 {
    case 1:
        v12 := GenS1{F0: -1}
        _ = v12
    case 0:
        var v13 string = ""
        _ = v13
    default:
        if len(v11) > 0 {
            v11[0] = 0
        }
    }
    switch int(v2) {
    case -1:
        var v14 string = "foo bar"
        _ = v14
    case 647:
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
