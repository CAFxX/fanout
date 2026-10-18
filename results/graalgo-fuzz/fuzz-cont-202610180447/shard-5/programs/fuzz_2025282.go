package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint8; F2 string }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 int8 }

func main() {
    var v1 int8 = -128
    var v2 int16 = -1
    var v3 int32 = 363
    switch 0 {
    case -617:
        if false {
            fmt.Println(v3)
            v4 := []int{}
            _ = v4
        }
    case -1:
        fmt.Println(v1, v1, v1)
    default:
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 bool = ((float32(1e10) > 3.14159) && false)
            v6 = (255 > 0)
            switch ("hello" + "") {
            case "foo bar":
                var v7 uint64 = 18446744073709551615
                _ = v7
            case "a":
                v8 := []int{}
                _ = v8
            case "":
                var v9 string = "foo bar"
                _ = v9
            default:
                v1 = -1
            }
            _ = fi5
            _ = v6
        }
    }
    if (("a" + "a") >= ("" + "\n\t")) {
        v2 -= v2
        var v10 uint = 251
        for fi11 := 0; fi11 < 19; fi11++ {
            if true {
                v12 := []int{fi11, -938, -977}
                var v13 uint8 = 238
                fmt.Println(v1, v2, v3)
                _ = v12
                _ = v13
            } else {
                fmt.Println(v3, v3)
                var v14 bool = (false && ((863 >= 373) && true))
                v3 = (^(v3))
                _ = v14
            }
            v3 = 85
            fmt.Println(fi11)
            _ = fi11
        }
        _ = v10
    } else {
        switch 127 {
        case 32767:
            v15 := []int{-128}
            _ = v15
        case 447:
            v16 := GenS1{F0: "hello"}
            _ = v16
        }
        v3 = 0
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        switch ((((("" + "\n\t") + "foo bar") + ((("a" + "hello") + (("hello" + ((("hello" + "foo bar") + ("foo bar" + "hello")) + "a")) + "")) + ("hello" + (("foo bar" + (("a" + "\n\t") + ("hello" + (((("\n\t" + (("a" + "foo bar") + "a")) + ("hello" + "\n\t")) + "a") + "")))) + "foo bar")))) + "") + "hello") {
        case "":
            for fi18 := 0; fi18 < 7; fi18++ {
                v2 -= 1000
                if (4294967295 != 2) {
                    var v19 bool = ((v1 / (((^(v1))) | 1)) <= (v1 + (v1 + (v1 * -128))))
                    var v20 float64 = ((1e10 / 0.1) + 1e10)
                    switch ("\n\t" + ("" + "foo bar")) {
                    case "a":
                        v19 = (v19 || (602 == 65535))
                    case "foo bar":
                        switch -178 {
                        case 441:
                            if v19 {
                                v1 -= (v1 ^ v1)
                                v3 = 0
                                var v21 int32 = v3
                                v1 = (^(v1))
                                _ = v21
                            }
                        default:
                            if v19 {
                                v20 = 3.14159
                            } else {
                                var v22 float64 = (v20 + v20)
                                switch "hello" {
                                case "\n\t":
                                    v19 = v19
                                case "a":
                                    v23 := GenS1{F0: "a"}
                                    _ = v23
                                default:
                                    var v24 uint8 = 255
                                    _ = v24
                                }
                                if v19 {
                                    var v25 string = "\n\t"
                                    v26 := []int{fi18}
                                    _ = v25
                                    _ = v26
                                }
                                _ = v22
                            }
                        }
                    case "\n\t":
                        var v27 string = ("a" + "foo bar")
                        var v28 uint16 = 3
                        _ = v27
                        _ = v28
                    }
                    _ = v19
                    _ = v20
                } else {
                    v2 += v2
                    v2 = (v2 >> ((912) & 63))
                    v29 := []int{fi18, 32768, (^(fi18))}
                    _ = v29
                }
                _ = fi18
            }
        }
        var v30 int64 = 42
        _ = fi17
        _ = v30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        if (3.14159 != -2.25) {
            var v32 string = ("hello" + "a")
            fmt.Println(v1)
            var v33 uint32 = 1
            _ = v32
            _ = v33
        } else {
            if (v3 == (v3 - v3)) {
                v34 := GenS2{F0: 760, F1: 223, F2: "a"}
                _ = v34
            }
            var v35 bool = (!false)
            v2 = v2
            _ = v35
        }
        if true {
            var v36 string = "\n\t"
            if (1 >= 0) {
                switch (fi31 & (fi31 / ((678) | 1))) {
                case -696:
                    v37 := []int{fi31, (fi31 % ((1) | 1))}
                    _ = v37
                case 128:
                    v38 := []int{(-(fi31)), fi31, fi31}
                    _ = v38
                case 2147483647:
                    for fi39 := 0; fi39 < 2; fi39++ {
                        v2 = v2
                        _ = fi39
                    }
                }
                v40 := GenS3{F0: 56, F1: (fi31 % ((fi31) | 1))}
                _ = v40
            }
            _ = v36
        } else {
            v41 := GenS4{F0: ("\n\t" + ("\n\t" + "foo bar"))}
            _ = v41
        }
        _ = fi31
    }
    fmt.Println(v2, v2)
    switch "foo bar" {
    case "":
        v42 := GenS5{F0: 1}
        _ = v42
    case "hello":
        switch "\n\t" {
        case "":
            if false {
                v2 = -1
                if (("" + "hello") <= "foo bar") {
                    var v43 int64 = -173
                    fmt.Println(v43, v1, v2)
                    _ = v43
                } else {
                    var v44 string = ("" + "a")
                    if (1 == 728) {
                        v3 -= v3
                    }
                    _ = v44
                }
                v45 := GenS6{F0: 1, F1: 127}
                _ = v45
            }
        case "hello":
            v46 := []int{-2147483648, -41, 0}
            _ = v46
        default:
            v2 = (v2 >> ((18446744073709551615) & 63))
        }
    default:
        fmt.Println(v3, v3)
    }
    for fi47 := 0; fi47 < 18; fi47++ {
        v1 += 100
        _ = fi47
    }
    var v48 int8 = v1
    var v49 int8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
}
