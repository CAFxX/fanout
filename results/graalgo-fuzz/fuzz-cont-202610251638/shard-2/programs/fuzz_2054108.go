package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64; F1 uint32; F2 int32 }
type GenS3 struct { F0 int8; F1 uint64 }
func genFn0() uint64 {
    v7 := []int{32767, 42, 9223372036854775807, -422}
    var v8 int32 = 0
    v9 := GenS1{F0: 0, F1: -54}
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() uint64 {
    var v10 uint8 = 127
    if false {
        v11 := GenS1{F0: -2147483648, F1: -128}
        _ = v11
    }
    v10 = 42
    _ = v10
    return uint64(v10)
}

func main() {
    var v1 uint32 = 4294967295
    var v2 string = "hello"
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v5 := []uint32{(v1 << ((255) & 63)), (^(v1)), (v1 * 300)}
    v6 := make(map[string]int)
    v1 -= v1
    switch v2 {
    case "foo bar":
        var v12 int64 = 841
        _ = v12
    case "\n\t":
        if ((1.5 - -2.25) > float32(-2.25)) {
            var v13 int16 = 1000
            _ = v13
        }
    case "a":
        if (uint16(v1) < 291) {
            var v14 int8 = 100
            v1 -= v1
            _ = v14
        }
    }
    v2 = "a"
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 int32 = 0
        _ = fi15
        _ = v16
    }
    v6[v2] = -674
    var v17 int = 42
    v18 := GenS2{F0: 1, F1: (v1 % ((v1) | 1)), F2: -75}
    if true {
        for fi19 := 0; fi19 < 17; fi19++ {
            switch (v2 + v2) {
            case "":
                v17 = v17
            case "a":
                switch ("\n\t" + (v2 + ("" + (("\n\t" + "") + v2)))) {
                case "\n\t":
                    var v20 int = fi19
                    _ = v20
                case "a":
                    var v21 uint16 = 255
                    _ = v21
                case "hello":
                    switch -1 {
                    case 42:
                        _ = len(v5)
                    default:
                        var v22 uint = 255
                        _ = v22
                    }
                default:
                    delete(v3, 2)
                }
            case "hello":
                v2 = v2
            }
            switch (v2 + (("foo bar" + "") + (("hello" + "a") + v2))) {
            case "\n\t":
                v17 = fi19
            case "":
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 string = v2
                    _ = fi23
                    _ = v24
                }
            case "a":
                if true {
                    v25 := v6[v2]
                    _ = v25
                }
            }
            v2 = (((v2 + "\n\t") + (v2 + v2)) + v2)
            _ = fi19
        }
        switch -606 {
        case -1:
            v3[v1] = (v1 | (v1 - v1))
        }
    }
    v2 = "hello"
    var v26 int = v17
    var v27 int64 = -744
    fmt.Println(v27, v26)
    v1 = (^(v1))
    v27 -= -9223372036854775808
    switch v2 {
    case "hello":
        _ = len(v5)
    default:
        var v28 int64 = (v27 << ((5) & 63))
        _ = v28
    }
    if len(v5) > 0 {
        v5[0] = (v1 >> ((0) & 63))
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        switch (fi29 / ((v17) | 1)) {
        case 0:
            var v30 int = fi29
            _ = v30
        default:
            if true {
                fmt.Println(v26)
                v1 += (v1 + (v1 - 2147483647))
            }
        }
        v2 = "hello"
        if (649 == 162) {
            v31 := GenS3{F0: -5, F1: 1}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v32 string = v2
            _ = v31
            _ = v32
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
