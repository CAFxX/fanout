package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 int8 = 57
    if true {
        var v4 uint = 255
        v5 := GenS1{F0: true, F1: 885}
        _ = v4
        _ = v5
    }
    if (127 < v3) {
        var v6 uint16 = 42
        var v7 uint64 = 757
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS2{F0: "foo bar"}
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    switch "hello" {
    case "hello":
        for fi10 := 0; fi10 < 11; fi10++ {
            switch "" {
            case "hello":
                v11 := GenS1{F0: true, F1: 42}
                _ = v11
            case "":
                var v12 bool = false
                _ = v12
            }
            var v13 float64 = 0.0
            fmt.Println(fi10)
            _ = fi10
            _ = v13
        }
    case "a":
        var v14 uint = 488
        _ = v14
    }
    var v15 uint16 = 767
    fmt.Println(v3, v3)
    v3 += (v3 | 101)
    var v16 int16 = -32768
    if false {
        v17 := GenS2{F0: ""}
        _ = v17
    }
    switch 1 {
    case 127:
        switch 0 {
        case 32768:
            switch -2147483648 {
            case 2147483647:
                switch 4294967295 {
                case 799:
                    v18 := GenS1{F0: (v16 == v16), F1: 42}
                    _ = v18
                }
            case 4294967295:
                fmt.Println(v16, v15)
            case 249:
                for fi19 := 0; fi19 < 5; fi19++ {
                    for fi20 := 0; fi20 < 16; fi20++ {
                        _ = len(v2)
                        if true {
                            var v21 string = "a"
                            _ = v21
                        } else {
                            fmt.Println(fi20)
                            var v22 bool = (3.14159 <= (1e10 / (1e10 + 1e10)))
                            v1[fi20] = -9223372036854775808
                            var v23 int64 = 9223372036854775807
                            switch (-(fi20)) {
                            case 722:
                                v24 := []int{(fi19 >> ((679) & 63))}
                                _ = v24
                            case -48:
                                v23 = (^(v23))
                            }
                            _ = v22
                            _ = v23
                        }
                        v3 = (v3 + v3)
                        _ = fi20
                    }
                    _ = fi19
                }
            }
        default:
            switch 32768 {
            case 32767:
                switch "a" {
                case "a":
                    var v25 uint8 = 1
                    _ = v25
                default:
                    v26 := v1[127]
                    _ = v26
                }
            case -981:
                var v27 uint32 = 2147483648
                _ = v27
            }
        }
    case 1:
        v15 += (v15 * v15)
    case -955:
        v28 := GenS2{F0: ("\n\t" + "hello")}
        _ = v28
    default:
        var v29 bool = (false && false)
        _ = v29
    }
    if false {
        v16 -= -73
    } else {
        var v30 bool = true
        for fi31 := 0; fi31 < 2; fi31++ {
            fmt.Println(v3)
            switch "hello" {
            case "hello":
                if (!(!(74 == 42))) {
                    v32 := []int{fi31, fi31, fi31}
                    var v33 bool = (255 != 18446744073709551615)
                    _ = v32
                    _ = v33
                } else {
                    v30 = (-2.25 < ((1e10 - 0.0) * (float32((3.14159 / 0.1)) + (1e10 + 3.14159))))
                }
            }
            var v34 bool = true
            _ = fi31
            _ = v34
        }
        _ = v30
    }
    v35 := GenS2{F0: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v35
}
