package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0() bool {
    v6 := GenS1{F0: 173, F1: 4294967295}
    var v7 float32 = -2.25
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := []int{}
    var v2 int64 = 715
    v3 := make(map[int]uint32)
    var v4 int64 = v2
    var v5 uint = 65535
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    v8 := GenS1{F0: 241, F1: 1}
    v5 = 65535
    v3[-503] = 336
    switch 32767 {
    case 9223372036854775807:
        var v9 uint16 = 255
        _ = v9
    default:
        delete(v3, -831)
    }
    v5 += v5
    if (255 == 255) {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 uint16 = 2
            for fi12 := 0; fi12 < 11; fi12++ {
                switch "foo bar" {
                case "\n\t":
                    v3[(fi12 / ((fi10) | 1))] = 588
                case "":
                    switch -300 {
                    case -9223372036854775808:
                        switch "a" {
                        case "a":
                            v1 = append(v1, int(v11))
                        case "\n\t":
                            var v13 int8 = 1
                            _ = v13
                        }
                    case -771:
                        v5 = (v5 % (((v5 & 65535)) | 1))
                    case 460:
                        v4 = -1
                    }
                case "a":
                    v2 = v2
                }
                _ = fi12
            }
            if false {
                fmt.Println(v5)
                if (-2.25 > 1e10) {
                    v14 := GenS1{F0: 128, F1: 0}
                    for fi15 := 0; fi15 < 5; fi15++ {
                        if true {
                            fmt.Println(v5, v5)
                            v4 = 694
                        } else {
                            v11 -= v11
                            v16 := GenS1{F0: 128, F1: 42}
                            _ = v16
                        }
                        _ = fi15
                    }
                    _ = v14
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                var v17 uint32 = 2
                v18 := GenS1{F0: 202, F1: 4294967295}
                fmt.Println(v5, fi10)
                _ = v17
                _ = v18
            }
            _ = fi10
            _ = v11
        }
    }
    delete(v3, 205)
    for fi19 := 0; fi19 < 20; fi19++ {
        v5 -= v5
        _ = fi19
    }
    v5 = (v5 - (v5 ^ v5))
    for fi20 := 0; fi20 < 8; fi20++ {
        v5 -= v5
        _ = fi20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
