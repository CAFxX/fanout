package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
func genFn0() uint {
    switch ((((("\n\t" + ("hello" + "")) + "\n\t") + "hello") + "foo bar") + (("a" + "") + ("hello" + ""))) {
    case "foo bar":
        if true {
            var v5 bool = false
            v5 = (v5 && true)
            for fi6 := 0; fi6 < 20; fi6++ {
                v5 = (53 != -2147483648)
                fmt.Println(v5, v5)
                _ = fi6
            }
            _ = v5
        }
    case "\n\t":
        v7 := GenS1{F0: -31}
        _ = v7
    }
    var v8 uint = 531
    switch 4294967295 {
    case 9223372036854775807:
        v9 := []int{-2147483648}
        _ = v9
    case -1:
        for fi10 := 0; fi10 < 12; fi10++ {
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 uint8 = 222
                var v13 int16 = 1000
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 bool = true
                _ = fi14
                _ = v15
            }
            fmt.Println(v8)
            _ = fi10
        }
    case 32768:
        var v16 uint = 255
        _ = v16
    default:
        switch "a" {
        case "":
            v17 := []int{-880, 9223372036854775807, -821, 127}
            _ = v17
        }
    }
    if (1.5 == float64(1.5)) {
        v8 += 255
    } else {
        v18 := []int{-1, 2147483647, -9223372036854775808, 257}
        _ = v18
    }
    _ = v8
    return (v8 - v8)
}

func main() {
    v1 := []uint32{65535, 4294967295}
    v2 := []int{32768, 268}
    v3 := make(map[int]uint32)
    var v4 float32 = 0.1
    v19 := GenS2{F0: 914}
    if true {
        if ((v4 + -2.25) <= v4) {
            fmt.Println(v4)
            v4 = 0.1
        }
    } else {
        v4 -= (1.5 * 1.5)
        switch "hello" {
        case "foo bar":
            var v20 uint = 65535
            _ = v20
        case "":
            fmt.Println(v4)
        case "\n\t":
            var v21 string = "a"
            _ = v21
        }
        var v22 uint8 = 127
        _ = v22
    }
    fmt.Println(v4)
    _ = len(v3)
    fmt.Println(v4)
    fmt.Println(v4)
    fmt.Println(v4)
    if (705 == 127) {
        var v23 int32 = 0
        switch -2147483648 {
        case 9223372036854775807:
            var v24 uint16 = 937
            _ = v24
        case 128:
            v23 -= v23
        case 4294967295:
            v2 = append(v2, -168)
        default:
            v25 := GenS2{F0: 42}
            _ = v25
        }
        _ = v23
    }
    if true {
        v26 := GenS1{F0: 42}
        var v27 bool = (526 != 65535)
        _ = v26
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
