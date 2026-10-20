package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int) uint32 {
    var v5 string = ""
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: 0}
        v5 = "foo bar"
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return 897
}
func genFn1(p0 uint64, p1 uint64) int16 {
    p1 += 0
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint16 = 255
        v10 := GenS1{F0: p1}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    return -32768
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 0
    var v3 float64 = 3.14159
    var v4 int16 = 1000
    var v11 uint8 = 130
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            switch "hello" {
            case "a":
                v3 += v3
            case "hello":
                if true {
                    switch "" {
                    case "":
                        var v14 uint8 = 41
                        _ = v14
                    case "foo bar":
                        v3 = float64((((-2.25 + float32(v3)) + (3.14159 / (1e10 / float32(v3)))) + -2.25))
                    case "\n\t":
                        v3 += 1.5
                    }
                } else {
                    var v15 uint32 = genFn0(fi12)
                    _ = v15
                }
            }
            fmt.Println(fi12, v4)
            var v16 string = ""
            _ = fi13
            _ = v16
        }
        _ = fi12
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v3 += float64((0.1 / float32(-2.25)))
        _ = fi17
    }
    if (0 <= (v11 + (v11 + v11))) {
        v4 -= (v4 >> ((1) & 63))
    } else {
        if true {
            var v18 int64 = (-(v2))
            switch "" {
            case "a":
                var v19 uint16 = 2
                _ = v19
            case "hello":
                fmt.Println(v18)
            }
            v4 = v4
            _ = v18
        } else {
            switch -2147483648 {
            case 9223372036854775807:
                var v20 uint16 = 1
                _ = v20
            case 1:
                v21 := []int{-57, 4294967295, -9223372036854775808, 32768}
                _ = v21
            case 472:
                switch -1 {
                case -576:
                    var v22 uint = 1
                    _ = v22
                default:
                    v23 := []int{0, 478, -454, 4294967295}
                    _ = v23
                }
            }
            if false {
                v2 = 9223372036854775807
            } else {
                var v24 float64 = (v3 / (v3 + (0.1 + ((v3 / v3) - float64(1e10)))))
                if true {
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        v25 := []int{839, -687, 825}
                        _ = v25
                    case "\n\t":
                        for fi26 := 0; fi26 < 8; fi26++ {
                            fmt.Println(v4, v24)
                            if (v24 != 0.0) {
                                var v27 uint = 255
                                var v28 int16 = (v4 * -32768)
                                fmt.Println(v4, v28)
                                _ = v27
                                _ = v28
                            }
                            if true {
                                switch fi26 {
                                case -985:
                                    switch "foo bar" {
                                    case "hello":
                                        var v29 uint = 255
                                        _ = v29
                                    case "\n\t":
                                        v2 = (-(v2))
                                    }
                                case -627:
                                    var v30 uint16 = 42
                                    _ = v30
                                }
                                v11 = 104
                            }
                            _ = fi26
                        }
                    case "":
                        for fi31 := 0; fi31 < 8; fi31++ {
                            var v32 int = (fi31 >> ((255) & 63))
                            switch "hello" {
                            case "":
                                if false {
                                    v3 = v24
                                }
                            default:
                                for fi33 := 0; fi33 < 20; fi33++ {
                                    var v34 int8 = 100
                                    _ = fi33
                                    _ = v34
                                }
                            }
                            _ = fi31
                            _ = v32
                        }
                    default:
                        var v35 uint = 890
                        _ = v35
                    }
                    fmt.Println(v4, v2, v24)
                    v3 = 1.5
                } else {
                    var v36 float32 = float32(float64(1.5))
                    _ = len(v1)
                    _ = v36
                }
                _ = v24
            }
        }
    }
    v11 = v11
    for fi37 := 0; fi37 < 1; fi37++ {
        var v38 uint8 = (v11 << ((18446744073709551615) & 63))
        fmt.Println(v11, v3)
        v2 = (v2 + 42)
        _ = fi37
        _ = v38
    }
    switch ("a" + "") {
    case "":
        v3 += 3.14159
    case "a":
        v1[70] = 828
    case "\n\t":
        var v39 int8 = 1
        _ = v39
    }
    if false {
        v2 = 0
        var v40 int32 = 42
        _ = v40
    } else {
        switch 2147483647 {
        case 42:
            v11 += v11
        case 122:
            switch ("foo bar" + ("a" + "\n\t")) {
            case "\n\t":
                v41 := GenS1{F0: 0}
                _ = v41
            case "hello":
                v42 := []int{}
                _ = v42
            }
        }
        v2 = 9223372036854775807
    }
    var v43 uint16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v43
}
