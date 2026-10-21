package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 bool; F1 int16 }

func main() {
    var v1 int8 = 42
    v2 := []string{"a", "hello", ("foo bar" + "\n\t"), "\n\t"}
    v3 := []uint32{0, 65535}
    var v4 int64 = 1
    var v5 uint16 = 32767
    if len(v3) > 0 {
        v3[0] = 134
    }
    v1 = 1
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := GenS1{F0: 144}
        switch "" {
        case "foo bar":
            if false {
                var v8 float64 = 0.0
                for fi9 := 0; fi9 < 6; fi9++ {
                    var v10 int64 = (v4 * 1)
                    _ = fi9
                    _ = v10
                }
                switch fi6 {
                case -660:
                    v11 := GenS1{F0: 65535}
                    _ = v11
                case -2147483648:
                    var v12 uint16 = 65535
                    _ = v12
                default:
                    if (v8 != (v8 / float64((0.1 / 0.1)))) {
                        v13 := GenS2{F0: true, F1: 927}
                        _ = v13
                    }
                }
                _ = v8
            }
        case "":
            for fi14 := 0; fi14 < 3; fi14++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi14
            }
        default:
            var v15 uint8 = 128
            _ = v15
        }
        v5 += v5
        _ = fi6
        _ = v7
    }
    var v16 uint16 = (v5 - v5)
    var v17 uint32 = 2147483647
    var v18 uint16 = (v5 << ((132) & 63))
    _ = len(v2)
    fmt.Println(v5, v5)
    v1 = v1
    v4 += v4
    var v19 uint8 = 0
    fmt.Println(v17, v16)
    if (false && true) {
        var v20 int = -2147483648
        v2 = append(v2, "foo bar")
        v21 := GenS1{F0: 2}
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: ((1e10 / float64(-2.25)) > (3.14159 / float64(0.1))), F1: 68}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
