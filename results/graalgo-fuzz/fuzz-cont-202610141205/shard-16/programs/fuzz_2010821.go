package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string; F1 string; F2 uint8 }
func genFn0(p0 int16, p1 uint32, p2 int) int8 {
    v7 := []int{-310, p2, p2}
    p1 -= p1
    var v8 int = (p2 ^ -9223372036854775808)
    v8 -= (v8 ^ p2)
    v9 := GenS1{F0: 4294967295, F1: 982}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 float64 = (1.5 + 0.1)
    var v2 int = -397
    v3 := make(map[uint32]string)
    var v4 uint32 = 78
    v5 := make(map[string]bool)
    var v6 float32 = 0.0
    v1 += v1
    if (1 < -979) {
        fmt.Println(v2)
        var v10 bool = false
        if v10 {
            fmt.Println(v6, v1)
            var v11 int64 = -68
            var v12 int64 = v11
            _ = v11
            _ = v12
        } else {
            v4 = 0
            switch -2147483648 {
            case 1:
                var v13 int32 = 2147483647
                _ = v13
            case -2147483648:
                v2 += v2
            }
            switch 9223372036854775807 {
            case 127:
                v14 := GenS1{F0: 65535, F1: 144}
                _ = v14
            case 2147483647:
                fmt.Println(v6, v6)
            }
        }
        _ = v10
    } else {
        var v15 float64 = 3.14159
        var v16 string = ("hello" + "foo bar")
        for fi17 := 0; fi17 < 4; fi17++ {
            v15 = v1
            fmt.Println(v1)
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        switch "\n\t" {
        case "\n\t":
            var v19 int32 = 1
            _ = v19
        case "hello":
            fmt.Println(fi18, v1)
        }
        _ = fi18
    }
    if (v2 != (v2 | v2)) {
        v20 := GenS1{F0: 42, F1: 0}
        for fi21 := 0; fi21 < 17; fi21++ {
            v6 = (v6 + (float32(1e10) * ((1e10 + v6) * v6)))
            v22 := GenS1{F0: 596, F1: 42}
            _ = fi21
            _ = v22
        }
        v1 += v1
        _ = v20
    } else {
        delete(v5, "")
    }
    v23 := GenS2{F0: v4, F1: 88}
    v6 += 1e10
    v24 := GenS1{F0: 155, F1: 42}
    var v25 int32 = 1
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 uint32 = 65535
        _ = fi26
        _ = v27
    }
    v6 += (v6 * v6)
    if (100 >= 55) {
        if true {
            var v28 bool = true
            delete(v5, ("foo bar" + ("hello" + "")))
            _ = v28
        }
        v6 = v6
        var v29 uint = 913
        _ = v29
    } else {
        for fi30 := 0; fi30 < 16; fi30++ {
            v31 := GenS2{F0: v4, F1: 100}
            v6 += 3.14159
            _ = fi30
            _ = v31
        }
        v32 := []int{}
        _ = v32
    }
    v25 = (v25 / ((v25) | 1))
    if true {
        var v33 bool = (!false)
        _ = v33
    } else {
        v25 += 42
        v34 := GenS3{F0: "\n\t", F1: "\n\t", F2: 0}
        _ = v34
    }
    fmt.Println(v2)
    fmt.Println(v25, v1, v2)
    for fi35 := 0; fi35 < 12; fi35++ {
        for fi36 := 0; fi36 < 17; fi36++ {
            var v37 int8 = -1
            fmt.Println(fi35, v4, v25)
            var v38 float32 = 0.0
            v25 += (v25 + (-(v25)))
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    v39 := GenS2{F0: v4, F1: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
