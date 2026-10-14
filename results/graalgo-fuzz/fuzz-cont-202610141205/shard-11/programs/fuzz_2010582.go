package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32) int8 {
    v4 := []int{-825}
    fmt.Println(p0)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    _ = p0
    _ = v4
    return 1
}

func main() {
    var v1 uint16 = 325
    var v2 bool = (v1 == (v1 | v1))
    v3 := []uint32{870}
    v2 = ((99 == 249) && (2147483647 != 1))
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: ("a" + "a"), F1: 0, F2: 18446744073709551615}
            _ = fi6
            _ = v7
        }
        fmt.Println(v1)
        _ = fi5
    }
    if (-2.25 <= (((3.14159 - 0.0) / 1e10) + float32(((0.1 * 3.14159) + -2.25)))) {
        if false {
            var v8 uint8 = 0
            _ = v8
        } else {
            if (42 == -88) {
                switch "\n\t" {
                case "\n\t":
                    var v9 uint64 = 42
                    _ = v9
                default:
                    for fi10 := 0; fi10 < 15; fi10++ {
                        for fi11 := 0; fi11 < 8; fi11++ {
                            v2 = (2 != 105)
                            if (!((543 != 242) && (!true))) {
                                v1 -= (v1 << ((304) & 63))
                            }
                            _ = fi11
                        }
                        var v12 uint64 = 1
                        if len(v3) > 0 {
                            v3[0] = 312
                        }
                        _ = fi10
                        _ = v12
                    }
                }
                var v13 bool = (1 <= 18446744073709551615)
                fmt.Println(v1, v13)
                _ = v13
            } else {
                for fi14 := 0; fi14 < 1; fi14++ {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v2 = v2
                        v16 := GenS2{F0: (fi14 & (fi14 ^ (fi15 >> ((464) & 63))))}
                        _ = len(v3)
                        _ = fi15
                        _ = v16
                    }
                    var v17 string = "foo bar"
                    fmt.Println(v2, v17)
                    _ = fi14
                    _ = v17
                }
                v18 := GenS2{F0: 4294967295}
                _ = v18
            }
        }
        v19 := GenS3{F0: 0}
        if len(v3) > 0 {
            v3[0] = 42
        }
        _ = v19
    } else {
        var v20 int32 = -298
        if (!v2) {
            for fi21 := 0; fi21 < 9; fi21++ {
                if len(v3) > 0 {
                    v3[0] = 2
                }
                v22 := GenS2{F0: fi21}
                var v23 uint64 = 4294967295
                _ = fi21
                _ = v22
                _ = v23
            }
            switch 1 {
            case 43:
                switch -2147483648 {
                case -852:
                    for fi24 := 0; fi24 < 13; fi24++ {
                        for fi25 := 0; fi25 < 20; fi25++ {
                            var v26 int8 = 1
                            _ = fi25
                            _ = v26
                        }
                        var v27 uint32 = 2147483648
                        fmt.Println(v20, v27, v2)
                        _ = fi24
                        _ = v27
                    }
                }
            case 9223372036854775807:
                var v28 uint16 = v1
                _ = v28
            case 0:
                v29 := GenS3{F0: (-(v20))}
                _ = v29
            }
        } else {
            fmt.Println(v1, v20, v2)
            fmt.Println(v1, v2)
            for fi30 := 0; fi30 < 8; fi30++ {
                var v31 int64 = -1
                _ = fi30
                _ = v31
            }
        }
        _ = v20
    }
    var v32 uint = uint(v1)
    fmt.Println(v2)
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 int64 = -9223372036854775808
        v2 = v2
        v35 := GenS1{F0: "", F1: 65535, F2: (v32 << ((v32) & 63))}
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println(v32, v2, v32)
    for fi36 := 0; fi36 < 14; fi36++ {
        switch (fi36 + 4294967295) {
        case 127:
            if (0 != (v32 | v32)) {
                var v37 uint64 = 18446744073709551615
                _ = v37
            } else {
                switch (fi36 / ((0) | 1)) {
                case 9223372036854775807:
                    var v38 int8 = genFn0(uint32(v1))
                    _ = v38
                case 128:
                    fmt.Println(fi36, v2)
                default:
                    var v39 int32 = 0
                    _ = v39
                }
                v2 = true
            }
        default:
            switch "\n\t" {
            case "a":
                var v40 string = "a"
                _ = v40
            }
        }
        _ = fi36
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
}
