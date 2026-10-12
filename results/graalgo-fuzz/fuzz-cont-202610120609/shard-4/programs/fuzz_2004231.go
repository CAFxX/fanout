package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int; F2 string }
type GenS5 struct { F0 uint32 }
func genFn0() uint64 {
    switch ("foo bar" + "hello") {
    case "foo bar":
        if true {
            v5 := []int{32767}
            _ = v5
        } else {
            switch 0 {
            case 32767:
                switch "\n\t" {
                case "hello":
                    var v6 uint8 = 128
                    _ = v6
                case "\n\t":
                    var v7 uint8 = 1
                    _ = v7
                }
            case 42:
                v8 := []int{516}
                _ = v8
            }
        }
    case "":
        if (4294967295 != 42) {
            var v9 bool = (false && (false || true))
            v10 := []int{}
            _ = v10
            _ = v9
        }
    }
    var v11 float64 = 0.0
    if (104 <= 974) {
        v11 = float64(float32((-2.25 - v11)))
        v11 = ((v11 * (v11 + float64((0.0 * ((float32(v11) * 3.14159) / 3.14159))))) * float64(0.1))
    }
    _ = v11
    return 0
}

func main() {
    var v1 uint16 = 65535
    var v2 float64 = 0.0
    var v3 int16 = -634
    var v4 int16 = (v3 << ((862) & 63))
    fmt.Println(v1, v1)
    v12 := []int{}
    v13 := GenS1{F0: 127, F1: 255, F2: 111}
    for fi14 := 0; fi14 < 13; fi14++ {
        if true {
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 uint = 0
                var v17 int = (-(fi15))
                _ = fi15
                _ = v16
                _ = v17
            }
            v18 := GenS2{F0: (v4 | 689)}
            v1 += v1
            _ = v18
        } else {
            var v19 string = ("" + "\n\t")
            v20 := GenS2{F0: v4}
            if (1e10 > float32((v2 + v2))) {
                var v21 string = ("a" + "foo bar")
                v19 = (((v19 + v21) + v19) + v21)
                _ = v21
            }
            _ = v19
            _ = v20
        }
        _ = fi14
    }
    v22 := GenS3{F0: 67, F1: 1, F2: (v1 / ((2) | 1))}
    fmt.Println(v2)
    var v23 uint64 = 42
    if (32768 == 128) {
        v4 = -770
        switch "a" {
        case "a":
            v1 = v1
        default:
            fmt.Println(v1, v3, v4)
        }
        v12 = append(v12, 127)
    } else {
        fmt.Println(v3, v2, v4)
        fmt.Println(v23, v1)
        var v24 int16 = (v3 >> ((163) & 63))
        _ = v24
    }
    v25 := GenS4{F0: 2, F1: -2147483648, F2: (("\n\t" + "hello") + "a")}
    switch 1 {
    case 267:
        if false {
            if false {
                var v26 int16 = 1000
                var v27 uint64 = (v23 % ((v23) | 1))
                for fi28 := 0; fi28 < 20; fi28++ {
                    if false {
                        if len(v12) > 0 {
                            v12[0] = (fi28 ^ fi28)
                        }
                        for fi29 := 0; fi29 < 2; fi29++ {
                            var v30 uint = 4294967295
                            _ = fi29
                            _ = v30
                        }
                        v31 := GenS5{F0: 760}
                        _ = v31
                    } else {
                        _ = len(v12)
                        var v32 int64 = -1
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = v32
                    }
                    var v33 float32 = (1.5 * -2.25)
                    _ = fi28
                    _ = v33
                }
                _ = v26
                _ = v27
            }
        }
    }
    fmt.Println(v2, v4)
    var v34 bool = (2147483648 < 1)
    if len(v12) > 0 {
        v12[0] = 32768
    }
    v23 = (v23 ^ v23)
    if len(v12) > 0 {
        v12[0] = 2147483647
    }
    if len(v12) > 0 {
        v12[0] = -1
    }
    v34 = (v34 || (2 == (^(v1))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v34
    _ = v4
}
