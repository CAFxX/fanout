package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint64; F2 int8 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 string; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint8, p1 int32, p2 string) int64 {
    switch "hello" {
    case "hello":
        fmt.Println(p2, p0)
    }
    v4 := GenS1{F0: 32767, F1: 616, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -16
}

func main() {
    var v1 int16 = 0
    var v2 int16 = v1
    v3 := make(map[int]int)
    for fi5 := 0; fi5 < 12; fi5++ {
        v2 = 42
        _ = fi5
    }
    var v6 uint64 = 4294967295
    v2 = -550
    var v7 uint32 = 4294967295
    var v8 int16 = -407
    var v9 float64 = 1.5
    for fi10 := 0; fi10 < 8; fi10++ {
        v9 += ((1e10 * v9) - v9)
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint32 = v7
            if (405 < -1) {
                v13 := GenS2{F0: 115, F1: v6, F2: -30}
                var v14 string = "hello"
                if true {
                    v15 := GenS3{F0: -2147483648, F1: 128}
                    v9 += v9
                    v16 := []int{}
                    _ = v15
                    _ = v16
                } else {
                    v17 := GenS4{F0: (("a" + v14) + "a"), F1: v1, F2: 1}
                    switch fi11 {
                    case -1:
                        switch (^(fi11)) {
                        case 4294967295:
                            if false {
                                v18 := []int{(-(fi10))}
                                v8 -= (v2 << ((255) & 63))
                                var v19 uint = 691
                                v8 += (v8 | (v1 - (^(v8))))
                                _ = len(v3)
                                _ = v18
                                _ = v19
                            } else {
                                var v20 int = -9223372036854775808
                                switch 32768 {
                                case -139:
                                    v21 := []int{}
                                    v7 += (v7 >> ((811) & 63))
                                    _ = v21
                                case -2147483648:
                                    var v22 string = v14
                                    _ = v22
                                default:
                                    v23 := GenS5{F0: (v7 | (v12 - (v12 ^ v12))), F1: -791}
                                    var v24 bool = true
                                    fmt.Println(v24)
                                    _ = v23
                                    _ = v24
                                }
                                _ = v20
                            }
                        case 32767:
                            var v25 uint32 = (v7 * 2147483648)
                            _ = v25
                        case 2147483647:
                            v14 = (((v14 + v14) + (v14 + ((v14 + ("" + "")) + "foo bar"))) + "hello")
                        }
                    case 128:
                        var v26 int16 = v2
                        _ = v26
                    }
                    _ = v17
                }
                if true {
                    var v27 int8 = -86
                    _ = v27
                } else {
                    v1 = (v2 + v2)
                    var v28 int = fi11
                    var v29 bool = true
                    _ = v28
                    _ = v29
                }
                _ = len(v3)
                _ = v13
                _ = v14
            } else {
                switch ((("foo bar" + "foo bar") + "foo bar") + "hello") {
                case "foo bar":
                    switch (((("" + "hello") + ("hello" + "\n\t")) + (("hello" + "hello") + "hello")) + "foo bar") {
                    case "":
                        if true {
                            v8 = (v8 * v8)
                        }
                    case "foo bar":
                        fmt.Println(v12, v12, v2)
                    case "hello":
                        if true {
                            v8 += 1
                            var v30 bool = (!(42 > (v6 / (((v6 / (((v6 & v6)) | 1))) | 1))))
                            v31 := GenS4{F0: ("\n\t" + "hello"), F1: 0, F2: 255}
                            v30 = false
                            _ = v30
                            _ = v31
                        }
                    default:
                        var v32 uint = 18446744073709551615
                        _ = v32
                    }
                case "a":
                    var v33 string = "hello"
                    _ = v33
                case "hello":
                    var v34 uint32 = (^(v7))
                    _ = v34
                default:
                    switch fi11 {
                    case 128:
                        var v35 uint = 42
                        _ = v35
                    case -343:
                        v3[(fi10 + fi10)] = 42
                    case 32767:
                        v36 := GenS2{F0: 0, F1: v6, F2: -128}
                        _ = v36
                    default:
                        if false {
                            var v37 float32 = float32(v9)
                            _ = v37
                        } else {
                            fmt.Println(v8)
                        }
                    }
                }
            }
            v2 += v8
            _ = fi11
            _ = v12
        }
        if (false || false) {
            v7 -= 1
            var v38 uint8 = 143
            var v39 uint = 4294967295
            _ = v38
            _ = v39
        } else {
            v1 = (-(v8))
        }
        _ = fi10
    }
    switch 1 {
    case 42:
        var v40 float32 = 1e10
        _ = v40
    }
    var v41 float32 = 1.5
    switch (("hello" + "hello") + "a") {
    case "foo bar":
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
