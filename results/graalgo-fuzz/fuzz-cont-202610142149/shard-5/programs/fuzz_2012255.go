package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 uint32; F2 int }
func genFn0(p0 uint32) string {
    var v5 int8 = 42
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: 720}
        v5 -= v5
        v8 := []int{(fi6 & fi6), fi6, (^(fi6)), (fi6 ^ fi6)}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 128
    switch "foo bar" {
    case "foo bar":
        if (false || false) {
            for fi10 := 0; fi10 < 18; fi10++ {
                v9 = v9
                _ = fi10
            }
            if (0 <= -32768) {
                v11 := GenS1{F0: 2}
                _ = v11
            } else {
                v12 := []int{721, -9223372036854775808, 0, -68}
                v13 := GenS2{F0: 158, F1: 137}
                for fi14 := 0; fi14 < 14; fi14++ {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        switch "hello" {
                        case "\n\t":
                            if (true || ((v5 | v5) > (v5 - v5))) {
                                fmt.Println(p0, fi14, fi14)
                                var v16 string = (("a" + "hello") + "hello")
                                _ = v16
                            }
                        default:
                            var v17 uint16 = 168
                            _ = v17
                        }
                        var v18 uint = 65535
                        var v19 bool = true
                        _ = fi15
                        _ = v18
                        _ = v19
                    }
                    var v20 float32 = (0.0 + 0.1)
                    _ = fi14
                    _ = v20
                }
                _ = v12
                _ = v13
            }
        }
    case "":
        if (2147483647 >= 2147483647) {
            var v21 uint = 1
            p0 += (p0 & p0)
            fmt.Println(p0, v9, p0)
            _ = v21
        } else {
            switch "foo bar" {
            case "hello":
                p0 -= p0
            }
            v22 := GenS1{F0: 1}
            _ = v22
        }
    }
    _ = p0
    _ = v5
    _ = v9
    return "a"
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 42
    v3 := make(map[int]int)
    var v4 string = (((("foo bar" + "foo bar") + "\n\t") + "") + "")
    v23 := GenS3{F0: 1}
    for fi24 := 0; fi24 < 10; fi24++ {
        v2 = -9223372036854775808
        if (-91 < -128) {
            fmt.Println(fi24, v4)
            var v25 uint = 42
            v4 = v4
            _ = v25
        }
        if (true || true) {
            v4 = v4
            var v26 int16 = 0
            fmt.Println(v26, v4)
            _ = v26
        }
        _ = fi24
    }
    var v27 int = 32768
    if true {
        if (357 == 132) {
            switch (^(v27)) {
            case 32767:
                for fi28 := 0; fi28 < 18; fi28++ {
                    fmt.Println(v27)
                    _ = fi28
                }
            }
            switch v27 {
            case 467:
                delete(v1, v4)
            default:
                var v29 uint32 = 2
                _ = v29
            }
        }
        var v30 float64 = 3.14159
        var v31 uint8 = 0
        _ = v30
        _ = v31
    }
    if (!true) {
        if (65535 > 42) {
            v32 := GenS4{F0: 42, F1: 65535, F2: -32768}
            var v33 int = (v27 / (((^(v27))) | 1))
            _ = v32
            _ = v33
        } else {
            var v34 uint = 186
            fmt.Println(v2, v2)
            v35 := []int{v27}
            _ = v34
            _ = v35
        }
        v36 := GenS3{F0: 1000}
        if false {
            switch v27 {
            case -1:
                if true {
                    v37 := []int{}
                    _ = v37
                } else {
                    var v38 int8 = -96
                    v39 := []int{(v27 - -9223372036854775808), v27, (v27 / (((v27 | (v27 ^ v27))) | 1)), -318}
                    _ = v38
                    _ = v39
                }
            case 128:
                switch (v27 / ((v27) | 1)) {
                case 32768:
                    v27 = (v27 % ((32767) | 1))
                case 1:
                    if (!false) {
                        v2 += v2
                    } else {
                        v27 -= -105
                        v40 := GenS5{F0: 962}
                        switch (v27 / ((v27) | 1)) {
                        case 42:
                            fmt.Println(v2)
                        }
                        _ = v40
                    }
                }
            }
        }
        _ = v36
    } else {
        var v41 float32 = 1e10
        if (v4 < v4) {
            for fi42 := 0; fi42 < 13; fi42++ {
                v27 -= 32767
                var v43 uint64 = 774
                _ = fi42
                _ = v43
            }
            var v44 bool = (0 <= 32767)
            _ = v44
        } else {
            fmt.Println(v4)
            fmt.Println(v41, v27)
            var v45 string = v4
            _ = v45
        }
        v46 := GenS6{F0: (-(v2))}
        _ = v41
        _ = v46
    }
    if (255 <= 1) {
        for fi47 := 0; fi47 < 18; fi47++ {
            v4 = ""
            _ = fi47
        }
        for fi48 := 0; fi48 < 1; fi48++ {
            var v49 int16 = int16(fi48)
            _ = fi48
            _ = v49
        }
        fmt.Println(v4, v4, v4)
    }
    v50 := GenS3{F0: 0}
    v4 = (v4 + "hello")
    var v51 int = (v27 - v27)
    switch v51 {
    case -2147483648:
        v51 -= (v27 * v27)
    case 740:
        if (225 <= 142) {
            v51 += 4294967295
            var v52 uint32 = 0
            _ = v52
        }
    }
    if false {
        switch v4 {
        case "foo bar":
            v27 = v27
        case "":
            v2 -= v2
        default:
            v53 := GenS7{F0: 42, F1: 1, F2: 9223372036854775807}
            _ = v53
        }
    } else {
        if ((42 == -1) || ((v27 / ((v27) | 1)) != 127)) {
            var v54 bool = false
            var v55 float32 = -2.25
            _ = v54
            _ = v55
        }
    }
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v50
    _ = v51
}
