package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint; F2 uint }
func genFn0(p0 int32, p1 bool, p2 uint32) int {
    for fi7 := 0; fi7 < 20; fi7++ {
        if p1 {
            p1 = (1e10 >= (1.5 / 1.0))
            p1 = (!(!p1))
        }
        _ = fi7
    }
    p0 += (p0 * (p0 ^ p0))
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    var v1 float64 = (1.5 - 1e10)
    v2 := make(map[string]uint32)
    var v3 float32 = 1e10
    var v4 uint32 = 2147483647
    v5 := make(map[int]string)
    v6 := make(map[uint32]bool)
    fmt.Println(v1, v4)
    v4 += v4
    if true {
        v8 := []int{-2147483648}
        _ = v8
    }
    switch 0 {
    case -735:
        switch (("foo bar" + "") + ("a" + "foo bar")) {
        case "foo bar":
            v1 -= v1
        case "\n\t":
            if (!((3.14159 <= v1) || true)) {
                v3 = 0.1
            } else {
                v4 = 0
            }
        case "hello":
            v9 := []int{}
            _ = v9
        default:
            delete(v2, "hello")
        }
    }
    var v10 int16 = 1
    v11 := GenS1{F0: "foo bar", F1: -1, F2: 134}
    v12 := GenS2{F0: "hello"}
    var v13 uint64 = 2
    if false {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 bool = (4294967295 > 42)
            _ = fi14
            _ = v15
        }
    } else {
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS2{F0: "\n\t"}
            if ((!false) || false) {
                if true {
                    var v18 int32 = 2147483647
                    v19 := v5[(fi16 * 42)]
                    _ = v18
                    _ = v19
                } else {
                    v4 = v4
                    v3 = ((v3 / v3) + ((0.0 - v3) + 1.5))
                    fmt.Println(v13)
                }
                for fi20 := 0; fi20 < 13; fi20++ {
                    v10 -= (v10 % (((v10 % ((v10) | 1))) | 1))
                    var v21 string = ""
                    v22 := GenS1{F0: v21, F1: 42, F2: 197}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            _ = fi16
            _ = v17
        }
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        v4 = (^(v4))
        switch fi23 {
        case 32768:
            v24 := GenS3{F0: 0, F1: 18446744073709551615, F2: 519}
            _ = v24
        default:
            v25 := []int{-880}
            _ = v25
        }
        _ = fi23
    }
    var v26 int16 = 1000
    if true {
        var v27 uint32 = (v4 >> ((65535) & 63))
        _ = v27
    } else {
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := []int{(fi28 * int(v26)), fi28, genFn0(-2147483648, false, (v4 & 65535))}
            v3 -= -2.25
            _ = fi28
            _ = v29
        }
    }
    v3 = 1.5
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
