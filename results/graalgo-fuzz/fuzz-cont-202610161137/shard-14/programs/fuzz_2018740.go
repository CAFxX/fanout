package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint }
func genFn0() int8 {
    v4 := []int{2147483647, 41, 9223372036854775807, -353}
    switch 128 {
    case 32768:
        var v5 uint = 0
        _ = v5
    case 0:
        if len(v4) > 0 {
            v4[0] = 460
        }
    case -478:
        var v6 int16 = 662
        _ = v6
    default:
        if len(v4) > 0 {
            v4[0] = 32768
        }
    }
    var v7 uint32 = 4294967295
    switch "hello" {
    case "a":
        v8 := GenS1{F0: ("a" + "a")}
        _ = v8
    }
    _ = len(v4)
    _ = v4
    _ = v7
    return -31
}
func genFn1(p0 uint32, p1 int8, p2 uint64) int8 {
    var v9 uint64 = p2
    var v10 float64 = 0.0
    if (!(!(1.5 >= (1e10 + 0.0)))) {
        v11 := GenS1{F0: "foo bar"}
        if (!true) {
            var v12 float64 = -2.25
            switch "hello" {
            case "hello":
                var v13 int64 = -308
                _ = v13
            default:
                v12 = ((float64(-2.25) * v12) + 1.5)
            }
            _ = v12
        } else {
            p1 += (^(p1))
            v14 := []int{}
            _ = v14
        }
        fmt.Println(v10, p2, v10)
        _ = v11
    } else {
        p2 = v9
    }
    var v15 bool = (false || false)
    switch "foo bar" {
    case "a":
        if (false || true) {
            p0 -= p0
            fmt.Println(v9, p1, p2)
            p0 -= 4294967295
        } else {
            fmt.Println(p1)
            v16 := GenS1{F0: "a"}
            var v17 float32 = 0.1
            _ = v16
            _ = v17
        }
    case "hello":
        if false {
            if (!((2 < 808) || (((42 >= 1) && ((("a" + "foo bar") + "a") > "hello")) && v15))) {
                v18 := []int{32767}
                _ = v18
            } else {
                if false {
                    v19 := []int{}
                    v20 := GenS1{F0: "foo bar"}
                    _ = v19
                    _ = v20
                }
            }
            if true {
                v21 := []int{9223372036854775807, -1}
                _ = v21
            } else {
                p1 = p1
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v15
    _ = v9
    return p1
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]bool)
    var v3 float32 = 1.5
    var v22 bool = false
    var v23 int64 = -1
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 uint32 = 0
        _ = fi24
        _ = v25
    }
    var v26 int = 786
    v27 := GenS1{F0: ""}
    var v28 float32 = 1e10
    v23 = (v23 >> ((83) & 63))
    v1 = v1
    if v22 {
        v29 := GenS1{F0: (("\n\t" + ("hello" + "hello")) + "")}
        _ = v29
    }
    fmt.Println(v23, v28)
    var v30 int8 = -128
    v31 := GenS1{F0: ("foo bar" + "foo bar")}
    var v32 uint16 = 42
    if (float32(0.0) > 3.14159) {
        for fi33 := 0; fi33 < 9; fi33++ {
            v30 -= 100
            _ = fi33
        }
        v32 += (v32 / (((^(v32))) | 1))
    }
    v2[-623] = (!(v22 && v22))
    fmt.Println(v22, v28)
    switch "a" {
    case "hello":
        v34 := GenS1{F0: "foo bar"}
        _ = v34
    case "foo bar":
        var v35 uint8 = 0
        _ = v35
    case "":
        fmt.Println(v28, v3)
    }
    switch (^(v26)) {
    case -1:
        if false {
            var v36 int32 = 42
            v37 := []int{v26, (v26 ^ (v26 * 9223372036854775807)), v26, -1}
            _ = v36
            _ = v37
        } else {
            v38 := GenS2{F0: (566 <= 4294967295), F1: 1}
            _ = v38
        }
    case 58:
        v39 := []int{0, 4294967295, (v26 * 42)}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
