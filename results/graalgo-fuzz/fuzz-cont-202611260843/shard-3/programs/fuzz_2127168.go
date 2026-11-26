package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0() int {
    var v6 float32 = (1e10 / 1.0)
    var v7 float32 = v6
    v7 = 0.0
    _ = v6
    _ = v7
    return -936
}

func main() {
    var v1 bool = true
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]string)
    var v4 uint16 = 32767
    var v5 string = (("a" + ("\n\t" + "foo bar")) + "hello")
    switch -9223372036854775808 {
    case 4294967295:
        if (!(2 > 2)) {
            var v8 int64 = (-(v2))
            switch (("foo bar" + v5) + "a") {
            case "hello":
                for fi9 := 0; fi9 < 17; fi9++ {
                    var v10 uint16 = v4
                    v4 = (v4 << ((2) & 63))
                    if v1 {
                        fmt.Println(v2, v2, v10)
                        for fi11 := 0; fi11 < 4; fi11++ {
                            var v12 uint32 = 455
                            var v13 string = "\n\t"
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                    }
                    _ = fi9
                    _ = v10
                }
            case "foo bar":
                if false {
                    v8 = (v2 ^ v8)
                    var v14 float32 = 1e10
                    var v15 uint8 = 0
                    _ = v14
                    _ = v15
                } else {
                    v8 -= -111
                    switch (v5 + (v5 + v5)) {
                    case "foo bar":
                        var v16 bool = (0.1 == 0.0)
                        _ = v16
                    case "\n\t":
                        if (100 < 42) {
                            var v17 float32 = (float32(3.14159) / 1e10)
                            delete(v3, (((v5 + "foo bar") + "a") + v5))
                            _ = v17
                        } else {
                            delete(v3, ((((v5 + (v5 + ((v5 + v5) + ("" + v5)))) + ((("foo bar" + "foo bar") + v5) + (v5 + ((v5 + v5) + "hello")))) + (v5 + v5)) + v5))
                            v18 := GenS1{F0: "", F1: v1}
                            var v19 int64 = -557
                            _ = v18
                            _ = v19
                        }
                    case "a":
                        for fi20 := 0; fi20 < 6; fi20++ {
                            var v21 int8 = -1
                            for fi22 := 0; fi22 < 20; fi22++ {
                                var v23 float64 = 1.5
                                var v24 uint64 = 0
                                v8 = (v8 ^ (v2 + (v2 << ((42) & 63))))
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            var v25 int16 = 336
                            _ = fi20
                            _ = v21
                            _ = v25
                        }
                    }
                    v4 = (v4 & (v4 * v4))
                }
            case "\n\t":
                fmt.Println(v2, v4)
            }
            _ = v8
        }
    }
    v1 = v1
    v4 = (v4 / ((v4) | 1))
    v5 = "\n\t"
    for fi26 := 0; fi26 < 16; fi26++ {
        _ = len(v3)
        _ = fi26
    }
    var v27 int32 = 0
    if (((1.5 + (((float32((0.1 * float64(0.0))) / 1.5) + 3.14159) + 3.14159)) / (-2.25 * (0.0 - float32(1e10)))) > 0.0) {
        v2 += 761
        fmt.Println(v5, v27, v5)
    }
    v1 = v1
    var v28 int16 = -32768
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
