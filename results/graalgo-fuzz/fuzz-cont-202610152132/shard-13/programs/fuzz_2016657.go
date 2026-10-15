package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 string }
func genFn0(p0 bool, p1 int8, p2 bool) int {
    switch 32767 {
    case 32768:
        var v6 float64 = 0.0
        _ = v6
    case 4294967295:
        fmt.Println(p0, p1, p2)
    }
    if false {
        var v7 int16 = -201
        p2 = (true && p0)
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        switch "" {
        case "hello":
            v9 := []int{fi8, fi8}
            _ = v9
        case "":
            fmt.Println(p0, p0)
        }
        v10 := []int{}
        var v11 uint64 = 18446744073709551615
        _ = fi8
        _ = v10
        _ = v11
    }
    p0 = (102 > 2)
    _ = p0
    _ = p1
    _ = p2
    return -46
}

func main() {
    var v1 int32 = -211
    v2 := []string{}
    var v3 float32 = 1e10
    var v4 uint8 = 141
    v5 := make(map[int]bool)
    v3 = v3
    v12 := GenS1{F0: 127, F1: 324}
    v13 := v5[genFn0(true, -1, (!(2 > (v4 << ((561) & 63)))))]
    var v14 string = "a"
    if (v4 >= (^(v4))) {
        switch 4294967295 {
        case 127:
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 float64 = (-2.25 + (0.1 - -2.25))
                v1 += 0
                _ = fi15
                _ = v16
            }
        case -1:
            for fi17 := 0; fi17 < 2; fi17++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch -916 {
                case 42:
                    _ = len(v2)
                }
                v18 := GenS2{F0: (v4 - 2), F1: 65535}
                _ = fi17
                _ = v18
            }
        }
        switch "hello" {
        case "a":
            fmt.Println(v4)
        }
    } else {
        var v19 float32 = v3
        var v20 uint32 = 4294967295
        for fi21 := 0; fi21 < 14; fi21++ {
            _ = len(v2)
            _ = fi21
        }
        _ = v19
        _ = v20
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v1 = (^(v1))
        v23 := GenS3{F0: ((("" + "a") + ((v14 + "foo bar") + (v14 + ((v14 + v14) + (v14 + v14))))) + "hello")}
        if (v1 != v1) {
            var v24 float64 = (1.5 / 1.5)
            var v25 int8 = 0
            var v26 bool = false
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    var v27 int = 0
    var v28 bool = (v13 && (100 > 127))
    fmt.Println(v4, v13, v13)
    v5[(v27 | (v27 * -932))] = v28
    var v29 int = 32767
    if v28 {
        if false {
            v29 = genFn0(v28, 10, (1 < 9223372036854775807))
        } else {
            v14 = v14
        }
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 int32 = 0
            _ = fi30
            _ = v31
        }
    } else {
        var v32 float32 = 0.0
        if v28 {
            var v33 uint = 42
            _ = v33
        }
        v34 := GenS3{F0: v14}
        _ = v32
        _ = v34
    }
    v3 -= 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
