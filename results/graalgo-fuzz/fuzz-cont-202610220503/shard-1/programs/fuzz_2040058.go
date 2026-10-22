package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int32 = 1
    var v2 int16 = 1
    var v3 int16 = (v2 | 1000)
    v4 := make(map[uint32]bool)
    var v5 float64 = (-2.25 * 1.5)
    v6 := GenS1{F0: 671, F1: "\n\t", F2: 1}
    var v7 int16 = v2
    var v8 uint16 = 823
    var v9 uint = 1
    var v10 uint = 2
    v11 := GenS1{F0: 635, F1: ("hello" + "\n\t"), F2: 109}
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint = v10
        var v14 uint = (^(v13))
        _ = fi12
        _ = v13
        _ = v14
    }
    if false {
        if true {
            if (false || true) {
                switch 32767 {
                case 9223372036854775807:
                    v10 = v9
                case 69:
                    var v15 uint16 = 705
                    _ = v15
                case -583:
                    var v16 string = ("foo bar" + ((("foo bar" + "hello") + "foo bar") + "a"))
                    _ = v16
                }
                if (812 > 42) {
                    v9 += 120
                } else {
                    if false {
                        for fi17 := 0; fi17 < 1; fi17++ {
                            if true {
                                if ((((true && (!(4294967295 <= 747))) || (true && true)) || (!((v9 == v9) || (!false)))) || ((v5 != float64(1e10)) || ((v5 / float64((-2.25 + 0.0))) > v5))) {
                                    delete(v4, 4294967295)
                                    var v18 float64 = (1.5 * v5)
                                    var v19 int8 = -1
                                    _ = v18
                                    _ = v19
                                } else {
                                    if (-801 != 818) {
                                        for fi20 := 0; fi20 < 12; fi20++ {
                                            v21 := GenS1{F0: 1, F1: "\n\t", F2: -66}
                                            _ = fi20
                                            _ = v21
                                        }
                                        var v22 uint64 = 42
                                        _ = v22
                                    } else {
                                        v23 := []int{fi17, (fi17 & (fi17 * fi17)), int(v8)}
                                        v24 := GenS1{F0: 1, F1: (("" + "\n\t") + ((("" + "foo bar") + "") + ("hello" + "foo bar"))), F2: 127}
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v25 uint8 = 255
                                            _ = v25
                                        }
                                        _ = v23
                                        _ = v24
                                    }
                                    var v26 bool = (1 < 1)
                                    _ = v26
                                }
                                var v27 int = (fi17 & 32767)
                                var v28 int16 = (v7 | (v7 + (v7 << ((v9) & 63))))
                                _ = v27
                                _ = v28
                            }
                            v10 += (^(v9))
                            _ = fi17
                        }
                        v29 := GenS1{F0: 42, F1: (("foo bar" + (((("hello" + "") + "foo bar") + "") + "a")) + "hello"), F2: -114}
                        _ = v29
                    } else {
                        v30 := GenS1{F0: 18446744073709551615, F1: (("a" + "") + ""), F2: 1}
                        _ = v30
                    }
                    v31 := GenS2{F0: (v8 - (^(v8)))}
                    _ = v31
                }
            } else {
                v32 := []int{630, 2147483647}
                _ = v32
            }
            v9 = 42
            v1 = v1
        }
        var v33 int8 = -1
        fmt.Println(v5, v3)
        _ = v33
    }
    delete(v4, 0)
    v8 -= v8
    v34 := GenS2{F0: (v8 & v8)}
    var v35 uint8 = 127
    for fi36 := 0; fi36 < 15; fi36++ {
        fmt.Println(v35, v9, v2)
        _ = fi36
    }
    v37 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
