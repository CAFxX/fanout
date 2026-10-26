package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int; F2 int8 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 uint; F2 int }
type GenS8 struct { F0 int }
func genFn0(p0 uint) uint8 {
    fmt.Println(p0)
    v7 := GenS1{F0: false, F1: 0}
    p0 = p0
    _ = p0
    _ = v7
    return 2
}

func main() {
    v1 := []string{"", ("hello" + (("foo bar" + (("" + ("\n\t" + "")) + "\n\t")) + "foo bar")), "a", ("" + "foo bar")}
    v2 := make(map[string]uint32)
    var v3 uint16 = 792
    var v4 uint = 1
    var v5 float64 = 3.14159
    var v6 bool = true
    for fi8 := 0; fi8 < 17; fi8++ {
        switch ("\n\t" + "a") {
        case "a":
            var v9 uint64 = 18446744073709551615
            _ = v9
        }
        if (true || v6) {
            var v10 int = (fi8 * (-(fi8)))
            for fi11 := 0; fi11 < 8; fi11++ {
                fmt.Println(v6, v6)
                var v12 float64 = v5
                v13 := GenS2{F0: 202, F1: 831}
                if ((fi8 / ((fi8) | 1)) >= v10) {
                    switch "\n\t" {
                    case "hello":
                        v6 = (2147483647 <= -2147483648)
                    case "\n\t":
                        var v14 uint8 = 1
                        _ = v14
                    case "a":
                        var v15 uint32 = 2147483648
                        _ = v15
                    }
                    v2[""] = 65535
                    fmt.Println(v10, fi11)
                }
                var v16 uint64 = 1
                _ = fi11
                _ = v12
                _ = v13
                _ = v16
            }
            v17 := GenS3{F0: v3}
            _ = v10
            _ = v17
        }
        var v18 uint32 = 2147483648
        _ = fi8
        _ = v18
    }
    var v19 int16 = -873
    switch "hello" {
    case "foo bar":
        v20 := GenS1{F0: (!v6), F1: (v3 | (^(v3)))}
        _ = v20
    case "a":
        var v21 uint64 = 1
        _ = v21
    default:
        if v6 {
            v22 := GenS4{F0: v6, F1: 127, F2: 42}
            for fi23 := 0; fi23 < 12; fi23++ {
                if (v5 < (1e10 + float64(0.0))) {
                    var v24 int64 = -9223372036854775808
                    v25 := GenS5{F0: "", F1: 184}
                    _ = len(v1)
                    _ = v24
                    _ = v25
                }
                v26 := GenS2{F0: 53, F1: 510}
                v3 = (v3 | (^(v3)))
                _ = fi23
                _ = v26
            }
            switch -9223372036854775808 {
            case 9223372036854775807:
                for fi27 := 0; fi27 < 7; fi27++ {
                    if ((42 >= 868) || (!(v6 && v6))) {
                        if true {
                            var v28 uint8 = 2
                            v29 := GenS3{F0: v3}
                            _ = v28
                            _ = v29
                        } else {
                            switch -422 {
                            case 175:
                                var v30 uint8 = 0
                                _ = v30
                            default:
                                switch "" {
                                case "\n\t":
                                    if false {
                                        v4 = (v4 & 2)
                                        v19 = (v19 ^ 1)
                                    } else {
                                        if ((729 >= 322) && (-1 < -1)) {
                                            if (fi27 > (fi27 - fi27)) {
                                                v31 := GenS6{F0: (v3 / ((v3) | 1))}
                                                _ = v31
                                            } else {
                                                var v32 int = fi27
                                                v5 -= v5
                                                fmt.Println(v32, fi27)
                                                _ = v32
                                            }
                                            v5 = (1.5 + v5)
                                            delete(v2, "hello")
                                        } else {
                                            var v33 bool = v6
                                            v33 = (1.5 > 1e10)
                                            var v34 uint8 = 1
                                            _ = v33
                                            _ = v34
                                        }
                                        if (uint(v3) <= (v4 - v4)) {
                                            var v35 uint8 = 1
                                            _ = v35
                                        }
                                    }
                                case "":
                                    v2[("" + "a")] = 2
                                }
                            }
                        }
                        v36 := GenS3{F0: v3}
                        v19 = 42
                        _ = v36
                    }
                    _ = fi27
                }
            }
            _ = v22
        }
    }
    fmt.Println(v3, v5)
    for fi37 := 0; fi37 < 8; fi37++ {
        if (!(1 <= -1)) {
            v38 := v2[("\n\t" + ("a" + ("foo bar" + "a")))]
            fmt.Println(v5, v3)
            if ((-560 == 1) || v6) {
                var v39 int64 = 0
                var v40 float32 = 1e10
                _ = v39
                _ = v40
            }
            _ = v38
        } else {
            var v41 uint8 = 1
            var v42 uint = (^(v4))
            for fi43 := 0; fi43 < 18; fi43++ {
                var v44 string = ""
                v6 = v6
                switch (fi37 / (((fi43 / (((^(fi37))) | 1))) | 1)) {
                case -554:
                    var v45 uint8 = (v41 << (((v42 >> (((v42 | v4)) & 63))) & 63))
                    _ = v45
                default:
                    if (((fi43 % (((^(fi37))) | 1)) < (-(fi37))) && v6) {
                        v5 = (v5 * (((1.5 * v5) + float64((3.14159 + (0.1 - 1e10)))) + v5))
                        v46 := GenS6{F0: 2}
                        _ = v46
                    }
                }
                v4 = (v42 & 0)
                _ = fi43
                _ = v44
            }
            _ = v41
            _ = v42
        }
        _ = fi37
    }
    if true {
        var v47 uint16 = (^(v3))
        var v48 bool = (-2147483648 > 127)
        _ = v47
        _ = v48
    } else {
        if (!false) {
            var v49 uint32 = 334
            if v6 {
                v19 -= v19
                v4 = (v4 % ((v4) | 1))
                delete(v2, "foo bar")
            }
            _ = v49
        } else {
            v19 += (^(v19))
        }
        for fi50 := 0; fi50 < 15; fi50++ {
            v51 := GenS4{F0: v6, F1: fi50, F2: 127}
            for fi52 := 0; fi52 < 20; fi52++ {
                v4 += v4
                var v53 int32 = int32(v3)
                _ = fi52
                _ = v53
            }
            _ = fi50
            _ = v51
        }
    }
    var v54 int8 = -6
    v55 := GenS2{F0: 255, F1: 42}
    delete(v2, "a")
    switch 70 {
    case 9223372036854775807:
        if v6 {
            var v56 int8 = v54
            v4 += (v4 + 18446744073709551615)
            v57 := GenS5{F0: "\n\t", F1: 42}
            _ = v56
            _ = v57
        } else {
            v54 += v54
            v54 = 127
            v3 = 42
        }
    case 42:
        fmt.Println(v3, v19)
    }
    if (v3 <= 0) {
        v58 := GenS7{F0: -1, F1: 2, F2: -962}
        _ = v58
    } else {
        var v59 float64 = 3.14159
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if ((!true) || v6) {
            v60 := GenS8{F0: 42}
            _ = v60
        } else {
            fmt.Println(v3, v4)
            for fi61 := 0; fi61 < 19; fi61++ {
                var v62 bool = true
                _ = fi61
                _ = v62
            }
            var v63 uint32 = 42
            _ = v63
        }
        _ = v59
    }
    var v64 int16 = (v19 - 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
    _ = v64
}
