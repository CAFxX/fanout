package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int32; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = ""
    var v3 uint64 = 94
    var v4 string = v2
    fmt.Println(v4, v2, v2)
    var v5 float64 = (0.0 / float64(1.5))
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 79}
        _ = fi6
        _ = v7
    }
    v2 = (v2 + "")
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 float32 = -2.25
        _ = fi8
        _ = v9
    }
    if (false && (42 >= 42)) {
        _ = len(v1)
    }
    if false {
        if false {
            v10 := GenS2{F0: 756, F1: 1, F2: -128}
            for fi11 := 0; fi11 < 3; fi11++ {
                if (!(0 != 1)) {
                    v12 := []int{fi11, (fi11 / ((fi11) | 1)), 156, (fi11 << ((2) & 63))}
                    _ = v12
                } else {
                    v5 -= v5
                }
                fmt.Println(v3, v5, v5)
                _ = fi11
            }
            _ = v10
        }
        v13 := GenS2{F0: 42, F1: 1, F2: 0}
        v4 = v2
        _ = v13
    }
    v14 := GenS3{F0: uint32(v3)}
    for fi15 := 0; fi15 < 5; fi15++ {
        if false {
            v4 = "hello"
            var v16 uint16 = 65535
            for fi17 := 0; fi17 < 4; fi17++ {
                switch int(v16) {
                case 32767:
                    if false {
                        var v18 uint32 = 4294967295
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(fi17, v16)
                        }
                        var v19 uint8 = 112
                        _ = v18
                        _ = v19
                    }
                case 32768:
                    v20 := GenS4{F0: -1, F1: 2}
                    _ = v20
                case 1:
                    var v21 bool = true
                    _ = v21
                }
                _ = fi17
            }
            _ = v16
        }
        if (1 <= (v3 << ((42) & 63))) {
            _ = len(v1)
            for fi22 := 0; fi22 < 6; fi22++ {
                v3 -= v3
                _ = fi22
            }
        }
        _ = fi15
    }
    v23 := []int{0}
    v2 = (("\n\t" + "") + v2)
    if len(v23) > 0 {
        v23[0] = 849
    }
    switch (v2 + "a") {
    case "hello":
        v23 = append(v23, 42)
    case "\n\t":
        var v24 uint32 = 2
        _ = v24
    case "foo bar":
        v2 = ((v2 + v4) + v2)
    }
    v23 = append(v23, 9223372036854775807)
    var v25 bool = true
    switch 711 {
    case 0:
        var v26 string = v4
        _ = v26
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 uint16 = 42
        _ = fi27
        _ = v28
    }
    var v29 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
