package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint8; F2 int }
func genFn0() int64 {
    var v4 uint32 = 420
    for fi5 := 0; fi5 < 12; fi5++ {
        v4 = v4
        if false {
            if false {
                var v6 int = fi5
                _ = v6
            }
            if true {
                v4 = v4
                v7 := GenS1{F0: 65535}
                _ = v7
            } else {
                if false {
                    v4 -= (v4 * v4)
                    switch ("" + ("a" + "a")) {
                    case "foo bar":
                        fmt.Println(v4, v4)
                    case "\n\t":
                        v8 := GenS2{F0: -9223372036854775808}
                        _ = v8
                    }
                } else {
                    for fi9 := 0; fi9 < 8; fi9++ {
                        switch fi9 {
                        case 4294967295:
                            v4 = 42
                        case -864:
                            var v10 uint8 = 88
                            _ = v10
                        default:
                            v4 -= (v4 % ((857) | 1))
                        }
                        switch (fi5 / ((648) | 1)) {
                        case 32768:
                            var v11 float64 = -2.25
                            _ = v11
                        default:
                            var v12 uint16 = 255
                            _ = v12
                        }
                        if (3.14159 >= (-2.25 * -2.25)) {
                            v13 := []int{fi5, fi5, (fi9 % ((fi5) | 1)), fi5}
                            if true {
                                if (797 == 419) {
                                    var v14 int32 = 492
                                    _ = v14
                                } else {
                                    fmt.Println(fi9, fi5, fi9)
                                    if true {
                                        var v15 uint8 = 70
                                        var v16 uint64 = 4294967295
                                        v15 = (v15 & v15)
                                        _ = v15
                                        _ = v16
                                    } else {
                                        v4 = (v4 & (v4 | (v4 + (v4 << ((0) & 63)))))
                                        var v17 uint16 = 42
                                        fmt.Println(fi9, v4, v17)
                                        _ = v17
                                    }
                                }
                                v4 = v4
                            }
                            var v18 int8 = 1
                            _ = v13
                            _ = v18
                        }
                        _ = fi9
                    }
                    var v19 float64 = (0.1 - (1e10 - 1.5))
                    fmt.Println(fi5)
                    _ = v19
                }
                var v20 bool = true
                _ = v20
            }
        } else {
            var v21 uint64 = uint64(v4)
            if (true && (float64((3.14159 + 0.0)) >= (1e10 * -2.25))) {
                v21 -= v21
                var v22 string = "a"
                fmt.Println(v22)
                _ = v22
            }
            _ = v21
        }
        v23 := GenS1{F0: 2}
        _ = fi5
        _ = v23
    }
    switch ((("hello" + "a") + "a") + "hello") {
    case "hello":
        v24 := GenS3{F0: 154, F1: 302, F2: 141}
        _ = v24
    case "a":
        var v25 float32 = float32(1e10)
        _ = v25
    case "":
        var v26 int16 = -32768
        _ = v26
    default:
        v27 := GenS1{F0: 2}
        _ = v27
    }
    v28 := GenS4{F0: 1, F1: 255, F2: 942}
    _ = v28
    _ = v4
    return 42
}
func genFn1(p0 int64, p1 uint64) uint8 {
    fmt.Println(p0)
    p1 += 18446744073709551615
    p0 = genFn0()
    if (255 == 42) {
        var v29 int64 = (p0 % (((p0 & (p0 ^ genFn0()))) | 1))
        switch 32768 {
        case -171:
            if true {
                var v30 uint = 18446744073709551615
                for fi31 := 0; fi31 < 17; fi31++ {
                    var v32 int16 = 642
                    _ = fi31
                    _ = v32
                }
                _ = v30
            } else {
                fmt.Println(p1, p0, p0)
            }
        case 32768:
            var v33 bool = (4294967295 != 0)
            _ = v33
        case 128:
            var v34 int64 = (v29 | p0)
            _ = v34
        default:
            var v35 uint8 = 42
            _ = v35
        }
        p0 += 0
        _ = v29
    }
    _ = p0
    _ = p1
    return 188
}

func main() {
    var v1 int8 = 100
    var v2 int32 = 1
    var v3 int32 = -39
    v3 = (v2 | v2)
    for fi36 := 0; fi36 < 11; fi36++ {
        v37 := []int{9223372036854775807}
        _ = fi36
        _ = v37
    }
    v38 := GenS3{F0: 255, F1: 519, F2: 2}
    v39 := []int{42, 32767, int(v1), 127}
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        if len(v39) > 0 {
            v39[0] = 1
        }
    case "":
        var v40 float64 = float64(0.0)
        _ = v40
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        v1 -= v1
        _ = fi41
    }
    fmt.Println(v2, v3, v3)
    var v42 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v42
}
