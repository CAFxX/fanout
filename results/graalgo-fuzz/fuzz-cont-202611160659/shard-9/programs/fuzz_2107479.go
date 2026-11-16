package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 uint }
func genFn0() int32 {
    switch -9223372036854775808 {
    case 127:
        var v5 string = (("a" + "") + (("hello" + "foo bar") + "a"))
        _ = v5
    case -9223372036854775808:
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 int8 = 1
            _ = fi6
            _ = v7
        }
    }
    if (2 <= 2) {
        v8 := []int{9223372036854775807, 1, 546, 167}
        var v9 string = "hello"
        var v10 float64 = 0.0
        _ = v10
        _ = v8
        _ = v9
    } else {
        var v11 int16 = 425
        for fi12 := 0; fi12 < 14; fi12++ {
            v11 -= -360
            v13 := []int{}
            var v14 string = ((("" + "") + ("" + "a")) + "")
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 bool = true
            for fi17 := 0; fi17 < 7; fi17++ {
                if v16 {
                    v11 += v11
                } else {
                    v18 := []int{fi17, fi17}
                    _ = v18
                }
                var v19 int32 = 1
                _ = fi17
                _ = v19
            }
            v20 := []int{-61, 223}
            _ = fi15
            _ = v16
            _ = v20
        }
        _ = v11
    }
    var v21 int8 = 127
    _ = v21
    return 1
}
func genFn1() int {
    var v22 uint64 = 0
    v23 := GenS1{F0: (42 > 9223372036854775807), F1: -1, F2: 59}
    var v24 uint8 = 42
    var v25 int32 = 591
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    return -2147483648
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]bool)
    var v3 uint8 = 1
    var v4 uint = 435
    for fi26 := 0; fi26 < 14; fi26++ {
        fmt.Println(fi26, fi26, v4)
        _ = fi26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 uint = (v4 % ((4294967295) | 1))
        for fi29 := 0; fi29 < 11; fi29++ {
            v4 -= 18446744073709551615
            var v30 bool = true
            _ = fi29
            _ = v30
        }
        fmt.Println(fi27, v1, v28)
        _ = fi27
        _ = v28
    }
    var v31 uint = 65535
    var v32 int16 = 825
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 string = ("foo bar" + "")
        v4 += (v31 + v31)
        v1 = (v1 - v1)
        _ = fi33
        _ = v34
    }
    v35 := GenS2{F0: false, F1: 2, F2: ("foo bar" + ("foo bar" + "\n\t"))}
    var v36 int16 = 32767
    v37 := GenS3{F0: v4}
    var v38 uint32 = 1
    var v39 int8 = 113
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
