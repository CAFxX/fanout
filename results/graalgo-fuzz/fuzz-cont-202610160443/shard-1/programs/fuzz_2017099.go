package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }
func genFn0() uint64 {
    var v4 uint32 = 724
    v5 := GenS1{F0: 24}
    var v6 int16 = 32767
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1() uint8 {
    if (1 < -32768) {
        var v7 float64 = -2.25
        if (!(false || (((((-2.25 / 3.14159) - 0.1) / 0.1) * 1e10) >= -2.25))) {
            switch ("\n\t" + "hello") {
            case "foo bar":
                v8 := []int{-9223372036854775808, 996}
                _ = v8
            }
        } else {
            v7 = (v7 / v7)
            var v9 int = -656
            var v10 int32 = -825
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        var v11 int8 = 106
        switch "foo bar" {
        case "foo bar":
            switch ("" + "foo bar") {
            case "":
                var v12 uint8 = 64
                _ = v12
            case "foo bar":
                var v13 uint8 = 127
                _ = v13
            case "a":
                var v14 string = "foo bar"
                _ = v14
            default:
                var v15 float32 = ((3.14159 / -2.25) / -2.25)
                _ = v15
            }
        case "hello":
            v11 -= v11
        case "a":
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 int32 = 937
                v18 := GenS2{F0: 383}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        switch ("a" + "") {
        case "foo bar":
            v19 := []int{9223372036854775807}
            _ = v19
        default:
            v20 := GenS3{F0: 266}
            _ = v20
        }
        _ = v11
    }
    var v21 bool = ((!(34 != 42)) || (-632 < 900))
    v22 := GenS3{F0: 1}
    switch 9223372036854775807 {
    case 9223372036854775807:
        v21 = v21
    case 1:
        var v23 float32 = 0.0
        _ = v23
    default:
        v24 := GenS3{F0: 2147483648}
        _ = v24
    }
    _ = v21
    _ = v22
    return 2
}

func main() {
    v1 := []string{"", "hello"}
    v2 := make(map[string]bool)
    var v3 bool = (2 != 230)
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 7; fi26++ {
            v3 = true
            if false {
                v3 = true
            } else {
                v3 = v3
                var v27 uint = 861
                _ = v27
            }
            _ = fi26
        }
        fmt.Println(fi25, v3)
        _ = fi25
    }
    switch 32768 {
    case 2147483647:
        fmt.Println(v3)
    default:
        if len(v1) > 0 {
            v1[0] = "hello"
        }
    }
    v3 = v3
    if (((-2.25 - 3.14159) - -2.25) >= 0.1) {
        v3 = (0 == -73)
        v28 := GenS2{F0: 1}
        switch 42 {
        case 781:
            for fi29 := 0; fi29 < 11; fi29++ {
                if (!v3) {
                    if false {
                        v3 = true
                        v3 = (4294967295 == 65535)
                    }
                    switch (fi29 * fi29) {
                    case 918:
                        switch "hello" {
                        case "\n\t":
                            var v30 int16 = 266
                            _ = v30
                        case "":
                            v3 = (v3 || (!(32768 < 1)))
                        }
                    default:
                        fmt.Println(fi29, v3)
                    }
                }
                for fi31 := 0; fi31 < 8; fi31++ {
                    v3 = (v3 && (true || (34 != 2)))
                    if v3 {
                        _ = len(v2)
                        var v32 float64 = 3.14159
                        var v33 int8 = 100
                        _ = v32
                        _ = v33
                    } else {
                        if v3 {
                            v3 = (((-2147483648 < -1) && v3) || ((418 <= genFn0()) && true))
                        } else {
                            v3 = v3
                        }
                        var v34 uint8 = 0
                        _ = v34
                    }
                    var v35 uint64 = 2
                    _ = fi31
                    _ = v35
                }
                var v36 float32 = (1e10 + 1.5)
                _ = fi29
                _ = v36
            }
        case -2147483648:
            var v37 uint = 4294967295
            _ = v37
        }
        _ = v28
    }
    var v38 int64 = 1
    v3 = (772 != -1)
    v3 = (!(!false))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 0 {
    case -519:
        var v39 float64 = 3.14159
        _ = v39
    default:
        v3 = (3.14159 > -2.25)
    }
    var v40 float32 = 0.1
    for fi41 := 0; fi41 < 3; fi41++ {
        switch "hello" {
        case "":
            var v42 string = ("\n\t" + "foo bar")
            _ = v42
        case "hello":
            var v43 int = fi41
            _ = v43
        }
        var v44 uint8 = 229
        _ = fi41
        _ = v44
    }
    v45 := GenS2{F0: 0}
    v46 := v2["hello"]
    var v47 uint16 = 2
    v48 := GenS4{F0: 42}
    switch 32768 {
    case 32768:
        var v49 uint = 946
        _ = v49
    case -2147483648:
        v40 = float32(((-2.25 + float64((((1.5 - v40) * float32((3.14159 / (0.1 / -2.25)))) + (v40 - 1.5)))) / -2.25))
    case 219:
        fmt.Println(v46, v38)
    }
    fmt.Println(v3, v46, v46)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v40
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
