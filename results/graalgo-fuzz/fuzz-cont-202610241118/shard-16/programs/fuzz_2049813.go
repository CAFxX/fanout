package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int16 }

func main() {
    var v1 uint32 = 838
    var v2 int8 = -40
    var v3 int16 = -1
    v4 := make(map[string]string)
    v5 := []string{"hello", "a"}
    if (v3 < v3) {
        var v6 bool = false
        _ = v6
    }
    if (!((0.1 * 1e10) < float32(1e10))) {
        switch "a" {
        case "a":
            for fi7 := 0; fi7 < 15; fi7++ {
                switch fi7 {
                case -1:
                    fmt.Println(v1, v3)
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi7
            }
        case "":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "hello":
            _ = len(v5)
        default:
            switch "foo bar" {
            case "":
                var v8 uint64 = 0
                _ = v8
            }
        }
    }
    fmt.Println(v2, v1)
    var v9 uint64 = 42
    delete(v4, "hello")
    v10 := v4[(("a" + "\n\t") + "")]
    switch 715 {
    case 1:
        v11 := GenS1{F0: 18446744073709551615, F1: (v9 ^ (^(v9))), F2: v3}
        _ = v11
    default:
        v2 += v2
    }
    fmt.Println(v10, v1, v10)
    v10 = (v10 + (v10 + v10))
    for fi12 := 0; fi12 < 19; fi12++ {
        if false {
            if false {
                if (true && false) {
                    var v13 int64 = 1
                    var v14 int32 = -2147483648
                    _ = v13
                    _ = v14
                } else {
                    if (false || false) {
                        if (83 != (v2 << ((1) & 63))) {
                            switch "a" {
                            case "foo bar":
                                var v15 float32 = -2.25
                                _ = v15
                            case "a":
                                fmt.Println(v1, v3, v2)
                            case "":
                                fmt.Println(v10, v9)
                            default:
                                var v16 uint = 4294967295
                                _ = v16
                            }
                            var v17 uint8 = 2
                            var v18 int64 = 42
                            _ = v17
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 17; fi19++ {
                            var v20 int32 = 793
                            v21 := GenS1{F0: 0, F1: 0, F2: (v3 & (v3 >> ((uint(v1)) & 63)))}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    for fi22 := 0; fi22 < 12; fi22++ {
                        switch ("hello" + v10) {
                        case "\n\t":
                            if true {
                                v10 = "a"
                            }
                        case "foo bar":
                            v23 := GenS1{F0: 220, F1: 575, F2: (-(v3))}
                            _ = v23
                        }
                        v1 = 4294967295
                        _ = fi22
                    }
                }
            }
        }
        var v24 int8 = (^(v2))
        _ = fi12
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
