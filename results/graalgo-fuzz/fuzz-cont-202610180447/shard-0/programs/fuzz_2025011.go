package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 string }
func genFn0() int8 {
    v6 := GenS1{F0: ("" + "foo bar"), F1: 2, F2: 618}
    v7 := []int{42, 127}
    var v8 int16 = 0
    v9 := GenS1{F0: "foo bar", F1: 978, F2: 1}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int16, p1 int32) uint64 {
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(p1, fi10, p0)
        _ = fi10
    }
    if false {
        var v11 float32 = (float32(1.5) / float32(1e10))
        p0 -= p0
        _ = v11
    }
    var v12 uint64 = 2
    if ("\n\t" != "") {
        p1 = (p1 % (((-(p1))) | 1))
        v13 := []int{}
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 string = ""
                var v17 string = (v16 + v16)
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
        _ = v13
    } else {
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int8 = 100
            _ = fi18
            _ = v19
        }
    }
    v20 := []int{2147483647}
    _ = p0
    _ = p1
    _ = v12
    _ = v20
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 1
    var v3 uint64 = 42
    v4 := []int{448, 128}
    var v5 float32 = 0.0
    fmt.Println(v5)
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 uint64 = v3
        v5 -= 3.14159
        v23 := GenS2{F0: "a"}
        _ = fi21
        _ = v22
        _ = v23
    }
    if ((1.5 / 1.0) == -2.25) {
        var v24 float32 = v5
        var v25 int8 = -1
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        v1[(("a" + "") + "")] = false
        if false {
            if (false && ((v2 * (v2 | (v2 & v2))) <= (v2 + (v2 - v2)))) {
                for fi27 := 0; fi27 < 18; fi27++ {
                    v28 := GenS2{F0: "\n\t"}
                    v3 += (^(v3))
                    if (uint8(v2) == 255) {
                        fmt.Println(fi26, fi26, v2)
                        var v29 float32 = 1e10
                        v5 += v5
                        _ = v29
                    }
                    v5 = (0.1 + 0.1)
                    var v30 float32 = 0.0
                    var v31 uint16 = (v2 * v2)
                    _ = fi27
                    _ = v28
                    _ = v30
                    _ = v31
                }
                var v32 int = fi26
                _ = v32
            }
        } else {
            v3 = 14
        }
        _ = fi26
    }
    switch 1 {
    case -86:
        v3 = v3
    case -295:
        switch 9223372036854775807 {
        case -9223372036854775808:
            v3 = (v3 + v3)
        default:
            var v33 uint64 = (v3 & (v3 - (^(v3))))
            _ = v33
        }
    case 555:
        fmt.Println(v2)
    }
    fmt.Println(v2, v3)
    fmt.Println(v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v34 uint64 = genFn1(-1, 1)
    var v35 uint8 = 120
    var v36 uint = 255
    for fi37 := 0; fi37 < 2; fi37++ {
        v38 := GenS1{F0: "\n\t", F1: 2147483647, F2: uint(v35)}
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        switch "" {
        case "a":
            switch -2147483648 {
            case 127:
                for fi40 := 0; fi40 < 8; fi40++ {
                    var v41 uint64 = v3
                    var v42 uint64 = v3
                    var v43 int16 = 32767
                    _ = fi40
                    _ = v41
                    _ = v42
                    _ = v43
                }
            }
        }
        v44 := GenS1{F0: "\n\t", F1: 2147483647, F2: (v36 + (v36 * (v36 ^ (v36 + v36))))}
        _ = fi39
        _ = v44
    }
    v4 = append(v4, 42)
    if len(v4) > 0 {
        v4[0] = 946
    }
    v1[("hello" + "a")] = true
    v45 := GenS1{F0: "foo bar", F1: uint32(v36), F2: v36}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v45
    _ = v5
}
