package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint = 0
    v2 := []int{}
    var v3 bool = (((1e10 == 0.0) || (true || (9223372036854775807 >= -1))) || (float64(-2.25) <= float64(0.1)))
    var v4 int8 = -110
    if len(v2) > 0 {
        v2[0] = 86
    }
    for fi5 := 0; fi5 < 15; fi5++ {
        if v3 {
            var v6 uint64 = 885
            var v7 float32 = (0.1 + 3.14159)
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    v1 += v1
    fmt.Println(v1, v4, v4)
    v8 := GenS1{F0: 42}
    v9 := GenS1{F0: 0}
    v3 = ((("foo bar" + "foo bar") + "") == "foo bar")
    for fi10 := 0; fi10 < 1; fi10++ {
        switch ("foo bar" + "") {
        case "":
            var v11 uint8 = 37
            _ = v11
        case "foo bar":
            v12 := GenS1{F0: 42}
            _ = v12
        default:
            v4 += v4
        }
        var v13 uint16 = 32767
        _ = fi10
        _ = v13
    }
    v3 = v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 32767 {
    case -2147483648:
        v14 := GenS1{F0: 789}
        _ = v14
    case 157:
        v3 = false
    case 2147483647:
        if (2147483647 == 639) {
            if v3 {
                var v15 uint8 = 22
                if (v3 && ((^(v15)) == 127)) {
                    if len(v2) > 0 {
                        v2[0] = 32768
                    }
                } else {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v3 = true
                        v15 -= v15
                        _ = fi16
                    }
                    var v17 uint = (^(v1))
                    _ = v17
                }
                v15 += 42
                _ = v15
            } else {
                if true {
                    fmt.Println(v3, v4, v1)
                    v18 := GenS1{F0: 9223372036854775807}
                    _ = len(v2)
                    _ = v18
                } else {
                    if true {
                        if ((100 == (v4 ^ (v4 + (v4 / (((v4 / ((1) | 1))) | 1))))) && (2147483648 > 65535)) {
                            switch ("" + "hello") {
                            case "a":
                                v3 = v3
                            default:
                                if v3 {
                                    switch (("hello" + (("hello" + "\n\t") + ("a" + "hello"))) + "foo bar") {
                                    case "\n\t":
                                        if v3 {
                                            if (false && false) {
                                                var v19 uint8 = 2
                                                var v20 uint8 = v19
                                                var v21 int8 = v4
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                            } else {
                                                v4 += (v4 - 0)
                                            }
                                        }
                                    }
                                } else {
                                    if false {
                                        v4 += 0
                                        if (!(false && (1 == 42))) {
                                            switch ("a" + "") {
                                            case "a":
                                                if v3 {
                                                    for fi22 := 0; fi22 < 15; fi22++ {
                                                        var v23 float32 = 0.1
                                                        var v24 int64 = 0
                                                        _ = fi22
                                                        _ = v23
                                                        _ = v24
                                                    }
                                                }
                                            }
                                        } else {
                                            var v25 uint32 = 1
                                            _ = v25
                                        }
                                    }
                                    switch "" {
                                    case "a":
                                        var v26 float64 = 1e10
                                        _ = v26
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = 2147483647
                                        }
                                    default:
                                        v1 = 18446744073709551615
                                    }
                                    v4 = (v4 << ((18446744073709551615) & 63))
                                }
                            }
                        } else {
                            for fi27 := 0; fi27 < 4; fi27++ {
                                for fi28 := 0; fi28 < 20; fi28++ {
                                    var v29 float64 = -2.25
                                    v2 = append(v2, 758)
                                    if (-397 != fi27) {
                                        fmt.Println(v4, v4)
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                                _ = fi27
                            }
                            if (v3 && (!v3)) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v3 = v3
                            } else {
                                v30 := GenS1{F0: 0}
                                for fi31 := 0; fi31 < 11; fi31++ {
                                    v4 = v4
                                    fmt.Println(v4)
                                    _ = fi31
                                }
                                _ = v30
                            }
                            var v32 bool = v3
                            _ = v32
                        }
                    } else {
                        v33 := GenS1{F0: 443}
                        _ = v33
                    }
                    v4 = 127
                    switch "" {
                    case "foo bar":
                        v34 := GenS2{F0: -635}
                        _ = v34
                    case "a":
                        v2 = append(v2, -9223372036854775808)
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
            }
            v4 -= (^(v4))
            fmt.Println(v1, v4)
        }
    }
    fmt.Println(v1, v4)
    if true {
        var v35 int = 32767
        v36 := GenS2{F0: 32767}
        _ = v35
        _ = v36
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
