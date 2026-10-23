package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int16) string {
    v7 := GenS1{F0: 203, F1: 2}
    fmt.Println(p0)
    var v8 float32 = 3.14159
    p0 = p0
    var v9 int32 = 42
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 int16 = 32767
    var v2 float64 = (1e10 + 1e10)
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    v6 := []uint32{784, 4294967295}
    v6 = append(v6, 2147483647)
    fmt.Println(v2, v2)
    var v10 uint16 = 308
    v2 -= 1.5
    v11 := GenS1{F0: 4294967295, F1: 2}
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: 4294967295, F1: 18446744073709551615}
        if (357 < 451) {
            if true {
                switch (("hello" + "a") + "a") {
                case "":
                    switch -2147483648 {
                    case 0:
                        for fi14 := 0; fi14 < 13; fi14++ {
                            var v15 uint64 = 42
                            v6 = append(v6, 4294967295)
                            var v16 uint = 268
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    default:
                        v6 = append(v6, 717)
                    }
                case "a":
                    var v17 uint32 = 1
                    _ = v17
                }
            }
            var v18 uint = 4294967295
            if false {
                if len(v6) > 0 {
                    v6[0] = 2147483648
                }
                var v19 float64 = v2
                _ = v19
            }
            _ = v18
        }
        _ = fi12
        _ = v13
    }
    var v20 uint = 65535
    switch int(v10) {
    case -610:
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 uint32 = 2147483647
            v23 := GenS1{F0: 2147483647, F1: 32}
            if false {
                for fi24 := 0; fi24 < 15; fi24++ {
                    v1 = v1
                    fmt.Println(v22, v10)
                    _ = fi24
                }
            }
            _ = fi21
            _ = v22
            _ = v23
        }
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println(v1, v10)
    for fi25 := 0; fi25 < 16; fi25++ {
        _ = len(v6)
        v6 = append(v6, 42)
        _ = fi25
    }
    if len(v6) > 0 {
        v6[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
