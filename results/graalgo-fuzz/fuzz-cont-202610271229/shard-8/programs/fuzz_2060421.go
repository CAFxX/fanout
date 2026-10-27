package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]int)
    v3 := []int{-47, 2147483647}
    v4 := []int{933, 2147483647}
    fmt.Println(v1)
    for fi5 := 0; fi5 < 6; fi5++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v6 := GenS1{F0: 42}
        v1 -= v1
        _ = fi5
        _ = v6
    }
    if true {
        if ("foo bar" < (("" + ("foo bar" + "foo bar")) + ((((("\n\t" + "") + ("a" + "foo bar")) + "\n\t") + "foo bar") + "foo bar"))) {
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 int32 = -713
                v8 = v8
                _ = fi7
                _ = v8
            }
        } else {
            _ = len(v3)
        }
        fmt.Println(v1)
    } else {
        v1 = v1
    }
    if ((true && (32767 == 1000)) || (!(18446744073709551615 < 0))) {
        v1 = (v1 * (v1 & v1))
        fmt.Println(v1)
        if false {
            if false {
                if (false || (true && true)) {
                    v1 = (v1 * (^(v1)))
                    var v9 int32 = 0
                    _ = v9
                }
                for fi10 := 0; fi10 < 15; fi10++ {
                    v1 = v1
                    _ = fi10
                }
                for fi11 := 0; fi11 < 2; fi11++ {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v1 = (v1 | v1)
                        if (607 >= 42) {
                            v13 := GenS2{F0: ("a" > ""), F1: 1}
                            v1 = v1
                            fmt.Println(v1, v1, v1)
                            _ = v13
                        } else {
                            var v14 int8 = 0
                            if true {
                                v1 = 1
                                fmt.Println(fi12, fi11, v1)
                                fmt.Println(v1)
                            } else {
                                var v15 int32 = 0
                                _ = v15
                            }
                            _ = v14
                        }
                        var v16 int16 = 721
                        _ = fi12
                        _ = v16
                    }
                    _ = fi11
                }
            } else {
                fmt.Println(v1)
                v1 += v1
                v17 := GenS3{F0: 65535, F1: 48}
                _ = v17
            }
            var v18 uint8 = 127
            var v19 uint = 1
            _ = v18
            _ = v19
        } else {
            var v20 float64 = 3.14159
            if true {
                fmt.Println(v20, v1)
                var v21 uint = 347
                switch -2147483648 {
                case 4294967295:
                    var v22 int32 = 918
                    _ = v22
                default:
                    var v23 int32 = -2147483648
                    _ = v23
                }
                _ = v21
            } else {
                v1 += v1
                v24 := GenS3{F0: 2147483648, F1: 127}
                _ = v24
            }
            _ = v20
        }
    }
    fmt.Println(v1)
    var v25 int64 = -909
    var v26 uint16 = 255
    v1 -= 255
    var v27 bool = (!true)
    if v27 {
        v25 = 1
    }
    var v28 int32 = -1
    var v29 float64 = (1e10 * (1e10 / (1.5 / 0.1)))
    var v30 float64 = 0.0
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 uint8 = 255
        _ = fi31
        _ = v32
    }
    v33 := GenS4{F0: (v26 + (v26 | (v26 << ((906) & 63)))), F1: -81}
    if v27 {
        var v34 bool = v27
        _ = v34
    } else {
        if v27 {
            v35 := GenS5{F0: v26}
            var v36 int64 = 0
            v29 = float64((1e10 - 0.0))
            _ = v35
            _ = v36
        }
        delete(v2, 868)
        var v37 string = "hello"
        _ = v37
    }
    v38 := GenS1{F0: 103}
    v39 := v2[1]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v38
    _ = v39
    _ = v4
}
