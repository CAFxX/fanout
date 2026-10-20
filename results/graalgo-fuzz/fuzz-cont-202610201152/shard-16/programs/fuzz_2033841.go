package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint32 = (v1 - v1)
    v3 := make(map[int]bool)
    var v4 string = ("" + "foo bar")
    v5 := []int{32768}
    v3[0] = (42 == 1)
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 1, F1: fi6}
        _ = fi6
        _ = v7
    }
    v3[32768] = (!true)
    if true {
        switch 950 {
        case 32767:
            v2 = 2
        case -9223372036854775808:
            v8 := GenS2{F0: (145 != 115), F1: 777}
            _ = v8
        case 9223372036854775807:
            v9 := GenS1{F0: 128, F1: 127}
            _ = v9
        }
        if (((("\n\t" + (v4 + v4)) + ((v4 + "foo bar") + v4)) + "") == "") {
            fmt.Println(v1, v2)
        }
    }
    switch -28 {
    case 597:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v4 = v4
        fmt.Println(v1, v4, v4)
        _ = fi10
    }
    var v11 int16 = 32767
    _ = len(v5)
    var v12 bool = ((-1 == 194) && true)
    fmt.Println(v12, v2)
    var v13 uint32 = (v1 % ((42) | 1))
    if (255 != 114) {
        v14 := GenS1{F0: 81, F1: 0}
        _ = v14
    } else {
        var v15 bool = v12
        _ = v15
    }
    var v16 float64 = (0.1 / (float64(-2.25) * float64(((((0.0 + ((1.5 * float32((0.0 - 0.0))) * 1.5)) - (float32(float64(((-2.25 + 3.14159) / 1.0))) - (0.0 + 1e10))) / 3.14159) - 1.5))))
    switch 9223372036854775807 {
    case 32768:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case -422:
        switch 4294967295 {
        case 209:
            switch "" {
            case "":
                v4 = v4
            case "hello":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "foo bar":
                v16 = 0.0
            }
        case -9223372036854775808:
            if v12 {
                for fi17 := 0; fi17 < 2; fi17++ {
                    v16 += 3.14159
                    _ = fi17
                }
            }
        case 2147483647:
            var v18 uint16 = 506
            _ = v18
        default:
            for fi19 := 0; fi19 < 1; fi19++ {
                v5 = append(v5, (fi19 - 1))
                if v12 {
                    delete(v3, fi19)
                }
                _ = fi19
            }
        }
    case 4294967295:
        v20 := GenS1{F0: 76, F1: -1}
        _ = v20
    default:
        v5 = append(v5, -984)
    }
    fmt.Println(v1, v1)
    var v21 bool = v12
    switch ("a" + v4) {
    case "":
        v2 = (v13 * (v2 ^ (^(v2))))
    case "foo bar":
        if ((1.5 + (v16 * v16)) != float64(-2.25)) {
            var v22 uint64 = 1
            var v23 float32 = 1.5
            _ = v22
            _ = v23
        }
    case "hello":
        var v24 string = v4
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        fmt.Println(v4, v12, v16)
        if ("\n\t" == ("" + "")) {
            switch (^(fi25)) {
            case 2147483647:
                fmt.Println(fi25, v21)
            case -9223372036854775808:
                var v26 int64 = -9223372036854775808
                _ = v26
            case 32767:
                fmt.Println(v11, v12, v12)
            default:
                v27 := GenS1{F0: 48, F1: 2147483647}
                _ = v27
            }
            switch "" {
            case "\n\t":
                for fi28 := 0; fi28 < 5; fi28++ {
                    fmt.Println(v11, fi25, fi25)
                    v12 = false
                    v16 = v16
                    _ = fi28
                }
            case "foo bar":
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 int64 = -232
                    v31 := GenS2{F0: (!(0 != 20)), F1: 86}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            default:
                for fi32 := 0; fi32 < 15; fi32++ {
                    if false {
                        switch -786 {
                        case -9223372036854775808:
                            v33 := GenS1{F0: 127, F1: fi25}
                            _ = v33
                        case 1:
                            if (v4 != "") {
                                v34 := GenS1{F0: 42, F1: (fi25 / ((fi32) | 1))}
                                v2 -= 207
                                _ = v34
                            } else {
                                v2 = v2
                            }
                        case 128:
                            _ = len(v3)
                        }
                        var v35 uint16 = 65535
                        _ = v35
                    }
                    fmt.Println(fi25, v1, v4)
                    switch ((v4 + v4) + (v4 + "\n\t")) {
                    case "hello":
                        if v12 {
                            var v36 float64 = v16
                            v5 = append(v5, 1)
                            _ = v36
                        }
                    default:
                        v1 = (^(v2))
                    }
                    _ = fi32
                }
            }
        } else {
            var v37 uint8 = 1
            if (1e10 != 0.0) {
                var v38 uint8 = v37
                _ = v38
            } else {
                var v39 uint32 = 4294967295
                _ = v39
            }
            _ = len(v5)
            _ = v37
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
