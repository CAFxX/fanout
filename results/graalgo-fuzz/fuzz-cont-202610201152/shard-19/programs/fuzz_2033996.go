package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 bool; F1 bool; F2 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint32, p1 int8, p2 uint) uint {
    for fi5 := 0; fi5 < 11; fi5++ {
        if (!(false || (!(false || (true && (!true)))))) {
            v6 := GenS1{F0: 2147483647, F1: 119, F2: fi5}
            var v7 uint = p2
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 uint32 = (p0 / ((65535) | 1))
                v10 := []int{(-(fi8)), -2147483648}
                p1 += p1
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    switch -678 {
    case 1:
        var v11 float64 = 1e10
        _ = v11
    default:
        switch int(p2) {
        case -402:
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 int16 = 1000
                if ((false && false) || false) {
                    v14 := []int{(fi12 - -636), -658, int(p1), fi12}
                    _ = v14
                } else {
                    v15 := GenS1{F0: 2147483647, F1: 8, F2: fi12}
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
        default:
            v16 := GenS1{F0: -136, F1: 42, F2: -440}
            _ = v16
        }
    }
    p0 = (p0 | p0)
    _ = p0
    _ = p1
    _ = p2
    return uint(p0)
}

func main() {
    var v1 int8 = 127
    var v2 bool = (-2.25 == (3.14159 + 1e10))
    v3 := make(map[string]int)
    v4 := []string{"a"}
    var v17 uint = 401
    v2 = (-33 == v1)
    for fi18 := 0; fi18 < 14; fi18++ {
        v19 := GenS2{F0: 235, F1: 47}
        _ = fi18
        _ = v19
    }
    _ = len(v4)
    v20 := GenS3{F0: 999, F1: 255, F2: 560}
    var v21 uint = v17
    v17 = 1
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v22 string = ""
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v21 = v21
    var v23 int32 = 923
    v24 := GenS4{F0: (!v2), F1: v2, F2: v1}
    for fi25 := 0; fi25 < 12; fi25++ {
        fmt.Println(v2, v17)
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS5{F0: (v1 + (v1 << ((v17) & 63)))}
            var v28 uint = genFn0(2, -98, (v17 ^ (v17 & 65535)))
            _ = fi26
            _ = v27
            _ = v28
        }
        if (!true) {
            _ = len(v4)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if v2 {
                v29 := GenS5{F0: 1}
                _ = v29
            }
        }
        _ = fi25
    }
    v23 = (v23 - (v23 ^ (v23 & (^(v23)))))
    v30 := GenS5{F0: 100}
    v22 = v22
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
}
