package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint8; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }
func genFn0() uint8 {
    switch 0 {
    case -112:
        var v5 int16 = 447
        _ = v5
    default:
        v6 := []int{0, 4294967295}
        _ = v6
    }
    var v7 uint32 = 65535
    var v8 uint32 = 2147483648
    v9 := GenS1{F0: 0, F1: false}
    var v10 float64 = ((-2.25 - 1e10) * ((-2.25 * ((float64(0.1) - 0.0) * 1.5)) + 1e10))
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 183
}
func genFn1(p0 uint, p1 uint32) int {
    switch "hello" {
    case "a":
        p1 = (^(p1))
    case "":
        p1 -= p1
    case "foo bar":
        var v11 uint = p0
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS1{F0: 322, F1: false}
        fmt.Println(p1, p1)
        _ = fi12
        _ = v13
    }
    if (p1 == (^(p1))) {
        if true {
            p1 -= 1
        }
    } else {
        if (32767 <= -509) {
            var v14 uint = (p0 & (p0 ^ p0))
            for fi15 := 0; fi15 < 4; fi15++ {
                p1 = (p1 << ((p0) & 63))
                var v16 float64 = 1.5
                var v17 uint64 = 42
                _ = fi15
                _ = v16
                _ = v17
            }
            switch "foo bar" {
            case "a":
                var v18 string = "a"
                _ = v18
            case "":
                v19 := []int{128, -1, 32767}
                _ = v19
            case "\n\t":
                var v20 bool = (!(18446744073709551615 <= (^(v14))))
                _ = v20
            default:
                var v21 uint16 = 1
                _ = v21
            }
            _ = v14
        }
        v22 := GenS2{F0: 0}
        v23 := GenS3{F0: "hello"}
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 float32 = 0.0
    var v2 int16 = 784
    var v3 int = -9223372036854775808
    var v4 float64 = (float64((v1 / v1)) * 0.0)
    v1 = 0.0
    var v24 int = (v3 & (v3 % ((genFn1(0, 4294967295)) | 1)))
    v25 := []int{(v24 | genFn1(65535, 42)), (v24 ^ v3), 895, (v3 << ((255) & 63))}
    var v26 int8 = 100
    v27 := GenS1{F0: 18446744073709551615, F1: ((1 != 1) || true)}
    for fi28 := 0; fi28 < 1; fi28++ {
        for fi29 := 0; fi29 < 1; fi29++ {
            v26 -= v26
            v30 := GenS4{F0: 0, F1: 232, F2: "a"}
            _ = fi29
            _ = v30
        }
        v3 = (fi28 & v24)
        _ = fi28
    }
    var v31 bool = (!(false && true))
    if (!v31) {
        var v32 uint8 = 0
        var v33 uint64 = 0
        if (v33 >= v33) {
            var v34 uint = 404
            _ = v34
        } else {
            if (("" + "\n\t") == ("a" + "a")) {
                fmt.Println(v4)
                switch v3 {
                case -1:
                    v31 = (v31 && true)
                case 42:
                    v35 := GenS3{F0: "hello"}
                    _ = v35
                case 2147483647:
                    if (v31 || (!false)) {
                        var v36 bool = false
                        switch (v3 ^ (v3 | v24)) {
                        case 2147483647:
                            switch (v24 * -972) {
                            case 794:
                                var v37 uint16 = 42
                                _ = v37
                            case 843:
                                var v38 bool = true
                                _ = v38
                            case 9223372036854775807:
                                var v39 int32 = 42
                                _ = v39
                            }
                        case 128:
                            if len(v25) > 0 {
                                v25[0] = (v24 | v3)
                            }
                        default:
                            var v40 int16 = (^(v2))
                            _ = v40
                        }
                        _ = v36
                    } else {
                        var v41 float64 = v4
                        _ = len(v25)
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        _ = v41
                    }
                default:
                    var v42 uint64 = v33
                    _ = v42
                }
                v3 = (-(v3))
            } else {
                v4 = 1.5
            }
            var v43 bool = v31
            fmt.Println(v32)
            _ = v43
        }
        _ = v32
        _ = v33
    } else {
        var v44 uint = 582
        if (!false) {
            var v45 int32 = 1
            v26 -= v26
            switch v3 {
            case 42:
                var v46 float32 = (v1 * v1)
                _ = v46
            case 118:
                var v47 uint64 = 18446744073709551615
                _ = v47
            case -953:
                var v48 int8 = int8(v44)
                _ = v48
            default:
                var v49 uint8 = 42
                _ = v49
            }
            _ = v45
        }
        if (v4 <= 0.1) {
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            switch ("" + ((("" + "\n\t") + "a") + "hello")) {
            case "hello":
                v50 := GenS2{F0: -9223372036854775808}
                _ = v50
            case "\n\t":
                var v51 float64 = float64((1e10 * 1e10))
                _ = v51
            case "foo bar":
                var v52 uint64 = 0
                _ = v52
            }
            for fi53 := 0; fi53 < 6; fi53++ {
                var v54 uint64 = 18446744073709551615
                v26 = (v26 / (((-(v26))) | 1))
                if true {
                    v55 := GenS5{F0: -167}
                    for fi56 := 0; fi56 < 15; fi56++ {
                        switch "a" {
                        case "foo bar":
                            var v57 uint8 = 38
                            _ = v57
                        case "":
                            switch genFn1(v44, 2147483647) {
                            case -1:
                                v58 := GenS6{F0: (v54 | (v54 >> (((^(v44))) & 63)))}
                                _ = v58
                            }
                        case "hello":
                            if (42 >= 65535) {
                                v54 += 59
                            }
                        }
                        _ = fi56
                    }
                    var v59 float64 = 0.1
                    _ = v55
                    _ = v59
                }
                _ = fi53
                _ = v54
            }
        }
        _ = v44
    }
    switch ("\n\t" + "foo bar") {
    case "a":
        v31 = ((-77 > -105) || (0 != -1))
    case "foo bar":
        var v60 int8 = v26
        _ = v60
    case "\n\t":
        v61 := GenS1{F0: 0, F1: ((v2 >> ((464) & 63)) == v2)}
        _ = v61
    }
    var v62 uint32 = 42
    var v63 uint = 601
    for fi64 := 0; fi64 < 13; fi64++ {
        switch (v24 * 128) {
        case 0:
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
        case 32768:
            v62 += (^(v62))
        case -341:
            fmt.Println(v62)
        }
        _ = fi64
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v62
    _ = v63
}
