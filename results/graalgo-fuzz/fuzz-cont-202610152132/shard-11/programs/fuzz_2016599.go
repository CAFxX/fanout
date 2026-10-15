package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint64; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int8; F2 uint }
type GenS6 struct { F0 int }
func genFn0(p0 uint16) string {
    switch ("\n\t" + "hello") {
    case "":
        v4 := []int{4294967295, -2147483648}
        _ = v4
    case "a":
        p0 -= 1
    case "hello":
        if (true && false) {
            switch 4294967295 {
            case 1:
                var v5 string = ("foo bar" + "a")
                _ = v5
            case 32768:
                var v6 float64 = (((1.5 * float64(3.14159)) + (3.14159 * (0.1 - (0.1 / (-2.25 / 1.0))))) * 1e10)
                _ = v6
            }
            v7 := []int{127}
            _ = v7
        }
    default:
        p0 = p0
    }
    switch -706 {
    case 32768:
        for fi8 := 0; fi8 < 11; fi8++ {
            if true {
                v9 := []int{fi8, (fi8 + -910), (fi8 ^ (fi8 << ((255) & 63))), 127}
                p0 -= p0
                _ = v9
            } else {
                p0 -= p0
            }
            _ = fi8
        }
    default:
        v10 := GenS1{F0: 42, F1: 32767, F2: false}
        _ = v10
    }
    p0 = p0
    p0 = p0
    _ = p0
    return ("foo bar" + "")
}
func genFn1(p0 string, p1 bool, p2 uint) uint16 {
    p1 = p1
    switch 42 {
    case -9223372036854775808:
        p1 = (!true)
    default:
        switch "foo bar" {
        case "a":
            var v11 bool = (p0 == p0)
            _ = v11
        case "foo bar":
            v12 := GenS2{F0: (p0 + p0)}
            _ = v12
        default:
            v13 := GenS3{F0: 0, F1: 2, F2: p0}
            _ = v13
        }
    }
    p2 += p2
    switch -2147483648 {
    case -1:
        p0 = "a"
    case 42:
        v14 := []int{}
        _ = v14
    default:
        fmt.Println(p0)
    }
    var v15 float32 = (1.5 * 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 255
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = v1
    v3 := []string{}
    var v16 int16 = 401
    var v17 uint64 = 18446744073709551615
    var v18 float64 = float64(0.1)
    var v19 bool = (false && true)
    switch ("a" + "a") {
    case "a":
        var v20 int8 = -1
        _ = v20
    case "\n\t":
        if (!false) {
            for fi21 := 0; fi21 < 4; fi21++ {
                v16 -= 32767
                v2 += (-(v2))
                _ = fi21
            }
            fmt.Println(v2, v18)
        } else {
            v22 := GenS4{F0: -128}
            fmt.Println(v2)
            if v19 {
                v16 += v16
            }
            _ = v22
        }
    }
    v23 := GenS2{F0: ""}
    for fi24 := 0; fi24 < 6; fi24++ {
        v18 = ((v18 * (3.14159 * v18)) / 1.5)
        _ = fi24
    }
    v17 = 38
    var v25 uint8 = 1
    if v19 {
        v26 := GenS3{F0: 65535, F1: (^(v17)), F2: "\n\t"}
        switch -1 {
        case -1:
            fmt.Println(v25)
        case 127:
            fmt.Println(v17, v17, v2)
        }
        v16 = (v16 & v16)
        _ = v26
    } else {
        switch 32768 {
        case 350:
            v17 = (^(v17))
        case 2147483647:
            fmt.Println(v17, v16, v17)
        case -621:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            if v19 {
                for fi27 := 0; fi27 < 15; fi27++ {
                    v2 += v1
                    switch (fi27 * 0) {
                    case -9223372036854775808:
                        var v28 int = fi27
                        _ = v28
                    case 0:
                        v16 += -1
                    }
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v2 = 0
                        v1 += v2
                        _ = fi29
                    }
                    _ = fi27
                }
            } else {
                v16 = v16
                switch 4294967295 {
                case -1:
                    fmt.Println(v16)
                case -563:
                    for fi30 := 0; fi30 < 8; fi30++ {
                        v31 := GenS5{F0: (v1 - v1), F1: 127, F2: 65535}
                        _ = fi30
                        _ = v31
                    }
                case -540:
                    v32 := GenS4{F0: -128}
                    _ = v32
                default:
                    v17 -= v17
                }
                if (int8(v1) < 1) {
                    v1 += (v2 - (v2 ^ v1))
                }
            }
        }
    }
    var v33 int16 = (^(v16))
    v2 = v2
    fmt.Println(v33)
    v34 := GenS6{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v33
    _ = v34
}
