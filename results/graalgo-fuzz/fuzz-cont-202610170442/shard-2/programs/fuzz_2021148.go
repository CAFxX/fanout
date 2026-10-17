package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint8; F2 int }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 int64; F2 int32 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 int32; F1 int }

func main() {
    var v1 uint16 = 552
    var v2 int = 706
    var v3 uint = 4294967295
    var v4 uint64 = 4294967295
    for fi5 := 0; fi5 < 8; fi5++ {
        if (1 != int64(v3)) {
            v6 := GenS1{F0: uint32(v4)}
            switch "hello" {
            case "foo bar":
                if (42 > 0) {
                    v4 = (v4 % (((v4 ^ (v4 / ((v4) | 1)))) | 1))
                    fmt.Println(v1, v2, fi5)
                    var v7 uint64 = (v4 << ((v3) & 63))
                    _ = v7
                }
            default:
                var v8 uint = 0
                _ = v8
            }
            _ = v6
        } else {
            if true {
                for fi9 := 0; fi9 < 18; fi9++ {
                    v10 := GenS2{F0: -17, F1: 227, F2: fi9}
                    fmt.Println(v2)
                    _ = fi9
                    _ = v10
                }
                var v11 string = ("" + "a")
                if ((^(v1)) == 0) {
                    var v12 string = v11
                    _ = v12
                } else {
                    var v13 string = v11
                    var v14 string = (v13 + "\n\t")
                    switch (((v11 + ((("\n\t" + "") + v14) + "a")) + "a") + (v11 + ((v11 + v13) + "hello"))) {
                    case "\n\t":
                        fmt.Println(v2, v4, v13)
                    case "hello":
                        for fi15 := 0; fi15 < 2; fi15++ {
                            fmt.Println(v14)
                            v16 := GenS1{F0: 1}
                            _ = fi15
                            _ = v16
                        }
                    case "a":
                        switch (("foo bar" + v13) + "hello") {
                        case "":
                            for fi17 := 0; fi17 < 2; fi17++ {
                                var v18 int16 = 449
                                _ = fi17
                                _ = v18
                            }
                        case "\n\t":
                            for fi19 := 0; fi19 < 17; fi19++ {
                                v20 := GenS2{F0: 704, F1: 42, F2: (fi5 << ((v3) & 63))}
                                v2 = v2
                                _ = fi19
                                _ = v20
                            }
                        default:
                            var v21 uint = (v3 / (((v3 * (v3 % (((^(v3))) | 1)))) | 1))
                            _ = v21
                        }
                    }
                    _ = v13
                    _ = v14
                }
                _ = v11
            }
        }
        v1 = v1
        var v22 uint = v3
        _ = fi5
        _ = v22
    }
    var v23 int32 = 42
    if false {
        fmt.Println(v2, v23, v2)
        if (42 != -1) {
            v24 := []int{(v2 * v2), v2}
            v25 := GenS3{F0: 2, F1: 127}
            _ = v24
            _ = v25
        } else {
            v26 := []int{(v2 & (v2 ^ (v2 - (^(v2))))), v2}
            _ = v26
        }
    }
    var v27 int = (v2 | v2)
    if (true && false) {
        v3 -= 831
        v28 := []int{(v27 * (^(v27))), v2, 127}
        _ = v28
    }
    var v29 float32 = 3.14159
    v30 := []int{v2, v27, (v27 ^ 815)}
    switch "a" {
    case "\n\t":
        v23 = v23
    case "hello":
        v1 += (^(v1))
    }
    v23 += 715
    v31 := GenS2{F0: (v23 | (v23 / ((-2147483648) | 1))), F1: 94, F2: (-(v2))}
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    switch 4294967295 {
    case 128:
        v4 = (v4 | v4)
    case 127:
        if false {
            v32 := GenS4{F0: uint32(v4)}
            switch "a" {
            case "hello":
                var v33 float32 = 1e10
                _ = v33
            case "\n\t":
                v3 -= v3
            case "a":
                var v34 uint = v3
                _ = v34
            }
            _ = v32
        } else {
            var v35 uint = v3
            for fi36 := 0; fi36 < 4; fi36++ {
                var v37 int64 = -1
                for fi38 := 0; fi38 < 8; fi38++ {
                    switch (fi36 % (((v2 << (((v3 << ((v3) & 63))) & 63))) | 1)) {
                    case 153:
                        if (!true) {
                            var v39 string = "foo bar"
                            var v40 string = ("a" + ("\n\t" + "\n\t"))
                            switch "a" {
                            case "foo bar":
                                fmt.Println(v35)
                            default:
                                var v41 bool = true
                                _ = v41
                            }
                            _ = v39
                            _ = v40
                        }
                        var v42 int32 = v23
                        _ = v42
                    case 0:
                        v43 := GenS5{F0: 65535, F1: 42, F2: (v23 & (v23 - (v23 ^ (v23 + 2147483647))))}
                        _ = v43
                    default:
                        if true {
                            v35 -= (^(v3))
                            fmt.Println(v4, v29)
                        }
                    }
                    var v44 int = v27
                    _ = fi38
                    _ = v44
                }
                var v45 uint32 = uint32(v4)
                _ = fi36
                _ = v37
                _ = v45
            }
            _ = v35
        }
    case 9223372036854775807:
        v46 := GenS6{F0: (v4 + (v4 % ((18446744073709551615) | 1))), F1: 728}
        _ = v46
    default:
        for fi47 := 0; fi47 < 3; fi47++ {
            v30 = append(v30, fi47)
            _ = fi47
        }
    }
    v48 := GenS7{F0: (v23 ^ (^(v23))), F1: (v2 * -277)}
    var v49 float32 = (0.0 * (3.14159 - (v29 / (v29 - v29))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v48
    _ = v49
}
