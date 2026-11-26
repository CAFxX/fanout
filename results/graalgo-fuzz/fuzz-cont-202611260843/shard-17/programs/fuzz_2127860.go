package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() bool {
    var v5 float64 = (0.0 + 1.5)
    var v6 uint32 = 65535
    _ = v5
    _ = v6
    return (42 < 2)
}

func main() {
    v1 := []uint32{1}
    var v2 float32 = -2.25
    var v3 uint16 = 42
    v4 := make(map[int]string)
    for fi7 := 0; fi7 < 8; fi7++ {
        if (1.5 != 1e10) {
            var v8 uint16 = (v3 ^ v3)
            v3 -= 1
            _ = v8
        } else {
            _ = len(v4)
            if (!false) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v9 int16 = 0
                v9 -= int16(v3)
                _ = v9
            } else {
                var v10 int32 = -570
                _ = v10
            }
        }
        switch ("foo bar" + "hello") {
        case "hello":
            fmt.Println(fi7, v2, fi7)
        case "a":
            v3 -= (v3 << ((314) & 63))
        case "":
            for fi11 := 0; fi11 < 5; fi11++ {
                if true {
                    switch "foo bar" {
                    case "a":
                        var v12 uint8 = 255
                        _ = v12
                    }
                }
                _ = fi11
            }
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            switch fi13 {
            case 128:
                delete(v4, -792)
                fmt.Println(v2, v2)
            case 222:
                var v14 uint64 = 393
                _ = v14
            case 127:
                v2 -= (v2 + ((v2 - (((v2 * 3.14159) + v2) - 0.1)) * float32(((float64(1.5) + float64(((v2 / (v2 + v2)) + (v2 - v2)))) * 1e10))))
            default:
                fmt.Println(fi13, v2)
            }
            v15 := GenS1{F0: 0}
            _ = fi13
            _ = v15
        }
        _ = fi7
    }
    _ = len(v4)
    _ = len(v4)
    v16 := GenS1{F0: 2147483647}
    fmt.Println(v2)
    var v17 uint32 = 2
    v18 := GenS1{F0: -253}
    var v19 int64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
