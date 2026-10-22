package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 string; F1 int; F2 int64 }
type GenS4 struct { F0 int; F1 int32 }
func genFn0(p0 int64) uint16 {
    switch (("" + "foo bar") + "foo bar") {
    case "":
        fmt.Println(p0)
    case "foo bar":
        p0 += p0
    default:
        var v4 float32 = float32(0.0)
        _ = v4
    }
    var v5 uint64 = 2
    _ = p0
    _ = v5
    return 1
}

func main() {
    v1 := []uint32{65535}
    v2 := []int{128, 9223372036854775807}
    var v3 string = ("a" + "\n\t")
    v3 = "foo bar"
    var v6 int16 = -32768
    var v7 uint = 978
    if true {
        var v8 int16 = (v6 * v6)
        v8 = v8
        _ = v8
    } else {
        if (0.1 == ((3.14159 + 1e10) / 1e10)) {
            var v9 float64 = 3.14159
            _ = v9
        }
    }
    fmt.Println(v3)
    v10 := GenS1{F0: 42}
    if len(v2) > 0 {
        v2[0] = 32767
    }
    if (-2147483648 >= -694) {
        switch -316 {
        case 455:
            var v11 int8 = 100
            _ = v11
        case 425:
            if false {
                v3 = v3
            }
        default:
            switch 0 {
            case 9223372036854775807:
                v7 = 424
            case -9223372036854775808:
                var v12 uint32 = 1
                _ = v12
            case 143:
                v7 = 4294967295
            }
        }
    } else {
        v3 = (("hello" + v3) + v3)
        fmt.Println(v7, v6)
        var v13 float32 = 0.1
        _ = v13
    }
    v7 = (v7 - 0)
    if false {
        var v14 uint16 = 65535
        var v15 uint8 = 1
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS1{F0: (fi16 / ((fi16) | 1))}
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    } else {
        var v18 uint8 = 144
        for fi19 := 0; fi19 < 2; fi19++ {
            v18 += (v18 / (((v18 | (v18 | (v18 >> ((255) & 63))))) | 1))
            _ = fi19
        }
        if ((!false) && (((v3 <= v3) || true) || false)) {
            switch v3 {
            case "":
                var v20 string = ""
                _ = v20
            case "foo bar":
                v6 += (-(v6))
            }
            for fi21 := 0; fi21 < 15; fi21++ {
                if len(v2) > 0 {
                    v2[0] = int(v18)
                }
                v22 := GenS1{F0: -573}
                var v23 uint16 = 102
                _ = fi21
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 uint = v7
                _ = fi24
                _ = v25
            }
        }
        _ = v18
    }
    switch v3 {
    case "a":
        for fi26 := 0; fi26 < 19; fi26++ {
            for fi27 := 0; fi27 < 19; fi27++ {
                switch (fi26 ^ (fi26 / ((4294967295) | 1))) {
                case 127:
                    switch -516 {
                    case 0:
                        if len(v2) > 0 {
                            v2[0] = 224
                        }
                    case -497:
                        switch v3 {
                        case "a":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            var v28 int32 = 2147483647
                            _ = v28
                        }
                    case -2147483648:
                        if true {
                            v29 := GenS1{F0: -223}
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                v30 := GenS2{F0: 9223372036854775807, F1: 2}
                                var v31 string = v3
                                _ = v30
                                _ = v31
                            }
                            v32 := GenS1{F0: fi27}
                            if true {
                                var v33 float64 = 0.1
                                var v34 int16 = v6
                                _ = v33
                                _ = v34
                            }
                            var v35 uint16 = 906
                            _ = v29
                            _ = v32
                            _ = v35
                        } else {
                            fmt.Println(v3, fi27)
                            fmt.Println(fi26, v6)
                        }
                    default:
                        switch v3 {
                        case "hello":
                            if false {
                                fmt.Println(fi27, v6)
                                v2 = append(v2, (fi27 + (fi26 | (fi27 << ((v7) & 63)))))
                            }
                        case "foo bar":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                            var v36 int16 = v6
                            _ = v36
                        case "a":
                            v6 += (v6 << ((0) & 63))
                            switch (v3 + "\n\t") {
                            case "hello":
                                v3 = ""
                            case "foo bar":
                                v7 += 0
                            default:
                                v37 := GenS3{F0: (("" + v3) + ("foo bar" + v3)), F1: (fi26 ^ 32768), F2: 1}
                                _ = v37
                            }
                            v2 = append(v2, (fi26 & 42))
                        }
                    }
                case 4294967295:
                    v7 = 42
                case 32767:
                    var v38 int8 = -128
                    _ = v38
                }
                if true {
                    if ((1 <= -67) || (((!((!true) && true)) || (0 <= 42)) || true)) {
                        _ = len(v2)
                        if false {
                            if (!(!(false && false))) {
                                v7 = (v7 | 18446744073709551615)
                                if len(v1) > 0 {
                                    v1[0] = 2147483648
                                }
                            } else {
                                var v39 bool = (true || false)
                                _ = v39
                            }
                            v6 -= v6
                            if (fi26 != fi26) {
                                if true {
                                    v6 += (v6 * 32767)
                                    fmt.Println(fi26, fi26)
                                } else {
                                    fmt.Println(fi27)
                                }
                            }
                        }
                        v7 += 945
                    } else {
                        fmt.Println(v6, fi26)
                    }
                    var v40 string = ((v3 + (v3 + (v3 + (v3 + (v3 + v3))))) + v3)
                    var v41 uint8 = 1
                    _ = v40
                    _ = v41
                }
                _ = len(v1)
                _ = fi27
            }
            _ = fi26
        }
    case "foo bar":
        var v42 float32 = (1.5 * -2.25)
        _ = v42
    case "\n\t":
        var v43 uint = 65535
        _ = v43
    }
    if false {
        fmt.Println(v3, v3)
        var v44 int64 = 1
        _ = v44
    } else {
        v7 += v7
    }
    for fi45 := 0; fi45 < 10; fi45++ {
        if true {
            var v46 float32 = 0.0
            v47 := GenS4{F0: (fi45 >> (((v7 % (((v7 >> ((255) & 63))) | 1))) & 63)), F1: 42}
            _ = v46
            _ = v47
        } else {
            var v48 bool = ((fi45 & fi45) <= (fi45 - 4294967295))
            fmt.Println(v48, v7)
            var v49 int64 = 0
            _ = v48
            _ = v49
        }
        _ = fi45
    }
    v3 = ("foo bar" + v3)
    var v50 uint = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v50
    _ = v6
    _ = v7
}
