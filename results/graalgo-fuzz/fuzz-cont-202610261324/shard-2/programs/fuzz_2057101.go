package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }

func main() {
    var v1 uint64 = 634
    v2 := make(map[string]int)
    var v3 uint32 = 65535
    v4 := make(map[string]int)
    var v5 uint8 = 91
    v5 += (v5 << ((2) & 63))
    var v6 uint64 = v1
    var v7 int16 = 1
    v6 += (v1 & v6)
    v7 = (v7 << ((42) & 63))
    v1 = 758
    if (-9223372036854775808 < 127) {
        switch 21 {
        case -1:
            switch -449 {
            case 128:
                for fi8 := 0; fi8 < 10; fi8++ {
                    fmt.Println(v5, v3)
                    fmt.Println(v1, fi8)
                    _ = fi8
                }
            case 127:
                if ((!true) && (3.14159 == 1.5)) {
                    if false {
                        for fi9 := 0; fi9 < 17; fi9++ {
                            v10 := []int{1}
                            _ = fi9
                            _ = v10
                        }
                    } else {
                        for fi11 := 0; fi11 < 13; fi11++ {
                            for fi12 := 0; fi12 < 19; fi12++ {
                                v13 := GenS1{F0: "", F1: int(v3)}
                                var v14 int32 = 2147483647
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                            v15 := GenS1{F0: "a", F1: fi11}
                            _ = fi11
                            _ = v15
                        }
                        v3 = (^(v3))
                        v6 += (v6 - v6)
                    }
                }
            default:
                var v16 string = "\n\t"
                _ = v16
            }
        case 32767:
            if (127 < -510) {
                var v17 int8 = 127
                v18 := []int{-2147483648, -830, -1, -9223372036854775808}
                _ = v17
                _ = v18
            } else {
                v5 += v5
                if (32767 > 736) {
                    v6 -= (v1 & (v1 << ((18446744073709551615) & 63)))
                    for fi19 := 0; fi19 < 17; fi19++ {
                        var v20 uint16 = 65535
                        _ = fi19
                        _ = v20
                    }
                }
                if true {
                    switch 1 {
                    case 4294967295:
                        fmt.Println(v5, v1, v3)
                    }
                }
            }
        }
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 float32 = -2.25
        _ = fi21
        _ = v22
    }
    var v23 float32 = float32((float64((((-2.25 + (0.1 + 1.5)) / 1.0) * -2.25)) * 3.14159))
    var v24 uint8 = 91
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS1{F0: ("hello" + "a"), F1: fi25}
        v27 := GenS1{F0: "", F1: fi25}
        _ = fi25
        _ = v26
        _ = v27
    }
    switch "a" {
    case "a":
        var v28 uint8 = v24
        _ = v28
    case "\n\t":
        v29 := GenS1{F0: ("hello" + ""), F1: -2147483648}
        _ = v29
    default:
        v7 -= -744
    }
    var v30 int64 = -706
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
