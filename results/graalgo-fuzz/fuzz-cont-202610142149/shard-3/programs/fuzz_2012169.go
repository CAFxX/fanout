package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 int32 = -362
    var v2 string = "\n\t"
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    for fi5 := 0; fi5 < 13; fi5++ {
        switch -2147483648 {
        case 32767:
            v6 := GenS1{F0: (v1 + -234), F1: int64(fi5)}
            _ = v6
        }
        switch "\n\t" {
        case "foo bar":
            var v7 uint8 = 42
            _ = v7
        case "a":
            fmt.Println(v2, fi5, v1)
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int32 = (v1 - (v1 + v1))
        v10 := GenS2{F0: 4294967295, F1: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 float64 = 0.1
    fmt.Println(v1, v2, v11)
    if false {
        var v12 int64 = 358
        var v13 uint64 = 18446744073709551615
        v11 = v11
        _ = v12
        _ = v13
    } else {
        if (true || (1 > 1)) {
            v14 := GenS1{F0: (v1 ^ v1), F1: 1}
            for fi15 := 0; fi15 < 18; fi15++ {
                for fi16 := 0; fi16 < 16; fi16++ {
                    if (83 < -101) {
                        v17 := GenS3{F0: 65535}
                        _ = v17
                    } else {
                        v18 := GenS1{F0: (v1 / (((v1 % ((v1) | 1))) | 1)), F1: -771}
                        var v19 uint = 1
                        var v20 uint64 = 4294967295
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    var v21 uint8 = 2
                    var v22 uint8 = 128
                    _ = fi16
                    _ = v21
                    _ = v22
                }
                v23 := GenS4{F0: (!true), F1: (fi15 % ((fi15) | 1))}
                v24 := GenS5{F0: 96}
                _ = fi15
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 float64 = float64(float32(v11))
                if (true || false) {
                    var v27 uint16 = 32767
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
            _ = v14
        }
    }
    switch v2 {
    case "a":
        for fi28 := 0; fi28 < 5; fi28++ {
            if (v11 < float64(1e10)) {
                v29 := []int{(fi28 * 4294967295), 2147483647, (^(fi28))}
                _ = v29
            } else {
                v1 -= -1
                var v30 int32 = (v1 >> ((76) & 63))
                v2 = ((v2 + (("foo bar" + v2) + "\n\t")) + "")
                _ = v30
            }
            var v31 uint8 = 127
            _ = fi28
            _ = v31
        }
    case "foo bar":
        var v32 float64 = (3.14159 + v11)
        _ = v32
    }
    _ = len(v4)
    v33 := []int{}
    fmt.Println(v1, v2, v11)
    var v34 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
