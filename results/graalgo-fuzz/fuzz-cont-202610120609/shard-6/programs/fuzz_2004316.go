package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 bool }
func genFn0(p0 int) bool {
    var v4 int32 = 1
    var v5 int8 = 127
    var v6 uint16 = 814
    if (0.0 != float64(0.0)) {
        v7 := GenS1{F0: (v6 >> ((42) & 63))}
        if ("a" > "hello") {
            var v8 int8 = -123
            _ = v8
        }
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (!false)
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int64 = 42
    var v3 int16 = 32767
    v9 := GenS1{F0: 912}
    v10 := []int{-913}
    if false {
        switch ((("\n\t" + ("foo bar" + "foo bar")) + "hello") + "hello") {
        case "hello":
            var v11 uint64 = 566
            _ = v11
        case "foo bar":
            var v12 string = ("a" + ("" + ("hello" + "a")))
            _ = v12
        default:
            for fi13 := 0; fi13 < 4; fi13++ {
                v2 = int64(v1)
                fmt.Println(v3, fi13, v2)
                _ = fi13
            }
        }
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int64 = v2
            switch "foo bar" {
            case "":
                fmt.Println(v3)
            case "hello":
                if (0 >= 42) {
                    v3 += int16(v2)
                } else {
                    v16 := GenS1{F0: 0}
                    if ("foo bar" < "a") {
                        if (((1e10 * 3.14159) / 3.14159) <= (((float32((1e10 - float64(((0.0 - (1e10 + 0.1)) / -2.25)))) * (float32(float64(float32(float64(3.14159)))) - (float32(-2.25) - 0.1))) * -2.25) / (1.5 / float32(3.14159)))) {
                            v3 = (v3 * 32767)
                        }
                    } else {
                        v10 = append(v10, fi14)
                        _ = len(v10)
                        var v17 uint = 42
                        _ = v17
                    }
                    if len(v10) > 0 {
                        v10[0] = int(v15)
                    }
                    _ = v16
                }
            case "\n\t":
                v18 := GenS1{F0: 0}
                _ = v18
            }
            _ = fi14
            _ = v15
        }
        v3 = 32767
    }
    v19 := GenS1{F0: 1}
    v20 := GenS2{F0: "hello"}
    v1 += 376
    v21 := GenS3{F0: v3, F1: true}
    v10 = append(v10, 0)
    var v22 int16 = (v3 * (v3 % ((v3) | 1)))
    var v23 uint32 = (v1 + 1)
    if len(v10) > 0 {
        v10[0] = -197
    }
    fmt.Println(v1, v2, v1)
    var v24 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v9
}
