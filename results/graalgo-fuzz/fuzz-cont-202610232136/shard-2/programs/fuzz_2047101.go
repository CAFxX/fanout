package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 int; F2 uint }
func genFn0() bool {
    switch 32768 {
    case 4294967295:
        var v5 float32 = -2.25
        _ = v5
    case 9223372036854775807:
        var v6 int8 = 42
        _ = v6
    case 32767:
        var v7 int8 = 49
        _ = v7
    default:
        switch "a" {
        case "\n\t":
            var v8 int32 = 2147483647
            _ = v8
        }
    }
    v9 := []int{0, 687, -2147483648, 42}
    v9 = append(v9, 127)
    _ = v9
    return true
}
func genFn1(p0 uint32, p1 int) uint64 {
    for fi10 := 0; fi10 < 14; fi10++ {
        fmt.Println(fi10)
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS1{F0: 1000, F1: -658}
            var v13 uint = 18446744073709551615
            p1 -= (fi11 * (fi11 - fi11))
            _ = fi11
            _ = v12
            _ = v13
        }
        p1 += fi10
        _ = fi10
    }
    p1 -= 42
    v14 := []int{0, p1, p1, p1}
    _ = p0
    _ = p1
    _ = v14
    return 875
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 int8 = 100
    v4 := make(map[uint32]string)
    fmt.Println(v3)
    switch (((("" + ("hello" + "\n\t")) + (("hello" + "hello") + ("\n\t" + ""))) + (("foo bar" + "hello") + "a")) + ("foo bar" + "")) {
    case "":
        var v15 int64 = 9223372036854775807
        _ = v15
    case "\n\t":
        v16 := v4[2147483647]
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 16; fi17++ {
            v3 = (v3 / ((v3) | 1))
            if false {
                for fi18 := 0; fi18 < 8; fi18++ {
                    if true {
                        if true {
                            v3 -= 42
                            v19 := GenS1{F0: 243, F1: -9223372036854775808}
                            _ = v19
                        } else {
                            if (!(!true)) {
                                v3 = v3
                                switch ("hello" + ("" + "\n\t")) {
                                case "a":
                                    v3 = (-(v3))
                                case "hello":
                                    v3 -= v3
                                default:
                                    if true {
                                        v20 := []int{fi17}
                                        fmt.Println(v3, v3, fi18)
                                        _ = v20
                                    }
                                }
                            }
                            var v21 uint64 = 2
                            _ = v21
                        }
                    }
                    v3 = 0
                    switch "a" {
                    case "foo bar":
                        if ((float32(((((-2.25 * 0.0) + (1e10 * 3.14159)) * ((0.1 / (0.1 - 3.14159)) * 0.0)) - ((float64((0.1 / 1.0)) - 1.5) * (0.0 + 1.5)))) / 0.1) != -2.25) {
                            if (553 == genFn1(1, fi18)) {
                                v22 := []int{(-(fi17)), (fi17 * fi17), (fi17 - (fi18 + fi18)), fi18}
                                _ = v22
                            } else {
                                if (!true) {
                                    v3 += 59
                                    switch ("\n\t" + "hello") {
                                    case "foo bar":
                                        switch 1 {
                                        case 42:
                                            v23 := GenS2{F0: 127}
                                            _ = v23
                                        }
                                    default:
                                        v3 = (-(v3))
                                    }
                                }
                            }
                            v24 := GenS3{F0: 4294967295}
                            _ = v24
                        } else {
                            v3 = -1
                            switch ("\n\t" + ("a" + "\n\t")) {
                            case "hello":
                                switch "\n\t" {
                                case "":
                                    if (((1e10 + 3.14159) * 3.14159) >= 0.1) {
                                        switch "a" {
                                        case "a":
                                            v3 -= v3
                                        case "":
                                            v3 += v3
                                        case "hello":
                                            v25 := GenS1{F0: 1000, F1: 9223372036854775807}
                                            _ = v25
                                        }
                                    } else {
                                        v3 = (v3 | v3)
                                        if (21 >= v3) {
                                            fmt.Println(fi17, fi17, fi17)
                                            v3 = v3
                                        }
                                    }
                                default:
                                    var v26 bool = true
                                    _ = v26
                                }
                            case "":
                                switch "hello" {
                                case "foo bar":
                                    switch "a" {
                                    case "\n\t":
                                        if ((fi18 | (-(fi17))) == fi17) {
                                            fmt.Println(v3, v3, fi17)
                                            v3 = (v3 % ((v3) | 1))
                                        } else {
                                            var v27 uint64 = genFn1(604, (fi18 - fi17))
                                            _ = v27
                                        }
                                    case "":
                                        v28 := GenS4{F0: -165}
                                        _ = v28
                                    case "foo bar":
                                        v29 := GenS2{F0: 128}
                                        _ = v29
                                    default:
                                        delete(v4, 0)
                                    }
                                case "\n\t":
                                    if (true || false) {
                                        var v30 int16 = 0
                                        _ = v30
                                    } else {
                                        var v31 bool = true
                                        v32 := []int{(fi18 ^ 298), (fi17 << ((2) & 63)), 209}
                                        _ = len(v32)
                                        _ = len(v4)
                                        _ = v31
                                        _ = v32
                                    }
                                case "":
                                    v3 -= 37
                                }
                            }
                        }
                    case "hello":
                        v3 = (v3 * 127)
                    default:
                        v3 -= (v3 >> ((460) & 63))
                    }
                    _ = fi18
                }
            } else {
                var v33 uint = 229
                var v34 int = 127
                _ = v33
                _ = v34
            }
            var v35 string = ("foo bar" + "a")
            _ = fi17
            _ = v35
        }
        switch 127 {
        case 4294967295:
            switch "hello" {
            case "hello":
                v3 = v3
            case "":
                var v36 string = "\n\t"
                _ = v36
            case "\n\t":
                var v37 uint16 = 42
                _ = v37
            }
        case 9223372036854775807:
            v3 = v3
        case 32768:
            fmt.Println(v3)
        default:
            v38 := []int{}
            _ = v38
        }
        var v39 uint64 = genFn1(976, 1)
        _ = v39
    }
    v3 += (v3 % ((v3) | 1))
    v40 := []int{127, 4294967295}
    v3 = v3
    if true {
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
        if false {
            v41 := GenS5{F0: 0, F1: -1, F2: 65535}
            v3 += (v3 % ((v3) | 1))
            if true {
                for fi42 := 0; fi42 < 19; fi42++ {
                    if false {
                        for fi43 := 0; fi43 < 10; fi43++ {
                            var v44 int32 = -679
                            _ = fi43
                            _ = v44
                        }
                    } else {
                        var v45 int64 = 0
                        _ = v45
                    }
                    var v46 float64 = 0.0
                    _ = fi42
                    _ = v46
                }
            }
            _ = v41
        }
        v3 += v3
    }
    var v47 float32 = 3.14159
    fmt.Println(v3, v3)
    for fi48 := 0; fi48 < 18; fi48++ {
        v3 += (v3 << ((105) & 63))
        _ = fi48
    }
    v3 = (v3 - (v3 | v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
}
