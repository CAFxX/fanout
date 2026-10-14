package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int32 }

func main() {
    var v1 bool = true
    var v2 uint8 = 56
    var v3 int16 = 42
    var v4 int64 = -1
    var v5 string = (("\n\t" + "hello") + (("hello" + ("a" + "a")) + "a"))
    var v6 int64 = v4
    v5 = v5
    v3 += 32767
    if (65535 > 473) {
        if v1 {
            var v7 uint64 = 4294967295
            fmt.Println(v4)
            var v8 bool = v1
            _ = v7
            _ = v8
        }
        v2 -= v2
        switch (("foo bar" + ("\n\t" + "hello")) + v5) {
        case "hello":
            if (v1 && (true && (1 >= 186))) {
                v6 += -9223372036854775808
                v1 = (0 == 4294967295)
                fmt.Println(v1, v5)
            }
        default:
            v9 := []int{-9223372036854775808, 626, 322, int(v2)}
            _ = v9
        }
    } else {
        fmt.Println(v2)
        if ((v1 || (42 <= (^(v6)))) && v1) {
            if true {
                for fi10 := 0; fi10 < 15; fi10++ {
                    for fi11 := 0; fi11 < 5; fi11++ {
                        v12 := GenS1{F0: (v2 << ((65535) & 63))}
                        v13 := GenS1{F0: v2}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = fi10
                }
            }
            fmt.Println(v1)
            var v14 float64 = 3.14159
            _ = v14
        } else {
            var v15 uint64 = 4294967295
            _ = v15
        }
        var v16 float64 = -2.25
        _ = v16
    }
    var v17 uint16 = 42
    var v18 string = ("\n\t" + v5)
    if (!(0.0 >= (1e10 + float64((-2.25 * 0.1))))) {
        v5 = (v18 + "a")
        v18 = "\n\t"
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := GenS2{F0: -83, F1: 0}
            var v21 uint64 = 4294967295
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        switch 4294967295 {
        case 42:
            var v22 string = "foo bar"
            _ = v22
        case 33:
            switch ("" + ("\n\t" + (v5 + "hello"))) {
            case "a":
                v2 = v2
            default:
                for fi23 := 0; fi23 < 20; fi23++ {
                    v24 := []int{(fi23 * fi23), 42}
                    for fi25 := 0; fi25 < 14; fi25++ {
                        var v26 int = int(v6)
                        _ = len(v24)
                        _ = fi25
                        _ = v26
                    }
                    var v27 int16 = -1
                    _ = fi23
                    _ = v24
                    _ = v27
                }
            }
        default:
            var v28 float64 = ((1e10 + float64(-2.25)) + float64((float32(-2.25) / 1.0)))
            _ = v28
        }
        v29 := GenS2{F0: -1, F1: 388}
        var v30 uint32 = 509
        _ = v29
        _ = v30
    }
    var v31 int64 = int64(v2)
    if true {
        v6 = 0
        if false {
            v18 = (v18 + "foo bar")
            v5 = (v5 + (v5 + "\n\t"))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
