package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int16 }
type GenS2 struct { F0 int8; F1 uint; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 int16; F2 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 uint; F2 bool }
func genFn0(p0 uint) uint8 {
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(fi5, fi5)
        _ = fi5
    }
    v6 := GenS1{F0: "", F1: 259, F2: 32767}
    v7 := GenS2{F0: -1, F1: 4294967295, F2: "foo bar"}
    var v8 uint8 = 92
    if true {
        var v9 string = "\n\t"
        _ = v9
    } else {
        v10 := []int{32767}
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS3{F0: 2147483647}
            _ = fi11
            _ = v12
        }
        v8 = v8
        _ = v10
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (v8 + 1)
}
func genFn1(p0 uint) int8 {
    v13 := GenS2{F0: -128, F1: 65535, F2: "a"}
    if true {
        for fi14 := 0; fi14 < 18; fi14++ {
            if true {
                var v15 uint32 = uint32(p0)
                v15 += 2147483647
                var v16 int8 = -39
                _ = v15
                _ = v16
            } else {
                if (!(42 == 417)) {
                    var v17 float64 = float64(((3.14159 + ((float32(-2.25) - 1.5) / 1.0)) - 0.1))
                    var v18 int16 = 1000
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(fi14, fi14)
                    if (!false) {
                        switch "\n\t" {
                        case "":
                            p0 = 623
                        case "\n\t":
                            for fi19 := 0; fi19 < 6; fi19++ {
                                v20 := GenS1{F0: "", F1: int16(fi19), F2: 1}
                                var v21 int16 = 1000
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                        default:
                            var v22 float32 = (1.5 * ((0.1 + (1e10 / 1.5)) + 0.0))
                            _ = v22
                        }
                        var v23 uint8 = 0
                        _ = v23
                    } else {
                        if (4294967295 == 0) {
                            fmt.Println(p0)
                            var v24 uint16 = 2
                            _ = v24
                        }
                        p0 = 4294967295
                        for fi25 := 0; fi25 < 10; fi25++ {
                            switch (("a" + "") + "\n\t") {
                            case "hello":
                                var v26 uint16 = 1
                                _ = v26
                            case "\n\t":
                                v27 := GenS4{F0: (false && true), F1: -1, F2: -9223372036854775808}
                                _ = v27
                            }
                            p0 += (p0 & p0)
                            _ = fi25
                        }
                    }
                    v28 := GenS2{F0: 42, F1: 255, F2: "hello"}
                    _ = v28
                }
                v29 := GenS2{F0: 47, F1: p0, F2: ("a" + "")}
                _ = v29
            }
            _ = fi14
        }
    }
    switch "hello" {
    case "":
        v30 := GenS1{F0: "a", F1: 0, F2: -725}
        _ = v30
    case "\n\t":
        p0 -= (^(p0))
    }
    _ = p0
    _ = v13
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 int64 = 1
    var v3 float32 = 0.1
    var v4 uint8 = 0
    for fi31 := 0; fi31 < 3; fi31++ {
        v32 := []int{int(v2), fi31}
        v1 += v1
        _ = fi31
        _ = v32
    }
    switch "a" {
    case "hello":
        if false {
            switch ("\n\t" + (("\n\t" + ("a" + ("\n\t" + "\n\t"))) + ("hello" + "foo bar"))) {
            case "a":
                for fi33 := 0; fi33 < 2; fi33++ {
                    v34 := []int{(-(fi33)), fi33, 2147483647, fi33}
                    var v35 uint64 = 4294967295
                    fmt.Println(v3)
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            case "foo bar":
                var v36 int64 = v2
                _ = v36
            }
        } else {
            v37 := []int{}
            _ = v37
        }
    case "\n\t":
        var v38 uint16 = 32767
        _ = v38
    case "a":
        v39 := []int{}
        _ = v39
    default:
        v40 := GenS4{F0: (!(!(false && (!(173 <= 594))))), F1: 487, F2: v2}
        _ = v40
    }
    if (!(1000 != -1)) {
        if ((v2 ^ v2) == (v2 >> ((4294967295) & 63))) {
            v1 = 1
        } else {
            fmt.Println(v4, v1, v2)
            var v41 int = -560
            _ = v41
        }
        fmt.Println(v2, v2, v2)
    }
    v42 := GenS5{F0: true}
    v1 = v1
    for fi43 := 0; fi43 < 4; fi43++ {
        var v44 int8 = 1
        _ = fi43
        _ = v44
    }
    var v45 int32 = 596
    for fi46 := 0; fi46 < 11; fi46++ {
        var v47 string = "a"
        _ = fi46
        _ = v47
    }
    for fi48 := 0; fi48 < 13; fi48++ {
        var v49 float64 = ((float64(0.1) + 1e10) - 0.1)
        var v50 int16 = 0
        _ = fi48
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 15; fi51++ {
        switch (fi51 - -713) {
        case -1:
            var v52 uint = uint(v4)
            _ = v52
        }
        _ = fi51
    }
    v2 -= (v2 >> ((4294967295) & 63))
    for fi53 := 0; fi53 < 4; fi53++ {
        for fi54 := 0; fi54 < 13; fi54++ {
            v55 := GenS6{F0: 201, F1: 18446744073709551615, F2: true}
            v56 := []int{(fi54 / ((fi53) | 1)), fi54, fi53}
            if true {
                switch (fi54 / ((fi53) | 1)) {
                case 128:
                    if false {
                        switch (fi54 + -199) {
                        case 9223372036854775807:
                            var v57 string = ("foo bar" + "a")
                            _ = v57
                        case 2147483647:
                            fmt.Println(v1, v2)
                        }
                    }
                case 4294967295:
                    v1 -= v1
                }
            }
            fmt.Println(v2)
            _ = fi54
            _ = v55
            _ = v56
        }
        _ = fi53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
}
