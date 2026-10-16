package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
type GenS4 struct { F0 uint; F1 int32; F2 uint64 }
func genFn0(p0 string, p1 uint32, p2 uint16) int {
    p1 += (p1 - (p1 - 1))
    p0 = "\n\t"
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: 4294967295, F1: int8(p1), F2: p1}
        v6 := []int{128}
        for fi7 := 0; fi7 < 13; fi7++ {
            p0 = p0
            v8 := GenS2{F0: fi4, F1: 42}
            var v9 uint8 = 2
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 uint64) int {
    p0 -= 18446744073709551615
    var v10 uint = 180
    _ = p0
    _ = v10
    return 345
}

func main() {
    var v1 string = ("\n\t" + "")
    var v2 uint32 = 882
    var v3 int64 = 9223372036854775807
    v11 := []int{32768, 2147483647, 2147483647}
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS2{F0: (-(fi12)), F1: 0}
        _ = fi12
        _ = v13
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v14 uint = 223
    if ((v14 ^ (v14 >> ((v14) & 63))) > v14) {
        var v15 float64 = float64(-2.25)
        _ = v15
    }
    switch -1 {
    case 42:
        if (65535 == 65535) {
            v1 = ("foo bar" + "hello")
        } else {
            var v16 bool = false
            v11 = append(v11, genFn0(("foo bar" + "hello"), (v2 & v2), 42))
            _ = v16
        }
    default:
        v17 := GenS3{F0: -9223372036854775808, F1: -107, F2: 1}
        _ = v17
    }
    var v18 uint16 = 255
    v19 := GenS3{F0: 125, F1: 127, F2: -2147483648}
    v20 := GenS3{F0: 128, F1: 94, F2: -158}
    v14 -= 189
    v18 = v18
    switch ("\n\t" + (v1 + (((v1 + ("hello" + v1)) + (((v1 + ((v1 + (("a" + "\n\t") + "foo bar")) + (v1 + ""))) + "a") + "a")) + v1))) {
    case "hello":
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 int64 = (v3 >> ((v14) & 63))
            fmt.Println(v14)
            if (65535 >= 65535) {
                v1 = v1
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 int16 = 1000
                    v25 := GenS4{F0: (v14 << ((v14) & 63)), F1: 42, F2: 243}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            } else {
                v26 := GenS1{F0: 353, F1: 0, F2: (^(v2))}
                v11 = append(v11, 127)
                _ = v26
            }
            _ = fi21
            _ = v22
        }
    }
    _ = len(v11)
    var v27 float32 = (-2.25 + (float32(3.14159) - 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
}
