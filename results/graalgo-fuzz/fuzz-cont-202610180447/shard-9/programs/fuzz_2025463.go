package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int, p1 uint8) uint {
    var v5 uint64 = 42
    var v6 uint16 = 2
    if (0 > 1) {
        v7 := []int{p0, 668}
        if (!(!false)) {
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 int32 = 0
                var v10 uint16 = v6
                if true {
                    _ = len(v7)
                    v10 = 255
                } else {
                    p1 = 2
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v12 := GenS1{F0: (int32(v5) > (^(v9))), F1: 42, F2: -1}
                        _ = fi11
                        _ = v12
                    }
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(v5, p1)
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 int64 = 1
                _ = fi13
                _ = v14
            }
        }
        v5 -= 633
        _ = v7
    } else {
        switch (p0 / ((p0) | 1)) {
        case 32768:
            v5 = v5
        case 4294967295:
            p1 += (p1 * p1)
        }
        p0 += (p0 + p0)
        v15 := []int{p0, (p0 >> ((54) & 63))}
        _ = v15
    }
    v6 -= 42
    switch "" {
    case "foo bar":
        var v16 int16 = -32768
        _ = v16
    case "\n\t":
        var v17 int32 = 2147483647
        _ = v17
    case "":
        p1 -= 2
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return uint(v6)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 18446744073709551615
    var v3 int16 = -913
    v4 := make(map[int]bool)
    v4[127] = (426 == 168)
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 int32 = 422
        fmt.Println(v3, v19, fi18)
        if false {
            for fi20 := 0; fi20 < 17; fi20++ {
                v21 := []int{(fi18 * (fi18 + fi20)), (fi20 / (((-(fi20))) | 1)), fi18, 1}
                if (2147483647 <= 65535) {
                    v22 := GenS2{F0: 141, F1: 42}
                    v23 := GenS1{F0: (true && true), F1: 116, F2: (v3 / ((v3) | 1))}
                    _ = v22
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            v2 = 152
            var v24 int8 = 0
            _ = v24
        } else {
            for fi25 := 0; fi25 < 11; fi25++ {
                fmt.Println(v3)
                var v26 float32 = (0.1 + 1e10)
                v27 := v4[1]
                _ = fi25
                _ = v26
                _ = v27
            }
            if (!false) {
                fmt.Println(v3, v3)
            }
        }
        _ = fi18
        _ = v19
    }
    var v28 uint32 = 42
    v29 := v4[-1]
    delete(v1, 32767)
    _ = len(v1)
    var v30 string = ""
    fmt.Println(v3, v3)
    var v31 float64 = float64(((float32(3.14159) - (-2.25 * ((1.5 / 1.0) - float32((-2.25 - 0.0))))) - -2.25))
    var v32 uint = genFn0(-564, 231)
    var v33 uint8 = 2
    var v34 string = "hello"
    if (!v29) {
        v30 = v30
    }
    var v35 float64 = v31
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
