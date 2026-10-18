package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 uint8; F2 uint64 }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    v1 += v1
    var v4 uint64 = 0
    v5 := []int{1, 822, 1}
    v6 := GenS1{F0: -845}
    if (true || true) {
        switch ("\n\t" + ("a" + "a")) {
        case "foo bar":
            var v7 uint = v1
            _ = v7
        case "":
            v4 -= (^(v4))
        }
        var v8 float64 = (0.0 * 1e10)
        v8 -= (v8 * 1.5)
        _ = v8
    }
    var v9 uint32 = 642
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int64 = 0
        _ = fi10
        _ = v11
    }
    v4 -= v4
    if (0 >= 65535) {
        if false {
            var v12 int = 2147483647
            _ = v12
        } else {
            switch 128 {
            case 2147483647:
                v1 += v1
            default:
                if false {
                    fmt.Println(v4, v4, v9)
                } else {
                    v1 = (v1 + v1)
                    v1 -= (v1 >> ((v1) & 63))
                }
            }
            var v13 bool = (!(!(false || false)))
            _ = v13
        }
        var v14 int32 = 42
        var v15 float32 = (-2.25 + ((3.14159 * -2.25) * 1.5))
        _ = v14
        _ = v15
    } else {
        _ = len(v5)
    }
    if true {
        for fi16 := 0; fi16 < 19; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := GenS2{F0: 100, F1: -32768}
                _ = fi17
                _ = v18
            }
            v4 -= 818
            fmt.Println(fi16)
            _ = fi16
        }
    } else {
        var v19 int8 = 42
        v20 := GenS2{F0: v19, F1: 1}
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 float64 = (0.1 * -2.25)
            for fi23 := 0; fi23 < 2; fi23++ {
                if true {
                    switch "a" {
                    case "hello":
                        fmt.Println(v9, v4, v4)
                    case "":
                        var v24 int = 42
                        _ = v24
                    case "\n\t":
                        fmt.Println(v19)
                    }
                    var v25 uint = (^(v1))
                    v5 = append(v5, 128)
                    switch (fi23 - 2147483647) {
                    case 4294967295:
                        v2["hello"] = fi21
                    case -1:
                        switch (fi21 % ((fi21) | 1)) {
                        case 127:
                            v26 := GenS2{F0: v19, F1: -32768}
                            _ = v26
                        default:
                            v25 = (v1 + (^(v25)))
                        }
                    default:
                        v27 := GenS3{F0: 42, F1: 1}
                        _ = v27
                    }
                    var v28 float32 = (1.5 + ((1.5 + (0.0 - -2.25)) * 0.1))
                    var v29 uint32 = (v9 >> ((162) & 63))
                    var v30 int = 42
                    var v31 int16 = 1
                    _ = len(v5)
                    _ = v25
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    var v32 int32 = -1
                    v22 -= v22
                    v33 := GenS4{F0: ((v9 + v9) <= v9)}
                    switch ((((("hello" + (("a" + (("a" + "\n\t") + (("" + "hello") + "hello"))) + "\n\t")) + "hello") + ("a" + "hello")) + ("foo bar" + "hello")) + ("" + "\n\t")) {
                    case "hello":
                        var v34 int32 = -2147483648
                        _ = v34
                    case "":
                        var v35 uint64 = (v4 << ((v1) & 63))
                        _ = v35
                    case "a":
                        fmt.Println(v9, v32)
                    }
                    _ = v32
                    _ = v33
                }
                if (v1 > v1) {
                    v36 := GenS4{F0: ((v19 > (v19 >> ((v1) & 63))) || (!false))}
                    v9 += v9
                    fmt.Println(v4)
                    var v37 uint = (v1 + v1)
                    _ = v36
                    _ = v37
                } else {
                    v4 -= v4
                }
                v38 := GenS2{F0: v19, F1: int16(fi21)}
                _ = fi23
                _ = v38
            }
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 uint32 = (v9 | (v9 * 2147483648))
        v9 -= (v9 ^ v40)
        v41 := GenS1{F0: 0}
        _ = fi39
        _ = v40
        _ = v41
    }
    if ((!(1000 < 32767)) || ((!false) || false)) {
        v9 += (v9 << (((v1 >> ((v1) & 63))) & 63))
        delete(v3, 394)
        var v42 float32 = ((-2.25 + float32((0.0 / 1e10))) - 3.14159)
        _ = v42
    } else {
        v43 := GenS5{F0: -32768, F1: 0, F2: v4}
        fmt.Println(v9)
        _ = v43
    }
    v9 -= (v9 / (((v9 << ((v1) & 63))) | 1))
    var v44 bool = (!(1 != -1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v9
}
