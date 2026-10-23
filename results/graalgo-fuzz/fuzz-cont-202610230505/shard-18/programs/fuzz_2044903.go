package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []string{"foo bar"}
    v2 := []string{""}
    v3 := make(map[uint32]int)
    var v4 int64 = -728
    var v5 string = ""
    var v6 int16 = 333
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 float32 = 3.14159
            _ = fi8
            _ = v9
        }
        if false {
            v5 = v5
            v6 += -461
        }
        _ = fi7
    }
    if true {
        v6 = 32767
        if (930 <= 42) {
            var v10 int8 = -109
            for fi11 := 0; fi11 < 10; fi11++ {
                for fi12 := 0; fi12 < 12; fi12++ {
                    v3[2147483647] = -1
                    v10 = (v10 + 127)
                    _ = fi12
                }
                v10 = 127
                var v13 int = fi11
                _ = fi11
                _ = v13
            }
            fmt.Println(v4, v6)
            _ = v10
        } else {
            fmt.Println(v4, v5, v5)
        }
    }
    v14 := GenS1{F0: true}
    var v15 uint32 = 348
    var v16 uint32 = 1
    var v17 int64 = v4
    if (712 >= 4294967295) {
        switch 40 {
        case 1:
            var v18 string = "a"
            _ = v18
        default:
            switch 32767 {
            case -2147483648:
                if true {
                    var v19 string = (v5 + v5)
                    v20 := GenS1{F0: false}
                    if (!true) {
                        fmt.Println(v4, v19, v16)
                    }
                    _ = v19
                    _ = v20
                }
            case -9223372036854775808:
                if true {
                    v1 = append(v1, "\n\t")
                    for fi21 := 0; fi21 < 18; fi21++ {
                        v22 := GenS1{F0: (0 < 22)}
                        _ = fi21
                        _ = v22
                    }
                } else {
                    if (0 > v4) {
                        switch v5 {
                        case "\n\t":
                            v23 := GenS1{F0: true}
                            _ = v23
                        default:
                            if (-1 > 1) {
                                fmt.Println(v17)
                                var v24 int64 = v17
                                var v25 uint8 = 2
                                _ = v24
                                _ = v25
                            }
                        }
                        for fi26 := 0; fi26 < 19; fi26++ {
                            var v27 int8 = 42
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v4 = v4
                            _ = fi28
                        }
                        var v29 int8 = -128
                        delete(v3, v16)
                        _ = v29
                    }
                    v30 := GenS1{F0: true}
                    var v31 float64 = float64(3.14159)
                    _ = v30
                    _ = v31
                }
            case 0:
                fmt.Println(v17, v17)
            default:
                var v32 uint = 4294967295
                _ = v32
            }
        }
    } else {
        fmt.Println(v17)
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    var v33 int16 = (^(v6))
    _ = len(v1)
    _ = len(v3)
    v1 = append(v1, (("hello" + (("foo bar" + (("hello" + v5) + v5)) + v5)) + ("hello" + (("foo bar" + v5) + v5))))
    v17 = v4
    v6 = (v33 % ((v33) | 1))
    v17 = 745
    if false {
        v4 = (-(v4))
    }
    v1 = append(v1, v5)
    if (127 < 127) {
        v34 := GenS1{F0: (false || false)}
        v15 += v16
        _ = v34
    } else {
        var v35 int = -2147483648
        if (!(!false)) {
            _ = len(v1)
        } else {
            var v36 uint16 = 2
            _ = v36
        }
        fmt.Println(v4, v15)
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
