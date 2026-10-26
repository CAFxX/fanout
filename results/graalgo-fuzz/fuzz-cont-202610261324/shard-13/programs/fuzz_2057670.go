package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
func genFn0(p0 uint64, p1 uint64) uint8 {
    for fi4 := 0; fi4 < 12; fi4++ {
        p1 = 2
        p0 = (^(p1))
        p1 += p1
        _ = fi4
    }
    v5 := []int{128, -390, -726, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1() uint32 {
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(fi6)
        fmt.Println(fi6)
        _ = fi6
    }
    switch 32767 {
    case 32768:
        var v7 float32 = ((float32(0.0) - -2.25) / (float32(1e10) * 3.14159))
        _ = v7
    }
    var v8 uint = 65535
    _ = v8
    return 2
}

func main() {
    var v1 uint = 95
    var v2 uint = v1
    var v3 uint64 = 237
    v1 -= (v2 & 42)
    v9 := GenS1{F0: (("\n\t" + "") + "\n\t"), F1: 32767}
    fmt.Println(v3)
    if false {
        switch "foo bar" {
        case "a":
            var v10 bool = false
            _ = v10
        default:
            if true {
                v11 := GenS1{F0: "foo bar", F1: 1}
                v12 := GenS1{F0: ("" + "foo bar"), F1: 42}
                _ = v11
                _ = v12
            }
        }
        var v13 int32 = -862
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := []int{fi14, fi14, -417}
            v13 = (v13 / (((v13 ^ v13)) | 1))
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(v2, fi16, v3)
        if (-543 < 0) {
            v1 += v1
        }
        v1 = v2
        _ = fi16
    }
    if (("foo bar" + "a") > ("\n\t" + "")) {
        v3 = (v3 * v3)
    } else {
        var v17 uint = v2
        v18 := GenS1{F0: ("" + ""), F1: 303}
        _ = v17
        _ = v18
    }
    v2 -= v1
    v19 := []int{127, 127, 128}
    var v20 int32 = 1
    v20 -= -292
    v3 += (v3 / (((v3 >> ((65535) & 63))) | 1))
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v19 = append(v19, 1)
    switch "hello" {
    case "foo bar":
        if len(v19) > 0 {
            v19[0] = 42
        }
    case "a":
        v21 := GenS1{F0: "foo bar", F1: 418}
        _ = v21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 uint64 = (v3 ^ (v3 | 42))
        v24 := GenS1{F0: "\n\t", F1: 32767}
        _ = fi22
        _ = v23
        _ = v24
    }
    v20 += v20
    v3 += 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
