package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"", (("hello" + "hello") + (((((("\n\t" + "foo bar") + "\n\t") + ((((("\n\t" + "a") + "foo bar") + "") + "foo bar") + "\n\t")) + ("a" + "hello")) + "hello") + ("" + ("foo bar" + ""))))}
    v3 := []int{-912, 0, -9223372036854775808, 42}
    var v4 uint64 = 2
    v4 = v4
    v5 := GenS1{F0: -9223372036854775808}
    v4 = (^(v4))
    var v6 int = 690
    var v7 uint = 337
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            switch 4294967295 {
            case 9223372036854775807:
                var v10 int = 32768
                var v11 string = ("a" + ("a" + ("foo bar" + "hello")))
                _ = v10
                _ = v11
            case 724:
                if false {
                    var v12 float64 = (1e10 - -2.25)
                    var v13 float32 = 1.5
                    _ = v12
                    _ = v13
                } else {
                    var v14 uint32 = 42
                    if true {
                        if (-1 > 42) {
                            v15 := GenS1{F0: v6}
                            v4 -= v4
                            if (true && ("\n\t" > "")) {
                                if true {
                                    v7 = v7
                                    var v16 int = v6
                                    _ = v16
                                }
                                v14 -= (v14 & 42)
                                var v17 int = fi9
                                var v18 uint32 = (v14 + v14)
                                var v19 int = (v17 >> ((v7) & 63))
                                _ = v17
                                _ = v18
                                _ = v19
                            } else {
                                v7 = (v7 | (v7 + (v7 << ((v7) & 63))))
                            }
                            fmt.Println(v4)
                            v6 = fi8
                            _ = v15
                        }
                        v20 := v1[(^(v14))]
                        v21 := GenS1{F0: (fi9 - v20)}
                        v7 = 97
                        if (920 > 42) {
                            v20 = v6
                            var v22 uint64 = (v4 * 230)
                            _ = v22
                        } else {
                            v6 += fi8
                            fmt.Println(v6, v4)
                            var v23 string = "hello"
                            _ = v23
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v20
                        _ = v21
                    }
                    _ = v14
                }
            default:
                var v24 uint64 = v4
                _ = v24
            }
            v1[2147483648] = (v6 << ((v7) & 63))
            v7 += (v7 << (((v7 - (v7 * (v7 & 800)))) & 63))
            _ = fi9
        }
        _ = fi8
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 bool = false
        if true {
            if (1 >= 127) {
                var v27 int = (v6 / ((-177) | 1))
                var v28 int8 = 127
                _ = v27
                _ = v28
            }
        }
        _ = fi25
        _ = v26
    }
    v7 -= v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
