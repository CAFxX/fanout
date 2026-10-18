package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[string]bool)
    var v3 uint32 = 2147483648
    v4 := make(map[string]int)
    var v5 string = ("foo bar" + (("foo bar" + "hello") + (("\n\t" + "foo bar") + "\n\t")))
    v1 -= v3
    v3 -= 42
    if (v1 != (v1 % (((^(v1))) | 1))) {
        var v6 uint16 = 255
        var v7 string = v5
        _ = v6
        _ = v7
    } else {
        v8 := []int{}
        v9 := GenS1{F0: v3, F1: ((0 > -1) || false)}
        _ = v8
        _ = v9
    }
    v1 += 1
    var v10 int16 = 1000
    var v11 uint16 = 42
    if ((4294967295 == 834) || ((v3 * 897) <= 2)) {
        var v12 uint64 = 173
        v13 := GenS1{F0: 2147483647, F1: false}
        _ = v12
        _ = v13
    } else {
        var v14 uint64 = 298
        v15 := v2[v5]
        v10 += v10
        _ = v14
        _ = v15
    }
    v1 -= (v1 ^ 817)
    for fi16 := 0; fi16 < 12; fi16++ {
        if (float64(0.1) <= 0.1) {
            var v17 uint8 = 128
            for fi18 := 0; fi18 < 16; fi18++ {
                v1 = 42
                _ = fi18
            }
            v17 = (v17 >> ((2) & 63))
            _ = v17
        }
        if true {
            var v19 string = ((v5 + v5) + "foo bar")
            fmt.Println(v10, v10)
            if false {
                var v20 uint = 65535
                var v21 uint32 = uint32(v11)
                _ = v20
                _ = v21
            }
            _ = v19
        } else {
            v10 -= 0
        }
        _ = fi16
    }
    var v22 uint64 = 2
    if (!false) {
        switch (v5 + v5) {
        case "foo bar":
            if (!true) {
                for fi23 := 0; fi23 < 20; fi23++ {
                    for fi24 := 0; fi24 < 19; fi24++ {
                        if true {
                            v10 = v10
                            if (!(false || false)) {
                                fmt.Println(v11)
                            } else {
                                v11 = v11
                            }
                            var v25 string = ("foo bar" + "foo bar")
                            _ = v25
                        } else {
                            var v26 float64 = float64(0.0)
                            var v27 float32 = (0.1 * (1e10 - 0.1))
                            _ = v26
                            _ = v27
                        }
                        v10 = 1
                        _ = fi24
                    }
                    var v28 uint16 = (v11 + (^(v11)))
                    _ = fi23
                    _ = v28
                }
                var v29 int64 = 1
                _ = v29
            } else {
                var v30 int64 = 328
                _ = v30
            }
        default:
            if true {
                var v31 float32 = 1e10
                _ = v31
            }
        }
        for fi32 := 0; fi32 < 5; fi32++ {
            var v33 uint8 = 0
            var v34 float32 = -2.25
            _ = fi32
            _ = v33
            _ = v34
        }
        if false {
            var v35 uint8 = 42
            v22 = v22
            var v36 float32 = 0.1
            _ = v35
            _ = v36
        } else {
            v1 -= 2
        }
    }
    switch 498 {
    case 939:
        switch v5 {
        case "a":
            fmt.Println(v3)
        default:
            fmt.Println(v1, v3, v22)
        }
    case 4294967295:
        v2[(v5 + v5)] = (!(("\n\t" <= v5) || (!((0 != v11) && false))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
