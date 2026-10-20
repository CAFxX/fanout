package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32 }
func genFn0() string {
    var v7 uint = 255
    if (-9223372036854775808 == -9223372036854775808) {
        v8 := GenS1{F0: 105, F1: 1}
        _ = v8
    }
    v9 := GenS2{F0: 2147483648}
    v7 = v7
    _ = v7
    _ = v9
    return (("hello" + "\n\t") + "a")
}
func genFn1(p0 uint16) bool {
    v10 := []int{-1, 32768, 32768, 32767}
    var v11 float32 = 0.1
    v12 := GenS2{F0: 4294967295}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return ((!false) && (p0 < p0))
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 9223372036854775807
    v3 := []string{"hello"}
    var v4 uint16 = 0
    v5 := make(map[string]string)
    v6 := []int{v2, (v2 % ((-40) | 1)), v2, (v2 | (v2 % (((v2 / ((v2) | 1))) | 1)))}
    v4 -= 42
    fmt.Println(v4)
    v2 = (-(v2))
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := v5["foo bar"]
        for fi15 := 0; fi15 < 6; fi15++ {
            v2 = -2147483648
            if (!(-944 != 2147483647)) {
                delete(v5, v14)
                var v16 uint16 = v4
                _ = v16
            }
            v3 = append(v3, v14)
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    var v17 bool = (!(4294967295 >= v2))
    v18 := GenS2{F0: 2147483647}
    fmt.Println(v17, v4, v4)
    if ((0.0 / 0.1) <= 3.14159) {
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := GenS2{F0: 976}
            for fi21 := 0; fi21 < 15; fi21++ {
                v22 := GenS2{F0: 2}
                _ = fi21
                _ = v22
            }
            var v23 uint8 = 255
            _ = fi19
            _ = v20
            _ = v23
        }
        var v24 int64 = 9223372036854775807
        _ = v24
    }
    switch v2 {
    case -2147483648:
        var v25 int16 = 32767
        _ = v25
    case -749:
        var v26 string = "a"
        _ = v26
    case 2147483647:
        v2 += v2
    }
    var v27 uint8 = 2
    var v28 uint16 = 1
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
