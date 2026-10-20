package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int8; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 uint32 }
func genFn0(p0 string, p1 string) int8 {
    p1 = (p0 + "a")
    var v6 string = ("hello" + "a")
    p1 = "\n\t"
    _ = p0
    _ = p1
    _ = v6
    return -128
}

func main() {
    var v1 uint32 = 42
    var v2 uint16 = 1
    var v3 int8 = 100
    v4 := make(map[string]string)
    var v5 int = 4294967295
    switch ("a" + "hello") {
    case "\n\t":
        if ((false && (v1 > 2147483648)) && true) {
            var v7 int8 = v3
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 int16 = 1000
                if true {
                    _ = len(v4)
                }
                fmt.Println(v3, v7)
                _ = fi8
                _ = v9
            }
            v5 += 32768
            _ = v7
        }
    case "foo bar":
        v10 := GenS1{F0: 978}
        _ = v10
    default:
        v1 = 65535
    }
    if true {
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(v3)
            _ = fi11
        }
        _ = len(v4)
        v2 = v2
    } else {
        v1 = (v1 >> ((65535) & 63))
    }
    v12 := []int{}
    var v13 uint8 = 127
    if true {
        var v14 uint8 = 255
        for fi15 := 0; fi15 < 5; fi15++ {
            switch ("" + "\n\t") {
            case "\n\t":
                fmt.Println(fi15)
            case "foo bar":
                v1 = (v1 ^ (v1 * v1))
            default:
                _ = len(v12)
            }
            v16 := v4[""]
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    v5 = (v5 + v5)
    var v17 int8 = int8(v1)
    for fi18 := 0; fi18 < 11; fi18++ {
        switch fi18 {
        case 0:
            if (1000 == -434) {
                v19 := GenS1{F0: -1}
                for fi20 := 0; fi20 < 19; fi20++ {
                    fmt.Println(v3, fi18)
                    var v21 float64 = float64((float32(-2.25) * (0.0 - (float32(((1e10 * (3.14159 - 3.14159)) - ((float64((1e10 - 1.5)) + 0.0) * float64(((-2.25 + (1.5 * ((float32((0.0 / (float64(0.1) - 0.0))) + 3.14159) + ((3.14159 - ((3.14159 * 1e10) * 1.5)) + (1e10 / 1.0))))) - -2.25))))) / (1.5 * 1.5)))))
                    fmt.Println(v13, v3)
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                switch ("a" + (("\n\t" + "a") + "foo bar")) {
                case "\n\t":
                    _ = len(v12)
                case "foo bar":
                    v1 = 4294967295
                }
            }
        case 127:
            var v22 int = fi18
            _ = v22
        }
        _ = fi18
    }
    var v23 uint8 = 174
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 float64 = 0.1
        _ = fi24
        _ = v25
    }
    var v26 string = "foo bar"
    for fi27 := 0; fi27 < 4; fi27++ {
        if (!("a" == (v26 + (v26 + v26)))) {
            v28 := GenS2{F0: 0, F1: 0, F2: 647}
            for fi29 := 0; fi29 < 8; fi29++ {
                v23 += v13
                v13 -= v23
                fmt.Println(fi27, fi29, fi27)
                _ = fi29
            }
            var v30 float64 = -2.25
            _ = v28
            _ = v30
        } else {
            if (v23 >= uint8(v1)) {
                v3 += 42
                switch fi27 {
                case 42:
                    var v31 uint16 = v2
                    _ = v31
                case 1:
                    for fi32 := 0; fi32 < 5; fi32++ {
                        var v33 uint64 = 18446744073709551615
                        var v34 int8 = genFn0("hello", v26)
                        if true {
                            fmt.Println(v1)
                            fmt.Println(v34)
                        } else {
                            fmt.Println(fi32)
                            var v35 float32 = float32(0.0)
                            var v36 int8 = v3
                            _ = v35
                            _ = v36
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                default:
                    var v37 uint = 0
                    _ = v37
                }
            }
            if true {
                _ = len(v12)
                var v38 uint8 = v13
                _ = v38
            } else {
                if (((3.14159 * 1.5) <= 0.1) && true) {
                    var v39 float32 = 3.14159
                    v1 += v1
                    for fi40 := 0; fi40 < 8; fi40++ {
                        v4[v26] = ((v26 + (v26 + v26)) + "foo bar")
                        v41 := GenS2{F0: 1, F1: 42, F2: 0}
                        if len(v12) > 0 {
                            v12[0] = (fi40 / ((v5) | 1))
                        }
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                } else {
                    v42 := GenS2{F0: 2, F1: 9223372036854775807, F2: v1}
                    _ = v42
                }
                switch ("" + v26) {
                case "":
                    fmt.Println(v17, v2)
                }
                v43 := v4[v26]
                _ = v43
            }
        }
        v44 := v4[v26]
        if (false && (0 >= 42)) {
            if false {
                var v45 float64 = 1.5
                fmt.Println(v3, fi27, v13)
                var v46 uint32 = (v1 % ((0) | 1))
                _ = v45
                _ = v46
            }
            for fi47 := 0; fi47 < 17; fi47++ {
                v48 := GenS3{F0: 1, F1: 100, F2: ((v2 << ((430) & 63)) != (v2 >> ((1) & 63)))}
                fmt.Println(v1, v5)
                _ = fi47
                _ = v48
            }
            if true {
                if (!false) {
                    for fi49 := 0; fi49 < 18; fi49++ {
                        v50 := GenS4{F0: (v3 % ((v3) | 1))}
                        v12 = append(v12, (fi27 >> ((963) & 63)))
                        v26 = "\n\t"
                        _ = fi49
                        _ = v50
                    }
                } else {
                    if (v44 == "hello") {
                        if false {
                            fmt.Println(v13)
                            v1 -= v1
                        } else {
                            v51 := GenS3{F0: (v1 >> ((325) & 63)), F1: (v3 - -119), F2: (250 == (v13 ^ (v23 ^ (v23 | (v13 * 0)))))}
                            delete(v4, ("foo bar" + v44))
                            var v52 float32 = float32((0.1 / 3.14159))
                            _ = v51
                            _ = v52
                        }
                        fmt.Println(v26)
                        v3 -= -23
                    }
                    fmt.Println(v26)
                }
            }
        }
        _ = fi27
        _ = v44
    }
    v53 := GenS4{F0: (v17 | (v3 ^ v3))}
    var v54 uint16 = (v2 >> ((255) & 63))
    v55 := GenS5{F0: 18446744073709551615, F1: v1}
    if ((((-32768 != -260) || false) || (true || (v2 >= v54))) || (0 >= int32(v1))) {
        v56 := GenS1{F0: -153}
        v5 = v5
        _ = v56
    }
    var v57 bool = ("" <= "\n\t")
    fmt.Println(v13, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v57
}
