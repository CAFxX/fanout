package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int8; F1 int32; F2 int16 }
func genFn0(p0 bool, p1 int8) int8 {
    var v7 uint = 564
    var v8 string = ""
    v9 := GenS1{F0: 1, F1: 4294967295}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 float64 = -2.25
    v2 := []string{}
    var v3 uint16 = 42
    var v4 string = "\n\t"
    var v5 bool = true
    v6 := []int{297, 195, 32767}
    v10 := GenS1{F0: 158, F1: 0}
    fmt.Println(v1, v4)
    var v11 bool = v5
    var v12 int = 42
    var v13 int8 = -1
    if ((v13 * 78) != (v13 << ((18446744073709551615) & 63))) {
        v5 = false
        if v11 {
            v14 := GenS2{F0: v13, F1: 34, F2: 986}
            fmt.Println(v5, v11, v12)
            _ = v14
        } else {
            v15 := GenS1{F0: 69, F1: 4294967295}
            var v16 int = (v12 + (v12 / ((-722) | 1)))
            _ = v15
            _ = v16
        }
        v13 += v13
    }
    switch (v12 - -953) {
    case -2147483648:
        switch "" {
        case "":
            fmt.Println(v5, v5)
        case "foo bar":
            var v17 uint = 1
            _ = v17
        case "\n\t":
            if false {
                if (96 < 325) {
                    switch 32768 {
                    case 128:
                        _ = len(v6)
                    case 127:
                        if false {
                            if ((v3 * 0) != v3) {
                                for fi18 := 0; fi18 < 5; fi18++ {
                                    for fi19 := 0; fi19 < 5; fi19++ {
                                        var v20 string = v4
                                        _ = fi19
                                        _ = v20
                                    }
                                    _ = fi18
                                }
                            }
                            switch v12 {
                            case 42:
                                v21 := GenS1{F0: 0, F1: 0}
                                _ = v21
                            default:
                                var v22 uint16 = 32767
                                _ = v22
                            }
                            v1 -= float64((float32(0.0) - -2.25))
                        } else {
                            var v23 uint64 = 397
                            var v24 uint = 1
                            var v25 int = (v12 >> ((782) & 63))
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    default:
                        v1 = v1
                    }
                    v3 = v3
                    var v26 int16 = int16(v12)
                    _ = v26
                }
                _ = len(v2)
                var v27 bool = v11
                _ = v27
            }
        default:
            if (!(127 >= 2)) {
                for fi28 := 0; fi28 < 12; fi28++ {
                    for fi29 := 0; fi29 < 10; fi29++ {
                        v30 := GenS1{F0: 127, F1: 0}
                        var v31 uint8 = 255
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    if true {
                        switch (fi28 >> ((0) & 63)) {
                        case 316:
                            fmt.Println(v1, v1, fi28)
                        case -179:
                            var v32 uint16 = v3
                            _ = v32
                        case 4294967295:
                            v11 = (true || v5)
                        }
                        var v33 string = "foo bar"
                        if true {
                            if (0 <= 148) {
                                v34 := GenS2{F0: 1, F1: 759, F2: 917}
                                _ = v34
                            } else {
                                v4 = v33
                                _ = len(v2)
                                v2 = append(v2, v4)
                            }
                            v35 := GenS2{F0: (v13 * v13), F1: 1, F2: 915}
                            if (v4 > (v4 + "foo bar")) {
                                if (v5 || v11) {
                                    fmt.Println(v12, fi28, v1)
                                }
                            } else {
                                var v36 uint64 = 886
                                var v37 uint = 703
                                _ = v36
                                _ = v37
                            }
                            _ = v35
                        } else {
                            v13 -= v13
                        }
                        _ = v33
                    } else {
                        var v38 uint8 = 0
                        switch "a" {
                        case "a":
                            for fi39 := 0; fi39 < 8; fi39++ {
                                v5 = false
                                switch "" {
                                case "":
                                    switch (fi39 & fi28) {
                                    case 862:
                                        v5 = (v11 || v5)
                                    case 42:
                                        v5 = (!v5)
                                    }
                                default:
                                    v40 := GenS1{F0: (v38 * 206), F1: 415}
                                    _ = v40
                                }
                                fmt.Println(v4, v13, v4)
                                switch fi28 {
                                case -2147483648:
                                    var v41 uint8 = v38
                                    _ = v41
                                default:
                                    v4 = "foo bar"
                                }
                                _ = fi39
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = (v4 + "a")
                        }
                        _ = v38
                    }
                    v13 = genFn0(false, 1)
                    _ = fi28
                }
            } else {
                _ = len(v6)
                v12 = v12
            }
        }
    case 0:
        v3 += v3
    case 280:
        v1 = float64(3.14159)
    }
    v5 = v5
    if len(v6) > 0 {
        v6[0] = v12
    }
    if len(v2) > 0 {
        v2[0] = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
