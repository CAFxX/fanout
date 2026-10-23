package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 string }

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = -9223372036854775808
    var v3 int = int(v1)
    v4 := make(map[int]string)
    var v5 int64 = v2
    v6 := []uint32{}
    fmt.Println(v5, v1, v5)
    v2 = 42
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: ("foo bar" + "foo bar"), F1: 0}
        _ = fi7
        _ = v8
    }
    _ = len(v4)
    v5 += v5
    for fi9 := 0; fi9 < 1; fi9++ {
        if false {
            var v10 int16 = 1000
            var v11 bool = (v10 == v10)
            _ = v10
            _ = v11
        }
        if false {
            for fi12 := 0; fi12 < 8; fi12++ {
                _ = len(v4)
                v13 := GenS1{F0: "", F1: 180}
                fmt.Println(fi12, v2, v3)
                _ = fi12
                _ = v13
            }
            if true {
                if (true && false) {
                    v14 := GenS1{F0: "\n\t", F1: 0}
                    fmt.Println(v3, v2, v2)
                    _ = v14
                }
            }
        }
        _ = fi9
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 float32 = (-2.25 + (3.14159 - -2.25))
        _ = fi15
        _ = v16
    }
    if ((-(v1)) > v1) {
        var v17 uint32 = 1
        var v18 bool = false
        if (((v17 & (v17 ^ (v17 << ((0) & 63)))) <= (v17 * v17)) && v18) {
            var v19 float64 = 1.5
            var v20 uint = 4294967295
            _ = v19
            _ = v20
        }
        _ = v17
        _ = v18
    } else {
        v2 -= 42
        for fi21 := 0; fi21 < 16; fi21++ {
            if true {
                fmt.Println(v1, v2, v2)
                if true {
                    var v22 int8 = -1
                    _ = v22
                }
            }
            var v23 int64 = v5
            switch "hello" {
            case "hello":
                if (!false) {
                    switch "" {
                    case "hello":
                        var v24 string = "\n\t"
                        _ = v24
                    case "a":
                        switch (^(v3)) {
                        case -1:
                            fmt.Println(v2, fi21, fi21)
                        case 127:
                            if (!false) {
                                v23 = v5
                                if false {
                                    v5 = -9223372036854775808
                                    for fi25 := 0; fi25 < 17; fi25++ {
                                        var v26 float32 = 0.1
                                        _ = fi25
                                        _ = v26
                                    }
                                    v27 := GenS1{F0: "\n\t", F1: 228}
                                    _ = v27
                                } else {
                                    switch (v3 % ((fi21) | 1)) {
                                    case 625:
                                        for fi28 := 0; fi28 < 18; fi28++ {
                                            v29 := GenS1{F0: "", F1: 42}
                                            var v30 uint64 = 18446744073709551615
                                            if true {
                                                if false {
                                                    var v31 string = "a"
                                                    v32 := GenS2{F0: v30}
                                                    fmt.Println(fi28, v2)
                                                    _ = v31
                                                    _ = v32
                                                }
                                                v23 += 0
                                            } else {
                                                var v33 string = ("foo bar" + ("foo bar" + ""))
                                                _ = v33
                                            }
                                            _ = fi28
                                            _ = v29
                                            _ = v30
                                        }
                                    case 860:
                                        v34 := GenS3{F0: 1000}
                                        _ = v34
                                    case 127:
                                        v35 := GenS4{F0: 128, F1: "\n\t"}
                                        _ = v35
                                    }
                                }
                                if false {
                                    fmt.Println(v3)
                                } else {
                                    v36 := GenS1{F0: "foo bar", F1: 239}
                                    if false {
                                        v3 -= fi21
                                    } else {
                                        for fi37 := 0; fi37 < 7; fi37++ {
                                            switch v3 {
                                            case -1:
                                                fmt.Println(fi21, fi21)
                                            case 1:
                                                fmt.Println(v1, fi37, v2)
                                            }
                                            var v38 int8 = 100
                                            delete(v4, fi21)
                                            _ = fi37
                                            _ = v38
                                        }
                                    }
                                    for fi39 := 0; fi39 < 15; fi39++ {
                                        v2 = v2
                                        v23 += 42
                                        var v40 int8 = 0
                                        v1 -= v1
                                        switch (fi39 & fi21) {
                                        case 124:
                                            if (false && (true || true)) {
                                                v41 := GenS1{F0: "hello", F1: 2}
                                                var v42 string = (("foo bar" + "foo bar") + "hello")
                                                _ = v41
                                                _ = v42
                                            } else {
                                                v3 += v3
                                            }
                                        case 127:
                                            v3 = fi21
                                            var v43 float64 = (0.0 / 1e10)
                                            _ = v43
                                        }
                                        _ = fi39
                                        _ = v40
                                    }
                                    _ = v36
                                }
                            }
                        case 877:
                            var v44 uint = 0
                            _ = v44
                        default:
                            fmt.Println(fi21, v3, v3)
                        }
                    case "foo bar":
                        v45 := GenS3{F0: 319}
                        _ = v45
                    }
                }
            case "a":
                if len(v6) > 0 {
                    v6[0] = 65535
                }
            case "\n\t":
                var v46 uint8 = 2
                _ = v46
            }
            _ = fi21
            _ = v23
        }
        for fi47 := 0; fi47 < 20; fi47++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v48 int32 = (v1 * (-(v1)))
            for fi49 := 0; fi49 < 3; fi49++ {
                var v50 uint64 = 0
                var v51 float32 = 0.1
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = fi47
            _ = v48
        }
    }
    var v52 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
