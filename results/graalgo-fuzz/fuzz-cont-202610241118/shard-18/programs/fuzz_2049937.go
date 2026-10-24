package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
type GenS2 struct { F0 uint16; F1 bool }

func main() {
    var v1 string = "\n\t"
    var v2 int8 = -128
    v3 := []string{v1, v1}
    v4 := make(map[string]int)
    var v5 uint32 = 647
    if true {
        var v6 int8 = (-(v2))
        v4[v1] = 28
        if (30 < -46) {
            if (!(4294967295 != 0)) {
                v1 = ((v1 + v1) + (v1 + "a"))
                var v7 int16 = 42
                _ = v7
            }
            fmt.Println(v5, v2)
            v3 = append(v3, ((v1 + (v1 + ("foo bar" + v1))) + v1))
        }
        _ = v6
    }
    var v8 uint = 564
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v9 float64 = float64(-2.25)
    for fi10 := 0; fi10 < 12; fi10++ {
        v5 = 42
        fmt.Println(v2)
        v1 = (v1 + "foo bar")
        _ = fi10
    }
    switch "hello" {
    case "a":
        v8 += 275
    case "foo bar":
        for fi11 := 0; fi11 < 4; fi11++ {
            v2 -= 100
            var v12 int32 = -2147483648
            _ = fi11
            _ = v12
        }
    case "\n\t":
        v13 := GenS1{F0: -1, F1: v1, F2: (0 <= 0)}
        _ = v13
    default:
        var v14 float32 = 0.0
        _ = v14
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (v9 < v9) {
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := GenS1{F0: 0, F1: ((v1 + "foo bar") + (v1 + v1)), F2: true}
            v17 := GenS2{F0: 42, F1: (v2 < (v2 & 0))}
            for fi18 := 0; fi18 < 5; fi18++ {
                if true {
                    if (!true) {
                        v8 -= (v8 << (((v8 - 65535)) & 63))
                        fmt.Println(v1)
                    } else {
                        v8 = (v8 ^ v8)
                        v5 = 0
                        var v19 int = fi18
                        _ = v19
                    }
                    switch (v1 + (v1 + "hello")) {
                    case "\n\t":
                        var v20 uint64 = 664
                        _ = v20
                    case "a":
                        fmt.Println(v5, v1, v5)
                    case "foo bar":
                        v2 = (v2 / ((0) | 1))
                    }
                } else {
                    switch v1 {
                    case "foo bar":
                        if (fi18 != -2147483648) {
                            switch (fi18 & 4294967295) {
                            case 0:
                                delete(v4, v1)
                            }
                        }
                    case "":
                        v1 = ("a" + "")
                    default:
                        v1 = "a"
                        switch fi15 {
                        case 903:
                            var v21 bool = (false && true)
                            _ = v21
                        }
                        var v22 uint64 = 42
                        _ = v22
                    }
                    var v23 uint16 = 1
                    fmt.Println(v5, v1, v9)
                    _ = v23
                }
                var v24 int8 = 100
                if true {
                    if (!true) {
                        fmt.Println(v24)
                        var v25 int16 = 0
                        var v26 int8 = -1
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS2{F0: 407, F1: ((v8 % (((v8 - (v8 % (((v8 + v8)) | 1)))) | 1)) != 65535)}
                    var v28 uint = 4294967295
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(v2, fi18)
                }
                _ = fi18
                _ = v24
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            var v30 string = ("hello" + "")
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 15; fi31++ {
            var v32 float32 = float32(((-2.25 * (1.5 / float64(float32((v9 - 0.1))))) / float64(-2.25)))
            v3 = append(v3, "\n\t")
            _ = fi31
            _ = v32
        }
    } else {
        switch "" {
        case "":
            v33 := GenS1{F0: 0, F1: v1, F2: false}
            _ = v33
        case "\n\t":
            var v34 int = 128
            _ = v34
        }
        if (-1 < 9223372036854775807) {
            _ = len(v4)
            if (-2147483648 < -906) {
                var v35 string = "hello"
                var v36 float32 = ((((float32((v9 / -2.25)) * float32(v9)) * 0.1) - 1e10) + 1.5)
                var v37 int = 42
                _ = v35
                _ = v36
                _ = v37
            }
        } else {
            var v38 uint16 = 255
            _ = v38
        }
        v39 := GenS2{F0: 32767, F1: true}
        _ = v39
    }
    fmt.Println(v8)
    if len(v3) > 0 {
        v3[0] = (v1 + "a")
    }
    if true {
        var v40 uint32 = v5
        v8 = v8
        var v41 uint16 = 1
        _ = v40
        _ = v41
    } else {
        v3 = append(v3, "\n\t")
        for fi42 := 0; fi42 < 10; fi42++ {
            fmt.Println(v8, fi42)
            _ = fi42
        }
    }
    switch "foo bar" {
    case "hello":
        v43 := GenS2{F0: 304, F1: (true || true)}
        _ = v43
    case "":
        var v44 int8 = v2
        _ = v44
    default:
        var v45 int8 = 39
        _ = v45
    }
    switch ((("\n\t" + v1) + v1) + "") {
    case "":
        if (!((true && (-755 <= 0)) && (!true))) {
            var v46 int32 = -1
            for fi47 := 0; fi47 < 10; fi47++ {
                v1 = v1
                switch (v1 + v1) {
                case "hello":
                    v1 = "a"
                case "":
                    if true {
                        var v48 string = (v1 + v1)
                        _ = v48
                    }
                case "foo bar":
                    for fi49 := 0; fi49 < 18; fi49++ {
                        var v50 uint32 = v5
                        _ = fi49
                        _ = v50
                    }
                }
                _ = fi47
            }
            _ = v46
        } else {
            v2 += v2
            v9 -= (float64((0.1 * ((-2.25 * (-2.25 * ((1e10 / -2.25) * 3.14159))) - 1.5))) + ((v9 * 0.1) * v9))
            if (v2 >= 1) {
                var v51 int32 = -123
                var v52 int8 = (v2 % (((v2 << ((929) & 63))) | 1))
                _ = v51
                _ = v52
            }
        }
    case "a":
        v3 = append(v3, ("a" + "a"))
    case "\n\t":
        var v53 string = v1
        _ = v53
    }
    v1 = v1
    if true {
        v54 := GenS1{F0: 1, F1: (v1 + ""), F2: (!true)}
        if ((v5 == 0) || (float64(0.1) >= v9)) {
            v55 := GenS2{F0: 0, F1: false}
            v9 += v9
            var v56 uint = 0
            _ = v55
            _ = v56
        }
        v57 := GenS2{F0: 65535, F1: true}
        _ = v54
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
