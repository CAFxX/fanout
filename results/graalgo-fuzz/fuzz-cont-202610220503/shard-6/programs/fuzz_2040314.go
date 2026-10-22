package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint32; F2 string }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint }
func genFn0(p0 int16, p1 int64, p2 bool) int64 {
    if (p2 && p2) {
        var v7 uint8 = 2
        _ = v7
    }
    var v8 float64 = 0.0
    v9 := []int{}
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int16 = p0
            var v13 bool = false
            v14 := GenS1{F0: 18446744073709551615}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int64, p1 uint64, p2 int16) uint16 {
    v15 := []int{-2147483648, 127, -29}
    var v16 int8 = 100
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := GenS2{F0: (fi17 ^ -554), F1: 2147483648, F2: ""}
        var v19 bool = false
        fmt.Println(fi17, fi17)
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 635
}

func main() {
    var v1 int8 = 1
    var v2 float64 = 0.1
    var v3 int8 = (v1 + (v1 ^ v1))
    var v4 uint8 = 42
    v5 := make(map[string]string)
    var v6 int16 = 1
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 float64 = ((0.0 / v2) / (v2 + ((v2 / float64(1.5)) - float64(0.0))))
        _ = fi20
        _ = v21
    }
    var v22 uint32 = 65535
    var v23 uint64 = 2
    v24 := GenS2{F0: -724, F1: 697, F2: "a"}
    for fi25 := 0; fi25 < 17; fi25++ {
        switch "" {
        case "foo bar":
            v22 += (v22 * (v22 / ((v22) | 1)))
        case "":
            var v26 int = fi25
            _ = v26
        case "a":
            v1 -= (v3 << ((18446744073709551615) & 63))
        }
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 int = -467
            _ = fi27
            _ = v28
        }
        _ = fi25
    }
    if (genFn0(v6, 0, ((float64(1e10) >= (v2 - 3.14159)) || (true || false))) <= 1) {
        v5[""] = "\n\t"
        var v29 uint32 = (v22 / ((v22) | 1))
        fmt.Println(v1)
        _ = v29
    } else {
        v30 := GenS3{F0: -127, F1: -390}
        v22 = uint32(v23)
        for fi31 := 0; fi31 < 9; fi31++ {
            for fi32 := 0; fi32 < 9; fi32++ {
                v33 := []int{(fi31 >> ((65535) & 63)), fi31, (fi31 - 4294967295)}
                var v34 uint8 = (v4 >> ((4294967295) & 63))
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi31
        }
        _ = v30
    }
    var v35 uint8 = v4
    fmt.Println(v1, v35, v22)
    for fi36 := 0; fi36 < 20; fi36++ {
        v37 := []int{(fi36 % (((fi36 - fi36)) | 1))}
        _ = fi36
        _ = v37
    }
    if true {
        fmt.Println(v35)
        var v38 uint16 = 65535
        _ = v38
    }
    if true {
        switch 32767 {
        case -1:
            v3 += (v1 - v1)
        case 4294967295:
            switch "a" {
            case "hello":
                if false {
                    var v39 uint64 = (v23 & (v23 << ((928) & 63)))
                    v40 := GenS4{F0: genFn1(42, uint64(v35), 32767), F1: (v39 ^ v39)}
                    _ = v39
                    _ = v40
                }
            case "foo bar":
                switch ((("\n\t" + "foo bar") + "\n\t") + "a") {
                case "hello":
                    var v41 uint32 = v22
                    _ = v41
                default:
                    for fi42 := 0; fi42 < 18; fi42++ {
                        v43 := GenS5{F0: 9223372036854775807}
                        if (int64(v23) != int64(fi42)) {
                            var v44 string = ""
                            v4 -= 72
                            _ = v44
                        } else {
                            v45 := GenS6{F0: 23}
                            _ = v45
                        }
                        v46 := GenS7{F0: 65535}
                        _ = fi42
                        _ = v43
                        _ = v46
                    }
                }
            }
        default:
            for fi47 := 0; fi47 < 8; fi47++ {
                fmt.Println(v2)
                _ = fi47
            }
        }
        switch "hello" {
        case "a":
            fmt.Println(v2, v23)
        case "":
            var v48 int = 128
            _ = v48
        case "\n\t":
            for fi49 := 0; fi49 < 14; fi49++ {
                for fi50 := 0; fi50 < 11; fi50++ {
                    switch "a" {
                    case "\n\t":
                        v35 = uint8(v22)
                    case "foo bar":
                        if true {
                            switch "hello" {
                            case "a":
                                fmt.Println(v2, fi50, v4)
                            case "\n\t":
                                var v51 int = fi50
                                _ = v51
                            case "":
                                var v52 int32 = int32(v22)
                                _ = v52
                            }
                        }
                    default:
                        fmt.Println(v22)
                    }
                    fmt.Println(v4, v23)
                    v53 := GenS3{F0: v6, F1: 2147483647}
                    switch 32767 {
                    case 939:
                        switch fi50 {
                        case -2147483648:
                            var v54 int8 = v1
                            _ = v54
                        default:
                            v55 := GenS1{F0: 0}
                            _ = v55
                        }
                    case -2147483648:
                        v3 += int8(v22)
                    case 32767:
                        fmt.Println(v4, v35)
                    default:
                        var v56 uint64 = uint64(v22)
                        _ = v56
                    }
                    v4 = v4
                    _ = fi50
                    _ = v53
                }
                v3 = 127
                _ = fi49
            }
        }
        switch "\n\t" {
        case "a":
            v35 -= (^(v4))
        case "":
            var v57 float32 = 1e10
            _ = v57
        case "foo bar":
            v23 = v23
        default:
            v1 = v3
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
