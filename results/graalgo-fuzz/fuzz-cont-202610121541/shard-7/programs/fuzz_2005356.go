package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0() int8 {
    var v7 int64 = -465
    v8 := []int{128}
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int32) int32 {
    p0 += (p0 | p0)
    v9 := GenS1{F0: 42, F1: 255}
    var v10 uint64 = 1
    fmt.Println(v10)
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{fi11, -324}
        p0 -= 0
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v9
    return p0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    v4 := []uint32{2, 552, 782}
    var v5 uint = 4294967295
    var v6 uint32 = 65535
    fmt.Println(v6)
    for fi13 := 0; fi13 < 19; fi13++ {
        if false {
            if (259 != 42) {
                if true {
                    if ((42 < 0) || (0 >= 127)) {
                        v6 -= (^(v6))
                    }
                    v6 += v6
                } else {
                    if (v6 == 1) {
                        switch (fi13 - (^(fi13))) {
                        case -1:
                            if (481 == 42) {
                                var v14 float32 = 0.0
                                delete(v2, "a")
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    v14 = v14
                                    _ = fi15
                                }
                                _ = v14
                            }
                        default:
                            switch (^(fi13)) {
                            case 2147483647:
                                v16 := GenS1{F0: (v6 << (((v5 ^ (v5 ^ (v5 | 1)))) & 63)), F1: 0}
                                _ = v16
                            }
                        }
                    } else {
                        switch "foo bar" {
                        case "a":
                            v17 := GenS1{F0: (v6 - 0), F1: 1}
                            _ = v17
                        default:
                            var v18 uint16 = 1
                            _ = v18
                        }
                        var v19 uint32 = uint32(v5)
                        _ = v19
                    }
                    var v20 uint = uint(v6)
                    _ = v20
                }
                if true {
                    switch (fi13 + 760) {
                    case 42:
                        v6 = (^(v6))
                    case -9223372036854775808:
                        v6 = (^(v6))
                    case 417:
                        var v21 uint64 = 42
                        _ = v21
                    }
                    switch fi13 {
                    case 4294967295:
                        v6 -= (^(v6))
                    case -292:
                        var v22 uint64 = 0
                        _ = v22
                    }
                    if true {
                        for fi23 := 0; fi23 < 9; fi23++ {
                            var v24 uint8 = 0
                            _ = fi23
                            _ = v24
                        }
                        var v25 int16 = 0
                        _ = v25
                    }
                }
            } else {
                v26 := GenS2{F0: (v6 | 2147483648), F1: uint16(v6)}
                v6 -= v6
                _ = v26
            }
            var v27 int64 = 214
            _ = v27
        }
        _ = fi13
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        switch "" {
        case "\n\t":
            v5 = (v5 << (((v5 % (((v5 >> ((42) & 63))) | 1))) & 63))
        }
        v6 += v6
        _ = fi28
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        switch (fi29 >> (((^(v5))) & 63)) {
        case 128:
            switch "" {
            case "hello":
                fmt.Println(fi29, v5, v5)
            case "":
                v6 = (v6 % (((v6 % (((^(v6))) | 1))) | 1))
            case "foo bar":
                v5 = 18446744073709551615
            }
        }
        if ((-2147483648 < -1) && (false || true)) {
            var v30 int16 = 0
            _ = v30
        }
        var v31 int = fi29
        _ = fi29
        _ = v31
    }
    v32 := GenS1{F0: v6, F1: 0}
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 uint8 = 174
        v5 = v5
        _ = fi33
        _ = v34
    }
    if true {
        var v35 bool = false
        _ = v35
    }
    _ = len(v4)
    v5 = 891
    var v36 float64 = 0.0
    for fi37 := 0; fi37 < 18; fi37++ {
        v6 = v6
        v36 -= (1.5 / (v36 + v36))
        var v38 int64 = 42
        _ = fi37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
