package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint }
func genFn0() int8 {
    if false {
        var v7 int32 = -215
        v7 = v7
        _ = v7
    } else {
        if (1 > -15) {
            var v8 uint16 = 255
            _ = v8
        } else {
            var v9 int32 = 2147483647
            fmt.Println(v9)
            _ = v9
        }
        var v10 int8 = 42
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            if true {
                var v13 int64 = 42
                _ = v13
            } else {
                var v14 int8 = 42
                var v15 int8 = 42
                _ = v14
                _ = v15
            }
            var v16 int = (fi11 % ((fi11) | 1))
            _ = fi12
            _ = v16
        }
        var v17 string = "foo bar"
        _ = fi11
        _ = v17
    }
    var v18 bool = true
    if (3.14159 > 0.0) {
        if (1 > 0) {
            v18 = true
        } else {
            if (false || ("hello" == "a")) {
                var v19 uint64 = 649
                if v18 {
                    v19 += 2
                    for fi20 := 0; fi20 < 4; fi20++ {
                        fmt.Println(v18, v18, v19)
                        v18 = (((255 <= 26) && v18) || v18)
                        if (false || v18) {
                            switch fi20 {
                            case 2147483647:
                                var v21 float32 = (0.1 + float32(1.5))
                                _ = v21
                            default:
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    var v23 uint = 87
                                    v24 := []int{(fi20 ^ (fi22 - 224)), 128}
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                            }
                        } else {
                            var v25 int32 = 0
                            var v26 int8 = -86
                            _ = v25
                            _ = v26
                        }
                        _ = fi20
                    }
                }
                switch 128 {
                case 2147483647:
                    for fi27 := 0; fi27 < 5; fi27++ {
                        v18 = ((fi27 >> ((4294967295) & 63)) < (fi27 - fi27))
                        for fi28 := 0; fi28 < 6; fi28++ {
                            v18 = (42 != 20)
                            if (!((127 > 100) && false)) {
                                switch -594 {
                                case -250:
                                    var v29 float32 = float32(0.1)
                                    _ = v29
                                case 128:
                                    var v30 int16 = 42
                                    _ = v30
                                case 0:
                                    v31 := []int{252, 71}
                                    _ = v31
                                }
                            }
                            _ = fi28
                        }
                        _ = fi27
                    }
                }
                _ = v19
            } else {
                v32 := []int{-9223372036854775808}
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v18)
                case "a":
                    for fi33 := 0; fi33 < 10; fi33++ {
                        v34 := GenS1{F0: -9223372036854775808, F1: 498}
                        _ = fi33
                        _ = v34
                    }
                case "hello":
                    _ = len(v32)
                default:
                    switch "hello" {
                    case "":
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v18)
                        }
                    case "foo bar":
                        v35 := GenS1{F0: 1, F1: 1}
                        _ = v35
                    case "\n\t":
                        _ = len(v32)
                    default:
                        var v36 int64 = 42
                        _ = v36
                    }
                }
                for fi37 := 0; fi37 < 10; fi37++ {
                    v18 = v18
                    v18 = (724 > 1)
                    _ = fi37
                }
                _ = v32
            }
            var v38 int16 = -32768
            _ = v38
        }
    } else {
        var v39 uint64 = 566
        v40 := GenS1{F0: 9223372036854775807, F1: 2147483647}
        _ = v39
        _ = v40
    }
    v41 := GenS2{F0: 926}
    _ = v18
    _ = v41
    return 0
}
func genFn1() uint64 {
    v42 := []int{-2147483648, -2147483648, 622}
    for fi43 := 0; fi43 < 11; fi43++ {
        var v44 int32 = 0
        v44 = (v44 + (v44 / ((2147483647) | 1)))
        fmt.Println(fi43, v44)
        _ = fi43
        _ = v44
    }
    _ = v42
    return 1
}

func main() {
    var v1 float32 = (-2.25 + ((3.14159 * -2.25) / 3.14159))
    var v2 int32 = -2147483648
    var v3 uint16 = 2
    var v4 uint64 = 998
    var v5 int = -9223372036854775808
    var v6 string = "\n\t"
    fmt.Println(v2, v3, v5)
    var v45 int8 = -1
    var v46 int8 = (v45 + (^(v45)))
    v47 := []int{v5, -9223372036854775808, (v5 / ((v5) | 1))}
    v48 := GenS1{F0: -1, F1: 1}
    fmt.Println(v5)
    var v49 uint8 = 255
    v50 := GenS1{F0: -984, F1: 42}
    if len(v47) >= 2 {
        v47 = v47[1:len(v47)]
    }
    v2 += (-(v2))
    v51 := GenS1{F0: -1, F1: 641}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
