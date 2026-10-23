package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint; F2 string }
func genFn0(p0 uint64) bool {
    v6 := []int{}
    v7 := GenS1{F0: 1, F1: 322}
    switch 128 {
    case 1:
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 float32 = 1e10
            _ = fi8
            _ = v9
        }
    default:
        v10 := GenS2{F0: 733}
        _ = v10
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    return ((p0 - p0) < (p0 + (^(p0))))
}

func main() {
    v1 := []string{"\n\t", "\n\t", "hello"}
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 uint16 = 1
    var v5 uint64 = 42
    v11 := GenS2{F0: 444}
    var v12 int32 = -2147483648
    var v13 float64 = (((0.0 + (((0.0 + 3.14159) + 0.1) + (0.0 * ((3.14159 / 0.1) + float64(1.5))))) + 0.1) + (0.0 - (float64(0.1) / 3.14159)))
    if (-618 == 9223372036854775807) {
        var v14 bool = (!(574 >= 976))
        _ = v14
    }
    switch ("\n\t" + "\n\t") {
    case "":
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 int = fi15
            if true {
                if false {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        switch 32768 {
                        case -9223372036854775808:
                            var v18 int64 = 42
                            _ = v18
                        default:
                            var v19 bool = (9223372036854775807 == 0)
                            _ = v19
                        }
                        v5 -= 474
                        v20 := GenS2{F0: 2147483647}
                        _ = fi17
                        _ = v20
                    }
                } else {
                    fmt.Println(v12, v12, v13)
                    var v21 int8 = 127
                    _ = v21
                }
            }
            v22 := GenS2{F0: 2147483648}
            _ = fi15
            _ = v16
            _ = v22
        }
    case "hello":
        v23 := v3[0]
        _ = v23
    case "a":
        fmt.Println(v12, v4)
    default:
        if (-1 >= -992) {
            var v24 int = 127
            v4 += v4
            _ = v24
        }
    }
    v5 -= (v5 - v5)
    v25 := GenS1{F0: 304, F1: v4}
    if ((v12 + (v12 * v12)) <= (^(v12))) {
        if ((^(v12)) <= (v12 - -520)) {
            var v26 float32 = -2.25
            v13 = v13
            _ = v26
        } else {
            var v27 int16 = -32768
            switch "" {
            case "":
                for fi28 := 0; fi28 < 1; fi28++ {
                    v4 += v4
                    fmt.Println(v13)
                    _ = fi28
                }
            case "a":
                v13 -= v13
            }
            _ = v27
        }
    } else {
        var v29 int8 = 1
        var v30 int16 = -1
        _ = v29
        _ = v30
    }
    v4 += v4
    _ = len(v3)
    var v31 uint64 = v5
    v5 += (v5 - 42)
    switch -2147483648 {
    case -9223372036854775808:
        v32 := GenS3{F0: v31, F1: 4294967295, F2: ("\n\t" + "a")}
        _ = v32
    }
    var v33 uint8 = 201
    if (true || true) {
        v31 += (v5 | (v5 << ((4294967295) & 63)))
    }
    v34 := GenS3{F0: v31, F1: 613, F2: "a"}
    delete(v2, "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
