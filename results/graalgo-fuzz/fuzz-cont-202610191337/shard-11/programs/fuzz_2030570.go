package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 int = 805
    var v2 uint = 4294967295
    var v3 bool = (-2147483648 == 0)
    var v4 int64 = 42
    v1 -= (v1 & v1)
    if (!(v3 || (v3 && v3))) {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v1)
        default:
            v5 := GenS1{F0: int8(v4), F1: 395}
            _ = v5
        }
        v1 = -883
    } else {
        var v6 uint16 = 912
        _ = v6
    }
    v1 += 2147483647
    v7 := []int{(v1 / ((v1) | 1))}
    v8 := GenS1{F0: -1, F1: 65535}
    var v9 int64 = (v4 % ((-1) | 1))
    switch (v1 * (-(v1))) {
    case 42:
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS2{F0: 1}
            _ = fi10
            _ = v11
        }
    case -2147483648:
        v3 = (!v3)
    case 145:
        if (!v3) {
            fmt.Println(v1, v3)
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS3{F0: -128, F1: v2, F2: 85}
                for fi14 := 0; fi14 < 6; fi14++ {
                    v2 += 895
                    _ = fi14
                }
                _ = fi12
                _ = v13
            }
            if v3 {
                if v3 {
                    fmt.Println(v3, v4)
                    if v3 {
                        switch ("foo bar" + "a") {
                        case "\n\t":
                            v2 = v2
                        case "a":
                            var v15 uint32 = 2147483647
                            _ = v15
                        }
                    }
                } else {
                    var v16 uint8 = 128
                    v17 := GenS4{F0: 2147483648}
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 4; fi18++ {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v7 = append(v7, fi18)
                        var v20 bool = v3
                        _ = fi19
                        _ = v20
                    }
                    v21 := GenS4{F0: 0}
                    _ = fi18
                    _ = v21
                }
            } else {
                var v22 float64 = (1.5 * 0.1)
                if v3 {
                    v4 -= 1
                }
                v23 := GenS5{F0: 255}
                _ = v22
                _ = v23
            }
        } else {
            v1 = v1
            v4 -= v9
            if true {
                var v24 uint64 = 42
                var v25 float64 = 1.5
                if len(v7) > 0 {
                    v7[0] = (v1 | 314)
                }
                _ = v24
                _ = v25
            }
        }
    }
    var v26 uint32 = 2147483647
    var v27 float32 = 3.14159
    if v3 {
        if true {
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 bool = v3
                _ = fi28
                _ = v29
            }
            if len(v7) > 0 {
                v7[0] = 127
            }
        } else {
            if (((107 < 1) && false) && (v3 && (-106 >= 100))) {
                v4 += v9
                fmt.Println(v1, v1, v26)
                var v30 string = (((("a" + "a") + "a") + "") + "\n\t")
                _ = v30
            } else {
                v31 := GenS5{F0: 255}
                fmt.Println(v3, v2)
                for fi32 := 0; fi32 < 14; fi32++ {
                    switch v1 {
                    case 2147483647:
                        var v33 float64 = float64(v27)
                        _ = v33
                    case 32768:
                        v34 := GenS5{F0: 1}
                        _ = v34
                    case 843:
                        v35 := GenS6{F0: 4294967295}
                        _ = v35
                    }
                    var v36 uint64 = 1
                    if v3 {
                        fmt.Println(v1, v26, v27)
                        v36 += (v36 + v36)
                        for fi37 := 0; fi37 < 10; fi37++ {
                            var v38 int16 = -172
                            _ = fi37
                            _ = v38
                        }
                    } else {
                        var v39 uint = (v2 / ((0) | 1))
                        switch (v1 + fi32) {
                        case 1:
                            for fi40 := 0; fi40 < 9; fi40++ {
                                v2 -= v2
                                var v41 uint32 = v26
                                v3 = v3
                                _ = fi40
                                _ = v41
                            }
                        default:
                            for fi42 := 0; fi42 < 4; fi42++ {
                                v26 = v26
                                _ = fi42
                            }
                        }
                        _ = v39
                    }
                    _ = fi32
                    _ = v36
                }
                _ = v31
            }
        }
    } else {
        if len(v7) > 0 {
            v7[0] = v1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
