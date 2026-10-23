package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int16 }
type GenS2 struct { F0 int8; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint8 }
func genFn0() int16 {
    var v6 float64 = -2.25
    if (-1 < 536) {
        if true {
            switch "hello" {
            case "hello":
                switch 32768 {
                case 2147483647:
                    v6 += 1e10
                case -826:
                    for fi7 := 0; fi7 < 18; fi7++ {
                        v8 := GenS1{F0: 0, F1: fi7, F2: 0}
                        _ = fi7
                        _ = v8
                    }
                case 590:
                    var v9 float32 = 0.1
                    _ = v9
                }
            case "\n\t":
                var v10 int32 = -252
                _ = v10
            default:
                v11 := []int{}
                _ = v11
            }
        }
    } else {
        var v12 uint16 = 65535
        if (true && ((v6 + v6) >= float64(3.14159))) {
            v13 := []int{127, 9223372036854775807, -9223372036854775808}
            fmt.Println(v6, v12)
            v6 = v6
            _ = v13
        }
        v14 := GenS1{F0: 42, F1: 582, F2: 1000}
        _ = v12
        _ = v14
    }
    fmt.Println(v6)
    v15 := GenS2{F0: 0, F1: 1, F2: 164}
    v6 += -2.25
    _ = v15
    _ = v6
    return 42
}
func genFn1() int64 {
    var v16 uint64 = 2
    if (-1 > -522) {
        v17 := []int{-412}
        v16 = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            if (596 > 42) {
                if false {
                    if (-799 < 1) {
                        var v20 string = ("a" + ("" + ("a" + "a")))
                        _ = v20
                    }
                    var v21 uint32 = 2
                    _ = v21
                }
                v22 := GenS3{F0: false}
                var v23 int16 = -522
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        var v24 string = ""
                        _ = v24
                    case "":
                        v16 -= v16
                    case "foo bar":
                        v25 := []int{(fi19 ^ (-(fi18))), (^(fi19)), (fi19 >> ((1) & 63)), (fi18 - (^(fi19)))}
                        _ = v25
                    }
                    if true {
                        v23 -= (v23 - genFn0())
                    }
                } else {
                    switch (fi18 - (fi19 - (fi19 >> ((18446744073709551615) & 63)))) {
                    case 32767:
                        v26 := []int{(-(fi19)), (-(fi18)), fi18}
                        _ = v26
                    }
                }
                _ = v22
                _ = v23
            }
            v16 = 913
            _ = fi19
        }
        _ = fi18
    }
    v27 := GenS4{F0: -1, F1: 2, F2: -1}
    v28 := []int{1, -9223372036854775808}
    _ = v16
    _ = v27
    _ = v28
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 0
    v3 := make(map[string]int)
    var v4 uint = (^(v2))
    v5 := make(map[string]uint32)
    var v29 float64 = float64(1e10)
    for fi30 := 0; fi30 < 14; fi30++ {
        if true {
            var v31 float64 = ((v29 + v29) / float64(float32((v29 - -2.25))))
            v32 := []int{1, fi30}
            _ = v31
            _ = v32
        }
        _ = fi30
    }
    if true {
        if false {
            if false {
                v29 += 1e10
                var v33 int8 = -1
                if ("a" != ("hello" + "a")) {
                    var v34 float32 = 1e10
                    _ = v34
                }
                _ = v33
            }
            v29 -= v29
        }
    }
    var v35 int16 = 1
    v29 = (v29 / -2.25)
    _ = len(v5)
    for fi36 := 0; fi36 < 13; fi36++ {
        var v37 uint = (v2 & 255)
        _ = fi36
        _ = v37
    }
    var v38 int16 = (-(v35))
    v39 := []int{328, 4294967295, 127, -97}
    v4 = 102
    v38 = v38
    var v40 int = -4
    for fi41 := 0; fi41 < 20; fi41++ {
        v38 = genFn0()
        _ = fi41
    }
    var v42 string = "a"
    for fi43 := 0; fi43 < 13; fi43++ {
        var v44 int16 = v38
        for fi45 := 0; fi45 < 3; fi45++ {
            var v46 int = v40
            _ = fi45
            _ = v46
        }
        _ = fi43
        _ = v44
    }
    _ = len(v39)
    if (v4 >= (v2 + (^(v2)))) {
        for fi47 := 0; fi47 < 11; fi47++ {
            v38 = 1
            v48 := GenS1{F0: 4294967295, F1: (v40 << ((165) & 63)), F2: (v35 & (-(v35)))}
            _ = fi47
            _ = v48
        }
        fmt.Println(v4)
    } else {
        v29 = (0.1 * (((1.5 + v29) - v29) - v29))
        v49 := GenS5{F0: 0}
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
