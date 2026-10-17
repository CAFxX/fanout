package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }
type GenS3 struct { F0 int }
func genFn0() uint {
    var v6 int32 = 42
    v7 := GenS1{F0: 2147483648, F1: 417, F2: 1}
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int = -1
    var v2 string = "hello"
    v3 := []string{v2, "hello", "foo bar", "a"}
    var v4 bool = (false && ((v2 + v2) > ((("hello" + "hello") + "\n\t") + v2)))
    var v5 uint16 = 2
    v1 = (v1 ^ (-(v1)))
    if false {
        v8 := GenS2{F0: 2, F1: 1, F2: 0}
        _ = v8
    }
    var v9 float64 = (1e10 - 1.5)
    switch (^(v1)) {
    case 969:
        for fi10 := 0; fi10 < 14; fi10++ {
            _ = len(v3)
            if (true || ((!(!(-32768 >= 32767))) || ((!(102 < 14)) || false))) {
                for fi11 := 0; fi11 < 13; fi11++ {
                    if len(v3) > 0 {
                        v3[0] = ("a" + v2)
                    }
                    _ = fi11
                }
                v2 = (v2 + ((v2 + v2) + v2))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                v4 = ((v1 % ((fi10) | 1)) <= fi10)
                var v12 float32 = float32(v9)
                var v13 uint16 = (^(v5))
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    case 2147483647:
        v4 = (!v4)
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS1{F0: 1, F1: 0, F2: 0}
        if (v4 || v4) {
            switch (v1 | (^(v1))) {
            case -2147483648:
                v2 = ""
            case 32768:
                v16 := GenS2{F0: uint(v5), F1: -315, F2: 2147483647}
                _ = v16
            default:
                if v4 {
                    if v4 {
                        switch (fi14 & v1) {
                        case 4294967295:
                            fmt.Println(fi14, v1, v1)
                        case -2147483648:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    } else {
                        fmt.Println(v2)
                        var v17 uint64 = 213
                        _ = v17
                    }
                    if ((-2.25 + 3.14159) < (-2.25 - ((-2.25 * -2.25) / 1.0))) {
                        if false {
                            fmt.Println(v4, v4)
                        } else {
                            fmt.Println(v2, v5, v2)
                            fmt.Println(v1, v2, v5)
                            for fi18 := 0; fi18 < 9; fi18++ {
                                switch v2 {
                                case "":
                                    var v19 uint16 = v5
                                    _ = v19
                                case "hello":
                                    v20 := GenS1{F0: 165, F1: 255, F2: 128}
                                    _ = v20
                                case "a":
                                    var v21 string = ((((v2 + v2) + ("\n\t" + v2)) + "foo bar") + ((v2 + v2) + v2))
                                    _ = v21
                                }
                                v1 = (fi18 >> ((2) & 63))
                                _ = fi18
                            }
                        }
                        fmt.Println(v4, fi14, v1)
                    }
                    var v22 int = (v1 ^ 113)
                    _ = v22
                } else {
                    v23 := GenS3{F0: v1}
                    _ = v23
                }
            }
            v3 = append(v3, (("\n\t" + (((v2 + v2) + "a") + v2)) + "foo bar"))
        } else {
            v5 -= (v5 + (v5 | v5))
            fmt.Println(v1)
            v3 = append(v3, v2)
        }
        v24 := GenS3{F0: (^(fi14))}
        _ = fi14
        _ = v15
        _ = v24
    }
    var v25 int64 = 1
    var v26 float64 = (3.14159 + ((v9 + (-2.25 * v9)) / 1.5))
    v26 -= (v26 + 0.0)
    v9 = 0.1
    var v27 int16 = -1
    switch v1 {
    case 231:
        if ((v5 <= v5) || v4) {
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 uint = genFn0()
                fmt.Println(v4)
                var v30 int16 = 0
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = len(v3)
            for fi31 := 0; fi31 < 20; fi31++ {
                if len(v3) > 0 {
                    v3[0] = ("hello" + v2)
                }
                _ = fi31
            }
        }
    case 952:
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 bool = false
            v26 = ((((v9 + ((1.5 / 0.1) * v9)) * (((v9 - 0.1) * 1.5) * 3.14159)) + v26) - float64(3.14159))
            _ = fi32
            _ = v33
        }
    case 913:
        v26 = (-2.25 + v9)
    }
    if false {
        var v34 string = (v2 + "a")
        if v4 {
            if len(v3) > 0 {
                v3[0] = v2
            }
        } else {
            v5 -= (v5 << ((65535) & 63))
            if false {
                v35 := GenS3{F0: v1}
                v3 = append(v3, "")
                _ = v35
            } else {
                var v36 string = "\n\t"
                _ = v36
            }
        }
        _ = v34
    } else {
        fmt.Println(v2)
    }
    v3 = append(v3, v2)
    var v37 uint16 = (v5 % ((v5) | 1))
    v4 = (v4 || v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v9
}
