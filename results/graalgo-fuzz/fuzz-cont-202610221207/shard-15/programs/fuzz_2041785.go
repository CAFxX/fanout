package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int8; F2 int32 }

func main() {
    var v1 int = -2147483648
    var v2 int8 = 127
    var v3 string = "foo bar"
    var v4 uint8 = 255
    var v5 float64 = 3.14159
    v6 := make(map[string]uint32)
    switch v3 {
    case "hello":
        v7 := GenS1{F0: v3}
        _ = v7
    case "\n\t":
        switch v1 {
        case 127:
            fmt.Println(v2)
        case -1:
            fmt.Println(v3, v3)
        default:
            if true {
                v4 = v4
                v3 = "a"
            }
        }
    }
    v8 := GenS1{F0: "foo bar"}
    v9 := GenS2{F0: 1000, F1: 65535}
    v6[""] = 0
    _ = len(v6)
    _ = len(v6)
    fmt.Println(v5, v3)
    v10 := GenS3{F0: -1}
    var v11 uint = 18446744073709551615
    for fi12 := 0; fi12 < 1; fi12++ {
        if (v3 != (v3 + v3)) {
            var v13 int = 127
            if false {
                var v14 int32 = -2147483648
                _ = v14
            } else {
                v15 := []int{-1, v13, (v13 * v1), 9223372036854775807}
                switch (fi12 + (v13 | (v1 >> (((^(v11))) & 63)))) {
                case 127:
                    fmt.Println(v3, fi12, v11)
                }
                _ = v15
            }
            var v16 uint32 = 240
            _ = v13
            _ = v16
        } else {
            v5 -= 1e10
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 int64 = 0
                v19 := v6[v3]
                var v20 uint16 = 255
                var v21 int8 = v2
                var v22 float64 = 1e10
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
                _ = v21
                _ = v22
            }
            switch (v3 + "a") {
            case "foo bar":
                for fi23 := 0; fi23 < 20; fi23++ {
                    if (false || true) {
                        var v24 int32 = -2147483648
                        _ = v24
                    }
                    fmt.Println(fi23, v3)
                    var v25 uint8 = v4
                    _ = fi23
                    _ = v25
                }
            case "a":
                if (-76 < 1) {
                    var v26 string = (v3 + v3)
                    v27 := GenS2{F0: -32768, F1: 2147483648}
                    _ = v26
                    _ = v27
                } else {
                    var v28 float64 = (v5 + ((-2.25 + 1.5) / v5))
                    v28 = (3.14159 - v28)
                    switch v3 {
                    case "\n\t":
                        fmt.Println(v2)
                    case "":
                        var v29 bool = (42 < 1)
                        _ = v29
                    case "foo bar":
                        var v30 uint8 = (^(v4))
                        _ = v30
                    default:
                        v31 := GenS1{F0: (v3 + ("" + "\n\t"))}
                        _ = v31
                    }
                    _ = v28
                }
            }
        }
        v11 = v11
        if false {
            v32 := GenS4{F0: -1, F1: (v2 / (((v2 - v2)) | 1)), F2: -1}
            v33 := []int{fi12, (v1 & fi12)}
            _ = v32
            _ = v33
        } else {
            for fi34 := 0; fi34 < 1; fi34++ {
                v35 := GenS4{F0: -1, F1: v2, F2: -2147483648}
                var v36 string = ((v3 + v3) + ((v3 + ("" + v3)) + ((v3 + "a") + "\n\t")))
                _ = fi34
                _ = v35
                _ = v36
            }
            var v37 int = v1
            _ = v37
        }
        _ = fi12
    }
    v11 -= (v11 - v11)
    var v38 uint64 = 943
    for fi39 := 0; fi39 < 6; fi39++ {
        var v40 bool = false
        var v41 int = v1
        _ = fi39
        _ = v40
        _ = v41
    }
    switch v3 {
    case "foo bar":
        v38 -= 858
    case "hello":
        var v42 int64 = 42
        _ = v42
    default:
        switch (v3 + v3) {
        case "hello":
            v43 := GenS4{F0: 9223372036854775807, F1: v2, F2: -1}
            _ = v43
        case "foo bar":
            v5 -= 3.14159
        case "a":
            switch -592 {
            case -9223372036854775808:
                fmt.Println(v5)
            default:
                if true {
                    if (v3 > "a") {
                        var v44 float32 = 0.0
                        switch v3 {
                        case "foo bar":
                            for fi45 := 0; fi45 < 20; fi45++ {
                                var v46 int16 = 1000
                                var v47 int = fi45
                                _ = fi45
                                _ = v46
                                _ = v47
                            }
                        case "\n\t":
                            v48 := GenS2{F0: -32768, F1: 4294967295}
                            _ = v48
                        }
                        _ = v44
                    }
                } else {
                    v1 = (v1 | int(v2))
                }
            }
        }
    }
    for fi49 := 0; fi49 < 3; fi49++ {
        v50 := []int{(v1 % (((fi49 * v1)) | 1)), int(v2)}
        for fi51 := 0; fi51 < 1; fi51++ {
            v2 = (v2 / ((v2) | 1))
            if (!(true || (((v5 * 0.1) * 3.14159) >= (v5 - -2.25)))) {
                var v52 uint32 = 65535
                _ = v52
            }
            v1 -= v1
            _ = fi51
        }
        _ = fi49
        _ = v50
    }
    var v53 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v8
    _ = v9
}
