package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int32; F2 int16 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 int }
func genFn0(p0 uint64) int16 {
    p0 = 0
    p0 += (p0 ^ p0)
    if (!true) {
        var v7 int64 = -407
        _ = v7
    } else {
        switch -954 {
        case 9223372036854775807:
            p0 = (p0 * p0)
        case 2147483647:
            p0 -= 0
        case -598:
            fmt.Println(p0)
        default:
            switch -2147483648 {
            case -526:
                v8 := GenS1{F0: 100}
                _ = v8
            case -530:
                var v9 float64 = (-2.25 + -2.25)
                _ = v9
            default:
                v10 := []int{42}
                _ = v10
            }
        }
        v11 := GenS2{F0: 111, F1: 0, F2: -209}
        _ = v11
    }
    fmt.Println(p0)
    _ = p0
    return -32768
}

func main() {
    var v1 string = "hello"
    v2 := []uint32{65535, 1, 2147483647}
    v3 := []int{}
    var v4 float32 = float32((0.1 + -2.25))
    var v5 int64 = 42
    var v6 string = ("hello" + v1)
    switch v1 {
    case "hello":
        var v12 int64 = -1
        _ = v12
    case "":
        var v13 int32 = 0
        _ = v13
    default:
        v14 := GenS2{F0: 0, F1: 1, F2: 1000}
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 float64 = float64(1e10)
            fmt.Println(v4, v16, v16)
            var v17 string = v1
            _ = fi15
            _ = v16
            _ = v17
        }
        if ((-14 == -1) || ((575 != 4294967295) || (-1 <= 1))) {
            var v18 int = 0
            if false {
                v1 = v1
            }
            var v19 uint64 = 913
            _ = v18
            _ = v19
        }
    } else {
        _ = len(v2)
        switch "" {
        case "\n\t":
            var v20 uint16 = 32767
            _ = v20
        case "hello":
            var v21 int64 = v5
            _ = v21
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            v5 = -9223372036854775808
        }
    }
    switch 869 {
    case 4294967295:
        v2 = append(v2, 2)
    case 2147483647:
        v1 = (v6 + v6)
    }
    var v22 int32 = -2147483648
    var v23 string = v6
    v1 = "a"
    fmt.Println(v23, v5)
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 19; fi25++ {
            var v26 uint = 2
            v6 = v23
            _ = len(v2)
            v27 := GenS3{F0: 65535, F1: true}
            v5 += v5
            _ = fi25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 15; fi28++ {
            _ = len(v2)
            switch 797 {
            case 128:
                var v29 uint16 = 1
                _ = v29
            case -455:
                if true {
                    if (!(!(!(false && true)))) {
                        _ = len(v2)
                        v1 = v6
                    }
                    switch ((v1 + (("foo bar" + (v6 + "a")) + v23)) + "") {
                    case "":
                        switch (fi24 * (fi24 * fi24)) {
                        case 9223372036854775807:
                            v4 += -2.25
                        case -1:
                            v5 += (v5 & v5)
                        case 127:
                            v23 = v1
                        default:
                            fmt.Println(v23, v5, v5)
                        }
                    case "a":
                        if false {
                            var v30 int32 = (-(v22))
                            _ = v30
                        }
                    case "hello":
                        if true {
                            var v31 uint8 = 123
                            v2 = append(v2, 263)
                            _ = v31
                        }
                    }
                    switch (fi28 + fi24) {
                    case 3:
                        v3 = append(v3, (fi24 | (fi24 * (fi28 / (((fi24 * 32768)) | 1)))))
                    default:
                        if true {
                            var v32 int = 42
                            v33 := GenS1{F0: 100}
                            _ = v32
                            _ = v33
                        }
                    }
                } else {
                    switch fi24 {
                    case -9223372036854775808:
                        v22 = (v22 * v22)
                    case 2147483647:
                        var v34 uint = 18446744073709551615
                        _ = v34
                    case -330:
                        var v35 uint16 = 42
                        _ = v35
                    }
                }
                fmt.Println(v23, v1, fi24)
                v6 = v1
            }
            var v36 uint8 = 184
            switch -179 {
            case -1:
                var v37 int32 = v22
                _ = v37
            }
            _ = fi28
            _ = v36
        }
        for fi38 := 0; fi38 < 11; fi38++ {
            v22 = 42
            v22 -= v22
            _ = fi38
        }
        _ = fi24
    }
    switch (v23 + "a") {
    case "foo bar":
        var v39 uint32 = 623
        _ = v39
    case "":
        v40 := GenS3{F0: 2, F1: true}
        _ = v40
    }
    if true {
        v23 = (v1 + "\n\t")
    } else {
        v41 := GenS4{F0: 100, F1: v22}
        switch ((v1 + ("hello" + v23)) + v6) {
        case "hello":
            v42 := GenS4{F0: 121, F1: v22}
            _ = v42
        case "a":
            v43 := GenS4{F0: -94, F1: (v22 << ((2) & 63))}
            _ = v43
        case "\n\t":
            v3 = append(v3, 242)
        }
        if (-32768 < 32767) {
            var v44 int = 0
            v6 = (v1 + v23)
            if true {
                var v45 float64 = float64(0.1)
                for fi46 := 0; fi46 < 14; fi46++ {
                    v47 := GenS3{F0: 414, F1: (false || true)}
                    v1 = "\n\t"
                    _ = fi46
                    _ = v47
                }
                v22 = v22
                _ = v45
            } else {
                v22 += v22
                for fi48 := 0; fi48 < 9; fi48++ {
                    var v49 uint64 = 53
                    _ = fi48
                    _ = v49
                }
                var v50 int = -1
                _ = v50
            }
            _ = v44
        }
        _ = v41
    }
    v5 += v5
    var v51 string = "\n\t"
    fmt.Println(v5)
    v1 = v23
    for fi52 := 0; fi52 < 16; fi52++ {
        if ((65535 < 621) || (("" <= v23) && true)) {
            var v53 bool = (127 == 127)
            for fi54 := 0; fi54 < 14; fi54++ {
                var v55 uint32 = 2
                if (1 != -9223372036854775808) {
                    v56 := GenS1{F0: 100}
                    _ = v56
                }
                switch (("a" + (v6 + ((v23 + "\n\t") + (v51 + ((v51 + "\n\t") + (v51 + "hello")))))) + ((v6 + "hello") + ((v23 + "foo bar") + v1))) {
                case "foo bar":
                    v51 = v6
                case "a":
                    var v57 uint64 = 2
                    v55 -= v55
                    v58 := GenS5{F0: false}
                    _ = v57
                    _ = v58
                }
                _ = fi54
                _ = v55
            }
            var v59 uint8 = 36
            _ = v53
            _ = v59
        } else {
            v60 := GenS6{F0: 127, F1: fi52}
            v61 := GenS3{F0: 65535, F1: true}
            var v62 uint32 = 631
            _ = v60
            _ = v61
            _ = v62
        }
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
