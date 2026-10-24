package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint8 }

func main() {
    v1 := []string{}
    var v2 uint8 = 127
    v3 := make(map[uint32]uint32)
    v2 = v2
    if true {
        v4 := v3[65535]
        var v5 float64 = 0.1
        var v6 string = ""
        _ = v4
        _ = v5
        _ = v6
    } else {
        _ = len(v3)
        var v7 uint8 = (v2 + v2)
        _ = v7
    }
    v8 := GenS1{F0: false, F1: 684, F2: v2}
    v9 := GenS1{F0: (659 >= 900), F1: 142, F2: (v2 >> ((747) & 63))}
    v10 := GenS1{F0: true, F1: 1, F2: (v2 * v2)}
    delete(v3, 0)
    var v11 uint32 = 949
    v11 = v11
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v11 += 4294967295
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint64 = 2
            v15 := GenS1{F0: (!true), F1: v11, F2: (v2 / (((v2 & 42)) | 1))}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    if true {
        for fi16 := 0; fi16 < 16; fi16++ {
            if true {
                var v17 uint16 = 32767
                var v18 int = (fi16 | 4294967295)
                _ = v17
                _ = v18
            } else {
                if (3.14159 > (1.5 / 1.0)) {
                    var v19 int32 = 0
                    _ = v19
                }
                var v20 string = ("a" + "")
                var v21 string = "hello"
                _ = v20
                _ = v21
            }
            _ = fi16
        }
    }
    delete(v3, v11)
    var v22 bool = false
    var v23 int = 2147483647
    v11 += (^(v11))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v8
    _ = v9
}
