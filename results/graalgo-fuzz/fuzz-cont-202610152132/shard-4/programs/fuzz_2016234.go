package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 int }
func genFn0(p0 int32) uint64 {
    p0 -= p0
    if false {
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: ((-1 >= 127) && (0 < 307)), F1: fi6, F2: (("hello" + "hello") + ((("" + "foo bar") + "hello") + "\n\t"))}
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 int64 = -757
                switch ("foo bar" + "foo bar") {
                case "a":
                    v9 += 140
                }
                _ = fi8
                _ = v9
            }
            switch fi6 {
            case -2147483648:
                var v10 uint32 = 616
                _ = v10
            default:
                for fi11 := 0; fi11 < 6; fi11++ {
                    v12 := GenS1{F0: true, F1: -2147483648, F2: "hello"}
                    var v13 uint16 = 195
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            _ = fi6
            _ = v7
        }
    } else {
        var v14 uint16 = 255
        _ = v14
    }
    v15 := GenS2{F0: -1}
    _ = p0
    _ = v15
    return 42
}
func genFn1(p0 int8, p1 bool) uint16 {
    v16 := GenS1{F0: (p1 && ((4294967295 <= 2147483647) && (!false))), F1: 0, F2: "a"}
    switch 32768 {
    case 2147483647:
        p0 += (p0 << ((4294967295) & 63))
    case -620:
        p1 = true
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v16
    return 793
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    var v3 uint = 1
    var v4 bool = (66 <= 255)
    var v5 bool = ((false && true) || (!false))
    fmt.Println(v3)
    if false {
        v1 = "\n\t"
    } else {
        v17 := GenS2{F0: 230}
        fmt.Println(v5)
        v18 := GenS2{F0: 1}
        _ = v17
        _ = v18
    }
    var v19 uint8 = 242
    v5 = (true && false)
    v20 := GenS2{F0: -9223372036854775808}
    switch -321 {
    case 4294967295:
        if false {
            var v21 uint32 = 4294967295
            _ = v21
        }
    case 42:
        if (42 <= 1000) {
            if v4 {
                fmt.Println(v5, v19, v4)
                for fi22 := 0; fi22 < 14; fi22++ {
                    v23 := GenS2{F0: (fi22 & 42)}
                    v1 = "foo bar"
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v25 := GenS2{F0: (^(fi24))}
                        var v26 int64 = -9223372036854775808
                        var v27 uint16 = genFn1(96, v4)
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = fi22
                    _ = v23
                }
            } else {
                var v28 bool = v4
                v19 = v19
                _ = v28
            }
        } else {
            var v29 string = (v1 + v1)
            _ = v29
        }
    default:
        if v4 {
            var v30 float64 = (float64((1.5 - 3.14159)) - 1e10)
            _ = v30
        }
    }
    v5 = (genFn1(-19, ((v4 || v5) || (42 >= 127))) != uint16(v19))
    var v31 string = v1
    v32 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
