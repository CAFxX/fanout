package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    var v1 string = "hello"
    var v2 string = v1
    var v3 int8 = -26
    var v4 float64 = 1e10
    var v5 bool = (1 > 128)
    v1 = "\n\t"
    v6 := GenS1{F0: 255, F1: 528}
    if true {
        fmt.Println(v1, v2, v3)
        switch (v2 + ("" + ("" + v1))) {
        case "foo bar":
            v7 := GenS1{F0: 205, F1: 42}
            _ = v7
        }
        var v8 bool = (v4 != float64(0.1))
        _ = v8
    }
    var v9 float64 = 1.5
    var v10 uint = 18446744073709551615
    if true {
        for fi11 := 0; fi11 < 14; fi11++ {
            if (v5 || (2147483647 >= 722)) {
                fmt.Println(fi11)
                v1 = (("a" + "hello") + ((v2 + (v1 + v2)) + "foo bar"))
            }
            if (((v10 ^ v10) > v10) && v5) {
                v9 -= -2.25
            } else {
                for fi12 := 0; fi12 < 9; fi12++ {
                    if (32767 > 141) {
                        if (v5 || v5) {
                            v10 += (v10 + 18446744073709551615)
                            var v13 float32 = 1e10
                            switch ((v1 + v1) + "") {
                            case "foo bar":
                                switch v1 {
                                case "hello":
                                    var v14 int64 = 1
                                    _ = v14
                                case "a":
                                    v15 := GenS1{F0: 127, F1: 1}
                                    _ = v15
                                case "":
                                    if true {
                                        fmt.Println(v13, v5, v13)
                                        if (!((!(false || v5)) || (2147483648 < 507))) {
                                            var v16 string = (v1 + (v2 + v2))
                                            if v5 {
                                                fmt.Println(fi12, v2, v3)
                                                v17 := []int{(fi12 + 31), fi11}
                                                if false {
                                                    var v18 bool = ((42 != 2147483647) || ((!(!(v9 >= 0.1))) || ((2 < v10) || v5)))
                                                    v19 := GenS1{F0: 0, F1: 0}
                                                    _ = v18
                                                    _ = v19
                                                } else {
                                                    var v20 uint64 = 18446744073709551615
                                                    v2 = v1
                                                    v13 -= v13
                                                    _ = v20
                                                }
                                                _ = v17
                                            }
                                            _ = v16
                                        } else {
                                            fmt.Println(v4, v2, v10)
                                            v5 = (-262 != 42)
                                        }
                                        v10 += (v10 << ((v10) & 63))
                                    } else {
                                        var v21 int = fi12
                                        var v22 uint8 = 255
                                        _ = v21
                                        _ = v22
                                    }
                                default:
                                    v23 := GenS1{F0: 42, F1: 32767}
                                    _ = v23
                                }
                            case "":
                                v13 += (v13 - (v13 - v13))
                            default:
                                var v24 int8 = v3
                                _ = v24
                            }
                            v25 := []int{fi12, 796, (fi12 | int(v3))}
                            _ = v13
                            _ = v25
                        }
                        v26 := []int{}
                        v2 = ""
                        _ = v26
                    }
                    _ = fi12
                }
            }
            _ = fi11
        }
        v5 = v5
        switch -502 {
        case 32767:
            var v27 uint16 = 255
            _ = v27
        case 0:
            v28 := []int{}
            _ = v28
        case 127:
            fmt.Println(v10, v3)
        default:
            if (true || (v10 != v10)) {
                v29 := []int{32768, -420, 1, 757}
                v30 := GenS2{F0: uint8(v10), F1: ((!(((v10 << (((v10 | 881)) & 63)) < (v10 ^ v10)) && (v5 && v5))) || v5)}
                _ = v29
                _ = v30
            } else {
                fmt.Println(v2)
                v9 = 0.1
            }
        }
    } else {
        v31 := []int{}
        _ = v31
    }
    switch 9223372036854775807 {
    case 426:
        v32 := []int{int(v10), 127, 42}
        _ = v32
    }
    switch v2 {
    case "hello":
        if (-2147483648 <= 0) {
            if (-1 > -918) {
                for fi33 := 0; fi33 < 16; fi33++ {
                    v34 := []int{}
                    _ = fi33
                    _ = v34
                }
            } else {
                v35 := []int{904, -854, 2147483647, 402}
                fmt.Println(v1, v9, v3)
                _ = v35
            }
            switch -2147483648 {
            case 32767:
                v1 = (("" + (v2 + ("foo bar" + "foo bar"))) + v1)
            case 4294967295:
                v5 = (v5 || (("hello" + v2) == v1))
            }
        } else {
            for fi36 := 0; fi36 < 4; fi36++ {
                if true {
                    for fi37 := 0; fi37 < 7; fi37++ {
                        var v38 uint = 2
                        _ = fi37
                        _ = v38
                    }
                } else {
                    for fi39 := 0; fi39 < 16; fi39++ {
                        var v40 int16 = -32768
                        v41 := GenS2{F0: 128, F1: ((v10 < (v10 % ((647) | 1))) && true)}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    if false {
                        v9 = v9
                    }
                    v9 -= v9
                }
                if v5 {
                    v42 := []int{1}
                    _ = v42
                } else {
                    fmt.Println(v3, v10)
                }
                v5 = v5
                _ = fi36
            }
        }
    default:
        if v5 {
            fmt.Println(v2, v5)
        }
    }
    v10 = v10
    v3 += (v3 % ((v3) | 1))
    if v5 {
        v1 = "foo bar"
        fmt.Println(v1, v9, v1)
    }
    fmt.Println(v9, v2)
    if v5 {
        v1 = (v2 + ((v1 + "hello") + v2))
        var v43 string = "hello"
        v1 = "\n\t"
        _ = v43
    } else {
        v5 = (false || ((v9 / 1.0) > v9))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
