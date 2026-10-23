package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0() uint32 {
    switch "hello" {
    case "":
        var v7 int8 = -128
        _ = v7
    case "a":
        if false {
            switch "\n\t" {
            case "\n\t":
                var v8 uint32 = 4294967295
                _ = v8
            default:
                v9 := GenS1{F0: 820, F1: 23, F2: (1.5 <= 0.1)}
                _ = v9
            }
            var v10 uint32 = 948
            if true {
                v11 := []int{9223372036854775807, 1}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                var v12 uint32 = (v10 | (v10 * 1))
                _ = v11
                _ = v12
            } else {
                if (1 > 127) {
                    var v13 int64 = 0
                    _ = v13
                }
                v14 := []int{127}
                fmt.Println(v10)
                _ = v14
            }
            _ = v10
        } else {
            v15 := GenS1{F0: 0, F1: 2, F2: true}
            var v16 int = 1
            _ = v15
            _ = v16
        }
    }
    if false {
        var v17 bool = false
        _ = v17
    } else {
        var v18 uint32 = 2147483647
        _ = v18
    }
    return 671
}

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 42
    v3 := []uint32{0}
    var v4 uint64 = 679
    var v5 float32 = (3.14159 * 1e10)
    v6 := make(map[int]string)
    if true {
        v19 := GenS1{F0: v4, F1: 128, F2: (!((255 != 2) || true))}
        switch "hello" {
        case "hello":
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 int64 = -1
                fmt.Println(v21, v21)
                _ = fi20
                _ = v21
            }
        case "a":
            for fi22 := 0; fi22 < 3; fi22++ {
                switch (("\n\t" + "\n\t") + (("foo bar" + "") + "a")) {
                case "":
                    v1 = 1e10
                case "\n\t":
                    if true {
                        switch "" {
                        case "hello":
                            var v23 int = (fi22 % (((fi22 & -1)) | 1))
                            _ = v23
                        }
                        v2 += (v2 ^ 255)
                        var v24 int64 = 42
                        _ = v24
                    }
                case "foo bar":
                    _ = len(v6)
                default:
                    v25 := GenS1{F0: v4, F1: 244, F2: ("\n\t" >= ("" + ("hello" + "")))}
                    _ = v25
                }
                _ = fi22
            }
        case "foo bar":
            v5 = (-2.25 * v5)
        }
        _ = v19
    }
    v26 := GenS2{F0: 42, F1: 2}
    if true {
        if (573 == 0) {
            var v27 uint = 211
            _ = v27
        }
    }
    v5 = v5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v28 int64 = 1
    switch -21 {
    case -907:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v29 uint8 = 2
    var v30 float64 = float64((float32((v1 / v1)) / 0.1))
    var v31 int16 = 42
    switch -1 {
    case 153:
        var v32 uint8 = v29
        _ = v32
    case -9223372036854775808:
        v29 -= v29
    case 2147483647:
        delete(v6, 127)
    }
    var v33 uint = 1
    var v34 uint64 = v4
    var v35 uint64 = (v4 << ((v33) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
