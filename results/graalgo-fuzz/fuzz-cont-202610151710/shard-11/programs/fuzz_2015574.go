package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 bool; F2 int64 }
type GenS5 struct { F0 uint16; F1 int64; F2 int32 }

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[uint32]bool)
    var v3 int = v1
    fmt.Println(v3, v1)
    v2[181] = true
    for fi4 := 0; fi4 < 15; fi4++ {
        switch "foo bar" {
        case "a":
            v5 := v2[4294967295]
            _ = v5
        case "hello":
            if false {
                var v6 float64 = 3.14159
                _ = v6
            }
        }
        if true {
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: (-235 != 1000), F1: 636}
                _ = fi7
                _ = v8
            }
        }
        _ = fi4
    }
    fmt.Println(v1)
    if (2147483648 <= 2) {
        v9 := GenS1{F0: false, F1: 0}
        var v10 int64 = 42
        _ = v10
        _ = v9
    } else {
        v3 = v3
        v11 := GenS2{F0: 4294967295, F1: 1}
        v3 = (-(v1))
        _ = v11
    }
    v1 -= v3
    var v12 uint = 1
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint = (^(v12))
        var v15 uint8 = 0
        fmt.Println(fi13)
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: v12, F1: 65535}
    v17 := GenS3{F0: (v12 ^ v12)}
    if true {
        for fi18 := 0; fi18 < 8; fi18++ {
            fmt.Println(v1, v3)
            v3 += v1
            fmt.Println(v1, v3, fi18)
            _ = fi18
        }
        v19 := []int{(v1 ^ (v3 >> (((v12 ^ (v12 ^ 42))) & 63))), (v3 * -138), (v3 / ((v1) | 1))}
        if (1 != 42) {
            v20 := GenS1{F0: true, F1: 227}
            v21 := GenS1{F0: ((0.0 < (float64(0.1) + 1e10)) && true), F1: 407}
            _ = v20
            _ = v21
        } else {
            v3 = v1
        }
        _ = v19
    } else {
        if ((("hello" + "") + "") < "foo bar") {
            if false {
                v22 := []int{(v1 * 52), 0, v3, 2147483647}
                var v23 string = "hello"
                _ = v22
                _ = v23
            }
            v24 := GenS4{F0: (!false), F1: true, F2: 9223372036854775807}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 8; fi25++ {
                fmt.Println(v3, fi25, v3)
                v26 := GenS5{F0: 0, F1: 42, F2: 120}
                var v27 float32 = (3.14159 + float32((float64(0.0) * ((float64(0.1) - 0.1) / (3.14159 / -2.25)))))
                _ = fi25
                _ = v26
                _ = v27
            }
            if (146 >= 325) {
                switch ((("foo bar" + ("\n\t" + ("a" + "hello"))) + ("hello" + "\n\t")) + ("" + ("\n\t" + "\n\t"))) {
                case "":
                    if (0 <= 836) {
                        v28 := GenS3{F0: (v12 ^ (^(v12)))}
                        switch "" {
                        case "foo bar":
                            var v29 int32 = -2147483648
                            _ = v29
                        case "\n\t":
                            fmt.Println(v3, v1, v3)
                        case "hello":
                            var v30 float32 = 0.1
                            _ = v30
                        default:
                            var v31 uint16 = 1
                            _ = v31
                        }
                        fmt.Println(v3)
                        _ = v28
                    } else {
                        var v32 uint64 = 42
                        switch (v1 % ((v3) | 1)) {
                        case 174:
                            var v33 int16 = -76
                            _ = v33
                        }
                        var v34 int8 = -98
                        _ = v32
                        _ = v34
                    }
                case "a":
                    var v35 uint8 = 120
                    _ = v35
                }
            }
            for fi36 := 0; fi36 < 10; fi36++ {
                v37 := []int{42}
                fmt.Println(fi36, fi36, v3)
                _ = fi36
                _ = v37
            }
        }
        switch v3 {
        case 2147483647:
            fmt.Println(v1, v3)
        case -1:
            var v38 uint32 = 333
            _ = v38
        default:
            v12 -= 43
        }
        delete(v2, 2147483647)
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 int64 = -9223372036854775808
        v3 += (v3 / (((v1 >> ((4294967295) & 63))) | 1))
        v3 += (v3 % ((42) | 1))
        _ = fi39
        _ = v40
    }
    switch "" {
    case "\n\t":
        fmt.Println(v3, v12)
    }
    v41 := v2[1]
    if (true || true) {
        if false {
            var v42 uint = v12
            _ = v42
        }
    } else {
        if (!v41) {
            var v43 bool = (!((65535 >= 255) || ((771 != 707) && false)))
            var v44 bool = ((-37 < -1) || true)
            v41 = false
            _ = v43
            _ = v44
        } else {
            if v41 {
                v1 = v1
                v12 = (v12 + v12)
                var v45 uint8 = 222
                _ = v45
            } else {
                for fi46 := 0; fi46 < 13; fi46++ {
                    if ((!true) || v41) {
                        if v41 {
                            v47 := []int{v3}
                            switch (v3 | fi46) {
                            case -9223372036854775808:
                                if (42 != 120) {
                                    if ((!((!(!(int16(v3) >= 32767))) && v41)) || (-128 < 0)) {
                                        v48 := GenS3{F0: (v12 << ((854) & 63))}
                                        _ = v48
                                    }
                                    for fi49 := 0; fi49 < 18; fi49++ {
                                        v12 = v12
                                        _ = fi49
                                    }
                                } else {
                                    fmt.Println(v41, v41, fi46)
                                    for fi50 := 0; fi50 < 16; fi50++ {
                                        var v51 int8 = -1
                                        _ = fi50
                                        _ = v51
                                    }
                                }
                            }
                            v41 = false
                            _ = v47
                        } else {
                            fmt.Println(v1, fi46)
                            var v52 float64 = float64(0.0)
                            _ = v52
                        }
                        v1 = fi46
                    } else {
                        v53 := GenS2{F0: 875, F1: 0}
                        var v54 float32 = 1e10
                        var v55 uint16 = 168
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                    fmt.Println(v3)
                    _ = fi46
                }
                var v56 int32 = -690
                _ = v56
            }
            v57 := GenS1{F0: (234 == 255), F1: 65535}
            _ = v57
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v41
}
