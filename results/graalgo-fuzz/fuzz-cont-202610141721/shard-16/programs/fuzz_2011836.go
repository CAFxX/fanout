package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint; F2 int8 }
type GenS6 struct { F0 int16 }
func genFn0() int64 {
    var v6 int64 = -975
    switch "foo bar" {
    case "foo bar":
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 uint64 = 836
                var v10 uint64 = 18446744073709551615
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
    case "a":
        switch "" {
        case "":
            var v11 int8 = 1
            _ = v11
        }
    default:
        var v12 string = ("a" + "hello")
        _ = v12
    }
    var v13 string = ("" + "")
    var v14 float64 = 1.5
    _ = v13
    _ = v14
    _ = v6
    return (v6 >> ((751) & 63))
}

func main() {
    var v1 uint16 = 611
    var v2 int16 = -541
    v3 := make(map[string]bool)
    var v4 int16 = (v2 >> ((255) & 63))
    var v5 int = 32768
    if ((float64(0.1) > ((3.14159 / -2.25) - -2.25)) && true) {
        if true {
            var v15 uint32 = 2147483648
            _ = v15
        } else {
            var v16 uint = 65535
            var v17 uint = 0
            _ = v16
            _ = v17
        }
        v4 += v2
        if true {
            var v18 int16 = 0
            _ = v18
        }
    }
    if true {
        if true {
            var v19 float64 = 0.1
            _ = v19
        } else {
            v1 = v1
            var v20 bool = false
            _ = v20
        }
    } else {
        var v21 float64 = 3.14159
        _ = v21
    }
    var v22 uint8 = 1
    for fi23 := 0; fi23 < 3; fi23++ {
        v2 = (v4 + (v4 % ((1000) | 1)))
        _ = fi23
    }
    _ = len(v3)
    var v24 uint32 = 65535
    fmt.Println(v24, v24, v5)
    if true {
        v22 += v22
    }
    switch v5 {
    case 128:
        var v25 float64 = 0.1
        _ = v25
    case 9223372036854775807:
        if ((1e10 + -2.25) > -2.25) {
            switch "a" {
            case "foo bar":
                v26 := GenS1{F0: 4294967295}
                _ = v26
            default:
                v24 = v24
            }
        } else {
            v27 := []int{v5}
            if true {
                if true {
                    var v28 string = "a"
                    _ = v28
                }
            } else {
                v22 = v22
            }
            _ = v27
        }
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v3[""] = true
    case "a":
        v5 = v5
    default:
        fmt.Println(v4)
    }
    v29 := v3["hello"]
    fmt.Println(v24, v2, v24)
    if (42 == 2) {
        var v30 int64 = -9223372036854775808
        fmt.Println(v5)
        var v31 bool = (4294967295 > 2)
        _ = v30
        _ = v31
    }
    var v32 float64 = float64(0.0)
    v32 = 1.5
    if (v29 || v29) {
        fmt.Println(v1, v2)
        for fi33 := 0; fi33 < 5; fi33++ {
            fmt.Println(v22, v5)
            var v34 uint32 = (^(v24))
            v22 -= 1
            _ = fi33
            _ = v34
        }
        var v35 float32 = 1e10
        _ = v35
    } else {
        if v29 {
            v36 := GenS1{F0: 448}
            _ = v36
        } else {
            v4 -= v4
            switch v5 {
            case -2147483648:
                v37 := GenS2{F0: 42, F1: 755, F2: 452}
                _ = v37
            default:
                if ((v5 - (v5 + (v5 >> ((4294967295) & 63)))) == 9223372036854775807) {
                    var v38 uint = uint(v1)
                    _ = v38
                } else {
                    switch "a" {
                    case "hello":
                        if ((("\n\t" < "\n\t") || v29) || v29) {
                            var v39 float64 = v32
                            if (v1 != v1) {
                                v4 = (v4 % ((-59) | 1))
                                var v40 uint64 = 814
                                for fi41 := 0; fi41 < 6; fi41++ {
                                    if true {
                                        v42 := GenS1{F0: 2}
                                        _ = v42
                                    }
                                    for fi43 := 0; fi43 < 11; fi43++ {
                                        if v29 {
                                            v32 = float64(((-2.25 / 0.1) - float32(1e10)))
                                            v24 -= v24
                                        }
                                        v44 := GenS1{F0: 0}
                                        fmt.Println(v39, v4, v39)
                                        _ = fi43
                                        _ = v44
                                    }
                                    _ = fi41
                                }
                                _ = v40
                            } else {
                                v45 := GenS3{F0: (v22 ^ v22)}
                                fmt.Println(v24, v39)
                                v46 := GenS4{F0: 892}
                                _ = v45
                                _ = v46
                            }
                            v47 := GenS5{F0: -1, F1: 4294967295, F2: 0}
                            _ = v39
                            _ = v47
                        }
                    case "":
                        fmt.Println(v4, v24)
                    case "a":
                        delete(v3, ("a" + "a"))
                    }
                    v48 := GenS6{F0: v4}
                    _ = v48
                }
            }
        }
        var v49 uint = 0
        var v50 uint = 65535
        _ = v49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
