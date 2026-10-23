package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int; F1 int; F2 uint }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 127
    v3 := make(map[uint32]uint32)
    switch "a" {
    case "a":
        for fi4 := 0; fi4 < 5; fi4++ {
            for fi5 := 0; fi5 < 1; fi5++ {
                var v6 int8 = v2
                _ = fi5
                _ = v6
            }
            _ = fi4
        }
    case "hello":
        v2 -= v2
    case "":
        v2 = (v2 << ((42) & 63))
    }
    v7 := []int{4294967295, 1, 42}
    for fi8 := 0; fi8 < 17; fi8++ {
        if (false && false) {
            var v9 float64 = float64(0.0)
            v10 := v3[1]
            v3[v10] = (^(v10))
            _ = v10
            _ = v9
        } else {
            v2 -= v2
            if (116 < 1) {
                var v11 float64 = 1.5
                switch ("" + "a") {
                case "a":
                    v12 := GenS1{F0: 18446744073709551615, F1: (v2 >> ((963) & 63))}
                    _ = v12
                case "foo bar":
                    switch fi8 {
                    case 166:
                        for fi13 := 0; fi13 < 9; fi13++ {
                            switch ("hello" + "\n\t") {
                            case "a":
                                v11 = v11
                            case "hello":
                                fmt.Println(fi13, v2)
                            }
                            switch ("" + "foo bar") {
                            case "hello":
                                v2 = (v2 ^ (v2 * (v2 * (v2 & v2))))
                            case "a":
                                _ = len(v7)
                            default:
                                v14 := GenS2{F0: fi8, F1: fi13, F2: 0}
                                _ = v14
                            }
                            _ = fi13
                        }
                    }
                }
                v15 := GenS2{F0: (fi8 >> ((4294967295) & 63)), F1: (^(fi8)), F2: 4294967295}
                _ = v11
                _ = v15
            }
        }
        _ = fi8
    }
    v2 = v2
    var v16 int64 = 42
    v17 := GenS1{F0: 2, F1: (^(v2))}
    v18 := GenS3{F0: 42}
    for fi19 := 0; fi19 < 5; fi19++ {
        fmt.Println(fi19, fi19)
        v20 := GenS1{F0: 911, F1: (v2 - (^(v2)))}
        _ = fi19
        _ = v20
    }
    v16 -= -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v7
}
