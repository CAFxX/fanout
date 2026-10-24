package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 string; F2 string }

func main() {
    var v1 string = "a"
    var v2 uint = 1
    v3 := make(map[string]uint32)
    var v4 int = 639
    for fi5 := 0; fi5 < 3; fi5++ {
        v4 -= 0
        v2 = v2
        _ = fi5
    }
    if true {
        switch v1 {
        case "":
            v2 = v2
        case "foo bar":
            switch (v1 + (((v1 + v1) + v1) + (((v1 + v1) + "a") + "a"))) {
            case "":
                if true {
                    var v6 bool = true
                    _ = v6
                } else {
                    if true {
                        fmt.Println(v2, v1)
                        var v7 uint = v2
                        _ = v7
                    }
                    v4 = 561
                }
            case "a":
                for fi8 := 0; fi8 < 3; fi8++ {
                    v2 -= v2
                    v4 -= fi8
                    _ = fi8
                }
            case "hello":
                if (954 != (-(v4))) {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        _ = len(v3)
                        for fi10 := 0; fi10 < 17; fi10++ {
                            v11 := GenS1{F0: 532}
                            _ = fi10
                            _ = v11
                        }
                        _ = fi9
                    }
                }
            }
        default:
            v4 += v4
        }
        v4 = (v4 & 1)
        v12 := v3[""]
        _ = v12
    }
    fmt.Println(v1, v2)
    if false {
        v13 := GenS1{F0: 65535}
        _ = v13
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS1{F0: 65535}
            if ((3.14159 < 0.0) && ((true && false) && ((4294967295 <= 2147483648) || false))) {
                if (65535 == 65535) {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        switch (v4 | fi14) {
                        case 32768:
                            fmt.Println(v4, v1, v2)
                        }
                        _ = fi16
                    }
                    var v17 float64 = (3.14159 * 0.0)
                    _ = v17
                }
                if true {
                    v18 := []int{(v4 & (v4 | 1)), v4, -63}
                    if true {
                        v19 := GenS1{F0: 567}
                        var v20 uint64 = 804
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v4 += fi14
                            if (!(!(true && (v20 >= v20)))) {
                                if (11 >= 0) {
                                    v1 = ""
                                }
                                v20 += v20
                            } else {
                                v2 -= v2
                            }
                            _ = fi21
                        }
                        _ = v19
                        _ = v20
                    } else {
                        var v22 float64 = 0.1
                        delete(v3, "a")
                        _ = v22
                    }
                    _ = v18
                }
                var v23 int = 1
                _ = v23
            }
            v24 := GenS2{F0: 42, F1: v1, F2: ("\n\t" + "hello")}
            _ = fi14
            _ = v15
            _ = v24
        }
    }
    var v25 int32 = -358
    var v26 uint64 = 4294967295
    v27 := GenS1{F0: 0}
    if (!true) {
        v28 := GenS2{F0: 42, F1: (("foo bar" + v1) + "\n\t"), F2: (v1 + (v1 + ((v1 + v1) + v1)))}
        v29 := GenS1{F0: 65535}
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
