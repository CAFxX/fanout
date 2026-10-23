package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 int16, p1 uint16) int32 {
    if false {
        if true {
            var v4 float64 = 3.14159
            switch "foo bar" {
            case "a":
                var v5 uint32 = 4294967295
                _ = v5
            case "foo bar":
                for fi6 := 0; fi6 < 12; fi6++ {
                    var v7 int64 = -962
                    var v8 uint = 18446744073709551615
                    var v9 int = 42
                    _ = fi6
                    _ = v7
                    _ = v8
                    _ = v9
                }
            }
            _ = v4
        }
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            p0 -= -114
            if ((-2.25 * -2.25) != -2.25) {
                v12 := []int{fi11, 1}
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v12
            } else {
                fmt.Println(p1, fi11)
            }
            p1 = (p1 & p1)
            p0 -= (p0 % ((p0) | 1))
            _ = fi11
        }
        _ = fi10
    }
    var v13 uint32 = 2147483648
    switch (("hello" + (("" + "hello") + ("a" + ""))) + "") {
    case "foo bar":
        var v14 bool = (!(1 >= -726))
        _ = v14
    case "\n\t":
        var v15 int64 = -139
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    return 1
}

func main() {
    var v1 uint32 = 1
    var v2 float64 = 1.5
    v3 := []uint32{(^(v1)), (v1 << ((255) & 63))}
    v16 := GenS1{F0: 1, F1: 1, F2: (("" + "hello") + ("a" + "a"))}
    v17 := GenS2{F0: ("\n\t" + (("" + "") + "a"))}
    for fi18 := 0; fi18 < 20; fi18++ {
        _ = len(v3)
        switch 2147483647 {
        case -745:
            var v19 int = (fi18 % (((fi18 / ((fi18) | 1))) | 1))
            _ = v19
        case 4294967295:
            for fi20 := 0; fi20 < 2; fi20++ {
                switch "foo bar" {
                case "hello":
                    v2 = 3.14159
                case "a":
                    if (!true) {
                        fmt.Println(v1, fi18)
                        var v21 uint16 = 120
                        v22 := GenS2{F0: ("foo bar" + "a")}
                        _ = v21
                        _ = v22
                    }
                case "\n\t":
                    fmt.Println(fi18, fi20, v2)
                }
                switch ("hello" + "foo bar") {
                case "a":
                    if ((1.5 != v2) && (!true)) {
                        v2 = (v2 + -2.25)
                    } else {
                        v2 += v2
                    }
                }
                _ = fi20
            }
        case 42:
            switch 127 {
            case 207:
                switch "\n\t" {
                case "foo bar":
                    if false {
                        for fi23 := 0; fi23 < 6; fi23++ {
                            var v24 float64 = 1.5
                            var v25 int32 = genFn0(125, 0)
                            switch ("a" + "") {
                            case "a":
                                v1 = 2147483648
                            }
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    }
                case "\n\t":
                    fmt.Println(v1)
                case "a":
                    v26 := GenS1{F0: 18446744073709551615, F1: 1, F2: "a"}
                    _ = v26
                }
            }
        default:
            v3 = append(v3, (v1 ^ v1))
        }
        _ = fi18
    }
    fmt.Println(v2, v1)
    var v27 int16 = 409
    var v28 int32 = 0
    var v29 float64 = ((v2 - float64(-2.25)) / v2)
    if true {
        v3 = append(v3, (v1 * (^(v1))))
    } else {
        switch 4294967295 {
        case 32768:
            var v30 int64 = -1
            _ = v30
        case -496:
            var v31 bool = true
            _ = v31
        case -2147483648:
            var v32 int = 186
            _ = v32
        }
        v33 := GenS3{F0: 57, F1: true}
        var v34 int64 = 9223372036854775807
        _ = v33
        _ = v34
    }
    if false {
        for fi35 := 0; fi35 < 20; fi35++ {
            if len(v3) > 0 {
                v3[0] = v1
            }
            fmt.Println(v27, v2)
            v36 := GenS4{F0: 42, F1: 65535, F2: (v28 + v28)}
            _ = fi35
            _ = v36
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v29 -= v2
    }
    fmt.Println(v1, v2, v29)
    fmt.Println(v27, v27, v27)
    fmt.Println(v27, v28)
    v27 -= (v27 & (v27 / (((v27 >> ((255) & 63))) | 1)))
    v27 = v27
    fmt.Println(v2, v1)
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 uint = 1
        var v39 float64 = 0.1
        v40 := GenS4{F0: 255, F1: (v38 | 2), F2: 0}
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
