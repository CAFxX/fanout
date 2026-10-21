package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int16; F1 bool }

func main() {
    var v1 int16 = 728
    var v2 float64 = 3.14159
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]int)
    v2 = (((v2 - v2) * float64(1.5)) / ((-2.25 * 3.14159) / v2))
    v2 -= v2
    if false {
        for fi5 := 0; fi5 < 12; fi5++ {
            v6 := []int{(fi5 % (((^(fi5))) | 1)), 127}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi5
            _ = v6
        }
        var v7 int16 = 0
        _ = v7
    } else {
        if ((!(true && (3.14159 > -2.25))) || false) {
            v1 = v1
        } else {
            v8 := GenS1{F0: (v1 & v1)}
            var v9 int16 = v1
            v2 += v2
            _ = v8
            _ = v9
        }
        if ((false && true) && true) {
            v2 = 0.1
        } else {
            fmt.Println(v1, v1)
        }
        v1 = -32768
    }
    v1 += v1
    v10 := GenS1{F0: 1}
    if (((918 >= 2147483647) || (false || false)) || true) {
        if true {
            if (!false) {
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 uint64 = 877
                    v4[""] = fi11
                    _ = fi11
                    _ = v12
                }
            } else {
                v13 := []int{-9223372036854775808}
                _ = v13
            }
            for fi14 := 0; fi14 < 8; fi14++ {
                if ((v1 >> ((2) & 63)) == -1) {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        if (!(v1 >= v1)) {
                            var v16 bool = false
                            switch (fi14 - -68) {
                            case -2147483648:
                                var v17 string = "hello"
                                _ = v17
                            case -1:
                                var v18 int32 = 2147483647
                                _ = v18
                            }
                            _ = v16
                        }
                        _ = fi15
                    }
                    for fi19 := 0; fi19 < 12; fi19++ {
                        var v20 float32 = -2.25
                        v20 += v20
                        _ = fi19
                        _ = v20
                    }
                } else {
                    if (100 == 127) {
                        v21 := v4[("foo bar" + ("\n\t" + "a"))]
                        var v22 uint64 = 0
                        _ = v21
                        _ = v22
                    }
                    var v23 float64 = v2
                    _ = v23
                }
                if false {
                    if (3.14159 == (-2.25 + 1.5)) {
                        fmt.Println(v2)
                        v2 += v2
                    }
                    if true {
                        switch "hello" {
                        case "foo bar":
                            switch fi14 {
                            case 434:
                                var v24 int = (fi14 * fi14)
                                _ = v24
                            case 352:
                                var v25 uint32 = 4294967295
                                _ = v25
                            case 9223372036854775807:
                                v2 = float64(float32(v2))
                            default:
                                v26 := GenS1{F0: -327}
                                _ = v26
                            }
                        }
                        for fi27 := 0; fi27 < 7; fi27++ {
                            var v28 int32 = -29
                            v29 := GenS2{F0: int(v1), F1: fi27}
                            fmt.Println(v1, v28)
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        v1 -= 1
                    }
                    switch "foo bar" {
                    case "foo bar":
                        var v30 string = "a"
                        _ = v30
                    case "":
                        for fi31 := 0; fi31 < 15; fi31++ {
                            v32 := GenS3{F0: v1, F1: v1}
                            v33 := []int{}
                            var v34 int = fi31
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    }
                }
                for fi35 := 0; fi35 < 18; fi35++ {
                    v36 := GenS1{F0: v1}
                    _ = fi35
                    _ = v36
                }
                _ = fi14
            }
            v37 := GenS4{F0: 58, F1: 2147483647}
            _ = v37
        } else {
            if false {
                v38 := []int{553, 127, 881, 760}
                var v39 float32 = 0.0
                _ = v38
                _ = v39
            }
            v40 := GenS2{F0: 75, F1: 42}
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 18; fi41++ {
            var v42 uint = 1
            switch (((("" + "") + "") + "a") + ("" + "hello")) {
            case "":
                delete(v4, "\n\t")
            case "a":
                v2 -= ((v2 - v2) / v2)
            }
            v43 := GenS5{F0: 42, F1: true}
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        v1 = v1
        var v45 uint64 = 399
        var v46 int32 = 194
        _ = fi44
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        var v48 string = ("foo bar" + ("" + "foo bar"))
        v48 = v48
        var v49 int16 = v1
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
