package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint16; F1 int; F2 uint16 }
type GenS4 struct { F0 int32; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int64 }

func main() {
    v1 := []uint32{2147483647, 462}
    var v2 float32 = 0.1
    v3 := []string{"\n\t"}
    v4 := []string{"", (("\n\t" + "foo bar") + "\n\t"), "\n\t", "\n\t"}
    var v5 int32 = -325
    var v6 int8 = -1
    switch 32767 {
    case 127:
        fmt.Println(v6)
    default:
        var v7 float32 = (v2 - float32(float64((1e10 * (v2 * v2)))))
        _ = v7
    }
    switch ("a" + ("foo bar" + "")) {
    case "a":
        var v8 int = 127
        _ = v8
    case "\n\t":
        fmt.Println(v6, v5, v5)
    case "foo bar":
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                v6 = (v6 >> ((1) & 63))
                v5 = v5
                if ((("hello" + ("a" + "")) + "\n\t") == (("a" + ("\n\t" + "")) + "hello")) {
                    var v11 int = fi9
                    v11 += -31
                    v11 = (fi9 | v11)
                    _ = v11
                }
                _ = fi10
            }
            _ = fi9
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v12 := GenS1{F0: 2, F1: 853, F2: v5}
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS2{F0: 128, F1: (v5 | (v5 - (v5 - 42)))}
        v15 := GenS3{F0: 255, F1: fi13, F2: 42}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint64 = 1
    v6 += v6
    var v17 int16 = 656
    if (0 != 2) {
        v18 := GenS4{F0: v5, F1: (v17 << ((2) & 63)), F2: 2}
        v16 += (^(v16))
        v19 := GenS5{F0: 499, F1: 1}
        _ = v18
        _ = v19
    }
    switch 9223372036854775807 {
    case 32767:
        var v20 uint16 = 65535
        _ = v20
    case -2147483648:
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 bool = true
            if false {
                if (2147483647 == -1) {
                    v23 := GenS1{F0: 32767, F1: 963, F2: (v5 & -2147483648)}
                    var v24 int8 = v6
                    v5 -= (v5 >> ((205) & 63))
                    _ = v23
                    _ = v24
                }
                v5 = v5
                for fi25 := 0; fi25 < 12; fi25++ {
                    v2 += (v2 * (v2 + v2))
                    v26 := GenS4{F0: v5, F1: v17, F2: 127}
                    fmt.Println(v16, v6)
                    _ = fi25
                    _ = v26
                }
            } else {
                v4 = append(v4, "\n\t")
            }
            _ = fi21
            _ = v22
        }
    }
    var v27 bool = (!(255 < 736))
    var v28 uint8 = 127
    var v29 int = -308
    fmt.Println(v17, v17, v6)
    if len(v1) > 0 {
        v1[0] = uint32(v16)
    }
    if (v17 == (v17 % ((v17) | 1))) {
        v30 := GenS1{F0: 255, F1: 431, F2: 0}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v30
    }
    var v31 int64 = -135
    fmt.Println(v5)
    var v32 float64 = (0.1 + -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
