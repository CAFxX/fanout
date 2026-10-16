package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int64 }
func genFn0(p0 uint16) uint8 {
    v5 := []int{-9223372036854775808}
    switch "hello" {
    case "":
        switch "foo bar" {
        case "a":
            var v6 uint16 = 0
            _ = v6
        case "\n\t":
            fmt.Println(p0)
        }
    case "hello":
        for fi7 := 0; fi7 < 6; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 uint32 = 2
                switch (fi8 - fi8) {
                case 128:
                    p0 += p0
                default:
                    if (1 == 0) {
                        fmt.Println(fi8, v9, fi7)
                        switch "hello" {
                        case "foo bar":
                            if false {
                                v9 -= uint32(p0)
                                if len(v5) > 0 {
                                    v5[0] = fi7
                                }
                                p0 -= p0
                            }
                        case "hello":
                            v10 := GenS1{F0: 42, F1: -850, F2: -9223372036854775808}
                            _ = v10
                        }
                        var v11 uint32 = (v9 | (^(v9)))
                        _ = v11
                    }
                }
                v5 = append(v5, (^(fi8)))
                _ = fi8
                _ = v9
            }
            fmt.Println(p0, fi7)
            _ = fi7
        }
    default:
        v12 := GenS1{F0: 679, F1: 496, F2: -1}
        _ = v12
    }
    var v13 float64 = 0.1
    _ = p0
    _ = v13
    _ = v5
    return 1
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint8 = 2
    v3 := make(map[string]string)
    var v4 int16 = 1
    var v14 int32 = (v1 + -470)
    if true {
        var v15 uint = 89
        _ = v15
    }
    if false {
        v16 := []int{}
        v14 += int32(v4)
        v14 -= int32(v4)
        _ = v16
    } else {
        v17 := v3["\n\t"]
        v2 -= v2
        _ = v17
    }
    switch "a" {
    case "hello":
        for fi18 := 0; fi18 < 10; fi18++ {
            v4 = v4
            _ = fi18
        }
    case "a":
        _ = len(v3)
    default:
        switch 0 {
        case 32767:
            switch 32767 {
            case -9223372036854775808:
                var v19 int32 = (v1 >> ((65535) & 63))
                _ = v19
            case -2147483648:
                var v20 int16 = (v4 % ((-366) | 1))
                _ = v20
            }
        case 654:
            if true {
                v14 += v14
                v21 := []int{127}
                _ = v21
            }
        case 9223372036854775807:
            fmt.Println(v2, v1, v2)
        default:
            if (true || ((!false) || ((!false) || (false && (174 >= 42))))) {
                for fi22 := 0; fi22 < 7; fi22++ {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        var v24 uint = 65535
                        var v25 string = ((("a" + "a") + "\n\t") + (((("" + "") + ("foo bar" + "hello")) + "a") + "hello"))
                        v26 := GenS1{F0: v14, F1: 131, F2: 1}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    var v27 string = ""
                    v14 -= 2147483647
                    _ = fi22
                    _ = v27
                }
                for fi28 := 0; fi28 < 7; fi28++ {
                    v29 := []int{}
                    _ = fi28
                    _ = v29
                }
                v14 = v14
            }
        }
    }
    v4 -= v4
    var v30 int64 = 1
    v31 := []int{913, 9223372036854775807}
    for fi32 := 0; fi32 < 15; fi32++ {
        switch ("foo bar" + ("foo bar" + (("\n\t" + "a") + "hello"))) {
        case "foo bar":
            var v33 float64 = (1.5 * -2.25)
            _ = v33
        case "a":
            for fi34 := 0; fi34 < 6; fi34++ {
                v14 = v14
                v35 := v3["hello"]
                _ = fi34
                _ = v35
            }
        }
        var v36 uint8 = (v2 * (v2 >> ((525) & 63)))
        _ = fi32
        _ = v36
    }
    var v37 int32 = (v1 + (v1 / ((v14) | 1)))
    switch "foo bar" {
    case "\n\t":
        var v38 uint16 = 1
        _ = v38
    }
    if ((int8(v2) == 42) || false) {
        var v39 float32 = (3.14159 - float32(-2.25))
        for fi40 := 0; fi40 < 20; fi40++ {
            v39 = float32(0.1)
            if (598 == 65535) {
                for fi41 := 0; fi41 < 15; fi41++ {
                    _ = len(v3)
                    v2 += (v2 | v2)
                    v4 = -1
                    _ = fi41
                }
            }
            v30 -= -756
            _ = fi40
        }
        v30 += (-(v30))
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v4
}
