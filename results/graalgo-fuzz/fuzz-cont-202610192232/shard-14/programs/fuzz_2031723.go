package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 string = "a"
    var v2 uint = 4294967295
    v3 := make(map[string]int)
    v4 := make(map[string]int)
    v5 := []uint32{810, 2147483647}
    switch 0 {
    case 731:
        var v6 uint32 = 2
        _ = v6
    case 42:
        var v7 float64 = ((((float64(((float32(float64(-2.25)) + 0.0) / -2.25)) + -2.25) * 0.1) / 1e10) - 0.1)
        _ = v7
    case 128:
        if len(v5) > 0 {
            v5[0] = 1
        }
    }
    if true {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint64 = 4294967295
            fmt.Println(fi8, fi8, v2)
            if true {
                if true {
                    _ = len(v5)
                    fmt.Println(v2, v1, v9)
                }
                v5 = append(v5, 510)
            }
            _ = fi8
            _ = v9
        }
        v4["a"] = -2147483648
        fmt.Println(v1)
    } else {
        if false {
            v10 := v3[v1]
            for fi11 := 0; fi11 < 11; fi11++ {
                switch (v10 - v10) {
                case 9223372036854775807:
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v10 = (fi11 & -1)
                        v1 = ""
                        var v13 uint16 = 42
                        _ = fi12
                        _ = v13
                    }
                }
                switch v1 {
                case "a":
                    var v14 float32 = -2.25
                    _ = v14
                default:
                    var v15 uint8 = 2
                    _ = v15
                }
                _ = fi11
            }
            _ = v10
        }
        fmt.Println(v2, v2)
    }
    fmt.Println(v1)
    v1 = v1
    var v16 int = 856
    if len(v5) > 0 {
        v5[0] = 494
    }
    var v17 uint32 = 2147483648
    for fi18 := 0; fi18 < 14; fi18++ {
        for fi19 := 0; fi19 < 14; fi19++ {
            if ((v2 | v2) != (v2 & 0)) {
                v1 = ("hello" + v1)
                var v20 uint = v2
                var v21 int32 = -144
                if (!true) {
                    if (fi19 <= 982) {
                        switch fi19 {
                        case 32767:
                            var v22 uint16 = 856
                            _ = v22
                        case -1:
                            v21 -= (v21 & v21)
                        case -2147483648:
                            v23 := GenS1{F0: 122}
                            var v24 int16 = -369
                            v17 -= v17
                            _ = v23
                            _ = v24
                        default:
                            var v25 uint32 = 294
                            _ = v25
                        }
                        fmt.Println(v16, v1)
                        var v26 int16 = 0
                        v27 := GenS1{F0: 42}
                        _ = v26
                        _ = v27
                    } else {
                        if false {
                            var v28 int64 = -1
                            _ = v28
                        }
                        v1 = v1
                    }
                    if false {
                        var v29 int64 = 916
                        if true {
                            switch ((v1 + "") + (v1 + v1)) {
                            case "foo bar":
                                var v30 string = ("hello" + v1)
                                _ = v30
                            case "":
                                if (((-110 > -58) || (false && (v21 == 925))) && false) {
                                    v3["a"] = (^(v16))
                                    var v31 int64 = v29
                                    _ = v31
                                } else {
                                    if true {
                                        v1 = ""
                                    } else {
                                        v32 := GenS2{F0: -4}
                                        var v33 int16 = 1000
                                        _ = v32
                                        _ = v33
                                    }
                                    v1 = "hello"
                                    v21 = v21
                                }
                                var v34 uint32 = 65535
                                _ = v34
                            case "hello":
                                var v35 uint64 = 582
                                _ = v35
                            }
                        }
                        _ = v29
                    } else {
                        v2 = (v20 - 972)
                        if false {
                            var v36 uint64 = 18446744073709551615
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v36
                        } else {
                            var v37 float64 = 1e10
                            _ = v37
                        }
                    }
                    _ = len(v3)
                }
                _ = v20
                _ = v21
            }
            _ = fi19
        }
        switch (^(v16)) {
        case 128:
            if len(v5) > 0 {
                v5[0] = v17
            }
        }
        _ = fi18
    }
    v38 := GenS1{F0: 42}
    for fi39 := 0; fi39 < 20; fi39++ {
        var v40 bool = (!false)
        _ = fi39
        _ = v40
    }
    v3[(v1 + (v1 + "hello"))] = v16
    v16 += (v16 ^ -2147483648)
    switch v1 {
    case "a":
        for fi41 := 0; fi41 < 20; fi41++ {
            var v42 uint = (v2 / ((42) | 1))
            _ = fi41
            _ = v42
        }
    case "foo bar":
        var v43 int64 = 42
        _ = v43
    case "\n\t":
        for fi44 := 0; fi44 < 14; fi44++ {
            var v45 uint32 = (^(v17))
            var v46 uint8 = 2
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    var v47 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v47
    _ = v5
}
