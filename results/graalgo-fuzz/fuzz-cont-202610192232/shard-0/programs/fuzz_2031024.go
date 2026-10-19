package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint16 = 0
    var v2 int64 = 283
    v3 := []uint32{873, 492, 2}
    var v4 int8 = 100
    v5 := []string{"hello"}
    var v6 string = ""
    fmt.Println(v1, v6)
    for fi7 := 0; fi7 < 16; fi7++ {
        if (!(true || true)) {
            fmt.Println(v6, v4)
        }
        v8 := GenS1{F0: 127}
        for fi9 := 0; fi9 < 9; fi9++ {
            if true {
                var v10 string = v6
                _ = v10
            } else {
                v4 -= -1
            }
            if len(v3) > 0 {
                v3[0] = 1
            }
            var v11 int32 = 2147483647
            var v12 uint8 = 0
            var v13 int = fi9
            _ = fi9
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    var v14 uint = 4294967295
    fmt.Println(v1, v1, v2)
    v2 += (-(v2))
    var v15 int32 = -2147483648
    v6 = ("\n\t" + v6)
    var v16 int64 = int64(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
