package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint32 = 993
    var v2 float32 = ((float32(3.14159) * (3.14159 - 0.1)) / -2.25)
    var v3 uint32 = 42
    switch "" {
    case "a":
        for fi4 := 0; fi4 < 18; fi4++ {
            if (2147483647 < (v3 / (((^(v1))) | 1))) {
                fmt.Println(fi4)
                for fi5 := 0; fi5 < 8; fi5++ {
                    var v6 int32 = 1
                    var v7 uint = 18446744073709551615
                    v8 := []int{0}
                    _ = fi5
                    _ = v6
                    _ = v7
                    _ = v8
                }
            }
            _ = fi4
        }
    default:
        v3 = v1
    }
    v2 = 1e10
    v9 := GenS1{F0: int64(v3), F1: 1}
    var v10 int = 1
    v1 += v3
    var v11 uint64 = 42
    var v12 float32 = (0.0 / v2)
    var v13 bool = true
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS2{F0: -128}
        v12 = v12
        v16 := GenS3{F0: 0, F1: 32767}
        _ = fi14
        _ = v15
        _ = v16
    }
    v13 = (!v13)
    if (!true) {
        v17 := []int{v10, v10}
        _ = v17
    }
    fmt.Println(v10, v13)
    for fi18 := 0; fi18 < 4; fi18++ {
        v3 = 737
        switch "a" {
        case "foo bar":
            v1 -= 657
        }
        switch ("hello" + "a") {
        case "foo bar":
            if (0 == 402) {
                var v19 int8 = 110
                var v20 string = ("a" + ((("a" + ("hello" + "")) + ("" + "foo bar")) + "\n\t"))
                _ = v19
                _ = v20
            } else {
                if v13 {
                    var v21 int16 = -1
                    switch -1 {
                    case 42:
                        switch v10 {
                        case -1:
                            v22 := GenS2{F0: 127}
                            _ = v22
                        case 496:
                            for fi23 := 0; fi23 < 12; fi23++ {
                                var v24 int8 = 127
                                _ = fi23
                                _ = v24
                            }
                        default:
                            v2 += v12
                        }
                    case 4294967295:
                        v25 := []int{}
                        _ = v25
                    default:
                        var v26 uint64 = v11
                        _ = v26
                    }
                    fmt.Println(v12, v12, v11)
                    _ = v21
                } else {
                    v12 += (float32((3.14159 - -2.25)) - float32(-2.25))
                }
                var v27 uint64 = (v11 * v11)
                var v28 int32 = -424
                _ = v27
                _ = v28
            }
        case "":
            var v29 int16 = 0
            _ = v29
        case "a":
            for fi30 := 0; fi30 < 15; fi30++ {
                var v31 uint64 = 112
                v32 := []int{fi30}
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        _ = fi18
    }
    v13 = v13
    var v33 bool = false
    for fi34 := 0; fi34 < 17; fi34++ {
        v35 := []int{-2147483648, (fi34 << ((65535) & 63)), -2147483648}
        v36 := GenS3{F0: 0, F1: -1}
        v37 := GenS3{F0: 0, F1: 32767}
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    v38 := GenS4{F0: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v9
}
