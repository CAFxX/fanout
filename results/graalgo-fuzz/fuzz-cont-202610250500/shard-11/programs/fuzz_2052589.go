package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int8 }

func main() {
    v1 := []uint32{}
    var v2 uint32 = 464
    var v3 int8 = -45
    var v4 string = ((((("\n\t" + ("a" + ("foo bar" + "hello"))) + ("hello" + "hello")) + "hello") + ("hello" + "a")) + ("hello" + "\n\t"))
    var v5 float64 = float64(3.14159)
    v3 = v3
    if (2147483648 > v2) {
        for fi6 := 0; fi6 < 8; fi6++ {
            switch v4 {
            case "a":
                if ((185 >= 1) || ((!(!true)) && false)) {
                    fmt.Println(v2, v5)
                    for fi7 := 0; fi7 < 1; fi7++ {
                        if (1000 != 0) {
                            var v8 uint32 = 65535
                            v9 := GenS1{F0: -814, F1: (v8 & v2)}
                            _ = v8
                            _ = v9
                        }
                        v10 := GenS1{F0: -232, F1: v2}
                        _ = fi7
                        _ = v10
                    }
                    if len(v1) > 0 {
                        v1[0] = (v2 & v2)
                    }
                }
            }
            _ = fi6
        }
        v11 := GenS1{F0: 256, F1: 0}
        v12 := GenS1{F0: -1, F1: v2}
        _ = v11
        _ = v12
    }
    v1 = append(v1, v2)
    for fi13 := 0; fi13 < 7; fi13++ {
        v3 = v3
        fmt.Println(v3, v4)
        if true {
            switch (((((v4 + v4) + "") + v4) + ((("\n\t" + v4) + (v4 + (v4 + ""))) + v4)) + (v4 + "hello")) {
            case "a":
                if ((0 == 255) || false) {
                    if true {
                        switch int(v3) {
                        case 841:
                            fmt.Println(v4)
                        }
                        if (true && true) {
                            v14 := GenS2{F0: 202}
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 6; fi15++ {
                            var v16 int = 32768
                            var v17 int16 = -32768
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                } else {
                    _ = len(v1)
                    var v18 uint64 = 42
                    _ = v18
                }
            default:
                switch fi13 {
                case -1:
                    var v19 int = (fi13 >> ((42) & 63))
                    _ = v19
                }
            }
            if true {
                v5 -= (float64(1.5) + v5)
                var v20 string = ((((v4 + v4) + ((v4 + "foo bar") + ("foo bar" + "\n\t"))) + "hello") + (("" + v4) + v4))
                fmt.Println(v20, v5)
                _ = v20
            } else {
                v21 := GenS3{F0: 1}
                var v22 int64 = 1
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS1{F0: 150, F1: v2}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v25 := GenS1{F0: 625, F1: v2}
                v26 := GenS1{F0: 771, F1: v2}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
        }
        _ = fi13
    }
    var v27 string = v4
    if len(v1) > 0 {
        v1[0] = v2
    }
    v5 += (3.14159 - v5)
    if (0 == v2) {
        v28 := GenS1{F0: 9223372036854775807, F1: v2}
        var v29 int64 = -840
        _ = v28
        _ = v29
    } else {
        if false {
            if ((412 > 871) || (v5 > (1e10 - 1.5))) {
                v5 = v5
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 int8 = 127
                    _ = fi30
                    _ = v31
                }
                v32 := GenS4{F0: v2, F1: v3}
                _ = v32
            }
        } else {
            fmt.Println(v5)
            v33 := GenS1{F0: 0, F1: v2}
            _ = v33
        }
        v3 += v3
    }
    v5 += float64((0.1 - float32(v5)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
