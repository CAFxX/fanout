package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    v3 := make(map[int]uint32)
    v4 := []uint32{}
    var v5 uint64 = 18446744073709551615
    delete(v1, -1)
    v2 = (1.5 <= ((3.14159 - (float64((0.0 * -2.25)) - 3.14159)) + ((0.1 - 1.5) - 1.5)))
    v6 := GenS1{F0: "foo bar"}
    for fi7 := 0; fi7 < 19; fi7++ {
        if len(v4) > 0 {
            v4[0] = 2147483648
        }
        v8 := GenS2{F0: 0}
        _ = fi7
        _ = v8
    }
    v5 = 18446744073709551615
    if (v2 || false) {
        v9 := GenS1{F0: "\n\t"}
        _ = v9
    }
    if v2 {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 float32 = 0.0
            if v2 {
                if true {
                    _ = len(v1)
                }
            }
            fmt.Println(v5, v11, v11)
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: "\n\t"}
        _ = v12
    } else {
        fmt.Println(v5)
        for fi13 := 0; fi13 < 12; fi13++ {
            switch (^(fi13)) {
            case -497:
                v5 += (v5 - (v5 << ((42) & 63)))
            case 4294967295:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 0:
                var v14 int16 = int16(fi13)
                _ = v14
            }
            _ = fi13
        }
    }
    v5 = v5
    var v15 uint64 = v5
    v16 := GenS3{F0: 284}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
