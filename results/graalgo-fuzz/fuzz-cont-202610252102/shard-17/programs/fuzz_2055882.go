package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint8; F2 string }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := []string{"a"}
    v2 := make(map[int]bool)
    var v3 string = "a"
    var v4 uint32 = 1
    var v5 uint8 = 127
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 3; fi7++ {
        switch v3 {
        case "\n\t":
            v8 := GenS1{F0: 18446744073709551615}
            _ = v8
        case "foo bar":
            if false {
                fmt.Println(fi7)
                v9 := GenS2{F0: (v3 + ("hello" + "foo bar"))}
                _ = v9
            } else {
                if true {
                    v5 -= 56
                }
            }
        default:
            v4 -= 4294967295
        }
        _ = fi7
    }
    if false {
        if (false || false) {
            if (false || (false && true)) {
                var v10 uint = v6
                _ = len(v1)
                for fi11 := 0; fi11 < 15; fi11++ {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        switch fi12 {
                        case -407:
                            var v13 int64 = -1
                            _ = v13
                        case 1:
                            var v14 uint8 = (v5 ^ v5)
                            _ = v14
                        case 2147483647:
                            if true {
                                if false {
                                    v2[(fi12 + -2147483648)] = (false || false)
                                    v3 = ("a" + "")
                                    v15 := GenS2{F0: ""}
                                    _ = v15
                                }
                                var v16 float32 = 0.1
                                var v17 uint = uint(v5)
                                _ = v16
                                _ = v17
                            } else {
                                var v18 string = v3
                                _ = v18
                            }
                        default:
                            fmt.Println(fi12)
                        }
                        _ = fi12
                    }
                    _ = fi11
                }
                _ = v10
            }
        } else {
            _ = len(v1)
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 int32 = -213
                switch v3 {
                case "":
                    _ = len(v1)
                }
                for fi21 := 0; fi21 < 6; fi21++ {
                    v2[(fi19 - 641)] = true
                    _ = fi21
                }
                _ = fi19
                _ = v20
            }
            v5 = (v5 * 0)
        }
        var v22 bool = (4294967295 != v4)
        _ = v22
    } else {
        var v23 uint64 = 4294967295
        for fi24 := 0; fi24 < 14; fi24++ {
            if (false || (!(v3 <= (v3 + v3)))) {
                v1 = append(v1, "foo bar")
                delete(v2, fi24)
                fmt.Println(v4, v23, v23)
            }
            v25 := GenS3{F0: (v5 + 93), F1: (v5 | (v5 & 102)), F2: v3}
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 uint8 = (v5 * v5)
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        v29 := GenS4{F0: 385, F1: (-779 < -1)}
        _ = fi28
        _ = v29
    }
    var v30 int8 = -61
    if (147 >= (v4 - (v4 * (^(v4))))) {
        var v31 int32 = -539
        var v32 float32 = -2.25
        if true {
            v30 += v30
        }
        _ = v31
        _ = v32
    }
    var v33 uint8 = v5
    v1 = append(v1, "")
    var v34 uint32 = (v4 | v4)
    fmt.Println(v5)
    if (0.1 < -2.25) {
        v35 := GenS1{F0: 2}
        _ = v35
    }
    var v36 uint32 = 4294967295
    switch "" {
    case "foo bar":
        v3 = "\n\t"
    case "hello":
        v37 := GenS1{F0: 1}
        _ = v37
    case "a":
        fmt.Println(v5)
    }
    switch v3 {
    case "":
        if true {
            for fi38 := 0; fi38 < 1; fi38++ {
                v36 += (v4 ^ 2147483648)
                var v39 uint64 = 4294967295
                var v40 int16 = -1
                _ = fi38
                _ = v39
                _ = v40
            }
            for fi41 := 0; fi41 < 15; fi41++ {
                v42 := v2[320]
                var v43 uint16 = 42
                v3 = v3
                _ = fi41
                _ = v42
                _ = v43
            }
        } else {
            var v44 int64 = 42
            v1 = append(v1, v3)
            _ = v44
        }
    case "\n\t":
        var v45 float32 = 3.14159
        _ = v45
    default:
        var v46 int8 = (v30 & 100)
        _ = v46
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        v33 += 2
        if false {
            if (((!(((1.5 * 1.5) / -2.25) <= (0.0 - 1.5))) || true) || (-197 > 1)) {
                v48 := GenS2{F0: "\n\t"}
                delete(v2, (fi47 % (((fi47 << ((v6) & 63))) | 1)))
                for fi49 := 0; fi49 < 4; fi49++ {
                    v30 += (-(v30))
                    v50 := GenS5{F0: v34, F1: 2}
                    v51 := GenS1{F0: uint64(v33)}
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                _ = v48
            } else {
                v52 := v2[fi47]
                _ = v52
            }
            v53 := GenS1{F0: 1}
            var v54 int64 = 42
            _ = v53
            _ = v54
        }
        _ = fi47
    }
    var v55 int16 = -769
    if false {
        var v56 string = ("\n\t" + ((v3 + ("foo bar" + (v3 + v3))) + "hello"))
        _ = v56
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
