package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    v1 = -743
    fmt.Println(v1)
    for fi5 := 0; fi5 < 5; fi5++ {
        v1 -= v1
        _ = fi5
    }
    v6 := GenS1{F0: 2147483647}
    v1 += v1
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: fi7}
        switch (fi7 | fi7) {
        case 9223372036854775807:
            switch (fi7 ^ (fi7 & (fi7 - fi7))) {
            case 997:
                if (1 == 0) {
                    v9 := v3[42]
                    if ((1 < 1) && false) {
                        v10 := []int{-722, fi7, (fi7 << ((255) & 63)), -1}
                        for fi11 := 0; fi11 < 4; fi11++ {
                            v9 = v9
                            _ = fi11
                        }
                        switch "" {
                        case "":
                            for fi12 := 0; fi12 < 17; fi12++ {
                                switch 2147483647 {
                                case -1:
                                    v9 = ""
                                }
                                v9 = "\n\t"
                                v1 = 786
                                _ = fi12
                            }
                        case "a":
                            for fi13 := 0; fi13 < 12; fi13++ {
                                if true {
                                    fmt.Println(v1)
                                } else {
                                    if (731 < -32768) {
                                        v9 = (v9 + "foo bar")
                                    } else {
                                        fmt.Println(fi13, v1, v1)
                                        v14 := GenS2{F0: 128}
                                        v9 = v9
                                        _ = v14
                                    }
                                    _ = len(v10)
                                    if false {
                                        if false {
                                            var v15 int32 = (v1 >> ((18446744073709551615) & 63))
                                            _ = v15
                                        } else {
                                            fmt.Println(v1, v1)
                                            fmt.Println(fi13, v1, v1)
                                        }
                                        v1 = (-(v1))
                                    } else {
                                        v9 = ("" + (v9 + v9))
                                        fmt.Println(v9, v9, v1)
                                        v1 += 805
                                    }
                                }
                                _ = fi13
                            }
                        default:
                            v1 = v1
                        }
                        _ = v10
                    } else {
                        var v16 int8 = -47
                        v1 = (v1 >> ((255) & 63))
                        v17 := []int{}
                        _ = v16
                        _ = v17
                    }
                    _ = v9
                }
            case 874:
                if true {
                    v1 = (-(v1))
                    fmt.Println(fi7, fi7)
                    v18 := GenS1{F0: fi7}
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        fmt.Println(fi7)
                        _ = fi19
                    }
                }
            case 912:
                if (1e10 <= 1e10) {
                    var v20 int8 = -128
                    _ = v20
                }
            }
        }
        _ = fi7
        _ = v8
    }
    v21 := GenS1{F0: 0}
    if false {
        var v22 int8 = -128
        _ = v22
    } else {
        v23 := GenS2{F0: 128}
        var v24 int32 = v1
        _ = v23
        _ = v24
    }
    switch 1 {
    case 9223372036854775807:
        switch 32767 {
        case -66:
            for fi25 := 0; fi25 < 17; fi25++ {
                v1 -= (v1 ^ v1)
                var v26 int = fi25
                _ = fi25
                _ = v26
            }
        case 32768:
            var v27 bool = false
            _ = v27
        }
    default:
        v28 := GenS2{F0: 42}
        _ = v28
    }
    var v29 int16 = -32768
    v30 := GenS1{F0: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v6
}
