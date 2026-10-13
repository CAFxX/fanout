package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int64; F1 int8; F2 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 string; F2 bool }
func genFn0(p0 uint64, p1 uint32, p2 int16) string {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint = 42
        if (p1 >= p1) {
            fmt.Println(p0, fi7)
            var v9 float32 = -2.25
            fmt.Println(p2, p0, fi7)
            _ = v9
        } else {
            fmt.Println(v8, p0)
            if false {
                v10 := GenS1{F0: fi7}
                fmt.Println(v8, v8)
                _ = v10
            } else {
                var v11 bool = true
                switch ("" + (("a" + "hello") + "")) {
                case "hello":
                    if v11 {
                        for fi12 := 0; fi12 < 17; fi12++ {
                            if v11 {
                                v13 := GenS1{F0: 128}
                                _ = v13
                            } else {
                                fmt.Println(fi7, p1, v8)
                                var v14 int8 = 42
                                _ = v14
                            }
                            v15 := GenS2{F0: v8, F1: (p2 * 1000), F2: (v8 - (v8 ^ v8))}
                            p0 = 42
                            _ = fi12
                            _ = v15
                        }
                    }
                case "\n\t":
                    for fi16 := 0; fi16 < 1; fi16++ {
                        fmt.Println(p2, v8, p2)
                        _ = fi16
                    }
                }
                var v17 int8 = 1
                _ = v11
                _ = v17
            }
            var v18 uint = v8
            _ = v18
        }
        _ = fi7
        _ = v8
    }
    if (-490 != -2147483648) {
        var v19 uint32 = p1
        switch -58 {
        case -1:
            var v20 uint16 = 232
            _ = v20
        case -768:
            if ((p1 - v19) >= (^(p1))) {
                var v21 float32 = 3.14159
                var v22 float64 = -2.25
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 7; fi23++ {
                    switch ("a" + "a") {
                    case "\n\t":
                        switch fi23 {
                        case 127:
                            p1 = p1
                        }
                    case "foo bar":
                        var v24 uint8 = 255
                        _ = v24
                    case "":
                        var v25 int32 = 1
                        _ = v25
                    }
                    v26 := []int{fi23, 32767, 127, fi23}
                    _ = fi23
                    _ = v26
                }
                v27 := GenS3{F0: 255, F1: 0}
                _ = v27
            }
        case 9223372036854775807:
            p2 += (p2 >> ((18446744073709551615) & 63))
        default:
            fmt.Println(p2)
        }
        var v28 string = ("\n\t" + "foo bar")
        _ = v19
        _ = v28
    } else {
        fmt.Println(p1)
        v29 := []int{-1, -467, 9223372036854775807, -237}
        var v30 uint16 = 42
        _ = v29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    var v1 bool = false
    var v2 bool = v1
    var v3 uint16 = 32767
    var v4 int64 = 9223372036854775807
    v5 := make(map[string]string)
    var v6 string = "hello"
    var v31 uint32 = 2147483648
    if (1.5 > (-2.25 + 1.5)) {
        if (v4 > 0) {
            v32 := GenS3{F0: v3, F1: 144}
            if v2 {
                v33 := GenS4{F0: (v4 & v4), F1: 127, F2: 42}
                _ = v33
            }
            _ = v32
        } else {
            for fi34 := 0; fi34 < 10; fi34++ {
                delete(v5, "\n\t")
                v35 := GenS2{F0: 254, F1: 0, F2: 4294967295}
                v31 += 631
                _ = fi34
                _ = v35
            }
        }
    }
    v36 := v5[(("foo bar" + v6) + "foo bar")]
    v37 := GenS4{F0: (-(v4)), F1: -128, F2: 42}
    v38 := GenS4{F0: 0, F1: -43, F2: 35}
    if v2 {
        for fi39 := 0; fi39 < 15; fi39++ {
            v6 = v36
            if v1 {
                for fi40 := 0; fi40 < 4; fi40++ {
                    if v1 {
                        v41 := GenS1{F0: fi40}
                        var v42 int = -1
                        _ = v41
                        _ = v42
                    }
                    delete(v5, (("hello" + "\n\t") + ((((v36 + "\n\t") + "\n\t") + (("" + "hello") + (("hello" + ("foo bar" + (v36 + (v36 + ((v36 + v36) + v6))))) + (v6 + ("a" + "\n\t"))))) + ("a" + ("\n\t" + ("" + ""))))))
                    var v43 int = 249
                    _ = fi40
                    _ = v43
                }
            } else {
                switch fi39 {
                case 9223372036854775807:
                    v36 = v36
                }
                v3 = 2
            }
            v36 = v6
            _ = fi39
        }
        v44 := []int{0, int(v3), 0}
        if (!(false && (0 <= 32767))) {
            v45 := GenS4{F0: v4, F1: 100, F2: 859}
            _ = v45
        }
        _ = v44
    } else {
        var v46 int32 = -947
        _ = v46
    }
    var v47 int64 = (^(v4))
    v4 = v47
    v48 := v5[(("" + v6) + v6)]
    v4 = (v4 >> ((0) & 63))
    v1 = true
    var v49 uint64 = 2
    if v2 {
        switch "a" {
        case "hello":
            var v50 int32 = 2147483647
            _ = v50
        case "a":
            for fi51 := 0; fi51 < 17; fi51++ {
                v48 = (v36 + v48)
                _ = fi51
            }
        case "foo bar":
            var v52 float64 = 3.14159
            _ = v52
        default:
            v5[v36] = "foo bar"
        }
    } else {
        for fi53 := 0; fi53 < 2; fi53++ {
            v49 += (v49 & v49)
            var v54 uint64 = v49
            _ = fi53
            _ = v54
        }
        v49 -= v49
        if (!v2) {
            v55 := GenS5{F0: -41}
            v2 = v1
            for fi56 := 0; fi56 < 13; fi56++ {
                v31 -= (v31 / ((v31) | 1))
                delete(v5, v36)
                _ = fi56
            }
            _ = v55
        }
    }
    var v57 float32 = float32(-2.25)
    switch v36 {
    case "\n\t":
        v5[""] = "\n\t"
    case "hello":
        var v58 uint8 = 0
        _ = v58
    }
    v59 := GenS6{F0: 0, F1: (v36 + v6), F2: v2}
    v47 += (v4 * (v4 + v47))
    var v60 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v57
    _ = v59
    _ = v6
    _ = v60
}
