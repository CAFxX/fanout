package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 bool = false
    var v2 int16 = -1
    var v3 uint16 = 2
    fmt.Println(v3, v3, v3)
    var v4 int = 2147483647
    switch (^(v4)) {
    case 42:
        var v5 uint32 = 42
        _ = v5
    case 2147483647:
        if v1 {
            var v6 uint32 = 65535
            for fi7 := 0; fi7 < 9; fi7++ {
                fmt.Println(v1, v6, v1)
                _ = fi7
            }
            v8 := []int{v4, (-(v4)), 4294967295}
            _ = v6
            _ = v8
        }
    case 127:
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := GenS1{F0: 42, F1: 2, F2: 1}
            _ = fi9
            _ = v10
        }
    }
    if (0 <= 100) {
        var v11 bool = v1
        _ = v11
    }
    var v12 uint16 = 32767
    var v13 float64 = 0.0
    switch "foo bar" {
    case "hello":
        var v14 uint8 = 76
        _ = v14
    case "foo bar":
        v15 := []int{v4, (v4 + v4), 0}
        _ = v15
    case "":
        v16 := GenS1{F0: 9223372036854775807, F1: 255, F2: 42}
        _ = v16
    default:
        switch "foo bar" {
        case "":
            var v17 int8 = 42
            _ = v17
        case "\n\t":
            v18 := GenS2{F0: 4294967295, F1: (v3 & (v3 / ((118) | 1)))}
            _ = v18
        default:
            v19 := GenS1{F0: 61, F1: 65535, F2: 127}
            _ = v19
        }
    }
    v20 := []int{v4, 9223372036854775807}
    for fi21 := 0; fi21 < 19; fi21++ {
        v12 += v12
        v22 := GenS1{F0: -647, F1: 881, F2: 74}
        for fi23 := 0; fi23 < 14; fi23++ {
            _ = len(v20)
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 float64 = float64(1e10)
        switch "hello" {
        case "\n\t":
            var v26 int8 = 109
            _ = v26
        default:
            if false {
                fmt.Println(v4, v3, v4)
            }
        }
        if (((!(!(!(2 > 4294967295)))) && (127 >= 0)) || true) {
            v27 := GenS1{F0: -1, F1: 2, F2: -128}
            switch v4 {
            case -85:
                if v1 {
                    _ = len(v20)
                    var v28 int16 = (v2 + (-(v2)))
                    _ = v28
                }
            case 1:
                var v29 uint = 1
                _ = v29
            }
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    var v30 uint32 = 1
    fmt.Println(v2)
    v1 = (v1 && (v1 && (true || (-2147483648 >= -94))))
    switch v4 {
    case 12:
        fmt.Println(v3)
    case 127:
        var v31 int16 = v2
        _ = v31
    }
    if true {
        if v1 {
            if v1 {
                var v32 uint = 42
                _ = v32
            }
            switch (v4 / (((v4 + -2147483648)) | 1)) {
            case 42:
                v33 := GenS3{F0: 42}
                _ = v33
            }
        } else {
            v12 -= 65535
            if ((v12 <= (v3 - 582)) || (v30 > v30)) {
                var v34 int16 = v2
                for fi35 := 0; fi35 < 11; fi35++ {
                    v12 = v12
                    var v36 int16 = (v34 << ((18446744073709551615) & 63))
                    _ = fi35
                    _ = v36
                }
                _ = v34
            } else {
                if ((v2 >> ((1) & 63)) == v2) {
                    var v37 int32 = 130
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 3; fi38++ {
                        if len(v20) > 0 {
                            v20[0] = v4
                        }
                        _ = fi38
                    }
                }
            }
        }
        _ = len(v20)
        for fi39 := 0; fi39 < 17; fi39++ {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            var v40 int16 = -353
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = fi39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
}
