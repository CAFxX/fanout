package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int8, p1 int32) string {
    v5 := []int{-9223372036854775808, -2147483648, 0}
    switch 974 {
    case 32767:
        _ = len(v5)
    default:
        for fi6 := 0; fi6 < 1; fi6++ {
            p1 += p1
            _ = fi6
        }
    }
    p1 = (p1 >> ((2) & 63))
    if (0 < 2) {
        v7 := GenS1{F0: 42}
        fmt.Println(p1, p1)
        switch 32767 {
        case 9223372036854775807:
            var v8 uint = 0
            _ = v8
        }
        _ = v7
    } else {
        var v9 int16 = 935
        if false {
            switch ("foo bar" + ("\n\t" + "foo bar")) {
            case "hello":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "":
                if len(v5) > 0 {
                    v5[0] = 462
                }
            default:
                switch 32768 {
                case -2147483648:
                    p0 -= p0
                case 2147483647:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    if ((^(v9)) > (v9 >> ((342) & 63))) {
                        if (-1 < -9223372036854775808) {
                            for fi10 := 0; fi10 < 5; fi10++ {
                                var v11 string = ("\n\t" + ("\n\t" + (("hello" + "\n\t") + "foo bar")))
                                _ = fi10
                                _ = v11
                            }
                            var v12 int64 = 42
                            _ = v12
                        }
                    } else {
                        var v13 float64 = (0.0 / (0.1 - -2.25))
                        _ = v13
                    }
                }
            }
        }
        _ = v9
    }
    var v14 uint = 4294967295
    _ = p0
    _ = p1
    _ = v14
    _ = v5
    return "\n\t"
}

func main() {
    var v1 uint = 1
    v2 := make(map[int]string)
    var v3 uint32 = 596
    var v4 uint16 = 255
    var v15 bool = (42 >= 42)
    v15 = true
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := []int{(fi16 | fi16), fi16}
        _ = fi16
        _ = v17
    }
    var v18 float64 = 1e10
    v3 -= v3
    var v19 int16 = -608
    v18 += v18
    var v20 int32 = -654
    for fi21 := 0; fi21 < 1; fi21++ {
        delete(v2, (^(fi21)))
        if v15 {
            _ = len(v2)
            var v22 bool = v15
            _ = v22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            delete(v2, 2147483647)
            _ = fi23
        }
        _ = fi21
    }
    v24 := []int{}
    v4 = v4
    if v15 {
        if (1 < (v1 << ((732) & 63))) {
            delete(v2, 4294967295)
            var v25 int32 = v20
            _ = v25
        }
    } else {
        var v26 string = (("foo bar" + "\n\t") + "\n\t")
        for fi27 := 0; fi27 < 12; fi27++ {
            var v28 uint64 = 42
            v19 -= v19
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
