package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int16; F2 int64 }
func genFn0(p0 string, p1 uint16) int32 {
    v7 := GenS1{F0: 1}
    if false {
        switch 32767 {
        case 32768:
            fmt.Println(p1, p0)
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 57
}
func genFn1(p0 int32, p1 int) string {
    if ((!(1.5 < 1e10)) && true) {
        for fi8 := 0; fi8 < 16; fi8++ {
            if (65535 > 1) {
                var v9 uint8 = 255
                for fi10 := 0; fi10 < 18; fi10++ {
                    v11 := []int{fi10, fi10, (^(p1))}
                    _ = fi10
                    _ = v11
                }
                if (false || true) {
                    v12 := GenS1{F0: 18446744073709551615}
                    v13 := []int{(p1 % ((-251) | 1))}
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        var v14 int32 = (p0 % (((p0 / ((p0) | 1))) | 1))
                        _ = v14
                    case "a":
                        var v15 float64 = 3.14159
                        _ = v15
                    case "\n\t":
                        v16 := GenS1{F0: 2}
                        _ = v16
                    }
                    _ = v12
                    _ = v13
                } else {
                    v9 -= (v9 - (v9 - 84))
                }
                _ = v9
            }
            _ = fi8
        }
    } else {
        p0 = (p0 >> ((548) & 63))
        var v17 int16 = 0
        _ = v17
    }
    if (0.1 <= 1e10) {
        var v18 int8 = -128
        _ = v18
    } else {
        fmt.Println(p1)
    }
    p1 += p1
    _ = p0
    _ = p1
    return "hello"
}

func main() {
    var v1 int = -2147483648
    var v2 uint32 = 302
    var v3 string = (("" + "a") + (("hello" + ("foo bar" + "a")) + ("foo bar" + ("hello" + "\n\t"))))
    var v4 string = v3
    var v5 int64 = -1
    var v6 int8 = 106
    for fi19 := 0; fi19 < 4; fi19++ {
        if (!true) {
            if (127 != 128) {
                for fi20 := 0; fi20 < 12; fi20++ {
                    v21 := GenS1{F0: 42}
                    v22 := []int{118}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            fmt.Println(v3)
            if (!true) {
                var v23 int8 = 1
                v4 = v3
                _ = v23
            }
        }
        var v24 uint16 = uint16(v2)
        v25 := []int{32768, (fi19 << ((0) & 63)), (fi19 | v1), v1}
        _ = fi19
        _ = v24
        _ = v25
    }
    fmt.Println(v5, v6, v5)
    if false {
        v26 := GenS1{F0: 42}
        if true {
            var v27 int8 = v6
            _ = v27
        } else {
            v4 = (v3 + v3)
        }
        _ = v26
    } else {
        v28 := GenS2{F0: -1, F1: 4294967295, F2: 0}
        _ = v28
    }
    switch (^(v1)) {
    case -519:
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 int64 = (v5 - -9223372036854775808)
            var v31 uint32 = (v2 >> ((uint(v2)) & 63))
            _ = fi29
            _ = v30
            _ = v31
        }
    case 824:
        var v32 int64 = 163
        _ = v32
    }
    v3 = (v3 + v3)
    for fi33 := 0; fi33 < 4; fi33++ {
        switch v1 {
        case 1:
            for fi34 := 0; fi34 < 16; fi34++ {
                v35 := []int{v1, fi33, (fi33 << ((255) & 63)), int(v6)}
                v36 := GenS3{F0: int8(v2)}
                v37 := GenS2{F0: -32768, F1: 455, F2: 18446744073709551615}
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
        case -9223372036854775808:
            if true {
                v5 += -89
                switch (v1 * 4294967295) {
                case -1:
                    v3 = ((((("foo bar" + "") + (v4 + v4)) + ("hello" + v3)) + v4) + v4)
                case 32768:
                    var v38 uint64 = 1
                    _ = v38
                case 2147483647:
                    fmt.Println(fi33, v1, v1)
                }
            }
        default:
            v1 += fi33
        }
        fmt.Println(v2, fi33, v1)
        _ = fi33
    }
    v6 = v6
    if ((v2 == (^(v2))) && true) {
        fmt.Println(v1)
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        for fi40 := 0; fi40 < 17; fi40++ {
            if (v5 >= (v5 * (v5 * (v5 % ((291) | 1))))) {
                v41 := []int{(v1 << ((18446744073709551615) & 63)), -511, 1}
                _ = v41
            }
            _ = fi40
        }
        var v42 int64 = (v5 & v5)
        v43 := GenS2{F0: 0, F1: 188, F2: 600}
        _ = fi39
        _ = v42
        _ = v43
    }
    var v44 int16 = 32767
    var v45 uint16 = 0
    v6 = v6
    v1 += 1
    v4 = v3
    v46 := GenS4{F0: 18446744073709551615, F1: v44, F2: v5}
    v47 := []int{}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
