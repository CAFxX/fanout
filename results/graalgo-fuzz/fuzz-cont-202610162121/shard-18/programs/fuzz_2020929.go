package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 string }
type GenS2 struct { F0 int; F1 uint64; F2 string }
type GenS3 struct { F0 uint16; F1 uint32; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 int; F2 uint }

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    fmt.Println(v1)
    fmt.Println(v1)
    v4 := GenS1{F0: 314, F1: 0, F2: ("\n\t" + "")}
    v1 = 65535
    var v5 float64 = ((1.5 * 3.14159) / (1.5 * -2.25))
    v6 := GenS1{F0: 561, F1: 393, F2: (("" + "foo bar") + "a")}
    for fi7 := 0; fi7 < 18; fi7++ {
        v1 += (v1 % (((v1 >> ((uint(v1)) & 63))) | 1))
        if (-1 <= (fi7 >> ((70) & 63))) {
            v8 := GenS1{F0: uint(v1), F1: -941, F2: ("\n\t" + (("\n\t" + "hello") + "\n\t"))}
            _ = v8
        } else {
            var v9 int64 = -9223372036854775808
            _ = v9
        }
        if (!true) {
            v10 := GenS1{F0: 744, F1: 1, F2: ("hello" + "foo bar")}
            var v11 bool = true
            _ = v10
            _ = v11
        } else {
            v1 = (v1 << ((545) & 63))
        }
        _ = fi7
    }
    v1 -= v1
    switch (("\n\t" + ("\n\t" + ("" + "foo bar"))) + ("\n\t" + ("\n\t" + ("foo bar" + "hello")))) {
    case "":
        var v12 uint64 = 1
        _ = v12
    case "a":
        v13 := GenS2{F0: -9223372036854775808, F1: 4294967295, F2: (("foo bar" + "\n\t") + ("hello" + ("" + ("hello" + ("foo bar" + ("\n\t" + "hello"))))))}
        _ = v13
    }
    v1 = (v1 ^ v1)
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS1{F0: 380, F1: -2147483648, F2: "a"}
            switch fi14 {
            case 42:
                if true {
                    v17 := GenS3{F0: 65535, F1: 666, F2: (-(fi15))}
                    v5 -= v5
                    _ = v17
                }
            default:
                v18 := GenS2{F0: (fi15 | (^(fi14))), F1: 0, F2: "hello"}
                _ = v18
            }
            v19 := []int{42, (fi15 / ((0) | 1))}
            v5 += (float64(-2.25) / v5)
            v5 += v5
            _ = fi15
            _ = v16
            _ = v19
        }
        delete(v2, (fi14 / ((fi14) | 1)))
        _ = fi14
    }
    fmt.Println(v5, v1)
    var v20 bool = false
    v21 := v3["hello"]
    var v22 int64 = 0
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS3{F0: v1, F1: uint32(v1), F2: fi23}
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := GenS4{F0: 239}
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    var v27 int8 = 0
    switch (v21 >> ((1) & 63)) {
    case 1:
        if (42 > 1000) {
            v20 = v20
            fmt.Println(v22, v20, v27)
        }
    case 32767:
        if (v27 == 42) {
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 int = 32768
                _ = fi28
                _ = v29
            }
            v30 := GenS5{F0: -1, F1: v21, F2: 947}
            _ = v30
        } else {
            v21 += 51
            v31 := []int{v21}
            v31 = append(v31, v21)
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
