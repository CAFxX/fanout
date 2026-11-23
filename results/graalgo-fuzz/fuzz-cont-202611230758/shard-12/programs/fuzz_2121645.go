package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int8; F1 bool; F2 int8 }
func genFn0(p0 uint, p1 int64, p2 int32) uint16 {
    fmt.Println(p1)
    p0 = p0
    for fi4 := 0; fi4 < 7; fi4++ {
        p0 = 2
        var v5 float32 = 1e10
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{-544, 1}
    var v3 float32 = 1.5
    if (true || true) {
        v3 = v3
        _ = len(v1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v6 int8 = 0
    switch "" {
    case "foo bar":
        if true {
            _ = len(v1)
            v6 = 118
        }
    case "a":
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: 0}
            _ = fi7
            _ = v8
        }
    case "hello":
        v9 := GenS2{F0: false, F1: 44}
        _ = v9
    }
    v3 += v3
    if true {
        for fi10 := 0; fi10 < 18; fi10++ {
            if true {
                if len(v2) > 0 {
                    v2[0] = fi10
                }
            } else {
                var v11 int64 = 363
                _ = v11
            }
            var v12 uint64 = 0
            v3 += v3
            _ = fi10
            _ = v12
        }
        if true {
            for fi13 := 0; fi13 < 13; fi13++ {
                v3 -= ((0.0 + (3.14159 / v3)) / 1e10)
                fmt.Println(v6, fi13)
                _ = fi13
            }
            switch -844 {
            case 42:
                v3 -= 1.5
            }
            v6 -= (v6 >> ((2) & 63))
        }
    }
    v3 -= 1.5
    var v14 int = 32767
    v2 = append(v2, v14)
    if (42 != 768) {
        for fi15 := 0; fi15 < 3; fi15++ {
            if (false || false) {
                v6 = (v6 % (((v6 % ((int8(v14)) | 1))) | 1))
            }
            var v16 bool = false
            delete(v1, "\n\t")
            _ = fi15
            _ = v16
        }
        var v17 float64 = (1.5 * 0.0)
        v6 += 100
        _ = v17
    } else {
        if ((1e10 - (0.0 - (float64((v3 * 1e10)) * (float64(v3) / 3.14159)))) <= ((3.14159 / 1e10) * 0.1)) {
            var v18 bool = (!((v6 + (v6 ^ 1)) < (^(v6))))
            var v19 uint16 = 483
            _ = v18
            _ = v19
        } else {
            v20 := GenS3{F0: v6, F1: (0 == 671), F2: -128}
            _ = v20
        }
    }
    var v21 bool = (!((!(1 < (v6 % ((-6) | 1)))) || (203 == genFn0(821, -9223372036854775808, -59))))
    var v22 int8 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v6
}
