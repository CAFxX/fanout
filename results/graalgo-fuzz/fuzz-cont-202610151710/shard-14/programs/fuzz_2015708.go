package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int32; F1 int; F2 int32 }
func genFn0(p0 uint64, p1 bool) uint {
    v5 := GenS1{F0: true, F1: 32767}
    var v6 bool = (-26 >= -77)
    if v6 {
        p1 = ((p0 ^ 592) > (p0 % (((p0 + (^(p0)))) | 1)))
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 uint32, p1 int8) int64 {
    p0 = p0
    if (-719 > 749) {
        p0 = 65535
        switch ("a" + (("\n\t" + "a") + ("foo bar" + ""))) {
        case "foo bar":
            var v7 int32 = 0
            _ = v7
        }
        for fi8 := 0; fi8 < 16; fi8++ {
            p1 = -108
            p0 += (p0 ^ 2147483648)
            v9 := []int{}
            _ = fi8
            _ = v9
        }
    } else {
        if false {
            v10 := []int{448}
            _ = v10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                p1 = p1
                _ = fi12
            }
            p1 -= (p1 & (p1 + (p1 % ((p1) | 1))))
            var v13 uint64 = 18446744073709551615
            _ = fi11
            _ = v13
        }
        fmt.Println(p0, p1)
    }
    p1 -= p1
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = 633
    var v3 uint8 = 168
    v4 := make(map[uint32]uint32)
    switch -212 {
    case -2147483648:
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 string = "a"
            _ = fi14
            _ = v15
        }
    case 42:
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 float64 = 1e10
            if true {
                var v18 int8 = -1
                v19 := GenS1{F0: (v17 <= v17), F1: 2}
                _ = v18
                _ = v19
            } else {
                fmt.Println(v1, v1, v2)
                switch "" {
                case "":
                    fmt.Println(v1, fi16)
                case "hello":
                    var v20 bool = false
                    _ = v20
                case "a":
                    var v21 uint32 = 52
                    _ = v21
                default:
                    v22 := []int{1, (fi16 * 42), -386}
                    _ = v22
                }
            }
            _ = fi16
            _ = v17
        }
    case -1:
        var v23 uint8 = (^(v3))
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        if false {
            v1 += 483
            var v25 int16 = 32767
            _ = v25
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            if false {
                v27 := GenS2{F0: 128, F1: (839 <= 0)}
                v1 = 4294967295
                switch "a" {
                case "hello":
                    if (0.0 != (3.14159 + ((float64(3.14159) * 1e10) - 0.0))) {
                        v2 -= v2
                    } else {
                        v28 := GenS3{F0: genFn1(2, 108), F1: 433, F2: 2147483648}
                        if false {
                            v29 := GenS4{F0: true}
                            v3 = v3
                            var v30 int64 = 9223372036854775807
                            _ = v29
                            _ = v30
                        }
                        _ = v28
                    }
                case "\n\t":
                    fmt.Println(fi26)
                case "a":
                    v2 = v2
                default:
                    v1 -= 2
                }
                if true {
                    if (42 == 255) {
                        v4[2147483647] = 2147483647
                        v31 := GenS3{F0: 0, F1: 2147483647, F2: 42}
                        _ = v31
                    } else {
                        switch (^(fi26)) {
                        case 574:
                            v2 = int32(fi26)
                        case 995:
                            v1 += v1
                        }
                        var v32 uint64 = (v1 % ((v1) | 1))
                        var v33 uint32 = 42
                        _ = v32
                        _ = v33
                    }
                    v2 = (-(v2))
                }
                _ = v27
            }
            _ = len(v4)
            v34 := v4[2147483647]
            _ = fi26
            _ = v34
        }
        fmt.Println(fi24, v1)
        _ = fi24
    }
    v4[2147483648] = 793
    if ((v1 * v1) != (v1 - 1)) {
        if (!true) {
            v3 += v3
            if true {
                for fi35 := 0; fi35 < 9; fi35++ {
                    for fi36 := 0; fi36 < 18; fi36++ {
                        v2 += v2
                        var v37 uint = genFn0((v1 >> ((uint(v1)) & 63)), true)
                        v37 = 0
                        _ = fi36
                        _ = v37
                    }
                    if false {
                        v4[2147483648] = 4294967295
                    } else {
                        v38 := GenS5{F0: "foo bar", F1: 0}
                        _ = v38
                    }
                    var v39 uint32 = 2147483648
                    _ = fi35
                    _ = v39
                }
                fmt.Println(v2, v1)
            }
            var v40 uint8 = (^(v3))
            _ = v40
        }
        v41 := GenS5{F0: ((("a" + "a") + ("hello" + "")) + "\n\t"), F1: 1}
        v3 += (v3 - v3)
        _ = v41
    } else {
        var v42 int64 = int64(v2)
        switch 42 {
        case 581:
            v42 = v42
        }
        _ = v42
    }
    v3 = v3
    if (true && ((!true) || true)) {
        for fi43 := 0; fi43 < 4; fi43++ {
            v1 -= (^(v1))
            delete(v4, 1)
            _ = fi43
        }
        fmt.Println(v3, v2, v1)
        for fi44 := 0; fi44 < 16; fi44++ {
            var v45 int = fi44
            v46 := GenS6{F0: v2, F1: -2147483648, F2: v2}
            var v47 string = (("\n\t" + (("" + ("hello" + "foo bar")) + ("hello" + ("foo bar" + "hello")))) + "a")
            _ = fi44
            _ = v45
            _ = v46
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        var v49 bool = (((("hello" + "a") + "hello") + "foo bar") == ("hello" + ""))
        _ = fi48
        _ = v49
    }
    v50 := []int{627, 42}
    v1 += (v1 >> ((uint(v3)) & 63))
    var v51 uint = 124
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
}
