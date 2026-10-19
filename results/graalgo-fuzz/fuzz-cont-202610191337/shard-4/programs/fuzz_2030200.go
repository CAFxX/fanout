package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 bool; F2 uint16 }

func main() {
    var v1 float64 = (0.0 / (float64(1.5) / 1.0))
    var v2 uint16 = 323
    var v3 uint8 = 42
    v2 = v2
    var v4 uint64 = 242
    v4 = 18446744073709551615
    v4 += (v4 << ((65535) & 63))
    fmt.Println(v2, v1, v3)
    v5 := GenS1{F0: 42}
    if (255 < 839) {
        for fi6 := 0; fi6 < 6; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 uint64 = v4
                var v9 float32 = 0.0
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
        v1 -= v1
        var v10 int32 = 2147483647
        _ = v10
    } else {
        switch 4294967295 {
        case -277:
            v11 := GenS2{F0: 0}
            _ = v11
        case -711:
            switch ("" + ("foo bar" + ("foo bar" + "\n\t"))) {
            case "":
                var v12 uint = 1
                _ = v12
            case "a":
                for fi13 := 0; fi13 < 1; fi13++ {
                    if (("a" + "hello") == (((((("a" + "a") + "a") + "a") + "hello") + "foo bar") + "foo bar")) {
                        fmt.Println(v1, v4, v1)
                    } else {
                        for fi14 := 0; fi14 < 7; fi14++ {
                            if true {
                                v15 := GenS3{F0: (v3 == 1), F1: (3.14159 < (1e10 / 3.14159)), F2: v2}
                                _ = v15
                            } else {
                                if true {
                                    v1 = v1
                                    if (!(false || true)) {
                                        v16 := []int{(fi14 >> ((65535) & 63)), fi14, (fi13 / ((9223372036854775807) | 1)), fi13}
                                        _ = v16
                                    }
                                }
                                var v17 float32 = 1e10
                                v1 -= v1
                                _ = v17
                            }
                            _ = fi14
                        }
                        v2 = 42
                        var v18 uint32 = 65535
                        _ = v18
                    }
                    v1 += (((-2.25 - (float64((((0.0 * 0.1) * (3.14159 * 0.1)) + 3.14159)) / -2.25)) + v1) + float64(((3.14159 * -2.25) - 1.5)))
                    _ = fi13
                }
            default:
                v19 := []int{}
                _ = v19
            }
        case -1:
            v2 = uint16(v3)
        }
        v20 := GenS3{F0: (0 > 4294967295), F1: ((v2 ^ v2) <= (v2 + v2)), F2: (v2 - v2)}
        _ = v20
    }
    var v21 string = ("hello" + "hello")
    var v22 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
