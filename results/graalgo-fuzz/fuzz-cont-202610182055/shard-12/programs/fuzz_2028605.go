package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    v3 := []int{}
    switch ("foo bar" + "a") {
    case "foo bar":
        if false {
            if ((0 == -1) && (255 == 2)) {
                for fi4 := 0; fi4 < 5; fi4++ {
                    if true {
                        var v5 int64 = 42
                        _ = v5
                    }
                    _ = fi4
                }
            } else {
                if (0 < 100) {
                    var v6 float64 = 1.5
                    v7 := GenS1{F0: 255}
                    v6 -= ((v6 - v6) - (v6 * float64((0.0 - 1.5))))
                    _ = v6
                    _ = v7
                } else {
                    v8 := GenS1{F0: 776}
                    var v9 float32 = 0.1
                    _ = v8
                    _ = v9
                }
            }
            if true {
                var v10 int64 = 565
                v1[-227] = (("a" + ("" + "a")) + "hello")
                _ = v10
            }
            v11 := GenS1{F0: 688}
            _ = v11
        } else {
            var v12 int = -9223372036854775808
            for fi13 := 0; fi13 < 15; fi13++ {
                fmt.Println(v12)
                switch ("foo bar" + "") {
                case "\n\t":
                    _ = len(v3)
                default:
                    v12 -= 798
                }
                _ = fi13
            }
            _ = v12
        }
    default:
        delete(v2, 2147483648)
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: 2}
        var v16 float32 = ((3.14159 + 0.1) * ((0.1 * (-2.25 / 1.0)) + (((float32(1.5) / (-2.25 - 0.1)) + 3.14159) / 1.5)))
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int16 = 1
    var v18 int16 = v17
    if (4294967295 < 32768) {
        v17 -= v17
        if false {
            v19 := GenS2{F0: (v17 - (v18 & -1))}
            v18 = 1
            _ = v19
        }
    }
    if true {
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS3{F0: 0}
            _ = fi20
            _ = v21
        }
        v22 := GenS1{F0: 42}
        fmt.Println(v18, v17)
        _ = v22
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v17 = v18
    v17 = 1000
    v17 = v17
    v23 := v1[2147483647]
    var v24 bool = (true || false)
    v25 := GenS2{F0: (v17 & (v18 & 1000))}
    if v24 {
        if (v24 || ((469 == -902) && (v24 && v24))) {
            fmt.Println(v24, v17, v24)
        }
        v3 = append(v3, 42)
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        v18 += v17
        _ = fi26
    }
    switch -637 {
    case 127:
        if (1 >= 18446744073709551615) {
            fmt.Println(v23)
        }
    case -587:
        for fi27 := 0; fi27 < 6; fi27++ {
            v24 = v24
            var v28 int32 = -110
            _ = fi27
            _ = v28
        }
    }
    v23 = v23
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
