package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int8 = 100
    v2 := []int{}
    var v3 int16 = -673
    var v4 uint64 = 4294967295
    var v5 uint8 = 241
    v6 := GenS1{F0: v3}
    fmt.Println(v1, v5)
    var v7 uint = 487
    for fi8 := 0; fi8 < 17; fi8++ {
        if false {
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := GenS1{F0: v3}
                _ = fi9
                _ = v10
            }
            var v11 int16 = v3
            _ = v11
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v5 -= (^(v5))
        _ = fi8
    }
    v12 := GenS1{F0: 1}
    if len(v2) > 0 {
        v2[0] = 568
    }
    var v13 uint64 = v4
    var v14 int8 = v1
    var v15 uint = (v7 % ((553) | 1))
    if ((v13 >> ((v15) & 63)) > uint64(v5)) {
        if (v7 > v15) {
            if true {
                switch ((("hello" + "") + "foo bar") + "hello") {
                case "a":
                    switch 4294967295 {
                    case -375:
                        if (1 >= uint32(v7)) {
                            var v16 string = "a"
                            var v17 uint32 = 610
                            switch ((v16 + ("\n\t" + (v16 + "foo bar"))) + v16) {
                            case "hello":
                                if len(v2) > 0 {
                                    v2[0] = -808
                                }
                            case "a":
                                v5 = (v5 ^ (^(v5)))
                            }
                            _ = v16
                            _ = v17
                        } else {
                            if false {
                                v18 := GenS1{F0: (v3 << (((v7 << ((v15) & 63))) & 63))}
                                v3 = v3
                                v4 = 42
                                _ = v18
                            } else {
                                v13 = 332
                                for fi19 := 0; fi19 < 5; fi19++ {
                                    for fi20 := 0; fi20 < 8; fi20++ {
                                        switch (("a" + ("" + "foo bar")) + "") {
                                        case "\n\t":
                                            var v21 float32 = (((3.14159 - ((-2.25 * 0.0) / -2.25)) * 0.0) / 0.1)
                                            _ = v21
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v22 := GenS1{F0: v3}
                                        _ = fi20
                                        _ = v22
                                    }
                                    v3 = v3
                                    if true {
                                        fmt.Println(v13)
                                    } else {
                                        if ((v5 | v5) < (v5 * (v5 >> (((v15 >> ((4294967295) & 63))) & 63)))) {
                                            var v23 uint32 = 2147483647
                                            _ = v23
                                        }
                                        fmt.Println(fi19, fi19)
                                        for fi24 := 0; fi24 < 9; fi24++ {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = fi24
                                        }
                                    }
                                    _ = fi19
                                }
                                var v25 uint = 65535
                                _ = v25
                            }
                            var v26 int64 = 42
                            v2 = append(v2, 4294967295)
                            _ = v26
                        }
                    case 127:
                        _ = len(v2)
                    }
                case "hello":
                    for fi27 := 0; fi27 < 4; fi27++ {
                        if false {
                            v28 := GenS1{F0: (v3 - v3)}
                            _ = v28
                        }
                        _ = fi27
                    }
                }
            } else {
                for fi29 := 0; fi29 < 15; fi29++ {
                    if false {
                        var v30 float64 = (-2.25 / -2.25)
                        var v31 int32 = 2147483647
                        _ = v30
                        _ = v31
                    }
                    var v32 float32 = 0.1
                    _ = fi29
                    _ = v32
                }
                fmt.Println(v5, v4, v13)
            }
            v33 := GenS2{F0: (v4 * 18446744073709551615)}
            v34 := GenS3{F0: (v15 * v15)}
            _ = v33
            _ = v34
        }
        fmt.Println(v13, v13, v13)
        var v35 float64 = 1e10
        _ = v35
    } else {
        v7 += 889
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
