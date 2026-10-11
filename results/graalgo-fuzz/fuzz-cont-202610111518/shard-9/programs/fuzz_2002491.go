package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int32, p1 int, p2 int) uint16 {
    if false {
        if (!true) {
            v7 := GenS1{F0: 1}
            v8 := []int{p1}
            _ = v7
            _ = v8
        } else {
            var v9 uint64 = 4294967295
            _ = v9
        }
        var v10 int64 = 9223372036854775807
        p2 += (^(p1))
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 int8 = 100
        p1 = -745
        fmt.Println(fi11)
        _ = fi11
        _ = v12
    }
    v13 := []int{106}
    var v14 uint32 = 65535
    v15 := GenS1{F0: 138}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 23
}
func genFn1(p0 int32) uint32 {
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            if (1 > 454) {
                var v18 uint16 = 569
                _ = v18
            }
            _ = fi17
        }
        _ = fi16
    }
    var v19 uint64 = 2
    _ = p0
    _ = v19
    return 42
}

func main() {
    var v1 uint8 = 155
    var v2 float32 = 0.0
    v3 := make(map[uint32]int)
    var v4 int8 = -118
    var v5 float64 = 0.0
    var v6 bool = true
    v20 := GenS1{F0: (v1 | v1)}
    v2 -= 3.14159
    switch "\n\t" {
    case "":
        for fi21 := 0; fi21 < 5; fi21++ {
            v2 -= (v2 / v2)
            _ = fi21
        }
    case "\n\t":
        if ((v1 - (v1 % (((v1 * 0)) | 1))) != 42) {
            v4 += (v4 - (^(v4)))
            var v22 int64 = -402
            _ = v22
        } else {
            var v23 int8 = v4
            v24 := GenS1{F0: (^(v1))}
            v3[4294967295] = 1
            _ = v23
            _ = v24
        }
    case "a":
        var v25 int64 = -9223372036854775808
        _ = v25
    default:
        for fi26 := 0; fi26 < 11; fi26++ {
            fmt.Println(fi26)
            for fi27 := 0; fi27 < 2; fi27++ {
                switch fi27 {
                case -9223372036854775808:
                    var v28 float32 = 3.14159
                    _ = v28
                case 4294967295:
                    if (!(!v6)) {
                        var v29 int8 = -121
                        v1 = (v1 ^ 127)
                        _ = v29
                    }
                }
                v5 -= v5
                v2 = ((((0.0 * (v2 * v2)) * (3.14159 * -2.25)) - (-2.25 / 1.0)) * (v2 / float32(0.1)))
                var v30 string = "foo bar"
                _ = fi27
                _ = v30
            }
            _ = fi26
        }
    }
    v4 = (v4 + (v4 >> ((42) & 63)))
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := []int{(-(fi31)), (fi31 & (fi31 ^ fi31)), (^(fi31))}
        v4 = (v4 * (v4 | 100))
        var v33 float64 = v5
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 float64 = 0.1
    fmt.Println(v34, v4, v2)
    v4 = v4
    v34 += v34
    for fi35 := 0; fi35 < 9; fi35++ {
        v6 = v6
        switch fi35 {
        case 32767:
            fmt.Println(v2, v5)
        case 128:
            var v36 int8 = v4
            _ = v36
        case 420:
            var v37 int32 = -2147483648
            _ = v37
        }
        _ = fi35
    }
    v34 = v5
    v2 = 1e10
    v2 += v2
    fmt.Println(v6)
    switch (("a" + "a") + ("hello" + "foo bar")) {
    case "foo bar":
        var v38 bool = (genFn1(662) > 1)
        _ = v38
    case "\n\t":
        for fi39 := 0; fi39 < 6; fi39++ {
            v6 = v6
            var v40 int32 = -491
            _ = fi39
            _ = v40
        }
    case "hello":
        var v41 int32 = 0
        _ = v41
    default:
        v5 = (v34 - float64(v2))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
