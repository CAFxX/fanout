package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 string }
func genFn0(p0 int32, p1 int8) uint16 {
    p0 = (p0 & (p0 & p0))
    v7 := []int{32767, 127}
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1() int16 {
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{4294967295, 466, fi8}
        fmt.Println(fi8)
        fmt.Println(fi8)
        _ = fi8
        _ = v9
    }
    var v10 uint64 = 1
    var v11 float64 = 1e10
    for fi12 := 0; fi12 < 13; fi12++ {
        if true {
            fmt.Println(v11, fi12)
        } else {
            v11 += (((float64(1.5) - 0.1) * (v11 - 3.14159)) * float64((0.1 - float32(-2.25))))
            var v13 uint64 = 42
            _ = v13
        }
        v10 -= (v10 - (v10 ^ v10))
        _ = fi12
    }
    v14 := []int{771, -2147483648, -9223372036854775808}
    _ = v10
    _ = v11
    _ = v14
    return -815
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 1
    v3 := make(map[int]int)
    var v4 string = ("a" + ("" + "foo bar"))
    var v5 uint64 = 827
    v6 := make(map[string]int)
    if (true && ((!("hello" == v4)) || false)) {
        var v15 uint8 = 113
        var v16 uint32 = 1
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := []int{fi17, (-(fi17)), int(v2), -2147483648}
            fmt.Println(v5, v5, fi17)
            _ = fi17
            _ = v18
        }
        v2 += (v2 | (^(v2)))
    }
    v19 := GenS1{F0: -1, F1: 0}
    var v20 bool = (false && true)
    v20 = v20
    v21 := GenS1{F0: -34, F1: 255}
    v4 = v4
    var v22 int32 = 42
    v23 := GenS2{F0: 1, F1: (v4 + "\n\t")}
    v24 := GenS2{F0: 0, F1: ("a" + "hello")}
    if false {
        v4 = "a"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
