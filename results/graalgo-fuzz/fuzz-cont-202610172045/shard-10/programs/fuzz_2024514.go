package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int16 }

func main() {
    v1 := []uint32{4294967295}
    v2 := []uint32{342, 65535, 0, 2}
    v3 := make(map[uint32]string)
    v4 := []string{""}
    var v5 int = 389
    var v6 uint64 = 42
    fmt.Println(v5, v6)
    for fi7 := 0; fi7 < 3; fi7++ {
        v5 = (v5 ^ v5)
        var v8 bool = true
        var v9 string = "\n\t"
        _ = fi7
        _ = v8
        _ = v9
    }
    if (0 < 42) {
        var v10 int = v5
        switch ("\n\t" + ("hello" + "a")) {
        case "a":
            if (true || (!true)) {
                v3[1] = "a"
                v10 += -2147483648
            } else {
                for fi11 := 0; fi11 < 17; fi11++ {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        if false {
                            if true {
                                v6 = v6
                                v13 := GenS1{F0: 42, F1: -9223372036854775808}
                                var v14 uint16 = 32767
                                var v15 uint = 65535
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            v16 := GenS1{F0: 1, F1: 370}
                            var v17 int32 = 1
                            _ = v16
                            _ = v17
                        }
                        v18 := GenS1{F0: 0, F1: 1}
                        v6 = v6
                        _ = fi12
                        _ = v18
                    }
                    _ = fi11
                }
            }
        case "\n\t":
            fmt.Println(v10, v5, v6)
        }
        _ = v10
    } else {
        v19 := GenS2{F0: 65535, F1: 4294967295}
        switch v5 {
        case 32767:
            var v20 int = (v5 / ((v5) | 1))
            _ = v20
        case 128:
            var v21 int8 = 100
            _ = v21
        }
        _ = v19
    }
    switch v5 {
    case -444:
        for fi22 := 0; fi22 < 17; fi22++ {
            v3[65535] = ("a" + "")
            _ = fi22
        }
    default:
        v23 := GenS3{F0: 782}
        _ = v23
    }
    fmt.Println(v6, v5)
    for fi24 := 0; fi24 < 13; fi24++ {
        v5 += v5
        _ = fi24
    }
    var v25 int64 = -996
    v26 := GenS4{F0: (v25 | (v25 >> ((2) & 63))), F1: true}
    v27 := GenS5{F0: 1}
    var v28 bool = false
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 uint32 = 4294967295
        var v31 uint8 = 0
        var v32 uint16 = 255
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
