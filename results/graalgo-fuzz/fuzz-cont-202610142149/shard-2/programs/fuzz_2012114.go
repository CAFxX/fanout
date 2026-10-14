package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }

func main() {
    var v1 bool = false
    v2 := []int{128}
    var v3 uint16 = 2
    var v4 float32 = (float32(1e10) + (3.14159 * 0.1))
    var v5 int32 = 34
    var v6 string = ((("foo bar" + "") + "\n\t") + ("a" + "a"))
    for fi7 := 0; fi7 < 6; fi7++ {
        v6 = v6
        _ = fi7
    }
    v5 = v5
    switch 2147483647 {
    case 127:
        v4 = 0.1
    case 32768:
        var v8 bool = false
        _ = v8
    default:
        v9 := GenS1{F0: 604, F1: 42, F2: 128}
        _ = v9
    }
    v10 := GenS1{F0: 2, F1: -279, F2: -1}
    v2 = append(v2, 32768)
    var v11 uint32 = 1
    for fi12 := 0; fi12 < 11; fi12++ {
        v11 -= 4294967295
        _ = fi12
    }
    switch -289 {
    case 32767:
        switch v6 {
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = -1
            }
        }
    case 127:
        v13 := GenS1{F0: 0, F1: -659, F2: 833}
        _ = v13
    case 128:
        if v1 {
            switch 2147483647 {
            case -91:
                var v14 int32 = int32(v3)
                _ = v14
            case 2147483647:
                if len(v2) > 0 {
                    v2[0] = 549
                }
            }
            if (18446744073709551615 >= 42) {
                for fi15 := 0; fi15 < 9; fi15++ {
                    v16 := GenS1{F0: 249, F1: 32767, F2: fi15}
                    var v17 int64 = -604
                    switch -880 {
                    case -9223372036854775808:
                        var v18 bool = (!((v5 | 235) > (v5 | -404)))
                        _ = v18
                    default:
                        _ = len(v2)
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 int = (fi19 * fi19)
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 int = 127
                        _ = fi21
                        _ = v22
                    }
                    if ((32767 > 32767) && (!(((("a" + (v6 + "a")) + "a") + v6) <= v6))) {
                        var v23 uint32 = v11
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v23 += v23
                            if len(v2) > 0 {
                                v2[0] = (fi19 * (fi19 - 1))
                            }
                            var v25 int32 = (v5 | v5)
                            _ = fi24
                            _ = v25
                        }
                        _ = v23
                    }
                    _ = fi19
                    _ = v20
                }
            }
            if len(v2) > 0 {
                v2[0] = 248
            }
        }
    }
    v2 = append(v2, -50)
    for fi26 := 0; fi26 < 10; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 int32 = 500
            _ = fi27
            _ = v28
        }
        switch "a" {
        case "foo bar":
            fmt.Println(v5, fi26, v5)
        }
        _ = fi26
    }
    if (42 < 196) {
        v5 = v5
    }
    v2 = append(v2, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
