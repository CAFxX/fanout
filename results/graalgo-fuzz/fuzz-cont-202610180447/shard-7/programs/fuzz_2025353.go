package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint32 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0(p0 int, p1 string, p2 uint32) uint32 {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 int = (p0 / ((-1) | 1))
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: (p2 | 0), F1: p1, F2: p2}
    var v7 int32 = -2147483648
    v8 := GenS2{F0: "a", F1: 1}
    p2 -= 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (p2 << ((1) & 63))
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    var v3 int64 = 9223372036854775807
    var v9 float32 = 3.14159
    if true {
        v9 -= (float32(-2.25) / 1.0)
        if (841 <= 32767) {
            v3 -= (v3 & (v3 * 42))
            v3 -= v3
        } else {
            v10 := GenS3{F0: 2}
            v11 := GenS3{F0: 42}
            for fi12 := 0; fi12 < 1; fi12++ {
                delete(v1, (^(fi12)))
                _ = fi12
            }
            _ = v10
            _ = v11
        }
        v1[128] = 4294967295
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            if false {
                for fi14 := 0; fi14 < 5; fi14++ {
                    v3 = v3
                    switch ("hello" + "") {
                    case "a":
                        v15 := v1[797]
                        _ = v15
                    default:
                        v3 = v3
                        var v16 int32 = 1
                        _ = v16
                    }
                    _ = fi14
                }
            } else {
                var v17 float64 = -2.25
                v3 = (v3 | (v3 & -618))
                var v18 float64 = v17
                _ = v17
                _ = v18
            }
            _ = fi13
        }
        var v19 int32 = -1
        var v20 uint16 = 0
        _ = v19
        _ = v20
    }
    delete(v1, 42)
    v3 = v3
    if ((!(!(2147483647 > 128))) && (!true)) {
        var v21 int64 = -9223372036854775808
        _ = v21
    } else {
        var v22 float32 = v9
        fmt.Println(v22, v3, v3)
        fmt.Println(v9, v9)
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        v24 := GenS3{F0: 0}
        var v25 int8 = 0
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        switch fi26 {
        case 94:
            delete(v1, (fi26 >> ((2) & 63)))
        case 4294967295:
            if false {
                var v27 float64 = 0.1
                v3 -= v3
                _ = v27
            }
        default:
            v28 := GenS3{F0: 4294967295}
            _ = v28
        }
        _ = fi26
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        if true {
            var v30 uint64 = 2
            _ = v30
        }
        _ = len(v2)
        if true {
            if false {
                switch (fi29 * int(v3)) {
                case -2147483648:
                    _ = len(v2)
                case -9223372036854775808:
                    v2 = append(v2, "a")
                }
            }
            for fi31 := 0; fi31 < 17; fi31++ {
                var v32 uint16 = 2
                delete(v1, 9223372036854775807)
                var v33 uint8 = 8
                if false {
                    v34 := GenS4{F0: 255, F1: 42}
                    v2 = append(v2, "")
                    switch "hello" {
                    case "hello":
                        v35 := GenS4{F0: 18446744073709551615, F1: 350}
                        _ = v35
                    }
                    _ = v34
                }
                if len(v2) > 0 {
                    v2[0] = ("foo bar" + "hello")
                }
                var v36 int64 = (^(v3))
                _ = fi31
                _ = v32
                _ = v33
                _ = v36
            }
            v9 -= v9
        }
        _ = fi29
    }
    v2 = append(v2, "a")
    var v37 bool = (true || true)
    v2 = append(v2, "hello")
    v9 = float32(0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v9
}
