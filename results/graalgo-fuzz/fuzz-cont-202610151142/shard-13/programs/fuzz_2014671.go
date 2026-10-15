package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int32 }
type GenS2 struct { F0 int }

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    v6 := make(map[string]uint32)
    switch 2147483647 {
    case 42:
        v1 -= 0
    default:
        fmt.Println(v1)
    }
    v1 = v1
    v3 = append(v3, 65535)
    v7 := GenS1{F0: "hello", F1: 0, F2: -2147483648}
    var v8 uint8 = 0
    v9 := GenS1{F0: ("hello" + ("\n\t" + "a")), F1: 618, F2: -2147483648}
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            v1 += 42
            _ = fi11
        }
        _ = fi10
    }
    var v12 string = "foo bar"
    fmt.Println(v12)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 += (v1 * 0)
    v13 := GenS1{F0: "\n\t", F1: 398, F2: 47}
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(fi14, v8, v1)
        v5[((v12 + v12) + (("foo bar" + v12) + ((v12 + "hello") + v12)))] = v12
        _ = fi14
    }
    v8 += 15
    if ((v1 < (-(v1))) || true) {
        var v15 int = 1
        fmt.Println(v1, v1, v12)
        if ((false && true) && true) {
            fmt.Println(v8)
            if true {
                var v16 uint32 = 465
                v17 := GenS2{F0: 857}
                var v18 uint32 = (v16 << ((811) & 63))
                _ = v16
                _ = v17
                _ = v18
            } else {
                v12 = v12
                var v19 uint8 = v8
                for fi20 := 0; fi20 < 1; fi20++ {
                    if ((v15 >> ((65535) & 63)) < fi20) {
                        if (v19 > v8) {
                            v15 -= 636
                            v15 = 32767
                        }
                        switch v12 {
                        case "\n\t":
                            v19 += (v8 / ((v19) | 1))
                        case "":
                            for fi21 := 0; fi21 < 9; fi21++ {
                                var v22 uint = 0
                                var v23 float32 = 0.0
                                if (-75 < 0) {
                                    v23 -= 3.14159
                                    var v24 bool = (100 == 38)
                                    v15 -= 42
                                    var v25 uint16 = 255
                                    _ = v24
                                    _ = v25
                                }
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = fi20
                }
                _ = v19
            }
            fmt.Println(v12)
        } else {
            v26 := GenS2{F0: (v15 << ((46) & 63))}
            var v27 float64 = 0.1
            _ = v26
            _ = v27
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
