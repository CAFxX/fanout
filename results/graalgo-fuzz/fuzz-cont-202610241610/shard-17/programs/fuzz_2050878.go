package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 int64 }

func main() {
    var v1 string = "\n\t"
    var v2 int64 = -426
    var v3 uint8 = 127
    var v4 int8 = -24
    v5 := GenS1{F0: 32, F1: false}
    v1 = v1
    v6 := []int{128}
    if true {
        v7 := GenS1{F0: 42, F1: false}
        v3 = 1
        _ = v7
    } else {
        if (!((true || (18446744073709551615 > 0)) || ((569 == 2) || true))) {
            var v8 string = v1
            _ = v8
        } else {
            if len(v6) > 0 {
                v6[0] = 421
            }
            v9 := GenS1{F0: 28, F1: false}
            _ = v9
        }
    }
    if true {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint32 = 2147483648
            _ = fi10
            _ = v11
        }
        if len(v6) > 0 {
            v6[0] = -798
        }
        if false {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 uint8 = v3
                v14 := GenS2{F0: v4, F1: 0}
                for fi15 := 0; fi15 < 7; fi15++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v16 uint32 = 1
                    _ = fi15
                    _ = v16
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            v17 := GenS1{F0: 1, F1: ((0.1 * (1.5 + (0.0 - (1.5 * ((((1e10 / 1e10) - 3.14159) * 1.5) * (-2.25 * -2.25)))))) <= 1.5)}
            _ = v17
        }
    } else {
        if (2147483647 != 1) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v4 = -128
        } else {
            for fi18 := 0; fi18 < 2; fi18++ {
                v4 -= 33
                for fi19 := 0; fi19 < 9; fi19++ {
                    switch v1 {
                    case "a":
                        if false {
                            var v20 uint16 = 255
                            switch "a" {
                            case "a":
                                fmt.Println(fi18)
                            case "\n\t":
                                if true {
                                    if len(v6) > 0 {
                                        v6[0] = 32768
                                    }
                                    fmt.Println(v3)
                                }
                            case "":
                                v6 = append(v6, fi18)
                                var v21 int64 = (v2 % ((v2) | 1))
                                var v22 int64 = -771
                                _ = v21
                                _ = v22
                            default:
                                v2 -= v2
                            }
                            var v23 bool = true
                            _ = v20
                            _ = v23
                        }
                    }
                    if (!false) {
                        v24 := GenS3{F0: 969}
                        _ = v24
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    _ = fi19
                }
                _ = fi18
            }
            v25 := GenS4{F0: true}
            for fi26 := 0; fi26 < 16; fi26++ {
                v6 = append(v6, int(v4))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi26
            }
            _ = v25
        }
        var v27 uint32 = 4294967295
        _ = v27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        for fi29 := 0; fi29 < 10; fi29++ {
            if len(v6) > 0 {
                v6[0] = 298
            }
            var v30 float64 = -2.25
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v31 := GenS5{F0: 581, F1: v2}
    for fi32 := 0; fi32 < 17; fi32++ {
        v33 := GenS3{F0: 906}
        var v34 uint16 = 872
        var v35 float64 = ((3.14159 + -2.25) / (3.14159 + -2.25))
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    v4 += 0
    fmt.Println(v4, v2)
    fmt.Println(v4)
    var v36 uint8 = v3
    var v37 uint8 = v3
    v4 = 37
    if len(v6) > 0 {
        v6[0] = int(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
