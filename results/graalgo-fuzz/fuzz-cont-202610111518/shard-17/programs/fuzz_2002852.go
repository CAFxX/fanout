package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint32 }
func genFn0(p0 string, p1 int64, p2 string) int32 {
    p1 -= p1
    switch (p2 + "a") {
    case "a":
        var v6 uint64 = 1
        _ = v6
    case "foo bar":
        v7 := []int{-1}
        _ = v7
    case "\n\t":
        v8 := GenS1{F0: 506}
        _ = v8
    }
    var v9 uint = 890
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        switch 32767 {
        case 1:
            if (760 <= 237) {
                v9 += 18446744073709551615
                var v12 int16 = 1000
                _ = v12
            }
        default:
            switch (-(fi11)) {
            case -621:
                v13 := GenS1{F0: 18446744073709551615}
                _ = v13
            case 503:
                v14 := []int{}
                _ = v14
            default:
                if (42 != 127) {
                    switch (p0 + (p0 + p0)) {
                    case "\n\t":
                        switch "hello" {
                        case "a":
                            var v15 uint32 = 267
                            _ = v15
                        case "\n\t":
                            var v16 uint16 = 67
                            _ = v16
                        default:
                            switch fi11 {
                            case 128:
                                switch (fi11 << ((2) & 63)) {
                                case -1:
                                    var v17 uint16 = 2
                                    _ = v17
                                }
                            default:
                                p0 = ("hello" + "\n\t")
                            }
                        }
                    default:
                        p2 = p0
                    }
                }
            }
        }
        var v18 uint64 = 18446744073709551615
        _ = fi11
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    var v1 float64 = 0.0
    var v2 bool = true
    v3 := []int{0, 689}
    v4 := []int{647, -914, 127, -940}
    var v5 bool = (-28 != 42)
    for fi19 := 0; fi19 < 12; fi19++ {
        if (v2 && v5) {
            v2 = (827 < 4294967295)
            for fi20 := 0; fi20 < 6; fi20++ {
                fmt.Println(v5, fi20)
                var v21 bool = v2
                v5 = true
                _ = fi20
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 int32 = 2147483647
            _ = fi22
            _ = v23
        }
        _ = fi19
    }
    v24 := GenS1{F0: 18446744073709551615}
    v5 = (true && v5)
    for fi25 := 0; fi25 < 1; fi25++ {
        fmt.Println(fi25, v1)
        switch (fi25 / (((^(fi25))) | 1)) {
        case 32767:
            switch fi25 {
            case 870:
                if true {
                    var v26 float64 = v1
                    _ = v26
                }
            }
        case 755:
            if len(v3) > 0 {
                v3[0] = -1
            }
        case 0:
            for fi27 := 0; fi27 < 17; fi27++ {
                if (32767 <= 1) {
                    if (!(float64(float32((1.5 + 1e10))) > (v1 * float64(float32(1e10))))) {
                        var v28 uint8 = 255
                        v28 = 67
                        var v29 uint = 2
                        _ = v28
                        _ = v29
                    }
                }
                v1 += float64(3.14159)
                _ = fi27
            }
        }
        v30 := GenS2{F0: 1, F1: 4294967295}
        _ = fi25
        _ = v30
    }
    v2 = (v1 > (-2.25 + v1))
    v5 = v5
    v1 += ((0.0 / float64(3.14159)) * 0.1)
    _ = len(v3)
    v4 = append(v4, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
