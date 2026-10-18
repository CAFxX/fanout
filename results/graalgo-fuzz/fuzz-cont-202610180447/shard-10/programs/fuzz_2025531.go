package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
func genFn0() int16 {
    v4 := GenS1{F0: 128}
    switch 9223372036854775807 {
    case 32768:
        if (1 >= 532) {
            var v5 float64 = 3.14159
            v6 := GenS2{F0: 472}
            if false {
                var v7 string = "\n\t"
                var v8 int8 = -128
                _ = v7
                _ = v8
            } else {
                var v9 bool = (-1 > 1)
                v5 = 0.1
                for fi10 := 0; fi10 < 12; fi10++ {
                    v9 = v9
                    _ = fi10
                }
                _ = v9
            }
            _ = v5
            _ = v6
        } else {
            for fi11 := 0; fi11 < 1; fi11++ {
                if true {
                    v12 := []int{-1, 862, fi11, (fi11 & -2147483648)}
                    var v13 uint64 = 4294967295
                    _ = v12
                    _ = v13
                } else {
                    switch fi11 {
                    case 378:
                        v14 := GenS3{F0: -120}
                        _ = v14
                    default:
                        var v15 uint8 = 0
                        _ = v15
                    }
                }
                var v16 int = fi11
                _ = fi11
                _ = v16
            }
            switch 222 {
            case -630:
                var v17 bool = true
                _ = v17
            case 32768:
                v18 := GenS3{F0: 1}
                _ = v18
            default:
                var v19 uint8 = 2
                _ = v19
            }
        }
    case 89:
        v20 := []int{4294967295, -956}
        _ = v20
    }
    _ = v4
    return -583
}

func main() {
    var v1 int32 = -1
    var v2 int16 = -600
    var v3 uint32 = 2147483648
    v2 -= 1000
    var v21 int = 9223372036854775807
    var v22 uint16 = 255
    for fi23 := 0; fi23 < 11; fi23++ {
        var v24 uint16 = v22
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 float64 = 3.14159
        var v27 int16 = v2
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v2, v21, v21)
    for fi28 := 0; fi28 < 20; fi28++ {
        if ((9223372036854775807 != 630) || (!false)) {
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 uint8 = 255
                switch "\n\t" {
                case "hello":
                    if (("foo bar" + "hello") >= ("foo bar" + "\n\t")) {
                        fmt.Println(v21)
                        fmt.Println(v2, v21, fi29)
                        v22 += (v22 ^ 42)
                    }
                case "":
                    v21 = 21
                case "foo bar":
                    if true {
                        if true {
                            if (!(!(4294967295 == 42))) {
                                v2 -= -32768
                                switch "foo bar" {
                                case "\n\t":
                                    v2 = genFn0()
                                default:
                                    if (!(!(!false))) {
                                        v31 := []int{-1}
                                        if (3.14159 >= 1e10) {
                                            var v32 string = ""
                                            if len(v31) >= 2 {
                                                v31 = v31[1:len(v31)]
                                            }
                                            v32 = (((v32 + (("a" + v32) + "a")) + (("\n\t" + v32) + "a")) + v32)
                                            _ = len(v31)
                                            v33 := GenS3{F0: 0}
                                            v1 -= (v1 - (v1 + -597))
                                            v3 = v3
                                            _ = v32
                                            _ = v33
                                        }
                                        var v34 uint64 = uint64(v22)
                                        _ = v31
                                        _ = v34
                                    }
                                }
                                fmt.Println(fi29, v1)
                            } else {
                                var v35 uint64 = 18446744073709551615
                                var v36 uint16 = v22
                                v37 := GenS2{F0: -9223372036854775808}
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                        }
                        v1 -= (^(v1))
                        var v38 float64 = float64((3.14159 + ((float32((1.5 - 1.5)) - (0.0 - -2.25)) / 1.5)))
                        _ = v38
                    } else {
                        v22 = 116
                        v21 -= 32768
                    }
                }
                v1 += v1
                _ = fi29
                _ = v30
            }
        }
        v3 = 2
        v39 := []int{-2147483648, (fi28 + (-(fi28))), int(v3), 0}
        _ = fi28
        _ = v39
    }
    switch ("hello" + "a") {
    case "hello":
        v1 -= 2147483647
    case "":
        v1 -= (v1 >> ((uint(v22)) & 63))
    case "foo bar":
        v3 += (v3 >> ((18446744073709551615) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
