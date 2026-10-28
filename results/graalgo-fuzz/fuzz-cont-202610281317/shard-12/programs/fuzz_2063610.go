package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int }

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 0
    var v3 int32 = 42
    var v4 bool = (false || true)
    var v5 uint8 = 111
    v6 := []int{-439, 525, 32768}
    if (true && v4) {
        v7 := GenS1{F0: 529, F1: 0}
        _ = v7
    } else {
        var v8 int32 = -2147483648
        switch "a" {
        case "foo bar":
            var v9 int64 = -9223372036854775808
            _ = v9
        case "hello":
            switch "\n\t" {
            case "":
                switch (("hello" + "foo bar") + ("a" + "foo bar")) {
                case "hello":
                    for fi10 := 0; fi10 < 14; fi10++ {
                        for fi11 := 0; fi11 < 3; fi11++ {
                            var v12 int = 9223372036854775807
                            v13 := GenS2{F0: 1}
                            var v14 int16 = -1
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        var v15 string = "a"
                        _ = fi10
                        _ = v15
                    }
                case "foo bar":
                    if (!(v2 == v2)) {
                        if len(v6) > 0 {
                            v6[0] = 4294967295
                        }
                        var v16 uint64 = 42
                        for fi17 := 0; fi17 < 17; fi17++ {
                            if v4 {
                                v16 -= (v16 - 1)
                                var v18 int = (fi17 % ((fi17) | 1))
                                _ = v18
                            } else {
                                var v19 int32 = (v3 & 50)
                                _ = v19
                            }
                            _ = fi17
                        }
                        _ = v16
                    } else {
                        var v20 uint64 = 278
                        v21 := GenS2{F0: uint(v2)}
                        _ = len(v1)
                        _ = v20
                        _ = v21
                    }
                case "":
                    v4 = true
                default:
                    v4 = (!v4)
                }
            case "hello":
                fmt.Println(v2, v3, v5)
            case "a":
                v4 = (!v4)
            }
        }
        _ = v8
    }
    var v22 int8 = 42
    if true {
        v23 := GenS3{F0: 0, F1: 128}
        if (1e10 <= 1e10) {
            fmt.Println(v4, v22, v3)
            switch 128 {
            case -2147483648:
                if len(v6) > 0 {
                    v6[0] = 32768
                }
            case 42:
                v2 = v2
            }
            switch "\n\t" {
            case "":
                var v24 int32 = v3
                _ = v24
            case "foo bar":
                var v25 float32 = ((1e10 / 1.0) + (0.1 + (0.1 * 1e10)))
                _ = v25
            default:
                if false {
                    v1["hello"] = (("\n\t" + "a") == "a")
                } else {
                    var v26 bool = false
                    _ = len(v1)
                    _ = v26
                }
            }
        }
        var v27 bool = ("hello" >= "hello")
        _ = v23
        _ = v27
    } else {
        var v28 uint64 = 918
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 float64 = 1e10
            _ = fi29
            _ = v30
        }
        if true {
            switch "foo bar" {
            case "a":
                if (!v4) {
                    if false {
                        v2 += (v2 | v2)
                        var v31 uint64 = v28
                        _ = v31
                    }
                    v3 -= 42
                } else {
                    var v32 int32 = 1
                    _ = v32
                }
            case "\n\t":
                var v33 int32 = (v3 << ((65535) & 63))
                _ = v33
            case "hello":
                switch 32768 {
                case 32768:
                    v5 += (v5 << ((2) & 63))
                default:
                    v22 = 79
                }
            }
        }
        _ = v28
    }
    var v34 int8 = v22
    v22 -= -27
    switch "" {
    case "":
        var v35 uint16 = v2
        _ = v35
    default:
        switch 616 {
        case 4294967295:
            v36 := GenS1{F0: 255, F1: 32767}
            _ = v36
        }
    }
    var v37 float64 = 1e10
    v38 := GenS1{F0: 282, F1: 0}
    var v39 int64 = 1
    var v40 int64 = 42
    fmt.Println(v37, v22)
    v22 += (v34 / ((v34) | 1))
    v41 := GenS1{F0: 18446744073709551615, F1: 2147483647}
    v6 = append(v6, 4294967295)
    if (!(!v4)) {
        if v4 {
            var v42 string = "hello"
            fmt.Println(v3, v3)
            switch 4294967295 {
            case 32767:
                var v43 int = -1
                _ = v43
            case 944:
                var v44 uint32 = 2147483647
                _ = v44
            case -9223372036854775808:
                v34 += (v22 ^ v34)
            }
            _ = v42
        } else {
            switch -2147483648 {
            case -2147483648:
                if (v4 && true) {
                    fmt.Println(v34, v5, v3)
                }
            default:
                if (!v4) {
                    for fi45 := 0; fi45 < 12; fi45++ {
                        fmt.Println(v3, v34)
                        _ = fi45
                    }
                }
            }
            if ((((42 < 32767) && v4) && (uint64(v2) == 1)) && v4) {
                if v4 {
                    var v46 uint64 = 1
                    v46 = 2
                    _ = v46
                } else {
                    v22 += (v34 + v34)
                    var v47 uint = 255
                    _ = v47
                }
            }
        }
        switch 42 {
        case 32767:
            fmt.Println(v3, v5)
        case -747:
            v34 = (v22 * v22)
        case 127:
            v2 += (v2 % ((32767) | 1))
        }
    } else {
        if false {
            v22 -= 1
            var v48 float32 = 0.1
            delete(v1, (("a" + "\n\t") + "\n\t"))
            _ = v48
        }
    }
    v49 := v1[("" + ("hello" + ("\n\t" + "")))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v49
    _ = v5
    _ = v6
}
