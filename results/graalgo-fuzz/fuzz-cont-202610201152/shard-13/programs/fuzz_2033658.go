package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    var v4 uint8 = 127
    var v5 float32 = 3.14159
    var v6 int = -163
    v1 += v1
    fmt.Println(v1)
    v7 := GenS1{F0: 352, F1: 47}
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 bool = false
            var v11 string = ""
            var v12 float64 = -2.25
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    var v13 int64 = -9223372036854775808
    v1 = v1
    if false {
        var v14 int32 = 1
        v6 = (v6 - v6)
        switch ("\n\t" + "hello") {
        case "\n\t":
            _ = len(v3)
        default:
            v15 := []int{32768, v6, v6}
            _ = v15
        }
        _ = v14
    } else {
        delete(v3, (^(v6)))
        for fi16 := 0; fi16 < 3; fi16++ {
            v13 += v13
            var v17 int16 = -389
            _ = fi16
            _ = v17
        }
    }
    switch v6 {
    case -764:
        for fi18 := 0; fi18 < 16; fi18++ {
            v1 += v1
            _ = fi18
        }
    default:
        var v19 float64 = (1.5 + -2.25)
        _ = v19
    }
    v6 += v6
    if false {
        var v20 uint = 42
        fmt.Println(v5, v4)
        _ = v20
    }
    v3[v6] = 2147483647
    v13 = (v13 << ((4294967295) & 63))
    for fi21 := 0; fi21 < 20; fi21++ {
        if false {
            v22 := GenS1{F0: 519, F1: 2}
            _ = v22
        } else {
            var v23 uint16 = 65535
            if (v13 <= v13) {
                v24 := GenS1{F0: 555, F1: 989}
                _ = v24
            } else {
                switch "\n\t" {
                case "":
                    if false {
                        if (false && (v6 <= (v6 ^ -9223372036854775808))) {
                            for fi25 := 0; fi25 < 15; fi25++ {
                                var v26 uint16 = v23
                                _ = fi25
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 13; fi27++ {
                                _ = len(v3)
                                var v28 uint8 = (v4 + (v4 - (v4 - 0)))
                                fmt.Println(v1)
                                _ = fi27
                                _ = v28
                            }
                        }
                    } else {
                        v29 := GenS2{F0: -802, F1: v1, F2: (v23 * (v23 - (v23 ^ (v23 * 223))))}
                        v30 := []int{32767, (fi21 - (v6 << ((65535) & 63))), (v6 ^ 0), fi21}
                        _ = v29
                        _ = v30
                    }
                default:
                    switch 128 {
                    case 324:
                        v5 -= (v5 * ((v5 / 0.1) - 1e10))
                    }
                }
                for fi31 := 0; fi31 < 19; fi31++ {
                    v6 -= (v6 ^ (v6 ^ fi31))
                    var v32 int32 = 1
                    var v33 int8 = 100
                    v6 += fi31
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = len(v2)
            }
            var v34 uint = 2
            _ = v23
            _ = v34
        }
        _ = fi21
    }
    var v35 int32 = -2147483648
    fmt.Println(v5, v4, v1)
    var v36 bool = (!(true || (324 == 255)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
