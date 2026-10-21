package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int8 }
func genFn0() uint16 {
    var v5 int64 = -1
    v5 += (v5 - v5)
    var v6 int16 = 1
    v6 = v6
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint = 65535
    var v2 string = (("" + "foo bar") + "a")
    v3 := make(map[int]string)
    var v4 float64 = (float64((((float32((0.0 / (1.5 + 3.14159))) * 1e10) + (1e10 - 0.1)) / 1.0)) / 1e10)
    v2 = (v2 + v2)
    v7 := GenS1{F0: (true && (!true)), F1: v2, F2: 100}
    var v8 int16 = 1
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 bool = (false || ((false || (-1 == v8)) || false))
            var v12 int = (^(fi10))
            var v13 uint16 = 2
            switch v2 {
            case "foo bar":
                v13 = v13
            case "a":
                var v14 int8 = -1
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        v1 -= (^(v1))
        _ = fi9
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS1{F0: false, F1: (v2 + v2), F2: 46}
        if (456 >= 2) {
            var v17 uint64 = 4294967295
            fmt.Println(v1)
            v18 := []int{127}
            _ = v17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v19 := []int{}
    v20 := GenS1{F0: (!(1 == 854)), F1: ("hello" + "foo bar"), F2: -50}
    var v21 uint = v1
    fmt.Println(v1, v1)
    var v22 int64 = 9223372036854775807
    var v23 uint16 = 1
    switch -868 {
    case 128:
        var v24 uint64 = 4294967295
        _ = v24
    case 9223372036854775807:
        for fi25 := 0; fi25 < 8; fi25++ {
            v26 := GenS1{F0: (!(!false)), F1: "\n\t", F2: 0}
            for fi27 := 0; fi27 < 8; fi27++ {
                v1 = (v1 + (v21 * v1))
                v8 += 1
                _ = fi27
            }
            switch fi25 {
            case 32768:
                v28 := v3[(fi25 | (-(fi25)))]
                _ = v28
            case 9223372036854775807:
                if (!false) {
                    if (!true) {
                        _ = len(v3)
                        v23 = (^(v23))
                        v23 += 923
                    }
                    v22 -= (v22 & v22)
                    if true {
                        var v29 int8 = 89
                        if true {
                            var v30 uint8 = 1
                            _ = v30
                        }
                        _ = v29
                    }
                }
            case 2147483647:
                v23 += 322
            }
            _ = fi25
            _ = v26
        }
    case -9223372036854775808:
        v1 = (v1 >> (((v21 - (v1 + v1))) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
