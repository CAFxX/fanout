package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int64; F1 int64 }
func genFn0(p0 uint8, p1 int16, p2 uint8) bool {
    var v6 uint8 = (p2 + (p2 >> ((42) & 63)))
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int8 = int8(fi8)
            _ = fi8
            _ = v9
        }
        var v10 float64 = 0.1
        _ = fi7
        _ = v10
    }
    switch 954 {
    case 32767:
        switch -2147483648 {
        case 0:
            if true {
                v11 := GenS1{F0: 2, F1: 1, F2: 4294967295}
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 int16 = (p1 ^ p1)
                    if true {
                        p2 = (v6 << ((18446744073709551615) & 63))
                    } else {
                        switch (fi12 * fi12) {
                        case -731:
                            if (1 >= 354) {
                                for fi14 := 0; fi14 < 2; fi14++ {
                                    switch ("" + "hello") {
                                    case "\n\t":
                                        v13 -= 887
                                    }
                                    v6 += p0
                                    p0 = 42
                                    _ = fi14
                                }
                            }
                        case 285:
                            var v15 string = (("foo bar" + "hello") + "a")
                            _ = v15
                        default:
                            v16 := []int{(-(fi12)), fi12}
                            _ = v16
                        }
                    }
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
        case 1:
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS1{F0: 42, F1: 127, F2: 122}
                for fi19 := 0; fi19 < 12; fi19++ {
                    fmt.Println(p1, p2, fi19)
                    _ = fi19
                }
                v20 := GenS2{F0: 1}
                _ = fi17
                _ = v18
                _ = v20
            }
        case -629:
            if (4294967295 > 18446744073709551615) {
                v21 := []int{}
                switch -1 {
                case 0:
                    if true {
                        var v22 uint16 = uint16(p0)
                        _ = v22
                    } else {
                        v6 = (^(p2))
                        if false {
                            fmt.Println(p0, p0, p1)
                        } else {
                            if true {
                                var v23 uint = 2
                                _ = v23
                            }
                            switch ("" + ((("\n\t" + "foo bar") + ("" + "")) + "foo bar")) {
                            case "a":
                                for fi24 := 0; fi24 < 10; fi24++ {
                                    v25 := GenS2{F0: 1}
                                    _ = fi24
                                    _ = v25
                                }
                            }
                        }
                        var v26 uint8 = 1
                        _ = v26
                    }
                case 32768:
                    p1 -= 32767
                case -1:
                    var v27 int8 = -91
                    _ = v27
                }
                _ = v21
            } else {
                var v28 uint16 = 273
                _ = v28
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}
func genFn1(p0 int64, p1 int16, p2 uint8) uint64 {
    var v29 float32 = (((1.5 * 1e10) + 3.14159) * 3.14159)
    if (255 < 42) {
        if true {
            v30 := []int{int(p0)}
            _ = v30
        } else {
            var v31 int16 = 1
            switch "a" {
            case "a":
                for fi32 := 0; fi32 < 3; fi32++ {
                    var v33 int32 = 42
                    _ = fi32
                    _ = v33
                }
            case "":
                v34 := []int{-775, 32767}
                _ = v34
            default:
                v35 := GenS3{F0: 2147483648}
                _ = v35
            }
            _ = v31
        }
        if false {
            v36 := []int{32768, 42, 0, 264}
            for fi37 := 0; fi37 < 4; fi37++ {
                v38 := GenS3{F0: 4294967295}
                fmt.Println(fi37, v29)
                _ = fi37
                _ = v38
            }
            _ = v36
        } else {
            v39 := GenS4{F0: -1, F1: 756, F2: 0}
            _ = v39
        }
        var v40 int32 = 42
        _ = v40
    } else {
        v41 := []int{1, 32767}
        fmt.Println(p0)
        _ = v41
    }
    var v42 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    _ = v42
    return 42
}

func main() {
    var v1 float32 = (1e10 - 1e10)
    var v2 int16 = 32767
    var v3 int64 = -1
    v4 := make(map[string]bool)
    var v5 uint32 = 65535
    var v43 uint = 18446744073709551615
    v44 := GenS3{F0: (v5 | (^(v5)))}
    var v45 uint16 = 0
    for fi46 := 0; fi46 < 1; fi46++ {
        v45 = v45
        _ = fi46
    }
    var v47 uint64 = 2
    v43 -= (v43 / ((v43) | 1))
    v3 += (v3 ^ (^(v3)))
    v48 := GenS1{F0: v45, F1: 42, F2: (v43 | 0)}
    var v49 int64 = (-(v3))
    if (false || true) {
        v50 := GenS5{F0: v3, F1: (^(v49))}
        _ = v50
    }
    fmt.Println(v49, v1, v47)
    var v51 int = 9223372036854775807
    if false {
        switch v51 {
        case 4294967295:
            if ((!(float32(1e10) >= 1.5)) || (false && false)) {
                v49 = -1
                if true {
                    v43 = v43
                }
            } else {
                fmt.Println(v51, v1, v45)
            }
        default:
            switch v51 {
            case 32767:
                var v52 uint = (v43 | 1)
                _ = v52
            case 42:
                _ = len(v4)
            case 803:
                var v53 bool = false
                _ = v53
            default:
                v47 -= v47
            }
        }
        v47 += (v47 + (v47 ^ v47))
        var v54 int64 = v49
        _ = v54
    }
    if ((v49 & 0) != (v3 * 83)) {
        var v55 int64 = (-(v49))
        v1 -= 0.0
        v55 += (-(v49))
        _ = v55
    } else {
        v49 = (v3 >> ((uint(v47)) & 63))
        fmt.Println(v3, v3)
    }
    v56 := GenS1{F0: (^(v45)), F1: -1, F2: v43}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v51
    _ = v56
}
