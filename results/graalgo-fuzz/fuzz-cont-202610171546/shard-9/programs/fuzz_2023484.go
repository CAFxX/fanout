package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint64; F2 int32 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0(p0 int8, p1 string, p2 int8) uint8 {
    switch 2147483647 {
    case 4294967295:
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{0, -9223372036854775808}
            _ = fi7
            _ = v8
        }
    case 501:
        if false {
            switch (p1 + "\n\t") {
            case "hello":
                var v9 int16 = 1
                _ = v9
            case "a":
                fmt.Println(p2, p1)
            }
            p0 = p2
            p0 -= p2
        } else {
            for fi10 := 0; fi10 < 12; fi10++ {
                switch (p1 + p1) {
                case "hello":
                    if (-147 > 2147483647) {
                        if (false && (((true && (1 <= 1)) || true) && true)) {
                            v11 := []int{fi10, fi10, 32767}
                            for fi12 := 0; fi12 < 3; fi12++ {
                                var v13 int = fi12
                                _ = len(v11)
                                _ = fi12
                                _ = v13
                            }
                            _ = v11
                        } else {
                            fmt.Println(p0, p2, p1)
                            p2 = p2
                        }
                        v14 := []int{(fi10 - -2147483648), fi10, fi10}
                        var v15 bool = (-991 > 42)
                        _ = v14
                        _ = v15
                    }
                default:
                    var v16 uint16 = 65535
                    _ = v16
                }
                p2 = 46
                for fi17 := 0; fi17 < 6; fi17++ {
                    p1 = p1
                    _ = fi17
                }
                _ = fi10
            }
            fmt.Println(p1, p2)
            for fi18 := 0; fi18 < 13; fi18++ {
                fmt.Println(p1, fi18)
                _ = fi18
            }
        }
    }
    var v19 int8 = (p0 * (p0 / ((p0) | 1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 20
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 1
    var v3 uint16 = 0
    var v4 float32 = float32(1e10)
    var v5 uint64 = 826
    v6 := make(map[int]bool)
    switch 694 {
    case 315:
        v20 := GenS1{F0: 255, F1: 255}
        _ = v20
    case 4294967295:
        switch "foo bar" {
        case "hello":
            for fi21 := 0; fi21 < 8; fi21++ {
                v3 -= v3
                _ = fi21
            }
        case "a":
            switch (v2 & (v2 / ((v2) | 1))) {
            case -617:
                v4 = 1.5
            case -2147483648:
                var v22 int32 = 1
                _ = v22
            }
        case "\n\t":
            if false {
                switch v2 {
                case -1:
                    v23 := GenS1{F0: 153, F1: 42}
                    _ = v23
                case 115:
                    var v24 uint = 690
                    _ = v24
                case 9223372036854775807:
                    if true {
                        v5 = (v5 | 2)
                    }
                }
            }
        }
    case -2147483648:
        v25 := []int{}
        _ = v25
    default:
        switch (v2 - (v2 | (^(v2)))) {
        case -201:
            v4 -= v4
        case 32767:
            v26 := GenS2{F0: (0 >= -50), F1: (v5 & 280), F2: 2147483647}
            _ = v26
        default:
            v4 = v4
        }
    }
    v27 := GenS2{F0: (161 == 0), F1: (v5 & v5), F2: 42}
    v2 = (-(v2))
    v28 := []int{}
    var v29 float64 = 0.0
    var v30 float64 = v29
    switch (^(v2)) {
    case 4294967295:
        v3 = v3
    case 1:
        v28 = append(v28, 42)
    case 128:
        v29 -= (v29 - v29)
    }
    var v31 bool = false
    switch "foo bar" {
    case "hello":
        v32 := GenS3{F0: 42, F1: 0}
        _ = v32
    case "a":
        var v33 int8 = 26
        _ = v33
    case "foo bar":
        v29 = 0.1
    default:
        for fi34 := 0; fi34 < 9; fi34++ {
            v35 := GenS1{F0: 42, F1: genFn0(100, "hello", -100)}
            switch (("hello" + "a") + "") {
            case "":
                for fi36 := 0; fi36 < 1; fi36++ {
                    v3 = (v3 / ((0) | 1))
                    v37 := GenS4{F0: -96, F1: -32768}
                    var v38 int32 = 185
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            default:
                v30 += v30
            }
            _ = fi34
            _ = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
