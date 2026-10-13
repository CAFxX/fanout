package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 255
    var v3 uint = 18446744073709551615
    var v4 uint = (^(v3))
    fmt.Println(v2)
    if (1 <= 65535) {
        switch "\n\t" {
        case "a":
            v5 := GenS1{F0: -517}
            _ = v5
        case "\n\t":
            if false {
                v3 += v4
                if true {
                    v2 -= v2
                    var v6 uint64 = 42
                    switch 127 {
                    case 32767:
                        var v7 float64 = 1.5
                        _ = v7
                    case 127:
                        v2 = 42
                    default:
                        for fi8 := 0; fi8 < 1; fi8++ {
                            v9 := GenS1{F0: int16(v6)}
                            switch "\n\t" {
                            case "":
                                v10 := GenS1{F0: -958}
                                _ = v10
                            case "\n\t":
                                switch (("\n\t" + "\n\t") + "a") {
                                case "hello":
                                    var v11 uint16 = v2
                                    _ = v11
                                case "a":
                                    v12 := []int{fi8, fi8}
                                    _ = v12
                                default:
                                    var v13 float32 = (0.0 + (-2.25 / (((float32(1.5) + 0.1) * ((0.0 * (float32((-2.25 - (3.14159 * float64(0.0)))) / 3.14159)) / 0.1)) + (float32(float64(0.1)) - 0.0))))
                                    _ = v13
                                }
                            }
                            for fi14 := 0; fi14 < 3; fi14++ {
                                delete(v1, "\n\t")
                                _ = fi14
                            }
                            _ = fi8
                            _ = v9
                        }
                    }
                    _ = v6
                }
                v15 := GenS2{F0: 41, F1: 0}
                _ = v15
            }
        }
        v1["hello"] = (!(-128 < 127))
        var v16 float32 = 1.5
        _ = v16
    }
    fmt.Println(v3, v4)
    v4 = (v3 >> ((v4) & 63))
    var v17 int8 = 127
    var v18 bool = false
    if ((("foo bar" + "hello") + "hello") != ("" + ("hello" + "foo bar"))) {
        v19 := GenS1{F0: 913}
        _ = v19
    }
    v2 = 112
    v20 := []int{1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
