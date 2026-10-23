package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int16; F1 bool; F2 uint32 }

func main() {
    var v1 uint = 42
    v2 := []string{"a"}
    var v3 int = 0
    var v4 float32 = 1.5
    if ((!false) && (!false)) {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 uint = (v1 * 42)
            if false {
                var v7 uint32 = 65535
                v7 += v7
                if true {
                    for fi8 := 0; fi8 < 6; fi8++ {
                        switch ("\n\t" + "hello") {
                        case "a":
                            if len(v2) > 0 {
                                v2[0] = (("a" + "\n\t") + "foo bar")
                            }
                        case "\n\t":
                            v9 := GenS1{F0: ((640 != 1) || (-2147483648 != -540)), F1: -1}
                            _ = v9
                        }
                        _ = fi8
                    }
                    fmt.Println(v1, v4, fi5)
                } else {
                    fmt.Println(v6, fi5)
                    v10 := GenS2{F0: 1, F1: 59}
                    var v11 int8 = 43
                    _ = v10
                    _ = v11
                }
                _ = v7
            }
            v12 := GenS3{F0: -441, F1: false, F2: 592}
            _ = fi5
            _ = v12
            _ = v6
        }
    }
    var v13 float32 = float32(1e10)
    for fi14 := 0; fi14 < 7; fi14++ {
        v13 = (v4 - (1.5 + (float32((1e10 / 1.0)) - float32((0.1 - (1.5 - (float64(3.14159) / (float64(((v4 / v4) * v4)) * (float64(float32(3.14159)) + 1e10)))))))))
        _ = fi14
    }
    var v15 uint64 = 388
    v1 = 18446744073709551615
    var v16 float64 = (1.5 + 0.0)
    var v17 uint64 = (v15 << ((v1) & 63))
    var v18 int16 = 1
    if (v13 > (v13 * 3.14159)) {
        for fi19 := 0; fi19 < 19; fi19++ {
            v3 = (fi19 >> ((v1) & 63))
            _ = fi19
        }
        var v20 int32 = -1
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
