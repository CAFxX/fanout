package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 9223372036854775807
    var v3 uint8 = 42
    switch (v2 % (((v2 ^ (v2 | v2))) | 1)) {
    case 128:
        v2 = 9223372036854775807
    case 42:
        switch v2 {
        case -301:
            var v4 uint64 = 18446744073709551615
            _ = v4
        case 0:
            var v5 float64 = 1.5
            _ = v5
        case 19:
            var v6 bool = (false || true)
            _ = v6
        default:
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := GenS1{F0: v2}
                v9 := GenS1{F0: v2}
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    case 127:
        var v10 int = (-(v2))
        _ = v10
    default:
        if true {
            switch "a" {
            case "foo bar":
                if ((-2.25 - ((-2.25 / -2.25) + 0.0)) < 0.1) {
                    for fi11 := 0; fi11 < 18; fi11++ {
                        var v12 int8 = 0
                        var v13 uint32 = 50
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS2{F0: -523}
                    v15 := []int{767, 9223372036854775807}
                    _ = v14
                    _ = v15
                }
            case "hello":
                v2 += (v2 << ((4294967295) & 63))
            default:
                if false {
                    switch "\n\t" {
                    case "\n\t":
                        v16 := GenS3{F0: "\n\t"}
                        _ = v16
                    case "a":
                        fmt.Println(v3, v2)
                    }
                    if false {
                        v1[-1] = false
                    } else {
                        switch v2 {
                        case 9223372036854775807:
                            for fi17 := 0; fi17 < 17; fi17++ {
                                switch fi17 {
                                case 32767:
                                    switch (v2 ^ (v2 / ((-2147483648) | 1))) {
                                    case 42:
                                        switch ("hello" + "") {
                                        case "a":
                                            var v18 int = (-(fi17))
                                            _ = v18
                                        case "\n\t":
                                            v19 := []int{}
                                            _ = v19
                                        default:
                                            var v20 bool = false
                                            _ = v20
                                        }
                                    case -904:
                                        v21 := GenS3{F0: "hello"}
                                        _ = v21
                                    case 1:
                                        fmt.Println(fi17, fi17)
                                    }
                                }
                                if (42 < 0) {
                                    var v22 uint16 = 565
                                    var v23 int = v2
                                    _ = v22
                                    _ = v23
                                }
                                var v24 uint32 = 2
                                _ = fi17
                                _ = v24
                            }
                        case 32768:
                            v25 := []int{}
                            _ = v25
                        case -718:
                            var v26 int8 = 0
                            _ = v26
                        default:
                            if false {
                                if true {
                                    fmt.Println(v3)
                                    v2 += 128
                                    var v27 uint64 = 1
                                    _ = v27
                                }
                                if true {
                                    for fi28 := 0; fi28 < 4; fi28++ {
                                        var v29 uint64 = 2
                                        _ = len(v1)
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                            }
                        }
                        for fi30 := 0; fi30 < 19; fi30++ {
                            v31 := GenS3{F0: "\n\t"}
                            _ = fi30
                            _ = v31
                        }
                        var v32 int = 32767
                        _ = v32
                    }
                }
            }
            if (9223372036854775807 < 1) {
                var v33 int8 = -47
                _ = v33
            } else {
                fmt.Println(v3)
                switch v2 {
                case -9223372036854775808:
                    if false {
                        if ((true && ((false || true) && true)) || true) {
                            v34 := GenS1{F0: -1}
                            v35 := GenS1{F0: (v2 | 42)}
                            _ = v34
                            _ = v35
                        } else {
                            switch v2 {
                            case 127:
                                for fi36 := 0; fi36 < 20; fi36++ {
                                    var v37 float64 = 1e10
                                    _ = fi36
                                    _ = v37
                                }
                            case 291:
                                v2 += v2
                            default:
                                if true {
                                    v2 = (v2 % ((v2) | 1))
                                    v3 = v3
                                } else {
                                    v38 := GenS1{F0: (^(v2))}
                                    _ = v38
                                }
                            }
                        }
                        var v39 int64 = 42
                        for fi40 := 0; fi40 < 17; fi40++ {
                            var v41 bool = ((v39 * (v39 >> ((2) & 63))) < -685)
                            _ = fi40
                            _ = v41
                        }
                        _ = v39
                    }
                case 128:
                    v2 = -776
                case -976:
                    switch "hello" {
                    case "\n\t":
                        v3 = (v3 % (((v3 | (v3 >> ((593) & 63)))) | 1))
                    case "foo bar":
                        var v42 string = "hello"
                        _ = v42
                    }
                }
                var v43 uint32 = 42
                _ = v43
            }
            v44 := []int{v2, v2, (v2 & v2), (v2 << ((0) & 63))}
            _ = v44
        }
    }
    v3 += (v3 + v3)
    v45 := v1[(v2 ^ 430)]
    for fi46 := 0; fi46 < 19; fi46++ {
        v47 := []int{v2, fi46, (fi46 ^ 9223372036854775807)}
        _ = fi46
        _ = v47
    }
    fmt.Println(v3, v3)
    var v48 uint32 = 1
    v49 := v1[v2]
    v50 := []int{}
    if (255 == 1) {
        v51 := GenS4{F0: 574, F1: -9223372036854775808}
        if v49 {
            v52 := GenS3{F0: "hello"}
            v49 = (v49 && (65535 == 299))
            _ = v52
        } else {
            var v53 string = ("hello" + (("" + "\n\t") + (("" + "\n\t") + "hello")))
            _ = v53
        }
        var v54 bool = false
        _ = v51
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v48
    _ = v49
    _ = v50
}
