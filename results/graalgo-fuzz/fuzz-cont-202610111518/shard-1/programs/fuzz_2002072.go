package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint8) string {
    if true {
        var v5 uint64 = 18446744073709551615
        if (-1 <= 0) {
            if (48 != 4294967295) {
                var v6 bool = (2 > 1)
                _ = v6
            } else {
                v7 := []int{}
                v8 := GenS1{F0: 1, F1: false, F2: 65535}
                var v9 int16 = 1
                _ = v7
                _ = v8
                _ = v9
            }
            p0 = (p0 % (((p0 % (((p0 - (^(p0)))) | 1))) | 1))
            v10 := []int{32767, 128, int(v5), 4294967295}
            _ = v10
        } else {
            v11 := []int{9223372036854775807, int(p0)}
            var v12 int = 32768
            _ = v11
            _ = v12
        }
        var v13 uint16 = 290
        _ = v13
        _ = v5
    }
    v14 := []int{}
    _ = p0
    _ = v14
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{513, 0}
    var v3 int32 = -190
    v4 := []int{-2147483648}
    v3 = -1
    _ = len(v1)
    var v15 bool = false
    var v16 uint8 = 128
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v17 := GenS2{F0: 2, F1: int8(v3)}
    v16 -= (v16 - (v16 + (v16 * 175)))
    switch (("" + ("a" + ("" + ((("" + "hello") + "a") + "hello")))) + "\n\t") {
    case "foo bar":
        var v18 bool = (0 >= 1)
        _ = v18
    default:
        if ((("\n\t" + "\n\t") >= ("\n\t" + "hello")) && (1.5 < -2.25)) {
            switch "a" {
            case "hello":
                v4 = append(v4, 32768)
            case "\n\t":
                var v19 string = ((("" + ("hello" + "")) + "\n\t") + "\n\t")
                _ = v19
            case "foo bar":
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 int16 = 32767
                    v16 = (v16 + v16)
                    fmt.Println(v15, v3, fi20)
                    _ = fi20
                    _ = v21
                }
            default:
                v15 = (-9223372036854775808 <= -695)
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
