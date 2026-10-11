package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int16 }
type GenS2 struct { F0 int64; F1 int }
func genFn0(p0 int8) uint32 {
    v4 := []int{-1, 0}
    p0 += -1
    _ = p0
    _ = v4
    return 65535
}
func genFn1(p0 int, p1 int16, p2 int64) int16 {
    if false {
        fmt.Println(p0, p0)
    }
    p2 = (p2 % (((p2 >> ((427) & 63))) | 1))
    if true {
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := GenS1{F0: 1, F1: (p2 << ((0) & 63)), F2: p1}
            switch (fi5 ^ (^(p0))) {
            case 692:
                p2 = -1
            case 2147483647:
                fmt.Println(p2, p0)
            }
            v7 := []int{}
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 << ((18446744073709551615) & 63))
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 0
    var v3 float32 = -2.25
    var v8 int64 = 42
    v1["foo bar"] = -576
    v9 := v1["a"]
    switch v9 {
    case -264:
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := []int{}
                v13 := GenS1{F0: 1, F1: v8, F2: 42}
                _ = fi11
                _ = v12
                _ = v13
            }
            v3 = v3
            _ = fi10
        }
    case 9223372036854775807:
        if (0 <= uint8(v2)) {
            v8 = (v8 & (-(v8)))
            v3 = ((v3 + v3) * ((1e10 * 0.1) + ((3.14159 / v3) / float32(-2.25))))
            var v14 int64 = (v8 & 0)
            _ = v14
        }
    case 32767:
        var v15 int16 = genFn1(-440, -32768, (v8 & v8))
        _ = v15
    default:
        v16 := v1["a"]
        _ = v16
    }
    v1["a"] = v9
    v8 = v8
    for fi17 := 0; fi17 < 8; fi17++ {
        v2 -= v2
        var v18 uint8 = 2
        v19 := []int{fi17, fi17, (v9 + (-(v9))), v9}
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 uint16 = 654
    fmt.Println(v20)
    var v21 uint8 = 1
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 bool = (!false)
        v24 := GenS2{F0: (-(v8)), F1: 127}
        var v25 int = fi22
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 uint16 = v20
    var v27 uint32 = (v2 - (^(v2)))
    if false {
        switch ("foo bar" + ("foo bar" + "\n\t")) {
        case "a":
            if ("a" == "\n\t") {
                v28 := GenS2{F0: (v8 * v8), F1: (v9 | 4294967295)}
                v1["hello"] = v9
                _ = v28
            } else {
                v8 = v8
            }
        case "foo bar":
            var v29 int32 = 42
            _ = v29
        default:
            switch "\n\t" {
            case "":
                var v30 int16 = genFn1(32768, -32768, (-(v8)))
                _ = v30
            default:
                for fi31 := 0; fi31 < 2; fi31++ {
                    var v32 int = (v9 ^ (v9 >> ((42) & 63)))
                    for fi33 := 0; fi33 < 8; fi33++ {
                        v34 := GenS2{F0: 9223372036854775807, F1: fi31}
                        v3 += v3
                        _ = fi33
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
            }
        }
        v35 := GenS1{F0: -81, F1: (^(v8)), F2: 1000}
        _ = v35
    } else {
        fmt.Println(v27)
        for fi36 := 0; fi36 < 5; fi36++ {
            switch ("a" + "") {
            case "":
                fmt.Println(v2, v3, v27)
            case "a":
                for fi37 := 0; fi37 < 16; fi37++ {
                    fmt.Println(v21, fi37)
                    v20 = 32767
                    _ = fi37
                }
            default:
                var v38 int32 = -448
                _ = v38
            }
            fmt.Println(v21, v27)
            fmt.Println(fi36, fi36, fi36)
            _ = fi36
        }
        v39 := []int{v9, -263, v9, (v9 & v9)}
        _ = v39
    }
    var v40 int8 = -13
    v9 = (v9 * v9)
    var v41 int64 = (v8 & (v8 | v8))
    fmt.Println(v21, v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v40
    _ = v41
    _ = v8
    _ = v9
}
