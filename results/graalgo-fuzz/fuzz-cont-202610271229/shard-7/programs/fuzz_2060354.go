package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 bool }
func genFn0() uint32 {
    switch -393 {
    case 2147483647:
        var v7 uint32 = 2147483648
        _ = v7
    case 127:
        if (2 != 128) {
            var v8 float64 = 0.0
            if false {
                v9 := []int{32768, 32768, 223, 104}
                if (-652 <= -1) {
                    var v10 int16 = 42
                    _ = v10
                }
                fmt.Println(v8)
                _ = v9
            }
            _ = v8
        }
    case -289:
        if true {
            v11 := GenS1{F0: -1, F1: "\n\t", F2: true}
            var v12 uint64 = 0
            _ = v11
            _ = v12
        } else {
            v13 := []int{-653, 0, -515}
            _ = v13
        }
    }
    var v14 float64 = float64(0.0)
    _ = v14
    return 2147483647
}

func main() {
    var v1 float32 = -2.25
    var v2 uint32 = 65535
    var v3 int = 42
    v4 := make(map[int]string)
    var v5 uint8 = 42
    v6 := make(map[string]int)
    fmt.Println(v5, v2, v3)
    if (80 <= 42) {
        var v15 uint64 = uint64(v2)
        _ = v15
    }
    var v16 bool = (2147483647 > 508)
    var v17 uint16 = 0
    var v18 int32 = -843
    v5 = v5
    v16 = ((1 >= -1) || ((1 > 1) || ((-(v18)) == 249)))
    v19 := []int{v3, v3, (v3 | v3), 32768}
    v2 += uint32(v17)
    for fi20 := 0; fi20 < 13; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            if v16 {
                v22 := GenS1{F0: -128, F1: "a", F2: v16}
                var v23 int8 = -1
                fmt.Println(v1, v18, v17)
                _ = v22
                _ = v23
            }
            if v16 {
                var v24 float32 = ((float32((-2.25 / -2.25)) - 1.5) + 0.0)
                v25 := GenS1{F0: 49, F1: "foo bar", F2: (!(223 >= 255))}
                _ = v24
                _ = v25
            } else {
                v16 = true
                v26 := GenS1{F0: 100, F1: (("foo bar" + "hello") + ("foo bar" + ("\n\t" + (("a" + ((("hello" + "foo bar") + "") + ((("hello" + ("hello" + "\n\t")) + ("hello" + "foo bar")) + "hello"))) + "\n\t")))), F2: ((-1 != 1) || ((v1 * ((v1 - 0.1) / -2.25)) < v1))}
                v27 := GenS1{F0: 100, F1: ("hello" + "foo bar"), F2: ((genFn0() != (v2 - (v2 << ((271) & 63)))) && ((0 < (v3 - -516)) || ((!false) || (v16 || v16))))}
                _ = v26
                _ = v27
            }
            _ = fi21
        }
        _ = fi20
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        for fi29 := 0; fi29 < 19; fi29++ {
            if v16 {
                v30 := GenS1{F0: 0, F1: "", F2: ((v16 || false) && (uint(v17) == 0))}
                _ = v30
            }
            var v31 bool = v16
            v32 := GenS1{F0: -128, F1: "foo bar", F2: v31}
            _ = fi29
            _ = v31
            _ = v32
        }
        _ = len(v19)
        if v16 {
            fmt.Println(v5)
            var v33 string = "a"
            _ = v33
        } else {
            v16 = v16
        }
        _ = fi28
    }
    v2 -= (v2 >> ((53) & 63))
    v34 := GenS1{F0: -1, F1: "", F2: (1 < 42)}
    var v35 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
