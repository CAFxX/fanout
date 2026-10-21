package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 uint64 }
func genFn0(p0 string) int16 {
    fmt.Println(p0)
    v4 := GenS1{F0: 2, F1: 145, F2: 2}
    var v5 float64 = 1e10
    _ = p0
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint32, p1 string) bool {
    var v6 int16 = genFn0(((p1 + p1) + (p1 + p1)))
    p0 -= (^(p0))
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS2{F0: ((fi7 | (fi7 - fi7)) < fi7), F1: (-(fi7))}
        v9 := []int{(-(fi7)), (fi7 % ((128) | 1)), (fi7 ^ 128)}
        if (0.1 != 0.0) {
            fmt.Println(p1, fi7, v6)
        } else {
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 int32 = 0
                switch -1 {
                case 128:
                    fmt.Println(p1, v11, fi10)
                case -184:
                    _ = len(v9)
                }
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 17; fi12++ {
                v6 -= 1000
                v9 = append(v9, -685)
                fmt.Println(p0)
                _ = fi12
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return false
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -681
    var v3 string = "foo bar"
    var v13 bool = true
    fmt.Println(v13)
    v2 += (v2 / ((949) | 1))
    for fi14 := 0; fi14 < 18; fi14++ {
        if (true || v13) {
            var v15 uint = 42
            _ = v15
        }
        _ = fi14
    }
    v13 = (!(725 > 809))
    v16 := GenS1{F0: 0, F1: -290, F2: 20}
    delete(v1, -378)
    v13 = v13
    v17 := GenS3{F0: "foo bar", F1: 65535}
    var v18 uint8 = 0
    if v13 {
        if v13 {
            v19 := GenS3{F0: "hello", F1: 2}
            var v20 int16 = 0
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 2; fi21++ {
                if (2147483648 < 1) {
                    var v22 int32 = 1
                    for fi23 := 0; fi23 < 14; fi23++ {
                        var v24 uint32 = 4294967295
                        _ = fi23
                        _ = v24
                    }
                    fmt.Println(fi21, v3, v2)
                    _ = v22
                } else {
                    fmt.Println(v2)
                }
                for fi25 := 0; fi25 < 13; fi25++ {
                    var v26 int32 = 1
                    v26 += (v2 | v26)
                    v3 = v3
                    fmt.Println(v26, fi21, v13)
                    _ = fi25
                    _ = v26
                }
                _ = fi21
            }
            switch v3 {
            case "a":
                v27 := GenS4{F0: 322}
                _ = v27
            }
            v28 := GenS5{F0: -100, F1: 42}
            _ = v28
        }
        v18 += (^(v18))
    }
    if v13 {
        var v29 bool = (!(!(v13 && (2147483647 > 42))))
        if (v13 || v13) {
            v30 := GenS3{F0: (v3 + v3), F1: 1}
            v1[2147483647] = 128
            _ = v30
        }
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
