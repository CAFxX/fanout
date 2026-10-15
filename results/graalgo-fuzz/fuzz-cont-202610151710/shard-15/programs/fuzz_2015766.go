package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint64, p1 bool, p2 uint32) int64 {
    var v4 int8 = 42
    v5 := GenS1{F0: v4, F1: 1}
    for fi6 := 0; fi6 < 19; fi6++ {
        p2 -= p2
        switch "\n\t" {
        case "foo bar":
            for fi7 := 0; fi7 < 6; fi7++ {
                var v8 float64 = (1.5 + 0.0)
                _ = fi7
                _ = v8
            }
        case "a":
            v9 := GenS1{F0: (v4 + (-(v4))), F1: 255}
            _ = v9
        }
        _ = fi6
    }
    var v10 int8 = 42
    var v11 bool = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v4
    _ = v5
    return -325
}

func main() {
    var v1 int32 = 1
    v2 := []int{-1}
    var v3 int16 = 0
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS2{F0: v1, F1: (v1 + 42), F2: v3}
        _ = fi12
        _ = v13
    }
    v1 = (v1 << ((42) & 63))
    if ((!false) && (!true)) {
        if (433 > 1) {
            if (0 >= 0) {
                for fi14 := 0; fi14 < 8; fi14++ {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v16 := GenS3{F0: 1}
                        var v17 float32 = 3.14159
                        fmt.Println(v1)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS3{F0: 18446744073709551615}
                    _ = len(v2)
                    _ = fi14
                    _ = v18
                }
                v1 = v1
            }
            _ = len(v2)
        }
        if false {
            var v19 uint32 = 341
            if true {
                if true {
                    var v20 float32 = float32((((0.1 + 1.5) / (((0.0 * 0.0) - (1e10 * (1.5 + 0.0))) * 3.14159)) - (-2.25 + (1e10 * -2.25))))
                    _ = v20
                } else {
                    v2 = append(v2, int(v1))
                    if true {
                        v3 += v3
                    } else {
                        v19 = (v19 * v19)
                        v3 += (v3 - v3)
                        var v21 int16 = (v3 >> ((42) & 63))
                        _ = v21
                    }
                }
                if false {
                    for fi22 := 0; fi22 < 5; fi22++ {
                        fmt.Println(v1, v19, v3)
                        var v23 uint = 879
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(v1, v3)
                } else {
                    var v24 bool = false
                    _ = v24
                }
            } else {
                switch "a" {
                case "\n\t":
                    v19 = (^(v19))
                case "foo bar":
                    for fi25 := 0; fi25 < 10; fi25++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi25
                    }
                case "a":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v26 uint64 = 618
                    _ = v26
                }
            }
            _ = v19
        } else {
            v1 += v1
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v27 bool = true
    v1 -= -1
    v27 = v27
    v2 = append(v2, -519)
    switch "hello" {
    case "\n\t":
        switch (("" + "") + ("" + "hello")) {
        case "":
            var v28 uint32 = 2147483647
            _ = v28
        }
    case "foo bar":
        if v27 {
            v27 = ((0.1 + 0.1) < -2.25)
            if v27 {
                var v29 int8 = 100
                _ = v29
            }
        }
    default:
        _ = len(v2)
    }
    switch 2147483647 {
    case -9223372036854775808:
        fmt.Println(v27, v3, v27)
    case 32768:
        v3 -= 42
    case -119:
        if v27 {
            var v30 string = "\n\t"
            if v27 {
                var v31 string = (v30 + "a")
                _ = v31
            }
            if v27 {
                v1 -= (v1 << ((423) & 63))
                var v32 int32 = v1
                _ = v32
            }
            _ = v30
        } else {
            v27 = v27
        }
    }
    v33 := GenS1{F0: 1, F1: 42}
    var v34 string = ("a" + "")
    var v35 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
