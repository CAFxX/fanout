package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 int64 }
func genFn0(p0 bool, p1 uint) int32 {
    if p0 {
        p0 = p0
        for fi5 := 0; fi5 < 4; fi5++ {
            v6 := []int{(fi5 * (fi5 * 128))}
            _ = fi5
            _ = v6
        }
    } else {
        switch "foo bar" {
        case "a":
            p1 = 65535
        case "foo bar":
            v7 := GenS1{F0: 42, F1: 789, F2: p0}
            _ = v7
        }
        var v8 uint = p1
        fmt.Println(p1)
        _ = v8
    }
    var v9 int = 2147483647
    fmt.Println(p0, p1)
    v9 = 32767
    if ((("" != "foo bar") || (3.14159 == (float32(1.5) - (1e10 + 1e10)))) || (!p0)) {
        if p0 {
            v10 := []int{}
            p0 = (v9 > (^(v9)))
            if (!(p0 && (0 > int32(v9)))) {
                var v11 int = (^(v9))
                var v12 int16 = -32768
                _ = v11
                _ = v12
            }
            _ = v10
        }
    } else {
        var v13 float32 = (1.5 - (0.1 / 1.5))
        var v14 bool = false
        switch "foo bar" {
        case "\n\t":
            var v15 int32 = -2147483648
            _ = v15
        }
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v9
    return -119
}
func genFn1(p0 int64, p1 bool, p2 int) uint64 {
    p1 = p1
    v16 := []int{128, 0, p2}
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int64 = p0
        for fi19 := 0; fi19 < 15; fi19++ {
            if (42 == 510) {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                var v20 float64 = 0.0
                _ = v20
            } else {
                var v21 string = "a"
                if (42 <= -60) {
                    var v22 int = fi19
                    _ = v22
                }
                _ = v21
            }
            if ((("" + "") + "") > "") {
                v23 := GenS2{F0: 127}
                _ = len(v16)
                if len(v16) > 0 {
                    v16[0] = -326
                }
                _ = v23
            } else {
                v24 := GenS1{F0: 128, F1: fi17, F2: (32767 < 2)}
                _ = v24
            }
            if (42 >= 0) {
                var v25 int16 = 555
                p2 -= 9223372036854775807
                v26 := GenS3{F0: p1}
                _ = v25
                _ = v26
            } else {
                var v27 float32 = 3.14159
                _ = v27
            }
            p1 = p1
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    v28 := GenS2{F0: -12}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v28
    return 4294967295
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    v29 := GenS4{F0: 32767}
    fmt.Println(v1)
    v30 := GenS5{F0: 9223372036854775807}
    fmt.Println(v1)
    if true {
        v1 = -1
        v1 -= v1
    } else {
        var v31 int16 = 1
        _ = v31
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        switch (-(fi32)) {
        case 0:
            if (!(!true)) {
                v33 := v2[42]
                _ = v33
            }
        default:
            var v34 float32 = 3.14159
            _ = v34
        }
        v1 += v1
        _ = fi32
    }
    var v35 int32 = int32(v1)
    if false {
        fmt.Println(v35)
        if false {
            var v36 int = -698
            _ = v36
        } else {
            for fi37 := 0; fi37 < 1; fi37++ {
                v38 := GenS2{F0: 1}
                var v39 int = (fi37 - 2147483647)
                v39 = (v39 | -9223372036854775808)
                _ = fi37
                _ = v38
                _ = v39
            }
            for fi40 := 0; fi40 < 11; fi40++ {
                if false {
                    var v41 uint32 = 2147483648
                    _ = v41
                }
                switch (("" + ("" + "hello")) + "hello") {
                case "a":
                    fmt.Println(fi40, fi40)
                case "foo bar":
                    v35 = int32(v1)
                default:
                    if true {
                        var v42 uint16 = 42
                        _ = v42
                    }
                }
                v35 -= 464
                _ = fi40
            }
            var v43 uint32 = 2147483648
            _ = v43
        }
        for fi44 := 0; fi44 < 13; fi44++ {
            if false {
                fmt.Println(fi44, v1, v1)
                if false {
                    v1 += -9223372036854775808
                }
            }
            if true {
                for fi45 := 0; fi45 < 14; fi45++ {
                    fmt.Println(fi44, fi45, fi44)
                    switch "\n\t" {
                    case "foo bar":
                        var v46 uint16 = 42
                        _ = v46
                    }
                    _ = fi45
                }
                v47 := GenS6{F0: 95, F1: (v1 ^ -9223372036854775808)}
                _ = v47
            }
            _ = fi44
        }
    } else {
        switch "foo bar" {
        case "hello":
            var v48 uint32 = 65535
            _ = v48
        case "foo bar":
            var v49 uint = 4294967295
            _ = v49
        case "":
            v50 := GenS4{F0: -1}
            _ = v50
        }
        for fi51 := 0; fi51 < 4; fi51++ {
            for fi52 := 0; fi52 < 6; fi52++ {
                fmt.Println(fi51, v1)
                var v53 bool = true
                var v54 bool = false
                _ = fi52
                _ = v53
                _ = v54
            }
            v35 += (-(v35))
            _ = fi51
        }
    }
    fmt.Println(v35)
    var v55 bool = false
    v35 += genFn0((1e10 <= ((((float32(0.1) / 3.14159) / (-2.25 * 1e10)) + 1e10) / (-2.25 - (1.5 + (3.14159 * -2.25))))), 18446744073709551615)
    fmt.Println(v55)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v55
}
