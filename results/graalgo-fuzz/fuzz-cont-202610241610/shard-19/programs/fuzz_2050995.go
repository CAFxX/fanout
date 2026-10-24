package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int64, p1 uint, p2 int8) string {
    var v4 int32 = int32(p1)
    for fi5 := 0; fi5 < 11; fi5++ {
        if true {
            fmt.Println(v4, v4, p0)
        }
        var v6 int16 = 32767
        var v7 int8 = p2
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "foo bar"
}
func genFn1(p0 int64) uint8 {
    var v8 uint = 769
    var v9 int64 = 0
    _ = p0
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    var v3 int = 4294967295
    v3 = v3
    var v10 int = 0
    v3 += v10
    var v11 uint32 = 90
    v12 := GenS1{F0: ("\n\t" + "foo bar")}
    switch ("" + "") {
    case "a":
        var v13 float64 = (((1e10 - (3.14159 + 0.1)) / 1e10) - float64(1e10))
        _ = v13
    case "hello":
        v2[1] = (v3 % ((v10) | 1))
    }
    v10 -= -2147483648
    fmt.Println(v11)
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 string = "a"
        fmt.Println(v1, v3)
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: ("foo bar" + "a")}
    var v17 string = "foo bar"
    var v18 uint32 = (v11 + (v11 + 192))
    v18 -= v11
    if ((-2.25 + 3.14159) >= ((-2.25 + -2.25) - (1.5 - -2.25))) {
        switch v17 {
        case "hello":
            v17 = "\n\t"
        case "":
            var v19 uint8 = genFn1(1)
            _ = v19
        case "a":
            var v20 uint64 = 2
            _ = v20
        }
    } else {
        for fi21 := 0; fi21 < 11; fi21++ {
            if v1 {
                v2[(v18 + 953)] = (v10 >> ((255) & 63))
                _ = len(v2)
            }
            _ = fi21
        }
        v2[0] = v10
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        if v1 {
            for fi23 := 0; fi23 < 12; fi23++ {
                v24 := []int{}
                v2[v11] = (v10 / ((v10) | 1))
                _ = fi23
                _ = v24
            }
            v25 := []int{542, (fi22 | 32768), (fi22 >> ((65535) & 63)), (v3 + fi22)}
            _ = v25
        } else {
            v2[(v11 << ((4294967295) & 63))] = 32768
            switch ("" + (v17 + ((v17 + "") + (v17 + v17)))) {
            case "":
                var v26 int16 = 42
                _ = v26
            default:
                fmt.Println(v18)
            }
            var v27 int16 = 42
            _ = v27
        }
        if v1 {
            v11 = (v18 / ((80) | 1))
        } else {
            v1 = v1
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
