package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int; F1 int64 }
func genFn0(p0 string) uint32 {
    v5 := GenS1{F0: 0, F1: 131}
    for fi6 := 0; fi6 < 1; fi6++ {
        if false {
            for fi7 := 0; fi7 < 15; fi7++ {
                switch fi7 {
                case -854:
                    p0 = p0
                }
                _ = fi7
            }
        }
        _ = fi6
    }
    var v8 int16 = 1000
    var v9 float32 = ((3.14159 - 1.5) / 1e10)
    if (!(v9 == 0.1)) {
        var v10 int16 = 376
        for fi11 := 0; fi11 < 1; fi11++ {
            if false {
                var v12 int8 = -126
                fmt.Println(fi11)
                _ = v12
            } else {
                if (!(-870 > fi11)) {
                    var v13 int8 = -127
                    p0 = p0
                    _ = v13
                }
                var v14 int = (fi11 % (((fi11 & fi11)) | 1))
                p0 = "hello"
                _ = v14
            }
            _ = fi11
        }
        _ = v10
    }
    _ = p0
    _ = v5
    _ = v8
    _ = v9
    return 784
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 4294967295
    var v3 uint64 = 278
    var v4 int8 = 72
    fmt.Println(v2, v3)
    fmt.Println(v2, v3, v2)
    var v15 uint8 = 42
    var v16 float64 = float64(3.14159)
    if false {
        delete(v1, v2)
        var v17 int8 = v4
        v15 = (v15 % (((v15 & (v15 & v15))) | 1))
        _ = v17
    }
    v18 := GenS2{F0: 656, F1: 0}
    if (int(v3) != 32767) {
        switch ("foo bar" + "") {
        case "\n\t":
            v16 -= (v16 * -2.25)
        case "hello":
            var v19 uint16 = 255
            _ = v19
        default:
            v4 = 100
        }
    }
    fmt.Println(v3)
    v16 -= v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
