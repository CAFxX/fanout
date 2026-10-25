package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int16; F1 int; F2 string }
func genFn0(p0 uint64) int16 {
    switch ("hello" + ("foo bar" + "\n\t")) {
    case "":
        v4 := []int{32767, -1, 32767, -2147483648}
        _ = v4
    case "hello":
        switch 32768 {
        case 32768:
            p0 = p0
        case 874:
            var v5 bool = ((100 == 100) || false)
            _ = v5
        case 555:
            fmt.Println(p0)
        default:
            fmt.Println(p0)
        }
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        if (128 > 2) {
            if (!(!false)) {
                fmt.Println(fi7, fi7)
                v8 := []int{-550, (fi7 + -2147483648), fi7, fi7}
                _ = v8
            }
            var v9 float64 = (((((-2.25 / (-2.25 + float64(1e10))) - -2.25) * (-2.25 + ((-2.25 * (3.14159 * (((3.14159 / (0.1 * (0.1 - -2.25))) / (1.5 * 0.1)) - -2.25))) * (((1e10 / 1e10) * 1e10) - 0.0)))) - 1.5) - 1e10)
            if false {
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 uint8 = 227
                    var v12 int16 = 1
                    v13 := GenS1{F0: fi7, F1: -441}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                var v14 float64 = -2.25
                var v15 uint = 517
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 12; fi16++ {
                    v17 := GenS1{F0: (fi16 & 0), F1: -2147483648}
                    _ = fi16
                    _ = v17
                }
                p0 = (p0 | (^(p0)))
            }
            _ = v9
        }
        _ = fi7
    }
    if true {
        fmt.Println(p0)
        var v18 int32 = -825
        switch -873 {
        case 694:
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int8 = -69
                _ = fi19
                _ = v20
            }
        case 0:
            var v21 string = "foo bar"
            _ = v21
        case 611:
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 string = ""
                for fi24 := 0; fi24 < 13; fi24++ {
                    v25 := GenS1{F0: (fi24 ^ fi24), F1: (v18 - (v18 - v18))}
                    v23 = v23
                    var v26 bool = (true && (100 <= 0))
                    v18 = (v18 + (v18 - v18))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                var v27 int8 = 42
                _ = fi22
                _ = v23
                _ = v27
            }
        default:
            p0 -= 18446744073709551615
        }
        _ = v18
    }
    switch "foo bar" {
    case "foo bar":
        for fi28 := 0; fi28 < 4; fi28++ {
            v29 := []int{}
            _ = fi28
            _ = v29
        }
    case "a":
        var v30 uint8 = 255
        _ = v30
    default:
        v31 := GenS2{F0: 32767, F1: 637, F2: "hello"}
        _ = v31
    }
    _ = p0
    return -1
}
func genFn1() uint32 {
    var v32 uint32 = 65535
    v33 := []int{-900, 1}
    _ = v32
    _ = v33
    return (v32 ^ v32)
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    v34 := GenS2{F0: 0, F1: 500, F2: ""}
    v1 -= 0.1
    v1 += v1
    switch "hello" {
    case "\n\t":
        switch ((("" + "foo bar") + "a") + "foo bar") {
        case "\n\t":
            var v35 string = ("hello" + ("" + "\n\t"))
            _ = v35
        default:
            fmt.Println(v1)
        }
    case "a":
        fmt.Println(v1)
    case "":
        v1 += -2.25
    }
    var v36 bool = ((801 != 2) || ((!false) && (42 < 255)))
    var v37 int16 = -32768
    fmt.Println(v1)
    v38 := GenS1{F0: 32767, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v38
}
