package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0(p0 int32, p1 int16, p2 int8) int32 {
    p0 = -2147483648
    p1 = p1
    v7 := []int{int(p0), 32767, -9223372036854775808, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 << ((867) & 63))
}
func genFn1(p0 int8) bool {
    v8 := GenS1{F0: 100}
    var v9 int = 42
    v10 := []int{(^(v9))}
    for fi11 := 0; fi11 < 13; fi11++ {
        if false {
            p0 -= p0
            var v12 uint32 = 417
            switch fi11 {
            case 1:
                if true {
                    v13 := GenS2{F0: -2147483648}
                    if len(v10) > 0 {
                        v10[0] = v9
                    }
                    if len(v10) > 0 {
                        v10[0] = int(p0)
                    }
                    _ = v13
                }
            case -353:
                var v14 float64 = float64((1e10 * 0.0))
                _ = v14
            case 291:
                fmt.Println(p0, v9, fi11)
            }
            _ = v12
        } else {
            var v15 uint = 152
            var v16 int16 = -32768
            fmt.Println(v9, fi11, fi11)
            _ = v15
            _ = v16
        }
        if true {
            v17 := GenS2{F0: 793}
            var v18 uint16 = 179
            if (59 >= 2147483647) {
                if false {
                    var v19 int = v9
                    fmt.Println(p0, v18)
                    v9 += v9
                    _ = v19
                }
                fmt.Println(v18)
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 int32 = 1
                    _ = fi20
                    _ = v21
                }
            }
            _ = v17
            _ = v18
        } else {
            var v22 uint64 = 0
            if true {
                var v23 int32 = 42
                v23 = (v23 / ((v23) | 1))
                var v24 int16 = 0
                _ = v23
                _ = v24
            }
            _ = v22
        }
        _ = fi11
    }
    var v25 bool = true
    _ = p0
    _ = v10
    _ = v25
    _ = v8
    _ = v9
    return v25
}

func main() {
    var v1 bool = (((70 < 152) || (!true)) || true)
    v2 := make(map[string]uint32)
    var v3 uint16 = 255
    var v4 uint64 = 4294967295
    var v5 int32 = -1
    var v6 uint8 = 0
    switch "hello" {
    case "a":
        v2[("\n\t" + ("hello" + "a"))] = 710
    case "\n\t":
        fmt.Println(v1, v1, v4)
    case "":
        fmt.Println(v1, v6)
    }
    v26 := GenS2{F0: 127}
    v27 := GenS1{F0: (v3 >> ((664) & 63))}
    switch "foo bar" {
    case "a":
        var v28 uint8 = v6
        _ = v28
    case "foo bar":
        v29 := []int{32768, 9223372036854775807, 32767, 32768}
        _ = v29
    case "":
        var v30 uint16 = v3
        _ = v30
    }
    v31 := GenS1{F0: v3}
    v1 = v1
    fmt.Println(v6, v6)
    if false {
        var v32 string = "foo bar"
        for fi33 := 0; fi33 < 14; fi33++ {
            var v34 float64 = (3.14159 / 1e10)
            _ = fi33
            _ = v34
        }
        _ = v32
    } else {
        v35 := GenS3{F0: 60, F1: 42}
        v3 = (^(v3))
        v6 -= (v6 & v6)
        _ = v35
    }
    switch -1 {
    case 128:
        for fi36 := 0; fi36 < 18; fi36++ {
            v3 += v3
            v1 = true
            v37 := []int{(fi36 << ((44) & 63)), (^(fi36)), (fi36 | (fi36 ^ 9223372036854775807)), fi36}
            _ = fi36
            _ = v37
        }
    case 4294967295:
        v3 -= v3
    case 42:
        v4 = (v4 << ((0) & 63))
    default:
        v5 = (v5 * v5)
    }
    _ = len(v2)
    v38 := GenS3{F0: 5, F1: 2147483647}
    for fi39 := 0; fi39 < 14; fi39++ {
        v40 := GenS2{F0: fi39}
        v6 -= v6
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
