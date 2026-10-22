package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0() int64 {
    var v4 string = "\n\t"
    var v5 uint32 = 0
    v6 := GenS1{F0: -1, F1: 465}
    var v7 float32 = 3.14159
    if (true || true) {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int16 = 32767
            _ = fi8
            _ = v9
        }
    }
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 9223372036854775807
}
func genFn1() uint16 {
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 int16 = 32767
            switch ("hello" + "hello") {
            case "\n\t":
                if true {
                    var v13 uint = 255
                    _ = v13
                }
            case "":
                v14 := GenS1{F0: -904, F1: 713}
                _ = v14
            case "foo bar":
                var v15 float32 = float32((0.0 * 0.1))
                _ = v15
            }
            var v16 uint = 42
            _ = fi11
            _ = v12
            _ = v16
        }
        _ = fi10
    }
    v17 := []int{32767}
    var v18 uint16 = 255
    v18 -= (v18 << ((4294967295) & 63))
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    _ = v17
    _ = v18
    return 32767
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 1
    var v3 string = "\n\t"
    v2 += 876
    v3 = ("" + "")
    v19 := GenS1{F0: 561, F1: 4294967295}
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := GenS1{F0: -2147483648, F1: 42}
            fmt.Println(v2, fi21, fi21)
            _ = fi21
            _ = v22
        }
        v23 := GenS1{F0: -846, F1: 42}
        _ = fi20
        _ = v23
    }
    fmt.Println(v3, v2)
    var v24 int32 = -640
    var v25 string = (v3 + v3)
    v25 = v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
