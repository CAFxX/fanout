package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint32; F1 bool; F2 int }
func genFn0() bool {
    var v7 uint32 = 1
    var v8 float32 = (((1.5 / (1e10 - ((0.0 - ((0.0 - 1e10) / 0.1)) * ((1e10 / (0.1 * 3.14159)) + (0.1 - 1e10))))) * 0.1) + (1.5 * float32(0.1)))
    var v9 string = ""
    v10 := GenS1{F0: false, F1: 0}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (true || ((v8 / ((1.5 / ((v8 - v8) * (v8 * (v8 * v8)))) - (1e10 / (v8 + (float32((3.14159 - (0.0 / 3.14159))) + v8))))) <= 3.14159))
}
func genFn1() bool {
    var v11 int32 = 1
    var v12 int32 = v11
    _ = v11
    _ = v12
    return (!(!true))
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]string)
    var v3 bool = true
    var v4 int64 = 0
    var v5 float64 = (float64((-2.25 / 1.0)) / 1e10)
    var v6 uint32 = 2147483647
    fmt.Println(v5, v5)
    v13 := GenS2{F0: (v6 % ((v6) | 1)), F1: v3, F2: 117}
    for fi14 := 0; fi14 < 3; fi14++ {
        switch (-(fi14)) {
        case -941:
            for fi15 := 0; fi15 < 16; fi15++ {
                v2[(v6 & v6)] = (v1 + "")
                if v3 {
                    switch (fi15 / ((-1) | 1)) {
                    case -828:
                        var v16 uint16 = 42
                        delete(v2, 1)
                        if (1000 == 138) {
                            var v17 uint64 = 18446744073709551615
                            v5 = (v5 * float64(0.1))
                            v4 = v4
                            _ = v17
                        }
                        _ = v16
                    case -659:
                        var v18 int8 = 1
                        _ = v18
                    }
                    v5 = v5
                    v19 := v2[(^(v6))]
                    _ = v19
                }
                _ = fi15
            }
        }
        v20 := []int{(^(fi14)), (fi14 | fi14), fi14}
        if v3 {
            var v21 string = (v1 + "")
            _ = v21
        } else {
            var v22 float32 = float32(1e10)
            fmt.Println(v22)
            _ = v22
        }
        _ = fi14
        _ = v20
    }
    var v23 string = "hello"
    v24 := []int{-104, 127, 32767}
    for fi25 := 0; fi25 < 7; fi25++ {
        switch (fi25 - 4294967295) {
        case 128:
            v5 = (v5 - 0.1)
        case 725:
            for fi26 := 0; fi26 < 8; fi26++ {
                switch (((v23 + (v1 + v1)) + ("hello" + v1)) + ("" + ((v1 + ("a" + "hello")) + ("foo bar" + v1)))) {
                case "a":
                    v2[v6] = "\n\t"
                case "hello":
                    var v27 int64 = int64(v6)
                    _ = v27
                case "":
                    var v28 string = (v23 + v23)
                    _ = v28
                default:
                    v29 := GenS1{F0: (42 == -954), F1: 42}
                    _ = v29
                }
                fmt.Println(v23, fi26)
                _ = fi26
            }
        }
        var v30 string = "hello"
        v4 = v4
        _ = fi25
        _ = v30
    }
    var v31 uint32 = (v6 + (^(v6)))
    if (!true) {
        var v32 int16 = -32768
        _ = v32
    }
    var v33 uint16 = 0
    switch 2147483647 {
    case -1:
        var v34 int8 = 1
        _ = v34
    case 390:
        v35 := GenS1{F0: ((!false) && (false || v3)), F1: 42}
        _ = v35
    case 9223372036854775807:
        var v36 string = v1
        _ = v36
    default:
        v23 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
