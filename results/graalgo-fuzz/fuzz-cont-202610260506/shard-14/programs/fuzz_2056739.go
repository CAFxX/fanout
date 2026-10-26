package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int8 {
    if (18446744073709551615 > 1) {
        var v7 uint8 = 255
        v8 := []int{707, -1}
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 18446744073709551615}
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(fi10)
        var v11 uint32 = 42
        _ = fi10
        _ = v11
    }
    _ = v9
    return 102
}
func genFn1(p0 int16, p1 uint8) uint64 {
    var v12 int8 = -128
    var v13 uint16 = 717
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    var v1 int8 = -44
    var v2 int32 = -1
    var v3 uint = 4294967295
    v4 := []string{"hello"}
    var v5 float64 = (-2.25 * (1.5 * ((1e10 / ((1.5 / float64(float32(1.5))) / 1e10)) + 0.0)))
    var v6 string = ("" + "a")
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 float32 = float32(0.1)
    var v15 string = ("foo bar" + v6)
    for fi16 := 0; fi16 < 10; fi16++ {
        switch fi16 {
        case 127:
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint32 = 65535
                _ = fi17
                _ = v18
            }
        case -670:
            var v19 bool = true
            _ = v19
        }
        _ = fi16
    }
    switch v6 {
    case "a":
        var v20 int8 = 1
        _ = v20
    case "hello":
        for fi21 := 0; fi21 < 9; fi21++ {
            v6 = ("hello" + v6)
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 int16 = 1000
                var v24 int16 = v23
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi21
        }
    default:
        switch "\n\t" {
        case "foo bar":
            v3 -= 2
        case "\n\t":
            switch 496 {
            case -9223372036854775808:
                v14 -= 1.5
            case 601:
                if (!(54 > v1)) {
                    if false {
                        fmt.Println(v15, v1)
                        fmt.Println(v15, v5)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        if true {
                            var v25 uint = v3
                            var v26 int32 = v2
                            _ = v25
                            _ = v26
                        }
                        var v27 uint64 = 407
                        _ = v27
                    }
                    var v28 float32 = v14
                    _ = v28
                }
            }
        }
    }
    v6 = "\n\t"
    v14 -= v14
    var v29 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
