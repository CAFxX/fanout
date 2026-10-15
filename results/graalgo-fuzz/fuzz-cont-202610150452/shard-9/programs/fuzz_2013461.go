package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
func genFn0() int {
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 uint32 = 2
        _ = fi4
        _ = v5
    }
    switch "" {
    case "hello":
        var v6 string = "foo bar"
        _ = v6
    case "a":
        if (!true) {
            var v7 int = 1
            if (18446744073709551615 == 18446744073709551615) {
                v8 := GenS1{F0: 807, F1: true}
                for fi9 := 0; fi9 < 2; fi9++ {
                    fmt.Println(fi9)
                    _ = fi9
                }
                _ = v8
            } else {
                var v10 bool = (false && true)
                v11 := GenS2{F0: (v7 + (^(v7))), F1: 0}
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: 2, F1: true}
            _ = v12
            _ = v7
        } else {
            var v13 uint16 = 567
            v13 -= v13
            _ = v13
        }
    case "":
        v14 := []int{-2147483648, 135, -2147483648, 127}
        _ = v14
    default:
        var v15 uint = 65535
        _ = v15
    }
    v16 := []int{127, 397, 987}
    _ = v16
    return 424
}
func genFn1(p0 uint, p1 uint8) int32 {
    v17 := GenS2{F0: 1, F1: 42}
    p0 -= 1
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := GenS3{F0: -1}
        for fi20 := 0; fi20 < 17; fi20++ {
            if (1.5 == 0.0) {
                fmt.Println(p1, p1)
                var v21 uint16 = 0
                _ = v21
            } else {
                var v22 bool = (("hello" + "") < ("hello" + "hello"))
                v22 = v22
                _ = v22
            }
            v23 := []int{int(p0), fi20, 32768, fi18}
            _ = fi20
            _ = v23
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(p1, p1)
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 float32 = -2.25
        var v26 uint = 18446744073709551615
        _ = fi24
        _ = v25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v17
    return 1
}

func main() {
    var v1 float32 = 1.5
    var v2 bool = false
    v3 := make(map[string]bool)
    delete(v3, ("a" + ""))
    var v27 uint64 = 4294967295
    var v28 int8 = 92
    v1 = float32(1.5)
    fmt.Println(v27, v1)
    v29 := GenS4{F0: 878}
    v1 += (v1 * (((v1 - v1) + (((((float32(3.14159) * (3.14159 / v1)) * v1) + (3.14159 * v1)) + 0.0) + v1)) + (v1 * (-2.25 - ((v1 - v1) * v1)))))
    delete(v3, "")
    fmt.Println(v1, v1, v1)
    v2 = (true && true)
    v27 -= 42
    v30 := GenS3{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
