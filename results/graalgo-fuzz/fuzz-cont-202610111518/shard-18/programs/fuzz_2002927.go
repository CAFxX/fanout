package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 string; F1 int }
func genFn0(p0 uint32, p1 uint, p2 uint16) bool {
    v7 := []int{109, -954, -14}
    fmt.Println(p0, p0, p0)
    v8 := GenS1{F0: (p2 ^ (^(p2))), F1: 196}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (!false)
}

func main() {
    var v1 bool = (!false)
    var v2 float64 = 1e10
    v3 := make(map[uint32]int)
    var v4 uint16 = 255
    var v5 int64 = -249
    var v6 uint32 = 2
    for fi9 := 0; fi9 < 20; fi9++ {
        if v1 {
            v4 = v4
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := []int{-1, fi10, 32768, (-(fi10))}
                _ = fi10
                _ = v11
            }
        } else {
            var v12 int = 0
            _ = v12
        }
        var v13 string = ("\n\t" + ("hello" + "foo bar"))
        _ = fi9
        _ = v13
    }
    fmt.Println(v5, v4, v5)
    for fi14 := 0; fi14 < 19; fi14++ {
        if false {
            v15 := GenS1{F0: 32767, F1: 0}
            _ = v15
        } else {
            var v16 int = fi14
            fmt.Println(v5, v1, v2)
            _ = v16
        }
        var v17 uint8 = 128
        _ = fi14
        _ = v17
    }
    var v18 int32 = 2147483647
    v19 := GenS2{F0: "foo bar", F1: 2147483647}
    delete(v3, v6)
    var v20 int = 32767
    if v1 {
        v6 = v6
    } else {
        for fi21 := 0; fi21 < 9; fi21++ {
            v22 := GenS2{F0: ("hello" + "hello"), F1: fi21}
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            _ = len(v3)
            if ((fi23 ^ int(v18)) != 127) {
                if v1 {
                    v24 := GenS2{F0: "\n\t", F1: fi23}
                    v25 := GenS2{F0: ("hello" + ("\n\t" + "")), F1: v20}
                    _ = v24
                    _ = v25
                }
                var v26 int8 = 18
                _ = v26
            }
            var v27 uint = 42
            _ = fi23
            _ = v27
        }
        var v28 int32 = 1
        _ = v28
    }
    if ((v5 - v5) == v5) {
        fmt.Println(v18)
        var v29 string = ("" + ("" + "foo bar"))
        _ = v29
    } else {
        fmt.Println(v18, v2, v1)
        var v30 float32 = 1.5
        _ = v30
    }
    v6 += 2
    switch (v20 / (((v20 + (^(v20)))) | 1)) {
    case -952:
        v31 := GenS1{F0: (^(v4)), F1: 333}
        _ = v31
    case -416:
        v1 = false
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
