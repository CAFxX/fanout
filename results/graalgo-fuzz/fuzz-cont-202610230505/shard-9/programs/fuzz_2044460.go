package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0(p0 bool) int16 {
    var v7 uint16 = 255
    switch "hello" {
    case "":
        fmt.Println(v7, v7)
    case "\n\t":
        v8 := GenS1{F0: 1, F1: 703}
        _ = v8
    case "foo bar":
        v9 := GenS1{F0: 2147483648, F1: 521}
        _ = v9
    default:
        if (-96 > 2147483647) {
            v10 := []int{-407}
            switch -784 {
            case 2147483647:
                var v11 uint = 42
                _ = v11
            }
            if false {
                for fi12 := 0; fi12 < 16; fi12++ {
                    p0 = true
                    v13 := GenS1{F0: 2, F1: 2}
                    if (("a" + ("\n\t" + "")) >= "a") {
                        v14 := GenS1{F0: 1, F1: 65535}
                        for fi15 := 0; fi15 < 4; fi15++ {
                            p0 = p0
                            v7 = 32767
                            _ = fi15
                        }
                        if p0 {
                            var v16 int8 = 1
                            var v17 int16 = -400
                            v10 = append(v10, int(v16))
                            _ = v16
                            _ = v17
                        }
                        _ = v14
                    } else {
                        var v18 uint64 = 2
                        p0 = p0
                        _ = v18
                    }
                    _ = fi12
                    _ = v13
                }
                v7 = (v7 + v7)
            }
            _ = v10
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int16 = 32767
                _ = fi19
                _ = v20
            }
            v21 := []int{9223372036854775807, 42, -9223372036854775808}
            _ = v21
        }
    }
    _ = p0
    _ = v7
    return -119
}

func main() {
    var v1 int8 = 127
    v2 := []uint32{2, 4, 283}
    var v3 uint16 = 14
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]bool)
    var v6 int8 = v1
    if false {
        v6 = (^(v1))
    } else {
        for fi22 := 0; fi22 < 3; fi22++ {
            switch (fi22 & fi22) {
            case 4294967295:
                v5[2147483647] = false
            default:
                var v23 int32 = 799
                _ = v23
            }
            if true {
                switch "foo bar" {
                case "foo bar":
                    var v24 string = "a"
                    _ = v24
                case "hello":
                    if false {
                        v6 = v6
                        fmt.Println(v3)
                    }
                }
            }
            _ = fi22
        }
    }
    var v25 int32 = 706
    v26 := GenS1{F0: 554, F1: uint32(v3)}
    v25 = (^(v25))
    v27 := GenS1{F0: 721, F1: uint32(v3)}
    fmt.Println(v6, v1)
    var v28 float64 = (((((0.1 + (0.0 / 0.1)) - 3.14159) * (-2.25 / 1e10)) + (-2.25 * ((0.0 - 0.1) * ((((3.14159 - (3.14159 + (1.5 - (1e10 * 0.0)))) - (1.5 / 1.0)) * -2.25) / float64(1.5))))) * 1.5)
    var v29 int8 = v1
    v3 += 32767
    delete(v5, 2147483647)
    var v30 int32 = int32(v3)
    var v31 uint16 = (v3 / (((v3 ^ (v3 - v3))) | 1))
    v30 = v30
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
