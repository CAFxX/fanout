package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }

func main() {
    v1 := []uint32{1}
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    var v4 string = "a"
    var v5 string = "a"
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint = 594
        var v8 float32 = float32(1.5)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint16 = 27
    var v10 string = "foo bar"
    var v11 int = 32767
    switch (v10 + v5) {
    case "":
        if len(v1) > 0 {
            v1[0] = 42
        }
    case "\n\t":
        var v12 uint16 = v9
        _ = v12
    }
    v1 = append(v1, 65535)
    v13 := GenS1{F0: 9223372036854775807, F1: 1}
    v1 = append(v1, 4294967295)
    v11 = (^(v11))
    var v14 float32 = -2.25
    v15 := GenS1{F0: 1, F1: 1}
    var v16 string = "a"
    var v17 string = ""
    switch v11 {
    case 52:
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 uint32 = 42
            fmt.Println(v11, v4)
            _ = fi18
            _ = v19
        }
    case 42:
        var v20 int = int(v9)
        _ = v20
    case -715:
        switch v10 {
        case "a":
            var v21 uint = 2
            _ = v21
        }
    default:
        _ = len(v3)
    }
    switch (v11 + -348) {
    case -238:
        for fi22 := 0; fi22 < 4; fi22++ {
            if ((float64((float32(0.0) / 1.0)) > (0.0 - 0.0)) && false) {
                v2 = (v2 ^ v2)
                var v23 int8 = -128
                var v24 string = "foo bar"
                _ = v23
                _ = v24
            } else {
                switch v11 {
                case 473:
                    if (!(v2 >= v2)) {
                        v10 = ("\n\t" + v5)
                        for fi25 := 0; fi25 < 17; fi25++ {
                            _ = len(v1)
                            if (!true) {
                                v4 = v16
                                v16 = v4
                            }
                            var v26 uint16 = v9
                            _ = fi25
                            _ = v26
                        }
                    }
                default:
                    v11 = (v11 ^ (fi22 + fi22))
                }
                var v27 string = "hello"
                _ = v27
            }
            _ = fi22
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
