package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint64; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int16; F2 int16 }
func genFn0() int64 {
    var v7 int64 = 514
    var v8 float32 = (1e10 + float32((1e10 - -2.25)))
    _ = v7
    _ = v8
    return (v7 & v7)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 0
    v3 := []int{-730}
    var v4 int16 = 1000
    var v5 int16 = v4
    var v6 bool = ("hello" < "\n\t")
    var v9 int32 = 886
    fmt.Println(v6, v5)
    v10 := GenS1{F0: (v2 >> ((420) & 63))}
    var v11 int32 = v9
    v12 := GenS2{F0: -128, F1: 320, F2: 698}
    switch 2147483647 {
    case -718:
        switch (("" + "foo bar") + ("" + (("\n\t" + ("foo bar" + ((("a" + ("\n\t" + (("a" + "\n\t") + ("foo bar" + "\n\t")))) + "hello") + ("" + ("" + "\n\t"))))) + "foo bar"))) {
        case "hello":
            delete(v1, 4294967295)
        case "a":
            for fi13 := 0; fi13 < 17; fi13++ {
                if v6 {
                    fmt.Println(v6)
                    if ((360 != 95) || v6) {
                        v11 += 0
                        var v14 string = "foo bar"
                        _ = v14
                    } else {
                        switch (fi13 + fi13) {
                        case -1:
                            v15 := GenS1{F0: (v2 / ((v2) | 1))}
                            _ = v15
                        }
                    }
                } else {
                    fmt.Println(v6, v6, v11)
                    var v16 int16 = v5
                    _ = v16
                }
                var v17 uint64 = 1
                var v18 uint8 = 1
                _ = fi13
                _ = v17
                _ = v18
            }
        case "\n\t":
            switch 2147483647 {
            case -1:
                var v19 int64 = 1
                _ = v19
            case -487:
                switch (("hello" + "a") + "hello") {
                case "\n\t":
                    var v20 int32 = v9
                    _ = v20
                case "a":
                    fmt.Println(v5)
                default:
                    v3 = append(v3, 722)
                }
            }
        default:
            if v6 {
                v21 := v1[2147483648]
                _ = v21
            }
        }
    case -183:
        var v22 bool = (0 < uint32(v2))
        _ = v22
    default:
        v6 = true
    }
    v23 := GenS3{F0: 1}
    var v24 uint64 = 967
    switch 1 {
    case 9223372036854775807:
        var v25 bool = (0.1 >= (((3.14159 - 1.5) / -2.25) * (1e10 + (3.14159 / 3.14159))))
        _ = v25
    case 0:
        var v26 float32 = 0.0
        _ = v26
    case 42:
        delete(v1, 1)
    default:
        var v27 int = 4294967295
        _ = v27
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        switch "" {
        case "":
            if (v24 >= v24) {
                v9 -= 2147483647
                v29 := GenS3{F0: 0}
                var v30 float64 = (1.5 - float64((3.14159 - (-2.25 + -2.25))))
                _ = v29
                _ = v30
            }
        }
        _ = fi28
    }
    v31 := GenS4{F0: v4, F1: (v5 | v4), F2: (v4 % ((1) | 1))}
    var v32 uint16 = 65535
    v4 += (v4 + int16(v32))
    v9 -= (v11 & (v11 * 582))
    var v33 int32 = v11
    for fi34 := 0; fi34 < 15; fi34++ {
        v1[0] = 2147483647
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
