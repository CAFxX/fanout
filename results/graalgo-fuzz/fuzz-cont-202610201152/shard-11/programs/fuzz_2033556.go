package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 bool; F1 int16; F2 int8 }
type GenS3 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = (-2.25 - 1.5)
    v3 := []int{32767, 641, 32767, 127}
    var v4 string = "foo bar"
    var v5 string = (v4 + (((("foo bar" + "\n\t") + v4) + "") + v4))
    v4 = ("hello" + v5)
    var v6 uint8 = 255
    v2 = 0.1
    v4 = "\n\t"
    v6 = v6
    if true {
        v5 = ""
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 float64 = -2.25
            switch v4 {
            case "a":
                for fi9 := 0; fi9 < 16; fi9++ {
                    v6 = (^(v6))
                    if len(v3) > 0 {
                        v3[0] = (fi7 ^ (fi9 >> ((42) & 63)))
                    }
                    _ = fi9
                }
            case "":
                var v10 uint16 = 1
                _ = v10
            case "\n\t":
                var v11 uint8 = v6
                _ = v11
            default:
                var v12 uint32 = 42
                _ = v12
            }
            _ = fi7
            _ = v8
        }
    }
    if ((-290 >= -1) || true) {
        v13 := GenS1{F0: 2147483647, F1: 214}
        v14 := v1[2147483648]
        v3 = append(v3, 164)
        _ = v13
        _ = v14
    } else {
        var v15 int64 = 0
        v3 = append(v3, 359)
        _ = v15
    }
    fmt.Println(v4, v5, v6)
    if (true || false) {
        delete(v1, 860)
    }
    v16 := GenS2{F0: true, F1: 1, F2: 100}
    var v17 uint = 407
    for fi18 := 0; fi18 < 4; fi18++ {
        switch (fi18 - fi18) {
        case 1:
            v19 := GenS2{F0: true, F1: -32768, F2: 127}
            _ = v19
        case 32768:
            if ((v4 <= (v4 + "a")) && true) {
                var v20 int16 = 1
                if false {
                    fmt.Println(v4, v17, v6)
                    v17 = v17
                }
                _ = v20
            } else {
                fmt.Println(v2, v5)
            }
        case 9223372036854775807:
            var v21 uint64 = 2
            _ = v21
        }
        v22 := GenS3{F0: -28, F1: 111}
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        _ = fi18
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
