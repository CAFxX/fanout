package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int16; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 489
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    v3[396] = ((float32(3.14159) < 1.5) || (-1 != 42))
    v4 = (v4 + v4)
    v5 := []int{-823, 1}
    delete(v1, v4)
    switch -2147483648 {
    case 42:
        var v6 int = 9223372036854775807
        _ = v6
    }
    if true {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 uint32 = 65535
            var v9 int16 = 42
            switch v4 {
            case "hello":
                switch (("hello" + "") + v4) {
                case "foo bar":
                    if true {
                        var v10 uint64 = 0
                        _ = v10
                    } else {
                        switch (v4 + v4) {
                        case "a":
                            fmt.Println(v2, v2, fi7)
                        case "hello":
                            v5 = append(v5, 32768)
                        default:
                            var v11 float32 = 0.0
                            _ = v11
                        }
                    }
                }
            case "":
                switch fi7 {
                case 0:
                    fmt.Println(v8, v8)
                case 128:
                    for fi12 := 0; fi12 < 20; fi12++ {
                        switch "\n\t" {
                        case "foo bar":
                            var v13 int8 = 114
                            _ = v13
                        }
                        v14 := GenS1{F0: 128}
                        _ = fi12
                        _ = v14
                    }
                }
            case "a":
                fmt.Println(v2, v9, v8)
            default:
                switch v4 {
                case "":
                    var v15 float64 = 3.14159
                    _ = v15
                case "\n\t":
                    if ("" < v4) {
                        fmt.Println(v8, v8, fi7)
                        v4 = v4
                        if true {
                            var v16 int = fi7
                            if ((554 != 2) || true) {
                                fmt.Println(v4, fi7, v2)
                            } else {
                                switch "" {
                                case "":
                                    for fi17 := 0; fi17 < 13; fi17++ {
                                        v8 -= v8
                                        if true {
                                            switch (((v4 + v4) + v4) + v4) {
                                            case "\n\t":
                                                var v18 int = fi7
                                                _ = v18
                                            }
                                            if false {
                                                if false {
                                                    switch "foo bar" {
                                                    case "":
                                                        var v19 uint8 = 131
                                                        _ = v19
                                                    case "foo bar":
                                                        fmt.Println(v4)
                                                    default:
                                                        if false {
                                                            if len(v5) > 0 {
                                                                v5[0] = (fi7 * (^(v16)))
                                                            }
                                                            v20 := GenS1{F0: 141}
                                                            _ = v20
                                                        } else {
                                                            fmt.Println(v9, v2)
                                                        }
                                                        var v21 int8 = 0
                                                        v5 = append(v5, 32767)
                                                        _ = v21
                                                    }
                                                } else {
                                                    if (42 >= 1) {
                                                        var v22 float32 = -2.25
                                                        var v23 int32 = 287
                                                        _ = v22
                                                        _ = v23
                                                    }
                                                }
                                                if (226 <= 42) {
                                                    var v24 uint = v2
                                                    v25 := GenS2{F0: true, F1: (-(v9)), F2: true}
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v24
                                                    _ = v25
                                                }
                                            } else {
                                                switch ((v4 + ("\n\t" + "hello")) + "\n\t") {
                                                case "a":
                                                    v4 = v4
                                                }
                                                v26 := GenS2{F0: true, F1: (v9 / (((^(v9))) | 1)), F2: true}
                                                _ = v26
                                            }
                                            var v27 int64 = 42
                                            _ = v27
                                        }
                                        _ = fi17
                                    }
                                case "foo bar":
                                    if (!true) {
                                        if true {
                                            var v28 uint64 = 258
                                            _ = v28
                                        } else {
                                            var v29 uint32 = (v8 << ((v2) & 63))
                                            _ = v29
                                        }
                                        v2 = 544
                                    }
                                }
                            }
                            _ = v16
                        } else {
                            v30 := v3[v8]
                            _ = v30
                        }
                    } else {
                        v8 -= 1
                        switch 0 {
                        case 4294967295:
                            if ((v2 - uint(v8)) != (v2 + 2)) {
                                if false {
                                    var v31 int16 = v9
                                    var v32 uint64 = 0
                                    _ = v31
                                    _ = v32
                                } else {
                                    var v33 uint64 = 18446744073709551615
                                    switch "a" {
                                    case "\n\t":
                                        v9 += (v9 + 42)
                                    }
                                    _ = v33
                                }
                                if len(v5) > 0 {
                                    v5[0] = fi7
                                }
                                v2 = v2
                            }
                        default:
                            var v34 int32 = -2147483648
                            _ = v34
                        }
                        v35 := GenS2{F0: ((1 != (v9 + (v9 % (((v9 - -1)) | 1)))) && false), F1: 32767, F2: false}
                        _ = v35
                    }
                case "foo bar":
                    var v36 bool = false
                    _ = v36
                default:
                    var v37 bool = true
                    _ = v37
                }
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        var v38 float32 = ((float32(-2.25) - -2.25) * ((-2.25 - (3.14159 * 0.0)) / 1.0))
        _ = v38
    } else {
        for fi39 := 0; fi39 < 20; fi39++ {
            for fi40 := 0; fi40 < 16; fi40++ {
                var v41 uint64 = 1
                v42 := GenS2{F0: (42 <= 42), F1: -620, F2: (4294967295 <= 402)}
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = fi39
        }
        v2 -= v2
    }
    fmt.Println(v4, v2)
    if (1.5 == 0.0) {
        for fi43 := 0; fi43 < 10; fi43++ {
            fmt.Println(v2, fi43)
            var v44 uint32 = 1
            _ = fi43
            _ = v44
        }
    } else {
        var v45 uint64 = 4294967295
        _ = v45
    }
    v46 := v1["a"]
    for fi47 := 0; fi47 < 20; fi47++ {
        v48 := GenS2{F0: false, F1: 42, F2: v46}
        _ = fi47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
