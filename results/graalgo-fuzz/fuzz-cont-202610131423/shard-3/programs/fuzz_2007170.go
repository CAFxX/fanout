package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int32; F2 uint }
type GenS5 struct { F0 uint16; F1 int; F2 uint32 }
type GenS6 struct { F0 uint32; F1 int64; F2 bool }
type GenS7 struct { F0 int }
func genFn0(p0 uint64) bool {
    if true {
        var v6 uint32 = 42
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 string = ((("foo bar" + ("a" + ("a" + "hello"))) + (("a" + "\n\t") + "\n\t")) + "foo bar")
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        v6 = v6
        _ = v6
    } else {
        var v10 float32 = (1.5 / 1e10)
        v11 := GenS1{F0: 0, F1: -1}
        p0 = p0
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 1000}
    var v13 int64 = 42
    fmt.Println(v13, v13)
    for fi14 := 0; fi14 < 4; fi14++ {
        p0 = p0
        for fi15 := 0; fi15 < 12; fi15++ {
            if false {
                fmt.Println(v13)
                v16 := []int{fi15, int(p0), (fi15 + (fi15 << ((4294967295) & 63)))}
                v13 = (v13 + v13)
                _ = v16
            }
            var v17 int16 = 42
            if (!(!(!(!true)))) {
                var v18 uint64 = (p0 ^ p0)
                _ = v18
            }
            _ = fi15
            _ = v17
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(v13, v13, fi19)
            _ = fi19
        }
        _ = fi14
    }
    _ = p0
    _ = v12
    _ = v13
    return true
}

func main() {
    var v1 int16 = -245
    v2 := []uint32{1, 42}
    var v3 int16 = (v1 - (-(v1)))
    var v4 bool = false
    var v5 string = ""
    v20 := GenS3{F0: 2}
    var v21 float64 = (((3.14159 * (0.1 - 1e10)) - (0.1 * float64((3.14159 * (1.5 / 1.0))))) + (3.14159 * 1.5))
    v3 = 0
    if v4 {
        v2 = append(v2, 451)
    } else {
        v22 := GenS2{F0: 1}
        if len(v2) > 0 {
            v2[0] = 646
        }
        var v23 int32 = -189
        _ = v22
        _ = v23
    }
    if v4 {
        for fi24 := 0; fi24 < 7; fi24++ {
            _ = len(v2)
            v25 := GenS4{F0: fi24, F1: -2147483648, F2: 2}
            _ = fi24
            _ = v25
        }
        v1 = (v3 & (v3 + (v3 - 42)))
        v26 := GenS1{F0: 32768, F1: (v3 % (((v3 % ((323) | 1))) | 1))}
        _ = v26
    }
    switch 128 {
    case -9223372036854775808:
        v27 := GenS2{F0: v1}
        _ = v27
    case 32768:
        v28 := GenS3{F0: 555}
        _ = v28
    case -949:
        var v29 int = 2147483647
        _ = v29
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    var v30 int64 = -1
    v31 := GenS4{F0: -898, F1: 42, F2: 4294967295}
    v1 += v3
    v32 := GenS5{F0: 659, F1: 179, F2: 42}
    fmt.Println(v30, v21, v4)
    if v4 {
        v33 := GenS6{F0: 2147483647, F1: v30, F2: v4}
        v4 = (false || v4)
        _ = v33
    } else {
        switch "foo bar" {
        case "foo bar":
            for fi34 := 0; fi34 < 1; fi34++ {
                var v35 uint8 = 0
                for fi36 := 0; fi36 < 6; fi36++ {
                    v4 = (-2147483648 == -628)
                    if v4 {
                        v2 = append(v2, 4294967295)
                        if false {
                            v2 = append(v2, 2147483648)
                            fmt.Println(v35, v5)
                            fmt.Println(v4)
                        }
                        switch 2147483647 {
                        case -1:
                            v30 += (v30 ^ (^(v30)))
                        }
                    }
                    v37 := GenS3{F0: 2}
                    _ = fi36
                    _ = v37
                }
                var v38 uint32 = 561
                _ = fi34
                _ = v35
                _ = v38
            }
        case "a":
            var v39 uint32 = 2147483648
            _ = v39
        }
    }
    v30 += v30
    fmt.Println(v3)
    if true {
        if true {
            _ = len(v2)
            for fi40 := 0; fi40 < 5; fi40++ {
                if false {
                    v5 = "a"
                    for fi41 := 0; fi41 < 17; fi41++ {
                        v1 += v3
                        v2 = append(v2, 880)
                        switch -2147483648 {
                        case -39:
                            if v4 {
                                var v42 int16 = (v1 ^ (-(v3)))
                                var v43 uint8 = 2
                                var v44 bool = (!((v30 <= v30) && false))
                                v30 += (v30 + (v30 - -9223372036854775808))
                                fmt.Println(v43)
                                v44 = v4
                                v3 += (v1 % ((32767) | 1))
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                        case -1:
                            fmt.Println(v30, v1)
                        }
                        v2 = append(v2, 531)
                        v2 = append(v2, 2)
                        _ = fi41
                    }
                } else {
                    for fi45 := 0; fi45 < 19; fi45++ {
                        if (1 != 214) {
                            var v46 uint64 = 602
                            v47 := GenS7{F0: fi40}
                            _ = v46
                            _ = v47
                        } else {
                            var v48 int = (fi40 - (-(fi45)))
                            switch v48 {
                            case 2147483647:
                                v21 += v21
                            case 128:
                                v1 = (v3 + 0)
                            case 42:
                                switch v5 {
                                case "":
                                    switch (v5 + "a") {
                                    case "foo bar":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "":
                                        v4 = false
                                    }
                                case "a":
                                    var v49 int16 = v3
                                    _ = v49
                                default:
                                    if v4 {
                                        v2 = append(v2, 4294967295)
                                        v50 := GenS3{F0: 986}
                                        _ = v50
                                    } else {
                                        if v4 {
                                            v51 := GenS1{F0: fi45, F1: 1000}
                                            v21 += float64((3.14159 * -2.25))
                                            _ = v51
                                        }
                                        v21 = v21
                                        var v52 float64 = ((1e10 - (float64(1.5) * v21)) + 0.0)
                                        _ = v52
                                    }
                                }
                            default:
                                var v53 float32 = ((float32(v21) - (1e10 * 0.1)) - -2.25)
                                _ = v53
                            }
                            _ = v48
                        }
                        if (235 < 0) {
                            var v54 uint32 = 2147483647
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                            v55 := GenS2{F0: (-(v3))}
                            _ = v54
                            _ = v55
                        }
                        _ = fi45
                    }
                }
                var v56 int64 = v30
                _ = fi40
                _ = v56
            }
        }
        var v57 uint32 = 210
        fmt.Println(v30, v5, v5)
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
