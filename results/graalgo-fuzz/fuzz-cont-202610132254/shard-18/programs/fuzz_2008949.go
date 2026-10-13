package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0() int {
    var v7 uint8 = 42
    v8 := []int{4294967295, -140, 176, 449}
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1() uint16 {
    v9 := []int{-9223372036854775808}
    var v10 uint8 = 182
    var v11 uint64 = 42
    var v12 uint16 = 255
    var v13 float32 = 1.5
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v9
    return 255
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint = 0
    var v3 bool = false
    v4 := make(map[string]int)
    var v5 uint = (v2 * 580)
    var v6 int32 = -2147483648
    var v14 int32 = (v6 - (v6 * (v6 | 2147483647)))
    for fi15 := 0; fi15 < 13; fi15++ {
        if (!v3) {
            if (!(!(!((v3 || false) || (!false))))) {
                var v16 uint8 = 42
                v3 = v3
                _ = v16
            } else {
                if v3 {
                    var v17 uint8 = 128
                    v17 = (v17 | (v17 >> ((v2) & 63)))
                    if (true && (153 >= 486)) {
                        v17 += (v17 >> (((^(v2))) & 63))
                        v6 = v14
                        for fi18 := 0; fi18 < 18; fi18++ {
                            v19 := GenS1{F0: fi18, F1: 2147483648}
                            _ = fi18
                            _ = v19
                        }
                    }
                    _ = v17
                }
            }
            var v20 int8 = -128
            var v21 int32 = v6
            _ = v20
            _ = v21
        }
        if (4294967295 <= 24) {
            v3 = (9223372036854775807 == -748)
            if (42 <= 2) {
                fmt.Println(fi15, v6, v2)
                v3 = v3
            }
        }
        _ = fi15
    }
    switch 9223372036854775807 {
    case 0:
        fmt.Println(v5, v5)
    default:
        v1 = v5
    }
    v22 := v4[("hello" + ("a" + ("\n\t" + "")))]
    fmt.Println(v6, v14)
    v5 += 4294967295
    var v23 float32 = 3.14159
    v23 += float32(3.14159)
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 uint16 = 906
        v14 += v6
        v26 := GenS1{F0: (fi24 % (((fi24 % ((int(v2)) | 1))) | 1)), F1: 2}
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int32 = v14
    var v28 float64 = (((1e10 + 3.14159) + float64((float32((((3.14159 - (1e10 * (1e10 + 0.0))) / 3.14159) + (float64(0.0) * 1e10))) + ((v23 - -2.25) * float32(-2.25))))) / 1.5)
    v29 := []int{128}
    v2 += (v2 / ((v2) | 1))
    for fi30 := 0; fi30 < 10; fi30++ {
        if (127 >= 247) {
            v23 -= (0.0 * 3.14159)
        } else {
            var v31 float32 = (1.5 - (1e10 + float32(-2.25)))
            var v32 uint32 = 961
            v29 = append(v29, -2147483648)
            _ = v31
            _ = v32
        }
        v2 = v1
        _ = fi30
    }
    if (42 <= 255) {
        for fi33 := 0; fi33 < 9; fi33++ {
            if (-1 > 194) {
                var v34 string = ("foo bar" + "hello")
                _ = v34
            }
            _ = fi33
        }
        delete(v4, "foo bar")
    } else {
        var v35 uint32 = 65535
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
