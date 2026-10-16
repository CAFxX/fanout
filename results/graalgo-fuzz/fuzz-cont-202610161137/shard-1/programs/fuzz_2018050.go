package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int8 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 bool; F1 uint8; F2 int16 }
type GenS6 struct { F0 bool; F1 int16; F2 bool }
type GenS7 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int16; F1 bool }
func genFn0(p0 int16, p1 uint16, p2 string) uint16 {
    fmt.Println(p0, p0, p0)
    var v5 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    var v1 float64 = 1e10
    var v2 int8 = 1
    var v3 float32 = float32(v1)
    var v4 int8 = 0
    v6 := GenS1{F0: 127, F1: 1, F2: v2}
    fmt.Println(v1, v4, v4)
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS2{F0: 0, F1: (fi7 >> ((65535) & 63))}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 174, F1: -9223372036854775808, F2: (v2 / ((v4) | 1))}
    v3 += (((v3 * ((1.5 * (0.1 - 1.5)) + (0.0 + v3))) * (1e10 - float32(v1))) * v3)
    fmt.Println(v1, v4, v2)
    v1 -= (v1 + v1)
    v1 += v1
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            if true {
                var v12 int64 = 9223372036854775807
                _ = v12
            }
            v2 += (-(v4))
            v13 := GenS1{F0: 2, F1: 42, F2: -32}
            v14 := GenS3{F0: 773, F1: 312}
            var v15 uint16 = 2
            if false {
                switch (fi10 & fi10) {
                case -217:
                    if false {
                        v3 -= float32((0.0 - (v1 + ((v1 + v1) * 3.14159))))
                        v4 -= 42
                    } else {
                        fmt.Println(v3, fi10)
                        switch (fi11 + fi10) {
                        case 514:
                            v16 := []int{fi11, (fi10 / ((32767) | 1)), -1, fi10}
                            _ = v16
                        case -1:
                            var v17 uint16 = genFn0(0, 255, "foo bar")
                            var v18 uint = 18446744073709551615
                            if false {
                                var v19 bool = true
                                _ = v19
                            } else {
                                v3 += (v3 - -2.25)
                            }
                            v1 += 3.14159
                            _ = v17
                            _ = v18
                        }
                        switch (("\n\t" + "foo bar") + "") {
                        case "":
                            v3 += float32(3.14159)
                        case "hello":
                            v20 := GenS2{F0: 0, F1: fi11}
                            _ = v20
                        default:
                            v21 := []int{(fi11 << ((42) & 63)), (fi11 & 743), (-(fi10)), fi11}
                            _ = v21
                        }
                        v22 := GenS3{F0: 1, F1: 2}
                        _ = v22
                    }
                case 128:
                    v2 = 127
                default:
                    v23 := GenS1{F0: 0, F1: 9223372036854775807, F2: v2}
                    _ = v23
                }
                v24 := GenS4{F0: v15, F1: v4}
                v25 := GenS5{F0: (140 == 42), F1: 255, F2: -1}
                v1 = v1
                var v26 uint8 = 246
                _ = v24
                _ = v25
                _ = v26
            } else {
                var v27 int32 = -2147483648
                v28 := GenS6{F0: true, F1: -32768, F2: true}
                v3 += v3
                _ = v27
                _ = v28
            }
            v3 = 0.0
            v3 = 0.1
            _ = fi11
            _ = v13
            _ = v14
            _ = v15
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            fmt.Println(fi10, fi10, v4)
            _ = fi29
        }
        _ = fi10
    }
    fmt.Println(v1)
    for fi30 := 0; fi30 < 7; fi30++ {
        v4 = v4
        var v31 uint16 = genFn0(-32768, genFn0(0, 42, "hello"), "a")
        fmt.Println(v1)
        _ = fi30
        _ = v31
    }
    v32 := GenS7{F0: 42, F1: 0, F2: 433}
    if (((4294967295 > 42) && false) && ("\n\t" != ("" + ("" + (((((("" + (("" + (("" + "\n\t") + "")) + "foo bar")) + "\n\t") + "foo bar") + ("hello" + "\n\t")) + "foo bar") + ("hello" + "\n\t")))))) {
        var v33 uint16 = 32767
        _ = v33
    } else {
        fmt.Println(v2, v3)
        v34 := GenS8{F0: 1000, F1: (-1 == 0)}
        var v35 int = 42
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v6
    _ = v9
}
