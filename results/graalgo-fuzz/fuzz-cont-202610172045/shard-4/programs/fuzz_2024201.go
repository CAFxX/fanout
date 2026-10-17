package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 bool) bool {
    var v4 uint16 = 32767
    switch "\n\t" {
    case "foo bar":
        var v5 int8 = 35
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 int32 = 42
            var v8 int16 = 1000
            var v9 bool = (v8 > v8)
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    case "a":
        v4 += 2
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v4 += v4
        v4 -= (v4 | (v4 * 1))
        _ = fi10
    }
    if p0 {
        fmt.Println(p0, v4)
    }
    _ = p0
    _ = v4
    return ((9223372036854775807 > -9223372036854775808) || true)
}
func genFn1(p0 bool) uint64 {
    var v11 uint = 921
    v12 := []int{128}
    var v13 int16 = 42
    if p0 {
        if ((!(!(true || false))) || p0) {
            v11 = v11
        }
        var v14 bool = (!p0)
        _ = v14
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 729
}

func main() {
    var v1 float64 = float64(1.5)
    var v2 uint16 = 32767
    var v3 uint = 737
    v2 = (v2 - v2)
    for fi15 := 0; fi15 < 14; fi15++ {
        if true {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int16 = 124
                var v18 bool = (-2.25 != (v1 / (float64((1.5 + (0.1 / 1.5))) + v1)))
                _ = fi16
                _ = v17
                _ = v18
            }
            switch ((("hello" + "a") + "foo bar") + "foo bar") {
            case "":
                v19 := []int{}
                _ = v19
            case "a":
                switch (-(fi15)) {
                case 127:
                    v20 := []int{fi15}
                    _ = v20
                case 4294967295:
                    v21 := GenS1{F0: -1}
                    _ = v21
                }
            case "\n\t":
                v2 = v2
            default:
                var v22 int32 = 2147483647
                _ = v22
            }
        }
        v23 := GenS1{F0: 934}
        var v24 int8 = 100
        _ = fi15
        _ = v23
        _ = v24
    }
    v25 := []int{}
    fmt.Println(v3, v3)
    for fi26 := 0; fi26 < 12; fi26++ {
        v1 -= v1
        _ = fi26
    }
    var v27 int32 = 1
    for fi28 := 0; fi28 < 8; fi28++ {
        v25 = append(v25, -9223372036854775808)
        _ = fi28
    }
    v2 += (v2 >> ((v3) & 63))
    _ = len(v25)
    fmt.Println(v27)
    v3 = (v3 % (((v3 & v3)) | 1))
    var v29 uint64 = genFn1((((((1.5 * (1e10 - 0.1)) + 1e10) < ((((((3.14159 + (-2.25 / (-2.25 / 0.1))) - 1e10) - 3.14159) / ((3.14159 - 1.5) + (3.14159 + ((float32((1e10 / v1)) - -2.25) - 0.1)))) * (0.0 * float32(float64(float32((float64(1.5) / 0.1)))))) + 1.5)) || false) || true))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v29
    _ = v3
}
