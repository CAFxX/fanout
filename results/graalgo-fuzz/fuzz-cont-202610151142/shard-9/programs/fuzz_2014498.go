package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int8 = 10
    var v2 float64 = 1.5
    var v3 uint64 = 980
    var v4 float64 = v2
    v5 := make(map[string]string)
    if ((1.5 + -2.25) >= (3.14159 * 1e10)) {
        switch 1 {
        case -9223372036854775808:
            v4 = v2
        case 42:
            var v6 uint8 = 0
            _ = v6
        case 32767:
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 int = (-(fi7))
                _ = fi7
                _ = v8
            }
        }
        switch "foo bar" {
        case "hello":
            if (!true) {
                if (2 == 123) {
                    fmt.Println(v4, v4, v1)
                    var v9 int32 = 0
                    var v10 int8 = v1
                    _ = v10
                    _ = v9
                }
            }
        default:
            v11 := []int{-9223372036854775808, 1, 127}
            _ = v11
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            v3 = v3
            v2 = 1e10
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 string = "foo bar"
                v14 = (v14 + v14)
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    fmt.Println(v3, v2, v1)
    v5[""] = "\n\t"
    fmt.Println(v3, v4, v3)
    if false {
        if (true || (!false)) {
            for fi15 := 0; fi15 < 7; fi15++ {
                for fi16 := 0; fi16 < 12; fi16++ {
                    switch "\n\t" {
                    case "":
                        v17 := []int{}
                        _ = v17
                    default:
                        v1 = (v1 & (v1 - -20))
                    }
                    v18 := []int{(fi16 + (fi16 * (fi15 >> ((4294967295) & 63)))), fi16, (fi16 | fi16), (-(fi15))}
                    _ = fi16
                    _ = v18
                }
                _ = fi15
            }
        } else {
            v19 := GenS1{F0: 1}
            v2 = 0.1
            _ = v19
        }
    }
    var v20 int16 = -1
    if (-2147483648 == -2147483648) {
        var v21 int64 = 42
        _ = v21
    }
    var v22 int64 = 1
    var v23 string = "hello"
    var v24 uint32 = 510
    for fi25 := 0; fi25 < 11; fi25++ {
        v26 := GenS2{F0: -1}
        switch v23 {
        case "":
            var v27 int64 = (v22 + v22)
            _ = v27
        case "hello":
            v28 := GenS2{F0: -1}
            _ = v28
        case "\n\t":
            fmt.Println(fi25)
        default:
            for fi29 := 0; fi29 < 8; fi29++ {
                if false {
                    delete(v5, ("\n\t" + ""))
                    fmt.Println(v3, v20, v2)
                    v22 = (-(v22))
                }
                v24 -= (v24 + v24)
                _ = fi29
            }
        }
        _ = fi25
        _ = v26
    }
    var v30 bool = (1.5 != 3.14159)
    v23 = (v23 + (v23 + "hello"))
    for fi31 := 0; fi31 < 1; fi31++ {
        for fi32 := 0; fi32 < 19; fi32++ {
            v33 := GenS1{F0: 505}
            var v34 int8 = v1
            _ = fi32
            _ = v33
            _ = v34
        }
        for fi35 := 0; fi35 < 13; fi35++ {
            switch (v23 + ("\n\t" + "foo bar")) {
            case "":
                v23 = "\n\t"
            }
            _ = fi35
        }
        fmt.Println(v20, v22, v22)
        _ = fi31
    }
    if ((v30 || ((((v24 >> ((42) & 63)) > (v24 >> ((258) & 63))) || v30) && false)) && (v30 || (true && true))) {
        fmt.Println(v22)
        var v36 uint16 = 65535
        switch v23 {
        case "":
            switch v23 {
            case "hello":
                if (v24 == 1) {
                    v37 := GenS3{F0: (v24 | (^(v24)))}
                    v5["foo bar"] = v23
                    v20 += v20
                    _ = v37
                } else {
                    v38 := []int{128, 4294967295, 4294967295}
                    var v39 int8 = (v1 >> ((891) & 63))
                    _ = v38
                    _ = v39
                }
            case "foo bar":
                v40 := GenS3{F0: (v24 + (v24 * v24))}
                _ = v40
            case "":
                var v41 int8 = (^(v1))
                _ = v41
            default:
                v42 := v5[v23]
                _ = v42
            }
        case "a":
            v36 = v36
        case "\n\t":
            var v43 int = -2147483648
            _ = v43
        }
        _ = v36
    } else {
        v44 := v5[v23]
        var v45 uint32 = v24
        _ = v44
        _ = v45
    }
    if (!false) {
        v46 := []int{-9223372036854775808, -25, 0, -849}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
