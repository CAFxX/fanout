package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 int }
type GenS3 struct { F0 bool }
func genFn0(p0 bool, p1 int, p2 bool) uint16 {
    if p2 {
        p0 = p2
    } else {
        v4 := []int{}
        _ = v4
    }
    for fi5 := 0; fi5 < 15; fi5++ {
        fmt.Println(p1, p0, p2)
        _ = fi5
    }
    v6 := []int{p1}
    if p2 {
        v7 := GenS1{F0: 4294967295, F1: true}
        for fi8 := 0; fi8 < 20; fi8++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = len(v6)
            _ = fi8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int64 = 9223372036854775807
            fmt.Println(fi9)
            _ = fi9
            _ = v10
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32767
}

func main() {
    v1 := []string{"", ("foo bar" + "")}
    v2 := make(map[int]string)
    v3 := []int{0, -698}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 float32 = 0.0
        v13 := GenS1{F0: 0, F1: false}
        switch ("a" + ("\n\t" + "\n\t")) {
        case "a":
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 int16 = 32767
                v1 = append(v1, "")
                var v16 int32 = 42
                if ((1.5 >= 0.0) || false) {
                    v17 := GenS1{F0: 4294967295, F1: false}
                    switch 42 {
                    case 32767:
                        var v18 uint64 = 2
                        _ = v18
                    case 4294967295:
                        v12 -= float32(0.1)
                    case -89:
                        v16 = (v16 / (((v16 * (v16 - 425))) | 1))
                    }
                    var v19 string = ("a" + ("hello" + ""))
                    var v20 float32 = (3.14159 - 1.5)
                    _ = v17
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS1{F0: 0, F1: (1e10 <= 3.14159)}
                    var v22 uint16 = 465
                    _ = v21
                    _ = v22
                }
                var v23 uint = 42
                var v24 uint32 = 4294967295
                v23 += 42
                _ = fi14
                _ = v15
                _ = v16
                _ = v23
                _ = v24
            }
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    var v25 float64 = ((3.14159 / (-2.25 / 0.1)) + -2.25)
    v25 = 1.5
    switch "" {
    case "foo bar":
        var v26 uint16 = 255
        _ = v26
    case "hello":
        var v27 string = ("hello" + "")
        _ = v27
    case "":
        for fi28 := 0; fi28 < 17; fi28++ {
            fmt.Println(fi28)
            for fi29 := 0; fi29 < 6; fi29++ {
                v25 -= 3.14159
                _ = len(v2)
                fmt.Println(fi28, fi29)
                _ = fi29
            }
            _ = fi28
        }
    default:
        var v30 uint16 = 2
        _ = v30
    }
    v31 := GenS2{F0: 127, F1: 850, F2: 32767}
    if (-527 < 42) {
        var v32 int = -1
        if (!(true || false)) {
            switch v32 {
            case -724:
                v33 := GenS1{F0: 2, F1: ((false || true) && (0.1 >= 3.14159))}
                _ = v33
            case 42:
                if (-2147483648 < 0) {
                    v25 -= v25
                } else {
                    var v34 int = -1
                    var v35 uint = 255
                    for fi36 := 0; fi36 < 5; fi36++ {
                        var v37 uint = 611
                        _ = fi36
                        _ = v37
                    }
                    _ = v34
                    _ = v35
                }
            case 127:
                v32 = (v32 & v32)
            default:
                v32 -= v32
            }
            var v38 int = (v32 ^ 2147483647)
            v38 += v38
            _ = v38
        } else {
            var v39 string = "a"
            v40 := GenS1{F0: 42, F1: true}
            if false {
                v41 := GenS2{F0: 0, F1: 42, F2: v32}
                _ = v41
            } else {
                if true {
                    switch (v32 % ((v32) | 1)) {
                    case -1:
                        fmt.Println(v39, v32, v39)
                    case 9223372036854775807:
                        v25 = v25
                    default:
                        if (0 < 2) {
                            for fi42 := 0; fi42 < 2; fi42++ {
                                v43 := GenS2{F0: 127, F1: 0, F2: (^(fi42))}
                                var v44 uint64 = 18446744073709551615
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                            for fi45 := 0; fi45 < 3; fi45++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch (fi45 & 2147483647) {
                                case -9223372036854775808:
                                    v46 := GenS1{F0: 2147483647, F1: false}
                                    _ = v46
                                }
                                _ = fi45
                            }
                            v32 = v32
                        } else {
                            var v47 float64 = float64((1e10 - ((0.1 + 0.1) / ((1.5 + ((0.1 / 0.1) + (float32(((v25 * -2.25) + -2.25)) + 0.1))) / 0.1))))
                            _ = v47
                        }
                    }
                    var v48 int32 = 2147483647
                    v39 = ("foo bar" + "\n\t")
                    _ = v48
                }
                v32 += v32
                v49 := GenS3{F0: (-1 > -32768)}
                _ = v49
            }
            _ = v39
            _ = v40
        }
        _ = v32
    }
    fmt.Println(v25)
    var v50 int = 0
    var v51 uint32 = 2147483647
    fmt.Println(v25, v50)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v50
    _ = v51
}
