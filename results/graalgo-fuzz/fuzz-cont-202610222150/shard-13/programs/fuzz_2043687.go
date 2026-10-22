package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint32, p1 bool, p2 string) string {
    v6 := []int{127, 951, 128, -62}
    var v7 string = "foo bar"
    v8 := GenS1{F0: p2, F1: 42, F2: 176}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return ("a" + "")
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]bool)
    var v3 int64 = -535
    var v4 float64 = ((((((((((0.1 + 0.1) + float64(3.14159)) * 1e10) / float64(3.14159)) / (1e10 * (-2.25 - (0.0 * 0.1)))) + ((3.14159 - 1.5) / (((1e10 * (float64((1e10 + (float32(1.5) * 1e10))) * (0.1 + (0.1 / 3.14159)))) - (0.1 - 1.5)) - 3.14159))) / (3.14159 * (((0.1 + 0.0) + 3.14159) * -2.25))) + 1.5) - 1.5) + ((0.0 / 3.14159) - 1.5))
    v5 := []int{}
    if (-2147483648 < -2147483648) {
        var v9 int8 = v1
        _ = v9
    }
    if false {
        v10 := GenS2{F0: 65535}
        _ = v10
    }
    var v11 int64 = v3
    var v12 int16 = 1000
    switch "" {
    case "hello":
        if (!(false || false)) {
            if true {
                if len(v5) > 0 {
                    v5[0] = -9223372036854775808
                }
                v4 += 0.1
            }
            var v13 uint16 = 32
            _ = v13
        } else {
            var v14 string = "a"
            var v15 uint32 = 2147483648
            _ = v14
            _ = v15
        }
    case "":
        for fi16 := 0; fi16 < 16; fi16++ {
            if false {
                fmt.Println(v3)
                v12 = v12
            }
            _ = fi16
        }
    case "a":
        if (104 != 32767) {
            if len(v5) > 0 {
                v5[0] = -9223372036854775808
            }
        }
    }
    switch "hello" {
    case "":
        fmt.Println(v1)
    default:
        v1 -= 42
    }
    v17 := GenS2{F0: 1}
    v18 := GenS3{F0: 674}
    var v19 float32 = -2.25
    v12 += (v12 ^ (v12 | (v12 >> ((4294967295) & 63))))
    fmt.Println(v4, v4, v12)
    var v20 float64 = -2.25
    for fi21 := 0; fi21 < 17; fi21++ {
        v3 = (v11 / (((v11 << ((255) & 63))) | 1))
        var v22 int32 = -540
        delete(v2, fi21)
        _ = fi21
        _ = v22
    }
    if false {
        for fi23 := 0; fi23 < 15; fi23++ {
            if ((-(v11)) > (v11 / ((v11) | 1))) {
                if (!((("\n\t" + "a") + "hello") >= ("a" + ("" + "a")))) {
                    _ = len(v5)
                    fmt.Println(v20, v19)
                }
                var v24 int = 2147483647
                for fi25 := 0; fi25 < 17; fi25++ {
                    var v26 int64 = v3
                    switch (("a" + "a") + "a") {
                    case "":
                        _ = len(v2)
                    default:
                        v27 := GenS1{F0: "a", F1: 42, F2: 240}
                        _ = v27
                    }
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
            _ = fi23
        }
    } else {
        var v28 int = 1
        fmt.Println(v1, v20)
        _ = v28
    }
    v5 = append(v5, 9223372036854775807)
    var v29 float64 = v20
    var v30 int64 = (v3 * (v11 + v11))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
