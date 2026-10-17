package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 string; F2 string }
type GenS4 struct { F0 int16; F1 int32; F2 uint32 }
func genFn0() uint8 {
    v7 := GenS1{F0: false, F1: false, F2: (-106 < -1)}
    var v8 uint32 = 0
    v8 -= (v8 * 335)
    v8 += (v8 << ((2) & 63))
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]uint32)
    var v3 uint32 = 930
    v4 := make(map[string]string)
    var v5 uint8 = 1
    v6 := make(map[int]int)
    fmt.Println(v5)
    v9 := GenS1{F0: true, F1: true, F2: (!true)}
    var v10 int16 = 1
    var v11 int32 = (^(v1))
    switch 128 {
    case 42:
        if false {
            switch "" {
            case "":
                if (!true) {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        fmt.Println(v1, v5)
                        _ = fi12
                    }
                    fmt.Println(v11, v11)
                } else {
                    var v13 uint32 = v3
                    for fi14 := 0; fi14 < 1; fi14++ {
                        if true {
                            var v15 int16 = (v10 * (v10 >> ((42) & 63)))
                            var v16 bool = ((int32(fi14) != (v11 & (v11 ^ (v11 * (v11 * v11))))) || true)
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS2{F0: -1}
                        for fi18 := 0; fi18 < 9; fi18++ {
                            v19 := GenS1{F0: ((717 < 967) || false), F1: (!(true || false)), F2: ((200 == 0) && (fi14 <= fi18))}
                            if (-16 > 100) {
                                v20 := GenS2{F0: -42}
                                var v21 uint8 = (^(v5))
                                _ = v20
                                _ = v21
                            }
                            _ = fi18
                            _ = v19
                        }
                        _ = fi14
                        _ = v17
                    }
                    v11 += (v1 + -801)
                    _ = v13
                }
            case "a":
                v22 := GenS3{F0: ((true && false) && false), F1: "foo bar", F2: ("" + "hello")}
                _ = v22
            case "foo bar":
                delete(v2, 32768)
            }
            var v23 int32 = v11
            v11 = (^(v23))
            _ = v23
        }
    case -597:
        var v24 string = "hello"
        _ = v24
    case -638:
        var v25 uint64 = 1
        _ = v25
    default:
        v11 = -1
    }
    v26 := GenS4{F0: (v10 << ((42) & 63)), F1: v11, F2: (v3 + 2147483647)}
    var v27 float32 = ((3.14159 / -2.25) / 3.14159)
    var v28 string = ("a" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
