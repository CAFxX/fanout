package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
func genFn0() bool {
    if (((false && true) && false) || (42 != 4294967295)) {
        switch "foo bar" {
        case "a":
            var v7 uint16 = 0
            _ = v7
        }
        for fi8 := 0; fi8 < 9; fi8++ {
            if (!(false && (((false || (!((true && true) || (127 >= (fi8 % ((fi8) | 1)))))) && false) || false))) {
                var v9 string = "a"
                _ = v9
            } else {
                switch (("a" + "a") + "a") {
                case "a":
                    var v10 uint = 4294967295
                    _ = v10
                case "foo bar":
                    var v11 int8 = -125
                    _ = v11
                }
                v12 := GenS1{F0: 2}
                fmt.Println(fi8)
                _ = v12
            }
            var v13 float64 = 3.14159
            var v14 uint8 = 42
            _ = fi8
            _ = v13
            _ = v14
        }
        var v15 uint8 = 1
        _ = v15
    } else {
        switch "hello" {
        case "a":
            v16 := []int{32768, -174}
            _ = v16
        case "\n\t":
            var v17 uint = 279
            _ = v17
        case "foo bar":
            if false {
                v18 := []int{}
                v19 := GenS1{F0: 127}
                _ = v18
                _ = v19
            }
        default:
            var v20 uint16 = 32767
            _ = v20
        }
        var v21 uint = 255
        _ = v21
    }
    var v22 int8 = 20
    if false {
        v22 -= (v22 & (v22 / ((v22) | 1)))
    } else {
        switch "" {
        case "foo bar":
            v22 -= 53
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            for fi24 := 0; fi24 < 2; fi24++ {
                switch "\n\t" {
                case "hello":
                    if (((42 <= 869) || true) || (false && (true || (false || (false && false))))) {
                        var v25 int32 = 2147483647
                        if (("foo bar" + (("a" + ("foo bar" + "")) + ("\n\t" + (("hello" + "a") + "")))) != "a") {
                            if true {
                                var v26 uint = 792
                                _ = v26
                            }
                        }
                        v27 := []int{(fi24 ^ 285)}
                        _ = v25
                        _ = v27
                    }
                case "\n\t":
                    v28 := GenS1{F0: 127}
                    if false {
                        var v29 bool = false
                        _ = v29
                    }
                    _ = v28
                default:
                    switch "" {
                    case "foo bar":
                        switch "hello" {
                        case "":
                            var v30 int16 = 1
                            _ = v30
                        case "a":
                            if (42 <= -1) {
                                var v31 bool = false
                                _ = v31
                            }
                        }
                    default:
                        if (1e10 != float64(1e10)) {
                            v32 := GenS2{F0: "foo bar", F1: 128, F2: 1}
                            var v33 float64 = 0.1
                            var v34 float32 = 0.0
                            _ = v32
                            _ = v33
                            _ = v34
                        } else {
                            v35 := GenS2{F0: ("\n\t" + "hello"), F1: 88, F2: 127}
                            _ = v35
                        }
                    }
                }
                _ = fi24
            }
            _ = fi23
        }
        for fi36 := 0; fi36 < 15; fi36++ {
            switch 128 {
            case 4294967295:
                v37 := GenS2{F0: "\n\t", F1: 42, F2: 77}
                _ = v37
            }
            _ = fi36
        }
    }
    fmt.Println(v22)
    switch -408 {
    case -322:
        switch "a" {
        case "foo bar":
            v38 := GenS3{F0: 1, F1: 814, F2: 1}
            _ = v38
        case "\n\t":
            var v39 uint64 = 391
            _ = v39
        case "hello":
            var v40 float32 = -2.25
            _ = v40
        default:
            if true {
                v41 := []int{}
                for fi42 := 0; fi42 < 1; fi42++ {
                    v43 := GenS4{F0: 4294967295}
                    _ = fi42
                    _ = v43
                }
                var v44 bool = (-1 != 1)
                _ = v41
                _ = v44
            }
        }
    default:
        v22 -= (-(v22))
    }
    _ = v22
    return true
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 string = ("\n\t" + (("" + "foo bar") + ""))
    var v5 int32 = 1
    var v6 uint16 = 65535
    switch -2147483648 {
    case -2147483648:
        v4 = (v4 + (v4 + ("foo bar" + "hello")))
    case 0:
        v4 = ("" + "a")
    }
    var v45 float64 = (1.5 / -2.25)
    if false {
        var v46 uint = 65535
        var v47 int8 = 1
        _ = v46
        _ = v47
    }
    v48 := GenS4{F0: 254}
    v49 := v2["foo bar"]
    v50 := GenS5{F0: 880}
    _ = len(v1)
    var v51 uint16 = 0
    var v52 uint = 2
    switch 32768 {
    case 128:
        v51 = 0
    case 32767:
        fmt.Println(v4, v49, v6)
    }
    for fi53 := 0; fi53 < 12; fi53++ {
        var v54 uint8 = 175
        _ = fi53
        _ = v54
    }
    var v55 int32 = 2147483647
    var v56 uint16 = v6
    var v57 uint = (v52 % ((18446744073709551615) | 1))
    switch -1 {
    case -847:
        switch 722 {
        case 4294967295:
            var v58 uint = v57
            _ = v58
        }
    case 127:
        var v59 int32 = (^(v5))
        _ = v59
    case 32768:
        switch 750 {
        case -317:
            switch int(v6) {
            case 832:
                if len(v1) > 0 {
                    v1[0] = v4
                }
            default:
                v52 += v57
            }
        case -1:
            var v60 bool = (true && (!v49))
            _ = v60
        case 32767:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            fmt.Println(v51)
        }
    default:
        var v61 uint64 = 324
        _ = v61
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v55
    _ = v56
    _ = v57
    _ = v6
}
