package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 uint16, p1 uint) int32 {
    v4 := GenS1{F0: -9223372036854775808, F1: 921, F2: -1}
    p0 = (p0 & (p0 >> ((p1) & 63)))
    _ = p0
    _ = p1
    _ = v4
    return 2147483647
}
func genFn1() string {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint32 = 930
        var v7 string = "a"
        v6 -= v6
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint = 1
    _ = v8
    return ("a" + ("\n\t" + "hello"))
}

func main() {
    var v1 float64 = (-2.25 * 1.5)
    var v2 int8 = 42
    v3 := make(map[string]bool)
    v2 = v2
    v2 = 1
    if true {
        v9 := v3["hello"]
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(v2)
            switch "a" {
            case "hello":
                v11 := []int{}
                _ = v11
            case "foo bar":
                var v12 uint = 42
                _ = v12
            }
            _ = fi10
        }
        _ = v9
    } else {
        v13 := GenS2{F0: 78}
        v14 := v3[("hello" + "")]
        v15 := GenS2{F0: 22}
        _ = v13
        _ = v14
        _ = v15
    }
    v1 = 3.14159
    v1 -= (float64((0.0 * (3.14159 * 1.5))) + ((((float64(3.14159) + float64((float32(v1) / (3.14159 - (float32((v1 - (1e10 / ((v1 + v1) * v1)))) / 1e10))))) / v1) + v1) + v1))
    var v16 int = 466
    v17 := []int{(^(v16)), (v16 % ((v16) | 1)), (-(v16))}
    if (0 > 127) {
        var v18 int64 = 42
        _ = v18
    } else {
        var v19 bool = true
        _ = v19
    }
    if (v16 < (v16 | 1)) {
        v16 += v16
        if true {
            var v20 string = ""
            var v21 float64 = float64(float32(0.1))
            v17 = append(v17, v16)
            _ = v20
            _ = v21
        }
    }
    if false {
        v22 := GenS2{F0: 61}
        fmt.Println(v16, v1, v2)
        if false {
            v16 += v16
        } else {
            v1 = 1.5
        }
        _ = v22
    }
    v1 -= (-2.25 / (v1 * 0.1))
    v23 := GenS3{F0: (v16 | (v16 | (v16 | (v16 << ((4294967295) & 63))))), F1: 4294967295}
    var v24 float64 = (v1 * v1)
    v25 := GenS3{F0: (v16 - v16), F1: 4294967295}
    v24 = ((v1 + ((v24 / v1) - float64(((1e10 + 3.14159) * (0.0 / (1e10 - 1.5)))))) * float64((1e10 + ((float32(v24) / 3.14159) / (1e10 / (1.5 - float32(1e10)))))))
    for fi26 := 0; fi26 < 14; fi26++ {
        v1 += v1
        _ = fi26
    }
    var v27 string = ("foo bar" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
}
