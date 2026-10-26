package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int32 = 42
    var v2 uint64 = 1
    var v3 float64 = -2.25
    var v4 int32 = (v1 % ((v1) | 1))
    if (-1 >= -1) {
        var v5 uint32 = 1
        if (1 < 32768) {
            if (95 > 908) {
                v4 += -2147483648
            } else {
                v2 += (^(v2))
            }
        }
        _ = v5
    } else {
        v6 := GenS1{F0: 1}
        for fi7 := 0; fi7 < 4; fi7++ {
            fmt.Println(fi7, v2, v1)
            _ = fi7
        }
        _ = v6
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        switch "hello" {
        case "":
            for fi9 := 0; fi9 < 3; fi9++ {
                if false {
                    switch "\n\t" {
                    case "a":
                        switch fi8 {
                        case 2147483647:
                            v10 := GenS1{F0: 42}
                            _ = v10
                        case -9223372036854775808:
                            v11 := []int{int(v4), fi9, (-(fi8)), 127}
                            _ = v11
                        }
                    case "":
                        v12 := []int{806, (fi9 ^ (fi9 - 2147483647)), -825}
                        _ = v12
                    case "foo bar":
                        fmt.Println(v3)
                    default:
                        v3 = ((((v3 - v3) + (v3 - v3)) / (v3 + v3)) * v3)
                    }
                } else {
                    v13 := GenS1{F0: 255}
                    _ = v13
                }
                v1 += -1
                v4 += v4
                _ = fi9
            }
        default:
            for fi14 := 0; fi14 < 1; fi14++ {
                v15 := []int{}
                if true {
                    var v16 int8 = -82
                    if (false && (32767 <= int16(v16))) {
                        var v17 uint8 = 2
                        _ = v17
                    } else {
                        var v18 uint8 = 255
                        fmt.Println(v18, v1)
                        var v19 uint64 = (v2 * v2)
                        _ = v18
                        _ = v19
                    }
                    var v20 int16 = 32767
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = v16
                    _ = v20
                }
                _ = fi14
                _ = v15
            }
        }
        _ = fi8
    }
    var v21 uint = 1
    for fi22 := 0; fi22 < 7; fi22++ {
        fmt.Println(v1, v21, v2)
        v4 += v1
        v23 := []int{(fi22 * (fi22 + (fi22 ^ fi22))), (fi22 | (fi22 + (fi22 >> ((v21) & 63)))), (fi22 % ((fi22) | 1)), 128}
        _ = fi22
        _ = v23
    }
    var v24 string = ""
    fmt.Println(v4, v21, v24)
    v4 -= (v1 & v4)
    v25 := GenS1{F0: 719}
    switch ("foo bar" + ((v24 + v24) + "hello")) {
    case "foo bar":
        if false {
            fmt.Println(v4)
            for fi26 := 0; fi26 < 18; fi26++ {
                v27 := GenS1{F0: uint16(v2)}
                var v28 int8 = 1
                var v29 int = fi26
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                for fi31 := 0; fi31 < 4; fi31++ {
                    v4 = (^(v1))
                    switch (^(fi30)) {
                    case 9223372036854775807:
                        fmt.Println(fi30, v2)
                    }
                    v32 := []int{}
                    if false {
                        v33 := GenS1{F0: 619}
                        _ = v33
                    }
                    v1 += v4
                    _ = fi31
                    _ = v32
                }
                if (false || true) {
                    switch (fi30 | fi30) {
                    case 4294967295:
                        fmt.Println(v24)
                    }
                    var v34 bool = true
                    _ = v34
                } else {
                    v35 := GenS1{F0: 65535}
                    _ = v35
                }
                _ = fi30
            }
        }
    case "a":
        var v36 int32 = (v4 ^ (-(v4)))
        _ = v36
    case "\n\t":
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 float32 = (1.5 * 3.14159)
            v39 := GenS1{F0: 2}
            var v40 int32 = v4
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
