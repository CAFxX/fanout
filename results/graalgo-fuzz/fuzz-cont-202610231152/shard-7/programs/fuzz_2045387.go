package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int64; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 bool }
func genFn0(p0 uint32) bool {
    p0 = p0
    var v7 int64 = -1
    v7 = v7
    v7 += v7
    v8 := GenS1{F0: ""}
    _ = p0
    _ = v7
    _ = v8
    return (0 <= 2147483647)
}
func genFn1() uint32 {
    var v9 uint32 = 1
    for fi10 := 0; fi10 < 4; fi10++ {
        switch (fi10 << ((449) & 63)) {
        case 2147483647:
            v11 := GenS2{F0: (v9 >> ((1) & 63)), F1: -438, F2: 42}
            _ = v11
        case 128:
            fmt.Println(v9, v9)
        case 32768:
            v9 -= v9
        default:
            var v12 uint64 = 4294967295
            _ = v12
        }
        _ = fi10
    }
    fmt.Println(v9)
    _ = v9
    return v9
}

func main() {
    var v1 int64 = 295
    var v2 uint64 = 346
    var v3 uint64 = 18446744073709551615
    var v4 int = 2147483647
    v5 := make(map[string]int)
    var v6 int = v4
    if ((-1 <= (v6 % ((-1) | 1))) || false) {
        var v13 uint16 = 0
        var v14 float32 = float32((((3.14159 + (1.5 / 0.1)) / 1.5) + 3.14159))
        _ = v13
        _ = v14
    }
    if (!true) {
        fmt.Println(v6, v4)
        for fi15 := 0; fi15 < 19; fi15++ {
            v6 = (fi15 | v6)
            v16 := []int{742, fi15, v6, (v6 | fi15)}
            for fi17 := 0; fi17 < 14; fi17++ {
                if (v3 <= v2) {
                    var v18 uint32 = 727
                    _ = v18
                }
                fmt.Println(v6, v4, v2)
                if (-716 > 42) {
                    var v19 int32 = 42
                    if (!((true || true) && true)) {
                        var v20 int8 = 69
                        v16 = append(v16, v6)
                        _ = v20
                    } else {
                        var v21 uint64 = (v2 * (v2 & (v2 * (v2 - 413))))
                        switch (fi17 * (v4 + fi15)) {
                        case 32767:
                            delete(v5, ("hello" + ("" + "hello")))
                        case -893:
                            delete(v5, "a")
                        case -9223372036854775808:
                            fmt.Println(v1, v3, v19)
                            fmt.Println(fi17, v21, fi17)
                            if true {
                                fmt.Println(v4, v19)
                            } else {
                                if (false && (false || (100 == -1))) {
                                    var v22 uint32 = 4294967295
                                    v6 = v6
                                    v19 = v19
                                    _ = v22
                                }
                                v23 := GenS3{F0: 479, F1: 255, F2: 248}
                                _ = v23
                            }
                        }
                        fmt.Println(fi15)
                        _ = v21
                    }
                    _ = v19
                }
                if true {
                    if len(v16) > 0 {
                        v16[0] = fi17
                    }
                    var v24 uint64 = v2
                    if true {
                        var v25 uint32 = 1
                        fmt.Println(fi15)
                        v26 := GenS4{F0: 127}
                        _ = v25
                        _ = v26
                    }
                    switch (fi17 + v4) {
                    case 185:
                        v6 += (fi17 ^ 1)
                    case 4294967295:
                        v24 = 964
                    }
                    _ = v24
                }
                var v27 uint = 65535
                _ = fi17
                _ = v27
            }
            _ = fi15
            _ = v16
        }
    } else {
        switch (v6 | v6) {
        case -2147483648:
            var v28 int64 = (^(v1))
            _ = v28
        case -684:
            switch ("foo bar" + ("foo bar" + "a")) {
            case "\n\t":
                for fi29 := 0; fi29 < 14; fi29++ {
                    for fi30 := 0; fi30 < 18; fi30++ {
                        v31 := GenS5{F0: v4, F1: (!true)}
                        v3 -= 18446744073709551615
                        _ = fi30
                        _ = v31
                    }
                    fmt.Println(v6, v2)
                    _ = fi29
                }
            default:
                v5["\n\t"] = v4
            }
        case -439:
            fmt.Println(v2)
        default:
            v32 := GenS4{F0: 128}
            _ = v32
        }
    }
    var v33 int = (-(v6))
    for fi34 := 0; fi34 < 6; fi34++ {
        fmt.Println(v2, v3)
        v35 := []int{(v33 ^ v6)}
        var v36 int8 = -67
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println(v2)
    v6 = (-(v4))
    v3 -= (v2 % (((v3 & v3)) | 1))
    switch "\n\t" {
    case "foo bar":
        v1 = v1
    }
    v37 := GenS4{F0: 2}
    fmt.Println(v6)
    var v38 string = "a"
    fmt.Println(v33)
    v39 := GenS2{F0: 153, F1: (v1 >> ((2) & 63)), F2: -251}
    switch v38 {
    case "":
        v40 := v5[""]
        _ = v40
    case "foo bar":
        var v41 uint16 = 42
        _ = v41
    case "\n\t":
        v42 := GenS1{F0: ""}
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
