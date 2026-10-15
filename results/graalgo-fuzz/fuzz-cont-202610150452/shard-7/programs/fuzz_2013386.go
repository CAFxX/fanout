package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int64 }
type GenS2 struct { F0 string; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 int) int8 {
    v6 := GenS1{F0: 42, F1: -426, F2: -423}
    var v7 uint64 = 4294967295
    v7 += (v7 ^ 4294967295)
    p0 = (p0 / (((p0 >> ((65535) & 63))) | 1))
    if false {
        var v8 uint32 = 2
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 string = "hello"
            _ = fi9
            _ = v10
        }
        switch "hello" {
        case "hello":
            var v11 int32 = 1
            _ = v11
        case "a":
            v12 := GenS1{F0: uint32(v7), F1: int64(v7), F2: 9223372036854775807}
            _ = v12
        }
        _ = v8
    } else {
        if (2 == 30) {
            v7 = (v7 - v7)
        }
        for fi13 := 0; fi13 < 18; fi13++ {
            switch "a" {
            case "\n\t":
                for fi14 := 0; fi14 < 14; fi14++ {
                    if (1 >= 467) {
                        switch "\n\t" {
                        case "a":
                            p0 = (^(p0))
                            var v15 int = (^(p0))
                            _ = v15
                        case "\n\t":
                            if (2 != 42) {
                                var v16 string = "hello"
                                switch p0 {
                                case 127:
                                    var v17 float32 = 3.14159
                                    _ = v17
                                case 832:
                                    if false {
                                        var v18 uint16 = 65535
                                        _ = v18
                                    }
                                case -304:
                                    v19 := []int{32768, -692, p0, fi14}
                                    _ = v19
                                }
                                _ = v16
                            } else {
                                v20 := []int{(p0 ^ -504), 9223372036854775807}
                                var v21 int16 = -740
                                var v22 float64 = float64((((0.1 / 1e10) / 1e10) - (float32(0.1) + ((3.14159 + 3.14159) + 1e10))))
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            if ((!false) && true) {
                                v23 := []int{42, (p0 | (fi13 * fi13)), p0, fi13}
                                var v24 uint32 = 65535
                                _ = v23
                                _ = v24
                            }
                            v7 += 1
                        }
                    }
                    v25 := []int{950}
                    var v26 int16 = -32768
                    _ = fi14
                    _ = v25
                    _ = v26
                }
            case "a":
                if (1 == 423) {
                    var v27 float32 = 0.1
                    _ = v27
                }
            }
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := []int{fi28, 42, (p0 % ((128) | 1))}
                v30 := GenS2{F0: ("hello" + ("a" + ("\n\t" + "hello"))), F1: 0, F2: 171}
                v31 := GenS3{F0: 0, F1: -32768}
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            v32 := GenS3{F0: 42, F1: 42}
            _ = fi13
            _ = v32
        }
        if (-233 >= -250) {
            v33 := []int{}
            fmt.Println(p0)
            _ = v33
        } else {
            v7 = 0
            fmt.Println(p0)
            v34 := []int{p0, (p0 - p0), p0}
            _ = v34
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 string = (((("a" + "a") + "hello") + "") + "a")
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 2
    var v4 int32 = 147
    var v5 bool = true
    v1 = "\n\t"
    v35 := GenS3{F0: 65535, F1: 0}
    var v36 uint16 = 516
    fmt.Println(v36)
    v4 = v4
    v37 := GenS3{F0: 76, F1: 32767}
    var v38 int = 42
    var v39 uint32 = 46
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
