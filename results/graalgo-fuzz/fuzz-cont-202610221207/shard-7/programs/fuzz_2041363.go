package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int64 }

func main() {
    var v1 int = 816
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 0
    var v4 int64 = -438
    var v5 uint16 = 2
    var v6 uint32 = 2147483647
    if ((v3 & (v3 * v3)) <= v3) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v4 += v4
            v8 := v2[(v6 * (v6 - v6))]
            _ = fi7
            _ = v8
        }
        if true {
            var v9 string = "hello"
            _ = v9
        }
        var v10 string = ("hello" + "a")
        _ = v10
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            v5 -= v5
            for fi12 := 0; fi12 < 19; fi12++ {
                if true {
                    v13 := GenS1{F0: 2}
                    var v14 float32 = 0.0
                    fmt.Println(v6, v1)
                    v3 += v3
                    _ = v13
                    _ = v14
                } else {
                    var v15 float32 = (0.0 * 3.14159)
                    _ = v15
                }
                v1 -= fi11
                v2[v6] = 0
                if true {
                    var v16 int8 = 1
                    switch "" {
                    case "foo bar":
                        var v17 uint64 = v3
                        _ = v17
                    case "hello":
                        v5 -= uint16(v3)
                    case "\n\t":
                        fmt.Println(fi11, v16)
                    default:
                        if (false || true) {
                            v1 -= fi12
                        } else {
                            var v18 float64 = 1.5
                            _ = v18
                        }
                    }
                    _ = v16
                } else {
                    switch ("\n\t" + ("" + "hello")) {
                    case "hello":
                        var v19 string = ""
                        _ = v19
                    case "\n\t":
                        delete(v2, v6)
                    case "":
                        if (1 < v4) {
                            v20 := GenS1{F0: uint8(v3)}
                            _ = v20
                        }
                    default:
                        v21 := []int{}
                        var v22 uint64 = 0
                        if (18446744073709551615 >= (^(v3))) {
                            var v23 bool = (1000 == 32767)
                            var v24 uint16 = v5
                            _ = v23
                            _ = v24
                        }
                        _ = v21
                        _ = v22
                    }
                    v25 := GenS2{F0: (v4 >> ((1) & 63)), F1: false}
                    v26 := GenS2{F0: 415, F1: ("" <= "foo bar")}
                    _ = v25
                    _ = v26
                }
                _ = fi12
            }
            v27 := []int{(-(fi11)), (v1 & -325), int(v4), (-(fi11))}
            _ = fi11
            _ = v27
        }
    }
    fmt.Println(v3)
    fmt.Println(v5, v4)
    v1 = v1
    var v28 uint8 = 128
    v29 := GenS3{F0: 0, F1: v28}
    var v30 int = (v1 + v1)
    v4 = -361
    v31 := GenS2{F0: (^(v4)), F1: ("" == "a")}
    v5 = v5
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 10; fi33++ {
            v28 = 0
            _ = fi33
        }
        _ = fi32
    }
    delete(v2, uint32(v28))
    var v34 float32 = 1e10
    v35 := GenS4{F0: 1, F1: (^(v6)), F2: (v4 & 1)}
    var v36 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
