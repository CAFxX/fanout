package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint; F2 uint64 }
func genFn0(p0 uint16, p1 int64) int32 {
    if false {
        var v6 uint64 = 18446744073709551615
        var v7 int8 = -128
        var v8 uint16 = p0
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int16 = 149
            v12 := []int{}
            p0 += (p0 * (p0 - p0))
            var v13 uint64 = 0
            if false {
                var v14 int64 = p1
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            p1 -= (-(p1))
            var v16 bool = true
            _ = fi15
            _ = v16
        }
        _ = fi9
    }
    v17 := GenS1{F0: 2}
    if ((true && false) || false) {
        var v18 int8 = 69
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v17
    return 2147483647
}

func main() {
    v1 := []string{}
    v2 := []int{127}
    var v3 int32 = 2147483647
    var v4 bool = true
    v5 := make(map[uint32]bool)
    if false {
        var v19 float32 = -2.25
        for fi20 := 0; fi20 < 16; fi20++ {
            v4 = true
            _ = len(v5)
            _ = fi20
        }
        _ = v19
    } else {
        delete(v5, 668)
    }
    v3 = (v3 + v3)
    for fi21 := 0; fi21 < 17; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            if true {
                var v23 string = ""
                v4 = (1 <= -1)
                v23 = "\n\t"
                _ = v23
            }
            fmt.Println(v4, v4)
            switch fi22 {
            case 561:
                v24 := GenS1{F0: 4294967295}
                _ = v24
            case 9223372036854775807:
                switch (((("\n\t" + "") + ("\n\t" + "")) + "\n\t") + "a") {
                case "hello":
                    v25 := v5[65535]
                    _ = v25
                case "a":
                    v4 = (-2.25 < ((1e10 / (float32(0.1) + -2.25)) / 1e10))
                case "\n\t":
                    switch (("a" + "hello") + (("hello" + "foo bar") + "a")) {
                    case "foo bar":
                        if false {
                            v3 = genFn0(42, -9223372036854775808)
                            v3 -= genFn0(255, -9223372036854775808)
                        } else {
                            var v26 int16 = 1
                            v27 := GenS2{F0: 1, F1: 2}
                            fmt.Println(v4, v4)
                            _ = v26
                            _ = v27
                        }
                    default:
                        fmt.Println(fi21, fi22)
                    }
                }
            }
            _ = fi22
        }
        switch "" {
        case "a":
            if (!(447 == 1000)) {
                switch (^(fi21)) {
                case 9223372036854775807:
                    v3 = (v3 ^ (v3 + -1))
                case 4294967295:
                    fmt.Println(v3, v4, fi21)
                }
                v3 = v3
            }
        }
        if len(v2) > 0 {
            v2[0] = (^(fi21))
        }
        _ = fi21
    }
    var v28 int64 = 1
    var v29 float32 = float32(1.5)
    for fi30 := 0; fi30 < 9; fi30++ {
        for fi31 := 0; fi31 < 2; fi31++ {
            var v32 uint64 = 18446744073709551615
            var v33 string = "foo bar"
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    _ = len(v5)
    _ = len(v5)
    v34 := GenS1{F0: 18446744073709551615}
    fmt.Println(v4, v28)
    if (!true) {
        if true {
            var v35 uint8 = 1
            v3 = 1
            _ = v35
        } else {
            _ = len(v5)
            var v36 bool = v4
            var v37 float64 = (1.5 + -2.25)
            _ = v36
            _ = v37
        }
        for fi38 := 0; fi38 < 13; fi38++ {
            v39 := GenS3{F0: 294, F1: 18446744073709551615, F2: 96}
            _ = fi38
            _ = v39
        }
    } else {
        fmt.Println(v3)
        for fi40 := 0; fi40 < 7; fi40++ {
            var v41 uint8 = 0
            v1 = append(v1, "")
            _ = fi40
            _ = v41
        }
        var v42 int16 = 42
        _ = v42
    }
    v3 = (v3 - -2147483648)
    if v4 {
        for fi43 := 0; fi43 < 19; fi43++ {
            switch "\n\t" {
            case "\n\t":
                for fi44 := 0; fi44 < 9; fi44++ {
                    if (!(506 == 0)) {
                        if true {
                            v45 := GenS1{F0: 785}
                            var v46 float64 = -2.25
                            _ = len(v2)
                            _ = v45
                            _ = v46
                        }
                        v29 += (float32(float64(0.1)) + float32(1.5))
                    }
                    v47 := GenS2{F0: 42, F1: 336}
                    var v48 int16 = 921
                    _ = fi44
                    _ = v47
                    _ = v48
                }
            }
            _ = fi43
        }
        v3 -= v3
    }
    var v49 int32 = 843
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v49
    _ = v5
}
