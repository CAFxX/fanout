package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int32; F1 int32 }
func genFn0(p0 int) string {
    v5 := []int{(p0 & (p0 | 0)), p0, p0}
    var v6 bool = (!(true && (0 > 42)))
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint64 = 18446744073709551615
        v9 := GenS1{F0: 42}
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := GenS1{F0: 127}
            var v12 bool = (42 < 2147483647)
            if false {
                var v13 string = "hello"
                var v14 string = v13
                switch (-(fi7)) {
                case 32767:
                    _ = len(v5)
                case 32768:
                    if v12 {
                        switch v14 {
                        case "a":
                            var v15 bool = (-77 == -128)
                            _ = v15
                        case "\n\t":
                            var v16 int = fi10
                            v17 := GenS2{F0: v12}
                            switch fi7 {
                            case 760:
                                if (!v6) {
                                    fmt.Println(v14, fi7)
                                }
                            default:
                                v18 := GenS2{F0: false}
                                _ = v18
                            }
                            _ = v16
                            _ = v17
                        }
                        var v19 float64 = float64(3.14159)
                        _ = v19
                    }
                case -301:
                    fmt.Println(v6)
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        switch p0 {
        case 0:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case 32767:
            var v21 bool = ((!false) || (!(4294967295 == 4294967295)))
            _ = v21
        }
        var v22 float64 = 1e10
        v23 := GenS3{F0: -9223372036854775808, F1: (fi20 - -299)}
        _ = fi20
        _ = v22
        _ = v23
    }
    switch ((("foo bar" + ((((("" + (("foo bar" + "\n\t") + "\n\t")) + "hello") + "hello") + "a") + (("a" + (("\n\t" + "a") + ("" + (("\n\t" + "\n\t") + "")))) + ("\n\t" + "foo bar")))) + "a") + ("hello" + ((("hello" + "\n\t") + "a") + "\n\t"))) {
    case "\n\t":
        var v24 string = ("\n\t" + "")
        _ = v24
    case "hello":
        v25 := GenS1{F0: 0}
        _ = v25
    default:
        p0 = 32767
    }
    _ = p0
    _ = v5
    _ = v6
    return ""
}
func genFn1() int32 {
    v26 := GenS3{F0: 42, F1: 32768}
    v27 := GenS1{F0: 128}
    var v28 int32 = 42
    _ = v26
    _ = v27
    _ = v28
    return v28
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 1
    v3 := make(map[int]bool)
    v4 := []string{("hello" + ("a" + "\n\t")), ("foo bar" + "a"), "\n\t", "hello"}
    v3[v2] = (42 < 32767)
    switch v2 {
    case 0:
        v2 = (v2 ^ v2)
    case 128:
        v2 -= v2
    case -749:
        v29 := GenS3{F0: 0, F1: (v2 * v2)}
        _ = v29
    default:
        v30 := GenS3{F0: 650, F1: 9223372036854775807}
        _ = v30
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 uint16 = 0
        v32 = v32
        var v33 float32 = 0.0
        _ = fi31
        _ = v32
        _ = v33
    }
    if true {
        fmt.Println(v2)
        if len(v4) > 0 {
            v4[0] = "a"
        }
    } else {
        v2 = v2
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    switch "hello" {
    case "":
        if (95 != -57) {
            for fi34 := 0; fi34 < 10; fi34++ {
                v35 := GenS1{F0: 86}
                if (0.1 > float32(0.1)) {
                    switch "hello" {
                    case "hello":
                        var v36 uint64 = 2
                        _ = v36
                    case "\n\t":
                        v37 := GenS1{F0: 2}
                        _ = v37
                    case "foo bar":
                        v38 := GenS1{F0: 0}
                        _ = v38
                    }
                }
                _ = fi34
                _ = v35
            }
            v2 -= (v2 | v2)
            switch "hello" {
            case "hello":
                v2 -= v2
            }
        }
    case "\n\t":
        var v39 int16 = 1000
        _ = v39
    case "a":
        fmt.Println(v2)
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        if (-2.25 == -2.25) {
            var v41 uint32 = 597
            fmt.Println(v2, v41)
            if (3.14159 >= 0.0) {
                for fi42 := 0; fi42 < 7; fi42++ {
                    if false {
                        v2 = v2
                        v41 += 907
                    } else {
                        var v43 int16 = 664
                        v2 += fi40
                        switch ("hello" + ("a" + "")) {
                        case "\n\t":
                            var v44 string = "hello"
                            _ = v44
                        case "":
                            v45 := GenS1{F0: 1}
                            _ = v45
                        default:
                            v41 = (^(v41))
                        }
                        _ = v43
                    }
                    v2 -= fi40
                    _ = len(v1)
                    _ = fi42
                }
                v41 = (v41 >> ((255) & 63))
            }
            _ = v41
        }
        v46 := GenS4{F0: 2147483647, F1: -1}
        switch fi40 {
        case -664:
            switch ("\n\t" + ("\n\t" + ("foo bar" + ("foo bar" + "hello")))) {
            case "a":
                v2 -= -9223372036854775808
            }
        }
        _ = fi40
        _ = v46
    }
    v2 = v2
    var v47 bool = true
    for fi48 := 0; fi48 < 11; fi48++ {
        var v49 uint = 1
        for fi50 := 0; fi50 < 13; fi50++ {
            fmt.Println(fi48)
            _ = fi50
        }
        _ = fi48
        _ = v49
    }
    v2 -= (v2 & v2)
    v2 = v2
    var v51 uint16 = 0
    for fi52 := 0; fi52 < 11; fi52++ {
        v53 := v3[fi52]
        _ = fi52
        _ = v53
    }
    v2 = v2
    var v54 uint = 921
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v51
    _ = v54
}
