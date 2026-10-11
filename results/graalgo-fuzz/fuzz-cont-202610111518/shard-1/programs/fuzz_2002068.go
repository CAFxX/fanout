package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
func genFn0(p0 uint32, p1 bool) int {
    switch "hello" {
    case "":
        if (670 == 65535) {
            if p1 {
                for fi7 := 0; fi7 < 19; fi7++ {
                    var v8 int16 = 42
                    _ = fi7
                    _ = v8
                }
                p1 = (93 > 0)
            } else {
                p1 = (128 >= 2147483647)
                fmt.Println(p0, p0)
                var v9 uint32 = 1
                _ = v9
            }
            if (!(65 > -1)) {
                if false {
                    var v10 float32 = 0.0
                    v11 := []int{-754, 1, 1, -622}
                    _ = v10
                    _ = v11
                }
            } else {
                v12 := []int{127}
                var v13 uint32 = (p0 - 65535)
                _ = v12
                _ = v13
            }
        }
    case "hello":
        v14 := GenS1{F0: 2, F1: -984, F2: ("hello" + ("foo bar" + "a"))}
        _ = v14
    case "a":
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 int = (fi15 ^ (fi15 ^ (fi15 ^ fi15)))
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := []int{-2147483648, fi17, -8}
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
    }
    p0 += (p0 | 463)
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 uint16) int32 {
    if true {
        fmt.Println(p0)
    }
    var v19 float64 = float64(0.0)
    v20 := []int{-926}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    _ = p0
    _ = v19
    _ = v20
    return -813
}

func main() {
    var v1 int8 = 42
    v2 := make(map[string]bool)
    v3 := []int{2147483647, 32768, 2147483647, 32768}
    v4 := make(map[uint32]string)
    var v5 float32 = float32(float64((0.0 - ((1.5 / 1.5) / -2.25))))
    var v6 int16 = 1000
    v5 += (v5 / v5)
    fmt.Println(v1, v6)
    v6 -= 1000
    fmt.Println(v5, v1, v5)
    var v21 uint32 = 1
    v22 := GenS1{F0: 534, F1: 961, F2: (("foo bar" + "hello") + ("\n\t" + "\n\t"))}
    var v23 float64 = -2.25
    var v24 uint = uint(v21)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
