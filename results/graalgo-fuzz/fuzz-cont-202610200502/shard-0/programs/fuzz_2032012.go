package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 string; F2 int16 }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 32768
    var v3 int8 = -113
    var v4 int64 = 9223372036854775807
    var v5 bool = (931 > 18446744073709551615)
    v6 := GenS1{F0: 2, F1: -32768}
    v4 -= (^(v4))
    v7 := GenS1{F0: 0, F1: 1}
    v8 := GenS1{F0: 4294967295, F1: 1}
    var v9 int64 = (v4 + (v4 ^ v4))
    v10 := GenS2{F0: 1}
    for fi11 := 0; fi11 < 19; fi11++ {
        v2 -= (fi11 | v2)
        _ = fi11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 float32 = 1e10
        _ = fi12
        _ = v13
    }
    if (!true) {
        v14 := v1[275]
        var v15 uint8 = 128
        _ = v14
        _ = v15
    } else {
        switch v2 {
        case 476:
            switch v2 {
            case 713:
                v16 := []int{v2, -84}
                _ = v16
            case 32767:
                v9 = (^(v4))
            case 32768:
                if false {
                    fmt.Println(v5, v3)
                    switch (v2 | v2) {
                    case 9223372036854775807:
                        v3 = v3
                    default:
                        v17 := v1[42]
                        _ = v17
                    }
                }
            default:
                var v18 uint = 18446744073709551615
                _ = v18
            }
        case 128:
            for fi19 := 0; fi19 < 5; fi19++ {
                delete(v1, 720)
                var v20 string = ("foo bar" + "")
                var v21 bool = (65535 < 64)
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        var v22 uint64 = 834
        var v23 uint = 323
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS2{F0: 18446744073709551615}
            if ((0.1 - (0.1 / (1.5 / 0.1))) > 0.1) {
                switch "hello" {
                case "hello":
                    switch "foo bar" {
                    case "\n\t":
                        delete(v1, 2147483647)
                    case "hello":
                        var v27 uint32 = 4294967295
                        _ = v27
                    case "a":
                        if (255 >= 859) {
                            fmt.Println(v2)
                            fmt.Println(v2, v5)
                            var v28 int16 = -875
                            _ = v28
                        }
                    default:
                        var v29 float64 = 3.14159
                        _ = v29
                    }
                    v9 = 0
                    switch "\n\t" {
                    case "foo bar":
                        switch "foo bar" {
                        case "foo bar":
                            var v30 bool = v5
                            _ = v30
                        case "hello":
                            v31 := []int{fi24, v2}
                            _ = v31
                        }
                    }
                    var v32 string = "\n\t"
                    _ = v32
                case "a":
                    v33 := GenS3{F0: 1, F1: "", F2: 668}
                    _ = v33
                case "foo bar":
                    var v34 uint64 = 42
                    _ = v34
                default:
                    v35 := []int{(fi25 / (((-(fi25))) | 1))}
                    _ = v35
                }
                v36 := GenS4{F0: fi24}
                v2 = fi24
                _ = v36
            } else {
                v9 += (v4 ^ (^(v9)))
                v9 = v9
            }
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    v3 = v3
    v9 += 1
    switch (v2 & 42) {
    case 32767:
        if (65535 > 0) {
            v37 := GenS3{F0: 1, F1: "a", F2: -908}
            for fi38 := 0; fi38 < 19; fi38++ {
                var v39 uint16 = 1
                v2 = (fi38 << ((255) & 63))
                _ = fi38
                _ = v39
            }
            fmt.Println(v9, v9)
            _ = v37
        } else {
            if v5 {
                v4 = v9
                v9 -= v4
            } else {
                fmt.Println(v2)
            }
        }
    case 379:
        for fi40 := 0; fi40 < 6; fi40++ {
            v41 := []int{(v2 ^ v2)}
            v2 = (fi40 - (^(v2)))
            _ = fi40
            _ = v41
        }
    default:
        var v42 int64 = v4
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
