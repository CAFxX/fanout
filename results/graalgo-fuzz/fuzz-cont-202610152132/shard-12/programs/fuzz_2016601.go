package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint8, p1 int32) uint {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float32 = -2.25
        v8 := []int{(fi6 << ((37) & 63)), fi6}
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 += (p1 - p1)
    v9 := []int{-280}
    _ = p0
    _ = p1
    _ = v9
    return 65535
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int16 = int16(v1)
    v3 := make(map[uint32]bool)
    var v4 uint16 = 389
    v5 := []string{("" + (("foo bar" + "\n\t") + ("a" + "\n\t")))}
    v4 += (v4 | 65535)
    v4 = 32767
    var v10 bool = false
    if v10 {
        if (!(!(float32(1.5) < 0.1))) {
            for fi11 := 0; fi11 < 16; fi11++ {
                v4 = (v4 % ((42) | 1))
                var v12 uint = 987
                for fi13 := 0; fi13 < 5; fi13++ {
                    if v10 {
                        v14 := GenS1{F0: (v2 >> ((v12) & 63)), F1: (v12 >> (((v12 ^ 1)) & 63))}
                        v12 += v12
                        v10 = true
                        _ = v14
                    }
                    _ = fi13
                }
                _ = fi11
                _ = v12
            }
            if (42 < 127) {
                v15 := GenS1{F0: v2, F1: 65535}
                v10 = true
                if (!v10) {
                    switch (((("hello" + ("a" + "\n\t")) + "\n\t") + "a") + (("foo bar" + ("hello" + "\n\t")) + "")) {
                    case "\n\t":
                        if true {
                            v4 += (v4 % ((v4) | 1))
                            v16 := GenS2{F0: 0}
                            _ = v16
                        } else {
                            var v17 int16 = (-(v2))
                            _ = v17
                        }
                    case "hello":
                        var v18 float32 = (float32(float64(0.0)) - (-2.25 * (1e10 - -2.25)))
                        _ = v18
                    default:
                        switch (("foo bar" + "foo bar") + (("a" + "") + "foo bar")) {
                        case "hello":
                            v4 += (v4 & 2)
                        case "":
                            for fi19 := 0; fi19 < 5; fi19++ {
                                v4 = (v4 >> ((18446744073709551615) & 63))
                                var v20 uint = 1
                                var v21 int32 = -1
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                        case "\n\t":
                            v22 := GenS3{F0: 0, F1: 127}
                            _ = v22
                        default:
                            if true {
                                var v23 uint32 = 0
                                _ = v23
                            }
                        }
                    }
                    fmt.Println(v2, v2, v2)
                } else {
                    v10 = (true || ((v4 * 925) < 32767))
                    if false {
                        for fi24 := 0; fi24 < 8; fi24++ {
                            var v25 uint32 = 2
                            _ = fi24
                            _ = v25
                        }
                    } else {
                        fmt.Println(v2, v4)
                    }
                    v4 = v4
                }
                _ = v15
            }
            fmt.Println(v4)
        }
    }
    v26 := GenS4{F0: (42 < 18446744073709551615), F1: 255}
    _ = len(v5)
    switch (v1 - v1) {
    case -1:
        var v27 uint32 = 4294967295
        _ = v27
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        v2 -= v2
        fmt.Println(fi28)
        fmt.Println(v1)
        _ = fi28
    }
    v1 -= 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
