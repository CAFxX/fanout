package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int16; F2 uint }
type GenS3 struct { F0 int16; F1 int8; F2 bool }
type GenS4 struct { F0 uint8; F1 int8; F2 uint32 }
func genFn0(p0 int) uint8 {
    fmt.Println(p0)
    v6 := []int{(p0 >> ((0) & 63)), p0}
    _ = len(v6)
    var v7 uint8 = 127
    var v8 int32 = -158
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 uint8, p1 uint16, p2 int8) string {
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 538}
        _ = fi9
        _ = v10
    }
    var v11 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (("\n\t" + "foo bar") + "foo bar")
}

func main() {
    var v1 float32 = -2.25
    var v2 int16 = -32768
    var v3 int64 = 0
    var v4 uint64 = 2
    v5 := make(map[string]string)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v2, v2, v1)
    case "hello":
        for fi12 := 0; fi12 < 11; fi12++ {
            if (0 == 2) {
                v3 = -1
                v3 = v3
            } else {
                if (0 != 255) {
                    if (605 < 2) {
                        v4 += 18446744073709551615
                    } else {
                        var v13 uint8 = 85
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 14; fi14++ {
                        delete(v5, "foo bar")
                        if false {
                            v15 := []int{fi12, (^(fi14)), 127, (fi14 << ((65535) & 63))}
                            _ = v15
                        } else {
                            fmt.Println(fi14, fi14)
                        }
                        _ = fi14
                    }
                }
                if ((fi12 / ((-1) | 1)) == (fi12 >> ((601) & 63))) {
                    switch ("" + "a") {
                    case "hello":
                        for fi16 := 0; fi16 < 11; fi16++ {
                            var v17 uint16 = 65535
                            if (true && (2147483647 >= -2147483648)) {
                                delete(v5, "a")
                                var v18 uint8 = 51
                                _ = v18
                            } else {
                                var v19 int16 = -1
                                switch int(v17) {
                                case 4294967295:
                                    var v20 uint = uint(v4)
                                    _ = v20
                                case -9223372036854775808:
                                    if (2 <= v17) {
                                        var v21 int = fi12
                                        var v22 int32 = -2147483648
                                        _ = v21
                                        _ = v22
                                    } else {
                                        switch ("\n\t" + "hello") {
                                        case "":
                                            var v23 uint = 1
                                            _ = v23
                                        case "hello":
                                            fmt.Println(fi16)
                                        }
                                        v4 = v4
                                    }
                                }
                                var v24 string = ("a" + "")
                                v25 := []int{fi16, fi12, fi12}
                                fmt.Println(v24, v24, v3)
                                _ = v19
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS1{F0: (^(v4))}
                            _ = fi16
                            _ = v17
                            _ = v26
                        }
                    case "\n\t":
                        if false {
                            var v27 float32 = 0.0
                            _ = v27
                        }
                    case "foo bar":
                        var v28 uint64 = (v4 >> ((255) & 63))
                        _ = v28
                    }
                    v29 := GenS2{F0: (fi12 / ((4294967295) | 1)), F1: (-(v2)), F2: 2}
                    _ = v29
                }
                var v30 float32 = (v1 / (float32(3.14159) - 1e10))
                _ = v30
            }
            v31 := GenS1{F0: v4}
            var v32 int8 = -128
            _ = fi12
            _ = v31
            _ = v32
        }
    }
    var v33 int64 = (-(v3))
    for fi34 := 0; fi34 < 5; fi34++ {
        if (730 == 108) {
            _ = len(v5)
            v2 = -399
        } else {
            v35 := []int{(fi34 ^ int(v2)), fi34, 32768}
            _ = v35
        }
        v36 := []int{-2147483648}
        _ = fi34
        _ = v36
    }
    fmt.Println(v1, v1, v1)
    v5["foo bar"] = ("\n\t" + ("a" + "\n\t"))
    var v37 uint16 = 0
    fmt.Println(v1)
    var v38 uint16 = 42
    v39 := GenS3{F0: v2, F1: 100, F2: (68 > 42)}
    for fi40 := 0; fi40 < 12; fi40++ {
        v41 := GenS4{F0: 2, F1: -51, F2: 65535}
        _ = fi40
        _ = v41
    }
    delete(v5, "a")
    if true {
        fmt.Println(v2, v38)
    } else {
        _ = len(v5)
        var v42 int = 2147483647
        _ = v42
    }
    v43 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v5
}
