package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 string = ("hello" + (("" + ("hello" + "a")) + ((("\n\t" + "\n\t") + "foo bar") + "foo bar")))
    var v3 int8 = 99
    v3 += -128
    fmt.Println(v3)
    var v4 int16 = 96
    v5 := GenS1{F0: -91, F1: -854}
    v2 = v2
    v6 := GenS2{F0: 128}
    v1[127] = (v2 + v2)
    var v7 uint32 = 65535
    _ = len(v1)
    v8 := GenS3{F0: 42, F1: int16(v3)}
    v9 := []int{int(v4), 232, -1, 32768}
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int = (^(fi10))
        if (-9223372036854775808 < 9223372036854775807) {
            if (295 == v7) {
                if (!true) {
                    fmt.Println(v3, v4, v11)
                    var v12 int8 = v3
                    _ = v12
                }
                var v13 float32 = 1.5
                _ = v13
            }
            if false {
                delete(v1, v11)
            }
            if len(v9) > 0 {
                v9[0] = -9223372036854775808
            }
        } else {
            v4 -= (-(v4))
        }
        if (true && (!true)) {
            v7 = v7
        }
        _ = fi10
        _ = v11
    }
    var v14 uint64 = 233
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
