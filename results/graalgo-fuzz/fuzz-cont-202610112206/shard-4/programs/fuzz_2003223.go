package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
type GenS3 struct { F0 int8; F1 bool }
func genFn0() uint64 {
    v7 := GenS1{F0: 864, F1: 32767, F2: 753}
    var v8 bool = (447 <= 2)
    v9 := GenS1{F0: 0, F1: -1, F2: 777}
    var v10 int64 = 0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint, p1 int, p2 uint64) int64 {
    var v11 uint8 = 1
    v11 = 255
    v11 -= 1
    v12 := []int{4294967295, p1, p1, (p1 >> ((p0) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 9223372036854775807
}

func main() {
    v1 := []int{217, -2147483648, 0, 128}
    var v2 int16 = 660
    var v3 int16 = 296
    v4 := []uint32{304, 607, 505, 316}
    var v5 uint8 = 0
    v6 := make(map[string]uint32)
    fmt.Println(v2, v2)
    switch "a" {
    case "":
        var v13 bool = (2 == 2147483647)
        _ = v13
    }
    fmt.Println(v2, v2, v3)
    v14 := GenS2{F0: 127, F1: -1, F2: 0}
    if true {
        var v15 int16 = v3
        _ = v15
    }
    v2 -= 0
    if (!(152 == 2147483647)) {
        var v16 int32 = -2147483648
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v5 = v5
        _ = fi17
    }
    var v18 string = "a"
    if (226 == 616) {
        _ = len(v6)
        v19 := GenS1{F0: -1, F1: (^(v2)), F2: 163}
        v18 = "hello"
        _ = v19
    } else {
        fmt.Println(v5)
    }
    v20 := GenS2{F0: (^(v5)), F1: 1, F2: genFn1(uint(v5), -1, 693)}
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := GenS2{F0: v5, F1: 1, F2: 0}
        switch v18 {
        case "\n\t":
            if (((((-2.25 / (0.1 - 3.14159)) / 1e10) == (-2.25 - -2.25)) && false) && false) {
                var v23 uint16 = 42
                v24 := GenS1{F0: -1, F1: v2, F2: 2147483648}
                _ = len(v6)
                _ = v23
                _ = v24
            } else {
                v25 := v6[((v18 + v18) + v18)]
                for fi26 := 0; fi26 < 7; fi26++ {
                    _ = len(v1)
                    _ = fi26
                }
                v25 = v25
                _ = v25
            }
        case "a":
            if false {
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 float32 = 1.5
                    var v29 int64 = 0
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                fmt.Println(v18)
                v30 := GenS1{F0: int32(v2), F1: v2, F2: 42}
                _ = v30
            } else {
                if (true && false) {
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        v5 += v5
                    }
                    var v31 int = fi21
                    _ = len(v4)
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 5; fi32++ {
                        v5 = v5
                        _ = fi32
                    }
                    v33 := GenS3{F0: -128, F1: false}
                    _ = v33
                }
                v3 = v2
                fmt.Println(v18, v18, fi21)
            }
        }
        switch v18 {
        case "hello":
            v18 = (v18 + "a")
        case "":
            if (false || false) {
                v18 = v18
                var v34 int64 = 493
                _ = len(v4)
                _ = v34
            } else {
                delete(v6, v18)
                v2 += v2
            }
        }
        _ = fi21
        _ = v22
    }
    v35 := GenS2{F0: (v5 + (^(v5))), F1: -1, F2: 1}
    var v36 uint = 255
    if (339 == 2147483647) {
        for fi37 := 0; fi37 < 7; fi37++ {
            v36 = v36
            switch (v18 + "") {
            case "":
                var v38 int = -2147483648
                _ = v38
            case "\n\t":
                fmt.Println(v5, fi37)
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi37
        }
    } else {
        var v39 int16 = 1
        _ = v39
    }
    switch -899 {
    case 9223372036854775807:
        var v40 uint8 = 2
        _ = v40
    default:
        fmt.Println(v36, v18)
    }
    if true {
        fmt.Println(v36, v5, v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
