package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint }
func genFn0() uint {
    if (!(float32(3.14159) > 0.1)) {
        var v6 int = 42
        v7 := GenS1{F0: ((!(!((float64(0.0) - 0.1) != 0.0))) && ("foo bar" <= "a")), F1: 76, F2: 18446744073709551615}
        _ = v6
        _ = v7
    } else {
        v8 := []int{-587, 677, -1}
        var v9 string = "\n\t"
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(v9, v9)
            switch v9 {
            case "foo bar":
                fmt.Println(fi10, v9)
            case "":
                v9 = v9
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    var v11 float64 = 0.1
    _ = v11
    return 2
}
func genFn1(p0 uint16, p1 uint, p2 string) int16 {
    p2 = (p2 + p2)
    var v12 string = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 42
    var v3 int8 = -43
    var v4 uint = 255
    v5 := make(map[string]string)
    if (838 < 1) {
        var v13 uint8 = 42
        _ = v13
    }
    var v14 int = -1
    if false {
        switch "a" {
        case "foo bar":
            switch "foo bar" {
            case "foo bar":
                v15 := v5["hello"]
                _ = v15
            case "":
                var v16 float64 = 0.1
                _ = v16
            default:
                v17 := GenS2{F0: 77}
                _ = v17
            }
        case "":
            for fi18 := 0; fi18 < 3; fi18++ {
                for fi19 := 0; fi19 < 1; fi19++ {
                    if (v4 > (v4 | (^(v4)))) {
                        v20 := GenS2{F0: (^(v4))}
                        _ = v20
                    } else {
                        if ((v2 > v2) && false) {
                            if (false || (!true)) {
                                v2 = v2
                                switch "hello" {
                                case "":
                                    v21 := []int{int(v3), 42, 0, 128}
                                    _ = v21
                                }
                            }
                            v4 += v4
                            var v22 float64 = -2.25
                            _ = v22
                        }
                        if (128 != (fi19 << (((v4 + genFn0())) & 63))) {
                            fmt.Println(fi18, fi18)
                        }
                    }
                    fmt.Println(v2, fi18)
                    _ = fi19
                }
                _ = fi18
            }
        case "hello":
            var v23 uint8 = 2
            _ = v23
        }
        for fi24 := 0; fi24 < 12; fi24++ {
            fmt.Println(v2)
            fmt.Println(v4, v2)
            v25 := v1[1]
            _ = fi24
            _ = v25
        }
        v26 := GenS2{F0: v4}
        _ = v26
    }
    fmt.Println(v4, v3, v14)
    var v27 int = v14
    for fi28 := 0; fi28 < 9; fi28++ {
        if false {
            if ((1e10 * -2.25) < (((3.14159 + 0.1) + float64(-2.25)) * 3.14159)) {
                v27 = (v14 | 257)
                fmt.Println(v14, v3)
            } else {
                if (42 < 769) {
                    fmt.Println(v27, v2, v4)
                }
                if false {
                    fmt.Println(v14)
                    v29 := GenS1{F0: true, F1: 42, F2: 584}
                    _ = v29
                }
            }
            v4 = 65535
        } else {
            switch -148 {
            case 0:
                fmt.Println(v27)
            default:
                fmt.Println(fi28)
            }
            v2 = v2
        }
        v4 -= 419
        _ = fi28
    }
    fmt.Println(v27, v27, v2)
    if false {
        v2 += 42
        v30 := GenS2{F0: v4}
        _ = v30
    } else {
        v3 += v3
        v31 := GenS2{F0: (v4 | 0)}
        _ = v31
    }
    v32 := []int{v27}
    switch "" {
    case "hello":
        var v33 uint64 = 128
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
