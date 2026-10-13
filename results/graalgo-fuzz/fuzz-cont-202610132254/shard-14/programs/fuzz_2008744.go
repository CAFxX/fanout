package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16; F1 int16; F2 string }
func genFn0(p0 bool, p1 int8) uint8 {
    p1 = p1
    p0 = p0
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(fi5, p1)
        _ = fi5
    }
    _ = p0
    _ = p1
    return 188
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 float64 = 0.1
    var v4 uint32 = 917
    var v6 int32 = -234
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int8 = -1
        _ = fi8
        _ = v9
    }
    var v10 bool = false
    if v10 {
        if (!false) {
            var v11 uint16 = 42
            _ = v11
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            switch fi12 {
            case 387:
                switch "\n\t" {
                case "":
                    v13 := []int{(-(fi12)), (fi12 ^ fi12), int(v6), fi12}
                    _ = v13
                }
            case -9223372036854775808:
                fmt.Println(v4, v7, v10)
            default:
                v6 += (v6 & 1)
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := GenS1{F0: 65535, F1: 65535, F2: (v4 >> ((42) & 63))}
                v4 -= (v4 | 784)
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 14; fi16++ {
                v3 = v3
                v17 := []int{}
                var v18 uint64 = v7
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi12
        }
        if false {
            switch (("\n\t" + ("foo bar" + "\n\t")) + "foo bar") {
            case "foo bar":
                v19 := GenS1{F0: 295, F1: 968, F2: v4}
                _ = v19
            default:
                var v20 string = "\n\t"
                _ = v20
            }
            fmt.Println(v3, v10, v10)
        } else {
            v21 := GenS1{F0: 1, F1: 255, F2: v4}
            var v22 int16 = 32767
            v7 -= (v7 / (((v7 << ((uint(v4)) & 63))) | 1))
            _ = v21
            _ = v22
        }
    }
    delete(v1, v4)
    if v10 {
        for fi23 := 0; fi23 < 19; fi23++ {
            fmt.Println(v6, fi23)
            v24 := []int{}
            _ = fi23
            _ = v24
        }
        v7 += (v7 ^ (v7 - v7))
        v4 = v4
    }
    var v25 uint64 = v7
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 float32 = (float32(v3) - (-2.25 - 0.1))
        v4 -= (v4 << ((521) & 63))
        var v28 uint = 65535
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 int32 = -2147483648
    v30 := GenS1{F0: uint16(v7), F1: 0, F2: (v4 ^ (v4 << ((452) & 63)))}
    if (false || v10) {
        v31 := GenS1{F0: 85, F1: 290, F2: v4}
        v7 -= 18446744073709551615
        for fi32 := 0; fi32 < 3; fi32++ {
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 int16 = -1
                delete(v2, (-(fi32)))
                _ = fi33
                _ = v34
            }
            v10 = ("foo bar" <= "hello")
            _ = fi32
        }
        _ = v31
    } else {
        v2[0] = -550
        v7 += (^(v25))
        v35 := []int{13, 128, 32768, -2147483648}
        _ = v35
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        v37 := GenS2{F0: -203, F1: -1, F2: (("a" + "a") + ("hello" + "hello"))}
        switch (fi36 + -9223372036854775808) {
        case 816:
            if (3.14159 < 1e10) {
                if false {
                    v38 := []int{30}
                    _ = v38
                } else {
                    v39 := []int{902, fi36}
                    _ = v39
                }
                var v40 float64 = float64(1e10)
                _ = v40
            }
        case 4294967295:
            var v41 uint = 86
            _ = v41
        case -1:
            v25 = (v25 * v25)
        }
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v6
    _ = v7
}
