package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int8) int64 {
    if true {
        if false {
            p0 += (p0 << ((18446744073709551615) & 63))
            var v5 int8 = p0
            fmt.Println(p0)
            _ = v5
        } else {
            var v6 int8 = (p0 * p0)
            fmt.Println(p0, v6)
            if (!(893 > 42)) {
                p0 -= v6
                p0 = p0
            }
            _ = v6
        }
        switch -2147483648 {
        case 722:
            if true {
                fmt.Println(p0)
            } else {
                var v7 int16 = -32768
                var v8 uint16 = 915
                _ = v7
                _ = v8
            }
        }
    } else {
        p0 -= p0
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint8 = 173
            var v11 uint32 = 942
            p0 += p0
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 int8 = p0
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            p0 += (v12 / ((v12) | 1))
            v15 := GenS1{F0: 0, F1: fi14}
            var v16 int = fi13
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 uint8 = 144
        _ = fi13
        _ = v17
    }
    v12 = -74
    var v18 float32 = -2.25
    _ = p0
    _ = v12
    _ = v18
    return -1
}

func main() {
    var v1 bool = true
    var v2 bool = false
    var v3 uint8 = 127
    v4 := make(map[int]int)
    v2 = (0 == 1000)
    var v19 int = -716
    if v1 {
        for fi20 := 0; fi20 < 15; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                v3 += v3
                v19 -= (v19 << ((uint(v3)) & 63))
                v22 := GenS2{F0: 959, F1: 122, F2: 0}
                _ = len(v4)
                _ = fi21
                _ = v22
            }
            for fi23 := 0; fi23 < 14; fi23++ {
                v24 := []int{(v19 >> ((18446744073709551615) & 63))}
                _ = fi23
                _ = v24
            }
            _ = fi20
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            switch (("\n\t" + "") + ("a" + "hello")) {
            case "":
                v1 = ((^(v3)) > 42)
            case "a":
                v26 := []int{}
                _ = v26
            case "foo bar":
                if v1 {
                    var v27 float64 = 1.5
                    _ = v27
                }
            }
            v28 := []int{v19, v19}
            v29 := GenS2{F0: 2, F1: 0, F2: 4294967295}
            _ = fi25
            _ = v28
            _ = v29
        }
        _ = len(v4)
    }
    v19 -= (v19 / (((^(v19))) | 1))
    fmt.Println(v1, v2)
    _ = len(v4)
    v30 := GenS3{F0: 0}
    v2 = v1
    v31 := v4[(v19 + v19)]
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 bool = (v1 && (32767 == -1))
        _ = fi32
        _ = v33
    }
    var v34 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
}
