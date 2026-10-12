package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
func genFn0() int64 {
    var v5 int8 = -4
    var v6 int16 = 682
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            if true {
                var v9 uint8 = 2
                if false {
                    var v10 float32 = (3.14159 + ((((0.1 + 3.14159) + 0.1) / 0.1) * 0.0))
                    _ = v10
                } else {
                    if true {
                        var v11 string = "foo bar"
                        _ = v11
                    }
                }
                _ = v9
            }
            fmt.Println(fi8)
            if true {
                v6 += 42
            }
            _ = fi8
        }
        var v12 int = fi7
        for fi13 := 0; fi13 < 19; fi13++ {
            if true {
                var v14 string = "foo bar"
                _ = v14
            }
            var v15 int16 = v6
            v16 := []int{(fi13 | -2147483648), int(v15), (-(v12))}
            _ = fi13
            _ = v15
            _ = v16
        }
        _ = fi7
        _ = v12
    }
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int = 122
    v2 := make(map[int]string)
    var v3 string = "foo bar"
    var v4 int16 = -32768
    v3 = ((v3 + "foo bar") + v3)
    if ((!(128 <= 0)) && (!((true || false) || true))) {
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint32 = 238
            _ = fi17
            _ = v18
        }
        v4 = 32767
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := GenS1{F0: 2, F1: 0, F2: 255}
        _ = fi19
        _ = v20
    }
    v21 := []int{v1, 1, (v1 | -2147483648), 42}
    var v22 int16 = (-(v4))
    switch -308 {
    case 42:
        var v23 uint16 = 32767
        _ = v23
    case 1:
        v2[v1] = (v3 + ("a" + "foo bar"))
    default:
        for fi24 := 0; fi24 < 16; fi24++ {
            v21 = append(v21, fi24)
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 int = (v1 + (fi25 % (((^(fi25))) | 1)))
                v4 = (-(v22))
                switch "a" {
                case "":
                    v27 := GenS1{F0: 1, F1: (v22 << ((65535) & 63)), F2: 704}
                    _ = v27
                case "foo bar":
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                case "hello":
                    if ((!false) && (((!(2 == 116)) || (0 > 0)) || false)) {
                        var v28 uint32 = 1
                        if false {
                            v22 = -286
                            v22 += int16(v28)
                        }
                        v29 := GenS1{F0: 207, F1: -1, F2: 956}
                        v22 -= (v22 | 599)
                        if false {
                            _ = len(v2)
                        }
                        _ = v28
                        _ = v29
                    } else {
                        v30 := v2[(fi24 << ((18446744073709551615) & 63))]
                        if true {
                            var v31 int8 = 1
                            fmt.Println(v26)
                            _ = v31
                        } else {
                            var v32 int8 = int8(fi24)
                            _ = v32
                        }
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        _ = v30
                    }
                default:
                    v26 = (^(v26))
                }
                v33 := GenS1{F0: 4294967295, F1: (v4 ^ v4), F2: 255}
                v4 -= (v22 & v4)
                _ = fi25
                _ = v26
                _ = v33
            }
            _ = fi24
        }
    }
    if true {
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 string = v3
            v22 = 0
            _ = fi34
            _ = v35
        }
        v1 += v1
    }
    v1 += 0
    _ = len(v21)
    switch (v1 % ((v1) | 1)) {
    case -399:
        fmt.Println(v1, v22, v3)
    case -2147483648:
        v3 = ""
    default:
        switch v3 {
        case "\n\t":
            fmt.Println(v3, v22, v3)
        case "":
            fmt.Println(v1)
        default:
            v36 := GenS1{F0: 1, F1: -918, F2: 978}
            _ = v36
        }
    }
    if (!true) {
        fmt.Println(v22, v22, v22)
    } else {
        v37 := GenS1{F0: 2, F1: (^(v22)), F2: 42}
        _ = v37
    }
    var v38 string = "hello"
    v22 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v38
    _ = v4
}
