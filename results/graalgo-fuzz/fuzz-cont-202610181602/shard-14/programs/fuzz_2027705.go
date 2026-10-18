package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint, p1 bool) int8 {
    p0 -= (p0 ^ p0)
    p0 = p0
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []uint32{838, 42, 726, 2147483647}
    v2 := make(map[string]uint32)
    var v3 string = ("\n\t" + ("foo bar" + ""))
    var v4 int8 = 0
    v4 += v4
    var v5 int8 = genFn0(65535, ((1e10 - 0.1) < -2.25))
    if true {
        switch ("\n\t" + ("a" + v3)) {
        case "\n\t":
            v6 := GenS1{F0: 0}
            _ = v6
        default:
            fmt.Println(v5, v5)
        }
        if len(v1) > 0 {
            v1[0] = 830
        }
        var v7 uint32 = 0
        _ = v7
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            switch -331 {
            case -605:
                _ = len(v1)
            default:
                v4 = (v4 | (v4 >> ((447) & 63)))
            }
        }
    }
    if false {
        for fi8 := 0; fi8 < 20; fi8++ {
            v4 += v5
            v4 -= v4
            _ = fi8
        }
    } else {
        var v9 uint32 = 2147483647
        switch 1 {
        case 1:
            var v10 int16 = 0
            _ = v10
        case -1:
            switch -361 {
            case 4294967295:
                for fi11 := 0; fi11 < 18; fi11++ {
                    fmt.Println(v3, v3, v5)
                    if true {
                        v3 = v3
                        v9 += v9
                        v12 := v2["foo bar"]
                        _ = v12
                    }
                    _ = fi11
                }
            case 128:
                if false {
                    var v13 int16 = 32767
                    switch v3 {
                    case "hello":
                        switch v3 {
                        case "":
                            var v14 bool = ((((!(0 <= uint8(v9))) && false) && (v5 < 57)) || true)
                            _ = v14
                        }
                    }
                    if (1 <= 451) {
                        switch ("" + "a") {
                        case "a":
                            v5 -= (v5 & 42)
                        case "":
                            switch v3 {
                            case "hello":
                                for fi15 := 0; fi15 < 7; fi15++ {
                                    if false {
                                        var v16 int16 = -32768
                                        _ = v16
                                    }
                                    if true {
                                        v9 -= v9
                                    }
                                    _ = fi15
                                }
                            }
                        case "foo bar":
                            v1 = append(v1, v9)
                        default:
                            var v17 uint8 = 127
                            _ = v17
                        }
                    }
                    _ = v13
                } else {
                    var v18 uint8 = 60
                    _ = v18
                }
            case 741:
                v19 := GenS1{F0: 0}
                _ = v19
            }
        }
        _ = v9
    }
    if true {
        switch -563 {
        case -224:
            if (759 < 2) {
                v20 := GenS1{F0: 1}
                var v21 int = -197
                _ = v20
                _ = v21
            }
        }
        v22 := GenS1{F0: 1}
        v1 = append(v1, 1)
        _ = v22
    } else {
        if ((("" + v3) > v3) || (0 <= 1000)) {
            var v23 uint = 633
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 bool = (1 >= 159)
                v23 -= v23
                _ = fi24
                _ = v25
            }
            _ = v23
        } else {
            for fi26 := 0; fi26 < 7; fi26++ {
                if (fi26 == fi26) {
                    var v27 uint64 = 896
                    var v28 uint32 = 2147483647
                    if (!(0.0 <= float32((1e10 - (((1e10 + 3.14159) + 0.0) + 0.0))))) {
                        v5 = v5
                        for fi29 := 0; fi29 < 16; fi29++ {
                            v5 += genFn0(991, false)
                            var v30 uint64 = v27
                            v4 = 26
                            if false {
                                v31 := GenS2{F0: 42}
                                var v32 int16 = 1000
                                v27 = v30
                                v4 = (^(v5))
                                var v33 uint32 = v28
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            _ = fi29
                            _ = v30
                        }
                        var v34 uint8 = 229
                        _ = v34
                    }
                    _ = v27
                    _ = v28
                }
                var v35 uint16 = 255
                _ = fi26
                _ = v35
            }
            var v36 bool = false
            _ = v36
        }
    }
    var v37 float64 = 0.1
    var v38 float32 = (1e10 + ((1e10 / -2.25) * 0.1))
    v4 = v4
    fmt.Println(v37)
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 uint64 = 0
        v37 = (3.14159 - 1e10)
        for fi41 := 0; fi41 < 12; fi41++ {
            delete(v2, (v3 + (v3 + "")))
            var v42 string = ("\n\t" + "\n\t")
            _ = fi41
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
