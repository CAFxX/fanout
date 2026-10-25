package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint8, p1 uint16, p2 int16) uint64 {
    var v6 uint8 = p0
    fmt.Println(p0, v6, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 int32 = 2147483647
    var v2 int64 = -9223372036854775808
    var v3 int64 = (v2 - v2)
    var v4 uint16 = 0
    var v5 int32 = v1
    switch 42 {
    case 32767:
        var v7 int8 = 1
        _ = v7
    case 127:
        var v8 float32 = (1e10 + -2.25)
        _ = v8
    }
    v3 -= (v2 * 9223372036854775807)
    if (!(!true)) {
        v9 := GenS1{F0: (v5 | (v1 >> ((1) & 63))), F1: 190}
        if true {
            v10 := []int{}
            if true {
                for fi11 := 0; fi11 < 14; fi11++ {
                    switch 228 {
                    case 32768:
                        var v12 uint32 = 2147483648
                        _ = v12
                    case 42:
                        if ((!(261 <= 617)) || (!(1.5 < 0.0))) {
                            for fi13 := 0; fi13 < 11; fi13++ {
                                var v14 uint32 = 415
                                switch int(v14) {
                                case -121:
                                    if (!(float32((1.5 * 3.14159)) != 3.14159)) {
                                        fmt.Println(fi11)
                                        v4 += v4
                                    } else {
                                        switch (fi13 & fi11) {
                                        case 1:
                                            fmt.Println(fi11, fi11)
                                        case -805:
                                            v4 = v4
                                        case 128:
                                            var v15 uint = 18446744073709551615
                                            _ = v15
                                        default:
                                            var v16 int64 = v2
                                            _ = v16
                                        }
                                    }
                                case 2147483647:
                                    v5 = 0
                                default:
                                    var v17 uint16 = 1
                                    _ = v17
                                }
                                var v18 string = "a"
                                _ = fi13
                                _ = v14
                                _ = v18
                            }
                            for fi19 := 0; fi19 < 2; fi19++ {
                                v3 = (v2 % (((v3 >> ((510) & 63))) | 1))
                                switch "a" {
                                case "":
                                    var v20 int = fi11
                                    _ = v20
                                case "foo bar":
                                    if len(v10) > 0 {
                                        v10[0] = (-(fi19))
                                    }
                                }
                                v5 -= v1
                                fmt.Println(v3)
                                _ = fi19
                            }
                        }
                    case -824:
                        var v21 uint8 = 128
                        _ = v21
                    default:
                        v4 = 21
                    }
                    v10 = append(v10, fi11)
                    _ = fi11
                }
            }
            v3 = (v2 + v2)
            _ = v10
        } else {
            var v22 int64 = v3
            _ = v22
        }
        _ = v9
    }
    var v23 int8 = 42
    var v24 uint64 = 4
    v24 += v24
    switch "foo bar" {
    case "hello":
        v5 -= 1
    case "\n\t":
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 int = fi25
            fmt.Println(v4, v2, v26)
            v24 += genFn0(1, (v4 + v4), 0)
            _ = fi25
            _ = v26
        }
    }
    v5 = (v5 >> ((920) & 63))
    if (false && false) {
        if false {
            for fi27 := 0; fi27 < 16; fi27++ {
                v2 -= (v3 * (v2 ^ -34))
                v28 := []int{32768, -2147483648, (fi27 - fi27)}
                switch ("foo bar" + ("hello" + "\n\t")) {
                case "\n\t":
                    v29 := GenS1{F0: -947, F1: 24}
                    _ = v29
                case "":
                    for fi30 := 0; fi30 < 15; fi30++ {
                        var v31 uint = 617
                        switch (fi27 | (fi27 / ((32768) | 1))) {
                        case -1:
                            if len(v28) > 0 {
                                v28[0] = fi30
                            }
                        case 603:
                            var v32 int64 = v2
                            _ = v32
                        case 328:
                            v31 = v31
                        }
                        if (!(!false)) {
                            var v33 int8 = int8(v4)
                            var v34 uint32 = 42
                            var v35 int = -9223372036854775808
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        _ = fi30
                        _ = v31
                    }
                case "a":
                    v2 += v2
                }
                _ = fi27
                _ = v28
            }
            var v36 int = 128
            v37 := []int{(v36 & v36)}
            _ = v36
            _ = v37
        }
    } else {
        v38 := []int{32768, int(v4), 9223372036854775807}
        for fi39 := 0; fi39 < 1; fi39++ {
            v5 = (v1 + int32(v24))
            _ = fi39
        }
        _ = v38
    }
    v1 -= v5
    var v40 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
