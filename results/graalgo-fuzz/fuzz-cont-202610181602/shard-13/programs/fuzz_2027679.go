package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int8) uint32 {
    var v5 int64 = 1
    var v6 uint = 2
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 uint32, p1 bool, p2 uint64) string {
    v7 := GenS1{F0: -103, F1: 54, F2: (p0 << ((0) & 63))}
    var v8 int64 = -64
    switch 32768 {
    case 127:
        v9 := []int{9223372036854775807}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ("a" + "foo bar")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 1e10
    var v3 float64 = 1e10
    var v4 int8 = 1
    var v10 float64 = v3
    v11 := []int{-254, 42, 950}
    v2 = 3.14159
    v10 = float64((v2 - v2))
    _ = len(v1)
    switch (("hello" + "a") + "\n\t") {
    case "foo bar":
        v11 = append(v11, 0)
    case "hello":
        v3 -= float64((v2 * v2))
    case "\n\t":
        for fi12 := 0; fi12 < 14; fi12++ {
            if (((!(((((v2 / v2) / -2.25) != v2) || ((-2147483648 == -1) && true)) && (fi12 <= -1))) || ((32768 == -744) && (("hello" <= "") && false))) || (!(!true))) {
                if ((3.14159 / v2) != float32(v3)) {
                    if len(v11) > 0 {
                        v11[0] = -2147483648
                    }
                } else {
                    var v13 uint16 = 1
                    v10 += (((v3 - v10) + 0.0) - 1.5)
                    var v14 bool = true
                    _ = v13
                    _ = v14
                }
            } else {
                if true {
                    var v15 int64 = -1
                    _ = v15
                }
            }
            _ = fi12
        }
    }
    v1[genFn0((v4 >> ((42) & 63)))] = true
    fmt.Println(v2, v10, v4)
    switch ("" + "") {
    case "":
        switch "a" {
        case "":
            v16 := GenS2{F0: 2}
            _ = v16
        case "hello":
            v17 := v1[2147483647]
            _ = v17
        default:
            var v18 int8 = (v4 * v4)
            _ = v18
        }
    case "a":
        if (8 != 127) {
            delete(v1, 80)
            fmt.Println(v2, v2)
        } else {
            v4 += v4
        }
    default:
        if ((!((v4 ^ v4) <= (v4 / (((v4 / ((v4) | 1))) | 1)))) && ("a" > "\n\t")) {
            switch 2147483647 {
            case 48:
                for fi19 := 0; fi19 < 17; fi19++ {
                    v10 -= v3
                    _ = fi19
                }
            }
            v20 := v1[911]
            _ = v20
        }
    }
    v2 = -2.25
    switch 202 {
    case -786:
        v1[2147483647] = (18446744073709551615 > 2)
    default:
        var v21 int64 = -864
        _ = v21
    }
    var v22 float64 = 1.5
    for fi23 := 0; fi23 < 2; fi23++ {
        if (!false) {
            var v24 string = "\n\t"
            for fi25 := 0; fi25 < 2; fi25++ {
                v26 := v1[1]
                _ = fi25
                _ = v26
            }
            v10 += v3
            _ = v24
        }
        _ = fi23
    }
    switch 1 {
    case -885:
        v1[42] = true
    case 32768:
        if false {
            v22 -= v22
            fmt.Println(v4, v10)
            switch 32768 {
            case -762:
                v2 -= 0.1
            case 4294967295:
                v1[42] = false
            }
        } else {
            fmt.Println(v3, v3, v3)
            for fi27 := 0; fi27 < 14; fi27++ {
                v10 = v22
                _ = fi27
            }
            v28 := GenS3{F0: 2, F1: 255}
            _ = v28
        }
    default:
        switch "a" {
        case "a":
            var v29 uint64 = 18446744073709551615
            _ = v29
        case "":
            var v30 uint16 = 42
            _ = v30
        case "\n\t":
            var v31 uint16 = 466
            _ = v31
        }
    }
    switch 128 {
    case 32767:
        v4 = v4
    }
    var v32 bool = ((18446744073709551615 <= 3) && (!((v10 + v10) == v10)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v4
}
