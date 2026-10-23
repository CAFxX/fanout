package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint16; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint16; F2 uint }
func genFn0() uint16 {
    var v4 float64 = -2.25
    var v5 int64 = 538
    v5 -= 9223372036854775807
    var v6 int32 = 1
    _ = v4
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v7 bool = true
    var v8 int8 = 42
    var v9 string = "a"
    if v7 {
        v7 = (0 == -110)
    } else {
        if v7 {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v10 int32 = 2147483647
            _ = v10
        }
        v8 = 100
    }
    if v7 {
        if (true && (((1.5 + 3.14159) * 0.1) < 0.1)) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch v9 {
            case "a":
                var v11 uint64 = 42
                _ = v11
            default:
                fmt.Println(v8, v7, v9)
            }
            switch ((((("\n\t" + v9) + ("foo bar" + v9)) + v9) + v9) + v9) {
            case "\n\t":
                for fi12 := 0; fi12 < 1; fi12++ {
                    switch fi12 {
                    case 32768:
                        if false {
                            if (!v7) {
                                fmt.Println(v9)
                            }
                        }
                    case 127:
                        var v13 int = (fi12 << ((501) & 63))
                        _ = v13
                    case 881:
                        v8 = (v8 | (v8 / (((v8 / ((v8) | 1))) | 1)))
                    default:
                        v14 := GenS1{F0: (fi12 + fi12)}
                        _ = v14
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if false {
                        fmt.Println(v7, fi12)
                        switch fi12 {
                        case 4294967295:
                            for fi15 := 0; fi15 < 16; fi15++ {
                                v16 := GenS1{F0: int(v8)}
                                _ = fi15
                                _ = v16
                            }
                        }
                    } else {
                        var v17 int = fi12
                        v8 = (v8 & v8)
                        _ = v17
                    }
                    _ = fi12
                }
            default:
                v7 = (!(false && ((v9 + "\n\t") < ("a" + ""))))
            }
        } else {
            var v18 int16 = 60
            if v7 {
                v8 = v8
                for fi19 := 0; fi19 < 19; fi19++ {
                    if (1 >= 42) {
                        switch fi19 {
                        case -764:
                            fmt.Println(v9, fi19)
                        case 99:
                            v7 = (fi19 != fi19)
                        case 231:
                            for fi20 := 0; fi20 < 1; fi20++ {
                                v7 = true
                                _ = fi20
                            }
                        }
                        for fi21 := 0; fi21 < 10; fi21++ {
                            if (!false) {
                                switch "foo bar" {
                                case "a":
                                    v22 := GenS1{F0: (fi21 - fi19)}
                                    _ = v22
                                }
                                var v23 uint32 = 447
                                _ = v23
                            } else {
                                var v24 bool = v7
                                v3[794] = 42
                                var v25 int8 = 57
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS2{F0: 65535, F1: 2, F2: v9}
                            _ = fi21
                            _ = v26
                        }
                        var v27 int8 = (-(v8))
                        _ = v27
                    }
                    _ = fi19
                }
                switch ("\n\t" + ("\n\t" + "\n\t")) {
                case "":
                    var v28 uint64 = 726
                    _ = v28
                case "hello":
                    if v7 {
                        if v7 {
                            v9 = v9
                            var v29 uint8 = 1
                            _ = v29
                        }
                    } else {
                        if (v9 > v9) {
                            var v30 int16 = int16(v8)
                            var v31 int64 = int64(v8)
                            fmt.Println(v18, v30, v18)
                            _ = v30
                            _ = v31
                        } else {
                            v8 += -128
                            switch v9 {
                            case "\n\t":
                                v32 := GenS3{F0: 804}
                                _ = v32
                            case "":
                                switch ("hello" + "") {
                                case "foo bar":
                                    v18 += v18
                                default:
                                    v7 = (v7 && (false || v7))
                                }
                            default:
                                for fi33 := 0; fi33 < 14; fi33++ {
                                    for fi34 := 0; fi34 < 9; fi34++ {
                                        v9 = v9
                                        v35 := GenS4{F0: 0, F1: 255, F2: 42}
                                        _ = len(v2)
                                        _ = fi34
                                        _ = v35
                                    }
                                    var v36 int64 = -9223372036854775808
                                    _ = fi33
                                    _ = v36
                                }
                            }
                            _ = len(v1)
                        }
                        for fi37 := 0; fi37 < 13; fi37++ {
                            switch "a" {
                            case "hello":
                                v38 := v3[128]
                                _ = v38
                            case "":
                                if v7 {
                                    for fi39 := 0; fi39 < 11; fi39++ {
                                        if len(v2) > 0 {
                                            v2[0] = ("hello" + v9)
                                        }
                                        var v40 uint = 4294967295
                                        if false {
                                            fmt.Println(v7, fi39, v18)
                                        }
                                        v9 = "foo bar"
                                        v41 := GenS4{F0: 0, F1: 2, F2: 0}
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                    var v42 uint64 = 42
                                    _ = v42
                                } else {
                                    switch (v9 + v9) {
                                    case "foo bar":
                                        fmt.Println(v7)
                                    }
                                }
                            case "\n\t":
                                for fi43 := 0; fi43 < 19; fi43++ {
                                    fmt.Println(v18)
                                    _ = fi43
                                }
                            }
                            _ = fi37
                        }
                    }
                case "foo bar":
                    v9 = ""
                default:
                    switch ("\n\t" + "") {
                    case "foo bar":
                        v8 -= (v8 >> ((4294967295) & 63))
                    case "\n\t":
                        switch "" {
                        case "foo bar":
                            v9 = v9
                        case "a":
                            var v44 bool = false
                            _ = v44
                        }
                    }
                }
            } else {
                fmt.Println(v18, v8)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v9 = (v9 + v9)
            }
            if ((-2.25 * -2.25) != (0.0 - 1e10)) {
                v45 := GenS3{F0: -2147483648}
                _ = v45
            }
            _ = v18
        }
    } else {
        for fi46 := 0; fi46 < 13; fi46++ {
            var v47 int8 = -96
            _ = fi46
            _ = v47
        }
        var v48 float64 = 3.14159
        _ = v48
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        for fi50 := 0; fi50 < 16; fi50++ {
            if ((42 != (fi50 * fi50)) || (3.14159 > (3.14159 / 1.0))) {
                v9 = v9
            }
            _ = fi50
        }
        var v51 uint8 = 207
        _ = fi49
        _ = v51
    }
    switch "\n\t" {
    case "\n\t":
        v9 = ((v9 + v9) + (("foo bar" + "") + v9))
    }
    v7 = (!(v9 <= "\n\t"))
    var v52 uint = 1
    v9 = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v7
    _ = v8
    _ = v9
}
