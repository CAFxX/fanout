package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint, p1 int16, p2 int32) int32 {
    var v7 float64 = -2.25
    fmt.Println(v7, p0, v7)
    switch "" {
    case "":
        v8 := []int{-2147483648, 128}
        _ = v8
    case "a":
        var v9 float32 = (-2.25 + 0.1)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 & (p2 % ((17) | 1)))
}
func genFn1(p0 uint8) uint32 {
    switch "hello" {
    case "a":
        v10 := []int{0}
        _ = v10
    case "\n\t":
        fmt.Println(p0)
    case "":
        p0 -= p0
    }
    var v11 int32 = 0
    v12 := GenS1{F0: -1}
    v13 := []int{-307}
    switch (("hello" + ("foo bar" + "")) + ("a" + "")) {
    case "a":
        for fi14 := 0; fi14 < 8; fi14++ {
            v11 = 42
            p0 = 2
            _ = fi14
        }
    default:
        if len(v13) > 0 {
            v13[0] = -9223372036854775808
        }
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 85
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 128
    var v3 float32 = 1.5
    var v4 uint32 = 2147483648
    var v5 int64 = -1
    var v6 uint32 = (v4 + v4)
    for fi15 := 0; fi15 < 4; fi15++ {
        if false {
            for fi16 := 0; fi16 < 9; fi16++ {
                v17 := GenS1{F0: -899}
                var v18 string = ""
                var v19 float32 = float32(0.0)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            if ((false && false) || true) {
                var v20 uint = 4294967295
                _ = v20
            }
            fmt.Println(v3, v2, v5)
        } else {
            v21 := v1[32767]
            var v22 float64 = 1.5
            v6 -= genFn1(59)
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS1{F0: v5}
            v4 += 1
            var v25 uint = 4294967295
            _ = fi23
            _ = v24
            _ = v25
        }
        v26 := []int{fi15, (fi15 << ((0) & 63)), -2147483648, (-(v2))}
        _ = fi15
        _ = v26
    }
    var v27 uint = 255
    var v28 int8 = 0
    if true {
        switch (v2 | (v2 * v2)) {
        case -103:
            switch 2147483647 {
            case -496:
                v27 -= v27
            case 9223372036854775807:
                _ = len(v1)
            default:
                v1[v2] = "a"
            }
        case -9223372036854775808:
            if (v4 != (v6 / ((0) | 1))) {
                v27 += 586
            } else {
                v29 := v1[(v2 ^ (v2 << ((v27) & 63)))]
                v30 := GenS1{F0: 42}
                _ = v29
                _ = v30
            }
        case 0:
            v5 = v5
        }
        if false {
            var v31 float32 = v3
            if true {
                v28 -= (^(v28))
                for fi32 := 0; fi32 < 14; fi32++ {
                    var v33 int8 = (v28 * (^(v28)))
                    var v34 int16 = 189
                    v35 := GenS1{F0: v5}
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                }
            } else {
                v6 = v4
            }
            v2 += (v2 % ((127) | 1))
            _ = v31
        }
        var v36 int16 = -32768
        _ = v36
    } else {
        for fi37 := 0; fi37 < 5; fi37++ {
            if true {
                v38 := GenS1{F0: (-(v5))}
                if (!("" >= "foo bar")) {
                    v39 := []int{}
                    _ = v39
                } else {
                    fmt.Println(v5)
                    fmt.Println(v4, v5, v2)
                }
                v40 := []int{}
                _ = v38
                _ = v40
            } else {
                var v41 int32 = -483
                fmt.Println(v27, v4, v41)
                _ = v41
            }
            _ = fi37
        }
        var v42 int8 = (-(v28))
        _ = v42
    }
    v5 -= v5
    fmt.Println(v4, v5)
    v2 -= 4294967295
    delete(v1, -9223372036854775808)
    v43 := GenS1{F0: -1}
    var v44 float32 = 0.0
    var v45 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
