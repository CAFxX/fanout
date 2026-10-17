package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 string }
func genFn0(p0 int16, p1 int32, p2 uint8) int8 {
    var v6 int16 = p0
    var v7 string = ("hello" + "a")
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 bool = (2147483647 != -2147483648)
    var v2 int64 = -9223372036854775808
    var v3 uint8 = 127
    var v4 int32 = -2147483648
    var v5 uint = 65535
    var v8 int64 = v2
    fmt.Println(v2)
    if v1 {
        v1 = ((("foo bar" + "foo bar") + ("" + "\n\t")) == "\n\t")
        var v9 int32 = v4
        for fi10 := 0; fi10 < 17; fi10++ {
            v2 += (v8 & (v2 >> ((v5) & 63)))
            var v11 bool = (!(((v9 / (((v9 & (v4 & 0))) | 1)) == v9) || true))
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        var v12 uint = (^(v5))
        v13 := []int{0, 85}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := []int{128}
        _ = fi14
        _ = v15
    }
    var v16 bool = (v5 > (^(v5)))
    v17 := []int{}
    if v1 {
        var v18 uint32 = 4294967295
        v19 := GenS1{F0: "foo bar", F1: 32767, F2: "hello"}
        for fi20 := 0; fi20 < 15; fi20++ {
            v1 = (-2.25 == 1e10)
            for fi21 := 0; fi21 < 18; fi21++ {
                v4 -= v4
                v5 += v5
                var v22 float32 = 3.14159
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        _ = v18
        _ = v19
    }
    var v23 int32 = v4
    if (!(1.5 != 1e10)) {
        if v1 {
            v24 := GenS1{F0: "a", F1: -1, F2: ""}
            _ = v24
        } else {
            v3 -= v3
            v23 = -1
        }
    } else {
        var v25 float64 = 0.1
        fmt.Println(v1, v5, v16)
        _ = v25
    }
    if (-45 == 1) {
        var v26 int64 = v8
        v1 = (v26 <= v2)
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
