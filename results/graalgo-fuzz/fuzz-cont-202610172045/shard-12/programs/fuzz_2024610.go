package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 bool; F2 uint }
type GenS3 struct { F0 int64; F1 uint64; F2 uint8 }

func main() {
    var v1 uint16 = 846
    var v2 string = "foo bar"
    var v3 int = -428
    var v4 string = ("hello" + (v2 + ("foo bar" + ((v2 + ("a" + ("foo bar" + "a"))) + v2))))
    var v5 uint = 4294967295
    if false {
        fmt.Println(v4)
        var v6 int16 = 1000
        _ = v6
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            if (v5 < v5) {
                var v8 uint64 = 18446744073709551615
                v8 = v8
                _ = v8
            } else {
                var v9 uint8 = 42
                if (uint(v1) <= (v5 / ((v5) | 1))) {
                    if (!(957 >= -2147483648)) {
                        fmt.Println(v2, v1)
                    }
                    var v10 string = (v4 + "foo bar")
                    if false {
                        if true {
                            v11 := GenS1{F0: 2}
                            for fi12 := 0; fi12 < 15; fi12++ {
                                v9 = (v9 | 2)
                                _ = fi12
                            }
                            _ = v11
                        }
                    } else {
                        v4 = v2
                        v13 := GenS1{F0: 2147483648}
                        switch (v3 - (v3 & (fi7 >> (((^(v5))) & 63)))) {
                        case 128:
                            var v14 string = v2
                            _ = v14
                        case -1:
                            var v15 bool = (v1 < v1)
                            _ = v15
                        default:
                            if false {
                                if false {
                                    v16 := []int{v3, (fi7 | (fi7 & 32768)), 32767, 480}
                                    _ = v16
                                } else {
                                    v4 = ""
                                    for fi17 := 0; fi17 < 19; fi17++ {
                                        v3 = 0
                                        _ = fi17
                                    }
                                    v18 := GenS1{F0: 42}
                                    _ = v18
                                }
                            } else {
                                var v19 uint16 = (v1 & (^(v1)))
                                _ = v19
                            }
                        }
                        _ = v13
                    }
                    _ = v10
                } else {
                    if true {
                        switch "hello" {
                        case "\n\t":
                            v2 = v4
                        case "a":
                            var v20 float64 = (float64(0.0) - ((float64((0.0 + float32(3.14159))) * 3.14159) * 0.0))
                            _ = v20
                        case "":
                            switch (("" + v4) + (v2 + ("hello" + "hello"))) {
                            case "foo bar":
                                var v21 int16 = 0
                                _ = v21
                            case "\n\t":
                                v9 -= (v9 ^ (^(v9)))
                            }
                        default:
                            v22 := []int{2147483647, int(v9), 2147483647}
                            _ = v22
                        }
                        v5 = (^(v5))
                        var v23 uint16 = v1
                        _ = v23
                    }
                }
                _ = v9
            }
            v4 = v4
            _ = fi7
        }
        v24 := GenS1{F0: 2147483647}
        if ((v4 + ("" + v4)) > ("a" + (v4 + v2))) {
            switch 2147483647 {
            case 127:
                var v25 int16 = -32768
                _ = v25
            case -9223372036854775808:
                var v26 int16 = 251
                _ = v26
            }
            v4 = (v2 + v4)
            var v27 int16 = -1
            _ = v27
        }
        _ = v24
    }
    v28 := []int{v3, -390}
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 uint8 = 1
        if ((v30 - 1) == (v30 << ((v5) & 63))) {
            if true {
                fmt.Println(v30, v2)
                for fi31 := 0; fi31 < 15; fi31++ {
                    v32 := GenS1{F0: 2147483648}
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                    _ = fi31
                    _ = v32
                }
                v33 := GenS1{F0: 1}
                _ = v33
            } else {
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                var v34 uint = v5
                _ = v34
            }
            if false {
                var v35 int16 = -32768
                switch (v4 + (v2 + (v2 + "foo bar"))) {
                case "a":
                    if (false && (false && false)) {
                        for fi36 := 0; fi36 < 12; fi36++ {
                            if true {
                                var v37 int8 = -128
                                _ = v37
                            } else {
                                if true {
                                    fmt.Println(v2)
                                    var v38 uint16 = v1
                                    _ = v38
                                }
                                switch ("a" + v4) {
                                case "foo bar":
                                    fmt.Println(v5, v5, v3)
                                case "a":
                                    v4 = (("" + (v4 + v2)) + v2)
                                default:
                                    var v39 uint8 = v30
                                    _ = v39
                                }
                                v40 := GenS2{F0: 18446744073709551615, F1: ((v1 * 32767) > 1), F2: (v5 >> ((uint(v30)) & 63))}
                                _ = v40
                            }
                            v41 := GenS1{F0: 351}
                            v42 := GenS2{F0: 0, F1: (!true), F2: 2}
                            _ = len(v28)
                            _ = fi36
                            _ = v41
                            _ = v42
                        }
                        var v43 uint16 = (v1 / ((v1) | 1))
                        _ = v43
                    }
                case "":
                    var v44 string = v2
                    _ = v44
                default:
                    if len(v28) > 0 {
                        v28[0] = -9223372036854775808
                    }
                }
                _ = v35
            }
        }
        _ = fi29
        _ = v30
    }
    v45 := GenS1{F0: 2147483648}
    v46 := GenS2{F0: 2, F1: (-2147483648 == 2147483647), F2: (v5 - v5)}
    fmt.Println(v1)
    v47 := GenS3{F0: -1, F1: 788, F2: 114}
    for fi48 := 0; fi48 < 8; fi48++ {
        switch int(v5) {
        case 4294967295:
            fmt.Println(v3, fi48)
        }
        v49 := GenS1{F0: 42}
        _ = fi48
        _ = v49
    }
    v2 = "a"
    v4 = v2
    fmt.Println(v1, v3)
    var v50 string = v2
    switch ((v50 + v2) + v2) {
    case "\n\t":
        v50 = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v50
}
