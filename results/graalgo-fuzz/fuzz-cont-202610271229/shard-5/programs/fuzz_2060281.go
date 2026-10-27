package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int8; F1 int64; F2 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 uint }
func genFn0(p0 int32) int {
    if false {
        p0 -= p0
    }
    var v4 int64 = 0
    var v5 uint32 = 2
    fmt.Println(v5, p0, v4)
    p0 = p0
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1() string {
    switch ("" + "\n\t") {
    case "":
        for fi6 := 0; fi6 < 2; fi6++ {
            if (0.0 == 0.0) {
                var v7 int = (fi6 << ((65535) & 63))
                _ = v7
            }
            var v8 int8 = 33
            switch (fi6 % ((fi6) | 1)) {
            case -845:
                if (42 > 351) {
                    if false {
                        switch "\n\t" {
                        case "hello":
                            if (!(79 >= 255)) {
                                v9 := []int{308}
                                _ = v9
                            } else {
                                v10 := GenS1{F0: 2147483647, F1: genFn0(-1)}
                                v11 := []int{fi6}
                                _ = v10
                                _ = v11
                            }
                        case "a":
                            fmt.Println(fi6)
                        default:
                            v12 := []int{(fi6 >> ((289) & 63)), (fi6 / ((32768) | 1))}
                            _ = v12
                        }
                        var v13 int = fi6
                        if ("foo bar" <= (("" + "a") + "foo bar")) {
                            var v14 int64 = 1
                            _ = v14
                        }
                        _ = v13
                    }
                }
            case 42:
                v8 += (v8 & (v8 ^ v8))
            }
            _ = fi6
            _ = v8
        }
    case "foo bar":
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 uint16 = 2
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    case "\n\t":
        v18 := []int{}
        _ = v18
    }
    switch "\n\t" {
    case "hello":
        v19 := []int{332, 32767}
        _ = v19
    case "a":
        v20 := []int{}
        _ = v20
    default:
        for fi21 := 0; fi21 < 12; fi21++ {
            fmt.Println(fi21)
            _ = fi21
        }
    }
    if (true || (!(9223372036854775807 == 1))) {
        if ((true && true) && (false && true)) {
            v22 := []int{128, 9223372036854775807, 127, 32767}
            if true {
                var v23 uint = 956
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 float32 = -2.25
                    _ = fi24
                    _ = v25
                }
                v23 = 1
                _ = v23
            }
            _ = v22
        }
        v26 := GenS1{F0: 0, F1: 269}
        _ = v26
    } else {
        var v27 int64 = -9223372036854775808
        v28 := GenS2{F0: 1, F1: false}
        fmt.Println(v27)
        _ = v27
        _ = v28
    }
    v29 := []int{}
    v30 := GenS3{F0: 107, F1: 100}
    _ = v29
    _ = v30
    return (("" + ("" + "a")) + "")
}

func main() {
    var v1 float64 = -2.25
    var v2 uint16 = 255
    var v3 string = "hello"
    for fi31 := 0; fi31 < 16; fi31++ {
        if false {
            v1 -= v1
        }
        var v32 int32 = -453
        _ = fi31
        _ = v32
    }
    fmt.Println(v3, v2)
    if (1 == 1) {
        if true {
            v3 = v3
        } else {
            v33 := []int{-654, 42, 32768}
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = v33
        }
        var v34 float32 = -2.25
        fmt.Println(v34, v3, v34)
        _ = v34
    }
    v35 := GenS4{F0: 100, F1: 718, F2: -128}
    v36 := []int{42, genFn0(0), -253}
    switch v3 {
    case "":
        if (1 == 32767) {
            for fi37 := 0; fi37 < 6; fi37++ {
                for fi38 := 0; fi38 < 18; fi38++ {
                    var v39 int = fi37
                    var v40 string = "\n\t"
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                fmt.Println(v3, v1)
                _ = fi37
            }
            for fi41 := 0; fi41 < 18; fi41++ {
                var v42 int64 = 1
                _ = fi41
                _ = v42
            }
        }
    case "hello":
        v43 := GenS5{F0: 963}
        _ = v43
    case "foo bar":
        if true {
            if true {
                var v44 int16 = 42
                _ = v44
            }
        }
    }
    var v45 int64 = -590
    switch 25 {
    case 0:
        if ((!(!(441 < 42))) && (v45 >= 0)) {
            v46 := GenS6{F0: -128, F1: 655}
            _ = v46
        } else {
            var v47 uint = 583
            for fi48 := 0; fi48 < 20; fi48++ {
                fmt.Println(v1, fi48)
                var v49 bool = false
                v2 = (v2 + v2)
                _ = fi48
                _ = v49
            }
            _ = v47
        }
    case 32768:
        var v50 float32 = -2.25
        _ = v50
    case -283:
        var v51 uint16 = v2
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v45
}
