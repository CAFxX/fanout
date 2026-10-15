package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint16 }
func genFn0() int64 {
    v5 := GenS1{F0: 2147483647, F1: 42}
    v6 := []int{-680, 32768, 128}
    var v7 uint32 = 2
    if len(v6) > 0 {
        v6[0] = 32767
    }
    var v8 int8 = 1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -455
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 uint64 = 4294967295
    var v4 float64 = (1e10 / 1e10)
    for fi9 := 0; fi9 < 12; fi9++ {
        switch fi9 {
        case 1:
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := []int{fi10}
                _ = fi10
                _ = v11
            }
        case 32768:
            var v12 float32 = 1.5
            _ = v12
        case 0:
            var v13 string = "a"
            _ = v13
        }
        var v14 int8 = 14
        _ = fi9
        _ = v14
    }
    var v15 string = "hello"
    fmt.Println(v4, v4, v4)
    v4 -= float64((-2.25 - (1e10 / -2.25)))
    v16 := []int{9223372036854775807}
    _ = len(v1)
    v2[-2147483648] = "foo bar"
    var v17 string = v15
    fmt.Println(v3)
    if ("\n\t" >= (v15 + v15)) {
        fmt.Println(v4, v15)
        fmt.Println(v3, v4)
        v15 = v15
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := GenS1{F0: 42, F1: fi18}
        if ((-128 >= -119) && false) {
            if true {
                v3 = v3
                var v20 int32 = 345
                _ = v20
            }
            var v21 int8 = 100
            _ = v21
        }
        switch v15 {
        case "\n\t":
            var v22 float32 = (((1.5 * (-2.25 + (1.5 - 1.5))) - float32((0.0 + v4))) * (1.5 / 1.5))
            _ = v22
        case "a":
            switch "a" {
            case "\n\t":
                var v23 int = fi18
                _ = v23
            default:
                v4 -= v4
            }
        case "hello":
            _ = len(v16)
        }
        _ = fi18
        _ = v19
    }
    var v24 string = v15
    v25 := GenS1{F0: -651, F1: -2147483648}
    switch 32768 {
    case -452:
        switch ("\n\t" + (("" + v24) + (v15 + (v17 + ("a" + v17))))) {
        case "\n\t":
            v26 := GenS1{F0: 42, F1: -9223372036854775808}
            _ = v26
        case "a":
            for fi27 := 0; fi27 < 19; fi27++ {
                v28 := GenS2{F0: -50, F1: 100}
                fmt.Println(v4, v24, v3)
                _ = fi27
                _ = v28
            }
        case "hello":
            if true {
                for fi29 := 0; fi29 < 9; fi29++ {
                    v17 = v15
                    _ = fi29
                }
                var v30 uint64 = 2
                _ = v30
            } else {
                switch 127 {
                case -499:
                    var v31 int8 = -128
                    _ = v31
                case 710:
                    var v32 uint16 = 443
                    _ = v32
                case 9223372036854775807:
                    if false {
                        var v33 string = "a"
                        _ = v33
                    } else {
                        for fi34 := 0; fi34 < 3; fi34++ {
                            var v35 bool = true
                            for fi36 := 0; fi36 < 7; fi36++ {
                                switch ((v15 + (("\n\t" + (v15 + ((v17 + ("a" + v15)) + "foo bar"))) + v17)) + "foo bar") {
                                case "a":
                                    _ = len(v1)
                                case "foo bar":
                                    var v37 int32 = -2147483648
                                    _ = v37
                                case "hello":
                                    var v38 uint = 2
                                    _ = v38
                                }
                                if v35 {
                                    var v39 uint64 = v3
                                    v40 := GenS3{F0: (!true), F1: 79, F2: 969}
                                    switch v15 {
                                    case "foo bar":
                                        switch v15 {
                                        case "foo bar":
                                            switch (v17 + "") {
                                            case "":
                                                var v41 uint = 0
                                                _ = v41
                                            default:
                                                fmt.Println(v3, v4, v17)
                                                v3 = 1
                                            }
                                        }
                                    }
                                    _ = v39
                                    _ = v40
                                }
                                _ = fi36
                            }
                            _ = fi34
                            _ = v35
                        }
                        v42 := v1["foo bar"]
                        var v43 int16 = 42
                        _ = v42
                        _ = v43
                    }
                default:
                    if (1.5 != 1.5) {
                        var v44 float32 = (3.14159 / 1.5)
                        _ = v44
                    } else {
                        var v45 int8 = 51
                        var v46 uint16 = 754
                        _ = v45
                        _ = v46
                    }
                }
                switch 2147483647 {
                case -672:
                    switch -304 {
                    case 4294967295:
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    case 9223372036854775807:
                        v47 := GenS3{F0: (-128 <= 91), F1: 968, F2: v3}
                        _ = v47
                    case 806:
                        v48 := GenS4{F0: 42, F1: 1}
                        _ = v48
                    }
                }
                v17 = (v17 + v15)
            }
        default:
            switch 4294967295 {
            case 127:
                for fi49 := 0; fi49 < 12; fi49++ {
                    var v50 uint = 400
                    _ = len(v16)
                    var v51 int16 = 0
                    _ = fi49
                    _ = v50
                    _ = v51
                }
            case -745:
                for fi52 := 0; fi52 < 9; fi52++ {
                    if true {
                        v3 = v3
                        var v53 float64 = v4
                        var v54 string = v17
                        _ = v53
                        _ = v54
                    }
                    v3 -= v3
                    _ = fi52
                }
            default:
                var v55 string = ((v15 + "hello") + v17)
                _ = v55
            }
        }
    default:
        _ = len(v1)
    }
    if true {
        switch 32768 {
        case -2147483648:
            for fi56 := 0; fi56 < 5; fi56++ {
                fmt.Println(v17)
                v4 = (v4 + v4)
                _ = fi56
            }
        default:
            var v57 int8 = 100
            _ = v57
        }
        var v58 int64 = genFn0()
        _ = v58
    }
    switch "hello" {
    case "":
        _ = len(v2)
    default:
        var v59 uint64 = v3
        _ = v59
    }
    var v60 uint32 = 581
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v60
}
