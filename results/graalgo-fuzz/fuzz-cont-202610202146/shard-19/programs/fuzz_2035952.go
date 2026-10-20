package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int = 75
    var v2 float32 = float32(0.0)
    var v3 int64 = 9223372036854775807
    var v4 uint8 = 156
    var v5 float32 = v2
    var v6 uint64 = 4294967295
    var v7 int16 = 32767
    if false {
        var v8 int = (v1 ^ (v1 & (v1 % ((-473) | 1))))
        _ = v8
    } else {
        if false {
            var v9 float32 = v2
            _ = v9
        }
    }
    fmt.Println(v6, v7)
    for fi10 := 0; fi10 < 15; fi10++ {
        switch "a" {
        case "a":
            switch (("hello" + "foo bar") + "") {
            case "a":
                var v11 uint = 847
                _ = v11
            case "foo bar":
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 int8 = 42
                    v5 = float32(-2.25)
                    _ = fi12
                    _ = v13
                }
            }
        default:
            var v14 uint8 = (v4 ^ 222)
            _ = v14
        }
        fmt.Println(v2, v6, v2)
        fmt.Println(v3, v1)
        _ = fi10
    }
    if (((v2 / ((0.1 * (float32(0.0) + 0.1)) * -2.25)) != v5) && (!false)) {
        v6 = v6
    } else {
        var v15 uint32 = uint32(v6)
        v16 := GenS1{F0: 42}
        _ = v15
        _ = v16
    }
    var v17 uint32 = 26
    if true {
        var v18 uint16 = 0
        var v19 uint = 55
        if (false && ((!false) && true)) {
            v17 = v17
        }
        _ = v18
        _ = v19
    } else {
        v20 := []int{}
        for fi21 := 0; fi21 < 3; fi21++ {
            for fi22 := 0; fi22 < 7; fi22++ {
                if false {
                    v17 = 2
                    var v23 int32 = 42
                    _ = v23
                } else {
                    if false {
                        v2 = 0.1
                        v24 := GenS1{F0: 0}
                        switch "" {
                        case "a":
                            var v25 bool = false
                            _ = v25
                        case "":
                            v4 = uint8(v17)
                        case "foo bar":
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v6, v4)
                            default:
                                v2 += float32(3.14159)
                            }
                        default:
                            v2 = v2
                        }
                        _ = v24
                    } else {
                        v5 -= (1.5 * 0.1)
                    }
                    var v26 uint = 65535
                    v27 := GenS2{F0: (true && false)}
                    fmt.Println(v17, v6)
                    _ = v26
                    _ = v27
                }
                _ = fi22
            }
            v4 += (v4 / ((42) | 1))
            _ = fi21
        }
        if (42 == 0) {
            var v28 float64 = 0.1
            _ = v28
        } else {
            var v29 bool = false
            if len(v20) > 0 {
                v20[0] = (v1 / ((v1) | 1))
            }
            _ = v29
        }
        _ = v20
    }
    v30 := []int{int(v4), 42}
    fmt.Println(v5)
    var v31 int64 = v3
    for fi32 := 0; fi32 < 18; fi32++ {
        if len(v30) > 0 {
            v30[0] = (fi32 | -2147483648)
        }
        v3 -= 9223372036854775807
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
