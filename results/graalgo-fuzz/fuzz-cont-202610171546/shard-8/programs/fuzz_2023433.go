package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (float32((-2.25 - (float64((1e10 + (1e10 + 1e10))) + (1.5 / 3.14159)))) + (3.14159 / 3.14159))
    v3 := []int{}
    var v4 float32 = (-2.25 * (v2 + v2))
    switch "foo bar" {
    case "\n\t":
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 float32 = (0.0 + 1.5)
            _ = fi5
            _ = v6
        }
    case "a":
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(v4)
            v1[-9223372036854775808] = 2147483647
            _ = fi7
        }
    default:
        var v8 int32 = -2147483648
        _ = v8
    }
    var v9 uint16 = 32767
    if false {
        v4 = v2
        switch -946 {
        case -573:
            v4 += 0.0
        case -638:
            for fi10 := 0; fi10 < 20; fi10++ {
                switch "hello" {
                case "\n\t":
                    delete(v1, fi10)
                default:
                    switch "\n\t" {
                    case "":
                        var v11 uint = 18446744073709551615
                        _ = v11
                    case "a":
                        for fi12 := 0; fi12 < 4; fi12++ {
                            v13 := GenS1{F0: 42, F1: (!((85 < 792) && false)), F2: 412}
                            switch "hello" {
                            case "a":
                                var v14 int64 = 1
                                _ = v14
                            }
                            _ = fi12
                            _ = v13
                        }
                    case "\n\t":
                        v15 := v1[(fi10 | fi10)]
                        _ = v15
                    default:
                        switch (fi10 / ((fi10) | 1)) {
                        case -9223372036854775808:
                            var v16 float64 = 0.0
                            _ = v16
                        case 129:
                            if ((false || true) && (!false)) {
                                delete(v1, 32767)
                            }
                        case 42:
                            v4 += ((0.0 + v4) * ((1e10 - v4) + 1e10))
                        default:
                            var v17 int32 = -2147483648
                            _ = v17
                        }
                    }
                }
                if len(v3) > 0 {
                    v3[0] = (fi10 >> ((2) & 63))
                }
                _ = fi10
            }
        case -997:
            var v18 int8 = 62
            _ = v18
        }
        if false {
            v4 += v4
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v19 := v1[0]
        var v20 uint32 = (^(v19))
        _ = v19
        _ = v20
    }
    v3 = append(v3, 32767)
    for fi21 := 0; fi21 < 13; fi21++ {
        fmt.Println(v9, v2)
        _ = fi21
    }
    fmt.Println(v9, v9)
    var v22 uint16 = 2
    for fi23 := 0; fi23 < 20; fi23++ {
        v9 += 1
        if (!false) {
            for fi24 := 0; fi24 < 1; fi24++ {
                v2 = (v2 / v2)
                _ = fi24
            }
            var v25 uint16 = (^(v9))
            _ = v25
        }
        if (true && ("hello" <= "foo bar")) {
            var v26 bool = true
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            var v27 int8 = -110
            _ = v26
            _ = v27
        }
        _ = fi23
    }
    var v28 uint64 = 1
    if true {
        if true {
            switch 198 {
            case -2147483648:
                var v29 uint8 = 1
                _ = v29
            default:
                var v30 float64 = 3.14159
                _ = v30
            }
            var v31 int = 1
            _ = len(v1)
            _ = v31
        }
        var v32 uint16 = (v22 + 32767)
        _ = v32
    }
    var v33 uint32 = 42
    var v34 uint8 = 1
    v33 = v33
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v9
}
