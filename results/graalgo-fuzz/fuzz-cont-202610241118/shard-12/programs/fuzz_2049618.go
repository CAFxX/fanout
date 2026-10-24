package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 bool; F2 int16 }
type GenS5 struct { F0 string; F1 uint8; F2 int8 }
func genFn0(p0 uint64, p1 int16, p2 int32) uint {
    if false {
        var v6 uint64 = 0
        for fi7 := 0; fi7 < 3; fi7++ {
            p2 -= p2
            _ = fi7
        }
        if ((42 > 32767) && ((42 == 50) || (!false))) {
            for fi8 := 0; fi8 < 5; fi8++ {
                p2 += p2
                p1 = p1
                v9 := GenS1{F0: (p1 >> ((4294967295) & 63)), F1: (v6 | v6)}
                _ = fi8
                _ = v9
            }
            var v10 uint = 42
            var v11 uint16 = 1
            _ = v10
            _ = v11
        } else {
            v12 := GenS2{F0: 2147483647, F1: 818}
            p2 = p2
            v13 := GenS3{F0: true}
            _ = v12
            _ = v13
        }
        _ = v6
    } else {
        v14 := []int{127, -9223372036854775808}
        var v15 uint8 = 128
        if (!true) {
            var v16 uint64 = (^(p0))
            _ = v16
        } else {
            for fi17 := 0; fi17 < 15; fi17++ {
                p2 = (^(p2))
                _ = fi17
            }
        }
        _ = v14
        _ = v15
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        p2 = (p2 % ((-858) | 1))
        v19 := GenS4{F0: p1, F1: (9223372036854775807 >= -917), F2: p1}
        _ = fi18
        _ = v19
    }
    var v20 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 255
}

func main() {
    var v1 bool = (1 >= -1)
    var v2 uint64 = 4294967295
    var v3 float32 = 1.5
    v4 := []int{2147483647}
    v5 := make(map[string]string)
    v2 += v2
    v5[(("\n\t" + "\n\t") + "hello")] = ("\n\t" + "\n\t")
    _ = len(v5)
    v21 := GenS2{F0: 0, F1: -1}
    var v22 uint = 815
    if false {
        switch "hello" {
        case "hello":
            v22 += genFn0(v2, -218, 2147483647)
        case "foo bar":
            v23 := GenS5{F0: "\n\t", F1: 210, F2: 127}
            _ = v23
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            fmt.Println(v2, fi24, v3)
            switch ("foo bar" + ("\n\t" + "\n\t")) {
            case "a":
                v22 = genFn0(18446744073709551615, 1, 496)
            default:
                fmt.Println(v22)
            }
            v22 = 65535
            _ = fi24
        }
        fmt.Println(v1, v22, v1)
    }
    v25 := v5["a"]
    fmt.Println(v2, v3)
    if v1 {
        if v1 {
            var v26 int8 = -128
            _ = v26
        } else {
            if true {
                if (v1 || true) {
                    _ = len(v5)
                }
                var v27 int8 = -1
                _ = v27
            }
            var v28 int8 = 127
            _ = v28
        }
    } else {
        if v1 {
            fmt.Println(v22, v3, v22)
        }
        fmt.Println(v2)
        if v1 {
            var v29 float64 = -2.25
            var v30 uint8 = 117
            _ = v29
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
