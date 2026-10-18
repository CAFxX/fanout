package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint }

func main() {
    var v1 int = 128
    var v2 uint32 = 4294967295
    var v3 int = v1
    var v4 uint32 = v2
    v5 := []uint32{2, (v2 & v4)}
    v6 := make(map[string]bool)
    v3 += 128
    fmt.Println(v3, v2, v4)
    if (!(false || (0 < -2147483648))) {
        v7 := GenS1{F0: 198}
        _ = v7
    }
    var v8 uint8 = 1
    var v9 float32 = float32((3.14159 + float64(1.5)))
    v10 := GenS2{F0: 127, F1: 1}
    switch "a" {
    case "foo bar":
        switch (^(v1)) {
        case 2147483647:
            if (v9 == -2.25) {
                _ = len(v5)
                var v11 int64 = 42
                _ = v11
            } else {
                switch v1 {
                case 0:
                    v12 := GenS1{F0: (v8 >> ((42) & 63))}
                    _ = v12
                }
            }
        }
    case "hello":
        v13 := v6[((("hello" + "hello") + "hello") + ("\n\t" + "\n\t"))]
        _ = v13
    default:
        if ((-868 > 0) || (21 == -37)) {
            if (!(9223372036854775807 <= 42)) {
                var v14 uint = 42
                _ = v14
            } else {
                if (!(!false)) {
                    if false {
                        fmt.Println(v4, v4, v2)
                    }
                    v4 -= (v2 << ((789) & 63))
                }
            }
            v8 = v8
        } else {
            var v15 uint32 = (^(v4))
            if len(v5) > 0 {
                v5[0] = (^(v4))
            }
            if true {
                switch int(v4) {
                case -953:
                    if (((1e10 / 3.14159) >= ((0.0 * -2.25) + 0.1)) || true) {
                        if len(v5) > 0 {
                            v5[0] = (v15 % ((v2) | 1))
                        }
                    }
                case 42:
                    for fi16 := 0; fi16 < 19; fi16++ {
                        if len(v5) > 0 {
                            v5[0] = v15
                        }
                        v17 := GenS2{F0: v3, F1: 0}
                        var v18 bool = (((v1 <= v3) && (false || true)) || false)
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                case 1:
                    v9 = v9
                default:
                    delete(v6, ("hello" + "foo bar"))
                }
                v1 -= v3
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                for fi19 := 0; fi19 < 16; fi19++ {
                    var v20 uint8 = 168
                    fmt.Println(v2, v3)
                    _ = fi19
                    _ = v20
                }
            }
            _ = v15
        }
    }
    var v21 int32 = 2147483647
    var v22 string = "foo bar"
    if (v2 >= (^(v2))) {
        _ = len(v6)
        v23 := GenS2{F0: (-(v3)), F1: 0}
        _ = v23
    } else {
        v9 += v9
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := GenS1{F0: v8}
            v22 = v22
            _ = fi24
            _ = v25
        }
        if (!(true && false)) {
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 float64 = (-2.25 - float64(v9))
                for fi28 := 0; fi28 < 15; fi28++ {
                    if len(v5) > 0 {
                        v5[0] = (^(v4))
                    }
                    _ = fi28
                }
                var v29 uint32 = (v2 - (^(v2)))
                _ = fi26
                _ = v27
                _ = v29
            }
            v30 := GenS3{F0: v8, F1: -1, F2: 420}
            if true {
                fmt.Println(v21, v22)
            }
            _ = v30
        } else {
            var v31 uint = 0
            _ = v31
        }
    }
    var v32 bool = true
    var v33 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
