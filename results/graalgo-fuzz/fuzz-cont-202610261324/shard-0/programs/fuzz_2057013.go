package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint16, p1 bool, p2 uint) bool {
    p2 = p2
    var v6 uint16 = 255
    var v7 uint64 = 18446744073709551615
    if p1 {
        for fi8 := 0; fi8 < 5; fi8++ {
            p0 += p0
            fmt.Println(fi8, p0)
            _ = fi8
        }
    } else {
        v9 := GenS1{F0: "hello", F1: p2}
        v6 = 0
        var v10 int = 9223372036854775807
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (!true)
}
func genFn1(p0 int64, p1 uint32, p2 int) uint {
    if false {
        var v11 uint = 4294967295
        _ = v11
    } else {
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 14; fi13++ {
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    var v14 uint64 = 18446744073709551615
                    _ = v14
                default:
                    var v15 uint32 = 466
                    v16 := GenS2{F0: (fi12 & (p2 - p2)), F1: ((1.5 / float64((3.14159 - (0.0 / -2.25)))) > 1.5)}
                    switch (-(p2)) {
                    case 2147483647:
                        var v17 bool = true
                        _ = v17
                    }
                    _ = v15
                    _ = v16
                }
                _ = fi13
            }
            v18 := []int{p2, -1, -9223372036854775808}
            if (!false) {
                p0 += int64(fi12)
                p2 = p2
                var v19 float64 = 1.5
                _ = v19
            }
            _ = fi12
            _ = v18
        }
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            fmt.Println(fi21, fi20)
            if (!((!((-2147483648 > -2147483648) || false)) || ((("\n\t" + "a") + ("hello" + "")) == "\n\t"))) {
                v22 := GenS3{F0: (p1 / ((p1) | 1)), F1: false}
                var v23 int = (fi21 % ((fi20) | 1))
                _ = v22
                _ = v23
            }
            var v24 bool = true
            _ = fi21
            _ = v24
        }
        var v25 int = p2
        v26 := []int{}
        _ = fi20
        _ = v25
        _ = v26
    }
    var v27 uint = 57
    v28 := GenS3{F0: (p1 / ((p1) | 1)), F1: false}
    var v29 string = ("foo bar" + "")
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    _ = v29
    return v27
}

func main() {
    var v1 uint32 = 295
    var v2 uint64 = 243
    var v3 int64 = -1
    v4 := make(map[string]int)
    var v5 int32 = 42
    v30 := GenS3{F0: v1, F1: (!false)}
    if true {
        v31 := GenS2{F0: 4294967295, F1: ((128 <= 255) || false)}
        fmt.Println(v1, v1)
        for fi32 := 0; fi32 < 18; fi32++ {
            v33 := GenS2{F0: fi32, F1: false}
            v1 = v1
            var v34 int = (^(fi32))
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v31
    } else {
        var v35 int = -288
        _ = v35
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        v37 := GenS4{F0: int16(v5), F1: 216}
        delete(v4, "")
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        switch (((("" + "hello") + "hello") + "a") + ("hello" + ("" + "hello"))) {
        case "":
            if true {
                if true {
                    var v39 string = "\n\t"
                    var v40 float64 = (1.5 * (0.0 / 0.1))
                    _ = v39
                    _ = v40
                }
            }
        case "hello":
            v3 -= 9223372036854775807
        case "\n\t":
            v2 -= v2
        default:
            fmt.Println(v5, v2)
        }
        if false {
            var v41 bool = (64 <= 127)
            var v42 uint16 = 32767
            fmt.Println(v3, v3, v42)
            _ = v41
            _ = v42
        } else {
            v43 := v4["a"]
            _ = len(v4)
            _ = v43
        }
        _ = fi38
    }
    if false {
        var v44 uint = 65535
        v2 += 582
        _ = v44
    }
    fmt.Println(v2, v1, v5)
    v45 := GenS2{F0: -587, F1: false}
    var v46 uint16 = 709
    v1 += uint32(v46)
    v1 -= (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
