package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    v1 := []string{"\n\t"}
    var v2 bool = false
    var v3 int64 = 9223372036854775807
    var v4 int32 = -904
    var v5 int8 = 100
    v6 := make(map[uint32]bool)
    var v7 int32 = v4
    v8 := GenS1{F0: -302, F1: 255}
    v9 := GenS1{F0: 306, F1: 153}
    v5 = v5
    v10 := GenS1{F0: -529, F1: 0}
    v2 = v2
    var v11 int8 = (v5 & v5)
    fmt.Println(v4, v2, v4)
    var v12 int16 = -1
    v6[1] = v2
    switch ("foo bar" + ("a" + (("" + "\n\t") + "\n\t"))) {
    case "":
        v6[366] = (v12 > 1)
    default:
        if v2 {
            v13 := GenS1{F0: v12, F1: 36}
            switch -734 {
            case -414:
                switch 2147483647 {
                case 9223372036854775807:
                    v5 = v11
                default:
                    for fi14 := 0; fi14 < 18; fi14++ {
                        if v2 {
                            for fi15 := 0; fi15 < 6; fi15++ {
                                v2 = v2
                                v11 -= (v11 * v5)
                                v16 := GenS1{F0: (v12 << ((65535) & 63)), F1: 127}
                                _ = fi15
                                _ = v16
                            }
                        } else {
                            var v17 int64 = 9223372036854775807
                            v11 -= v5
                            switch (fi14 | (fi14 / ((fi14) | 1))) {
                            case 804:
                                fmt.Println(v2)
                            case -9223372036854775808:
                                switch "foo bar" {
                                case "":
                                    if true {
                                        delete(v6, 65535)
                                    }
                                case "hello":
                                    for fi18 := 0; fi18 < 18; fi18++ {
                                        v19 := GenS2{F0: 2147483647, F1: 294}
                                        var v20 uint64 = 42
                                        v2 = v2
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                case "a":
                                    v6[0] = false
                                default:
                                    v7 += 42
                                }
                            case 127:
                                v12 = (v12 * v12)
                            default:
                                if true {
                                    var v21 int64 = (v17 >> ((4294967295) & 63))
                                    v5 -= -1
                                    for fi22 := 0; fi22 < 2; fi22++ {
                                        var v23 int = -9223372036854775808
                                        v7 -= (v7 + v4)
                                        delete(v6, 2)
                                        v3 -= v17
                                        _ = fi22
                                        _ = v23
                                    }
                                    _ = v21
                                }
                            }
                            _ = v17
                        }
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v4 = v4
                            fmt.Println(v2)
                            var v25 int64 = 0
                            _ = fi24
                            _ = v25
                        }
                        _ = fi14
                    }
                }
            case -15:
                var v26 uint = 42
                _ = v26
            case 4294967295:
                var v27 uint8 = 127
                _ = v27
            }
            if len(v1) > 0 {
                v1[0] = ""
            }
            _ = v13
        }
    }
    v12 -= v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
