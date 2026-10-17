package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int64 }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 int64) uint64 {
    var v6 bool = true
    fmt.Println(p0, p0)
    switch "a" {
    case "foo bar":
        p0 = p0
    case "hello":
        fmt.Println(p0)
    case "":
        var v7 uint8 = 255
        _ = v7
    default:
        v8 := GenS1{F0: 91, F1: 9223372036854775807, F2: p0}
        _ = v8
    }
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 uint8) uint {
    var v9 float64 = (0.1 / -2.25)
    switch "a" {
    case "foo bar":
        fmt.Println(p0)
    }
    var v10 int16 = 32767
    v10 = 966
    v11 := GenS2{F0: 42, F1: 65535, F2: ("foo bar" + "a")}
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 2147483647
    var v3 string = "a"
    var v4 int8 = -16
    var v5 int = -75
    v12 := v1[65535]
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v2)
        if true {
            var v14 uint32 = v12
            v1[(^(v14))] = 2147483647
            if false {
                v15 := v1[(v14 | v12)]
                _ = v15
            } else {
                if (v3 < (("foo bar" + "") + "\n\t")) {
                    v4 = (v4 + -128)
                }
                v16 := []int{v5, (v5 ^ v5), v5}
                for fi17 := 0; fi17 < 9; fi17++ {
                    _ = len(v1)
                    switch (fi13 >> ((4294967295) & 63)) {
                    case 1:
                        if true {
                            var v18 int32 = -670
                            var v19 float32 = 1e10
                            _ = v18
                            _ = v19
                        }
                    case 4294967295:
                        var v20 uint16 = 813
                        _ = v20
                    default:
                        var v21 bool = (0 <= 0)
                        v14 = (v12 * 4294967295)
                        _ = v21
                    }
                    switch (v5 ^ (fi17 + (fi17 % ((v5) | 1)))) {
                    case -1:
                        var v22 int32 = (v2 & 111)
                        _ = v22
                    }
                    _ = fi17
                }
                _ = v16
            }
            _ = v14
        } else {
            v23 := GenS2{F0: v2, F1: (v12 << ((genFn1(0)) & 63)), F2: ""}
            v24 := []int{(v5 * v5), v5, fi13, (v5 | (v5 | (v5 - (fi13 - -114))))}
            var v25 uint64 = genFn0(int64(fi13))
            _ = v23
            _ = v24
            _ = v25
        }
        var v26 int64 = -433
        _ = fi13
        _ = v26
    }
    v3 = "foo bar"
    var v27 uint8 = 255
    var v28 float32 = 0.1
    var v29 int8 = v4
    var v30 uint64 = 0
    v31 := GenS2{F0: 910, F1: 2147483647, F2: v3}
    var v32 int = v5
    v27 = (v27 ^ v27)
    var v33 uint = genFn1(v27)
    var v34 uint16 = 65535
    var v35 int8 = 127
    for fi36 := 0; fi36 < 8; fi36++ {
        v34 += (v34 << (((v33 / (((v33 >> ((0) & 63))) | 1))) & 63))
        if true {
            v37 := GenS3{F0: true, F1: v2}
            _ = v37
        }
        v32 += 0
        _ = fi36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
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
    _ = v5
}
