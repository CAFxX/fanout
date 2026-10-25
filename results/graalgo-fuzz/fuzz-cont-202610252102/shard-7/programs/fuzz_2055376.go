package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0(p0 int16, p1 bool, p2 uint16) int {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 99, F1: 42}
        _ = fi4
        _ = v5
    }
    fmt.Println(p1, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}
func genFn1(p0 uint) bool {
    p0 = p0
    p0 = (p0 << (((p0 / (((^(p0))) | 1))) & 63))
    _ = p0
    return (false || (-1 < -59))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    var v3 uint64 = 163
    v3 -= (v3 << ((255) & 63))
    switch genFn0(0, false, 2) {
    case 128:
        if (1 == 65535) {
            var v6 uint = 600
            _ = v6
        }
    case 127:
        if true {
            v3 = 4294967295
        }
    }
    v3 += (v3 >> ((255) & 63))
    v7 := []int{genFn0(-1, true, 255), 89, 4294967295}
    var v8 string = "\n\t"
    var v9 int64 = 152
    v10 := GenS1{F0: -1, F1: 42}
    var v11 string = ("\n\t" + ("a" + (v8 + (("hello" + "hello") + v8))))
    for fi12 := 0; fi12 < 12; fi12++ {
        if false {
            v9 -= 0
            if false {
                var v13 bool = false
                _ = v13
            }
            v8 = "\n\t"
        } else {
            v14 := GenS1{F0: -1, F1: 127}
            v15 := GenS1{F0: 63, F1: 42}
            _ = v14
            _ = v15
        }
        switch 680 {
        case 42:
            v7 = append(v7, (^(fi12)))
        case 179:
            switch 128 {
            case -9223372036854775808:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case 127:
                v9 -= (v9 - v9)
            case 2147483647:
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 float64 = -2.25
                    _ = fi16
                    _ = v17
                }
            default:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
        case 32767:
            var v18 uint32 = 0
            _ = v18
        default:
            var v19 uint8 = 237
            _ = v19
        }
        fmt.Println(v11)
        _ = fi12
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v11)
    v1[32768] = 1
    v8 = "hello"
    v20 := GenS1{F0: 127, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
