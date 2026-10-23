package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 int8 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint32 }
func genFn0() int16 {
    var v4 int8 = 0
    if ((42 == -1) || false) {
        v4 = 100
        switch 32767 {
        case 32767:
            v5 := []int{}
            _ = v5
        case 32768:
            var v6 float64 = 3.14159
            _ = v6
        default:
            v4 = v4
        }
    } else {
        switch (("foo bar" + (("" + "a") + "foo bar")) + "foo bar") {
        case "foo bar":
            var v7 int64 = -839
            _ = v7
        default:
            var v8 uint8 = 1
            _ = v8
        }
        v9 := GenS1{F0: 1, F1: true}
        var v10 uint = 2
        _ = v10
        _ = v9
    }
    _ = v4
    return 107
}

func main() {
    var v1 int64 = 1
    var v2 uint16 = 2
    v3 := make(map[string]uint32)
    for fi11 := 0; fi11 < 6; fi11++ {
        if true {
            fmt.Println(fi11)
            var v12 int8 = 110
            _ = v12
        }
        v2 = v2
        if (float64((((float32(0.1) / (0.0 - -2.25)) - ((0.1 / 1.0) / (3.14159 - (0.0 - 0.1)))) - 3.14159)) == (1.5 + 3.14159)) {
            switch "\n\t" {
            case "hello":
                v13 := GenS2{F0: 34, F1: v1, F2: -3}
                _ = v13
            case "\n\t":
                v14 := v3[""]
                _ = v14
            case "foo bar":
                var v15 uint8 = 255
                _ = v15
            }
        } else {
            if (2147483647 != 0) {
                fmt.Println(v1)
                v1 = (v1 / (((v1 ^ (v1 >> ((255) & 63)))) | 1))
                for fi16 := 0; fi16 < 3; fi16++ {
                    v2 -= (^(v2))
                    switch "hello" {
                    case "hello":
                        fmt.Println(v1, fi16)
                    default:
                        switch "" {
                        case "\n\t":
                            v17 := []int{fi16, fi11, fi16}
                            _ = v17
                        case "":
                            v18 := []int{(^(fi16)), fi16}
                            _ = v18
                        case "hello":
                            v1 = v1
                            var v19 uint8 = 26
                            _ = v19
                        }
                    }
                    _ = fi16
                }
            }
        }
        _ = fi11
    }
    if true {
        if (42 != (v1 * v1)) {
            var v20 int8 = 42
            _ = len(v3)
            for fi21 := 0; fi21 < 11; fi21++ {
                for fi22 := 0; fi22 < 18; fi22++ {
                    v2 = (^(v2))
                    _ = fi22
                }
                _ = fi21
            }
            _ = v20
        } else {
            if true {
                var v23 float64 = 0.1
                fmt.Println(v1, v2)
                var v24 uint8 = 1
                _ = v23
                _ = v24
            } else {
                var v25 string = (("" + "foo bar") + ("" + "hello"))
                v1 -= v1
                var v26 float64 = (0.0 + 3.14159)
                _ = v25
                _ = v26
            }
            v2 = (v2 - v2)
            v1 += v1
        }
        switch 0 {
        case 32768:
            delete(v3, (("\n\t" + ("\n\t" + (("hello" + "\n\t") + "a"))) + "a"))
        }
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        v28 := GenS2{F0: 255, F1: (v1 / (((^(v1))) | 1)), F2: -128}
        for fi29 := 0; fi29 < 14; fi29++ {
            if (-2147483648 > 431) {
                v2 += (v2 % ((v2) | 1))
                v2 = 518
                fmt.Println(fi29, fi29, fi29)
                if true {
                    if ("hello" < "") {
                        v1 = v1
                        var v30 string = (("\n\t" + "foo bar") + "")
                        if false {
                            if true {
                                switch 128 {
                                case 32768:
                                    v31 := []int{fi27, fi29}
                                    _ = v31
                                case 42:
                                    var v32 float64 = 0.0
                                    _ = v32
                                case 9223372036854775807:
                                    var v33 int8 = 0
                                    _ = v33
                                }
                            } else {
                                v34 := []int{fi27, 32768}
                                _ = v34
                            }
                            if true {
                                var v35 int16 = 1000
                                fmt.Println(v30, v35, fi29)
                                var v36 uint = 89
                                _ = v35
                                _ = v36
                            } else {
                                if (true || (4294967295 != 243)) {
                                    v37 := GenS3{F0: v1, F1: 814}
                                    v2 -= 1
                                    _ = v37
                                }
                            }
                            var v38 uint16 = v2
                            var v39 bool = ("a" == ((v30 + v30) + (v30 + "")))
                            _ = v38
                            _ = v39
                        }
                        _ = v30
                    }
                    v2 -= (v2 ^ 65535)
                }
                v2 += (v2 & v2)
            }
            v40 := v3[""]
            _ = fi29
            _ = v40
        }
        _ = fi27
        _ = v28
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        for fi42 := 0; fi42 < 19; fi42++ {
            if false {
                v43 := GenS1{F0: 65, F1: true}
                v2 += v2
                _ = v43
            }
            if (45 != uint8(v2)) {
                _ = len(v3)
            }
            _ = fi42
        }
        if (0 != -32768) {
            var v44 float32 = (3.14159 + ((0.0 - 0.1) / 1e10))
            for fi45 := 0; fi45 < 20; fi45++ {
                if true {
                    var v46 float64 = 1e10
                    _ = v46
                }
                switch fi45 {
                case 0:
                    v47 := GenS1{F0: 0, F1: false}
                    _ = v47
                }
                var v48 string = "foo bar"
                _ = fi45
                _ = v48
            }
            _ = v44
        } else {
            v3["hello"] = 844
            v1 = v1
        }
        _ = fi41
    }
    v2 -= v2
    var v49 uint = 1
    fmt.Println(v1, v2)
    v50 := GenS2{F0: 30, F1: (^(v1)), F2: 42}
    for fi51 := 0; fi51 < 6; fi51++ {
        switch (("hello" + "hello") + "hello") {
        case "\n\t":
            v2 += (^(v2))
        default:
            var v52 int64 = v1
            _ = v52
        }
        v53 := GenS4{F0: 2}
        v49 -= 42
        _ = fi51
        _ = v53
    }
    v54 := GenS4{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v54
}
