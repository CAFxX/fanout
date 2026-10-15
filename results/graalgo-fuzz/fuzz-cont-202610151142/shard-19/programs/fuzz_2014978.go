package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 int; F1 uint8 }

func main() {
    v1 := []uint32{2, 4294967295, 2147483648}
    v2 := make(map[string]string)
    var v3 float64 = 0.0
    v4 := make(map[uint32]bool)
    v3 = ((v3 + float64(float32(v3))) + ((1.5 / 1.0) * 0.0))
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(v3, fi5)
        if len(v1) > 0 {
            v1[0] = 42
        }
        _ = fi5
    }
    fmt.Println(v3)
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = 2
        var v8 float64 = v3
        var v9 int32 = 2147483647
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v3 = (((0.1 / v3) + v3) - v3)
    switch 32768 {
    case -1:
        var v10 uint8 = 128
        _ = v10
    case 326:
        for fi11 := 0; fi11 < 7; fi11++ {
            if false {
                v12 := GenS1{F0: false, F1: false, F2: true}
                for fi13 := 0; fi13 < 14; fi13++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v14 := v4[2]
                    var v15 uint32 = 1
                    var v16 uint8 = 2
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = v12
            } else {
                if false {
                    var v17 string = "\n\t"
                    for fi18 := 0; fi18 < 12; fi18++ {
                        if true {
                            v19 := GenS1{F0: (!(!true)), F1: true, F2: false}
                            if ((1 >= 1) && true) {
                                if len(v1) > 0 {
                                    v1[0] = 149
                                }
                                if (0 >= 1) {
                                    v3 = 0.0
                                    v2[v17] = v17
                                } else {
                                    v20 := v4[2147483648]
                                    var v21 string = ("foo bar" + "foo bar")
                                    _ = v20
                                    _ = v21
                                }
                            }
                            _ = v19
                        }
                        fmt.Println(v17, fi18)
                        switch (v17 + v17) {
                        case "a":
                            v22 := GenS1{F0: true, F1: false, F2: (0.1 >= ((((1.5 * ((float32((v3 + 3.14159)) * ((1e10 + 1e10) / -2.25)) - 1.5)) - 0.0) * -2.25) + 0.1))}
                            _ = v22
                        case "hello":
                            v3 += (0.1 - ((v3 * (v3 / ((v3 * 3.14159) * v3))) * ((-2.25 * (v3 * (v3 * 3.14159))) - (v3 / float64(1e10)))))
                        }
                        _ = fi18
                    }
                    _ = v17
                } else {
                    var v23 uint32 = 65535
                    for fi24 := 0; fi24 < 2; fi24++ {
                        v25 := GenS2{F0: (-(fi11)), F1: 1}
                        v26 := GenS2{F0: fi24, F1: 42}
                        fmt.Println(v3, v3)
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    _ = v23
                }
                v3 += (v3 * v3)
                v3 = (v3 * float64(1.5))
            }
            var v27 float64 = v3
            _ = fi11
            _ = v27
        }
    }
    if len(v1) > 0 {
        v1[0] = 227
    }
    if false {
        v28 := GenS1{F0: true, F1: false, F2: false}
        _ = v28
    } else {
        switch "a" {
        case "a":
            _ = len(v2)
        }
    }
    v3 += v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3)
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
