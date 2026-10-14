package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint; F2 uint16 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint32; F1 int64; F2 int32 }

func main() {
    var v1 uint = 835
    var v2 int64 = 1
    var v3 uint16 = 1
    var v4 int64 = (v2 - 42)
    var v5 int16 = 1000
    var v6 string = ""
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint16 = (v3 << ((v1) & 63))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := []int{(fi9 % ((42) | 1)), (fi9 / ((-502) | 1))}
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int8 = 42
            fmt.Println(v5, fi9)
            var v13 uint32 = 4294967295
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 int8 = -1
        _ = fi9
        _ = v10
        _ = v14
    }
    fmt.Println(v6, v4, v5)
    if false {
        var v15 float32 = float32((3.14159 * (-2.25 * 0.0)))
        _ = v15
    }
    v4 = -9223372036854775808
    v16 := []int{}
    for fi17 := 0; fi17 < 11; fi17++ {
        v3 = v3
        v5 = v5
        _ = fi17
    }
    v6 = "\n\t"
    var v18 int = 128
    v1 += 0
    var v19 uint32 = 0
    _ = len(v16)
    v20 := GenS1{F0: uint8(v1)}
    v21 := GenS1{F0: 140}
    fmt.Println(v4, v2, v1)
    switch v18 {
    case 2147483647:
        switch "" {
        case "\n\t":
            v18 += (-(v18))
        }
    default:
        v4 = (^(v2))
    }
    var v22 string = v6
    switch v6 {
    case "foo bar":
        for fi23 := 0; fi23 < 9; fi23++ {
            v22 = "a"
            _ = fi23
        }
    case "\n\t":
        switch 1 {
        case -704:
            switch (v18 | (-(v18))) {
            case 128:
                v24 := GenS2{F0: v18, F1: v1, F2: (v3 | (v3 + v3))}
                _ = v24
            case 861:
                var v25 float32 = (0.1 - float32(float64((1.5 * 0.1))))
                _ = v25
            }
        }
    case "":
        for fi26 := 0; fi26 < 7; fi26++ {
            for fi27 := 0; fi27 < 11; fi27++ {
                v28 := GenS3{F0: true, F1: 42}
                if false {
                    v1 = 1
                    if len(v16) > 0 {
                        v16[0] = -9223372036854775808
                    }
                    if (v1 == 376) {
                        fmt.Println(fi26)
                        _ = len(v16)
                    }
                }
                v16 = append(v16, fi27)
                _ = fi27
                _ = v28
            }
            v29 := GenS1{F0: 1}
            if true {
                if true {
                    if (true && true) {
                        var v30 uint32 = (v19 | v19)
                        _ = v30
                    } else {
                        for fi31 := 0; fi31 < 3; fi31++ {
                            var v32 int = 32767
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            var v33 uint16 = 255
                            v34 := GenS4{F0: (v19 ^ (v19 - v19)), F1: v2, F2: -1}
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    }
                    var v35 uint16 = (^(v3))
                    _ = v35
                }
                v2 = v2
            }
            _ = fi26
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
