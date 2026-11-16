package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    v3 := []uint32{518, 65535, 2147483648, 2}
    var v4 int8 = 127
    var v5 float64 = ((0.0 / -2.25) - 0.1)
    delete(v1, 2147483647)
    var v6 int64 = 0
    _ = len(v2)
    var v7 int32 = 42
    var v8 float64 = (v5 * float64(0.1))
    var v9 int32 = int32(v4)
    delete(v1, 0)
    for fi10 := 0; fi10 < 6; fi10++ {
        if (true && (((^(v9)) <= v7) && true)) {
            fmt.Println(v5)
            for fi11 := 0; fi11 < 13; fi11++ {
                if (1 < 2) {
                    v4 -= v4
                    delete(v1, 721)
                    v7 += v9
                }
                _ = fi11
            }
        }
        switch "foo bar" {
        case "hello":
            for fi12 := 0; fi12 < 3; fi12++ {
                v13 := GenS1{F0: 4294967295, F1: 42, F2: (-(fi12))}
                delete(v1, 2)
                _ = fi12
                _ = v13
            }
        case "foo bar":
            switch fi10 {
            case -812:
                if (128 > 2) {
                    v14 := GenS2{F0: ((v4 >= (v4 >> ((603) & 63))) && (1 < 255))}
                    v15 := GenS1{F0: 65535, F1: 42, F2: 128}
                    if (!true) {
                        if ((float64((float32((0.1 - v8)) + float32(v5))) >= v8) || false) {
                            delete(v1, 677)
                            var v16 uint8 = 0
                            if false {
                                if false {
                                    var v17 uint64 = 2
                                    var v18 float32 = 1e10
                                    if false {
                                        for fi19 := 0; fi19 < 16; fi19++ {
                                            fmt.Println(v16, v8)
                                            fmt.Println(fi19, fi19, v16)
                                            delete(v1, 885)
                                            fmt.Println(fi19, v6, v17)
                                            _ = fi19
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    _ = v17
                                    _ = v18
                                }
                                v3 = append(v3, 2)
                            }
                            _ = v16
                        } else {
                            fmt.Println(v9)
                            if true {
                                var v20 float32 = ((0.0 - 0.1) * (1e10 + ((float32(1.5) - 1e10) - 1e10)))
                                var v21 int16 = -32768
                                _ = v20
                                _ = v21
                            }
                        }
                        var v22 bool = (1 == 4294967295)
                        _ = v22
                    }
                    _ = v14
                    _ = v15
                }
            }
        default:
            v23 := GenS1{F0: 750, F1: -491, F2: fi10}
            _ = v23
        }
        _ = fi10
    }
    switch "hello" {
    case "foo bar":
        var v24 float64 = 1e10
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
