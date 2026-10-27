package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 bool }
func genFn0(p0 uint64, p1 int, p2 uint8) uint8 {
    if false {
        if (2 > p0) {
            if (1 > 1) {
                v7 := GenS1{F0: false}
                v8 := GenS1{F0: (42 != 2147483647)}
                _ = v7
                _ = v8
            }
            var v9 uint16 = 0
            var v10 int32 = -2147483648
            _ = v10
            _ = v9
        }
    } else {
        for fi11 := 0; fi11 < 9; fi11++ {
            p0 = 680
            fmt.Println(p0)
            if (65535 < 255) {
                for fi12 := 0; fi12 < 2; fi12++ {
                    if ((("\n\t" + ("hello" + "hello")) + "") >= "foo bar") {
                        if (0.0 <= -2.25) {
                            var v13 uint = 2
                            _ = v13
                        } else {
                            var v14 uint8 = (p2 & (p2 & p2))
                            _ = v14
                        }
                        if true {
                            var v15 int16 = -1
                            _ = v15
                        }
                        v16 := GenS1{F0: true}
                        _ = v16
                    }
                    _ = fi12
                }
                v17 := GenS1{F0: true}
                if (true && (!(1 == 2))) {
                    var v18 uint = 65535
                    _ = v18
                }
                _ = v17
            }
            _ = fi11
        }
        var v19 bool = false
        fmt.Println(p2)
        _ = v19
    }
    var v20 uint16 = 0
    v21 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return p2
}
func genFn1(p0 uint64) int {
    if (108 <= 127) {
        for fi22 := 0; fi22 < 14; fi22++ {
            switch fi22 {
            case 127:
                if (!true) {
                    p0 = (p0 | 4294967295)
                    switch fi22 {
                    case 1:
                        v23 := GenS2{F0: true, F1: ((float32(float64(float32((0.0 * 3.14159)))) / ((1.5 + 1e10) / -2.25)) > 1e10)}
                        _ = v23
                    }
                    v24 := []int{-2147483648, fi22, (fi22 + (fi22 + (-(fi22))))}
                    _ = v24
                }
            case -204:
                v25 := []int{fi22, -1}
                _ = v25
            case 1:
                v26 := []int{(fi22 | fi22), 9223372036854775807}
                _ = v26
            }
            if (!true) {
                for fi27 := 0; fi27 < 8; fi27++ {
                    v28 := []int{340, fi27}
                    fmt.Println(p0, fi27, fi22)
                    _ = fi27
                    _ = v28
                }
            }
            var v29 bool = true
            _ = fi22
            _ = v29
        }
        var v30 uint8 = 42
        _ = v30
    }
    var v31 int8 = 42
    _ = p0
    _ = v31
    return 2147483647
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 2
    var v5 int64 = -9223372036854775808
    var v6 uint16 = 85
    v32 := GenS1{F0: ((v6 / ((v6) | 1)) >= 65535)}
    for fi33 := 0; fi33 < 20; fi33++ {
        v6 += (v6 | 674)
        for fi34 := 0; fi34 < 3; fi34++ {
            if (!(1 == (v6 << ((255) & 63)))) {
                fmt.Println(v6, fi34, fi33)
                fmt.Println(v5, v4)
                v6 = (^(v6))
            }
            v35 := GenS2{F0: false, F1: true}
            _ = fi34
            _ = v35
        }
        switch ("\n\t" + ("hello" + "")) {
        case "foo bar":
            fmt.Println(v1)
        case "":
            var v36 uint32 = 513
            _ = v36
        default:
            switch (fi33 | (-(fi33))) {
            case 127:
                fmt.Println(v1, v5)
            case -9223372036854775808:
                for fi37 := 0; fi37 < 6; fi37++ {
                    var v38 uint8 = (v4 % ((v4) | 1))
                    _ = fi37
                    _ = v38
                }
            case 953:
                v3[947] = 623
            }
        }
        _ = fi33
    }
    v39 := GenS1{F0: true}
    for fi40 := 0; fi40 < 13; fi40++ {
        v6 -= (v6 & (^(v6)))
        v41 := v3[0]
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        if false {
            var v43 int = (fi42 | fi42)
            for fi44 := 0; fi44 < 2; fi44++ {
                v45 := GenS1{F0: false}
                v3[2147483648] = 1
                _ = fi44
                _ = v45
            }
            _ = v43
        }
        var v46 int8 = -1
        v47 := GenS1{F0: (false || (v5 == v1))}
        _ = fi42
        _ = v46
        _ = v47
    }
    var v48 string = "foo bar"
    var v49 float32 = ((1.5 / (0.0 + (1.5 / (0.1 / (((1e10 / ((float32((0.0 - 0.0)) + -2.25) + 0.0)) + (3.14159 - 3.14159)) / 0.1))))) / -2.25)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
