package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 float32 = (((float32(((1e10 - ((1.5 - (-2.25 - 1.5)) - float64(-2.25))) / 3.14159)) - (0.1 - 0.1)) / -2.25) / (-2.25 + 0.1))
    var v2 uint16 = 42
    v3 := make(map[string]int)
    var v4 int32 = -2147483648
    if true {
        v2 = (v2 | (v2 % (((v2 & 204)) | 1)))
        v2 -= 65535
    } else {
        if true {
            v2 -= v2
            for fi5 := 0; fi5 < 12; fi5++ {
                var v6 int = (fi5 & -536)
                v2 = 255
                switch "" {
                case "":
                    var v7 int8 = -128
                    _ = v7
                default:
                    v1 += 1.5
                }
                _ = fi5
                _ = v6
            }
            fmt.Println(v2, v1, v1)
        }
        delete(v3, ("" + ("" + "hello")))
        var v8 float64 = -2.25
        _ = v8
    }
    v4 -= v4
    v1 += v1
    v9 := GenS1{F0: 42}
    v10 := v3[(("\n\t" + "") + "foo bar")]
    if false {
        if (1 >= -9223372036854775808) {
            var v11 uint32 = 540
            _ = v11
        }
    }
    v12 := []int{}
    var v13 bool = false
    var v14 string = ("a" + "foo bar")
    v3[(v14 + v14)] = (v10 * 9223372036854775807)
    for fi15 := 0; fi15 < 16; fi15++ {
        switch ((v14 + v14) + v14) {
        case "hello":
            var v16 bool = (v13 && (3.14159 > float64(1e10)))
            _ = v16
        case "foo bar":
            var v17 uint32 = 0
            _ = v17
        }
        _ = fi15
    }
    v1 = v1
    v1 = (3.14159 / v1)
    switch (^(v10)) {
    case 1:
        fmt.Println(v1, v10, v1)
    default:
        var v18 bool = (v13 || false)
        _ = v18
    }
    delete(v3, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
