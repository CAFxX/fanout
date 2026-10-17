package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint8 }
func genFn0() uint32 {
    var v5 uint8 = 255
    for fi6 := 0; fi6 < 14; fi6++ {
        if true {
            v7 := []int{-9223372036854775808, -2147483648}
            var v8 uint32 = 0
            switch "" {
            case "hello":
                switch (fi6 + fi6) {
                case 2147483647:
                    fmt.Println(v5, v5, v5)
                case -62:
                    switch (fi6 * (-(fi6))) {
                    case 9223372036854775807:
                        for fi9 := 0; fi9 < 7; fi9++ {
                            fmt.Println(fi9, v5)
                            v8 = (v8 >> ((18446744073709551615) & 63))
                            v5 += (v5 * v5)
                            _ = fi9
                        }
                    case 128:
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    case 42:
                        var v10 string = ""
                        _ = v10
                    default:
                        if true {
                            switch (fi6 % (((fi6 / ((32768) | 1))) | 1)) {
                            case -154:
                                if true {
                                    v5 -= 42
                                    for fi11 := 0; fi11 < 10; fi11++ {
                                        var v12 string = "a"
                                        _ = fi11
                                        _ = v12
                                    }
                                }
                            case 32767:
                                fmt.Println(fi6, v5)
                            default:
                                v5 = (v5 % (((v5 - (v5 / ((v5) | 1)))) | 1))
                            }
                        } else {
                            v7 = append(v7, fi6)
                            v13 := GenS1{F0: "a", F1: 211}
                            switch (fi6 * (fi6 - 1)) {
                            case 9223372036854775807:
                                v8 += (^(v8))
                            case -548:
                                fmt.Println(fi6, fi6, v5)
                            }
                            _ = v13
                        }
                    }
                }
            case "a":
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    _ = v5
    return 2147483648
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 1
    var v4 string = "foo bar"
    switch 1 {
    case -516:
        v14 := GenS1{F0: ("\n\t" + "a"), F1: 2}
        _ = v14
    case -1:
        v15 := GenS2{F0: v1}
        _ = v15
    case 599:
        v2[v3] = (v3 >> ((832) & 63))
    }
    v1 += (v1 + v1)
    switch v4 {
    case "hello":
        v16 := []int{255, -895}
        _ = v16
    case "foo bar":
        v17 := GenS2{F0: (v1 / ((v1) | 1))}
        _ = v17
    }
    v18 := GenS1{F0: "a", F1: 18446744073709551615}
    for fi19 := 0; fi19 < 19; fi19++ {
        delete(v2, genFn0())
        switch "hello" {
        case "\n\t":
            fmt.Println(v1, v1)
        case "foo bar":
            var v20 uint = 42
            _ = v20
        }
        var v21 bool = (18446744073709551615 != 18446744073709551615)
        _ = fi19
        _ = v21
    }
    v3 += (v3 - (v3 * (v3 / (((v3 ^ (v3 << ((846) & 63)))) | 1))))
    v22 := []int{-2147483648}
    switch -9223372036854775808 {
    case 0:
        var v23 uint32 = (v3 * v3)
        _ = v23
    case 847:
        v3 += 433
    }
    fmt.Println(v3)
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 float32 = (1.5 * ((0.1 + (1.5 / 1.0)) - (1.5 + 1.5)))
        if true {
            var v26 int = fi24
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
