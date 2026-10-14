package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int16; F2 int16 }
func genFn0(p0 bool, p1 int32, p2 uint8) bool {
    if p0 {
        for fi5 := 0; fi5 < 12; fi5++ {
            p2 = (p2 * (p2 & 128))
            _ = fi5
        }
    }
    v6 := []int{127}
    for fi7 := 0; fi7 < 10; fi7++ {
        p1 -= (p1 | (p1 / ((-533) | 1)))
        for fi8 := 0; fi8 < 8; fi8++ {
            p1 = (p1 + (p1 - p1))
            _ = fi8
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (!(p0 || p0))
}

func main() {
    var v1 bool = (false || (!(!false)))
    var v2 uint32 = 4294967295
    v3 := make(map[string]int)
    v4 := []int{42, -849, 2147483647}
    fmt.Println(v2, v2)
    v2 -= (v2 / (((v2 | (v2 + (v2 >> ((4294967295) & 63))))) | 1))
    for fi9 := 0; fi9 < 4; fi9++ {
        v2 = (^(v2))
        _ = fi9
    }
    v10 := GenS1{F0: 123, F1: 2147483647}
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            if true {
                if len(v4) > 0 {
                    v4[0] = 127
                }
            } else {
                if v1 {
                    _ = len(v4)
                }
                v13 := GenS2{F0: fi12, F1: 1}
                if v1 {
                    v14 := GenS3{F0: -58, F1: 2}
                    v4 = append(v4, fi11)
                    switch "a" {
                    case "hello":
                        if (193 != -544) {
                            v2 = (v2 * (v2 * v2))
                            var v15 bool = (18446744073709551615 >= uint64(v2))
                            _ = v15
                        }
                    case "foo bar":
                        var v16 uint32 = 65535
                        if (v1 || (-7 != 100)) {
                            v1 = (("" + "\n\t") != "")
                        } else {
                            var v17 uint16 = 0
                            fmt.Println(v2, v17, v1)
                            v18 := GenS1{F0: 127, F1: 0}
                            _ = v17
                            _ = v18
                        }
                        _ = v16
                    case "":
                        var v19 int16 = 32767
                        _ = v19
                    }
                    _ = v14
                }
                v20 := GenS4{F0: 42, F1: -1}
                var v21 int32 = -2147483648
                _ = v13
                _ = v20
                _ = v21
            }
            _ = fi12
        }
        var v22 int = (fi11 | fi11)
        _ = fi11
        _ = v22
    }
    fmt.Println(v1, v1)
    for fi23 := 0; fi23 < 11; fi23++ {
        if v1 {
            for fi24 := 0; fi24 < 10; fi24++ {
                if v1 {
                    switch fi23 {
                    case -300:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v4 = append(v4, fi23)
                    v25 := GenS2{F0: (fi24 << ((4294967295) & 63)), F1: 0}
                    _ = v25
                }
                var v26 float64 = float64(1.5)
                _ = fi24
                _ = v26
            }
            v27 := GenS2{F0: 9223372036854775807, F1: -1}
            _ = v27
        } else {
            fmt.Println(v2, fi23)
        }
        var v28 int = fi23
        _ = fi23
        _ = v28
    }
    v1 = v1
    for fi29 := 0; fi29 < 17; fi29++ {
        v2 += v2
        _ = fi29
    }
    if v1 {
        v1 = (true && v1)
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 float32 = 1.5
            v31 = v31
            _ = fi30
            _ = v31
        }
        v2 = (v2 / ((v2) | 1))
    }
    if (-419 >= 0) {
        if (!(float32(0.1) == (0.0 * 3.14159))) {
            var v32 int = 1
            _ = v32
        }
        for fi33 := 0; fi33 < 7; fi33++ {
            for fi34 := 0; fi34 < 18; fi34++ {
                var v35 int32 = -2147483648
                _ = fi34
                _ = v35
            }
            for fi36 := 0; fi36 < 12; fi36++ {
                var v37 uint64 = 1
                var v38 uint16 = 803
                v39 := v3[("\n\t" + "")]
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
            var v40 string = ""
            _ = fi33
            _ = v40
        }
        var v41 int32 = 42
        _ = v41
    } else {
        switch "hello" {
        case "hello":
            v1 = false
        case "":
            v42 := GenS5{F0: 32768, F1: 171, F2: 0}
            _ = v42
        case "\n\t":
            for fi43 := 0; fi43 < 5; fi43++ {
                v2 -= v2
                switch "hello" {
                case "\n\t":
                    if (v1 || v1) {
                        v1 = false
                        var v44 uint8 = 127
                        _ = v44
                    }
                case "":
                    fmt.Println(v2, fi43)
                default:
                    var v45 uint8 = 127
                    _ = v45
                }
                for fi46 := 0; fi46 < 1; fi46++ {
                    v47 := GenS6{F0: 0}
                    v48 := v3["hello"]
                    fmt.Println(v2, v2)
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                _ = fi43
            }
        default:
            switch "foo bar" {
            case "a":
                for fi49 := 0; fi49 < 2; fi49++ {
                    switch (("a" + "a") + "\n\t") {
                    case "a":
                        v50 := GenS3{F0: 127, F1: 1}
                        _ = v50
                    case "foo bar":
                        v51 := GenS1{F0: 175, F1: 0}
                        _ = v51
                    case "hello":
                        switch ("foo bar" + "a") {
                        case "a":
                            v52 := GenS5{F0: fi49, F1: 235, F2: 2}
                            _ = v52
                        default:
                            v53 := GenS7{F0: 2, F1: 735, F2: 0}
                            _ = v53
                        }
                    }
                    _ = fi49
                }
            }
        }
    }
    fmt.Println(v1, v1)
    if true {
        v3["hello"] = 67
        var v54 int64 = 9223372036854775807
        _ = v54
    } else {
        for fi55 := 0; fi55 < 19; fi55++ {
            v1 = (!true)
            for fi56 := 0; fi56 < 9; fi56++ {
                v57 := GenS1{F0: 127, F1: -493}
                _ = fi56
                _ = v57
            }
            v2 -= 42
            _ = fi55
        }
    }
    var v58 int = -361
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v58
}
