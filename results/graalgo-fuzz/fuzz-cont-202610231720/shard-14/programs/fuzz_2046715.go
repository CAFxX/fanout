package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 uint64) int32 {
    fmt.Println(p0)
    var v4 string = ("a" + "\n\t")
    v5 := GenS1{F0: (0 == 25), F1: (^(p0)), F2: 42}
    v4 = v4
    _ = p0
    _ = v4
    _ = v5
    return -2147483648
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    var v6 int64 = 895
    switch "" {
    case "a":
        _ = len(v2)
    }
    fmt.Println(v6)
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 float32 = 1.5
            if true {
                v10 += v10
                var v11 float64 = 0.1
                if ((2 <= v7) || (("foo bar" + ("\n\t" + "")) != "a")) {
                    v12 := GenS1{F0: (338 <= 689), F1: 2, F2: 1}
                    fmt.Println(v11)
                    switch "a" {
                    case "foo bar":
                        var v13 int64 = (v6 + (v6 | 655))
                        fmt.Println(v6, v10)
                        v2 = append(v2, 280)
                        if true {
                            v14 := GenS1{F0: (true || (v11 >= 1e10)), F1: 18446744073709551615, F2: 18446744073709551615}
                            _ = v14
                        }
                        if true {
                            fmt.Println(v11, fi9)
                            var v15 string = "foo bar"
                            _ = v15
                        }
                        _ = v13
                    case "":
                        if (463 <= 32767) {
                            v16 := GenS1{F0: (false && (!(-2147483648 <= -1))), F1: 2, F2: 42}
                            _ = v16
                        }
                    case "\n\t":
                        v10 = (0.0 + -2.25)
                    default:
                        var v17 uint32 = 264
                        _ = v17
                    }
                    _ = v12
                }
                v7 += (v7 | v7)
                _ = v11
            } else {
                v6 = (v6 % ((-784) | 1))
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v18 string = ("hello" + ("" + ""))
    var v19 float32 = ((0.0 + (float32((3.14159 * 3.14159)) / 1e10)) + (((0.0 + (1.5 - 0.1)) + 3.14159) - float32(1e10)))
    v18 = "a"
    v20 := v3[v18]
    if (255 != 147) {
        fmt.Println(v18)
        switch "a" {
        case "":
            _ = len(v2)
        case "foo bar":
            if true {
                var v21 float64 = 0.0
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 bool = true
                    fmt.Println(v20)
                    var v24 uint = 266
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = v21
            }
        case "\n\t":
            var v25 uint = 210
            _ = v25
        default:
            var v26 int16 = 0
            _ = v26
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            v7 -= v7
            v18 = (("foo bar" + "") + (v18 + v18))
            fmt.Println(v18)
            _ = fi27
        }
    } else {
        if false {
            var v28 uint8 = 1
            switch v18 {
            case "":
                switch 32768 {
                case 32767:
                    if false {
                        if false {
                            for fi29 := 0; fi29 < 20; fi29++ {
                                v20 -= 65535
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v18 = "foo bar"
                                _ = fi29
                            }
                            v30 := GenS2{F0: 32767}
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 4; fi31++ {
                            v32 := GenS3{F0: (v6 * (v6 ^ int64(v7))), F1: 0}
                            var v33 int8 = 127
                            v28 -= (v28 - v28)
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 10; fi34++ {
                            var v35 uint64 = 164
                            for fi36 := 0; fi36 < 4; fi36++ {
                                var v37 uint16 = (v7 + 255)
                                _ = fi36
                                _ = v37
                            }
                            fmt.Println(fi34, v20, v7)
                            _ = fi34
                            _ = v35
                        }
                    }
                }
            case "hello":
                fmt.Println(v7)
            default:
                v38 := GenS2{F0: (^(v7))}
                _ = v38
            }
            _ = v28
        }
    }
    fmt.Println(v20, v7)
    var v39 uint8 = 127
    v40 := GenS1{F0: true, F1: 0, F2: 0}
    for fi41 := 0; fi41 < 6; fi41++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            v18 = (("" + "\n\t") + v18)
            v39 += 42
        }
        _ = fi41
    }
    v39 -= 2
    var v42 bool = (-2147483648 >= 2147483647)
    var v43 int = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v40
    _ = v42
    _ = v43
    _ = v6
    _ = v7
}
