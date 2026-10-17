package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint; F1 int16; F2 int64 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32 }
func genFn0() int8 {
    var v4 int16 = -1
    if true {
        switch (("foo bar" + "") + "foo bar") {
        case "hello":
            v5 := GenS1{F0: -2147483648, F1: 2}
            _ = v5
        case "\n\t":
            var v6 int16 = v4
            _ = v6
        case "foo bar":
            v7 := GenS1{F0: 127, F1: 255}
            _ = v7
        default:
            v8 := []int{32768, 42, 0}
            _ = v8
        }
        v9 := []int{32767, 1}
        _ = v9
    } else {
        var v10 int32 = -2147483648
        switch 0 {
        case 825:
            var v11 uint = 552
            _ = v11
        case 9223372036854775807:
            v12 := []int{-1, -1, 4294967295}
            _ = v12
        case -2147483648:
            v10 = v10
        default:
            fmt.Println(v4, v4)
        }
        _ = v10
    }
    _ = v4
    return 0
}
func genFn1(p0 string) int16 {
    p0 = "a"
    switch "foo bar" {
    case "foo bar":
        for fi13 := 0; fi13 < 13; fi13++ {
            if (2 < 2) {
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 int64 = -1
                    _ = fi14
                    _ = v15
                }
            }
            _ = fi13
        }
    default:
        fmt.Println(p0)
    }
    fmt.Println(p0)
    _ = p0
    return 1000
}

func main() {
    var v1 bool = true
    var v2 int = 128
    var v3 bool = v1
    if v3 {
        for fi16 := 0; fi16 < 8; fi16++ {
            if true {
                switch (fi16 / ((fi16) | 1)) {
                case -2147483648:
                    v2 = (fi16 & (fi16 + 337))
                case -810:
                    var v17 int64 = 0
                    _ = v17
                case 0:
                    fmt.Println(v1, v3)
                }
            } else {
                v2 += (-(fi16))
                if ((!(42 <= 18446744073709551615)) || (!(387 <= 4294967295))) {
                    fmt.Println(v1, v3, v3)
                    if v1 {
                        v18 := GenS2{F0: 535, F1: 584, F2: -9223372036854775808}
                        _ = v18
                    } else {
                        v3 = ((v1 || (2 <= 42)) || false)
                        v19 := GenS3{F0: -9223372036854775808, F1: 1}
                        _ = v19
                    }
                    v20 := []int{fi16, 2147483647}
                    _ = v20
                } else {
                    switch "" {
                    case "":
                        v21 := GenS2{F0: 0, F1: -731, F2: -829}
                        _ = v21
                    case "a":
                        var v22 uint = 255
                        _ = v22
                    case "hello":
                        v23 := []int{661, v2, 32768, 4294967295}
                        _ = v23
                    }
                    fmt.Println(fi16, fi16)
                    v3 = v3
                }
                var v24 uint8 = 127
                _ = v24
            }
            var v25 uint32 = 226
            v26 := GenS1{F0: v2, F1: 127}
            _ = fi16
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 17; fi27++ {
            for fi28 := 0; fi28 < 17; fi28++ {
                if ((!(99 == -9223372036854775808)) || v3) {
                    v29 := GenS4{F0: "", F1: ((fi27 % (((v2 | v2)) | 1)) >= fi28)}
                    switch (fi27 & (v2 & 32767)) {
                    case 127:
                        v1 = (-292 > 9223372036854775807)
                    }
                    _ = v29
                }
                var v30 uint16 = 65535
                if v3 {
                    var v31 int8 = genFn0()
                    _ = v31
                } else {
                    v32 := GenS2{F0: 0, F1: -32768, F2: 0}
                    var v33 uint = 1
                    _ = v32
                    _ = v33
                }
                _ = fi28
                _ = v30
            }
            v34 := GenS5{F0: 255}
            _ = fi27
            _ = v34
        }
    } else {
        v35 := GenS3{F0: -1, F1: 42}
        v36 := GenS2{F0: 18446744073709551615, F1: 19, F2: -1}
        v1 = v3
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        for fi38 := 0; fi38 < 2; fi38++ {
            if v1 {
                var v39 int = fi38
                switch -2147483648 {
                case 32768:
                    var v40 string = "hello"
                    _ = v40
                case 4294967295:
                    v2 -= 1
                default:
                    v3 = false
                }
                _ = v39
            } else {
                var v41 int = 128
                v42 := GenS3{F0: -1, F1: 32767}
                if (!(!false)) {
                    switch "hello" {
                    case "foo bar":
                        v43 := GenS1{F0: (v41 & 32767), F1: 127}
                        _ = v43
                    default:
                        var v44 uint64 = 0
                        _ = v44
                    }
                    var v45 uint64 = 2
                    v1 = false
                    _ = v45
                }
                _ = v41
                _ = v42
            }
            switch "\n\t" {
            case "hello":
                if (-33 > 1000) {
                    v46 := []int{fi38, 9223372036854775807}
                    _ = v46
                }
            default:
                var v47 float64 = 1e10
                _ = v47
            }
            _ = fi38
        }
        _ = fi37
    }
    if ((971 == 42) || v3) {
        v1 = (!true)
        v48 := []int{(v2 & 127), 435}
        _ = v48
    } else {
        if false {
            if (2 >= 1) {
                v3 = (1.5 >= (-2.25 / 1.0))
            }
            var v49 int8 = 42
            _ = v49
        }
    }
    v1 = true
    var v50 int32 = -194
    fmt.Println(v1)
    v2 = int(v50)
    var v51 int = (v2 & (v2 + (v2 | -565)))
    v52 := GenS6{F0: 4294967295}
    var v53 float64 = (((1.5 * 1e10) - ((1e10 - -2.25) * -2.25)) - 1e10)
    switch v2 {
    case 591:
        var v54 int16 = 1
        _ = v54
    case 32767:
        var v55 uint16 = 42
        _ = v55
    case 127:
        var v56 bool = false
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
