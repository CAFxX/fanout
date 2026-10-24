package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 0
    var v3 int8 = 100
    var v4 int32 = 1
    v5 := make(map[uint32]uint32)
    v6 := []int{}
    _ = len(v6)
    switch 76 {
    case -334:
        var v7 uint64 = 638
        _ = v7
    case 938:
        if (true || false) {
            if true {
                if (true && false) {
                    v4 = v4
                    v8 := GenS1{F0: v4}
                    _ = v8
                } else {
                    fmt.Println(v3)
                }
            }
            var v9 uint16 = 444
            _ = v9
        }
    default:
        v5[2] = 2147483647
    }
    var v10 int64 = -9223372036854775808
    v2 += (v10 / ((v10) | 1))
    if false {
        _ = len(v5)
        switch -9223372036854775808 {
        case 849:
            v4 = 3
        case -9223372036854775808:
            v11 := v5[2]
            _ = v11
        default:
            for fi12 := 0; fi12 < 15; fi12++ {
                v3 -= v3
                _ = fi12
            }
        }
        var v13 bool = false
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int16 = 32767
        switch fi14 {
        case 128:
            var v16 int32 = (v4 << ((0) & 63))
            _ = v16
        case -203:
            v17 := GenS1{F0: v4}
            _ = v17
        case 32768:
            v10 = -485
        default:
            fmt.Println(v2)
        }
        _ = fi14
        _ = v15
    }
    var v18 uint32 = 1
    delete(v1, "foo bar")
    if (571 > 0) {
        v18 -= (^(v18))
        var v19 string = "foo bar"
        _ = v19
    } else {
        for fi20 := 0; fi20 < 6; fi20++ {
            if false {
                var v21 int16 = 32767
                _ = v21
            }
            var v22 uint = 18446744073709551615
            _ = fi20
            _ = v22
        }
        switch (("hello" + ("foo bar" + "a")) + (("hello" + "hello") + "a")) {
        case "\n\t":
            var v23 bool = ((4294967295 >= 4294967295) || false)
            _ = v23
        default:
            for fi24 := 0; fi24 < 11; fi24++ {
                if ((v10 == 982) && (!(!(!(0.0 != (float64(1e10) - -2.25)))))) {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        v26 := GenS1{F0: v4}
                        v4 -= v4
                        fmt.Println(v10, fi24, v10)
                        v4 = 0
                        if true {
                            fmt.Println(v18)
                        } else {
                            if (-2.25 == 3.14159) {
                                v27 := GenS1{F0: -189}
                                v1[("foo bar" + ((("\n\t" + ("hello" + (("a" + "") + (("a" + ("\n\t" + ("hello" + "foo bar"))) + "\n\t")))) + (("\n\t" + "") + ("a" + "a"))) + ("a" + ("foo bar" + "\n\t"))))] = ("" + "a")
                                _ = v27
                            }
                            var v28 uint64 = 0
                            v3 = v3
                            _ = v28
                        }
                        _ = fi25
                        _ = v26
                    }
                } else {
                    v10 = (-(v10))
                }
                var v29 int16 = 402
                _ = fi24
                _ = v29
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
