package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint32; F1 string; F2 int32 }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int32; F1 uint64 }
func genFn0(p0 string, p1 int8) int32 {
    p1 = (p1 ^ (p1 | (p1 << ((708) & 63))))
    switch 32768 {
    case 128:
        var v7 uint = 4294967295
        _ = v7
    case -172:
        v8 := GenS1{F0: 54}
        _ = v8
    default:
        switch 32768 {
        case 458:
            if true {
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 uint8 = 2
                    v11 := []int{}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        case -600:
            v12 := GenS1{F0: 42}
            _ = v12
        }
    }
    var v13 int64 = -1
    _ = p0
    _ = p1
    _ = v13
    return 42
}
func genFn1(p0 int32, p1 int32, p2 uint16) int8 {
    v14 := GenS2{F0: -1}
    for fi15 := 0; fi15 < 19; fi15++ {
        if (0 != 40) {
            var v16 int32 = p1
            _ = v16
        }
        _ = fi15
    }
    var v17 int = 582
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v17
    return 42
}

func main() {
    var v1 int16 = 1
    var v2 int32 = 91
    var v3 uint = 2
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    v6 := make(map[int]uint32)
    var v18 uint = (^(v3))
    if ((-(v1)) >= (v1 | (^(v1)))) {
        v19 := GenS3{F0: genFn1((-(v2)), (^(v2)), 2), F1: 316, F2: 1}
        if false {
            v1 = v1
            v20 := GenS2{F0: (^(v1))}
            v21 := GenS4{F0: 2, F1: "hello", F2: v2}
            _ = v20
            _ = v21
        } else {
            v6[128] = 2147483648
        }
        v22 := GenS3{F0: 39, F1: 65535, F2: -642}
        _ = v19
        _ = v22
    } else {
        if true {
            v3 = 255
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 uint32 = 4294967295
                v1 -= (v1 + v1)
                _ = fi23
                _ = v24
            }
            switch "" {
            case "a":
                var v25 int32 = (v2 << ((v18) & 63))
                _ = v25
            case "":
                fmt.Println(v2, v3)
            default:
                _ = len(v4)
            }
        } else {
            v2 = genFn0("foo bar", -33)
        }
        v26 := GenS5{F0: 4294967295, F1: 1, F2: 0}
        _ = v26
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 float32 = (((3.14159 - (1e10 + 1.5)) + 0.1) / -2.25)
        _ = fi27
        _ = v28
    }
    switch -1 {
    case -223:
        if true {
            v1 = v1
            v3 -= 128
            fmt.Println(v1)
        }
    case 127:
        if false {
            var v29 int32 = -2147483648
            _ = v29
        }
    default:
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 int16 = v1
            _ = fi30
            _ = v31
        }
    }
    var v32 int16 = v1
    var v33 string = "\n\t"
    switch (v33 + (v33 + v33)) {
    case "":
        var v34 int = -9223372036854775808
        _ = v34
    case "foo bar":
        v35 := []int{}
        _ = v35
    case "hello":
        for fi36 := 0; fi36 < 12; fi36++ {
            switch v33 {
            case "":
                for fi37 := 0; fi37 < 4; fi37++ {
                    if false {
                        v32 += (v1 ^ (v32 - v32))
                        v2 = v2
                    }
                    _ = fi37
                }
            case "\n\t":
                for fi38 := 0; fi38 < 16; fi38++ {
                    v39 := GenS6{F0: -1, F1: 895}
                    _ = fi38
                    _ = v39
                }
            case "foo bar":
                v40 := GenS1{F0: 44}
                _ = v40
            }
            _ = fi36
        }
    }
    v41 := []int{-983}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
