package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int) int8 {
    switch (p0 / ((p0) | 1)) {
    case 127:
        if (((0.1 - 3.14159) - -2.25) <= 0.1) {
            switch (("\n\t" + "foo bar") + "\n\t") {
            case "foo bar":
                var v7 float32 = 1e10
                _ = v7
            default:
                fmt.Println(p0)
            }
            v8 := []int{541, -1}
            _ = v8
        } else {
            var v9 int = (p0 & p0)
            v10 := []int{}
            var v11 uint = 4294967295
            _ = v10
            _ = v11
            _ = v9
        }
    }
    v12 := GenS1{F0: 65535}
    _ = p0
    _ = v12
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-392}
    var v3 uint16 = 255
    v4 := make(map[uint32]bool)
    var v5 string = (("a" + "\n\t") + "a")
    var v6 float64 = (0.1 - float64((1.5 - 1e10)))
    if false {
        for fi13 := 0; fi13 < 7; fi13++ {
            switch v5 {
            case "hello":
                v2 = append(v2, (-(fi13)))
            case "foo bar":
                if ((496 < 2147483648) || false) {
                    fmt.Println(v5, v3, fi13)
                    for fi14 := 0; fi14 < 20; fi14++ {
                        if ((42 == -361) || (0.0 > 0.0)) {
                            if (fi14 != fi13) {
                                v15 := GenS1{F0: 0}
                                v16 := GenS1{F0: 2}
                                switch (v5 + (v5 + v5)) {
                                case "foo bar":
                                    var v17 int32 = 919
                                    _ = v17
                                default:
                                    var v18 int16 = 918
                                    _ = v18
                                }
                                _ = v15
                                _ = v16
                            }
                            _ = len(v4)
                        } else {
                            if (1 != 0) {
                                v19 := GenS2{F0: -9223372036854775808}
                                if (-1 >= 576) {
                                    fmt.Println(v3, v3)
                                    switch v5 {
                                    case "hello":
                                        switch 668 {
                                        case -9223372036854775808:
                                            fmt.Println(fi13, v5)
                                        case 128:
                                            var v20 uint32 = 636
                                            _ = v20
                                        }
                                    }
                                }
                                fmt.Println(v3, v6, fi14)
                                _ = v19
                            }
                            var v21 uint8 = 49
                            delete(v1, fi13)
                            _ = v21
                        }
                        _ = fi14
                    }
                    v22 := GenS2{F0: -433}
                    _ = v22
                }
            case "a":
                var v23 string = v5
                _ = v23
            }
            if ((fi13 | (fi13 ^ fi13)) > (^(fi13))) {
                if true {
                    switch (v5 + "foo bar") {
                    case "foo bar":
                        v6 -= 0.1
                    case "\n\t":
                        v5 = "\n\t"
                    case "a":
                        var v24 uint = 387
                        _ = v24
                    default:
                        var v25 uint = 65535
                        _ = v25
                    }
                }
            }
            if len(v2) > 0 {
                v2[0] = 0
            }
            _ = fi13
        }
        v3 = v3
        v6 = ((v6 + v6) + v6)
    }
    v6 = v6
    v26 := GenS1{F0: 0}
    fmt.Println(v3, v5, v3)
    var v27 uint32 = 1
    v28 := GenS3{F0: (^(v3)), F1: 179}
    v27 = v27
    v3 = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
