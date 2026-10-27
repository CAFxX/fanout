package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 int64 }

func main() {
    var v1 uint = 2
    var v2 uint16 = 268
    var v3 int32 = -169
    var v4 int8 = 0
    v5 := make(map[int]string)
    var v6 uint64 = 470
    v2 = v2
    v7 := GenS1{F0: -594}
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int16 = -1
        _ = fi8
        _ = v9
    }
    v2 = v2
    v3 = 1
    switch "a" {
    case "":
        var v10 bool = (76 > 219)
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(v6, fi11)
            _ = fi11
        }
    case "foo bar":
        v12 := GenS2{F0: 703}
        _ = v12
    }
    fmt.Println(v2)
    var v13 int16 = 32767
    v14 := []int{}
    v4 = v4
    v13 -= v13
    for fi15 := 0; fi15 < 10; fi15++ {
        if true {
            if ((v6 * 1) < 0) {
                if false {
                    switch "hello" {
                    case "foo bar":
                        var v16 int8 = (v4 | v4)
                        _ = v16
                    case "hello":
                        var v17 float64 = 1.5
                        _ = v17
                    }
                }
                var v18 uint = 148
                _ = v18
            } else {
                v19 := GenS3{F0: (^(v6)), F1: 0, F2: 0}
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                var v20 int64 = 290
                _ = v19
                _ = v20
            }
            if (!((v4 >= v4) || (42 > (v6 % (((v6 / (((v6 * v6)) | 1))) | 1))))) {
                v21 := GenS4{F0: (v13 + (v13 ^ -1)), F1: uint32(v1)}
                if true {
                    var v22 int8 = (v4 - (v4 & v4))
                    v4 = (v4 & (v22 / (((v22 * v4)) | 1)))
                    _ = len(v14)
                    _ = v22
                } else {
                    v6 = v6
                }
                _ = v21
            }
            v23 := v5[-1]
            _ = v23
        } else {
            var v24 float32 = 0.0
            switch "hello" {
            case "foo bar":
                for fi25 := 0; fi25 < 1; fi25++ {
                    var v26 uint32 = 2147483647
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    _ = fi25
                    _ = v26
                }
            }
            v14 = append(v14, (fi15 & (fi15 - fi15)))
            _ = v24
        }
        fmt.Println(v1)
        v3 = v3
        _ = fi15
    }
    if (v4 < 41) {
        v13 = (v13 ^ v13)
        for fi27 := 0; fi27 < 20; fi27++ {
            v5[(^(fi27))] = "a"
            if len(v14) > 0 {
                v14[0] = (-(fi27))
            }
            var v28 int32 = v3
            _ = fi27
            _ = v28
        }
        v29 := GenS5{F0: 128}
        _ = v29
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        if (!(0.1 != 3.14159)) {
            v5[370] = "foo bar"
            var v30 int = 0
            _ = v30
        }
    case "foo bar":
        v31 := GenS6{F0: 1, F1: 1}
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
