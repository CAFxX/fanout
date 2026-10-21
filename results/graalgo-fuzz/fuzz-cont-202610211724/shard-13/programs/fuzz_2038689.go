package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int) uint {
    if (false || (("hello" + "foo bar") >= ("foo bar" + "hello"))) {
        var v6 float32 = (0.1 / float32((0.1 + ((float64(0.0) * -2.25) / ((0.1 + 1.5) * (-2.25 + (float64((0.0 * float32((-2.25 + (1.5 - -2.25))))) - 3.14159)))))))
        _ = v6
    }
    var v7 float32 = 3.14159
    var v8 float32 = v7
    for fi9 := 0; fi9 < 10; fi9++ {
        v8 = -2.25
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -465
    var v3 uint64 = 269
    v4 := make(map[int]uint32)
    var v5 int32 = v2
    v10 := []int{-1}
    var v11 int = 42
    if (42 <= -143) {
        var v12 uint = 320
        _ = v12
    } else {
        switch "" {
        case "foo bar":
            v3 += v3
        case "hello":
            for fi13 := 0; fi13 < 5; fi13++ {
                if (true || true) {
                    v11 = v11
                    fmt.Println(v11)
                } else {
                    switch "a" {
                    case "":
                        var v14 int = (fi13 + (v11 + v11))
                        _ = v14
                    case "foo bar":
                        if true {
                            var v15 bool = false
                            _ = v15
                        }
                    case "hello":
                        v3 = 47
                    }
                }
                var v16 uint = 0
                if true {
                    var v17 float32 = ((3.14159 - 1.5) * 1e10)
                    _ = v17
                }
                _ = fi13
                _ = v16
            }
        case "a":
            v5 -= v2
        }
        if true {
            if (42 >= 450) {
                var v18 int32 = v5
                v5 = (^(v5))
                _ = v18
            } else {
                var v19 string = ""
                if (((!(-206 != 1)) && (v11 >= (v11 << ((0) & 63)))) && true) {
                    fmt.Println(v2)
                }
                _ = v19
            }
            v3 = v3
            v5 = -2147483648
        } else {
            var v20 uint8 = 225
            if (!false) {
                var v21 bool = (false && true)
                _ = v21
            }
            fmt.Println(v2, v5, v2)
            _ = v20
        }
    }
    v22 := GenS1{F0: 9223372036854775807}
    var v23 float64 = 1.5
    v3 = 0
    v24 := v1[4294967295]
    v3 -= (v3 << ((42) & 63))
    if (true && true) {
        v2 = (-(v2))
        v24 = v11
        v5 -= v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
