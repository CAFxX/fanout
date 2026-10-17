package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{32768, -9223372036854775808}
    var v3 uint32 = 4294967295
    var v4 uint64 = 0
    switch "hello" {
    case "foo bar":
        var v5 uint = 965
        _ = v5
    case "\n\t":
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: 128}
            switch (fi6 / ((-9223372036854775808) | 1)) {
            case 389:
                var v8 uint16 = 32767
                _ = v8
            case 127:
                if ("foo bar" > "a") {
                    delete(v1, v3)
                } else {
                    var v9 uint8 = 27
                    var v10 int = -1
                    _ = v10
                    _ = v9
                }
            }
            _ = fi6
            _ = v7
        }
    case "a":
        v3 = 4294967295
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int = fi11
        v12 += (fi11 % ((fi11) | 1))
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS1{F0: 127}
            var v16 int8 = 122
            _ = fi14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := GenS2{F0: (v4 | v4), F1: -1, F2: (fi13 - (fi13 + (fi17 | (fi17 >> ((2) & 63)))))}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            v4 = (v4 * (v4 * (v4 + (^(v4)))))
            v3 += v3
            var v20 int8 = 0
            _ = fi19
            _ = v20
        }
        _ = fi13
    }
    var v21 float64 = 0.1
    var v22 float64 = 0.0
    v23 := GenS1{F0: 212}
    if (42 <= 793) {
        var v24 float32 = 1e10
        v25 := v1[(v3 >> ((1) & 63))]
        _ = v24
        _ = v25
    }
    var v26 int64 = -9223372036854775808
    v27 := v1[(v3 / ((v3) | 1))]
    delete(v1, v3)
    v2 = append(v2, (v27 ^ -310))
    var v28 int64 = 1
    var v29 float32 = 3.14159
    var v30 int8 = -1
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
