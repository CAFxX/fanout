package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    var v1 uint = 4294967295
    v2 := []uint32{0}
    var v3 float64 = float64(-2.25)
    fmt.Println(v1, v1)
    var v4 int64 = 0
    if (2147483647 < -2147483648) {
        if (!(("" <= "foo bar") || false)) {
            for fi5 := 0; fi5 < 2; fi5++ {
                for fi6 := 0; fi6 < 16; fi6++ {
                    v2 = append(v2, 841)
                    fmt.Println(v1)
                    _ = fi6
                }
                _ = fi5
            }
        }
    } else {
        if (627 == 1) {
            v7 := GenS1{F0: 65535, F1: ("foo bar" + ((("hello" + "hello") + "\n\t") + "")), F2: 45}
            if (true && (v1 < (v1 >> ((707) & 63)))) {
                var v8 int32 = -1
                _ = v8
            }
            _ = v7
        }
    }
    var v9 int8 = 103
    v10 := GenS2{F0: 0, F1: v4}
    fmt.Println(v4, v3, v1)
    fmt.Println(v4)
    v4 = (v4 >> ((v1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
