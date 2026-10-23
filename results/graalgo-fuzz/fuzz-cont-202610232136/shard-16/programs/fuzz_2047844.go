package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS2 struct { F0 string; F1 int16; F2 uint8 }
type GenS3 struct { F0 string; F1 string }
func genFn0() string {
    var v6 int16 = 42
    if true {
        switch "foo bar" {
        case "hello":
            for fi7 := 0; fi7 < 9; fi7++ {
                if (!true) {
                    for fi8 := 0; fi8 < 20; fi8++ {
                        fmt.Println(fi8, fi7, v6)
                        if false {
                            v9 := []int{42, (fi7 * (fi8 << ((18446744073709551615) & 63)))}
                            v6 = v6
                            _ = v9
                        }
                        _ = fi8
                    }
                } else {
                    fmt.Println(fi7, fi7)
                    v6 = (v6 ^ v6)
                    v6 -= v6
                }
                _ = fi7
            }
        case "":
            v10 := []int{}
            _ = v10
        case "a":
            for fi11 := 0; fi11 < 18; fi11++ {
                for fi12 := 0; fi12 < 7; fi12++ {
                    v6 = v6
                    switch (^(fi11)) {
                    case 95:
                        var v13 uint64 = 4294967295
                        _ = v13
                    case 2147483647:
                        if true {
                            switch 32768 {
                            case -919:
                                if true {
                                    if false {
                                        v14 := GenS1{F0: (-2.25 > ((1e10 / 1.0) + 0.0)), F1: 995, F2: 4294967295}
                                        switch "" {
                                        case "foo bar":
                                            v15 := []int{4294967295, -429}
                                            _ = v15
                                        case "\n\t":
                                            fmt.Println(fi11, fi12, v6)
                                        default:
                                            v6 -= (v6 & v6)
                                        }
                                        _ = v14
                                    } else {
                                        v16 := []int{473}
                                        _ = v16
                                    }
                                    if (971 > 0) {
                                        if (false || (833 > -9223372036854775808)) {
                                            fmt.Println(v6)
                                            fmt.Println(v6)
                                            if true {
                                                v6 -= v6
                                            }
                                        } else {
                                            var v17 uint16 = 42
                                            var v18 float64 = 1e10
                                            _ = v17
                                            _ = v18
                                        }
                                        var v19 uint = 1
                                        _ = v19
                                    } else {
                                        var v20 bool = false
                                        var v21 uint64 = 42
                                        v22 := []int{(fi11 ^ 0), 0, fi12, (^(fi12))}
                                        var v23 uint32 = 2
                                        v23 += 2
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    }
                                }
                            case 4294967295:
                                fmt.Println(fi12, v6)
                            }
                            v6 = v6
                            v6 += -596
                        }
                    case 128:
                        var v24 uint = 38
                        _ = v24
                    }
                    v6 += v6
                    _ = fi12
                }
                v6 = 42
                if ((32767 == 65535) || true) {
                    v25 := GenS2{F0: "a", F1: (v6 / (((v6 * (v6 ^ (v6 % (((v6 + (-(v6)))) | 1))))) | 1)), F2: 128}
                    if (!(false || (!(!(!false))))) {
                        switch 42 {
                        case -9223372036854775808:
                            v6 -= v6
                        }
                        v6 = (v6 << ((818) & 63))
                    }
                    fmt.Println(fi11)
                    _ = v25
                }
                _ = fi11
            }
        }
    } else {
        var v26 string = "hello"
        v27 := []int{177, 127, 471}
        var v28 float64 = 0.0
        _ = v26
        _ = v27
        _ = v28
    }
    var v29 bool = false
    _ = v29
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 32767
    var v3 int64 = -1
    var v4 uint8 = 127
    var v5 int = (v2 - v2)
    fmt.Println(v5, v3, v4)
    v30 := GenS2{F0: ("hello" + ""), F1: 1000, F2: v4}
    var v31 int64 = (v3 / ((v3) | 1))
    if true {
        v32 := []int{128, v5, v5}
        var v33 int8 = 127
        v3 = v3
        _ = v32
        _ = v33
    }
    fmt.Println(v3, v2)
    for fi34 := 0; fi34 < 13; fi34++ {
        v35 := GenS2{F0: "hello", F1: -84, F2: 0}
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        for fi37 := 0; fi37 < 12; fi37++ {
            if (true && (668 != 1)) {
                v31 = v31
            }
            if false {
                v38 := v1[v5]
                _ = v38
            } else {
                v39 := GenS3{F0: "hello", F1: "hello"}
                v40 := GenS2{F0: ((("hello" + "foo bar") + "\n\t") + "hello"), F1: -404, F2: (v4 ^ v4)}
                var v41 uint8 = (^(v4))
                _ = v39
                _ = v40
                _ = v41
            }
            _ = fi37
        }
        _ = fi36
    }
    v42 := GenS2{F0: "foo bar", F1: 1, F2: 158}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v42
    _ = v5
}
