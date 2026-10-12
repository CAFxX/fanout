package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 string }
func genFn0() uint8 {
    switch 32768 {
    case 1:
        var v4 int32 = -848
        _ = v4
    case -1:
        v5 := []int{}
        _ = v5
    }
    v6 := GenS1{F0: true}
    var v7 float32 = (0.0 - -2.25)
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int8 = 42
        var v10 float64 = -2.25
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 uint = 18446744073709551615
            var v13 uint = v12
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v6
    _ = v7
    return 59
}
func genFn1(p0 uint32) int8 {
    var v14 uint32 = 1
    var v15 uint32 = 0
    v16 := GenS2{F0: 42, F1: 2}
    v15 += (p0 | 2147483647)
    switch -2147483648 {
    case -140:
        if (false || (true && true)) {
            if (float64(-2.25) < 1e10) {
                v17 := []int{}
                if false {
                    var v18 uint16 = 65535
                    _ = v18
                }
                _ = v17
            } else {
                for fi19 := 0; fi19 < 2; fi19++ {
                    p0 += 2
                    var v20 string = ("" + "")
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 19; fi21++ {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        p0 = (^(v14))
                        var v23 float64 = 0.0
                        _ = fi22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 6; fi24++ {
                        var v25 int = (fi21 ^ (fi21 & fi24))
                        var v26 string = "a"
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    _ = fi21
                }
            }
        } else {
            switch "a" {
            case "hello":
                fmt.Println(v14, p0)
            case "\n\t":
                var v27 uint = 18446744073709551615
                _ = v27
            }
            v28 := GenS2{F0: -32768, F1: 42}
            _ = v28
        }
    case 127:
        for fi29 := 0; fi29 < 18; fi29++ {
            if true {
                v14 -= (v15 - (v15 - v14))
                for fi30 := 0; fi30 < 6; fi30++ {
                    v31 := GenS3{F0: 4294967295}
                    _ = fi30
                    _ = v31
                }
            } else {
                v32 := GenS4{F0: (((-2.25 >= (((0.1 / (3.14159 - 0.1)) * (1.5 / 1e10)) + 3.14159)) && false) || ((^(fi29)) > (fi29 * (fi29 * fi29)))), F1: 1}
                _ = v32
            }
            v33 := GenS1{F0: (!true)}
            _ = fi29
            _ = v33
        }
    case 32767:
        var v34 float32 = 0.0
        _ = v34
    }
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return -1
}

func main() {
    var v1 uint = 997
    var v2 uint8 = 187
    var v3 bool = true
    v35 := []int{1, 127, -493, -261}
    v2 -= 127
    v36 := GenS5{F0: ("hello" + "")}
    v35 = append(v35, -58)
    if (123 != -1) {
        v3 = ((v2 | 255) >= (v2 - v2))
        if len(v35) >= 2 {
            v35 = v35[1:len(v35)]
        }
        v1 += (v1 - (^(v1)))
    }
    v37 := GenS3{F0: 1}
    for fi38 := 0; fi38 < 4; fi38++ {
        fmt.Println(fi38)
        if len(v35) >= 2 {
            v35 = v35[1:len(v35)]
        }
        _ = fi38
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        for fi40 := 0; fi40 < 13; fi40++ {
            switch ("foo bar" + "\n\t") {
            case "hello":
                fmt.Println(v3)
            case "foo bar":
                v41 := GenS5{F0: "hello"}
                _ = v41
            }
            _ = fi40
        }
        var v42 uint16 = 65535
        _ = fi39
        _ = v42
    }
    var v43 int64 = 1
    var v44 float64 = float64(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v43
    _ = v44
}
