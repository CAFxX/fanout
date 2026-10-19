package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int32 }
type GenS2 struct { F0 uint; F1 int8; F2 string }

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]string)
    var v3 int8 = 13
    v4 := []int{}
    var v5 bool = false
    var v6 string = ""
    v7 := GenS1{F0: (!v5), F1: v3, F2: 2147483647}
    v2[271] = v6
    v8 := GenS2{F0: 65535, F1: (v3 * v3), F2: "a"}
    if (-2147483648 < 2147483647) {
        for fi9 := 0; fi9 < 16; fi9++ {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 uint8 = 161
                _ = fi10
                _ = v11
            }
            if (!(true || false)) {
                v12 := GenS2{F0: 31, F1: v3, F2: v6}
                var v13 string = v6
                v6 = v6
                _ = v12
                _ = v13
            } else {
                for fi14 := 0; fi14 < 14; fi14++ {
                    v5 = (!(((0 != -525) || v5) && false))
                    v15 := GenS2{F0: 49, F1: v3, F2: v6}
                    var v16 int64 = 9223372036854775807
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                v1 = v1
            }
            _ = fi9
        }
    }
    if true {
        for fi17 := 0; fi17 < 11; fi17++ {
            v3 = (v3 | -128)
            var v18 uint64 = (^(v1))
            _ = fi17
            _ = v18
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v6, v6, v1)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
