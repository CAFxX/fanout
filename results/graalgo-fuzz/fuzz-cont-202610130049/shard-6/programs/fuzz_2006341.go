package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int32; F1 int; F2 int16 }
func genFn0(p0 string) int64 {
    v4 := GenS1{F0: -754, F1: ((p0 + p0) + p0), F2: -1}
    if (42 == 1) {
        v5 := GenS2{F0: (p0 + "\n\t"), F1: -4}
        v6 := []int{1, 32768}
        var v7 bool = (!true)
        _ = v5
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 3; fi8++ {
            p0 = ((p0 + p0) + "a")
            _ = fi8
        }
        switch p0 {
        case "foo bar":
            v9 := GenS2{F0: (p0 + ("" + (p0 + p0))), F1: 11}
            _ = v9
        }
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS2{F0: p0, F1: 42}
                var v13 int64 = -1
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    }
    v14 := []int{609, 128, 4294967295}
    if len(v14) > 0 {
        v14[0] = 1
    }
    _ = p0
    _ = v14
    _ = v4
    return -27
}

func main() {
    var v1 int64 = 0
    var v2 int64 = (-(v1))
    v3 := make(map[uint32]bool)
    var v15 uint8 = 255
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 string = "foo bar"
        _ = fi16
        _ = v17
    }
    _ = len(v3)
    fmt.Println(v15, v2, v2)
    var v18 uint8 = v15
    if true {
        fmt.Println(v1, v18)
        v19 := v3[0]
        v1 = v2
        _ = v19
    }
    var v20 uint32 = 679
    v15 = 1
    var v21 bool = (65535 != 42)
    var v22 float32 = float32(0.0)
    switch 32768 {
    case -315:
        v2 -= v2
    case 2147483647:
        var v23 string = "a"
        _ = v23
    case -420:
        v1 = (v2 ^ (v1 + genFn0(("hello" + "foo bar"))))
    }
    if (true && v21) {
        var v24 int64 = genFn0("foo bar")
        _ = v24
    }
    v25 := GenS3{F0: -529, F1: 32768, F2: -1}
    for fi26 := 0; fi26 < 12; fi26++ {
        if (491 <= 255) {
            var v27 uint32 = (v20 & 557)
            _ = v27
        }
        v28 := GenS1{F0: int32(v18), F1: ((("foo bar" + "a") + "a") + "hello"), F2: -128}
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 int32 = -354
            var v31 uint16 = 654
            v15 += v18
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi26
        _ = v28
    }
    switch 9223372036854775807 {
    case 42:
        var v32 string = "a"
        _ = v32
    case 128:
        v33 := GenS3{F0: -1, F1: 127, F2: 42}
        _ = v33
    default:
        if (!true) {
            v34 := []int{}
            _ = v34
        }
    }
    v35 := GenS2{F0: "\n\t", F1: 1}
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 string = ("\n\t" + "\n\t")
        var v38 string = v37
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v35
}
