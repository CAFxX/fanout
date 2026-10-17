package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int32 }
type GenS2 struct { F0 string; F1 int8; F2 uint64 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int16, p1 uint) uint16 {
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: "", F1: -128, F2: 0}
            v10 := GenS2{F0: "hello", F1: 42, F2: 1}
            _ = fi8
            _ = v10
            _ = v9
        }
        if (32767 != (p0 | 42)) {
            p1 += 4294967295
            var v11 uint8 = 128
            p1 += (p1 << ((p1) & 63))
            _ = v11
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS3{F0: 1}
            fmt.Println(fi12)
            _ = fi12
            _ = v13
        }
        _ = fi7
    }
    fmt.Println(p0, p0)
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 float32 = float32(3.14159)
            v17 := GenS2{F0: (("" + (("\n\t" + ((("\n\t" + "foo bar") + "foo bar") + "foo bar")) + (("" + "hello") + ""))) + "\n\t"), F1: 100, F2: 0}
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := []int{(fi14 % ((2147483647) | 1)), fi14, (fi14 & fi14)}
        _ = fi14
        _ = v18
    }
    fmt.Println(p1, p1)
    v19 := []int{-1, 2147483647, -992, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v19
    return 0
}

func main() {
    v1 := []string{}
    var v2 uint64 = 819
    var v3 uint16 = 32767
    var v4 string = "\n\t"
    var v5 uint = 294
    v6 := []string{("" + "foo bar"), "hello"}
    v3 = 742
    v6 = append(v6, (v4 + v4))
    var v20 float32 = -2.25
    _ = len(v6)
    v21 := GenS3{F0: -312}
    var v22 int32 = 42
    var v23 uint16 = 285
    v23 = (v23 | v23)
    v5 -= v5
    fmt.Println(v5, v5)
    switch "" {
    case "":
        if (42 >= -9223372036854775808) {
            for fi24 := 0; fi24 < 13; fi24++ {
                v23 -= genFn0(-112, (^(v5)))
                if len(v6) > 0 {
                    v6[0] = v4
                }
                if ((!false) || false) {
                    if false {
                        v1 = append(v1, v4)
                        v25 := GenS3{F0: -1}
                        _ = v25
                    }
                    if false {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v3 = 32767
                        v26 := GenS3{F0: 0}
                        _ = v26
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v2 -= (v2 | (v2 | (v2 % ((v2) | 1))))
                        var v27 uint8 = 74
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 4; fi28++ {
                        switch (fi24 - (^(fi24))) {
                        case 127:
                            var v29 uint64 = (^(v2))
                            _ = v29
                        default:
                            switch v4 {
                            case "a":
                                switch "hello" {
                                case "":
                                    v20 = (v20 + v20)
                                case "foo bar":
                                    v1 = append(v1, ("\n\t" + v4))
                                default:
                                    _ = len(v6)
                                }
                            case "hello":
                                if true {
                                    if false {
                                        var v30 uint32 = 1
                                        fmt.Println(v4, v22)
                                        fmt.Println(v3, fi24, fi28)
                                        _ = v30
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            case "":
                                var v31 int8 = -55
                                _ = v31
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = (v4 + "a")
                        }
                        _ = fi28
                    }
                }
                _ = fi24
            }
            v5 -= (^(v5))
            switch 42 {
            case 2147483647:
                if false {
                    v22 = v22
                    v2 -= 18446744073709551615
                    fmt.Println(v5, v23)
                } else {
                    if true {
                        var v32 bool = true
                        v4 = "foo bar"
                        if len(v6) > 0 {
                            v6[0] = (v4 + "hello")
                        }
                        _ = v32
                    }
                    v4 = v4
                }
            case 4294967295:
                var v33 uint64 = 4294967295
                _ = v33
            }
        }
    case "foo bar":
        fmt.Println(v4)
    case "hello":
        var v34 uint = (v5 << ((v5) & 63))
        _ = v34
    default:
        var v35 int = -216
        _ = v35
    }
    v1 = append(v1, "\n\t")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
