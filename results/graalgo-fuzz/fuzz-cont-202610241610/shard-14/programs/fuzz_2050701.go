package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0() int64 {
    switch "" {
    case "":
        v4 := GenS1{F0: 65535}
        _ = v4
    default:
        for fi5 := 0; fi5 < 5; fi5++ {
            v6 := GenS2{F0: 0}
            var v7 uint16 = 255
            v7 += 255
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    var v8 uint64 = 18446744073709551615
    v9 := GenS2{F0: 128}
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 0
    v3 := make(map[uint32]string)
    fmt.Println(v2)
    v10 := []int{}
    var v11 string = (("hello" + "") + "a")
    var v12 int64 = 42
    var v13 int = -867
    var v14 int64 = 9223372036854775807
    if (3.14159 == 1e10) {
        var v15 float32 = float32(((0.1 * 0.1) + 3.14159))
        _ = v15
    } else {
        v16 := GenS1{F0: 1}
        for fi17 := 0; fi17 < 7; fi17++ {
            v14 = -1
            for fi18 := 0; fi18 < 5; fi18++ {
                fmt.Println(v12, v11, v14)
                switch (v11 + ("\n\t" + "a")) {
                case "\n\t":
                    switch v13 {
                    case 128:
                        if (992 == 248) {
                            v2 = (-(v2))
                            v19 := GenS2{F0: 135}
                            if (false || true) {
                                v2 -= -256
                                var v20 int8 = -128
                                _ = v20
                            }
                            _ = v19
                        }
                    case -1:
                        if false {
                            if true {
                                if (false || true) {
                                    var v21 int64 = v12
                                    _ = v21
                                }
                                fmt.Println(v14, v14, v2)
                                var v22 float64 = (0.0 + (float64((0.1 - (float32(1e10) + float32(3.14159)))) / 1.0))
                                _ = v22
                            }
                        } else {
                            var v23 uint = 833
                            v14 -= genFn0()
                            var v24 int8 = 127
                            _ = v23
                            _ = v24
                        }
                    default:
                        if len(v10) > 0 {
                            v10[0] = int(v2)
                        }
                    }
                case "hello":
                    v25 := GenS2{F0: 128}
                    _ = v25
                }
                _ = fi18
            }
            _ = fi17
        }
        _ = v16
    }
    fmt.Println(v13, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
