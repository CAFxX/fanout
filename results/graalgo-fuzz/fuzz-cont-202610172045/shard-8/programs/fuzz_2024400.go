package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int; F1 int8; F2 string }
func genFn0() bool {
    v5 := GenS1{F0: ("foo bar" + ("" + (("a" + "hello") + ""))), F1: -32768, F2: 264}
    switch ("\n\t" + ("\n\t" + ("a" + ""))) {
    case "\n\t":
        v6 := []int{2147483647}
        _ = v6
    case "a":
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 float32 = -2.25
            _ = fi7
            _ = v8
        }
    case "":
        for fi9 := 0; fi9 < 5; fi9++ {
            if false {
                v10 := GenS1{F0: (("" + ("hello" + "a")) + "hello"), F1: -1, F2: -2147483648}
                v11 := []int{}
                _ = v10
                _ = v11
            } else {
                switch fi9 {
                case 32768:
                    v12 := []int{fi9, 4294967295}
                    _ = v12
                case 128:
                    v13 := []int{}
                    _ = v13
                }
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 int32 = 454
                    fmt.Println(v15, fi9, v15)
                    fmt.Println(fi9, fi9, fi14)
                    _ = fi14
                    _ = v15
                }
                if (255 <= 498) {
                    if true {
                        switch "hello" {
                        case "\n\t":
                            var v16 float64 = (3.14159 - 1e10)
                            _ = v16
                        default:
                            var v17 float32 = (3.14159 / float32(-2.25))
                            _ = v17
                        }
                    } else {
                        if true {
                            var v18 uint = 265
                            v19 := GenS1{F0: "", F1: 42, F2: 107}
                            _ = v18
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 20; fi20++ {
                            if (2147483647 > -2147483648) {
                                var v21 int8 = 96
                                v21 = 85
                                v22 := []int{474}
                                _ = v21
                                _ = v22
                            } else {
                                var v23 string = ("foo bar" + "a")
                                v24 := GenS2{F0: 2147483647, F1: false}
                                if false {
                                    var v25 uint8 = 42
                                    v26 := []int{(fi20 | (fi9 % (((fi9 % ((fi9) | 1))) | 1))), fi9, (fi20 / ((9223372036854775807) | 1)), -901}
                                    v27 := GenS2{F0: 133, F1: (-1 != 1000)}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                _ = v23
                                _ = v24
                            }
                            var v28 uint16 = 2
                            _ = fi20
                            _ = v28
                        }
                        fmt.Println(fi9)
                    }
                    switch ("\n\t" + "") {
                    case "":
                        var v29 uint16 = 42
                        _ = v29
                    }
                    var v30 bool = false
                    _ = v30
                }
            }
            var v31 int32 = 959
            _ = fi9
            _ = v31
        }
    default:
        var v32 bool = true
        _ = v32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 int32 = 1
        var v35 uint64 = 18446744073709551615
        _ = fi33
        _ = v34
        _ = v35
    }
    _ = v5
    return (true || false)
}

func main() {
    var v1 float32 = (0.0 / -2.25)
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    var v4 int8 = 127
    v4 -= v4
    v1 = (((1e10 + v1) + ((v1 - (v1 - 3.14159)) * v1)) * 1.5)
    if true {
        fmt.Println(v1)
    } else {
        for fi36 := 0; fi36 < 19; fi36++ {
            switch -756 {
            case 4294967295:
                for fi37 := 0; fi37 < 14; fi37++ {
                    v38 := GenS2{F0: 35, F1: (false || false)}
                    fmt.Println(v4, v4)
                    fmt.Println(v4, fi36, fi37)
                    _ = fi37
                    _ = v38
                }
            case -9223372036854775808:
                v1 -= float32(float64(v1))
            case -2147483648:
                fmt.Println(v4, v1, v1)
            }
            _ = fi36
        }
    }
    var v39 int64 = 42
    var v40 int8 = (v4 - 62)
    v40 -= v40
    v41 := GenS3{F0: 250, F1: (v40 * v4), F2: "hello"}
    var v42 float32 = 1.5
    var v43 float64 = (0.0 + 0.1)
    v44 := GenS2{F0: 2147483648, F1: ((v39 - (-(v39))) >= (v39 | -135))}
    for fi45 := 0; fi45 < 1; fi45++ {
        v39 -= 9223372036854775807
        switch 32768 {
        case 127:
            v4 += 42
        default:
            for fi46 := 0; fi46 < 19; fi46++ {
                fmt.Println(v1)
                v47 := GenS1{F0: "\n\t", F1: 858, F2: -666}
                var v48 uint = 488
                _ = fi46
                _ = v47
                _ = v48
            }
        }
        v49 := GenS1{F0: "foo bar", F1: -147, F2: 2147483647}
        _ = fi45
        _ = v49
    }
    var v50 bool = (42 < 127)
    v51 := GenS1{F0: "a", F1: 1000, F2: 2147483647}
    fmt.Println(v1, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v50
    _ = v51
}
