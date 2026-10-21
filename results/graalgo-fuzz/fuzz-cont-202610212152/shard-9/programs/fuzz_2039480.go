package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint64; F2 string }
type GenS6 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS7 struct { F0 string; F1 string }
func genFn0(p0 uint8) string {
    var v5 uint = 0
    v6 := []int{0, -9223372036854775808, 9223372036854775807}
    _ = p0
    _ = v5
    _ = v6
    return ""
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 1
    var v3 int32 = -1
    v4 := make(map[uint32]string)
    v3 += 0
    for fi7 := 0; fi7 < 13; fi7++ {
        if (42 > -6) {
            var v8 uint32 = 464
            v9 := GenS1{F0: v3, F1: v2}
            v8 = v2
            _ = v8
            _ = v9
        } else {
            v10 := GenS2{F0: 0}
            switch 9223372036854775807 {
            case 729:
                delete(v4, v2)
            default:
                delete(v4, (v2 / ((v2) | 1)))
            }
            if (-9223372036854775808 <= 32768) {
                v11 := []int{}
                v12 := GenS3{F0: (v2 - (v2 | 65535)), F1: (v3 + -1), F2: 30}
                var v13 int16 = 1000
                _ = v11
                _ = v12
                _ = v13
            } else {
                var v14 uint64 = 4294967295
                v15 := GenS2{F0: 0}
                var v16 int32 = v3
                _ = v14
                _ = v15
                _ = v16
            }
            _ = v10
        }
        _ = fi7
    }
    var v17 uint64 = 18446744073709551615
    if false {
        switch ("\n\t" + "a") {
        case "\n\t":
            var v18 uint64 = v17
            _ = v18
        case "hello":
            var v19 int16 = 1000
            _ = v19
        default:
            v1[""] = (v2 ^ (v2 | (v2 / ((v2) | 1))))
        }
        v17 = (v17 - (^(v17)))
        if false {
            v20 := GenS1{F0: (v3 ^ 1), F1: 65535}
            _ = v20
        } else {
            v21 := GenS2{F0: 249}
            if false {
                v2 = v2
            }
            _ = v21
        }
    } else {
        v22 := []int{}
        _ = v22
    }
    var v23 int64 = 42
    v24 := GenS2{F0: v23}
    var v25 float32 = 0.1
    if true {
        for fi26 := 0; fi26 < 11; fi26++ {
            for fi27 := 0; fi27 < 5; fi27++ {
                if true {
                    v17 += v17
                    v28 := GenS1{F0: v3, F1: v2}
                    if true {
                        var v29 int32 = 115
                        v30 := GenS1{F0: 2147483647, F1: (v2 ^ 22)}
                        v31 := GenS1{F0: v29, F1: (v2 ^ 744)}
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        var v32 int32 = v3
                        v33 := GenS4{F0: 42}
                        _ = v32
                        _ = v33
                    }
                    v2 -= v2
                    _ = v28
                } else {
                    switch fi27 {
                    case 42:
                        switch 789 {
                        case 215:
                            var v34 int64 = v23
                            _ = v34
                        case -2147483648:
                            var v35 float64 = -2.25
                            _ = v35
                        case 127:
                            v36 := GenS4{F0: v23}
                            _ = v36
                        }
                    default:
                        if true {
                            v23 += v23
                            var v37 float64 = 0.1
                            v38 := GenS5{F0: (fi27 ^ -2147483648), F1: v17, F2: (((("" + "\n\t") + ("foo bar" + "foo bar")) + "\n\t") + "foo bar")}
                            _ = v37
                            _ = v38
                        } else {
                            v39 := []int{fi26, int(v2), fi27, (fi26 & fi27)}
                            v40 := GenS6{F0: (v17 + (v17 % ((v17) | 1))), F1: 128, F2: 32767}
                            fmt.Println(v25)
                            _ = v39
                            _ = v40
                        }
                    }
                    if (!(255 <= 127)) {
                        v23 = v23
                    }
                }
                v17 = v17
                _ = fi27
            }
            for fi41 := 0; fi41 < 19; fi41++ {
                v4[4294967295] = "hello"
                switch "foo bar" {
                case "\n\t":
                    if false {
                        var v42 uint32 = 311
                        switch ("\n\t" + "\n\t") {
                        case "hello":
                            v25 -= 0.0
                        case "":
                            var v43 uint16 = 0
                            _ = v43
                        case "a":
                            if true {
                                v23 -= (v23 & v23)
                                v25 += (0.0 / (1.5 + v25))
                                v44 := GenS3{F0: 4294967295, F1: v3, F2: 0}
                                _ = v44
                            }
                        default:
                            v17 -= 805
                        }
                        _ = v42
                    }
                case "":
                    v23 += v23
                }
                if ((v17 + v17) >= v17) {
                    if ((float64((1e10 + 3.14159)) + 0.0) <= 3.14159) {
                        var v45 uint32 = (v2 ^ 2147483648)
                        _ = v45
                    }
                    fmt.Println(v25)
                    v2 += 537
                    fmt.Println(fi26, v3, v25)
                }
                _ = fi41
            }
            v2 += (^(v2))
            _ = fi26
        }
        v46 := v1["foo bar"]
        _ = v46
    } else {
        delete(v1, "foo bar")
    }
    v3 += -2147483648
    v47 := GenS7{F0: "hello", F1: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v47
}
